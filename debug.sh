#/bin/sh
CGO_ENABLED=0 go build

export GITREPO=ollijanatuinen@vs-ssh.visualstudio.com:v3/ollijanatuinen/hello-world/hello-world
export ENVIRONMENT=cluster-dev
export PAT=
export AZDEV_REPO_API_URL=https://dev.azure.com/ollijanatuinen/56b215dd-32cb-4e7d-ac42-b4591fafde1e/_apis/git/repositories/2ec01da4-d169-4094-8edb-cabe7313caa9
# export AZDEV_REPO_API_URL=https://webhook.site/fc888d12-927e-4713-acc6-9344b36bfeab
./k8s4u-gitops-agent
