package main

import (
    "log"
    "net/http"
    "os"

    "github.com/gorilla/mux"
    "github.com/joho/godotenv"
    "github.com/shreyanshtri26/geospatial-app/api"
    "github.com/shreyanshtri26/geospatial-app/database"
)

func main() {
    // Load environment variables (.env file)
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    // Initialize database connection
    db := database.InitDB()
    defer db.Close() // Close connection on program exit

    // Initialize router
    router := mux.NewRouter()

   // Authentication routes
    router.HandleFunc("/api/user/register", api.RegisterHandler).Methods("POST")
    router.HandleFunc("/api/user/login", api.LoginHandler).Methods("POST")

    // Geospatial routes (with AuthMiddleware)
    geoGroup := router.PathPrefix("/api/geo").Subrouter() 
    geoGroup.Use(api.AuthMiddleware) 
    geoGroup.HandleFunc("/upload", api.UploadGeoData).Methods("POST")
    geoGroup.HandleFunc("/", api.GetGeoData).Methods("GET") 

    // File routes (with AuthMiddleware)
    fileGroup := router.PathPrefix("/api/file").Subrouter()
    fileGroup.Use(api.AuthMiddleware)
    fileGroup.HandleFunc("/upload", api.UploadFileHandler).Methods("POST")

   // Start Server
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080" 
    }
    log.Printf("Server starting on port %s", port)
    log.Fatal(http.ListenAndServe(":"+port, router))
}

