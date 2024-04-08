#!/bin/bash

ls
ls -A
ls -A -1p | grep -vE '^\.{1,2}/?$' | tr '\n' ', ' | sed 's/,$//'
ls -ltu
ls -pum
exit 0