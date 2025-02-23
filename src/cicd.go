// Cicd handles webhook requests from Github and rebuilds the backend service.
// Only `push` webhooks are accepted on the master branch.
package cicd

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// Validate confirms whether or not the request comes from desired Github repo.
func validate(payload []byte, hash []byte) bool {
    mac := hmac.New(sha256.New, []byte(ProgramConfig.Secret))
    mac.Write(payload)

    return hmac.Equal(mac.Sum(nil), hash)
}

// HandleWebHook checks if the request is a valid webhook. If it is
// the service is rebuilt.
func HandleWebHook(w http.ResponseWriter, r *http.Request) {
    // NOTE: Always return OK (unless build process fails). That way erroneous requests
    // won't affect the build status.
    w.WriteHeader(http.StatusOK)

    githubEvent := r.Header.Get("x-github-event")

    if githubEvent != "push" {
        msg := fmt.Sprintf("Recieved an unrecognized Github event: %s", githubEvent)
        log.Println(msg)
        fmt.Fprintln(w, msg)
        return
    }

    body, err := io.ReadAll(r.Body)
    if err != nil {
        msg := "Could not read request body!"
        log.Println(msg)
        fmt.Fprintln(w, msg)
        return
    }

    if ProgramConfig.Secret != "" {
        githubHash := r.Header.Get("X-Hub-Signature-256")

        if !validate(body, []byte(githubHash)) {
            msg := "Invalid secret!"
            log.Println(msg)
            fmt.Fprintln(w, msg)
            return
        }

    }

    var webhook WebHookRequest

    err = json.Unmarshal(body, &webhook)
    if err != nil {
        msg := "Could not parse JSON!"
        log.Println(msg)
        fmt.Fprintln(w, msg)
        return
    }

    if webhook.Ref != "refs/heads/master" {
        msg := fmt.Sprintf("Push event is not on the desired branch: %s", webhook.Ref)
        log.Println(msg)
        fmt.Fprintln(w, msg)
        return
    }

    log.Println("Webhook captured successfully!")
}
