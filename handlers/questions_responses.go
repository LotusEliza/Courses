package handlers

//
import (
	"courses/common/db"
	"courses/common/streams"
	"courses/crsctx"
	"courses/models"
	"encoding/json"
	"fmt"

	//"fmt"
	"github.com/finalist736/gokit/response"
	"github.com/gocraft/health"
	"github.com/gocraft/web"
	"io/ioutil"
)

type (
	QuestRespResponse struct {
		Items []*models.QuestResp `json:"Items"`
	}
)
type (
	Questions struct {
		Items []*models.Question `json:"Items"`
	}
)

func QuestRespGet(ctx *crsctx.Ctx, rw web.ResponseWriter, req *web.Request) {

	var (
		err      error
		itemInfo *models.ItemInfoRequest
	)

	job := streams.NewJob("QuestRespGet")

	itemInfo = new(models.ItemInfoRequest)
	ba, err := ioutil.ReadAll(req.Body)
	if err != nil {
		job.EventErr("QuestRespGet error: ", err)
		response.ErrorBadRequest(response.ERROR_NO_CONTENT, rw)
		return
	}

	err = json.Unmarshal(ba, itemInfo)
	if err != nil {
		job.EventErr("QuestRespGet error: ", err)
		response.ErrorBadRequest(response.ERROR_REQUEST_DATA, rw)
		return
	}
	//SELECT COUNT(id) AS Amount FROM questions WHERE id_video=1;
	session := db.GetSession()
	var resp Questions
	_, err = session.
		Select("*").
		From(models.QuestionsTableName).
		Where("id_video=?", itemInfo.ID).
		Load(&resp.Items)
	if err != nil {
		response.Error(response.ERROR_DBERROR, rw)
		job.EventErr("QuestRespGet error: ", err)
		return
	}

	var resp2 QuestRespResponse
	for i := 1; i <= len(resp.Items); i++ {
		fmt.Printf(" %d", i)

		session := db.GetSession()
		var item models.QuestResp
		_, err = session.
			Select("text, id").
			From(models.QuestionsTableName).
			Where("id_video=? AND sequence=?", itemInfo.ID, i).
			Load(&item)
		if err != nil {
			response.Error(response.ERROR_DBERROR, rw)
			job.EventErr("QuestRespGet error: ", err)
			return
		}

		session2 := db.GetSession()
		//var item models.QuestResp
		_, err = session2.
			Select("*").
			From(models.ResponsesTableName).
			Where("id_question=?", &item.ID).
			Load(&item.Responses)
		if err != nil {
			response.Error(response.ERROR_DBERROR, rw)
			job.EventErr("QuestRespGet error: ", err)
			return
		}
		fmt.Printf(" %d", i)
		//response.JsonIntent(item, rw)
		resp2.Items = append(resp2.Items, &item)
	}
	response.JsonIntent(resp2, rw)
	job.Complete(health.Success)
}
