package jwtauth

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"go.uber.org/zap"
	"micro-shop-api/userop-web/global"
	"time"
)

// CustomClaims 自定义 Payload 信息
type CustomClaims struct {
	Id       uint   // 用户id
	Mobile   string // 手机号
	Nickname string // 用户昵称
	jwt.StandardClaims
}

func NewCustomClaimsDefault(id uint, mobile string, nickname string) *CustomClaims {
	beforeTime := time.Now().Unix()
	return &CustomClaims{
		Id:       id,
		Mobile:   mobile,
		Nickname: nickname,
		StandardClaims: jwt.StandardClaims{
			NotBefore: beforeTime,
			ExpiresAt: beforeTime + 60*60*24,
			Issuer:    "lzscxb",
		},
	}
}

type JWT struct {
	singKey []byte
}

var (
	TokenExpired     = errors.New("Token is expired")        // 令牌过期
	TokenNotValidYet = errors.New("Token not active yet")    // 令牌未生效
	TokenMalformed   = errors.New("that's not even a token") // 令牌不完整
	TokenInvalid     = errors.New("")                        // 无效令牌
)

// NewJWT 返回一个JWT
func NewJWT() *JWT {
	return &JWT{
		singKey: []byte(global.Config.JwtInfo.SingKey),
	}
}

// CreateToken 创建新的token
func (j *JWT) CreateToken(claims CustomClaims) (token string, err error) {
	withClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return withClaims.SignedString(j.singKey)
}

// ParseToken 验证Token
func (j *JWT) ParseToken(token string) (*CustomClaims, error) {
	withClaims, err := jwt.ParseWithClaims(token, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.singKey, nil
	})
	if err != nil {
		// 获取到 Jwt ValidationError 错误类型
		if ve, ok := err.(*jwt.ValidationError); ok {
			zap.S().Infof("获取到 Jwt ValidationError 原：%v 错误类型:%v", err, ve.Errors)
			if ve.Errors&jwt.ValidationErrorMalformed != 0 { // 令牌不完整
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 { // 令牌过期
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 { // 令牌还未生效
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
		return nil, TokenInvalid
	}
	if withClaims == nil {
		return nil, TokenInvalid
	}

	if claims, ok := withClaims.Claims.(*CustomClaims); ok {
		return claims, nil
	} else {
		return nil, TokenInvalid
	}

}
