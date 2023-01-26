package generator

import (
	"github.com/rs/zerolog/log"
)

func BlankIngressPayload() *IngressKind {
	return &IngressKind{
		APIVersion: "networking.k8s.io/v1",
		Kind:       "Ingress",
	}
}

func PopulateIngressPayloadWithTLS(hostname string, input *InputIngress) *IngressKind {
	payload := PopulateIngressPayloadWithoutTLS(hostname, input)
	payload.Spec.TLS = []TLS{
		{
			Hosts:      []string{hostname},
			SecretName: input.SecretName,
		},
	}
	return payload
}

func PopulateIngressPayloadWithoutTLS(hostname string, input *InputIngress) *IngressKind {
	payload := BlankIngressPayload()
	payload.Metadata = Metadata{Name: input.IngressName, Namespace: input.IngressNamespace}

	path := Paths{
		Backend:  Backend{Service: Service{Name: input.ServiceName, Port: Port{Number: input.ServicePort}}},
		Path:     "/",
		PathType: "Prefix",
	}

	payload.Spec.Rules = []Rules{
		{
			Host: hostname,
			HTTP: HTTP{
				Paths: []Paths{path},
			},
		},
	}

	return payload
}

func GenerateIngressesFromApps(hostnameSuffix string, apps []string) *[]IngressKind {
	ingressList := make([]IngressKind, 0)
	for _, app := range apps {
		log.Info().Msgf("Generating Ingress for:  %s ", app)
		switch app {
		case "minio":
			ingressList = append(ingressList, *PopulateIngressPayloadWithTLS(app+"."+hostnameSuffix, &MinioIngress))
			log.Info().Msg("Generating Ingress for MinioIngress")
		case "minio-console":
			ingressList = append(ingressList, *PopulateIngressPayloadWithTLS(app+"."+hostnameSuffix, &MinioConsoleIngress))
			log.Info().Msg("Generating Ingress for MinioConsoleIngress")
		case "vault":
			ingressList = append(ingressList, *PopulateIngressPayloadWithTLS(app+"."+hostnameSuffix, &VaultIngress))
			log.Info().Msg("Generating Ingress for VaultIngress")
		case "atlantis":
			ingressList = append(ingressList, *PopulateIngressPayloadWithTLS(app+"."+hostnameSuffix, &AtlantisIngress))
			log.Info().Msg("Generating Ingress for AtlantisIngress")
		case "chartmuseum":
			ingressList = append(ingressList, *PopulateIngressPayloadWithTLS(app+"."+hostnameSuffix, &ChartMuseumIngress))
			log.Info().Msg("Generating Ingress for ChartMuseumIngress")
		case "argo":
			ingressList = append(ingressList, *PopulateIngressPayloadWithTLS(app+"."+hostnameSuffix, &ArgoIngress))
			log.Info().Msg("Generating Ingress for ArgoIngress")
		case "kubefirst":
			ingressList = append(ingressList, *PopulateIngressPayloadWithTLS(app+"."+hostnameSuffix, &KubefirstIngress))
			log.Info().Msg("Generating Ingress for KubefirstIngress")
		case "argocd":
			ingressList = append(ingressList, *PopulateIngressPayloadWithoutTLS(app+"."+hostnameSuffix, &ArgoCDIngress))
			log.Info().Msg("Generating Ingress for ArgoCDIngress")
		}

	}
	return &ingressList
}
