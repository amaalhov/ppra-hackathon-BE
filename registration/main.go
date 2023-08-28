package main

import (
	"log"
	"net/http"

	"github.com/uptrace/bunrouter"
)

func main() {
	router := bunrouter.New()
	router.GET("/", func(w http.ResponseWriter, req bunrouter.Request) error {
		// req embeds *http.Request and has all the same fields and methods
		log.Println(req.Method, req.Route(), req.Params().Map())
		return bunrouter.JSON(w, bunrouter.H{
			"message": "success",
			"status":  true,
		})
	})
	log.Println("registration listening on http://localhost:3000")
	log.Println(http.ListenAndServe(":3000", router))
}
