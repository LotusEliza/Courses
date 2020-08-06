package models

const TokenRestorePassword = "token_restore_password"

type Token struct {
	ID             int    `db:"id"`
	IdCustomer     int    `db:"id_customer"`
	Token          string `db:"token"`
	ExpirationTime string `db:"expiration_time"`
}

type ItemToken struct {
	ID             int    `db:"id"`
	IdCustomer     int    `db:"id_customer"`
	Token          string `db:"token"`
	ExpirationTime string `db:"expiration_time"`
	EmailCustomer  string `db:"email"`
}

var TokenColumns = []string{
	"id_customer",
	"token",
	"expiration_time",
}
