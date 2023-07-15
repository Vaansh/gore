#!/bin/bash

# build the program
go build -o ../main ../cmd/api

# unix
if [[ "$OSTYPE" == "linux-gnu" || "$OSTYPE" == "darwin"* ]]; then
  cd ..
  ./main
# windows
elif [[ "$OSTYPE" == "msys" || "$OSTYPE" == "cygwin" || "$OSTYPE" == "win32" ]]; then
  cd ..
  main.exe
else
  echo "Unknown operating system"
fi
