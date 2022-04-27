# GitOps Agent
k8s4u (Kubernetes for you) is open source project which targets for sharing well tested Kubernetes management code between organizations so that everyone does not need to solve same issues over and over again.

This repository contains the core part of GitOps agent which can be used to enforce GitOps style managed code to environments.

Currently this GitOps agent can be triggered with GitHub or Azure DevOps webhook, it uses personal access token (PAT) to authenticate with both services so it can be used with both public and private repositories.

Reason to use PAT is that same token can be later used to authenticate with REST API to write sync status to both solutions:
* https://docs.github.com/en/rest/reference/commits#create-a-commit-status
* https://docs.microsoft.com/en-us/rest/api/azure/devops/git/statuses/create?view=azure-devops-rest-7.1


You can find example from https://github.com/k8s4u/gitops/commits/ea268113c4a2b90ec2f86bbf32a5056736fdfe2a and clicking that red X commit:

![image](https://user-images.githubusercontent.com/6213926/165469166-2801a03d-6cfe-4581-87c8-12798b29da86.png)
