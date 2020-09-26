package  auth 
import (
	"fmt" 
	"net/http"
	"context"
	"github.com/conglt10/web-golang/models"
	"github.com/conglt10/web-golang/auth "
	"github.com/conglt10/web-golang/utils"
	"github.com/conglt10/web-golang/database"
	"github.com/julienschmidt/httprouter"
	"github.com/asakevich/govalidator"
	"go.mongodb.org/mongo-driver/bson"
)


func Register (w http.ResponseWriter , r *http.Request , _httprouter.Params){
	username := r.PostFormValue ("username")
	email := r.PostFormValue ("email ")
	password := r.PostFormValue ( "password ")

	if govalidator.IsNull ( username ) || govalidator.IsNull(email) || govalidator(password) {
		res.JSON(w , 400 , " data can not empty ")
		return 

	}

	if !govalidator.IsEmail(email ) {
		res.JSON ( w , 400 , " email is incalid ")
		return 
	}
	 username = models.Santize(username )
	 email = models.Santize(email)
	password = models.Santize ( password) 

	colection := db.ConnectUsers ()
	 var result bson.M
	 errFindUsername := collection.FindOne (context.TODO() , bson.M {"username" : email}).Decode (&result)
	errFindEmail := collection.FindOne ( context.TODO () , bson.M {"email":email }).Decode(&result)

	if errFindUsername == || errFindEmail == nil {
		 res.JSON(w ,  409 ,"USER does exists ")
		 return 

	}

	password , err := models.Hash(password) 

	if err != nil {
		res.JSON (w , 500 , " register  has failed ") 
		return 
	}

	 newUser := bson.M {" username " : username , "email" :  email , "password " : password }
	_, errs := collection.InsertOne (context.TODO() , newUser )

	if errs != nil {
		res.JSON (w 500 , " register has failes ")
		return 
	}
	res.JSON(w ,201 ," register sccesfully ")
	
}