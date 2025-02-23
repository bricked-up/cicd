package main

import (
    "fmt"
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
        msg := fmt.Sprintf("Recieved an unrecognized Github event: %s\n", githubEvent)
        log.Println(msg)
        fmt.Fprint(w, msg)
        return
    }

    body, err := io.ReadAll(r.Body)
    if err != nil {
        msg := "Could not read request body!"
        log.Println(msg)
        fmt.Fprint(w, msg)
        return
    }

    var webhook WebHookRequest

    err = json.Unmarshal(body, &webhook)
    if err != nil {
        msg := "Could not parse JSON!"
        log.Println(msg)
        fmt.Fprint(w, msg)
        return
    }

    if webhook.Ref != "refs/heads/master" {
        msg := fmt.Sprintf("Push event is not on the desired branch: %s\n", webhook.Ref)
        log.Println(msg)
        fmt.Fprint(w, msg)
        return
    }

    log.Println("Webhook captured successfully!")
}
