#!/bin/bash

mkdir -p 2017/$1

curl -H "Cookie: session=53616c7465645f5f7f1675d7e85e97cf406f52ce6fcd5fa4cfe907c0d5108674ac60dbaaf263f35c190daa4e8f0cb964" "http://adventofcode.com/2017/day/$1/input" > 2017/$1/input
