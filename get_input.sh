#!/bin/sh

mkdir -p ${AOC_YEAR}

curl -H "Cookie: session=${SESSION_COOKIE}" "https://adventofcode.com/${AOC_YEAR}/day/$1/input" > ${AOC_YEAR}/input/input$1.txt

cd ${AOC_YEAR}
