package utils

import (
	"context"
	"time"

	"github.com/labstack/echo/v4"
)

const appTimeout = time.Second * 10

type OTPData struct {
    PhoneNumber string `json:"phoneNumber,omitempty" validate:"required"`
}

type VerifyData struct {
    User *OTPData  `json:"user,omitempty" validate:"required"`
    Code string `json:"code,omitempty" validate:"required"`
}

func SendSMS(e echo.Context) error {
	_, cancel := context.WithTimeout(context.Background(), appTimeout)	
	defer cancel()

	var payload OTPData

	data := OTPData{
		PhoneNumber: payload.PhoneNumber,
	}

	_, err := TwilioSendOTP(data.PhoneNumber)
	if err != nil {
		return HandlerError(e, 500, err)
	}

	return e.JSON(200, "OTP sent successfully")
}
