package cicd

import (
	"flag"
	"log"
	"os"

	"github.com/pelletier/go-toml/v2"
)

func defaultEndpoint() EndpointConfig {
    cfg := EndpointConfig {
        Name: "default",
        Event: "push",
        Branch: "refs/heads/master",
        Secret: "",
        Script: "",
    }

    return cfg
}

// SetupConfig configures the global program configuration.
func SetupConfig() {
    flag.StringVar(&ProgramConfig.Port, "port", ":7123", "Port to listen on.")
    flag.StringVar(&ProgramConfig.ConfigFile, "config", "cicd.toml", "Path to default configuration.")

    flag.Parse()

    _, err := os.Stat(ProgramConfig.ConfigFile)
    if os.IsNotExist(err) {
        cfg := defaultEndpoint()
        ProgramConfig.Actions = append(ProgramConfig.Actions, cfg)

        output, err := toml.Marshal(ProgramConfig)
        if err != nil {
            log.Fatal("Could not create TOML config!");
        }

        err = os.WriteFile(ProgramConfig.ConfigFile, output, 0644)
        if err != nil {
            log.Fatalf("Could not write to %s!", ProgramConfig.ConfigFile)
        }
    } else {
        tomlcfg, err := os.ReadFile(ProgramConfig.ConfigFile)
        if err != nil {
            log.Fatalf("Could not read %s!", ProgramConfig.ConfigFile)
        }

        err = toml.Unmarshal(tomlcfg, &ProgramConfig)
        if err != nil {
            log.Fatalf("Failed to parse TOML config from %s!", ProgramConfig.ConfigFile)
        }
    }
}
