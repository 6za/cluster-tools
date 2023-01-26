package generator

type IngressKind struct {
	APIVersion string   `yaml:"apiVersion"`
	Kind       string   `yaml:"kind"`
	Metadata   Metadata `yaml:"metadata"`
	Spec       Spec     `yaml:"spec"`
}
type Metadata struct {
	Name      string `yaml:"name"`
	Namespace string `yaml:"namespace"`
}
type Port struct {
	Number int `yaml:"number"`
}
type Service struct {
	Name string `yaml:"name"`
	Port Port   `yaml:"port"`
}
type Backend struct {
	Service Service `yaml:"service"`
}
type Paths struct {
	Backend  Backend `yaml:"backend"`
	Path     string  `yaml:"path"`
	PathType string  `yaml:"pathType"`
}
type HTTP struct {
	Paths []Paths `yaml:"paths"`
}
type Rules struct {
	Host string `yaml:"host"`
	HTTP HTTP   `yaml:"http"`
}
type TLS struct {
	Hosts      []string `yaml:"hosts"`
	SecretName string   `yaml:"secretName"`
}
type Spec struct {
	Rules []Rules `yaml:"rules"`
	TLS   []TLS   `yaml:"tls"`
}

type InputIngress struct {
	IngressName      string
	IngressNamespace string
	ServiceNamespace string
	ServiceName      string
	ServicePort      int
	SecretName       string
}
