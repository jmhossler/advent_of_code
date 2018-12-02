#!/bin/bash

mkdir -p ${AOC_YEAR}/$1

curl -H "Cookie: session=${SESSION_COOKIE}" "https://adventofcode.com/${AOC_YEAR}/day/$1/input" > ${AOC_YEAR}/$1/input

cd ${AOC_YEAR}/$1
