package main

import (
	"net/http"
	"gopkg.in/go-playground/webhooks.v5/bitbucket"
	"gopkg.in/go-playground/webhooks.v5/docker"
	"gopkg.in/go-playground/webhooks.v5/github"
	"gopkg.in/go-playground/webhooks.v5/gitlab"

	log "github.com/sirupsen/logrus"
)

func BitBucketWebhookHandler(w http.ResponseWriter, r *http.Request) {
	hook, _ := bitbucket.New()
	payload, err := hook.Parse(r, bitbucket.PullRequestCreatedEvent, bitbucket.RepoPushEvent)
	if err != nil {
		if err == bitbucket.ErrEventNotFound {
			// ok event wasn't one of the ones asked to be parsed
		}
	}
	switch payload.(type) {
	case bitbucket.PullRequest: 
		pullRequest := payload.(bitbucket.PullRequest)
		// Do whatever you want from here...
		log.WithFields(log.Fields{"Webhook": "Bitbucket", "PullRequest": pullRequest }).Info("Order:Webhook")
	case bitbucket.RepoPushPayload:
		push := payload.(bitbucket.RepoPushPayload)
		// Do whatever you want from here...
		log.WithFields(log.Fields{"Webhook": "Bitbucket", "Push": push }).Info("Order:Webhook")
	}
}

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
		log.WithFields(log.Fields{"Webhook": "Dockerhub", "Building": build, "Repository": build.Repository }).Info("Order:Webhook")
	}
}

func GithubWebhookHandler(w http.ResponseWriter, r *http.Request) {
	hook, _ := github.New(github.Options.Secret(Conf.GithubWebhookSecret))

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
			log.WithFields(log.Fields{"Webhook": "Github", "PullRequest": pullRequest }).Info("Order:Webhook")
		case github.PushPayload:
			push := payload.(github.PushPayload)
			// Do whatever you want from here...
			log.WithFields(log.Fields{"Webhook": "Github", "Push": push }).Info("Order:Webhook")
	}
}

func gitlabWebhookHandler(w http.ResponseWriter, r *http.Request) {
	hook, _ := gitlab.New(gitlab.Options.Secret(Conf.GithubWebhookSecret))

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
			log.WithFields(log.Fields{"Webhook": "Gitlab", "PullRequest": pullRequest }).Info("Order:Webhook")
		case gitlab.MergeRequestEventPayload:
			push := payload.(gitlab.PushEventPayload)
			// Do whatever you want from here...
			log.WithFields(log.Fields{"Webhook": "Gitlab", "Push": push }).Info("Order:Webhook")
	}
}