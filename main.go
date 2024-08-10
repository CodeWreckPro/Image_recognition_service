// main.go
package main

import (
    "log"
    "net/http"

    "image-recognition-service/handlers"
)

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/upload", handlers.UploadHandler)
    mux.HandleFunc("/recognize", handlers.RecognitionHandler)

    log.Println("Server started on :8080")
    if err := http.ListenAndServe(":8080", mux); err != nil {
        log.Fatal(err)
    }
}
