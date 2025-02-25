package cicd


// EndpointConfig configures options for an endpoint.
type EndpointConfig struct {
    Name string `toml:"Name" comment:"name of the endpoint"`
    Event string `toml:"Event" comment:"which webhook event to listen for"` 
    Branch string `toml:"Branch" comment:"which branch to listen for"` 
    Secret string `toml:"Secret" comment:"optional secret to validate webhooks"` 
    Script string `toml:"Secret" comment:"path to script for webhook"`
}

// GlobalConfig configures the options for the entire program.
type GlobalConfig struct {
    Actions []EndpointConfig `toml:"Actions" comment:"actions to configure"`// Actions to configure.
    Port string `toml:"Port" comment:"port to listen on localhost"`
    ConfigFile string `toml:"ConfigFile" comment:"path to default config"`
}

// Global program config.
var ProgramConfig GlobalConfig

// WebHookRequest contains the relevant information from the Github webhook.
type WebHookRequest struct {
    Ref string `json:"Ref"` // Ref refers to the repository branch.
}
