package handlers

import (
	"bytes"
	"courses/common/db"
	"courses/common/streams"
	"courses/crsctx"
	"courses/models"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/finalist736/gokit/response"
	"github.com/gocraft/dbr"
	"github.com/gocraft/health"
	"github.com/gocraft/web"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"log"
	"math/rand"
	"net"
	"net/http"
	"net/smtp"
	"strconv"
	"time"
)

// Create the JWT key used to create the signature
var jwtKey = []byte("my_secret_key")

// Create a struct that will be encoded to a JWT.
// We add jwt.StandardClaims as an embedded type, to provide fields like expiry time
type Claims struct {
	UserId   int    `json:"userid"`
	Username string `json:"username"`
	Admin    bool   `json:"admin"`
	jwt.StandardClaims
}

type (
	CustomersResponse struct {
		Items []*models.Info `json:"Items"`
	}
)

//-------------------------------------------------------------------------------------------------
/**
 * Registers the customer using landing page form, method - POST
 * @return | nil OR error
 */
func Form(ctx *crsctx.Ctx, rw web.ResponseWriter, req *web.Request) {
	var (
		err      error
		itemForm *models.Info
	)

	job := streams.NewJob("Form")

	current := time.Now()
	hash, _ := HashPassword(req.FormValue("password"))
	itemForm = &models.Info{
		DateCreate: current.Format("2006-01-02T15:04:05.000"),
		Email:      req.FormValue("email"),
		Name:       req.FormValue("name"),
		Tel:        req.FormValue("tel"),
		Password:   hash,
	}

	session := db.GetSession()
	_, err = session.
		InsertInto(models.CustomersTableName).
		Columns(models.CustomerColumns...).
		Record(itemForm).
		Exec()
	if err != nil {
		response.ErrorInternalServer(response.ERROR_DBERROR, rw)
		job.EventErr("Form error", err)
		return
	}
	response.Error("", rw)

	job.Complete(health.Success)
}

//-------------------------------------------------------------------------------------------------
/**
 * Checks credentials and sets a jwt if everything ok, method - POST
 * @return | sets jwt OR error
 */
func UserExist(ctx *crsctx.Ctx, rw web.ResponseWriter, req *web.Request) {
	var (
		err      error
		UserInfo *models.Info
	)

	job := streams.NewJob("UserExist")

	UserInfo = new(models.Info)
	ba, err := ioutil.ReadAll(req.Body)
	if err != nil {
		job.EventErr("UserExist error: ", err)
		response.ErrorBadRequest(response.ERROR_NO_CONTENT, rw)
		return
	}

	err = json.Unmarshal(ba, UserInfo)
	if err != nil {
		job.EventErr("UserExist error: ", err)
		response.ErrorBadRequest(response.ERROR_REQUEST_DATA, rw)
		return
	}

	session := db.GetSession()
	var resp models.InfoWithToken
	err = session.
		Select("*").
		From(models.CustomersTableName).
		Where("name=? and email=?", UserInfo.Name, UserInfo.Email).
		LoadOne(&resp)
	if err == dbr.ErrNotFound {
		fmt.Println("Not Found Such User In DB!!!!!")
		response.Error(response.ERROR_NO_SUCH_USER, rw)
		job.EventErr("UserExist error: ", err)
		return
	}

	result := CheckPasswordHash(UserInfo.Password, resp.Password)
	fmt.Println("Match: ", result)
	fmt.Println("Password From user: ", UserInfo.Password)
	fmt.Println("PAssword Hash from DB: ", resp.Password)

	//--------------------------------JWT----------------------------------------------

	// If a password exists for the given user
	// AND, if it is the same as the password we received, the we can move ahead
	// if NOT, then we return an "Unauthorized" status
	if !result {
		//job.EventErr("UserExist error:", err)
		fmt.Println("Wrong Password!!!!!")
		//rw.WriteHeader(http.StatusUnauthorized)
		response.Error(response.ERROR_WRONG_PASSWORD, rw)
		return
	}
	SetToken(rw, &resp)
	job.Complete(health.Success)
}

func checkToken(rw web.ResponseWriter, req *web.Request) (*Claims, error) {
	c, err := req.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			// If the cookie is not set, return an unauthorized status

			response.Error(response.ERROR_NO_ACCESS_TOKEN, rw)
			//rw.WriteHeader(http.StatusUnauthorized)
			//response.ErrorInternalServer(response.ERROR_NO_ACCESS_TOKEN, rw)
			return nil, err
		}
		// For any other type of error, return a bad request status
		//rw.WriteHeader(http.StatusBadRequest)
		response.ErrorBadRequest(response.ERROR_TOKEN_NOT_FOUND_OR_EXPIRED, rw)
		return nil, err
	}

	// Get the JWT string from the cookie
	tknStr := c.Value

	// Initialize a new instance of `Claims`
	claims := &Claims{}

	// Parse the JWT string and store the result in `claims`.
	// Note that we are passing the key in this method as well. This method will return an error
	// if the token is invalid (if it has expired according to the expiry time we set on sign in),
	// or if the signature does not match
	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			rw.WriteHeader(http.StatusUnauthorized)
			return nil, err
		}
		//rw.WriteHeader(http.StatusBadRequest)
		response.Error(response.ERROR_USER_NOT_CONFIRMED, rw)
		return nil, err
	}
	if !tkn.Valid {
		rw.WriteHeader(http.StatusUnauthorized)
		return nil, err
	}
	return claims, nil
}

//-------------------------------------------------------------------------------------------------
/**
 * Refreshes the jwt if it's life is less then 30 sec, else just 425 - too early, method - POST
 * @return | updates jwt exp time OR error
 */
func Refresh(ctx *crsctx.Ctx, rw web.ResponseWriter, req *web.Request) {

	job := streams.NewJob("Refresh Handler")
	claims, err := checkToken(rw, req)
	if err == nil {
		//// We ensure that a new token is not issued until enough time has elapsed
		//// In this case, a new token will only be issued if the old token is within
		//// 30 seconds of expiry. Otherwise, return a bad request status

		if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > 10*time.Minute {
			//job.EventErr(" Refresh Handler:: StatusBadRequest", err)
			rw.WriteHeader(http.StatusTooEarly)
			return
		}

		//// Now, create a new token for the current use, with a renewed expiration time
		expirationTime := time.Now().Add(2 * time.Hour)
		claims.ExpiresAt = expirationTime.Unix()
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString(jwtKey)
		if err != nil {
			//job.EventErr(" Refresh Handler:: StatusInternalServerError", err)
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		//// Set the new token as the users `token` cookie
		http.SetCookie(rw, &http.Cookie{
			Name:    "token",
			Value:   tokenString,
			Expires: expirationTime,
		})
	}

	job.Complete(health.Success)
}

//-------------------------------------------------------------------------------------------------
/**
 * Selects user info (email, name, ...) by jwt, method - POST
 * @return | email, name, tel etc OR error
 */
func Info(ctx *crsctx.Ctx, rw web.ResponseWriter, req *web.Request) {

	job := streams.NewJob("Info")

	claims, err := checkToken(rw, req)

	session := db.GetSession()
	var resp models.Info
	_, err = session.
		Select("*").
		From(models.CustomersTableName).
		Where("id=?", claims.UserId).
		Load(&resp)
	if err != nil {
		job.EventErr("Info error: ", err)
		response.Error(response.ERROR_DBERROR, rw)
		return
	}
	response.JsonIntent(resp, rw)

	job.Complete(health.Success)
}

//-------------------------------------------------------------------------------------------------
/**
 * Selects user info from db by email and name, generates restore_password_token, sends verification link
 * to email address with token, method - POST
 * @return | sends email to user OR error
 */
func RestoreStart(ctx *crsctx.Ctx, rw web.ResponseWriter, req *web.Request) {
	var (
		err       error
		UserInfo  *models.Info
		ItemToken *models.Token
	)

	job := streams.NewJob("RestoreStart")

	UserInfo = new(models.Info)
	ba, err := ioutil.ReadAll(req.Body)
	if err != nil {
		job.EventErr("RestoreStart error: ", err)
		response.ErrorBadRequest(response.ERROR_NO_CONTENT, rw)
		return
	}

	err = json.Unmarshal(ba, UserInfo)
	if err != nil {
		job.EventErr("RestoreStart error: ", err)
		response.ErrorBadRequest(response.ERROR_REQUEST_DATA, rw)
		return
	}

	session := db.GetSession()
	var resp models.Info
	err = session.
		Select("*").
		From(models.CustomersTableName).
		Where("name=? and email=?", UserInfo.Name, UserInfo.Email).
		LoadOne(&resp)
	if err == dbr.ErrNotFound {
		fmt.Println("Not Found Such User In DB!!!!!")
		response.Error(response.ERROR_NO_SUCH_USER, rw)
		job.EventErr("RestoreStart error: ", err)
		return
	}

	expirationTime := time.Now().Add(10 * time.Minute)
	var randToken = GenerateToken(resp.Email)

	time := expirationTime.Format(time.RFC3339)
	ItemToken = &models.Token{
		IdCustomer:     resp.ID,
		Token:          randToken,
		ExpirationTime: time,
	}
	_, err = session.
		InsertInto(models.TokenRestorePassword).
		Columns(models.TokenColumns...).
		Record(ItemToken).
		Exec()
	if err != nil {
		response.ErrorInternalServer(response.ERROR_DBERROR, rw)
		job.EventErr("RestoreStart error:", err)
		return
	}

	send("http://localhost:8081/#/restore?test="+ItemToken.Token, "lotuselizza@gmail.com")
	job.Complete(health.Success)
}
func CheckVerCode(rw web.ResponseWriter, req *web.Request) {
	var (
		err     error
		VerCode *models.VerCode
	)
	job := streams.NewJob("CheckVerCode")

	VerCode = new(models.VerCode)
	ba, err := ioutil.ReadAll(req.Body)
	if err != nil {
		job.EventErr("CheckVerCode error: ", err)
		response.ErrorBadRequest(response.ERROR_NO_CONTENT, rw)
		return
	}

	err = json.Unmarshal(ba, VerCode)
	if err != nil {
		job.EventErr("CheckVerCode error: ", err)
		response.ErrorBadRequest(response.ERROR_REQUEST_DATA, rw)
		return
	}

	session := db.GetSession()
	var resp models.VerCode
	err = session.
		Select("*").
		From(models.VerCodeRestorePassword).
		Where("id_customer=? and code=?", VerCode.IdCustomer, VerCode.Code).
		LoadOne(&resp)
	if err == dbr.ErrNotFound {
		//fmt.Println("Not Found Such User or Code In DB!!!!!")
		response.Error(response.ERROR_NOT_FOUND, rw)
		job.EventErr("CheckVerCode error: ", err)
		return
	}
	expTime, err := time.Parse(time.RFC3339, resp.ExpirationTime)
	if err != nil {
		job.EventErr("CheckVerCode error: ", err)
		return
	}
	if expTime.Sub(time.Now()) < 30*time.Second {
		fmt.Println("CheckVerCode:: Too late to use this restore pass code...")
		response.Error(response.ERROR_EXPIRED_TOKEN, rw)
		return
	} else {
		fmt.Println("CheckVerCode:: Great time to use this restore pass code...")
		response.Error("", rw)
	}
	//TODO check that code is not exp
	job.Complete(health.Success)

}

func RecoverPassword(ctx *crsctx.Ctx, rw web.ResponseWriter, req *web.Request) {
	var (
		err     error
		Recover *models.Recover
	)

	job := streams.NewJob("RecoverPassword")

	Recover = new(models.Recover)
	ba, err := ioutil.ReadAll(req.Body)
	if err != nil {
		job.EventErr("RestoreStart error: ", err)
		response.ErrorBadRequest(response.ERROR_NO_CONTENT, rw)
		return
	}

	err = json.Unmarshal(ba, Recover)
	if err != nil {
		job.EventErr("RestoreStart error: ", err)
		response.ErrorBadRequest(response.ERROR_REQUEST_DATA, rw)
		return
	}

	rand.Seed(time.Now().UnixNano())
	randomNum := random(1000, 9999)

	//TODO change ExpirationTime to 10 minutes
	itemCode := &models.VerCode{
		IdCustomer:     Recover.ID,
		Code:           randomNum,
		ExpirationTime: time.Now().Add(3 * time.Minute).Format(time.RFC3339),
	}
	session := db.GetSession()
	_, err = session.
		InsertInto(models.VerCodeRestorePassword).
		Columns(models.VerCodeColumns...).
		Record(itemCode).
		Exec()
	if err != nil {
		response.ErrorInternalServer(response.ERROR_DBERROR, rw)
		job.EventErr("Form error", err)
		return
	}
	//response.Error("", rw)
	if Recover.Method == "email" {

		//TODO change email to Recover.Email
		send("the code is "+strconv.Itoa(randomNum), "lotuselizza@gmail.com")

	} else if Recover.Method == "tel" {
		//sendSMS()
		data := `{"recipients":[` + Recover.Tel + `],"sms":{"sender":"Finalist","text":"Ваш код: ` + strconv.Itoa(randomNum) + `!"}}`
		makeRequest(data)
		fmt.Println("Send code with tel number")
	}

	job.Complete(health.Success)
}

func random(min int, max int) int {
	return rand.Intn(max-min) + min
}

//-------------------------------------------------------------------------------------------------
/**
 * Sends email to user email address
 * @return | sends email OR error
 */
func send(body string, email string) {
	job := streams.NewJob("Send")
	from := "lotuselizza@gmail.com"
	pass := "12122elizza"
	to := email

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: Hello there\n\n" +
		body

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))

	if err != nil {
		job.EventErr("Send error", err)
		return
	}
	job.Complete(health.Success)
}

//-------------------------------------------------------------------------------------------------
/**
 * Checks restore_password_token in db and exp. time, method - POST
 * @return | nil OR error
 */
func Confirmation(ctx *crsctx.Ctx, rw web.ResponseWriter, req *web.Request) {

	var (
		err       error
		itemToken *models.Token
	)
	job := streams.NewJob("Confirmation")

	itemToken = new(models.Token)
	ba, err := ioutil.ReadAll(req.Body)
	if err != nil {
		job.EventErr("Confirmation error: ", err)
		response.ErrorBadRequest(response.ERROR_NO_CONTENT, rw)
		return
	}

	err = json.Unmarshal(ba, itemToken)
	if err != nil {
		job.EventErr("Confirmation error: ", err)
		response.ErrorBadRequest(response.ERROR_REQUEST_DATA, rw)
		return
	}

	session := db.GetSession()
	var resp models.ItemToken
	_, err = session.
		Select("token_restore_password.id_customer, token_restore_password.token, token_restore_password.expiration_time, customers.email").
		From(models.TokenRestorePassword).
		LeftJoin(models.CustomersTableName, "customers.id=token_restore_password.id_customer").
		Where("token=?", itemToken.Token).
		Load(&resp)
	if err != nil {
		job.EventErr("Confirmation error", err)
		response.Error(response.ERROR_DBERROR, rw)
		return
	}

	expTime, err := time.Parse(time.RFC3339, resp.ExpirationTime)
	if expTime.Sub(time.Now()) <= 1*time.Second {
		fmt.Println("Expired!!!!!")
		response.Error(response.ERROR_EXPIRED_TOKEN, rw)
	} else {
		fmt.Println("Not Expired!!!!!")
		response.Error("", rw)
		job.Complete(health.Success)
	}
}

//-------------------------------------------------------------------------------------------------
/**
 * Updates password of a user if restore_password_token in db exists,
 * sends jwt with response to front end (login user automatically after password change), method - POST
 * @return | jwt OR error
 */
func UpdateInfo(ctx *crsctx.Ctx, rw web.ResponseWriter, req *web.Request) {
	var (
		err      error
		itemInfo *models.Info
	)
	job := streams.NewJob("UpdateInfo")
	_, err = checkToken(rw, req)
	if err != nil {
		job.EventErr("UpdateInfo error", err)
		response.ErrorInternalServer(response.ERROR_NO_ACCESS_TOKEN, rw)
		return
	}

	itemInfo = new(models.Info)
	ba, err := ioutil.ReadAll(req.Body)
	if err != nil {
		job.EventErr("UpdateInfo error: ", err)
		response.ErrorBadRequest(response.ERROR_NO_CONTENT, rw)
		return
	}

	err = json.Unmarshal(ba, itemInfo)
	if err != nil {
		job.EventErr("UpdateInfo error: ", err)
		response.ErrorBadRequest(response.ERROR_REQUEST_DATA, rw)
		return
	}

	//hash, _ := HashPassword(itemInfo.Password)

	session := db.GetSession()
	_, err = session.
		Update(models.CustomersTableName).
		Set("name", itemInfo.Name).
		Set("email", itemInfo.Email).
		Set("tel", itemInfo.Tel).
		Where("id=?", itemInfo.ID).
		Exec()
	if err != nil {
		job.EventErr("UpdateInfo error", err)
		response.ErrorInternalServer(response.ERROR_DBERROR, rw)
		return
	}
	response.Error("", rw)
	job.Complete(health.Success)
}

//Update password after ver code is checked
func SetNewPassword(rw web.ResponseWriter, req *web.Request) {
	var (
		err      error
		itemInfo *models.Info
	)
	job := streams.NewJob("SetNewPassword")
	_, err = checkToken(rw, req)
	if err != nil {
		job.EventErr("SetNewPassword error", err)
		response.ErrorInternalServer(response.ERROR_NO_ACCESS_TOKEN, rw)
		return
	}

	itemInfo = new(models.Info)
	ba, err := ioutil.ReadAll(req.Body)
	if err != nil {
		job.EventErr("SetNewPassword error: ", err)
		response.ErrorBadRequest(response.ERROR_NO_CONTENT, rw)
		return
	}

	err = json.Unmarshal(ba, itemInfo)
	if err != nil {
		job.EventErr("SetNewPassword error: ", err)
		response.ErrorBadRequest(response.ERROR_REQUEST_DATA, rw)
		return
	}

	hash, _ := HashPassword(itemInfo.Password)

	session := db.GetSession()
	_, err = session.
		Update(models.CustomersTableName).
		Set("password", hash).
		Where("id=?", itemInfo.ID).
		Exec()
	if err != nil {
		job.EventErr("SetNewPassword error", err)
		response.ErrorInternalServer(response.ERROR_DBERROR, rw)
		return
	}

	session = db.GetSession()
	_, err = session.
		DeleteFrom(models.VerCodeRestorePassword).
		Where("id_customer=?", itemInfo.ID).
		Exec()
	if err != nil {
		job.EventErr("SetNewPassword error", err)
		response.ErrorInternalServer(response.ERROR_DBERROR, rw)
		return
	}

	response.Error("", rw)
	job.Complete(health.Success)
}

func SavePassword(ctx *crsctx.Ctx, rw web.ResponseWriter, req *web.Request) {
	var (
		err      error
		ItemPass *models.ItemPass
	)
	job := streams.NewJob("SavePassword")
	_, err = checkToken(rw, req)
	if err != nil {
		//job.EventErr("SavePassword error", err)
		//response.ErrorInternalServer(response.ERROR_NO_ACCESS_TOKEN, rw)
		return
	}

	ItemPass = new(models.ItemPass)
	ba, err := ioutil.ReadAll(req.Body)
	if err != nil {
		job.EventErr("SavePassword error: ", err)
		response.ErrorBadRequest(response.ERROR_NO_CONTENT, rw)
		return
	}

	err = json.Unmarshal(ba, ItemPass)
	if err != nil {
		job.EventErr("SavePassword error: ", err)
		response.ErrorBadRequest(response.ERROR_REQUEST_DATA, rw)
		return
	}

	session := db.GetSession()
	var resp models.Info
	_, err = session.
		Select("*").
		From(models.CustomersTableName).
		Where("id=?", ItemPass.ID).
		Load(&resp)
	if err != nil {
		response.Error(response.ERROR_DBERROR, rw)
		job.EventErr("UpdatePassword error: ", err)
		return
	}

	//CheckPasswordHash
	//hash, _ := HashPa
	//hash, _ := HashPassword(ItemPass.CurrentPassword)
	result := CheckPasswordHash(ItemPass.CurrentPassword, resp.Password)

	fmt.Println("Hash from user:::", ItemPass.CurrentPassword)
	fmt.Println("Hash from db:::", resp.Password)

	if result != true {
		job.EventErr("SavePassword error: No such User ", err)
		rw.WriteHeader(http.StatusUnauthorized)
		return
	}

	newHash, _ := HashPassword(ItemPass.NewPassword)
	fmt.Println("newHash from user:::", newHash)
	fmt.Println("new pass:::", ItemPass.NewPassword)
	fmt.Println("old hash pass from db:::", resp.Password)
	session = db.GetSession()
	_, err = session.
		Update(models.CustomersTableName).
		Set("password", newHash).
		Where("id=?", ItemPass.ID).
		Exec()
	if err != nil {
		job.EventErr("SavePassword error", err)
		response.ErrorInternalServer(response.ERROR_DBERROR, rw)
		return
	}

	response.Error("", rw)
	job.Complete(health.Success)
}

//-------------------------------------------------------------------------------------------------
/**
 * Updates password of a user if restore_password_token in db exists,
 * sends jwt with response to front end (login user automatically after password change), method - POST
 * @return | jwt OR error
 */
func UpdatePassword(ctx *crsctx.Ctx, rw web.ResponseWriter, req *web.Request) {
	var (
		err      error
		itemInfo *models.ItemPassToken
	)
	job := streams.NewJob("Confirmation")

	itemInfo = new(models.ItemPassToken)
	ba, err := ioutil.ReadAll(req.Body)
	if err != nil {
		job.EventErr("UpdatePassword error: ", err)
		response.ErrorBadRequest(response.ERROR_NO_CONTENT, rw)
		return
	}

	err = json.Unmarshal(ba, itemInfo)
	if err != nil {
		job.EventErr("UpdatePassword error: ", err)
		response.ErrorBadRequest(response.ERROR_REQUEST_DATA, rw)
		return
	}

	hash, _ := HashPassword(itemInfo.Password)

	session := db.GetSession()
	var resp models.Token
	_, err = session.
		Select("*").
		From(models.TokenRestorePassword).
		Where("token=?", itemInfo.Token).
		Load(&resp)
	if err != nil {
		response.Error(response.ERROR_DBERROR, rw)
		job.EventErr("UpdatePassword error: ", err)
		return
	}

	session = db.GetSession()
	_, err = session.
		Update(models.CustomersTableName).
		Set("password", hash).
		Where("id=?", resp.IdCustomer).
		Exec()
	if err != nil {
		job.EventErr("UpdatePassword error", err)
		response.ErrorInternalServer(response.ERROR_DBERROR, rw)
		return
	}
	session = db.GetSession()
	var respUser models.InfoWithToken
	_, err = session.
		Select("*").
		From(models.CustomersTableName).
		Where("id=?", resp.IdCustomer).
		Load(&respUser)
	if err != nil {
		response.Error(response.ERROR_DBERROR, rw)
		job.EventErr("UserExist error: ", err)
		return
	}
	session = db.GetSession()
	_, err = session.
		DeleteFrom(models.TokenRestorePassword).
		Where("id_customer=?", resp.IdCustomer).
		Exec()
	if err != nil {
		job.EventErr("Confirmation error", err)
		response.ErrorInternalServer(response.ERROR_DBERROR, rw)
		return
	}
	//response.Error("", rw)
	SetToken(rw, &respUser)
	job.Complete(health.Success)
}

//-------------------------------------------------------------------------------------------------
/**
 * Returns a unique token based on the email string
 * @return | token for restoring password OR error
 */
func GenerateToken(email string) string {
	job := streams.NewJob("GenerateToken")
	hash, err := bcrypt.GenerateFromPassword([]byte(email), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Hash to store:", string(hash))

	hasher := md5.New()
	hasher.Write(hash)
	job.Complete(health.Success)
	return hex.EncodeToString(hasher.Sum(nil))
}

//-------------------------------------------------------------------------------------------------
/**
 * Hashes password (base64)
 * @return | string OR error
 */
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

//-------------------------------------------------------------------------------------------------
/**
 * Compares two passwords that hashed by base64
 * @return | true OR false
 */
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

//-------------------------------------------------------------------------------------------------
/**
 * Sets the client cookie for "token" as the generated JWT, method - POST
 * @return | sets jwt OR error
 */
func SetToken(rw web.ResponseWriter, resp *models.InfoWithToken) {
	// Declare the expiration time of the token
	// here, we have kept it as 5 minutes
	fmt.Println("Set Token::::   resp", resp)
	job := streams.NewJob("SetToken")
	expirationTime := time.Now().Add(2 * time.Hour)
	// Create the JWT claims, which includes the username and expiry time
	claims := &Claims{
		UserId: resp.ID,
		//TODO change table customers add field admin bool
		Admin:    true,
		Username: resp.Name,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}
	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		job.EventErr("UserExist error: ", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Finally, we set the client cookie for "token" as the JWT we just generated
	// we also set an expiry time which is the same as the token itself
	http.SetCookie(rw, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})
}

//************************************************VIBER**********************************************
//func sendSMS() {
//
//}
func makeRequest(jsonData string) {
	tr := &http.Transport{
		MaxIdleConns:          10,
		IdleConnTimeout:       30 * time.Second,
		DisableCompression:    true,
		ExpectContinueTimeout: time.Second * 10,
		ResponseHeaderTimeout: time.Second * 10,
		TLSHandshakeTimeout:   time.Second * 10,
		DialContext: (&net.Dialer{
			Timeout: time.Second * 10,
		}).DialContext,
	}
	client := &http.Client{
		Transport: tr,
	}
	client.Timeout = time.Second * 15

	//data := `{"recipients":["380936161694"],"sms":{"sender":"Finalist","text":"Finalist TEST FROM GOLANG!"}}`
	data := jsonData

	buff := bytes.NewBufferString(data)
	req, err := http.NewRequest(
		"POST",
		"https://api.turbosms.ua/message/send.json",
		buff)
	if err != nil {
		panic(err)
	}
	// Create a Bearer string by appending string access token
	var bearer = "Bearer " + "fb05ef460d4408803976319949f7d228e6579a3b"
	// add authorization header to the req
	req.Header.Set("Authorization", bearer)
	req.Header.Set("Content-type", "application/json")
	//req.Header.Add("Authorization", bearer)
	fmt.Println("req.Header")
	fmt.Println(req.Header)
	fmt.Println("req")
	fmt.Println(req)

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	ba, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	log.Printf("response: %s\n\n", ba)
}

func CustomersGet(ctx *crsctx.Ctx, rw web.ResponseWriter, req *web.Request) {

	var (
		err error
	)

	job := streams.NewJob("CustomersGet")

	session := db.GetSession()
	var resp CustomersResponse
	_, err = session.
		Select("*").
		From(models.CustomersTableName).
		Load(&resp.Items)
	if err != nil {
		response.Error(response.ERROR_DBERROR, rw)
		job.EventErr("CustomersGet error: ", err)
		return
	}
	response.JsonIntent(resp, rw)
	job.Complete(health.Success)
}


func CustomerRemove(rw web.ResponseWriter, req *web.Request) {
	var (
		err      error
		UserInfo *models.Info
	)

	job := streams.NewJob("UserRemove")

	UserInfo = new(models.Info)
	ba, err := ioutil.ReadAll(req.Body)
	if err != nil {
		job.EventErr("UserRemove error: ", err)
		response.ErrorBadRequest(response.ERROR_NO_CONTENT, rw)
		return
	}

	err = json.Unmarshal(ba, UserInfo)
	if err != nil {
		job.EventErr("UserRemove error: ", err)
		response.ErrorBadRequest(response.ERROR_REQUEST_DATA, rw)
		return
	}

	session := db.GetSession()
	_, err = session.
		DeleteFrom(models.CustomersTableName).
		//LeftJoin(models.CustomerCourseVideoTableName, "customers.id=customer_course_video.id_customer").
		Where("id=?", UserInfo.ID).
		Exec()
	if err != nil {
		job.EventErr("UserRemove error", err)
		response.ErrorInternalServer(response.ERROR_DBERROR, rw)
		return
	}

	response.Error("", rw)
	job.Complete(health.Success)
}
//************************************************VIBER**********************************************
//
////-------------------------------------------------------------------------------------------------
//func generatePassword(passwordLength, minSpecialChar, minNum, minUpperCase int) string {
//
//	job := streams.NewJob("generatePassword")
//	var password strings.Builder
//
//	//Set special character
//	for i := 0; i < minSpecialChar; i++ {
//		random := rand.Intn(len(specialCharSet))
//		password.WriteString(string(specialCharSet[random]))
//	}
//
//	//Set numeric
//	for i := 0; i < minNum; i++ {
//		random := rand.Intn(len(numberSet))
//		password.WriteString(string(numberSet[random]))
//	}
//
//	//Set uppercase
//	for i := 0; i < minUpperCase; i++ {
//		random := rand.Intn(len(upperCharSet))
//		password.WriteString(string(upperCharSet[random]))
//	}
//
//	remainingLength := passwordLength - minSpecialChar - minNum - minUpperCase
//	for i := 0; i < remainingLength; i++ {
//		random := rand.Intn(len(allCharSet))
//		password.WriteString(string(allCharSet[random]))
//	}
//	inRune := []rune(password.String())
//	rand.Shuffle(len(inRune), func(i, j int) {
//		inRune[i], inRune[j] = inRune[j], inRune[i]
//	})
//	job.Complete(health.Success)
//	return string(inRune)
//}

//Password generation settings for restoring pas
//var (
//	lowerCharSet   = "abcdedfghijklmnopqrst"
//	upperCharSet   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
//	specialCharSet = "!@#$%&*"
//	numberSet      = "0123456789"
//	allCharSet     = lowerCharSet + upperCharSet + specialCharSet + numberSet
//)

//
////-------------------------------------------------------------------------------------------------
//func SendEmail(ctx *crsctx.Ctx, rw web.ResponseWriter, req *web.Request) {
//	var (
//		err       error
//		UserInfo  *models.Info
//		ItemToken *models.Token
//	)
//
//	job := streams.NewJob("SendEmail")
//
//	UserInfo = new(models.Info)
//	ba, err := ioutil.ReadAll(req.Body)
//	if err != nil {
//		job.EventErr("SendEmail error: ", err)
//		response.ErrorBadRequest(response.ERROR_NO_CONTENT, rw)
//		return
//	}
//
//	err = json.Unmarshal(ba, UserInfo)
//	if err != nil {
//		job.EventErr("SendEmail error: ", err)
//		response.ErrorBadRequest(response.ERROR_REQUEST_DATA, rw)
//		return
//	}
//	session := db.GetSession()
//	var resp models.Info
//	_, err = session.
//		Select("*").
//		From(models.CustomersTableName).
//		Where("name=? and email=?", UserInfo.Name, UserInfo.Email).
//		Load(&resp)
//	if err != nil {
//		response.Error(response.ERROR_DBERROR, rw)
//		job.EventErr("SendEmail error: ", err)
//		return
//	}
//	//if resp.Name == "" {
//	//	job.EventErr("SendEmail error::: No such User:", err)
//	//	response.ErrorNotAuthorized(response.ERROR_DBERROR, rw)
//	//	return
//	//}
//
//	// Declare the expiration time of the token
//	// here, we have kept it as 5 minutes
//	expirationTime := time.Now().Add(1 * time.Minute)
//	var randToken = GenerateToken(resp.Email)
//
//	time := expirationTime.Format(time.RFC3339)
//	ItemToken = &models.Token{
//		IdCustomer:     resp.ID,
//		Token:          randToken,
//		ExpirationTime: time,
//	}
//	_, err = session.
//		InsertInto(models.TokenRestorePassword).
//		Columns(models.TokenColumns...).
//		Record(ItemToken).
//		Exec()
//	if err != nil {
//		response.ErrorInternalServer(response.ERROR_DBERROR, rw)
//		job.EventErr("SendEmail error::: insert token", err)
//		return
//	}
//
//	//send("http://localhost:3000/confirmation/"+ItemToken.Token, "lotuselizza@gmail.com")
//	send("http://localhost:8081/#/restore?test="+ItemToken.Token, "lotuselizza@gmail.com")
//	//TODO save info (token and exp time to db)
//	//TODO check token before each function???
//	job.Complete(health.Success)
//}

////-------------------------------------------------------------------------------------------------
//func Welcome(ctx *crsctx.Ctx, rw web.ResponseWriter, req *web.Request) {
//	// We can obtain the session token from the requests cookies, which come with every request
//	c, err := req.Cookie("token")
//	if err != nil {
//		if err == http.ErrNoCookie {
//			// If the cookie is not set, return an unauthorized status
//			rw.WriteHeader(http.StatusUnauthorized)
//			return
//		}
//		// For any other type of error, return a bad request status
//		rw.WriteHeader(http.StatusBadRequest)
//		return
//	}
//
//	// Get the JWT string from the cookie
//	tknStr := c.Value
//
//	// Initialize a new instance of `Claims`
//	claims := &Claims{}
//
//	// Parse the JWT string and store the result in `claims`.
//	// Note that we are passing the key in this method as well. This method will return an error
//	// if the token is invalid (if it has expired according to the expiry time we set on sign in),
//	// or if the signature does not match
//	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
//		return jwtKey, nil
//	})
//	if err != nil {
//		if err == jwt.ErrSignatureInvalid {
//			rw.WriteHeader(http.StatusUnauthorized)
//			return
//		}
//		rw.WriteHeader(http.StatusBadRequest)
//		return
//	}
//	if !tkn.Valid {
//		rw.WriteHeader(http.StatusUnauthorized)
//		return
//	}
//
//	// Finally, return the welcome message to the user, along with their
//	// username given in the token
//
//	//rw.Redirect(rw, req, url,  code)
//}
