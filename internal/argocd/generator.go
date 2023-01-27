package argocd

func GenerateApp(name string, appNamespace string, path string, repoURL string) (*ArgocdApp, error) {
	result := &ArgocdApp{
		APIVersion: "argoproj.io/v1alpha1",
		Kind:       "Application",
		Metadata: Metadata{
			Name:        name,
			Namespace:   "argocd",
			Annotations: Annotations{ArgocdArgoprojIoSyncWave: "10"},
		},
		Spec: Spec{
			Project: "default",
			Source: Source{
				RepoURL:        repoURL,
				Path:           path,
				TargetRevision: "HEAD",
			},
			Destination: Destination{
				Server:    "https://kubernetes.default.svc",
				Namespace: "ingress-alt",
			},
			SyncPolicy: SyncPolicy{
				Automated: Automated{
					Prune:    true,
					SelfHeal: true,
				},
				SyncOptions: []string{"CreateNamespace=true"},
			},
		},
	}
	return result, nil
}
