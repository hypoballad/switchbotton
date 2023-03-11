package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/rs/xid"
	"github.com/spf13/viper"
)

type signature struct {
	Token string `json:"token"`
	Nonce string `json:"nonce"`
	T     int64  `json:"t"`
	Sign  string `json:"sign"`
}

func init() {
	viper.SetConfigType("yml")
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalln("Error reading config file, ", err)
	}
}

// get signature from token, secret, nonce
func getSignature() signature {
	// Use viper to read the token and secret values from a configuration file called config.yml
	// The config.yml file should be in the same directory as the main.go file
	// The config.yml file should contain the following:
	// token: <your token>
	// secret: <your secret>
	// device: <your device id>

	token := viper.GetString("token")
	secret := viper.GetString("secret")
	// get nonce from xid
	nonce := xid.New().String()
	t := int64(time.Now().UnixNano() / int64(time.Millisecond))
	stringToSign := fmt.Sprintf("%s%d%s", token, t, nonce)
	key := []byte(secret)
	msg := []byte(stringToSign)

	h := hmac.New(sha256.New, key)
	h.Write(msg)
	sign := base64.StdEncoding.EncodeToString(h.Sum(nil))

	result := signature{
		Token: token,
		Nonce: nonce,
		T:     t,
		Sign:  sign,
	}
	return result
}

// turn on switchbot device
func turnOn(sign signature) {
	// device id
	deviceID := viper.GetString("deviceID")
	body := []byte(`{"command":"turnOn","parameter":"default","commandType":"command"}`)
	req, err := http.NewRequest("POST", fmt.Sprintf("https://api.switch-bot.com/v1.1/devices/%s/commands", deviceID), bytes.NewReader(body))
	if err != nil {
		fmt.Printf("Failed to create request: %s\n", err)
		return
	}
	req.Header.Set("Authorization", sign.Token)
	req.Header.Set("t", fmt.Sprintf("%d", sign.T))
	req.Header.Set("sign", sign.Sign)
	req.Header.Set("nonce", sign.Nonce)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Content-Length", fmt.Sprintf("%d", len(body)))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Failed to execute request: %s\n", err)
		return
	}
	defer resp.Body.Close()

	fmt.Printf("statusCode: %d\n", resp.StatusCode)

	// get response
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Failed to read response: %s\n", err)
		return
	}
	fmt.Printf("response: %s\n", body)

}

func main() {
	sign := getSignature()
	turnOn(sign)
	os.Exit(0)
}
