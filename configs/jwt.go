package configs

import "github.com/golang-jwt/jwt/v4"

type JWTClaims struct {
	UID                string `json:"uid"`      // 유저 ID
	PreLogin           string `json:"PreLogin"` // 사용안함
	Admin              string `json:"admin"`    // 사용안함
	User               string `json:"user"`     // 사용안함
	jwt.StandardClaims        // 표준 토큰 Claims
}

const JWTSecret = "a(sl*jk@f!#sdkl)|9./:[]{}a0<>?;12^3%&,234sa$~f18sd"
