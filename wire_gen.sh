#!/bin/bash
set -e

cd ./di
rm -f app_di.go

cd ./core
wire
mv wire_gen.go ../app_di.go
eho "Wire generation completed successfully"