package public

import (
	"errors"
	"time"

	"github.com/buddhachain/buddha/apiserver/factory/config"
	"github.com/buddhachain/buddha/common/define"
	"github.com/buddhachain/buddha/common/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

const (
	Token = "token"
)

//生成jwt signed token
func GenerateToken(id string) string {
	jwtConf := config.JWTConfig()
	claim := jwt.StandardClaims{
		//Audience:
		ExpiresAt: time.Now().Add(time.Hour * 5).Unix(),
		Id:        id,
		Issuer:    jwtConf.Issuer,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	ss, err := token.SignedString([]byte(jwtConf.Passwd))
	if err != nil {
		panic(err)
	}
	return ss
}

//parse jwt token
func ParseToken(ss string) (string, error) {
	claims := jwt.StandardClaims{}
	token, err := jwt.ParseWithClaims(ss, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.JWTConfig().Passwd), nil
	})
	if err != nil {
		return "", err
	}
	if !token.Valid {
		return "", errors.New("invalid token")
	}
	return claims.Id, nil
}

func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader(Token)
		if token == "" {
			utils.Response(c, define.ErrToken, define.TokenErr, nil)
			c.Abort()
			return
		}
		id, err := ParseToken(token)
		if err != nil {
			utils.Response(c, define.ErrToken, define.TokenErr, nil)
			c.Abort()
			return
		}
		c.Request.Header.Set("user", id)
		c.Next()
	}
}
