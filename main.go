package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type URLRedirect struct {
	Original  string `json:"original"`
	Shortened string `json:"shortened"`
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	/*
			authToken, ok := os.LookupEnv("FAUNADB_SERVER_SECRET")

			if !ok {
				return events.APIGatewayProxyResponse{
					StatusCode: 500,
					Body:       "Could not connect to Fauna!",
				}, nil
			}
		client := f.NewFaunaClient(authToken)
	*/

	return events.APIGatewayProxyResponse{
		Headers: map[string]string{
			"Location": "https://pythonmagic.com",
		},
		StatusCode: 301,
		Body: "", //      "Hello AWS Lambda and Netlify",
	}, nil
}

func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(handler)
}
