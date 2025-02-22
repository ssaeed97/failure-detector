# failure-detector
 
python -m venv venv
source venv/bin/activate
pip install grpcio grpcio-tools
python -m grpc_tools.protoc -I. --python_out=./pyserver --grpc_python_out=./pyserver swim.proto

# Failure Detector - SWIM Protocol Implementation

This project implements a simple failure detection component of the SWIM protocol using gRPC in Python. It uses both direct and indirect pings to detect failures among nodes in a distributed system.

## Table of Contents

- [Overview](#overview)
- [Project Structure](#project-structure)
- [Prerequisites](#prerequisites)
- [Setup & Local Testing](#setup--local-testing)
  - [1. Virtual Environment Setup](#1-virtual-environment-setup)
  - [2. Install Dependencies](#2-install-dependencies)
  - [3. Generate gRPC Stubs](#3-generate-grpc-stubs)
  - [4. Running Locally](#4-running-locally)
- [Containerization](#containerization)
  - [1. Dockerfile](#1-dockerfile)
  - [2. Running Containers Manually](#2-running-containers-manually)
  - [3. (Optional) Docker Compose](#3-optional-docker-compose)
- [Notes](#notes)

## Overview

This project implements a failure detector that:
- Sends a **Ping** every T' seconds to a random node.
- If no direct response is received, the node sends an **IndirectPing** via k proxy nodes.
- Logs are printed on both the client side (before making an RPC call) and the server side (upon receiving an RPC call) to trace the interactions.

## Project Structure

We test the implementation using a simple pyserver which is located in the subdir within the root. The stubs are created based on the proto fole present.


## Prerequisites

- Python 3.8 or later
- `pip`
- `virtualenv` (optional, built-in `venv` works as well)
- Docker (if you plan to containerize)
- (Optional) Docker Compose

## Setup & Local Testing

### 1. Virtual Environment Setup

Open your terminal in the project root and run:

```bash
python -m venv venv
```

### Activate the VENV

`source venv/bin/activate`

### 2. Install Dependencies

With the virtual environment activated, install the required packages:

Plain textANTLR4BashCC#CSSCoffeeScriptCMakeDartDjangoDockerEJSErlangGitGoGraphQLGroovyHTMLJavaJavaScriptJSONJSXKotlinLaTeXLessLuaMakefileMarkdownMATLABMarkupObjective-CPerlPHPPowerShell.propertiesProtocol BuffersPythonRRubySass (Sass)Sass (Scss)SchemeSQLShellSwiftSVGTSXTypeScriptWebAssemblyYAMLXML`   bashCopypip install grpcio grpcio-tools   `

### 3. Generate gRPC Stubs

Make sure your swim.proto file is in the root directory. Run the following command to generate the Python stubs:

`python -m grpc_tools.protoc -I. --python_out=./pyserver --grpc_python_out=./pyserver swim.proto `

This will create swim_pb2.py and swim_pb2_grpc.py inside the pyserver/ directory.

### 4. Running Locally

You can test the failure detector locally by running multiple instances in separate terminal windows. For example, to run three nodes:

`python pyserver/failure_detector.py`
    
*   When prompted, enter a node ID (e.g., 1, 2, or 3).
    

The code uses a simple mechanism to bind ports as follows:

*   Port = 50050 + node ID(e.g., node 1 runs on 50051, node 2 on 50052, etc.)
    

Logs will display RPC calls (both client and server sides) as nodes ping one another.

## Containerization

### 1. Dockerfile

An example Dockerfile is provided to containerize the application. It installs dependencies, copies the project files, and runs the failure detector. A sample Dockerfile:


### 2. Running Containers Manually

Before using Docker Compose, you can create a Docker network and run each container individually.

1.  `docker build -t failure-detector .`
    
2.  `docker network create failure-network`
    
3.  For three nodes, run these commands in separate terminals:
    
    *   `docker run -it --name node1 --network failure-network -p 50051:50051 -e NODE_ID=1 -e MEMBERS="node1:50051,node2:50052,node3:50053" failure-detector`
        
    *   `docker run -it --name node2  --network failure-network  -p 50052:50052  -e NODE_ID=2  -e MEMBERS="node1:50051,node2:50052,node3:50053"  failure-detector`
        
    *   `docker run -it  --name node3  --network failure-network  -p 50053:50053  -e NODE_ID=3  -e MEMBERS="node1:50051,node2:50052,node3:50053"  failure-detector`
        

You should see interactive outputs from each container, showing the RPC messages as nodes ping each other.

## Notes


*   **Environment Variables:**The application reads NODE_ID and MEMBERS from the environment. If not provided, it falls back to interactive input (which is useful for local testing but not in containerized environments).
    
*   **Networking in Docker:**When running in Docker, use container names (e.g., node1) instead of localhost for inter-container communication.
    
*   **Resource Usage:**If you encounter exit codes like 137, it may be due to resource limits. Monitor container logs and adjust memory limits if needed.
    
*   **Logs:**The application prints both client-side and server-side logs in the following format:
    
    *   Client Side:Component FailureDetector of Node sends RPC to Component FailureDetector of Node
        
    *   Server Side:Component FailureDetector of Node runs RPC called by Component FailureDetector of Node
        

This README should help you get started with both local and containerized setups for the failure detector project. Enjoy experimenting with distributed systems!