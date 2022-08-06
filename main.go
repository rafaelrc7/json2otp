package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"
)

type Otp struct {
	Secret, Issuer, Label, Type, Algorithm, Thumbnail string
	Digits, Last_used, Used_frequency, Period         uint
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		otps := []Otp{}
		if err := json.Unmarshal([]byte(scanner.Text()), &otps); err != nil {
			log.Fatalf("unmarshal: %s\n", err)
		}

		for _, otp := range otps {
			uriBase := fmt.Sprintf("otpauth://%s/%s", strings.ToLower(otp.Type), otp.Label)
			otpauth, err := url.Parse(uriBase)
			if err != nil {
				log.Fatalf("url.Parse: %s", err)
			}

			params := url.Values{}
			params.Add("secret", strings.Trim(otp.Secret, "="))
			params.Add("digits", fmt.Sprint(otp.Digits))
			params.Add("issuer", otp.Issuer)

			if algorithm := strings.ToLower(otp.Algorithm); algorithm != "sha1" {
				params.Add("algorithm", algorithm)
			}
			if period := otp.Period; period != 30 {
				params.Add("period", fmt.Sprint(period))
			}

			otpauth.RawQuery = params.Encode()

			fmt.Printf("%s/%s:\n%s\n\n", otp.Issuer, otp.Label, otpauth)
		}
	}
}
