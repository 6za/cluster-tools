package controller

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/google/go-github/v50/github"
	"github.com/rs/zerolog/log"
	"golang.org/x/oauth2"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

type GithubSession struct {
	Context     context.Context
	StaticToken oauth2.TokenSource
	OauthClient *http.Client
	GitClient   *github.Client
}

type K8sSession struct {
	Config    *rest.Config
	ClientSet *kubernetes.Clientset
}

func InitializeGithubAPI() (GithubSession, error) {
	token := os.Getenv("GITHUB_AUTH_TOKEN")
	if token == "" {
		log.Error().Msg("Missing GITHUB_AUTH_TOKEN env variable")
		return GithubSession{}, fmt.Errorf("unauthorized: no token present")
	}
	var gSession GithubSession
	gSession.Context = context.Background()
	gSession.StaticToken = oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	gSession.OauthClient = oauth2.NewClient(gSession.Context, gSession.StaticToken)
	gSession.GitClient = github.NewClient(gSession.OauthClient)
	return gSession, nil
}

func InitializeK8sInCluster() (K8sSession, error) {
	var k8sSession K8sSession
	config, err := rest.InClusterConfig()
	if err != nil {
		log.Error().Msg("Error getting Rest Config K8s")
		return K8sSession{}, err
	}
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Error().Msg("Error getting ClientSet")
		return K8sSession{}, err
	}
	k8sSession.Config = config
	k8sSession.ClientSet = clientSet
	return k8sSession, nil

}
