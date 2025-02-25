#!/bin/sh
if [ -z "$MEMBERS" ]; then
    echo "MEMBERS not provided. Fetching membership list from bootstrap node..."
    MEMBERS=$(python3 pyserver/join_client.py)
    export MEMBERS
    echo "Updated MEMBERS: $MEMBERS"
fi
# Start the Go Dissemination service in the background.
# Assuming your Go code is in, for example, go-dissemination/main.go
go run main.go &

# Start the Python Failure Detector service.
python3 pyserver/failure_detector.py
