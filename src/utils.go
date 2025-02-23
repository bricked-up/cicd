package cicd

import "flag"

// SetupConfig configures the global program configuration.
func SetupConfig() {
    flag.StringVar(&ProgramConfig.Secret, "secret", "", "Secret to authenticate webhhook request.")
    flag.StringVar(&ProgramConfig.Port, "port", ":7123", "Port to listen on.")

    flag.Parse()
}
