# failure-detector
 
python -m venv venv
source venv/bin/activate
pip install grpcio grpcio-tools
python -m grpc_tools.protoc -I. --python_out=./pyserver --grpc_python_out=./pyserver swim.proto