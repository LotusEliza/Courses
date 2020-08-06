package handlers

import (
	"courses/common/db"
	"courses/common/streams"
	"courses/crsctx"
	"courses/models"
	"encoding/json"
	"github.com/finalist736/gokit/response"
	"github.com/gocraft/health"
	"github.com/gocraft/web"
	"io/ioutil"
)

//type (
//	CousesResponse struct {
//		Items []*models.Course `json:"Items"`
//	}
//)
//
type (
	CurrentVideoResp struct {
		Items []*models.ItemCurrentVideo `json:"Items"`
	}
)

func CurrentVideoGet(ctx *crsctx.Ctx, rw web.ResponseWriter, req *web.Request) {

	var (
		err            error
		ItemIdCustomer *models.ItemIdCustomer
	)

	job := streams.NewJob("CurrentVideoGet")

	ItemIdCustomer = new(models.ItemIdCustomer)
	ba, err := ioutil.ReadAll(req.Body)
	if err != nil {
		job.EventErr("CurrentVideoGet error: ", err)
		response.ErrorBadRequest(response.ERROR_NO_CONTENT, rw)
		return
	}

	err = json.Unmarshal(ba, ItemIdCustomer)
	if err != nil {
		job.EventErr("CurrentVideoGet error: ", err)
		response.ErrorBadRequest(response.ERROR_REQUEST_DATA, rw)
		return
	}
	session := db.GetSession()
	var resp CurrentVideoResp
	_, err = session.
		Select("customer_course_video.id_customer, customer_course_video.completed, customer_course_video.id_course, customer_course_video.current_video, videos.id").
		From(models.CustomerCourseVideoTableName).
		LeftJoin(models.VideosTableName, "customer_course_video.current_video=videos.sequence AND customer_course_video.id_course=videos.id_course").
		Where("id_customer=?", ItemIdCustomer.ID).
		Load(&resp.Items)
	if err != nil {
		response.Error(response.ERROR_DBERROR, rw)
		job.EventErr("CoursesGet error: ", err)
		return
	}
	response.JsonIntent(resp, rw)
	job.Complete(health.Success)
	//SELECT customer_course_video.id_customer, customer_course_video.id_course, customer_course_video.current_video, videos.id
	//FROM customer_course_video
	//INNER JOIN videos
	//ON customer_course_video.current_video = videos.sequence
	//WHERE customer_course_video.id_customer=1;

	//SELECT customer_course_video.id_customer, customer_course_video.id_course, customer_course_video.current_video, videos.id
	//FROM customer_course_video
	//INNER JOIN videos
	//ON customer_course_video.current_video = videos.sequence AND customer_course_video.id_course=videos.id_course
	//WHERE customer_course_video.id_customer=1;
}
