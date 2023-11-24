package utils

import (
	"errors"

	"github.com/ashiruhabeeb/simpleTodoApp/config"
	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/verify/v2"
)

var client *twilio.RestClient = twilio.NewRestClientWithParams(twilio.ClientParams{
	Username: config.GetENV("TWILIO_ACCOUNT_SID"),
	Password: config.GetENV("TWILIO_AUTHTOKEN"),
})

func TwilioSendOTP(phoneNumber string) (string, error) {
	params := &twilioApi.CreateVerificationParams{}
	params.SetTo(phoneNumber)
	params.SetChannel("sms")

	resp, err := client.VerifyV2.CreateVerification(config.GetENV("TWILIO_SERVICES_ID"), params)
	if err != nil {
		return "", err
	}

	return *resp.Sid, nil
}

func TwilioVerifyOTP(phoneNumber string, code string) error {
	params := &twilioApi.CreateVerificationCheckParams{}
	params.SetTo(phoneNumber)
	params.SetCode(code)

	resp, err := client.VerifyV2.CreateVerificationCheck(config.GetENV("TWILIO_SERVICES_ID"), params)
	if err != nil {
		return err
	}

	// BREAKING CHANGE IN THE VERIFY API
	// https://www.twilio.com/docs/verify/quickstarts/verify-totp-change-in-api-response-when-authpayload-is-incorrect
	if *resp.Status != "approved" {
		return errors.New("not a valid code")
	}

	return nil
}
