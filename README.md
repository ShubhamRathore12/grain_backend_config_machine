# Machine Config Service

Go microservice serving machine configuration data via REST API. Auto-deploys on push to `main`.

## Quick Deploy (One-Time Server Setup)

```powershell
# 1. SSH into server
ssh -i C:\Users\Shubham\.ssh\ssh-key.key root@91.98.235.142

# 2. Push this code to GitHub first, then on the server:
cd /opt
git clone https://github.com/ShubhamRathore12/grain_backend_config_machine.git machine-config
cd /opt/machine-config
bash server-setup.sh
```

That's it. The script handles everything: Docker build, nginx config, webhook listener.

## 3. Configure GitHub Webhook (One-Time)

Go to your GitHub repo → **Settings** → **Webhooks** → **Add webhook**:

| Field | Value |
|-------|-------|
| Payload URL | `http://91.98.235.142:9001/webhook` |
| Content type | `application/json` |
| Secret | `machine-config-deploy-secret-2026` |
| Events | Just the push event |

After this, every push to `main` auto-deploys to the server.

## Auto-Deploy Flow

```
Push to main → GitHub Webhook → http://91.98.235.142:9001/webhook → deploy.sh → git pull + docker rebuild
```

## API Endpoints

| Method | Path | Description |
|--------|------|-------------|
| GET | `/machine-config/api/health` | Health check |
| GET | `/machine-config/api/machines` | List all machines |
| GET | `/machine-config/api/machines/{name}` | Full machine config |
| GET | `/machine-config/api/machines/{name}/auto` | Auto config |
| GET | `/machine-config/api/machines/{name}/auto-grain` | Grain config |
| GET | `/machine-config/api/machines/{name}/auto-paddy` | Paddy config |
| GET | `/machine-config/api/machines/{name}/outputs` | Outputs config |
| GET | `/machine-config/api/machines/{name}/analog` | Analog config |
| GET | `/machine-config/api/machines/{name}/tags` | Fault tags |
| GET | `/machine-config/api/machines/{name}/menu` | Menu config |

**Base URL**: `https://www.primeosys.com/machine-config/api/`

## Feature Flag

Toggle the service without removing the container:

```bash
# Disable
ssh root@91.98.235.142
cd /opt/machine-config
sed -i 's/MACHINE_CONFIG_ENABLED=true/MACHINE_CONFIG_ENABLED=false/' .env
docker compose up -d --build

# Enable
sed -i 's/MACHINE_CONFIG_ENABLED=false/MACHINE_CONFIG_ENABLED=true/' .env
docker compose up -d --build
```

## Architecture

```
Internet → Nginx (primeosys.com)
              ├── /backend/api/*         → grain-backend:3000   (existing, untouched)
              └── /machine-config/api/*  → machine-config:8080  (this service)

GitHub Push → Webhook:9001 → deploy.sh → docker compose rebuild
```

## Troubleshooting

```bash
# Check container
docker ps | grep machine-config
docker logs machine-config-service --tail 30

# Check webhook
systemctl status machine-config-webhook
curl http://127.0.0.1:9001/

# Check deploy log
tail -20 /var/log/machine-config-deploy.log

# Manual redeploy
cd /opt/machine-config
git pull origin main
docker compose up -d --build --force-recreate
```

## Local Development

```bash
go run .
# http://localhost:8080/api/health
```
