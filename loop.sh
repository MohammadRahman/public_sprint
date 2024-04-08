#!/bin/bash
times=$1

if test $((times)) -gt 100
then
    times=100
fi
for i in $(seq 1 $times)
do
    echo "$i times I've printed mohammadobaidurrahman"
done