package models

const OrdersTableName = "orders"

type OrderInfo struct {
	ID         int    `db:"id"`
	IdCustomer int    `db:"id_customer"`
	IdCourse   int    `db:"id_course"`
	Status     int    `db:"status"`
	DateCreate string `db:"date_create"`
}

var OrderColumns = []string{
	"id_customer",
	"id_course",
	"status",
	"date_create",
}
