#!/bin/bash

# 定义 generate_mocks 函数
generate_mocks() {
    local source_dir=$1
    local destination_dir=$2
    local package_name=$3

    files=$(find "$source_dir" -type f -name "*.go")

    for file in $files
    do
        base=$(basename "$file")
        name="${base%.*}"
        destination="$destination_dir/${name}_mock.go"
        mockgen -source="$file" -destination="$destination" -package="$package_name"
    done
}

# 检查参数数量
if [ $# -ne 3 ]; then
    echo "Usage: $0 <source_directory> <destination_directory> <package_name>"
    exit 1
fi

# 调用函数
generate_mocks "$1" "$2" "$3"