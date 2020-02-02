#!/bin/bash

for i in $(cat .env); do
  # echo "export $i"
  export $i
done
