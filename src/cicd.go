// Cicd handles webhook requests from Github and rebuilds the backend service.
// Only `push` webhooks are accepted on the master branch.
package cicd

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os/exec"
)

// Validate confirms whether or not the request comes from desired Github repo.
func (ep *EndpointConfig) validate(payload []byte, hash []byte) bool {
    mac := hmac.New(sha256.New, []byte(ep.Secret))
    mac.Write(payload)

    expected := hex.EncodeToString(mac.Sum(nil))

    // NOTE: Github appends sha256= prefix to the signature.
    return hmac.Equal([]byte("sha256=" + expected), hash)
}


// Handle executes an action based on the action's EndpointConfig.
func (ep *EndpointConfig) Handle(w http.ResponseWriter, r *http.Request) {
    // NOTE: Always return OK (unless build process fails). That way erroneous requests
    // won't affect the build status.
    w.WriteHeader(http.StatusOK)

    githubEvent := r.Header.Get("x-github-event")

    if ep.Event != "" && githubEvent != ep.Event {
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

    if ep.Secret != "" {
        githubHash := r.Header.Get("X-Hub-Signature-256")

        if !ep.validate(body, []byte(githubHash)) {
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

    if ep.Branch != "" && webhook.Ref != ep.Branch {
        msg := fmt.Sprintf("Push event is not on the desired branch: %s", webhook.Ref)
        log.Println(msg)
        fmt.Fprintln(w, msg)
        return
    }

    var cmd *exec.Cmd

    if ep.Script != "" {
        cmd = exec.Command(ep.Script)
    } else {
        cmd = exec.Command("echo 'Webhook processed successfully!'")
    }

    // NOTE: Since GitHub webhooks have a timeout of 10s
    // we need to execute the command in a goroutine 
    // (in case it takes a long time to run).
    go func() {
        err = cmd.Run()
        if err != nil {
            msg := fmt.Sprintf("Could not run script: %s", ep.Script)
            log.Println(msg)
        }

        log.Println("Command executed successfully!")
    }()
}
