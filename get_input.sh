#!/bin/bash

mkdir -p 2018/$1

curl -H "Cookie: session=${SESSION_COOKIE}" "https://adventofcode.com/2018/day/$1/input" > 2018/$1/input
