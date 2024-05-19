package main

import (
	"fmt"
	"log"
	"net/http"
	"poc-go-dynamodb/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.TestMode) //TODO: Change to release mode

	routerInit := routers.InitRouter()
	endPoint := fmt.Sprintf(":%d", 8080)

	sever := &http.Server{
		Addr:    endPoint,
		Handler: routerInit,
	}

	log.Printf("[info] start http server listening %s", endPoint)

	sever.ListenAndServe()
}
