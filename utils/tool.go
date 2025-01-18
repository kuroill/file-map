package utils

import (
	"encoding/json"
	"errors"
	"file-map-server/app/config"
	"fmt"
	"strings"
)

func AuthToken(token string) error {
	if token == "" {
		return errors.New("token is not found")
	}
	s := strings.Split(token, " ")
	if len(s) >= 2 {
		token = s[1]
	}

	secret := config.GetConfig().User.SecretKey
	pwd := config.GetConfig().User.Password
	decrypt, err := DecryptAES([]byte(secret), token)
	if err != nil {
		return err
	}
	if decrypt != pwd {
		return errors.New("password error")
	}
	return nil
}

func PrettyStruct(data interface{}) {
	val, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return
	}
	fmt.Println(string(val))
}
