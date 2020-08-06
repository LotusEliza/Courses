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

type (
	VideosResp struct {
		Items []*models.Video `json:"Items"`
	}
)

func VideosByCourseGet(ctx *crsctx.Ctx, rw web.ResponseWriter, req *web.Request) {

	var (
		err      error
		itemInfo *models.ItemInfo
	)

	job := streams.NewJob("VideosByCourseGet")

	itemInfo = new(models.ItemInfo)
	ba, err := ioutil.ReadAll(req.Body)
	if err != nil {
		job.EventErr("ActivatedCoursesGet error: ", err)
		response.ErrorBadRequest(response.ERROR_NO_CONTENT, rw)
		return
	}

	err = json.Unmarshal(ba, itemInfo)
	if err != nil {
		job.EventErr("ActivatedCoursesGet error: ", err)
		response.ErrorBadRequest(response.ERROR_REQUEST_DATA, rw)
		return
	}
	session := db.GetSession()
	var resp VideosResp
	_, err = session.
		Select("*").
		From(models.VideosTableName).
		Where("id_course=?", itemInfo.ID).
		OrderAsc("sequence").
		Load(&resp.Items)
	if err != nil {
		response.Error(response.ERROR_DBERROR, rw)
		job.EventErr("CoursesGet error: ", err)
		return
	}
	response.JsonIntent(resp, rw)
	job.Complete(health.Success)
}

func CurrentVideoUpdate(ctx *crsctx.Ctx, rw web.ResponseWriter, req *web.Request) {

	var (
		err               error
		ItemVidCustCourse *models.ItemVidCustCourse
	)

	job := streams.NewJob("CurrentVideoUpdate")

	ItemVidCustCourse = new(models.ItemVidCustCourse)
	ba, err := ioutil.ReadAll(req.Body)
	if err != nil {
		job.EventErr("CurrentVideoUpdate error: ", err)
		response.ErrorBadRequest(response.ERROR_NO_CONTENT, rw)
		return
	}

	err = json.Unmarshal(ba, ItemVidCustCourse)
	if err != nil {
		job.EventErr("CurrentVideoUpdate error: ", err)
		response.ErrorBadRequest(response.ERROR_REQUEST_DATA, rw)
		return
	}
	session := db.GetSession()
	//var resp VideosResp
	_, err = session.
		Update(models.CustomerCourseVideoTableName).
		Set("current_video", ItemVidCustCourse.CurrentVideo).
		Set("completed", ItemVidCustCourse.Completed).
		Where("id_course=? And id_customer=?", ItemVidCustCourse.IdCourse, ItemVidCustCourse.IdCustomer).
		Exec()
	if err != nil {
		response.Error(response.ERROR_DBERROR, rw)
		job.EventErr("CurrentVideoUpdate error: ", err)
		return
	}
	response.Error("", rw)
	//response.JsonIntent(resp, rw)
	job.Complete(health.Success)
}
