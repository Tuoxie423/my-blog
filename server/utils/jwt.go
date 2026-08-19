package utils

import (
	"errors"
	"server/global"
	"server/model/request"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JWT struct {
	AccessTokenSecret  []byte
	RefreshTokenSecret []byte
}

var (
	TokenExpired     error = errors.New("token expired")
	TokenNotValidYet error = errors.New("not a active token yet")
	TokenMalformed   error = errors.New("that's not even a token") // JWT格式不正确
	TokenInvalid     error = errors.New("couldn't handle this token")
)

func NewJWT() *JWT {
	return &JWT{
		AccessTokenSecret:  []byte(global.Config.Jwt.AccessTokenSecret),
		RefreshTokenSecret: []byte(global.Config.Jwt.RefreshTokenSecret),
	}
}

func (j *JWT) CreateAccessClaims(baseClaims request.BaseClaims) request.JwtCustomClaims {
	ep, _ := ParseDuration(global.Config.Jwt.AccessTokenExpiryTime)
	claims := request.JwtCustomClaims{
		BaseClaims: baseClaims,
		RegisteredClaims: jwt.RegisteredClaims{
			Audience:  jwt.ClaimStrings{"Go blog"},
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(ep)),
			Issuer:    global.Config.Jwt.Issuer,
		},
	}
	return claims
}

func (j *JWT) CreateAccessToken(claims request.JwtCustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.AccessTokenSecret)
}

func (j *JWT) CreateRefreshClaims(baseClaims request.BaseClaims) request.JwtCustomRefreshClaims {
	ep, _ := ParseDuration(global.Config.Jwt.RefreshTokenExpiryTime)
	claims := request.JwtCustomRefreshClaims{
		UserID: baseClaims.UserID,
		RegisteredClaims: jwt.RegisteredClaims{
			Audience:  jwt.ClaimStrings{"Go blog"},
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(ep)),
			Issuer:    global.Config.Jwt.Issuer,
		},
	}
	return claims
}

func (j *JWT) CreateRefreshToken(claims request.JwtCustomRefreshClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.RefreshTokenSecret)
}

func (j *JWT) ParseAccessToken(tokenString string) (*request.JwtCustomClaims, error) {
	claims, err := j.parseToken(tokenString, &request.JwtCustomClaims{}, j.AccessTokenSecret) // 解析 Token
	if err != nil {
		return nil, err
	}
	if customClaims, ok := claims.(*request.JwtCustomClaims); ok { // 确保解析出的 Claims 类型正确
		return customClaims, nil
	}
	return nil, TokenInvalid // 如果解析结果无效，返回 TokenInvalid 错误
}

// ParseRefreshToken 解析 Refresh Token，验证 Token 并返回 Claims 信息
func (j *JWT) ParseRefreshToken(tokenString string) (*request.JwtCustomRefreshClaims, error) {
	claims, err := j.parseToken(tokenString, &request.JwtCustomRefreshClaims{}, j.RefreshTokenSecret) // 解析 Token
	if err != nil {
		return nil, err
	}
	if refreshClaims, ok := claims.(*request.JwtCustomRefreshClaims); ok { // 确保解析出的 Claims 类型正确
		return refreshClaims, nil
	}
	return nil, TokenInvalid // 如果解析结果无效，返回 TokenInvalid 错误
}

func (j *JWT) parseToken(tokenString string, claims jwt.Claims, secretKey interface{}) (interface{}, error) {
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil // 返回密钥以验证 Token
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok { // 处理 Token 验证错误
			switch {
			case ve.Errors&jwt.ValidationErrorMalformed != 0:
				return nil, TokenMalformed // Token 格式错误
			case ve.Errors&jwt.ValidationErrorExpired != 0:
				return nil, TokenExpired // Token 已过期
			case ve.Errors&jwt.ValidationErrorNotValidYet != 0:
				return nil, TokenNotValidYet // Token 还未生效
			default:
				return nil, TokenInvalid // 其他错误返回 Token 无效
			}
		}
		return nil, TokenInvalid // 默认返回 Token 无效错误
	}

	if token.Valid { // 如果 Token 验证通过，返回 Claims
		return token.Claims, nil
	}
	return nil, TokenInvalid // Token 无效，返回错误
}
