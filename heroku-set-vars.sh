#!/bin/bash
# FOR LOCAL RUN

echo "Heroku | Setting ENV variables."
for i in $(cat .env); do
  heroku config:set $i
done
