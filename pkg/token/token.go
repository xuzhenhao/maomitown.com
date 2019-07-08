package token

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"

	"github.com/gin-gonic/gin"
)

var (
	// ErrorMissingHeader 没有token信息
	ErrorMissingHeader = errors.New("Authorization 为空")
)

// Context 用户上下文信息
type Context struct {
	ID       uint64
	Username string
}

// Sign 签名生成token信息
func Sign(ctx *gin.Context, c Context, secret string) (tokenString string, err error) {
	//如果没有指定秘钥，使用默认的
	if secret == "" {
		secret = viper.GetString("jwt_secret")
	}

	//生成token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       c.ID,
		"username": c.Username,
		"nbf":      time.Now().Unix(),
		"iat":      time.Now().Unix(),
	})

	tokenString, err = token.SignedString([]byte(secret))
	return
}

// Parse 解析token是否匹配
func Parse(tokenString, secret string) (*Context, error) {
	ctx := &Context{}

	token, err := jwt.Parse(tokenString, secretFunc(secret))

	if err != nil {
		return ctx, err
	} else if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		ctx.ID = uint64(claims["id"].(float64))
		ctx.Username = claims["username"].(string)

		return ctx, nil
	} else {
		return ctx, err
	}
}

// ParseRequest 解析请求
func ParseRequest(c *gin.Context) (*Context, error) {
	header := c.Request.Header.Get("Authorization")

	secret := viper.GetString("jwt_secret")
	if len(header) == 0 {
		return &Context{}, ErrorMissingHeader
	}

	return Parse(header, secret)
}

func secretFunc(secret string) jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}

		return []byte(secret), nil
	}
}
