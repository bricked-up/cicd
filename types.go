package main


// WebHookRequest contains the relevant information from the Github webhook.
type WebHookRequest struct {
    Ref string `json:"Ref"` // Ref refers to the repository branch.
}

const (
    ServerOk = 200 // ServerOk is for when the request was processed successfully.
    InternalServerError = 500 // InternalServerError is for when deployment failed.
)

