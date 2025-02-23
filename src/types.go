package cicd

// Global config for program.
type Config struct {
    Secret string // Secret through which to validate Gihub webhooks.
    Port string // Port to listen to on localhost.
}

// Global program config.
var ProgramConfig Config

// WebHookRequest contains the relevant information from the Github webhook.
type WebHookRequest struct {
    Ref string `json:"Ref"` // Ref refers to the repository branch.
}
