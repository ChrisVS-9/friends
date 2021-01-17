package bd

import(
	"context"
	"time"
	"fmt"
	"github.com/ChrisVS-9/friends/models"
	"go.mongodb.org/mongo-driver/bson"
)
/*ConsultoRelacion sirve para poder consultar nuestras relaciones*/
func ConsultoRelacion(t models.Relacion) (bool, error){
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("friends")
	col := db.Collection("relacion")

	condicion := bson.M{
		"usuarioid": t.UsuarioID,
		"usuariorelacionid": t.UsuarioRelacionID,
	}

	var resultado models.Relacion
	fmt.Println(resultado)
	err := col.FindOne(ctx, condicion).Decode(&resultado)
	if err!=nil{
		fmt.Println(err.Error())
		return false, err
	}
	return true, nil
}