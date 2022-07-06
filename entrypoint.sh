#!/bin/sh
set -e

# Check if the config file was specified
if [ -z $ONKYO_CONFIG ]; then 
    echo "No configuration file specified!"
    exit 1
fi

exec /app/onkyo-remote -config $ONKYO_CONFIG "$@"
