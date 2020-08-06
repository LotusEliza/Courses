package models

const VideosTableName = "videos"

type Video struct {
	ID       int    `db:"id"`
	IdCourse int    `db:"id_course"`
	Title    string `db:"title"`
	Name     string `db:"name"`
	Sequence int    `db:"sequence"`
	Poster   string `db:"poster"`
}

type ItemVidCustCourse struct {
	IdCourse     int `db:"id_course"`
	IdCustomer   int `db:"id_course"`
	CurrentVideo int `db:"current_video"`
	Completed    int `db:"completed"`
}

var VideoColumns = []string{
	"id_course",
	"title",
	"name",
	"sequence",
	"poster",
}
