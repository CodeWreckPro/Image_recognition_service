// handlers/image_recognition.go
package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
)

// RecognitionHandler handles image recognition requests
func RecognitionHandler(w http.ResponseWriter, r *http.Request) {
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
	file, _, err := r.FormFile("image")
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

	tempFile, err := os.CreateTemp(tempDir, "recognize-*.png")
	if err != nil {
		http.Error(w, "Error creating temporary file", http.StatusInternalServerError)
		return
	}
	defer os.Remove(tempFile.Name()) // Clean up the file after processing
	defer tempFile.Close()

	// Copy the uploaded file data to the temporary file
	_, err = io.Copy(tempFile, file)
	if err != nil {
		http.Error(w, "Error saving the file", http.StatusInternalServerError)
		return
	}

	// Call Python script to perform image recognition
	result, err := recognizeImage(tempFile.Name())
	if err != nil {
		http.Error(w, "Error recognizing image", http.StatusInternalServerError)
		return
	}

	// Respond with the recognition result
	fmt.Fprintf(w, "Recognition result: %s\n", result)
}

// recognizeImage calls a Python script to recognize the image and returns the result
func recognizeImage(filePath string) (string, error) {
	// Command to execute the Python script
	cmd := exec.Command("python3", "model/model.py", filePath)

	// Run the command and capture the output
	output, err := cmd.Output()
	if err != nil {
		log.Printf("Error executing Python script: %v\n", err)
		return "", err
	}

	return string(output), nil
}
