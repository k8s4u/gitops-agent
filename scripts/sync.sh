#!/bin/sh

# Create subfolder because /gitops is owned by root and git complains about it
if [ ! -d "/gitops/code" ]; then
  mkdir /gitops/code
fi

# Get personal access token as base64 encoded string
cd /gitops/code
PAT=$(cat /git-auth/pat)
B64_PAT=$(printf "%s""x-access-token:$PAT" | base64 -w 0)

# Initial clone
if [ ! -d "/gitops/code/.git" ]; then
  git -c http.extraHeader="Authorization: Basic ${B64_PAT}" -c http.version="HTTP/1.1" clone $GITREPO .
fi

git -c http.extraHeader="Authorization: Basic ${B64_PAT}" -c http.version="HTTP/1.1" pull

kubectl apply -k envs/$ENVIRONMENT
exit $?
