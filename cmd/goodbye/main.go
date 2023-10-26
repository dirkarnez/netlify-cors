package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func downloadAsBase64String(url string) string {
	// Send an HTTP GET request to the file URL
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error downloading file:", err)
		return ""
	}
	defer response.Body.Close()

	// Check if the response was successful
	if response.StatusCode != http.StatusOK {
		fmt.Println("Error downloading file. Server returned status:", response.StatusCode)
		return ""
	}
	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error writing to output file:", err)
		return ""
	}

	// Encode the body content to Base64
	base64String := base64.StdEncoding.EncodeToString(bodyBytes)

	return base64String
}

func handler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	fmt.Print("Got request for '/.netlify/functions/goodbye', this message is dumpled by 'cmd/goodbye/main.go'")
	return &events.APIGatewayProxyResponse{
		StatusCode:      200,
		IsBase64Encoded: true,
		Body:            downloadAsBase64String("https://raw.githubusercontent.com/mozilla/pdf.js/ba2edeae/examples/learning/helloworld.pdf"),
	}, nil
}

func main() {
	lambda.Start(handler)
}
