#!/bin/bash
# server-setup.sh - Complete one-time setup on the server
# Run this ONCE after SSH-ing into 91.98.235.142
#
# Usage: bash /tmp/server-setup.sh

set -e

echo "============================================"
echo "  Machine Config Service - Server Setup"
echo "============================================"
echo ""

PROJECT_DIR="/opt/machine-config"
GITHUB_REPO="https://github.com/ShubhamRathore12/grain_backend_config_machine.git"

# ─── Step 1: Clone the repo ──────────────────────────────────────────────────
echo "📥 Step 1: Cloning repository..."
if [ -d "$PROJECT_DIR" ]; then
    echo "   Directory exists, pulling latest..."
    cd "$PROJECT_DIR"
    git pull origin main
else
    git clone "$GITHUB_REPO" "$PROJECT_DIR"
    cd "$PROJECT_DIR"
fi

# The repo root IS the service directory
SERVICE_DIR="$PROJECT_DIR"

echo "   ✅ Repository ready"
echo ""

# ─── Step 2: Create .env file ────────────────────────────────────────────────
echo "📝 Step 2: Creating .env file..."
if [ ! -f "$SERVICE_DIR/.env" ]; then
    cat > "$SERVICE_DIR/.env" << 'EOF'
PORT=8080
MACHINE_CONFIG_ENABLED=true
EOF
    echo "   ✅ .env created"
else
    echo "   ⏭️  .env already exists, skipping"
fi
echo ""

# ─── Step 3: Build and start Docker container ────────────────────────────────
echo "🐳 Step 3: Building and starting Docker container..."
cd "$SERVICE_DIR"
docker compose up -d --build
echo "   ✅ Container started"
echo ""

# ─── Step 4: Setup webhook listener ──────────────────────────────────────────
echo "🎣 Step 4: Setting up GitHub webhook listener..."

# Make deploy script executable
chmod +x "$PROJECT_DIR/deploy.sh"

# Install systemd service
cp "$PROJECT_DIR/machine-config-webhook.service" /etc/systemd/system/machine-config-webhook.service
systemctl daemon-reload
systemctl enable machine-config-webhook
systemctl start machine-config-webhook

echo "   ✅ Webhook listener running on port 9001"
echo ""

# ─── Step 5: Configure Nginx ─────────────────────────────────────────────────
echo "🌐 Step 5: Configuring Nginx..."

NGINX_CONF="/etc/nginx/sites-available/primeosys.com"

# Check if machine-config location already exists
if grep -q "machine-config" "$NGINX_CONF" 2>/dev/null; then
    echo "   ⏭️  Nginx already configured for machine-config"
else
    # Add location block before the last closing brace of the server block
    # We'll create a snippet and include it
    SNIPPET="/etc/nginx/snippets/machine-config.conf"
    cat > "$SNIPPET" << 'EOF'
# Machine Config Service proxy
location /machine-config/api/ {
    proxy_pass http://127.0.0.1:8080/api/;
    proxy_http_version 1.1;
    proxy_set_header Upgrade $http_upgrade;
    proxy_set_header Connection 'upgrade';
    proxy_set_header Host $host;
    proxy_set_header X-Real-IP $remote_addr;
    proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
    proxy_set_header X-Forwarded-Proto $scheme;
    proxy_cache_bypass $http_upgrade;
}
EOF

    # Add include directive to the main config if not already there
    if ! grep -q "snippets/machine-config.conf" "$NGINX_CONF" 2>/dev/null; then
        # Insert before the last } in the file
        sed -i '/^}/i\    include /etc/nginx/snippets/machine-config.conf;' "$NGINX_CONF"
    fi

    # Test and reload
    nginx -t
    systemctl reload nginx
    echo "   ✅ Nginx configured and reloaded"
fi
echo ""

# ─── Step 6: Open firewall for webhook port ──────────────────────────────────
echo "🔥 Step 6: Opening firewall port 9001..."
if command -v ufw &> /dev/null; then
    ufw allow 9001/tcp 2>/dev/null || true
    echo "   ✅ UFW rule added for port 9001"
elif command -v firewall-cmd &> /dev/null; then
    firewall-cmd --permanent --add-port=9001/tcp 2>/dev/null || true
    firewall-cmd --reload 2>/dev/null || true
    echo "   ✅ Firewalld rule added for port 9001"
else
    echo "   ⚠️  No firewall detected, port 9001 should be open"
fi
echo ""

# ─── Step 7: Verify everything ───────────────────────────────────────────────
echo "🔍 Step 7: Verifying deployment..."
echo ""

sleep 3

# Check container
if docker ps | grep -q "machine-config-service"; then
    echo "   ✅ Docker container: RUNNING"
else
    echo "   ❌ Docker container: NOT RUNNING"
fi

# Check health
if curl -s http://127.0.0.1:8080/api/health | grep -q "ok"; then
    echo "   ✅ Health check: OK"
else
    echo "   ❌ Health check: FAILED"
fi

# Check webhook
if curl -s http://127.0.0.1:9001/ | grep -q "active"; then
    echo "   ✅ Webhook listener: ACTIVE"
else
    echo "   ❌ Webhook listener: NOT RUNNING"
fi

# Check nginx
if nginx -t 2>&1 | grep -q "successful"; then
    echo "   ✅ Nginx config: VALID"
else
    echo "   ❌ Nginx config: INVALID"
fi

echo ""
echo "============================================"
echo "  ✅ SETUP COMPLETE!"
echo "============================================"
echo ""
echo "  Service URL:  https://www.primeosys.com/machine-config/api/health"
echo "  Webhook URL:  http://91.98.235.142:9001/webhook"
echo ""
echo "  ┌─────────────────────────────────────────┐"
echo "  │  NEXT: Configure GitHub Webhook         │"
echo "  │                                         │"
echo "  │  Go to: GitHub Repo → Settings →        │"
echo "  │         Webhooks → Add webhook          │"
echo "  │                                         │"
echo "  │  Payload URL:                           │"
echo "  │    http://91.98.235.142:9001/webhook    │"
echo "  │                                         │"
echo "  │  Content type: application/json         │"
echo "  │  Secret: machine-config-deploy-secret-2026 │"
echo "  │  Events: Just the push event            │"
echo "  └─────────────────────────────────────────┘"
echo ""
echo "  After that, every push to 'main' will auto-deploy!"
echo ""
