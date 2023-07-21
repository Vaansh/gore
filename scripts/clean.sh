#!/bin/bash

# remove files based on operating system type
if [[ "$OSTYPE" == "linux-gnu" || "$OSTYPE" == "darwin"* ]]; then
  rm ./data/*.mp4
  rm ./log/*-*.log
  rm main
elif [[ "$OSTYPE" == "msys" || "$OSTYPE" == "cygwin" || "$OSTYPE" == "win32" ]]; then
  del .\data\*.mp4
  del .\log\*.log
  del main
else
  echo "Unknown operating system"
fi
