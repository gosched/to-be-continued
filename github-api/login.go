package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/anaskhan96/soup"
)

// html/template

// DeviceVerificationCode

// LoginURL .
const LoginURL = "https://github.com/login"

var client http.Client

func init() {
	// jar, err := cookiejar.New(nil)
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }

	redirect := func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}

	client = http.Client{
		CheckRedirect: redirect,
		// Jar:           jar,
	}
}

// GetAuthenticityToken .
func GetAuthenticityToken() (string, error) {
	resp, err := client.Get(LoginURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	document := soup.HTMLParse(string(body))

	token := document.Find("input", "name", "authenticity_token").Attrs()["value"]

	return token, nil
}

// RequestLogin .
func RequestLogin(loginField, password, token string) error {
	data := url.Values{}
	data.Set("authenticity_token", token)
	data.Set("login_field", loginField)
	data.Set("password", password)

	// req, err := http.NewRequest("POST", LoginURL, strings.NewReader(data.Encode()))
	// if err != nil {
	// 	return err
	// }

	// req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	// req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.88 Safari/537.36")
	// req.Header.Set("Referer", "https://github.com/")
	// req.Header.Set("Host", "github.com")

	// resp, err := client.Do(req)

	resp, err := client.PostForm(LoginURL, data)

	if err != nil {
		log.Println(err)
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// f, err := os.OpenFile("./body.html", os.O_CREATE|os.O_WRONLY, 0644)
	// if err != nil {
	// 	return err
	// }
	// f.Write(body)
	// f.Close()

	if resp.StatusCode == 200 {

	}

	cookies := resp.Cookies()

	_ = body
	_ = cookies

	// fmt.Println(resp)
	fmt.Println("StatusCode: ", resp.StatusCode)
	// fmt.Println("Body: ", string(body))
	// fmt.Println("Cookies: ", cookies)

	return nil
}

// GetDeviceVerificationCode .
func GetDeviceVerificationCode() {

}

// SubmitDeviceVerificationCode .
func SubmitDeviceVerificationCode() {

}

// Login .
func Login(loginField, password string) error {
	token, err := GetAuthenticityToken()
	if err != nil {
		return err
	}
	RequestLogin(loginField, password, token)
	GetDeviceVerificationCode()
	SubmitDeviceVerificationCode()
	return nil
}

// Logout .
func Logout() {
	// client.Get(LogoutURL)
}
