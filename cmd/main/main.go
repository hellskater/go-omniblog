package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hellskater/omniblog/pkg/db"
	"github.com/hellskater/omniblog/pkg/router"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	db := db.Connect()
	db.Migrate()
	defer db.Close()

	r := gin.Default()
	router.Initialize(r, db.DB)
	log.Fatal(http.ListenAndServe(":4000", r))
}
