package utils

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
)

func IsValidToken(token string) bool {
	arr := strings.Split(token, ".")
	if len(arr) != 3 {
		return false
	}
	if arr[2] != SignCreatFunc(arr[0]+"."+arr[1]) {
		return false
	}
	return true
}

func SignCreatFunc(HeaderPointBody string) string {
	h := sha256.New()
	secret := "salt"
	h.Write([]byte(HeaderPointBody + secret))
	return fmt.Sprintf("%x", h.Sum(nil))
}

func CreateJWT(username string) string {

	header, err := json.Marshal(map[string]string{
		"alg": "HS256",
		"typ": "JWT",
	})
	if err != nil {
		fmt.Println("marshal err")
		return ""
	}
	jwtHeader := base64.URLEncoding.EncodeToString(header)

	body, err := json.Marshal(map[string]interface{}{
		"iss": username,
		"exp": "20231216",
		"jti": "100",
	})
	if err != nil {
		fmt.Println("marshal err")
		return ""
	}
	jwtBody := base64.URLEncoding.EncodeToString(body)

	sign := SignCreatFunc(jwtHeader + "." + jwtBody)
	return jwtHeader + "." + jwtBody + "." + sign
}
