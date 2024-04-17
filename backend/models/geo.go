package models

import "github.com/paulmach/orb"

type GeoData struct {
    ID        uint      `gorm:"primaryKey"`
    UserID    uint      
    Name      string    
    GeoJSON   []byte 
    Geometry  orb.Geometry // If using the orb library 
}

