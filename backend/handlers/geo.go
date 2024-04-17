package api

import (
    "encoding/json"
    "net/http"

    "github.com/paulmach/orb" // If you're using this library
    "github.com/shreyanshtri26/geospatial-app/models"
    "github.com/shreyanshtri26/geospatial-app/database" 
)

func UploadGeoData(w http.ResponseWriter, r *http.Request) {
    // ... Authentication logic using AuthMiddleware ...

    var geoData models.GeoData
    err := json.NewDecoder(r.Body).Decode(&geoData)
    if err != nil {
        http.Error(w, "Invalid GeoJSON", http.StatusBadRequest)
        return
    }

    // Get UserID from authentication context
    userID := 1 // Replace with actual retrieval from auth

    geoData.UserID = userID

    // Basic validation (You'll likely want more robust validation)
    if geoData.GeoJSON == nil {
        http.Error(w, "GeoJSON data is required", http.StatusBadRequest)
        return
    }

    // Optional: If using the 'orb' library, parse into geometry
    // var geometry orb.Geometry
    // geometry, err = orb.UnmarshalJSON(geoData.GeoJSON) 

    if err := database.DB.Create(&geoData).Error; err != nil {
        http.Error(w, "Error saving GeoJSON", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    fmt.Fprintf(w, "GeoJSON uploaded successfully")
}

func GetGeoData(w http.ResponseWriter, r *http.Request) {
    // ... Authentication logic using AuthMiddleware ...

    userID := 1 // Replace with actual retrieval from auth

    var geoDataList []models.GeoData
    if err := database.DB.Where("user_id = ?", userID).Find(&geoDataList).Error; err != nil {
        http.Error(w, "Error fetching geodata", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(geoDataList)
}

