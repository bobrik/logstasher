#!/bin/sh

set -e

if [ -d /target ]; then
    echo "installing logstaher in /target"
    cp /bin/logstasher /target/logstasher
    exit 0
fi

exec /bin/logstasher $@
