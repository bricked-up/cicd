package main

import (
    "log"
    "net/http"
)

const PORT = ":3000"

func main() {
    http.HandleFunc("/", handleWebHook)

    log.Printf("CI/CD running on localhost%s/", PORT)
    log.Fatal(http.ListenAndServe(PORT, nil))
}
