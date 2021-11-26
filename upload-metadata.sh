#!/bin/bash -e

# login beforehand
# snapcraft login

NAME=$1

if [ -z $NAME ]; then
    echo "Missing snap name!"
    exit 1
fi

# build
snapcraft

# upload
snapcraft upload-metadata $1_noversion_amd64.snap --force