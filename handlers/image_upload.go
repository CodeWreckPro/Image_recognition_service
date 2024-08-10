// handlers/image_upload.go
package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

// UploadHandler handles image uploads and saves them to a temporary directory
func UploadHandler(w http.ResponseWriter, r *http.Request) {
	// Check if the request method is POST
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Parse the form data to handle file uploads
	err := r.ParseMultipartForm(10 << 20) // Limit upload size to 10 MB
	if err != nil {
		http.Error(w, "Error parsing form data", http.StatusBadRequest)
		return
	}

	// Retrieve the file from form data
	file, handler, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "Error retrieving the file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Create a temporary file to save the uploaded image
	tempDir := "uploads"
	if _, err := os.Stat(tempDir); os.IsNotExist(err) {
		err = os.Mkdir(tempDir, 0755)
		if err != nil {
			http.Error(w, "Error creating upload directory", http.StatusInternalServerError)
			return
		}
	}

	tempFile, err := os.Create(filepath.Join(tempDir, handler.Filename))
	if err != nil {
		http.Error(w, "Error creating temporary file", http.StatusInternalServerError)
		return
	}
	defer tempFile.Close()

	// Copy the uploaded file data to the temporary file
	_, err = io.Copy(tempFile, file)
	if err != nil {
		http.Error(w, "Error saving the file", http.StatusInternalServerError)
		return
	}

	// Log the uploaded file details
	log.Printf("Uploaded file: %s (Size: %d bytes)\n", handler.Filename, handler.Size)

	// Respond with a success message
	fmt.Fprintf(w, "File uploaded successfully: %s\n", handler.Filename)
}
