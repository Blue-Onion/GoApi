package middeware

import (
	"fmt"

	"net/http"

	"github.com/blue-onion/goApi/api"
	"github.com/blue-onion/goApi/internal/tools"

	"github.com/sirupsen/logrus"
)
var unAuthorizedError=fmt.Errorf("Unauthorized user")
func Authorization(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username:=r.URL.Query().Get("username")
		token:=r.Header.Get("Authorization")
		if username==""||token==""{
			logrus.Error(unAuthorizedError)
			api.ReqErrorHandler(w,unAuthorizedError)
			return 
		}
		db,err:=tools.NewDatabase()
		if err!=nil{
			return 
		}
		loginDetails:=(*db).GetUserDetails(username)
		if loginDetails==nil||token!=(*loginDetails).Token{
			logrus.Error(unAuthorizedError)
			api.ReqErrorHandler(w,unAuthorizedError)
			return 
		}
		next.ServeHTTP(w,r)
	})
}