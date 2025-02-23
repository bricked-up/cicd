package main

import (
	cicd "brickedup/cicd/src"
	"log"
	"net/http"
)

func main() {
    cicd.SetupConfig()

    http.HandleFunc("/", cicd.HandleWebHook)

    log.Printf("CI/CD running on localhost%s/", cicd.ProgramConfig.Port)
    log.Fatal(http.ListenAndServe(cicd.ProgramConfig.Port, nil))
}
