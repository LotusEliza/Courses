package main

import (
	"courses/common/db"
	"courses/common/streams"
	"courses/crsctx"
	"courses/handlers"
	"github.com/gocraft/health"
	"github.com/gocraft/web"
	"github.com/pressly/goose"
	"math/rand"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	//stream for logs
	job := streams.NewJob("starting")

	err := db.Init()
	if err != nil {
		job.EventErr("db connection", err)
		os.Exit(1)
	}

	goose.SetDialect("mysql")
	err = goose.Up(db.GetDB().DB, "./migrations")
	if err != nil {
		job.EventErr("goose up", err)
		os.Exit(1)
	}

	//---------------------ADMIN SERVER-------------------------------
	router := web.New(crsctx.Ctx{})
	router.Middleware((*crsctx.Ctx).Cors)
	//router.Post("/requests", handlers.RequestsGet)
	router.Post("/process", handlers.Form)
	router.Post("/courses", handlers.CoursesGet)
	router.Post("/courses/activated", handlers.ActivatedCoursesGet)
	router.Post("/videos", handlers.VideosByCourseGet)
	router.Post("/quiz", handlers.QuestRespGet)
	router.Post("/customer_course_video", handlers.CurrentVideoGet)
	router.Post("/current_video_update", handlers.CurrentVideoUpdate)
	router.Post("/place_order", handlers.PlaceOrder)
	router.Post("/api_call_back", handlers.CallBack)
	router.Post("/user_exist", handlers.UserExist)
	//router.Post("/welcome", handlers.Welcome)
	router.Post("/refresh", handlers.Refresh)
	router.Post("/info", handlers.Info)
	router.Post("/restore_start", handlers.RestoreStart)
	router.Post("/confirmation", handlers.Confirmation)
	router.Post("/update_password", handlers.UpdatePassword)
	router.Post("/info_update", handlers.UpdateInfo)
	router.Post("/save_password", handlers.SavePassword)
	router.Post("/recover_password", handlers.RecoverPassword)
	router.Post("/check_ver_code", handlers.CheckVerCode)
	router.Post("/set_new_password", handlers.SetNewPassword)
	//---------------------------Admin--------------------------------
	router.Post("/get_customers", handlers.CustomersGet)
	router.Post("/get_orders", handlers.OrdersGet)
	router.Post("/customer_remove", handlers.CustomerRemove)
	router.Post("/course_remove", handlers.CourseRemove)
	router.Post("/order_remove", handlers.OrderRemove)

	var (
		port     string
		serv     *http.Server
		listener net.Listener
	)
	port = os.Getenv("PORT")
	if port == "" {
		port = ":3000"
	} else if !strings.Contains(port, ":") {
		port = ":" + port
	}
	serv = &http.Server{Addr: port,
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	job.EventKv("listening", health.Kvs{"addr": port})
	serv.SetKeepAlivesEnabled(false)
	listener, err = net.Listen("tcp", serv.Addr)
	if err != nil {
		job.EventErr("net.Listen", err)
		os.Exit(1)
	}
	go serv.Serve(listener)
	job.Event("api started")

	var (
		staticDir  string
		staticPort string
	)

	staticDir = os.Getenv("STATIC_FILES_DIR")
	if staticDir == "" {
		staticDir = "./static"
	}
	staticPort = os.Getenv("STATIC_FILES_PORT")
	if staticPort == "" {
		staticPort = ":8000"
	} else if !strings.Contains(staticPort, ":") {
		staticPort = ":" + staticPort
	}
	job.EventKv("staticDir", health.Kvs{"path": staticDir})
	job.EventKv("staticPort", health.Kvs{"addr": staticPort})

	fs := http.FileServer(http.Dir(staticDir))
	http.Handle("/", fs)
	//http.HandleFunc("/process", handlers.Form)
	//http.HandleFunc("/static", handlers.UserInfoSave)
	go http.ListenAndServe(staticPort, nil)
	job.Event("static started")

	job.Event("wait for os signals")
	interruptChannel := make(chan os.Signal)
	signal.Notify(interruptChannel, os.Interrupt, os.Kill, syscall.SIGTERM)
	for {
		select {
		case killSignal := <-interruptChannel:
			job.EventKv("signal", health.Kvs{"signal": killSignal.String()})
			if killSignal == os.Interrupt {
				job.Event("interrupted")
				job.Complete(health.Success)
				return
			}
			job.Event("killed")
			job.Complete(health.Success)
			return
		}
	}
}
