package main

import (
    "log"
    "net/http"
    "encoding/json"
)


type WebHookRequest struct {
    Ref string `json:"Ref"`
}


func testWebHook(w http.ResponseWriter, r *http.Request) {
    githubEvent := r.Header.Get("x-github-event")

    if githubEvent != "push" {
        log.Println("Recieved an unrecognized Github event:", githubEvent)
        return
    }

    var body []byte
    var webhook WebHookRequest

    _, err := r.Body.Read(body)
    if err != nil {
        log.Println("Could not ready body data.")
        return
    }

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
