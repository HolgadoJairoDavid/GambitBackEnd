package routes

import (
	"encoding/json"
	"strconv"

	"github.com/HolgadoJairoDavid/GambitBackEnd/bd"
	"github.com/HolgadoJairoDavid/GambitBackEnd/models"
	// "github.com/aws/aws-lambda-go/events"
)

func InsertCategory(body string, User string) (int, string) {
	var t models.Category

	err := json.Unmarshal([]byte(body), &t)
	if err != nil {
		return 400, "Error en los datos recibidos " + err.Error()
	}

	if len(t.CategName) == 0 {
		return 400, "Debe especificar el Nombre (Title) de la Categoría"
	}

	if len(t.CategPath) == 0 {
		return 400, "Debe especificar el Nombre (Title) de la Categoría"
	}

	isAdmin, msg := bd.UserIsAdmin(User)
	if !isAdmin {
		return 400, msg
	}

	result, err2 := bd.InsertCategory(t)

	if err2 != nil {
		return 400, "Ocurrió un error al intentar realizar ek registro de la categoría " + t.CategName + " > " + err2.Error()
	}

	return 200, "{CategID: " + strconv.Itoa(int(result)) + "}"
}
