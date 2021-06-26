# /bin/bash

pushd blog/public
hugo -D
git init
git remote add origin https://github.com/XXX/XXX.github.io.git
git add -A
git commit -m "first commit"
git push -u origin graph