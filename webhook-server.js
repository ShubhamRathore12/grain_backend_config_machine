// webhook-server.js - Lightweight webhook listener for GitHub auto-deploy
// Runs on port 9001 to avoid conflict with grain-backend webhook on 9000

const http = require('http');
const crypto = require('crypto');
const { execSync } = require('child_process');

const PORT = 9001;
const SECRET = 'machine-config-deploy-secret-2026';
const DEPLOY_SCRIPT = '/opt/machine-config/deploy.sh';

function verifySignature(payload, signature) {
    if (!signature) return false;
    const hmac = crypto.createHmac('sha256', SECRET);
    const digest = 'sha256=' + hmac.update(payload).digest('hex');
    return crypto.timingSafeEqual(Buffer.from(digest), Buffer.from(signature));
}

const server = http.createServer((req, res) => {
    // Health check
    if (req.method === 'GET' && req.url === '/') {
        res.writeHead(200, { 'Content-Type': 'text/plain' });
        res.end('Machine Config Webhook listener active');
        return;
    }

    // Webhook endpoint
    if (req.method === 'POST' && req.url === '/webhook') {
        let body = '';
        req.on('data', chunk => { body += chunk; });
        req.on('end', () => {
            const signature = req.headers['x-hub-signature-256'];

            if (!verifySignature(body, signature)) {
                console.log(`[${new Date().toISOString()}] ❌ Invalid signature`);
                res.writeHead(401);
                res.end('Unauthorized');
                return;
            }

            try {
                const payload = JSON.parse(body);
                const branch = payload.ref;

                // Only deploy on push to main
                if (branch === 'refs/heads/main') {
                    console.log(`[${new Date().toISOString()}] 🚀 Deploying from push to main...`);
                    execSync(`bash ${DEPLOY_SCRIPT}`, { stdio: 'inherit' });
                    res.writeHead(200);
                    res.end('Deployed successfully');
                } else {
                    console.log(`[${new Date().toISOString()}] ⏭️  Ignoring push to ${branch}`);
                    res.writeHead(200);
                    res.end('Ignored - not main branch');
                }
            } catch (err) {
                console.error(`[${new Date().toISOString()}] ❌ Deploy error:`, err.message);
                res.writeHead(500);
                res.end('Deploy failed');
            }
        });
        return;
    }

    res.writeHead(404);
    res.end('Not found');
});

server.listen(PORT, () => {
    console.log(`🎣 Machine Config Webhook listening on port ${PORT}`);
});
