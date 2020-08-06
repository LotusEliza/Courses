package models

const QuestionsTableName = "questions"
const ResponsesTableName = "responses"

type ItemInfoRequest struct {
	ID int `db:"id_video"`
	//Sequence int `db:"sequence"`
}

//type Course struct {
//	ID          int    `db:"id"`
//	Name        string `db:"name"`
//	Description string `db:"description"`
//	Price       int    `db:"price"`
//	Duration    string `db:"duration"`
//	Program     string `db:"program"`
//	Img         string `db:"img"`
//}
type QuestResp struct {
	ID        int    `db:"id"`
	Text      string `db:"text"`
	Responses []*Response
}

type Response struct {
	Text    string `db:"text"`
	Correct bool   `db:"correct"`
}

type Question struct {
	Text string `db:"text"`
}

//var CoursesColumns = []string{
//	"date_create",
//	"email",
//	"name",
//	"tel",
//	"password",
//	"program",
//	"img",
//}
