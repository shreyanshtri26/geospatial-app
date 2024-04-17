package api

import (
    "encoding/json"
    "net/http"

    "github.com/shreyanshtri26/geospatial-app/models"
    "github.com/shreyanshtri26/geospatial-app/database" 
    "golang.org/x/crypto/bcrypt"
    "github.com/dgrijalva/jwt-go" // For JWT functionality
    "time"
)

// Replace these with your JWT logic specifics
var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

type Credentials struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

type Claims struct {
    UserID uint `json:"userId"`
    jwt.StandardClaims
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
    var newUser models.User
    err := json.NewDecoder(r.Body).Decode(&newUser)
    if err != nil {
        http.Error(w, "Invalid JSON", http.StatusBadRequest)
        return
    }

    err = newUser.BeforeSave() // Hash password
    if err != nil {
        http.Error(w, "Error hashing password", http.StatusInternalServerError)
        return
    }

    if err := database.DB.Create(&newUser).Error; err != nil {
        http.Error(w, "Error creating user", http.StatusInternalServerError)
        return
    }

    // Generate JWT
    token, err := generateJWT(newUser.ID)
    if err != nil {
        http.Error(w, "Error generating JWT", http.StatusInternalServerError)
        return
    }
    http.SetCookie(w, &http.Cookie{Name: "token", Value: token, HttpOnly: true}) 
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
    var creds Credentials
    err := json.NewDecoder(r.Body).Decode(&creds)
    if err != nil {
        http.Error(w, "Invalid JSON", http.StatusBadRequest)
        return
    }

    var user models.User
    database.DB.Where("username = ?", creds.Username).First(&user) 

    err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(creds.Password))
    if err != nil { 
        http.Error(w, "Invalid credentials", http.StatusUnauthorized)
        return
    }

    // Generate JWT
    token, err := generateJWT(user.ID)
    if err != nil {
        http.Error(w, "Error generating JWT", http.StatusInternalServerError)
        return
    }
    http.SetCookie(w, &http.Cookie{Name: "token", Value: token, HttpOnly: true})  
}

// AuthMiddleware (replace with JWT validation)
func AuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

        // Extract and validate JWT ... 

        next.ServeHTTP(w, r) 
    })
}

// Helper to generate JWT 
func generateJWT(userID uint) (string, error) {
      expirationTime := time.Now().Add(24 * time.Hour) // Example expiration

    claims := Claims{
        UserID: userID,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: expirationTime.Unix(),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(jwtSecret)
}
