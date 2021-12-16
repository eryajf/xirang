package middleware

import (
	"errors"

	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// JWTAuth 中间件，检查token
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.URL.Path != "/api/v1/login" {
			token := c.Request.Header.Get("Authorization")
			if token == "" {
				c.JSON(250, gin.H{
					"status": -1,
					"msg":    "请求未携带token，无权限访问",
				})
				c.Abort()
				return
			}

			j := NewJWT()
			// parseToken 解析token包含的信息
			claims, err := j.ParseToken(token)
			if err != nil {
				if err == TokenExpired {
					c.JSON(251, gin.H{
						"status": -1,
						"msg":    "授权已过期",
					})
					c.Abort()
					return
				}
				c.JSON(252, gin.H{
					"status": -1,
					"msg":    err.Error(),
				})
				c.Abort()
				return
			}
			// 继续交由下一个路由处理,并将解析出的信息传递下去
			c.Set("claims", claims)
		}
	}
}

// JWT 签名结构
type JWT struct {
	SigningKey []byte
}

// 一些常量
var (
	TokenExpired     error  = errors.New("Token is expired")
	TokenNotValidYet error  = errors.New("Token not active yet")
	TokenMalformed   error  = errors.New("That's not even a token")
	TokenInvalid     error  = errors.New("Couldn't handle this token:")
	SignKey          string = "newtrekWang"
)

// CustomClaims 载荷，可以加一些自己需要的信息
type CustomClaims struct {
	ID       int    `json:"user_id"`
	Username string `json:"username"`
	Mobile   string `json:"mobile"`
	Email    string `json:"email"`
	jwt.StandardClaims
}

// NewJWT 新建一个jwt实例
func NewJWT() *JWT {
	return &JWT{
		[]byte(GetSignKey()),
	}
}

// GetSignKey 获取signKey
func GetSignKey() string {
	return SignKey
}

// SetSignKey 这是SignKey
func SetSignKey(key string) string {
	SignKey = key
	return SignKey
}

// CreateToken 生成一个token
func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// ParseToken 解析Tokne
func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, TokenInvalid
}

// RefreshToken 更新token
func (j *JWT) RefreshToken(tokenString string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(1 * time.Hour).Unix() // 默认token有效期为1个小时
		return j.CreateToken(*claims)
	}
	return "", TokenInvalid
}
