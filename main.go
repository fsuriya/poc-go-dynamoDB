package main

import (
	"fmt"
	"log"
	"net/http"

	"poc-go-dynamodb/database"
	"poc-go-dynamodb/routers"

	"github.com/gin-gonic/gin"
)

func init() {
	_, err := database.ConnectDynamo()
	if err != nil {
		log.Fatalf("failed to connect to DynamoDB: %v", err)
	}
}

func main() {
	gin.SetMode(gin.ReleaseMode)

	routerInit := routers.InitRouter()
	endPoint := fmt.Sprintf(":%d", 3000)

	sever := &http.Server{
		Addr:    endPoint,
		Handler: routerInit,
	}

	log.Printf("[info] start http server listening %s", endPoint)

	sever.ListenAndServe()
}
