package utils

import (
    "errors"
    "github.com/golang-jwt/jwt/v4"
    "github.com/joho/godotenv"
    "os"
    "time"
)

// Load the secret key from .env
func getSecretKey() (string, error) {
    if err := godotenv.Load(); err != nil {
        return "", errors.New("failed to load .env file")
    }
    secret := os.Getenv("JWT_SECRET")
    if secret == "" {
        return "", errors.New("JWT_SECRET is not set in .env file")
    }
    return secret, nil
}

// GenerateJWT generates a JWT token for a given user ID
func GenerateJWT(userID int) (string, error) {
    secretKey, err := getSecretKey()
    if err != nil {
        return "", err
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "user_id": userID,
        "exp":     time.Now().Add(24 * time.Hour).Unix(),
    })

    return token.SignedString([]byte(secretKey))
}

// VerifyJWT verifies the token and extracts the user ID
func VerifyJWT(tokenStr string) (int, error) {
    secretKey, err := getSecretKey()
    if err != nil {
        return 0, err
    }

    token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
        return []byte(secretKey), nil
    })

    if err != nil {
        if errors.Is(err, jwt.ErrTokenExpired) {
            return 0, errors.New("token has expired")
        }
        return 0, err
    }

    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
        userID, ok := claims["user_id"].(float64)
        if !ok {
            return 0, errors.New("invalid user_id claim type")
        }
        return int(userID), nil
    }

    return 0, errors.New("invalid token")
}
