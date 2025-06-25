## Install
```bash
go env -w GOPRIVATE=github.com/cbank1/*
go get github.com/cbank1/auth-client
```

## Example
```go
package main

import (
    "fmt"
    "github.com/cbank1/auth-client/auth"
    "github.com/golang-jwt/jwt/v5"
    "github.com/google/uuid"
    "log"
    "os"
)

type Claims struct {
    ProjectID uuid.UUID `json:"id"`
    jwt.RegisteredClaims
}

func main() {
    // init public key
    auth.InitPublicKey(os.Getenv('AUTH_PUBLIC_KEY_BASE64'))

    // user input
    user_input_token := "eyJpZCI6IjU2NjI1NjM1LTVlZGItNDk5Ny05ZWMwLWUzNGFm"

    // validate token and claims
    var claims Claims
    _, err := auth.ParseToken(user_input_token, &claims, true)
    if err != nil {
        log.Fatalf("ParseToken Error: %v", err)
    }

    fmt.Println(claims.ProjectID)
}

```