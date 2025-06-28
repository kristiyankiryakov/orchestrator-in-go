#!/bin/bash


export DOCKER_API_VERSION=1.45
export CUBE_WORKER_HOST=0.0.0.0
export CUBE_WORKER_PORT=5555
export CUBE_MANAGER_HOST=0.0.0.0
export CUBE_MANAGER_PORT=5556



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
         -v /var/run/docker.sock:/var/run/docker.sock \
         -e DOCKER_API_VERSION=$DOCKER_API_VERSION \
         -e CUBE_WORKER_HOST=$CUBE_WORKER_HOST \
         -e CUBE_WORKER_PORT=$CUBE_WORKER_PORT \
         -e CUBE_MANAGER_HOST=$CUBE_MANAGER_HOST \
         -e CUBE_MANAGER_PORT=$CUBE_MANAGER_PORT \
         -p 5555:5555 \
         -p 5556:5556 \
         orchestrator-worker

   echo "Container started on ports $CUBE_WORKER_PORT and $CUBE_MANAGER_PORT"
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