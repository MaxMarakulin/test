package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/go-co-op/gocron"
	"log"
	"test/database"
	"time"
)

var addr = flag.String("addr", "0.0.0.0:7100", "http service address")
var hostname = flag.String("hostname", "127.0.0.1", "database hostname")
var port = flag.String("port", "5432", "database port")
var databaseName = flag.String("database", "rates", "database name")
var username = flag.String("username", "postgres", "database username")
var password = flag.String("password", "", "database password")

func main() {
	// Иницилизируем базу данных
	err := database.Init(*hostname, *port, *databaseName, *username, *password)
	if err != nil {
		log.Println("database init error:", err)
	}
	defer func() {
		err = database.Close()
		if err != nil {
			log.Println("database close error:", err)
		}
	}()

	log.SetFlags(log.Lshortfile | log.LstdFlags)
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	// Настраиваем метод для получения курсов
	r.GET("/exchangeRate", getExchangeRate)

	// Так как курс обновляется раз в день, можно запрашивать один раз из разных источников и класть данные в базу
	s := gocron.NewScheduler(time.UTC)
	_, err = s.Cron("0 0 * * ?").Do(exchangeRate)
	if err != nil {
		log.Println("Error:", err)
	}
	s.StartAsync()
	log.Fatal(r.Run(*addr))

}
