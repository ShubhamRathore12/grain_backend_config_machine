# Machine Config Service

A Go microservice that serves machine configuration data (PLC settings, sensors, outputs, menus) via REST API.

## Endpoints

| Method | Path | Description |
|--------|------|-------------|
| GET | `/api/health` | Health check |
| GET | `/api/machines` | List all machines |
| GET | `/api/machines/{name}` | Get full machine config |
| GET | `/api/machines/{name}/auto` | Auto config |
| GET | `/api/machines/{name}/auto-grain` | Grain config |
| GET | `/api/machines/{name}/auto-paddy` | Paddy config |
| GET | `/api/machines/{name}/outputs` | Outputs config |
| GET | `/api/machines/{name}/analog` | Analog config |
| GET | `/api/machines/{name}/tags` | Fault tags |
| GET | `/api/machines/{name}/menu` | Menu config |

## Feature Flag

The service supports a `SERVICE_ENABLED` environment variable:

```bash
# Enable (default)
SERVICE_ENABLED=true

# Disable - service exits immediately without starting
SERVICE_ENABLED=false
```

This lets you deploy the container but keep it dormant until you're ready to go live.

## Local Development

```bash
go run .
# Service starts on http://localhost:8080
```

## Docker

```bash
docker compose up -d --build
```

## Deploy to Server (91.98.235.142)

### Prerequisites
- SSH access to the server
- Docker and Docker Compose installed on server

### Deployment Steps

```powershell
# 1. SSH into the server
ssh -i C:\Users\Shubham\.ssh\ssh-key.key root@91.98.235.142

# 2. Run the deploy script (first time sets everything up)
bash /opt/machine-config/services/machine-config/deploy.sh

# 3. Add nginx location block (one-time setup)
nano /etc/nginx/sites-available/primeosys.com
# Add the location block from nginx-machine-config.conf
nginx -t && systemctl reload nginx
```

### Toggle Service On/Off (Feature Flag)

```bash
# Disable the service
cd /opt/machine-config/services/machine-config
echo "MACHINE_CONFIG_ENABLED=false" > .env
docker compose up -d --build

# Enable the service
echo "MACHINE_CONFIG_ENABLED=true" > .env
docker compose up -d --build
```

### Verify Deployment

```bash
# Local check on server
curl http://127.0.0.1:8080/api/health

# Public URL (after nginx config)
curl https://www.primeosys.com/machine-config/api/health

# List machines
curl https://www.primeosys.com/machine-config/api/machines
```

## Architecture

```
Internet → Nginx (primeosys.com)
              ├── /backend/api/*         → grain-backend:3000  (existing)
              └── /machine-config/api/*  → machine-config:8080 (this service)
```

Both services run independently. Stopping/disabling machine-config has zero impact on grain-backend.
# grain_backend_config_machine
