#!/bin/bash

# build the program
go build -o ./main ./cmd/api

# run based on the operating system type
if [[ "$OSTYPE" == "linux-gnu" || "$OSTYPE" == "darwin"* ]]; then
  ./main
elif [[ "$OSTYPE" == "msys" || "$OSTYPE" == "cygwin" || "$OSTYPE" == "win32" ]]; then
  main.exe
else
  echo "Unknown operating system"
fi
