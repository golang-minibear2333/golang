# /bin/bash

currPath=`pwd`
blog="blog"

function buildFun() {
    echo pwd: `pwd`
    cp -r blog blog.bak
    mv blog.bak
    hugo -D
    rm -rf ../../golang-minibear2333.github.io/*
    cp -r public/* ../../golang-minibear2333.github.io
    rm -rf public/*
}

if [ -d ${currPath}/${blog} ]; then
  pushd ${currPath}/
  buildFun
  popd
elif [ -d ${currPath}/../${blog} ]; then
  pushd ${currPath}/..
  buildFun
  popd
fi
