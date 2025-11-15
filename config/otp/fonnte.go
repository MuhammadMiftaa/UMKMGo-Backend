package otp

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

const (
	URL_SEND_FONNTE = "https://api.fonnte.com/send"
)

type FonnteOtp struct {
	Request     FonnteRequest  `form:"request"       json:"request"`
	Response    FonnteResponse `form:"response"      json:"response"`
	Vendor      OTPVendor      `form:"vendor"        json:"vendor"`
	ClientTitle string         `form:"client_title"  json:"client_title"`
	ActionName  string         `form:"action_name"   json:"action_name"`
}

type FonnteRequest struct {
	Target      string `json:"target"`
	Message     string `json:"message"`
	CountryCode string `json:"country_code"`
}

type FonnteResponse struct {
	Status    bool   `json:"status"`
	Detail    string `json:"detail"`
	RequestID int    `json:"requestid"`
	Process   string `json:"process"`
	Reason    string `json:"reason"`
}

func (fo *FonnteOtp) SendOtp(phone, otpCode string) (string, error) {
	fn := "SendOTP"

	response := ""

	fo.Request.Target = phone
	fo.Request.Message = fmt.Sprintf("*%s* adalah kode verifikasi Anda. Demi keamanan, jangan bagikan kode ini.", otpCode)

	// Double check if phone is empty
	if fo.Request.Target == "" {
		return response, fmt.Errorf("%s-01:phone empty cannot continue, %w", fn, errors.New(ERR_PHONE_EMPTY))
	}

	if fo.Request.CountryCode == "" {
		fo.Request.CountryCode = COUNTRY_CODE_INDONESIA
	}

	if fo.Vendor.APIKey == "" {
		return response, fmt.Errorf("%s-02:fonnte token is not set in the configuration", fn)
	}

	json_data, err := json.Marshal(fo.Request)
	if err != nil {
		return response, fmt.Errorf("%s-01:server internal error json, %w", fn, err)
	}

	url := URL_SEND_FONNTE
	if fo.Vendor.Url != "" {
		url = fo.Vendor.Url
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
	req.Header.Add("Authorization", fo.Vendor.APIKey)

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

	var result *FonnteResponse

	err = json.Unmarshal([]byte(body), &result)
	if err != nil {
		return response, fmt.Errorf("%s-02:server error, cannot unmarshal response, %w", fn, err)
	}

	if !result.Status {
		return response, fmt.Errorf("%s-03:otp provider error with request id: %d, %w", fn, result.RequestID, errors.New(result.Reason))
	}

	return result.Detail, nil
}
