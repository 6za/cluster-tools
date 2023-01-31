module github.com/6za/cluster-tools

go 1.18

require (
	github.com/go-git/go-git/v5 v5.5.2
	github.com/google/go-github/v50 v50.0.0
	github.com/rs/zerolog v1.29.0
	github.com/spf13/cobra v1.6.1
	golang.org/x/oauth2 v0.0.0-20211104180415-d3ed0bb246c8
	gopkg.in/yaml.v3 v3.0.1
	k8s.io/apimachinery v0.23.16
	k8s.io/client-go v0.0.0
)

require (
	github.com/Microsoft/go-winio v0.5.2 // indirect
	github.com/ProtonMail/go-crypto v0.0.0-20221026131551-cf6655e29de4 // indirect
	github.com/acomagu/bufpipe v1.0.3 // indirect
	github.com/cloudflare/circl v1.1.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/emirpasic/gods v1.18.1 // indirect
	github.com/go-git/gcfg v1.5.0 // indirect
	github.com/go-git/go-billy/v5 v5.4.0 // indirect
	github.com/go-logr/logr v1.2.0 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/go-cmp v0.5.9 // indirect
	github.com/google/go-querystring v1.1.0 // indirect
	github.com/google/gofuzz v1.1.0 // indirect
	github.com/googleapis/gnostic v0.5.5 // indirect
	github.com/imdario/mergo v0.3.13 // indirect
	github.com/inconshreveable/mousetrap v1.0.1 // indirect
	github.com/jbenet/go-context v0.0.0-20150711004518-d14ea06fba99 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/kevinburke/ssh_config v1.2.0 // indirect
	github.com/mattn/go-colorable v0.1.12 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pjbgf/sha1cd v0.2.3 // indirect
	github.com/sergi/go-diff v1.1.0 // indirect
	github.com/skeema/knownhosts v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/xanzy/ssh-agent v0.3.3 // indirect
	golang.org/x/crypto v0.3.0 // indirect
	golang.org/x/net v0.3.1-0.20221206200815-1e63c2f08a10 // indirect
	golang.org/x/sys v0.3.0 // indirect
	golang.org/x/term v0.3.0 // indirect
	golang.org/x/text v0.5.0 // indirect
	golang.org/x/time v0.0.0-20220210224613-90d013bbcef8 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/protobuf v1.27.1 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	gopkg.in/tomb.v1 v1.0.0-20141024135613-dd632973f1e7 // indirect
	gopkg.in/warnings.v0 v0.1.2 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	k8s.io/api v0.23.16 // indirect
	k8s.io/klog/v2 v2.30.0 // indirect
	k8s.io/kube-openapi v0.0.0-20211115234752-e816edb12b65 // indirect
	k8s.io/utils v0.0.0-20211116205334-6203023598ed // indirect
	sigs.k8s.io/json v0.0.0-20211020170558-c049b76a60c6 // indirect
	sigs.k8s.io/structured-merge-diff/v4 v4.2.3 // indirect
	sigs.k8s.io/yaml v1.2.0 // indirect

)

replace k8s.io/api => k8s.io/api v0.23.16

replace k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.23.16

replace k8s.io/apimachinery => k8s.io/apimachinery v0.23.17-rc.0

replace k8s.io/apiserver => k8s.io/apiserver v0.23.16

replace k8s.io/cli-runtime => k8s.io/cli-runtime v0.23.16

replace k8s.io/client-go => k8s.io/client-go v0.23.16

replace k8s.io/cloud-provider => k8s.io/cloud-provider v0.23.16

replace k8s.io/cluster-bootstrap => k8s.io/cluster-bootstrap v0.23.16

replace k8s.io/code-generator => k8s.io/code-generator v0.23.17-rc.0

replace k8s.io/component-base => k8s.io/component-base v0.23.16

replace k8s.io/component-helpers => k8s.io/component-helpers v0.23.16

replace k8s.io/controller-manager => k8s.io/controller-manager v0.23.16

replace k8s.io/cri-api => k8s.io/cri-api v0.23.17-rc.0

replace k8s.io/csi-translation-lib => k8s.io/csi-translation-lib v0.23.16

replace k8s.io/kube-aggregator => k8s.io/kube-aggregator v0.23.16

replace k8s.io/kube-controller-manager => k8s.io/kube-controller-manager v0.23.16

replace k8s.io/kube-proxy => k8s.io/kube-proxy v0.23.16

replace k8s.io/kube-scheduler => k8s.io/kube-scheduler v0.23.16

replace k8s.io/kubectl => k8s.io/kubectl v0.23.16

replace k8s.io/kubelet => k8s.io/kubelet v0.23.16

replace k8s.io/legacy-cloud-providers => k8s.io/legacy-cloud-providers v0.23.16

replace k8s.io/metrics => k8s.io/metrics v0.23.16

replace k8s.io/mount-utils => k8s.io/mount-utils v0.23.17-rc.0

replace k8s.io/pod-security-admission => k8s.io/pod-security-admission v0.23.16

replace k8s.io/sample-apiserver => k8s.io/sample-apiserver v0.23.16

replace k8s.io/sample-cli-plugin => k8s.io/sample-cli-plugin v0.23.16

replace k8s.io/sample-controller => k8s.io/sample-controller v0.23.16
