#!/bin/bash

# https://gozeon.github.io/cheatsheets/bash/npm%E5%8C%85%E4%B8%8B%E8%BD%BD/

npm_registry="https://registry.npmmirror.com"
output="./public/static"

download_package() {
    mkdir -p $output/$1
    wget -qO- $npm_registry/$1/-/$1-$2.tgz \
        | tar -xzv --strip-components=1 -C $output/$1
}

download_package "@popperjs/core" "2.11.8"
download_package "bootstrap" "5.3.3"
download_package "bootstrap-icons" "1.11.3"
download_package "jquery" "3.7.1"