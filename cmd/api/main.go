package main
import (
	"fmt"
	"net/http"

	"github.com/blue-onion/goApi/internal/handler"
	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)
func main()  {
	log.SetReportCaller(true)
	r:=chi.NewRouter()
	handler.Handler(r)
	fmt.Println("Go Api baby")
	err:=http.ListenAndServe("localhost:3067",r)
	if err!=nil{
		log.Error(err)
	}

}