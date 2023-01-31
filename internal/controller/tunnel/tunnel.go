package tunnel

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/6za/cluster-tools/internal/controller"
	"github.com/rs/zerolog/log"
	v1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const IN_CLUSTER_ATLANTIS_ADDRESS = "http://ngrok-agent.ngrok-agent.svc.cluster.local:4040/api/tunnels"
const LOOP_WAIT = 5 * time.Second

func ExecuteTunnel(owner, repo string) error {
	//while true
	//Call Some SDK to validate NGROK
	//CURL: curl http://localhost:3000/api/tunnels
	//Check if previous tunnel is the same of the new tunnel public address
	//if matches, do nothing - sleep/exit
	//if don't matches
	//USE GITHUB token from ENV Variable(it will come from vault later)
	// USe token to create a webhook
	// USe token to remove  old webhook
	//Where old webhook is stored?
	//  - configMap?
	//  - can it be queried by some label on github?
	// - start witha pre-defined name

	gitHubClient, err := controller.InitializeGithubAPI()
	if err != nil {
		return fmt.Errorf("error when connecting to github: %v", err)
	}

	k8sSession, err := controller.InitializeK8sInCluster()
	if err != nil {
		return fmt.Errorf("error when connecting to k8s: %v", err)
	}
	clientset := k8sSession.ClientSet

	atlantisSecretClient := clientset.CoreV1().Secrets("atlantis")
	ngrokConfigMapClient := clientset.CoreV1().ConfigMaps("ngrok-agent")
	for {
		time.Sleep(LOOP_WAIT)
		// lastTunnel
		lastTunnel := ""
		ngrokConfig, err := ngrokConfigMapClient.Get(context.TODO(), "ngrok-agent-config", metaV1.GetOptions{})
		if err != nil {
			//How to check if it is not found
			log.Warn().Err(err).Msgf("error getting ngrok-agent-config")
		} else {
			lastTunnel = string(ngrokConfig.Data["LAST_TUNNEL"])
		}

		// Read from configMap
		payload, err := CheckNgrokTunnel()
		if err != nil {
			log.Warn().Msgf("error checking status of tunnel: %v", err)
			// We will try again soon, once cluster is more ready
			continue
		}
		fmt.Printf("%v", payload)
		if len(payload.Tunnels) < 1 {
			log.Warn().Msg("error reading tunnel info:  no tunnels")
			// We will try again soon, once cluster is more ready
			continue
		}

		hookName := "ngrok_atlantis"
		hookURL := payload.Tunnels[0].PublicURL + "/events"
		if hookURL == lastTunnel {
			// Nothing to be done
			continue
		}

		secret, err := atlantisSecretClient.Get(context.TODO(), "atlantis-secrets", metaV1.GetOptions{})
		if err != nil {
			log.Error().Err(err).Msgf("error getting key")
		}
		hookSecret := string(secret.Data["ATLANTIS_GH_WEBHOOK_SECRET"])
		hookEvents := []string{"issue_comment", "pull_request", "pull_request_review", "push"}
		err = gitHubClient.UpdateWebhook(owner, repo, hookName, hookURL, lastTunnel, hookSecret, hookEvents)
		if err != nil {
			log.Error().Msgf("error when updating a webhook: %v", err)
			return fmt.Errorf("error when updating a webhook: %v", err)
		}
		lastTunnel = hookURL
		m := make(map[string]string)
		m["LAST_TUNNEL"] = hookURL
		configMap := v1.ConfigMap{
			ObjectMeta: metaV1.ObjectMeta{
				Name:      "ngrok-agent-config",
				Namespace: "ngrok-agent",
			},
			Data: m,
		}
		err = ngrokConfigMapClient.Delete(context.TODO(), "ngrok-agent-config", metaV1.DeleteOptions{})
		if err != nil {
			log.Warn().Msgf("error deleting configmap ngrok-agent-config: %v", err)
		}

		_, err = ngrokConfigMapClient.Create(context.TODO(), &configMap, metaV1.CreateOptions{})
		if err != nil {
			log.Warn().Msgf("error creating configmap ngrok-agent-config: %v", err)
		}

	}
	return nil
}

type NgrokTunnel struct {
	Tunnels []struct {
		Name      string `json:"name"`
		ID        string `json:"ID"`
		URI       string `json:"uri"`
		PublicURL string `json:"public_url"`
		Proto     string `json:"proto"`
		Config    struct {
			Addr    string `json:"addr"`
			Inspect bool   `json:"inspect"`
		} `json:"config"`
	} `json:"tunnels"`
	URI string `json:"uri"`
}

func CheckNgrokTunnel() (*NgrokTunnel, error) {

	url := IN_CLUSTER_ATLANTIS_ADDRESS

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Error().Msgf("%s", err)
		return &NgrokTunnel{}, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Error().Msgf("%s", err)
		return &NgrokTunnel{}, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Error().Msgf("%s", err)
		return &NgrokTunnel{}, err
	}

	var payload NgrokTunnel
	err = json.Unmarshal(body, &payload)
	if err != nil {
		log.Error().Msgf("%s", err)
	}
	return &payload, nil
}
