#!/bin/bash

cd cmd/bet-starters

command="go run main.go"

if [ "$RUN_MIGRATIONS" = "true" ]; then
  echo "Running migrations"
  command="$command -migrate"
fi

if [ "$RUN_SEEDS" = "true" ]; then
  echo "Running seeders"
  command="$command -seed"
fi

echo "Running main application"
eval $command

cd ../../