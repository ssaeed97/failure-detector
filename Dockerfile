# Use an official Go image which includes Go and a minimal Linux environment.
FROM golang:1.22

# Install Python3 and pip (for Debian-based images)
# Install system dependencies including protoc
RUN apt-get update && apt-get install -y \
    python3 \
    python3-pip \
    build-essential \
    python3-dev \
    protobuf-compiler


# Set the working directory.
WORKDIR /app

# Copy the entire project directory into the container.
COPY . .
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

RUN export PATH="$PATH:$(go env GOPATH)/bin"                  

RUN protoc --go_out=./proto --go-grpc_out=./proto swim.proto

RUN go mod init failure-detector
RUN go mod tidy
# Install Python dependencies.
RUN pip3 install --no-cache-dir --break-system-packages -r requirements.txt


# Expose the necessary ports.
# Python Failure Detector listens on port 50050 + NODE_ID (e.g., 50051, 50052, etc.)
# Go Dissemination service listens on port 50060 + NODE_ID (e.g., 50061, 50062, etc.)
# (Expose a range or specific ports as needed.)
EXPOSE 50051 50052 50053 50054 50055 50061 50062 50063 50064 50065

# Use a startup script to launch both services.
COPY start.sh .
RUN chmod +x start.sh

# Start the container by running the startup script.
CMD ["./start.sh"]
