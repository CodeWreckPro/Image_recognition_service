// utils/utils.go
package utils

import (
    "bytes"
    "os/exec"
)

// RecognizeImage uses Python script to perform image recognition
func RecognizeImage(filePath string) (string, error) {
    cmd := exec.Command("python3", "model/model.py", filePath)
    var out bytes.Buffer
    cmd.Stdout = &out
    err := cmd.Run()
    if err != nil {
        return "", err
    }
    return out.String(), nil
}
