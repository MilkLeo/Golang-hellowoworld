package utils

import (
	"github.com/dgrijalva/jwt-go"
	_ "github.com/dgrijalva/jwt-go"
)

type LoginInfo struct {
	Uid   string
	Utype int
}

//生成token
func GenerateToken(uid string, utype int, secret string, time int64) (string, error) {
	var token string
	var err error
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uid":   uid,
		"exp":   time,
		"utype": utype,
	})
	token, err = at.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return token, nil
}

//验证token
func CheckToken(token string, secret string) (LoginInfo, error) {
	var userinfo LoginInfo
	claim, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return userinfo, err
	}
	uid := claim.Claims.(jwt.MapClaims)["uid"].(string)
	utype := int(claim.Claims.(jwt.MapClaims)["utype"].(float64))

	userinfo.Uid = uid
	userinfo.Utype = utype

	return userinfo, nil
}
