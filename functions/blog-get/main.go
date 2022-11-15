package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type BlogResponse struct {
	Name   string `json:"name"`
	Active bool   `json:"show"`
}

type Response events.APIGatewayProxyResponse

func Handler(request events.APIGatewayProxyRequest) (Response, error) {

	name := request.PathParameters["name"]

	active := request.QueryStringParameters["active"]

	activeBool, _ := strconv.ParseBool(active)
	blogResponse := BlogResponse{name, activeBool}

	res, err := json.Marshal(blogResponse)
	if err != nil {
		panic(err)
	}

	return Response{
		Body:       string(res),
		StatusCode: http.StatusOK,
		Headers:    map[string]string{"Content-Type": "application/json"},
	}, nil
}

func main() {
	lambda.Start(Handler)
}
