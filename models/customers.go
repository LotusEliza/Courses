package models

const CustomersTableName = "customers"

type Info struct {
	ID         int    `db:"id"`
	DateCreate string `db:"date_create"`
	Email      string `db:"email"`
	Name       string `db:"name"`
	Tel        string `db:"tel"`
	Password   string `db:"password"`
}

type Recover struct {
	ID     int    `db:"id"`
	Method string `db:"method"`
	Email  string `db:"email"`
	Tel    string `db:"tel"`
}

type InfoWithToken struct {
	ID       int    `db:"id"`
	Email    string `db:"email"`
	Name     string `db:"name"`
	Password string `db:"password"`
	Token    string `db:"token"`
	Admin    string `db:"admin"`
}

type ItemPass struct {
	ID              int    `db:"id"`
	CurrentPassword string `db:"password"`
	NewPassword     string `db:"password"`
}

type ItemInfo struct {
	ID        int  `db:"id_course"`
	Completed bool `db:"completed"`
}

type ItemPassToken struct {
	Password string `db:"password"`
	Token    string `db:"token"`
}

type RespValid struct {
	Valid     bool `db:"id_course"`
	Completed bool `db:"completed"`
}

var CustomerColumns = []string{
	"date_create",
	"email",
	"name",
	"tel",
	"password",
}
