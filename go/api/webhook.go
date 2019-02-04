package api

import (
	"net/http"
	"os"

	"gopkg.in/go-playground/webhooks.v5/docker"
	"gopkg.in/go-playground/webhooks.v5/github"
	"gopkg.in/go-playground/webhooks.v5/gitlab"

	log "github.com/sirupsen/logrus"
)

// DockerhubWebhookHandler func
func DockerhubWebhookHandler(w http.ResponseWriter, r *http.Request) {
	hook, _ := docker.New()
	payload, err := hook.Parse(r, docker.BuildEvent)
	if err != nil {
		if err == docker.ErrParsingPayload {
			// ok event wasn't one of the ones asked to be parsed
		}
	}
	switch payload.(type) {
	case docker.BuildPayload:
		build := payload.(docker.BuildPayload)
		// Do whatever you want from here...
		log.WithFields(log.Fields{"Webhook": "Dockerhub", "Building": build, "Repository": build.Repository}).Info("Order:Webhook")
	}
}

// GithubWebhookHandler func
func GithubWebhookHandler(w http.ResponseWriter, r *http.Request) {
	hook, _ := github.New(github.Options.Secret(os.Getenv("NWN_ORDER_GITHUB_WEBHOOK_SECRET")))
	payload, err := hook.Parse(r, github.PullRequestEvent, github.PushEvent)
	if err != nil {
		if err == github.ErrEventNotFound {
			// ok event wasn't one of the ones asked to be parsed
		}
	}
	switch payload.(type) {
	case github.PullRequestPayload:
		pullRequest := payload.(github.PullRequestPayload)
		// Do whatever you want from here...
		log.WithFields(log.Fields{"Webhook": "Github", "PullRequest": pullRequest}).Info("Order:Webhook")
	case github.PushPayload:
		push := payload.(github.PushPayload)
		// Do whatever you want from here...
		log.WithFields(log.Fields{"Webhook": "Github", "Push": push}).Info("Order:Webhook")
	}
}

// GitlabWebhookHandler func
func GitlabWebhookHandler(w http.ResponseWriter, r *http.Request) {
	hook, _ := gitlab.New(gitlab.Options.Secret(os.Getenv("NWN_ORDER_GITLAB_WEBHOOK_SECRET")))

	payload, err := hook.Parse(r, gitlab.PushEvents, gitlab.MergeRequestEvents)
	if err != nil {
		if err == gitlab.ErrEventNotFound {
			// ok event wasn't one of the ones asked to be parsed
		}
	}
	switch payload.(type) {
	case gitlab.PushEventPayload:
		pullRequest := payload.(gitlab.PushEventPayload)
		// Do whatever you want from here...
		log.WithFields(log.Fields{"Webhook": "Gitlab", "PullRequest": pullRequest}).Info("Order:Webhook")
	case gitlab.MergeRequestEventPayload:
		push := payload.(gitlab.PushEventPayload)
		// Do whatever you want from here...
		log.WithFields(log.Fields{"Webhook": "Gitlab", "Push": push}).Info("Order:Webhook")
	}
}
