package main

import (
	"context"
	"strconv"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Response events.APIGatewayProxyResponse

type JsonResponse struct {
	Resultado   float64 `json:"resultado"`
}

func getSum (parameters map[string]string) float64 {
	var number float64

	for key := range parameters {
		
		value, err :=  strconv.ParseFloat(parameters[key], 64)
		
		if err != nil {
			number = number + 0
		} else {
			number = number + value
		}	
	}

	return number
}

func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (Response, error) {
	
	var result = getSum(request.QueryStringParameters)

	jsonResponse := JsonResponse{
		Resultado: result,
	}

	byteArray, err := json.Marshal(jsonResponse)

	if err != nil {
		return Response{Body: "erro ao tentar executar a soma", StatusCode: 500}, nil
	}
	
	return Response{Body: string(byteArray), StatusCode: 200}, nil
}

func main() {
	lambda.Start(Handler)
}
