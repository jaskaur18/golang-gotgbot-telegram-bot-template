#!/bin/bash

# Exit immediately if a command exits with a non-zero status
set -e

echo "Starting installation..."

# Update system
sudo apt-get update
sudo apt-get upgrade -y

# Install Docker prerequisites
sudo apt-get install -y ca-certificates curl gnupg

# Add Docker's official GPG key
sudo install -m 0755 -d /etc/apt/keyrings
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o /etc/apt/keyrings/docker.gpg
sudo chmod a+r /etc/apt/keyrings/docker.gpg

# Set up the Docker repository
echo \
  "deb [arch=\"$(dpkg --print-architecture)\" signed-by=/etc/apt/keyrings/docker.gpg] https://download.docker.com/linux/ubuntu \
  $(. /etc/os-release && echo \"$VERSION_CODENAME\") stable" | \
  sudo tee /etc/apt/sources.list.d/docker.list > /dev/null

# Install Docker Engine
sudo apt-get update
sudo apt-get install -y docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin

# Verify Docker installation
sudo docker --version

# Build and start services using Docker Compose
echo "Building and starting services..."
if [ -f "docker-compose.yml" ]; then
    sudo docker compose up -d --build
    echo "Services started successfully!"
    echo "You can check logs with: sudo docker compose logs -f"
else
    echo "Error: docker-compose.yml not found in current directory."
    exit 1
fi
