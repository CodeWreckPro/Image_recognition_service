// handlers/handlers.go
package handlers

import (
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "os"

    "image-recognition-service/utils"
)

// UploadHandler handles image uploads
func UploadHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }

    file, header, err := r.FormFile("image")
    if err != nil {
        http.Error(w, "Error reading image", http.StatusBadRequest)
        return
    }
    defer file.Close()

    tempFile, err := ioutil.TempFile("uploads", "upload-*.png")
    if err != nil {
        http.Error(w, "Error creating temporary file", http.StatusInternalServerError)
        return
    }
    defer tempFile.Close()

    fileBytes, err := ioutil.ReadAll(file)
    if err != nil {
        http.Error(w, "Error reading file bytes", http.StatusInternalServerError)
        return
    }

    tempFile.Write(fileBytes)
    log.Printf("Uploaded file: %s", header.Filename)

    fmt.Fprintf(w, "File uploaded successfully: %s\n", header.Filename)
}

// RecognitionHandler handles image recognition requests
func RecognitionHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }

    file, _, err := r.FormFile("image")
    if err != nil {
        http.Error(w, "Error reading image", http.StatusBadRequest)
        return
    }
    defer file.Close()

    tempFile, err := ioutil.TempFile("uploads", "recognize-*.png")
    if err != nil {
        http.Error(w, "Error creating temporary file", http.StatusInternalServerError)
        return
    }
    defer os.Remove(tempFile.Name())

    fileBytes, err := ioutil.ReadAll(file)
    if err != nil {
        http.Error(w, "Error reading file bytes", http.StatusInternalServerError)
        return
    }

    tempFile.Write(fileBytes)

    // Call Python script to recognize image
    result, err := utils.RecognizeImage(tempFile.Name())
    if err != nil {
        http.Error(w, "Error recognizing image", http.StatusInternalServerError)
        return
    }

    fmt.Fprintf(w, "Recognition result: %s\n", result)
}
