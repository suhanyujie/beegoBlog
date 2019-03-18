#!/bin/bash
#
# 编译 并运行博客程序
go build -o beegoBlog main.go
./beegoBlog > /dev/null &
