package secretm

import (
	"encoding/json"
	"fmt"

	"github.com/HolgadoJairoDavid/GambitBackEnd/awsgo"
	"github.com/HolgadoJairoDavid/GambitBackEnd/models"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

func GetSecret(nombreSecret string) (models.SecretRDSJson, error) {
	var datoSecret models.SecretRDSJson
	fmt.Println(" > Pido Secreto " + nombreSecret)

	svc := secretsmanager.NewFromConfig(awsgo.Cfg)

	clave, err := svc.GetSecretValue(awsgo.Ctx, &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(nombreSecret),
	})

	if err != nil {
		fmt.Println(err.Error())
		return datoSecret, err
	}

	json.Unmarshal([]byte(*clave.SecretString), &datoSecret)
	fmt.Println(" > Lectura secret ok " + nombreSecret)
	return datoSecret, nil
}
