package models

const VerCodeRestorePassword = "ver_code_restore_password"

type VerCode struct {
	ID             int    `db:"id"`
	IdCustomer     int    `db:"id_customer"`
	Code           int    `db:"code"`
	ExpirationTime string `db:"expiration_time"`
}

//type VerCodePass struct {
//	IdCustomer     int    `db:"id_customer"`
//	Code          int `db:"code"`
//	Password string `db:"password"`
//}

//type ItemToken struct {
//	ID             int    `db:"id"`
//	IdCustomer     int    `db:"id_customer"`
//	Token          string `db:"token"`
//	ExpirationTime string `db:"expiration_time"`
//	EmailCustomer  string `db:"email"`
//}

var VerCodeColumns = []string{
	"id_customer",
	"code",
	"expiration_time",
}
