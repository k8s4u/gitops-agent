#!/bin/sh
cd /gitops
if [ ! -d "/gitops/.git" ]; then
  git clone $GITREPO .
fi
git pull
kubectl apply -k envs/$ENVIRONMENT
exit $?
