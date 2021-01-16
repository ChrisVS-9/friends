package routers

import(
	"errors"
	"strings"
	
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/ChrisVS-9/friends/bd"
	"github.com/ChrisVS-9/friends/models"
)
/*Email valor de Email usado en todos los EndPoints*/
var Email string
/*IDUsuario es el ID devuelto del modelo, que se usara en todos los EndPoints*/
var IDUsuario string

/*ProcesoToken proceso token para extraer sus valores*/
func ProcesoToken(tk string) (*models.Claim, bool, string, error){
	miClave :=[]byte("MastersdelDesarrollo_grupodeFacebook")
	claims := &models.Claim{}

	splitToken := strings.Split(tk,"Bearer")
	if len(splitToken) != 2{
		return claims, false, string(""),errors.New("formato de token invalido")
	}
	tk = strings.TrimSpace(splitToken[1])

	tkn, err :=jwt.ParseWithClaims(tk, claims, func(token *jwt.Token)(interface{},error){
		return miClave, nil
	})
	if err ==nil{
		_, encontrado, ID := bd.ChequeoYaExisteUsuario(claims.Email)
		if encontrado ==true{
			Email =claims.Email
			IDUsuario =claims.ID.Hex()
		}
		return claims, encontrado, ID, nil
	}
	if !tkn.Valid{
		return claims, false, string(""), errors.New("toke Invalido")
	}
	return claims, false, string(""),err
}