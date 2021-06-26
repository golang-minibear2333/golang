# /bin/bash

currPath=`pwd`
blog="blog"

function buildFun() {
    echo pwd: `pwd`
    hugo -D
    cp -r public/* ../docs/
    rm -rf public/*
}

if [ -d ${currPath}/${blog} ]; then
  pushd ${currPath}/${blog}
  buildFun
  popd
elif [ -d ${currPath}/../${blog} ]; then
  pushd ${currPath}/../${blog}
  buildFun
  popd
fi
