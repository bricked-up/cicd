package main


// WebHookRequest contains the relevant information from the Github webhook.
type WebHookRequest struct {
    Ref string `json:"Ref"` // Ref refers to the repository branch.
}
