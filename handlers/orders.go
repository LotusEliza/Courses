package handlers

import (
	"courses/common/db"
	"courses/common/streams"
	"courses/crsctx"
	"courses/models"
	sha1 "crypto/sha1"
	b64 "encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/finalist736/gokit/response"
	"github.com/gocraft/health"
	"github.com/gocraft/web"
	"io/ioutil"
	"log"
	"strconv"
	"time"
)

type Data struct {
	//Name          string    `json:"name"`
	PublicKey   string `db:"public_key" json:"public_key"`
	Version     int    `db:"version" json:"version"`
	Action      string `db:"action" json:"action"`
	Amount      int    `db:"amount" json:"amount"`
	Currency    string `db:"currency" json:"currency"`
	Description string `db:"description" json:"description"`
	OrderId     string `db:"id" json:"order_id"`
	ServerUrl   string `db:"id" json:"server_url"`
}

type RespOrder struct {
	Data      string `db:"data" json:"data"`
	Signature string `db:"signature" json:"signature"`
}

type (
	OrdersResponse struct {
		Items []*models.OrderInfo `json:"Items"`
	}
)

func PlaceOrder(ctx *crsctx.Ctx, rw web.ResponseWriter, req *web.Request) {

	var (
		err       error
		itemOrder *models.OrderInfo
	)

	job := streams.NewJob("PlaceOrder")

	current := time.Now()
	itemOrder = &models.OrderInfo{
		Status:     0,
		DateCreate: current.Format("2006-01-02T15:04:05.000"),
	}

	ba, err := ioutil.ReadAll(req.Body)
	if err != nil {
		job.EventErr("PlaceOrder error: ", err)
		response.ErrorBadRequest(response.ERROR_NO_CONTENT, rw)
		return
	}

	err = json.Unmarshal(ba, itemOrder)
	if err != nil {
		job.EventErr("PlaceOrder error: ", err)
		response.ErrorBadRequest(response.ERROR_REQUEST_DATA, rw)
		return
	}

	session := db.GetSession()
	_, err = session.
		InsertInto(models.OrdersTableName).
		Columns(models.OrderColumns...).
		Record(itemOrder).
		Exec()
	if err != nil {
		response.ErrorInternalServer(response.ERROR_DBERROR, rw)
		job.EventErr("error place order", err)
		return
	}

	id := getLastId()
	rawDataData := &Data{
		PublicKey:   "sandbox_i79527556721",
		Version:     3,
		Action:      "pay",
		Amount:      1,
		Currency:    "UAH",
		Description: "test",
		OrderId:     strconv.Itoa(id),
		ServerUrl:   "http://localhost:3000/api_call_back",
	}

	var jsonData []byte
	jsonData, err = json.Marshal(rawDataData)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(jsonData))
	data := b64.StdEncoding.EncodeToString(jsonData)
	fmt.Println(data)
	privateKey := "sandbox_9Cis7YaWEp5RHjmmq8c0FvYot35LbiT9kmJuQLBj"
	//sha := HashStr(privateKey+data+privateKey)
	str := privateKey + data + privateKey
	h := sha1.New()
	h.Write([]byte(str))
	bs := h.Sum(nil)
	signature := b64.StdEncoding.EncodeToString(bs)
	fmt.Println("signature")
	fmt.Println(signature)

	resp := &RespOrder{
		Data:      data,
		Signature: signature,
	}
	response.JsonIntent(resp, rw)
	//job.Complete(health.Success)

}

// Get sha1 from string
func HashStr(Txt string) string {
	h := sha1.New()
	h.Write([]byte(Txt))
	sha1Hash := hex.EncodeToString(h.Sum(nil))
	return sha1Hash
}

//type RespOrder struct {
//	ID         int    `db:"id"`
//	public_key string  `db:"public_key"`
//	version string  `db:"version"`
//	action string  `db:"action"`
//	amount string  `db:"amount"`
//	currency string  `db:"currency"`
//	description string  `db:"description"`
//	order_id string  `db:"id"`
//	//"version":"3",
//	//"action":"pay",
//	//"amount":"1",
//	//"currency":"UAH",
//	//"description":"test",
//	//"order_id":"000001"
//}

func CallBack(ctx *crsctx.Ctx, rw web.ResponseWriter, req *web.Request) {
	job := streams.NewJob("CallBack")
	ba, err := ioutil.ReadAll(req.Body)
	if err != nil {
		job.EventErr("CallBack error: ", err)
		response.ErrorBadRequest(response.ERROR_NO_CONTENT, rw)
		return
	}
	fmt.Println("Call Back handler::")
	fmt.Println(string(ba))
}
func getLastId() int {
	var (
		err error
	)

	job := streams.NewJob("GetLastOrderId")

	session := db.GetSession()
	var resp int
	_, err = session.
		Select("id").
		From(models.OrdersTableName).
		OrderDesc("id").
		Limit(1).
		Load(&resp)
	if err != nil {
		//response.Error(response.ERROR_DBERROR, rw)
		job.EventErr("GetLastOrderId Error: ", err)
	}
	job.Complete(health.Success)
	return resp
}

func OrdersGet(rw web.ResponseWriter, req *web.Request) {

	var (
		err error
	)

	job := streams.NewJob("OrdersGet")

	session := db.GetSession()
	var resp OrdersResponse
	_, err = session.
		Select("*").
		From(models.OrdersTableName).
		Load(&resp.Items)
	if err != nil {
		response.Error(response.ERROR_DBERROR, rw)
		job.EventErr("OrdersGet error: ", err)
		return
	}
	response.JsonIntent(resp, rw)
	job.Complete(health.Success)
}

func OrderRemove(rw web.ResponseWriter, req *web.Request) {
	var (
		err      error
		OrderItem *models.OrderInfo
	)

	job := streams.NewJob("OrderRemove")

	OrderItem = new(models.OrderInfo)
	ba, err := ioutil.ReadAll(req.Body)
	if err != nil {
		job.EventErr("OrderRemove error: ", err)
		response.ErrorBadRequest(response.ERROR_NO_CONTENT, rw)
		return
	}

	err = json.Unmarshal(ba, OrderItem)
	if err != nil {
		job.EventErr("OrderRemove error: ", err)
		response.ErrorBadRequest(response.ERROR_REQUEST_DATA, rw)
		return
	}

	session := db.GetSession()
	_, err = session.
		DeleteFrom(models.OrdersTableName).
		//LeftJoin(models.CustomerCourseVideoTableName, "customers.id=customer_course_video.id_customer").
		Where("id=?", OrderItem.ID).
		Exec()
	if err != nil {
		job.EventErr("OrderRemove error", err)
		response.ErrorInternalServer(response.ERROR_DBERROR, rw)
		return
	}

	response.Error("", rw)
	job.Complete(health.Success)
}
