# Please always use the same golang version that
# gomponents uses to maximize compatibility
FROM golang:1.18 AS golang

# This is the actual image we will use
FROM debian:trixie

# Copy the golang binaries from the golang image
COPY --from=golang /usr/local/go /usr/local/go
ENV PATH "$PATH:/usr/local/go/bin"

# Install dependencies, tools and setup bashrc
RUN \
    # Create a temporary working directory
    mkdir -p /app/temp && \
    cd /app/temp && \
    
    # Install system dependencies
    apt-get update && \
    apt-get install -y wget git && \

    # Install additional tools
    wget https://github.com/cli/cli/releases/download/v2.46.0/gh_2.46.0_linux_amd64.tar.gz && \
    tar -xzf gh_2.46.0_linux_amd64.tar.gz && \
    mv gh_2.46.0_linux_amd64/bin/gh /usr/local/bin/gh && \
    wget https://github.com/go-task/task/releases/download/v3.34.1/task_linux_amd64.tar.gz && \
    tar -xzf task_linux_amd64.tar.gz && \
    mv ./task /usr/local/bin/task && \
    wget https://github.com/tdewolff/minify/releases/download/v2.20.34/minify_linux_amd64.tar.gz && \
    tar -xzf minify_linux_amd64.tar.gz && \
    mv ./minify /usr/local/bin/minify && \

    # Make sure the tools are executable
    chmod +x /usr/local/bin/* && \

    # Cleanup
    rm -rf /var/lib/apt/lists/* && \
    cd / && rm -rf /app/temp

WORKDIR /app

# Add the startup script on every bash session
RUN rm -rf /root/.bashrc
COPY .devcontainer/.bashrc /root/.bashrc

# Command just to keep the container running
CMD ["sleep", "infinity"]
