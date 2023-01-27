package ingressgen

import (
	"fmt"
	"os"
	"time"

	"github.com/6za/cluster-tools/internal/argocd"
	"github.com/6za/cluster-tools/internal/generator"
	git "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/rs/zerolog/log"
	yaml "gopkg.in/yaml.v3"
) // with go modules disabled

func GenerateHostsLine(hostDomain string, appList []string, ip string) error {
	etcContent := ip + " "
	for _, app := range appList {
		etcContent = etcContent + " " + string(app) + "." + hostDomain

	}
	fmt.Println("Add this to th end of your /etc/hosts")
	fmt.Println(etcContent)
	return nil
}
func GenerateArtifacts(hostSuffix string, apps []string, repo string, path string) error {

	//Clone Repo
	directory := "/var/tmp/repo"
	_ = os.RemoveAll(directory)
	token := os.Getenv("GITHUB_AUTH_TOKEN")
	r, err := git.PlainClone(directory, false, &git.CloneOptions{
		URL: repo,
		Auth: &http.BasicAuth{
			Username: "abc123", // yes, this can be anything except an empty string
			Password: token,
		},
		Progress: os.Stdout,
	})
	if err != nil {
		log.Error().Err(err).Msgf("error: %s", err)
		return err
	}

	_, err = r.Head()
	if err != nil {
		log.Error().Err(err).Msgf("error: %s", err)
		return err
	}

	//Add files
	w, err := r.Worktree()
	if err != nil {
		log.Error().Err(err).Msgf("error: %s", err)
		return err
	}

	// Add Artifacts
	//Generate ingress files
	componentsFolder := "components/ingress-alt"
	filepath := directory + "/" + componentsFolder
	_ = os.RemoveAll(filepath)

	err = os.Mkdir(filepath, 0755)
	if err != nil {
		log.Error().Err(err).Msgf("error: %s", err)
		return err
	}
	filesToCommit := []string{}
	ingressList := generator.GenerateIngressesFromApps(hostSuffix, apps)
	for _, ingress := range *ingressList {
		data, err := yaml.Marshal(&ingress)
		if err != nil {
			log.Error().Err(err).Msgf("error: %s", err)
			return err
		}
		ingressFilename := ingress.Metadata.Name + ".yaml"
		ingressFile := filepath + "/" + ingress.Metadata.Name + ".yaml"
		err = os.WriteFile(ingressFile, data, 0755)
		if err != nil {
			log.Error().Err(err).Msgf("error: %s", err)
			return err
		}
		log.Info().Msgf("File Ingress:  %s", ingressFile)
		filesToCommit = append(filesToCommit, componentsFolder+"/"+ingressFilename)
	}
	applicationStruct, err := argocd.GenerateApp("ingress-alt", "ingress-alt", componentsFolder, repo)
	if err != nil {
		log.Error().Err(err).Msgf("error: %s", err)
		return err
	}

	data, err := yaml.Marshal(&applicationStruct)
	if err != nil {
		log.Error().Err(err).Msgf("error: %s", err)
		return err
	}
	appFilename := path + "/ingress-alt.yaml"
	appFile := directory + "/" + appFilename
	err = os.WriteFile(appFile, data, 0755)
	if err != nil {
		log.Error().Err(err).Msgf("error: %s", err)
		return err
	}
	log.Info().Msgf("File Ingress:  %s", appFile)
	filesToCommit = append(filesToCommit, appFilename)

	//Generate Application Yaml File
	//status, err := w.Status()

	//Workaround due to issues on the lib to add `.`
	for _, filename := range filesToCommit {
		_, err = w.Add(filename)
		if err != nil {
			log.Error().Err(err).Msgf("error: %s", err)
			return err
		}

	}
	// Check status
	_, err = w.Status()
	if err != nil {
		log.Error().Err(err).Msgf("error: %s", err)
		return err
	}

	//Commit
	_, err = w.Commit("Add ingress corrections", &git.CommitOptions{
		Author: &object.Signature{
			Name: "Cluster Utils",
			When: time.Now(),
		},
	})
	if err != nil {
		log.Error().Err(err).Msgf("error: %s", err)
		return err

	}
	//Send to main
	err = r.Push(&git.PushOptions{
		Auth: &http.BasicAuth{
			Username: "abc123", // yes, this can be anything except an empty string
			Password: token,
		},
	})
	if err != nil {
		log.Error().Err(err).Msgf("error: %s", err)
		return err

	}
	return nil
}
