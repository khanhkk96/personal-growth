#!/bin/bash
set -e

rm -f ./app_di.go

cd ./di/core
wire
mv wire_gen.go ../app_di.go
eho "Wire generation completed successfully"