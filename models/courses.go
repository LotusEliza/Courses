package models

const CoursesTableName = "courses"
const CustomerCourseTableName = "customer_course"

type Course struct {
	ID          int    `db:"id"`
	Name        string `db:"name"`
	Description string `db:"description"`
	Price       int    `db:"price"`
	Duration    string `db:"duration"`
	Program     string `db:"program"`
	Img         string `db:"img"`
}

var CoursesColumns = []string{
	"date_create",
	"email",
	"name",
	"tel",
	"password",
	"program",
	"img",
}
