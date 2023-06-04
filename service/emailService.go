package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	utils "tasks/common"
)

type response struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}

func SendEmail() {
	apiKey := "0350B92141F9921673D8BBE98D8C5FEA67C92C57FF60227B8365DD2486C3D7A2C8E238262214304938383CC5A085AC2D"
	from := "gargibanerjee49@gmail.com"
	to := []string{"gargibanerjee49@gmail.com"}
	subject := "Hello from Elastic Email"
	message := "This is the content of the email."

	// Create the HTTP client
	client := &http.Client{}

	// Prepare the request URL
	requestURL := "https://api.elasticemail.com/v2/email/send"

	// Prepare the request parameters
	params := url.Values{}
	params.Set("apikey", apiKey)
	params.Set("from", from)
	params.Set("to", strings.Join(to, ","))
	params.Set("subject", subject)
	params.Set("bodyHtml", message)

	// Create the POST request
	req, err := http.NewRequest("POST", requestURL, strings.NewReader(params.Encode()))
	if err != nil {
		fmt.Println(err)
		return
	}

	// Set the request headers
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Unmarshal the response body into a struct
	responseBody := response{}
	err = json.Unmarshal(body, &responseBody)
	if err != nil {
		fmt.Println(err)
		return
	}
	// Check the response status

	if responseBody.Error == "Access Denied" || responseBody.Success == false {
		utils.Logger.Err(fmt.Errorf(responseBody.Error)).Msgf("Failed to send email.")
		return
	}

	fmt.Println("Email sent successfully!")
}
