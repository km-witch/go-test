package controller

import (
	"errors"
	"fmt"
	"pkg/configs"

	"github.com/golang-jwt/jwt/v4"
)

// ValidateJWT            			godoc
// @Summary      					유저 토큰 검증
// @Description  					유저의 토큰 유효성을 검증 후, 올바르다면 Claims 반환
func ValidateJWT(tokenString string) (*configs.JWTClaims, error) {
	// Claim 과 Public Sercret으로 토큰 유효성을 검증한다.
	token, err := jwt.ParseWithClaims(tokenString, &configs.JWTClaims{}, func(token *jwt.Token) (interface{}, error) {

		fmt.Println("ParseWithClaims tokenstring: ", tokenString)

		//tokenString = strings.TrimLeft(tokenString, "Bearer ")

		//fmt.Println("ParseWithClaims tokenstring parsed: ", tokenString)

		// Signing 메소드가 다르다면 실패
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}

		// Public Sercret으로 검증
		return []byte(configs.JWTSecret), nil
	})

	// 검증 실패시
	if err != nil {
		return nil, err
	}

	// 검증 성공시, 토큰 클레임을 반환한다.
	claims := token.Claims.(*configs.JWTClaims)
	return claims, nil
}
