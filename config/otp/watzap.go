package otp

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

const (
	URL_SEND_WATZAP = "https://api.watzap.id/v1/send_message"
)

type WatzapOtp struct {
	Request     WatzapRequest  `form:"request"       json:"request"`
	Response    WatzapResponse `form:"response"      json:"response"`
	Vendor      OTPVendor      `form:"vendor"        json:"vendor"`
	ClientTitle string         `form:"client_title"  json:"client_title"`
	ActionName  string         `form:"action_name"   json:"action_name"`
}

type WatzapRequest struct {
	ApiKey      string `json:"api_key"`
	NumberKey   string `json:"number_key"`
	PhoneNumber string `json:"phone_no"`
	Message     string `json:"message"`
}

type WatzapResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Ack     string `json:"ack"`
}

func (wo WatzapResponse) ToString() string {
	data, _ := json.Marshal(wo) // Convert to a json string
	return string(data)
}

func (wo *WatzapOtp) SendOtp(phone, otpCode string) (string, error) {
	fn := "Watzap-SendOTP"

	response := ""

	wo.Request.PhoneNumber = phone
	wo.Request.Message = strings.ReplaceAll(wo.Request.Message, MESSAGE_OTP_CODE, otpCode)
	wo.Request.Message = strings.ReplaceAll(wo.Request.Message, MESSAGE_ACTION_CODE, wo.ActionName)
	wo.Request.Message = strings.ReplaceAll(wo.Request.Message, MESSAGE_CLIENT_CODE, wo.ClientTitle)

	// Double check if phone is empty
	if wo.Request.PhoneNumber == "" {
		return response, fmt.Errorf("%s-01:phone empty cannot continue, %w", fn, errors.New(ERR_PHONE_EMPTY))
	}

	json_data, err := json.Marshal(wo.Request)
	if err != nil {
		return response, fmt.Errorf("%s-01:server internal error json, %w", fn, err)
	}

	url := URL_SEND_WATZAP
	if wo.Vendor.Url != "" {
		url = wo.Vendor.Url
	}

	req, err := http.NewRequest(
		"POST",
		url,
		bytes.NewReader(json_data),
	)
	if err != nil {
		return response, fmt.Errorf("%s-01:server error cannot make request, %w", fn, err)
	}

	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return response, fmt.Errorf("%s-02:server error, cannot make http call, %w", fn, err)
	}

	defer req.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return response, fmt.Errorf("%s-02:server error, cannot read http response, %w", fn, err)
	}

	var result *WatzapResponse

	err = json.Unmarshal([]byte(body), &result)
	if err != nil {
		return response, fmt.Errorf("%s-02:server error, cannot unmarshal response, %w", fn, err)
	}

	if resp.StatusCode != http.StatusOK {
		return response, fmt.Errorf("%s-03:otp provider error, %w", fn, err)
	}

	return result.ToString(), nil
}
