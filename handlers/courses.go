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
	CousesResponse struct {
		Items []*models.Course `json:"Items"`
	}
)

type (
	ActiveCourseResp struct {
		Items []*models.ItemInfo `json:"Items"`
	}
)

func CoursesGet(ctx *crsctx.Ctx, rw web.ResponseWriter, req *web.Request) {

	var (
		err error
	)

	job := streams.NewJob("get courses")

	session := db.GetSession()
	var resp CousesResponse
	_, err = session.
		Select("*").
		From(models.CoursesTableName).
		Load(&resp.Items)
	if err != nil {
		response.Error(response.ERROR_DBERROR, rw)
		job.EventErr("CoursesGet error: ", err)
		return
	}
	response.JsonIntent(resp, rw)
	job.Complete(health.Success)
}

func ActivatedCoursesGet(ctx *crsctx.Ctx, rw web.ResponseWriter, req *web.Request) {

	var (
		err      error
		itemInfo *models.ItemInfo
	)

	job := streams.NewJob("get courses")

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
	var resp ActiveCourseResp
	_, err = session.
		Select("customer_course.id_course, customer_course_video.completed").
		From(models.CustomerCourseTableName).
		LeftJoin(models.CustomerCourseVideoTableName, "customer_course.id_course=customer_course_video.id_course").
		Where("customer_course.id_customer=?", itemInfo.ID).
		Load(&resp.Items)
	if err != nil {
		response.Error(response.ERROR_DBERROR, rw)
		job.EventErr("CoursesGet error: ", err)
		return
	}
	//SELECT customer_course.id_course, customer_course_video.completed
	//FROM customer_course
	//LEFT JOIN customer_course_video ON customer_course.id_course = customer_course_video.id_course
	//WHERE customer_course.id_customer = 1

	response.JsonIntent(resp, rw)
	job.Complete(health.Success)
}

func CourseRemove(rw web.ResponseWriter, req *web.Request) {
	var (
		err      error
		CourseItem *models.Course
	)

	job := streams.NewJob("CourseRemove")

	CourseItem = new(models.Course)
	ba, err := ioutil.ReadAll(req.Body)
	if err != nil {
		job.EventErr("CourseRemove error: ", err)
		response.ErrorBadRequest(response.ERROR_NO_CONTENT, rw)
		return
	}

	err = json.Unmarshal(ba, CourseItem)
	if err != nil {
		job.EventErr("CourseRemove error: ", err)
		response.ErrorBadRequest(response.ERROR_REQUEST_DATA, rw)
		return
	}

	session := db.GetSession()
	_, err = session.
		DeleteFrom(models.CoursesTableName).
		//LeftJoin(models.CustomerCourseVideoTableName, "customers.id=customer_course_video.id_customer").
		Where("id=?", CourseItem.ID).
		Exec()
	if err != nil {
		job.EventErr("CourseRemove error", err)
		response.ErrorInternalServer(response.ERROR_DBERROR, rw)
		return
	}

	response.Error("", rw)
	job.Complete(health.Success)
}
