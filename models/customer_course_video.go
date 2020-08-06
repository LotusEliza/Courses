package models

const CustomerCourseVideoTableName = "customer_course_video"

type ItemCurrentVideo struct {
	IdCustomer   int `db:"id_customer"`
	IdCourse     int `db:"id_course"`
	IdVideo      int `db:"id"`
	CurrentVideo int `db:"current_video"`
	Completed    int `db:"completed"`
}

type ItemIdCustomer struct {
	ID int `db:"id_customer"`
}
