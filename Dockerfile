# Use an official lightweight Python image.
FROM python:3.8-slim

# Set the working directory inside the container.
WORKDIR /app

# Copy and install dependencies.
COPY requirements.txt requirements.txt
RUN pip install --no-cache-dir -r requirements.txt

# Copy the entire project.
COPY . .

# Optionally generate stubs if not already generated:
# RUN python -m grpc_tools.protoc -I. --python_out=./pyserver --grpc_python_out=./pyserver swim.proto

# Set environment variables for non-interactive execution.
# For instance, you could set default values here.
ENV NODE_ID=1
ENV MEMBERS="node1:50051,node2:50052,node3:50053,node4:50054,node5:50055"

# Run the failure detector server.
CMD ["python", "pyserver/failure_detector.py"]
