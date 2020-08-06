package crsctx

import "github.com/gocraft/web"

type Ctx struct {
}

func (s *Ctx) Cors(rw web.ResponseWriter, req *web.Request, next web.NextMiddlewareFunc) {
	if req.Method == "OPTIONS" {
		rw.Header().Set("Access-Control-Allow-Origin", "http://localhost:8081")
		rw.Header().Set("Access-Control-Allow-Credentials", "true")
		rw.Header().Set("Access-Control-Allow-Methods", "GET,PUT,POST,DELETE,PATCH,OPTIONS")
		rw.Header().Set("Access-Control-Allow-Headers", "Authorization, Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers")

		//rw.Header().Set("Access-Control-Allow-Origin", "http://localhost:8081")

		//rw.Header().Set("Access-Control-Allow-Credentials", "true")
		//rw.Header().Set("Access-Control-Allow-Origin", "http://localhost:8081")
		//rw.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		//rw.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Requested-With")
		return
	} else {
		rw.Header().Set("Access-Control-Allow-Origin", "http://localhost:8081")
		rw.Header().Set("Access-Control-Allow-Credentials", "true")
		//rw.Header().Set("Access-Control-Allow-Origin", "*")
	}
	next(rw, req)
}
