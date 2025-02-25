package main

import (
	cicd "brickedup/cicd/src"
	"log"
	"net/http"
	"strings"
)

func main() {
    cicd.SetupConfig()

    log.Printf("CI/CD running on localhost%s/", cicd.ProgramConfig.Port)

    for _, v := range cicd.ProgramConfig.Actions {
        var endpoint strings.Builder
        endpoint.WriteString("/")
        endpoint.WriteString(v.Name)
        http.HandleFunc(endpoint.String(), v.Handle)

        log.Printf("Serving %s ---> %s", endpoint.String(), v.Script)
    }

    log.Fatal(http.ListenAndServe(cicd.ProgramConfig.Port, nil))
}
