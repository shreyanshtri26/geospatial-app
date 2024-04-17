package models

import (
    "golang.org/x/crypto/bcrypt"
)

type User struct {
    ID           uint   `gorm:"primaryKey"`
    Username     string `gorm:"unique"`
    PasswordHash string 
}

// BeforeSave hook (GORM callback, triggered before saving to database)
func (u *User) BeforeSave() error {
   hash, err := bcrypt.GenerateFromPassword([]byte(u.PasswordHash), bcrypt.DefaultCost) 
    if err != nil {
        return err
    }
    u.PasswordHash = string(hash)
    return nil
}  
