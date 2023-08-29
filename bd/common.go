package bd

import (
	"database/sql"
	"fmt"
	"os"

	// "fmt"
	"github.com/HolgadoJairoDavid/GambitBackEnd/models"
	"github.com/HolgadoJairoDavid/GambitBackEnd/secretm"
	_ "github.com/go-sql-driver/mysql"
)

var SecretModel models.SecretRDSJson
var err error
var Db *sql.DB

func RedSecret() error {
	SecretModel, err = secretm.GetSecret(os.Getenv("SecretName"))

	return err
}

func DbConnect() error {
	Db, err = sql.Open("mysql", ConnStr(SecretModel))
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	err = Db.Ping()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("Conexión existosa de la base de datos")
	return nil
}

func ConnStr(clave models.SecretRDSJson) string {
	var dbUser, authToken, dbEndpoint, dbName string

	dbUser = clave.Username
	authToken = clave.Password
	dbEndpoint = clave.Host
	dbName = "gambit"
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?allowCleartextPasswords=true", dbUser, authToken, dbEndpoint, dbName)
	fmt.Println(dsn)
	return dsn
}

func UserIsAdmin(userUUID string) (bool, string) {
	fmt.Println("Comienza userIsAdmin")

	err := DbConnect()
	if err != nil {
		return false, err.Error()
	}

	defer Db.Close()

	sentencia := "SELECT 1 FROM users WHERE User_UUID='" + userUUID + "' AND User_Status = 0"

	rows, err := Db.Query(sentencia)

	if err != nil {
		return false, err.Error()
	}
	var valor string

	rows.Next()
	rows.Scan(&valor)

	fmt.Println("UserIsAdmin > Ejecución exitosa - valor devuelto " + valor)
	if valor == "1" {
		return true, ""
	}

	return false, "User is not Admin"
}
