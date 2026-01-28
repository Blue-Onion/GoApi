package tools

import (
	"time"

	"github.com/sirupsen/logrus"
)

type mockDb struct{}
var mockLoginDetail=map[string]LoginDetails{
	"adit":{
		Token:"1234",
		UserName: "Aditya",
	},
	"Vijay":{
		Token:"1235",
		UserName: "Vijay",
	},
	"ajiy":{
		Token:"1236",
		UserName: "ajiy",
	},
	"ehh":{
		Token:"1237",
		UserName: "ehh",
	},
}
var mockCoinDetail=map[string]CoinDetails{
	"adit":{
		Coins: 192,
		UserName: "Aditya",
	},
	"Vijay":{
		Coins: 197,
		UserName: "Vijay",
	},
	"ajiy":{
		Coins: 199,
		UserName: "ajiy",
	},
	"ehh":{
		Coins: 1967,
		UserName: "ehh",
	},
}
func (d *mockDb) GetUserDetails(userName string)*LoginDetails{
	time.Sleep(time.Second*1)
	clientLoginData:=LoginDetails{}
	clientLoginData,ok:=mockLoginDetail[userName]
	if !ok{
		logrus.Info(&clientLoginData)
		return nil
	}

	return &clientLoginData
}
func (d *mockDb) GetCoinDetails(userName string)*CoinDetails{
	time.Sleep(time.Second*1)
	clientCoinData:=CoinDetails{}
	clientCoinData,ok:=mockCoinDetail[userName]
	if !ok{
		return nil
	}
	return &clientCoinData
}
func (d *mockDb) SetUpDatabase() error  {
	return nil
}