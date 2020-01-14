package middleware

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"strconv"
	"time"

	//"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type jwtClaims struct {
	jwt.StandardClaims
	UserID int `json:"user_id"`
}

var (
	secret = "miniProject"  //salt
	ExpireTime = 3600  //token expire time
)

func JwtAAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
	/*	if token == "" {
			c.JSON(http.StatusForbidden, gin.H{
				"status": "403",
				"msg": "token Invalid",
			})
		}*/

		//refresh(c)
		claim, err := verifyToken(token)
		if err != nil {
			c.String(http.StatusBadRequest,err.Error())
			c.Abort()
		}
		c.Set("uid",claim.UserID)
		//fmt.Println(token)
		fmt.Println(claim.UserID)
		c.Next()
		//c.Abort()
	}
}

func ProduceToken(uid string) string {
	id,_ := strconv.Atoi(uid)
	claims := &jwtClaims{
		UserID: id,
	}
	claims.IssuedAt = time.Now().Unix()
	claims.ExpiresAt = time.Now().Add(time.Second * time.Duration(ExpireTime)).Unix()
	singedToken, err := genToken(claims)
	fmt.Println(singedToken,err)
	return singedToken
}

func genToken(claims *jwtClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func verifyToken(varifyToken string) (*jwtClaims, error) {
	token, err := jwt.ParseWithClaims(varifyToken, &jwtClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil,err
	}
	claims, ok := token.Claims.(*jwtClaims)
	if !ok {
		return nil, errors.New("token Invalid")
	}
	if err := token.Claims.Valid(); err != nil {
		return nil, errors.New("token Invalid")
	}
	return claims, nil
}

/*
func refresh(c *gin.Context) {
	strToken := c.Param("token")
	claims, err := verifyToken(strToken)
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
		return
	}
	claims.ExpiresAt = time.Now().Unix() + (claims.ExpiresAt - claims.IssuedAt)
	signedToken, err := genToken(claims)
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
		return
	}
	c.String(http.StatusOK, signedToken, ", ", claims.ExpiresAt)
}
*/