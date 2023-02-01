# cluster-tools

A set of tools to add in a k8s cluster to give some extra powers to it.


## Ingress Generantor(cluster-tools ingress-gen)

A tool used to help users of [kubefirst](https://github.com/kubefirst/kubefirst) to explore their local installation under a internal LAN, sharing the cluster with multiple machines in this network. 

Tested for release: 1.11 
```bash 
 docker run -it --rm -e GITHUB_AUTH_TOKEN="ghp_token" \
     cluster-tools  \
     /home/developer/app/cluster-tools ingress-gen \
     --host-domain mylocal.cloud.internal \
     --repo https://github.com/6za/gitops.git \
     --ip 10.10.10.5

```

| Flag          | Description                                                                          |
|:--------------|:-------------------------------------------------------------------------------------|
| --host-domain | A domain to be used on your nodes /etc/hosts file to identify your k8s cluster host. |
| --repo        | Where your gitops repo is defined, only supported HTTPS at the moment.               |
| --ip          | The kubernetes cluster host IP on your LAN(internal network)                         |



# Ngrok-AGENT


```yaml 
apiVersion: argoproj.io/v1alpha1
kind: Application
metadata:
  name: ngrok-agent
  namespace: argocd
  annotations:    
    argocd.argoproj.io/sync-wave: "0"
spec: 
  project: default
  source:
    repoURL: 'https://6za.github.io/cluster-tools'
    targetRevision: 0.5.0
    chart: tunnel-agent
  destination:
    server: 'https://kubernetes.default.svc'
    namespace: watcher-system
  syncPolicy:
      automated:
        prune: true
        selfHeal: true
      syncOptions:
        - CreateNamespace=true
      retry:
        limit: 5
        backoff:
          duration: 5s
          maxDuration: 5m0s
          factor: 2
```