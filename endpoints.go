package main

import (
    "log"
    "net/http"
)

func testWebHook(w http.ResponseWriter, r *http.Request) {
    githubEvent := r.Header.Get("x-github-event")

    err := r.ParseForm()
    if err != nil {
        log.Println("Invalid HTTP body data.")
        return
    }

    if githubEvent != "push" {
        log.Println("Recieved an unrecognized Github event:", githubEvent)
        return
    }

    branch := r.PostFormValue("refs")

    if branch != "refs/heads/master" {
        log.Println("Push event is not on the desired branch: ", branch)
        return
    }

    log.Println("Webhook captured successfully!")
}
