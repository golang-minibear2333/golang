#!/bin/sh

git filter-branch --env-filter '
an="$GIT_AUTHOR_NAME"
am="$GIT_AUTHOR_EMAIL"
cn="$GIT_COMMITTER_NAME"
cm="$GIT_COMMITTER_EMAIL"
if [ "$GIT_COMMITTER_EMAIL" = "minibear2333@example.com" ]
then
    cn="minibear2333"
    cm="minibear2333@qq.com"
fi
if [ "$GIT_AUTHOR_EMAIL" = "minibear2333@example.com" ]
then
    an="minibear2333"
    am="minibear2333@qq.com"
fi
    export GIT_AUTHOR_NAME="$an"
    export GIT_AUTHOR_EMAIL="$am"
    export GIT_COMMITTER_NAME="$cn"
    export GIT_COMMITTER_EMAIL="$cm"
'
