package argocd

type ArgocdApp struct {
	APIVersion string   `yaml:"apiVersion"`
	Kind       string   `yaml:"kind"`
	Metadata   Metadata `yaml:"metadata"`
	Spec       Spec     `yaml:"spec"`
}
type Annotations struct {
	ArgocdArgoprojIoSyncWave string `yaml:"argocd.argoproj.io/sync-wave"`
}
type Metadata struct {
	Name        string      `yaml:"name"`
	Namespace   string      `yaml:"namespace"`
	Annotations Annotations `yaml:"annotations"`
}
type Source struct {
	RepoURL        string `yaml:"repoURL"`
	Path           string `yaml:"path"`
	TargetRevision string `yaml:"targetRevision"`
}
type Destination struct {
	Server    string `yaml:"server"`
	Namespace string `yaml:"namespace"`
}
type Automated struct {
	Prune    bool `yaml:"prune"`
	SelfHeal bool `yaml:"selfHeal"`
}
type SyncPolicy struct {
	Automated   Automated `yaml:"automated"`
	SyncOptions []string  `yaml:"syncOptions"`
}
type Spec struct {
	Project     string      `yaml:"project"`
	Source      Source      `yaml:"source"`
	Destination Destination `yaml:"destination"`
	SyncPolicy  SyncPolicy  `yaml:"syncPolicy"`
}
