#!/bin/sh
# Start the Go Dissemination service in the background.
# Assuming your Go code is in, for example, go-dissemination/main.go
go run main.go &

# Start the Python Failure Detector service.
python3 pyserver/failure_detector.py
