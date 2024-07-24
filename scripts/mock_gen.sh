#!/bin/bash

files=$(find internal/repo/ -type f -name "*.go")

for file in $files
do
    base=$(basename "$file")
    name="${base%.*}"
    destination="test/repo/${name}_mock.go"
    mockgen -source="$file" -destination="$destination" -package=repo
done