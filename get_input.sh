#!/bin/bash

mkdir -p 2018/$1

curl -H "Cookie: session=53616c7465645f5f449edbfe74d47c955bade8432ee5e06da47c0067f7a228922c7d63296e4f7f1e95bd99c2adaefaa1" "https://adventofcode.com/2018/day/$1/input" > 2018/$1/input
