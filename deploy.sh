#!/bin/bash
# deploy.sh - Auto-deploy script triggered by GitHub webhook
# Location on server: /opt/machine-config/deploy.sh

set -e

LOG_FILE="/var/log/machine-config-deploy.log"
PROJECT_DIR="/opt/machine-config"

echo "$(date) - Deployment triggered" >> "$LOG_FILE"

cd "$PROJECT_DIR"

# Pull latest code
echo "$(date) - Pulling latest code..." >> "$LOG_FILE"
git pull origin main >> "$LOG_FILE" 2>&1

# Rebuild and restart container
echo "$(date) - Building and restarting container..." >> "$LOG_FILE"
docker compose up -d --build --force-recreate >> "$LOG_FILE" 2>&1

# Cleanup old images
docker image prune -f >> "$LOG_FILE" 2>&1

# Verify health
sleep 3
if curl -s http://127.0.0.1:8080/api/health | grep -q "ok"; then
    echo "$(date) - ✅ Deployment successful!" >> "$LOG_FILE"
else
    echo "$(date) - ❌ Deployment failed! Service not healthy." >> "$LOG_FILE"
    exit 1
fi
