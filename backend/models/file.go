package models
 
import "time"

type File struct {
    ID          uint      `gorm:"primaryKey"` 
    Name        string
    UserID      uint      
    UploadedAt  time.Time
    ContentType string    
}
