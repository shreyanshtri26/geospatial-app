package api

import (
    "fmt"
    "io"
    "net/http"
    "os"
    "strconv"

    "github.com/shreyanshtri26/geospatial-app/models"
    "github.com/shreyanshtri26/geospatial-app/database" 
)

func UploadFileHandler(w http.ResponseWriter, r *http.Request) {
    // ... Authentication logic using AuthMiddleware ...

    // Parse multipart form (limit 10MB)
    err := r.ParseMultipartForm(10 << 20) 
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Retrieve the file from the form data
    file, handler, err := r.FormFile("file")
    if err != nil {
        http.Error(w, "Error retrieving file", http.StatusBadRequest)
        return
    }
    defer file.Close()

    // Create a new file on the server 
    f, err := os.OpenFile(handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
    if err != nil {
        http.Error(w, "Error creating file", http.StatusInternalServerError)
        return
    }
    defer f.Close()

    // Copy the file data
    _, err = io.Copy(f, file)
    if err != nil {
        http.Error(w, "Error copying file", http.StatusInternalServerError)
        return
    }

    // Get UserID from authentication context
    userID := 1 // Replace with actual retrieval from auth

    // Create File record in database
    newFile := models.File{
        Name:       handler.Filename,
        UserID:     userID,
        UploadedAt: time.Now(),
    }
    if err := database.DB.Create(&newFile).Error; err != nil {
        http.Error(w, "Error saving file to database", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    fmt.Fprintf(w, "File uploaded successfully")
}
