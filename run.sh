#!/bin/bash

# Set environment variables
export DOCKER_API_VERSION=1.45
export CUBE_HOST=${CUBE_HOST:-localhost}
export CUBE_PORT=${CUBE_PORT:-5555}

# Function to run in Docker
run_docker() {
    echo "Building and running in Docker..."

    # Build the image
    docker build -t orchestrator-worker .

    # Stop existing container if running
    docker stop orchestrator-worker 2>/dev/null || true
    docker rm orchestrator-worker 2>/dev/null || true

    # Run the container
    docker run -d \
        --name orchestrator-worker \
        -p $CUBE_PORT:5555 \
        -v /var/run/docker.sock:/var/run/docker.sock \
        -e DOCKER_API_VERSION=$DOCKER_API_VERSION \
        -e CUBE_HOST=0.0.0.0 \
        -e CUBE_PORT=5555 \
        orchestrator-worker

    echo "Container started on port $CUBE_PORT"
    echo "View logs with: docker logs -f orchestrator-worker"
}

# Default to docker, but allow local for development
case "$1" in
    "local")
        echo "Running locally..."
        go run main.go
        ;;
    "logs")
        docker logs -f orchestrator-worker
        ;;
    "stop")
        docker stop orchestrator-worker
        docker rm orchestrator-worker
        ;;
    *)
        run_docker
        ;;
esac