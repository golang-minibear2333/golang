# /bin/bash

currPath=`pwd`
blog="blog"

function buildFun() {
    echo pwd: `pwd`
    cp -r blog blog.bak
    mv blog.bak ../
    cd ../blog.bak && hugo -D
    cd ..
    rm -rf ./golang-minibear2333.github.io/*
    mv blog.bak/public/* ./golang-minibear2333.github.io/
    rm -rf blog.bak/
    cd ./golang-minibear2333.github.io/
    git add -A
    git commit -m "update"
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