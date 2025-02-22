package main

import (
    "log"
    "net/http"
    "encoding/json"
    "strings"
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

    var buf []byte = make([]byte, 1024)
    var body strings.Builder
    var webhook WebHookRequest

    for {
        bytes_read, err := r.Body.Read(buf)
        if err != nil {
            log.Println("Could not ready body data.")
            return
        }

        body.WriteString(string(buf))

        if bytes_read == 0 {
            break
        }
    }

    err := json.Unmarshal([]byte(body.String()), &webhook)
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
