package main

import (
    "log"
    "net/http"
    "encoding/json"
    "io"
)


// HandleWebHook checks if the request is a valid webhook. If it is
// the service is rebuilt.
func handleWebHook(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)

    githubEvent := r.Header.Get("x-github-event")

    if githubEvent != "push" {
        log.Println("Recieved an unrecognized Github event:", githubEvent)
        return
    }

    body, err := io.ReadAll(r.Body)
    if err != nil {
        log.Println("Could not read request body!")
        return
    }

    var webhook WebHookRequest

    err = json.Unmarshal(body, &webhook)
    if err != nil {
        log.Println("Could not parse JSON!")
        return
    }

    if webhook.Ref != "refs/heads/master" {
        log.Println("Push event is not on the desired branch: ", webhook.Ref)
        return
    }

    log.Println("Webhook captured successfully!")
}
