#!/bin/bash

# Path to the .env file
ENV_FILE=".env"

if [ -f "$ENV_FILE" ]; then
    echo "Loading environment variables from $ENV_FILE..."
    # Secure the .env file permissions
    chmod 600 "$ENV_FILE"
    # Export variables from .env, ignoring comments and empty lines
    export $(grep -v '^#' "$ENV_FILE" | xargs)
else
    echo "Warning: $ENV_FILE not found. Using default environment."
fi

# Ensure the binary has permission to bind to low ports
sudo setcap 'cap_net_bind_service=+ep' ./website_server

# Kill any existing server process
pkill website_server || true

# Start the server in the background
echo "Starting website_server..."
nohup ./website_server > server.log 2>&1 &

echo "Server started. Check server.log for details."
