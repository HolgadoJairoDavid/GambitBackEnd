package main

import (
	// "fmt"
	"context"
	"os"
	"strings"

	"github.com/HolgadoJairoDavid/GambitBackEnd/awsgo"
	"github.com/aws/aws-lambda-go/events"

	"github.com/HolgadoJairoDavid/GambitBackEnd/bd"
	"github.com/HolgadoJairoDavid/GambitBackEnd/handlers"

	lambda "github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(EjecutoLambda)
}

func EjecutoLambda(ctx context.Context, request events.APIGatewayV2HTTPRequest) (*events.APIGatewayProxyResponse, error) {
	awsgo.InicializoAWS()
	if !ValidoParametros() {
		panic("Error en los parámetros. Debe enviar 'SecretName', 'UrlPrefix'")
	}

	var res *events.APIGatewayProxyResponse
	path := strings.Replace(request.RawPath, os.Getenv("UrlPrefix"), "", -1)
	method := request.RequestContext.HTTP.Method
	body := request.Body
	heaeder := request.Headers

	bd.RedSecret()

	status, message := handlers.Manejadores(path, method, body, heaeder, request)

	headersResp := map[string]string{
		"Content-Type": "application/json",
	}
	res = &events.APIGatewayProxyResponse{
		StatusCode: status,
		Body:       string(message),
		Headers:    headersResp,
	}

	return res, nil
}

func ValidoParametros() bool {
	_, traeParametro := os.LookupEnv("SecretName")
	if !traeParametro {
		return traeParametro
	}

	// _, traeParametro = os.LookupEnv("UserPoolId")
	// if !traeParametro {
	// 	return traeParametro
	// }
	// _, traeParametro = os.LookupEnv("Region")
	// if !traeParametro {
	// 	return traeParametro
	// }
	_, traeParametro = os.LookupEnv("UrlPrefix")
	if !traeParametro {
		return traeParametro
	}
	return traeParametro
}
