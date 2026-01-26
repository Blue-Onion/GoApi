package handler

import (
	"encoding/json"
	"net/http"

	"github.com/blue-onion/goApi/api"
	"github.com/blue-onion/goApi/internal/tools"
	"github.com/gorilla/schema"
	"github.com/sirupsen/logrus"
)

func GetCoinBalance(w http.ResponseWriter, r *http.Request) {
	params := &api.CoinBalanceParams{}
	decoder := schema.NewDecoder()
	err := decoder.Decode(params, r.URL.Query())
	if err != nil {
		api.InternalErrorHandler(w)
		return
	}
	db, err := tools.NewDatabase()
	if err != nil {
		api.InternalErrorHandler(w)
		return
	}

	tokenDetails := (*db).GetCoinDetails(params.UserName)
	if tokenDetails == nil {
		logrus.Error(err)
		api.InternalErrorHandler(w)
		return
	}
	res := api.CoinResponse{
		StatusCode: http.StatusOK,
		Balance:    (*tokenDetails).Coins,
	}
	w.Header().Set("Content-type", "application/Json")
	err=json.NewEncoder(w).Encode(res)
	if err!=nil{
		logrus.Error(err)
		api.InternalErrorHandler(w)
		return
	}


}
