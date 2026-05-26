#!/bin/bash
# deploy.sh - Deploy machine-config-service to /opt/machine-config on the server
# This script runs ON the server (not locally)

set -e

PROJECT_DIR="/opt/machine-config"
REPO_URL="https://github.com/ShubhamRathore12/grain_backend.git"
SERVICE_PATH="services/machine-config"

echo "=== Machine Config Service - Deployment ==="

# Create project directory if it doesn't exist
if [ ! -d "$PROJECT_DIR" ]; then
    echo "📁 Creating project directory..."
    mkdir -p "$PROJECT_DIR"
fi

# If repo not cloned yet, do a sparse checkout of just this service
if [ ! -d "$PROJECT_DIR/.git" ]; then
    echo "📥 Cloning repository (sparse checkout)..."
    cd /opt
    git clone --filter=blob:none --sparse "$REPO_URL" machine-config
    cd "$PROJECT_DIR"
    git sparse-checkout set "$SERVICE_PATH"
else
    echo "🔄 Pulling latest changes..."
    cd "$PROJECT_DIR"
    git pull origin main
fi

# Navigate to service directory
cd "$PROJECT_DIR/$SERVICE_PATH"

# Create .env if it doesn't exist
if [ ! -f .env ]; then
    echo "📝 Creating .env file..."
    cat > .env << 'EOF'
PORT=8080
MACHINE_CONFIG_ENABLED=true
EOF
fi

# Build and deploy
echo "🐳 Building and deploying Docker container..."
docker compose up -d --build --force-recreate

# Cleanup old images
docker image prune -f

# Verify
echo "⏳ Waiting for service to start..."
sleep 3

if curl -s http://127.0.0.1:8080/api/health | grep -q "ok"; then
    echo "✅ Machine Config Service is running!"
    echo "   Health: http://127.0.0.1:8080/api/health"
else
    echo "❌ Service failed to start. Check logs:"
    echo "   docker logs machine-config-service --tail 20"
    exit 1
fi

echo ""
echo "=== Deployment Complete ==="
echo "Next: Configure nginx to proxy /machine-config/api/* to port 8080"
