package bd

import(
	"context"
	"time"
	"github.com/ChrisVS-9/friends/models"
	"go.mongodb.org/mongo-driver/bson"
)
/*ChequeoyaExisteUsuario recibe un email de parametro y checa si ya existe en la BD*/
func ChequeoYaExisteUsuario(email string) (models.Usuario, bool, string){
	ctx, cancel := context.WithTimeout(context.Background(),15*time.Second)
	defer cancel()

	db := MongoCN.Database("friends")
	col :=db.Collection("usuarios")

	condicion := bson.M{"email":email}

	var resultado models.Usuario

	err := col.FindOne(ctx,condicion).Decode(&resultado)
	ID := resultado.ID.Hex()
	if err != nil{
		return resultado, false, ID
	}
	return resultado, true , ID
}