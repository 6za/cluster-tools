package generator

var MinioIngress = InputIngress{
	IngressName:      "minio-alt",
	IngressNamespace: "minio",
	ServiceName:      "minio",
	ServicePort:      9000,
	SecretName:       "minio-tls",
}

var MinioConsoleIngress = InputIngress{
	IngressName:      "minio-console-alt",
	IngressNamespace: "minio",
	ServiceName:      "minio-console",
	ServicePort:      9001,
	SecretName:       "minio-console-tls",
}

var VaultIngress = InputIngress{
	IngressName:      "vault-alt",
	IngressNamespace: "vault",
	ServiceName:      "vault",
	ServicePort:      8200,
	SecretName:       "vault-tls",
}

var AtlantisIngress = InputIngress{
	IngressName:      "atlantis-alt",
	IngressNamespace: "atlantis",
	ServiceName:      "atlantis",
	ServicePort:      80,
	SecretName:       "atlantis-tls",
}

var ChartMuseumIngress = InputIngress{
	IngressName:      "chartmuseum-alt",
	IngressNamespace: "chartmuseum",
	ServiceName:      "chartmuseum",
	ServicePort:      8080,
	SecretName:       "chartmuseum-tls",
}

var ArgoIngress = InputIngress{
	IngressName:      "argo-server-alt",
	IngressNamespace: "argo",
	ServiceName:      "argo-server",
	ServicePort:      2746,
	SecretName:       "argo-tls",
}

var KubefirstIngress = InputIngress{
	IngressName:      "kubefirst-console-alt",
	IngressNamespace: "kubefirst",
	ServiceName:      "kubefirst-console",
	ServicePort:      80,
	SecretName:       "kubefirst-tls",
}

var ArgoCDIngress = InputIngress{
	IngressName:      "argocd-redirect-alt",
	IngressNamespace: "argocd",
	ServiceName:      "argocd-server",
	ServicePort:      80,
}
