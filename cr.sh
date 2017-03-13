#!/usr/bin/env bash

./compile.sh
if [ $? -eq 0 ]; then
    echo OK
    ./run.sh
else
    echo FAIL
fi
