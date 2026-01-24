package api
import (
	"encoding/json"
	"net/http"
)
type coinBalanceParams struct {
	UserName string
}
type statusCode int
type coinResponse struct {
	StatusCode int
	Balance    int64
}
type ErrorRes struct{
	StatusCode int
	Message string
}

func writeError(w http.ResponseWriter,message string,code int){
	res:=ErrorRes{
		StatusCode: code,
		Message: message,
	}
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(res)
	
}
var(
	ReqErrorHandler=func (w http.ResponseWriter,err error)  {
		writeError(w,err.Error(),400)
	}
	InternalErrorHandler=func (w http.ResponseWriter)  {
		writeError(w,"An unexpected Error Occurred",500)
	}
	NotFoundErrorHandler=func (w http.ResponseWriter)  {
		writeError(w,"Route not found",404)
	}
)