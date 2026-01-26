package tools
type LoginDetails struct{
	Token string
	UserName string
}
type CoinDetails struct{
	Coins int64
	UserName string
}
type DatabaseInterface interface{
	GetUserDetails(username string) *LoginDetails
	GetCoinDetails(username string) *CoinDetails
	SetUpDatabase() error

}
func NewDatabase() (*DatabaseInterface,error){
	var db DatabaseInterface=&mockDb{}
	err:=db.SetUpDatabase()
	if err!=nil{
		return nil,err
	}
	return &db,nil
}