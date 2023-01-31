# cluster-tools

A set of tools to add in a k8s cluster to give some extra powers to it.


## Ingress Generantor(cluster-tools ingress-gen)

A tool used to help users of [kubefirst](https://github.com/kubefirst/kubefirst) to explore their local installation under a internal LAN, sharing the cluster with multiple machines in this network. 

Tested for release: 1.11 
```bash 
 docker run -it --rm -e GITHUB_AUTH_TOKEN="ghp_token"  cluster-tools  /home/developer/app/cluster-tools ingress-gen --host-domain mylocal.cloud.internal --repo https://github.com/6za/gitops.git --ip 10.10.10.5

```

| Flag           | Description                                                                          |
|:---------------|:-------------------------------------------------------------------------------------|
| --host-domain | A domain to be used on your nodes /etc/hosts file to identify your k8s cluster host. |
| --repo        | Where your gitops repo is defined, only supported HTTPS at the moment.               |
| --ip           | The kubernetes cluster host IP on your LAN(internal network)                         |


