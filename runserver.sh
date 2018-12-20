#!/bin/bash

echo "Compiling ..."
go build Server.go
echo "Running server ..."
go run Server.go
