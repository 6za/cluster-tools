package controller

import (
	"fmt"

	"github.com/google/go-github/v50/github"
	"github.com/rs/zerolog/log"
)

func (g GithubSession) CreateWebhookRepo(org, repo, hookName, hookUrl, hookSecret string, hookEvents []string) error {
	input := &github.Hook{
		Name:   &hookName,
		Events: hookEvents,
		Config: map[string]interface{}{
			"content_type": "json",
			"insecure_ssl": 0,
			"url":          hookUrl,
			"secret":       hookSecret,
		},
	}

	hook, _, err := g.GitClient.Repositories.CreateHook(g.Context, org, repo, input)

	if err != nil {
		return fmt.Errorf("error when creating a webhook: %v", err)
	}

	log.Printf("Successfully created hook (id): %v", hook.GetID())

	return nil
}

func (g GithubSession) UpdateWebhook(owner, repo, hookName, hookUrl, lastHookUrl, hookSecret string, hookEvents []string) error {
	// List webhooks
	// Get webhook with name that matches, if exist delete it.
	// Create new webhook with same name
	// Return sucess or fail
	// Note to not support pagination

	hooks, _, err := g.GitClient.Repositories.ListHooks(g.Context, owner, repo, &github.ListOptions{})
	if err != nil {
		return fmt.Errorf("error when listing a webhook: %v", err)
	}
	for i, hook := range hooks {
		log.Debug().Msgf("%s, %s", i, hook)
		oldHookURL := fmt.Sprint(hook.Config["url"])
		log.Error().Msgf("Found webhook: %v", *hook)
		log.Error().Msgf("Found webhook: %v", hook)
		log.Error().Msgf("Found webhook: %v", oldHookURL)
		//name of hook seems to be no-ops all are called "Web"
		if oldHookURL == lastHookUrl {
			log.Error().Msgf("Found match!! webhook: %v", *hook)
			log.Error().Msgf("Found match!! webhook: %v", hook)
			oldHookID := *hook.ID
			log.Error().Msgf("Found match!! webhook: %d", oldHookID)
			_, err := g.GitClient.Repositories.DeleteHook(g.Context, owner, repo, oldHookID)
			if err != nil {
				log.Error().Msgf("error when removing a webhook: %v", err)
			}
			break
		}
	}
	err = g.CreateWebhookRepo(owner, repo, hookName, hookUrl, hookSecret, hookEvents)
	if err != nil {
		return fmt.Errorf("error when creating a webhook: %v", err)
	}
	return nil

}
