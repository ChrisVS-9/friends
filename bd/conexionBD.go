package bd

import (
	"context"
    "log"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)
/*MongoCN es el objeto de conexion a la BD*/
var MongoCN=ConectarBD()

var clientOptions = options.Client().ApplyURI("mongodb+srv://Chris:bad91291olokika@friends.5e4b4.mongodb.net/friends?retryWrites=true&w=majority")
/*ConectarBD es la funcion que permite conectar la BD*/
func ConectarBD() *mongo.Client{
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil{
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil{
		log.Fatal(err.Error())
		return client
	}
	log.Println("Conexion Exitosa con la BD")
	return client
}
/*ChequeoConnection es el Ping a la BD*/
func ChequeoConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil{
		return 0
	}
	return 1
}