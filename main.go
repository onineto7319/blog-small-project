package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/blog-small-project/global"
	"github.com/blog-small-project/internal/router"
	"github.com/blog-small-project/pkg/database"
	"github.com/blog-small-project/pkg/setting"
)

func init() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}

	err = setupDBEngine()
	if err != nil {
		log.Fatalf("init.setupDBEngine err: %v", err)
	}

	err = createMysqlTables()
	if err != nil {
		log.Fatalf("init.createMysqlTables err: %v", err)
	}
}

// @title 部落格系統
// @version v1.0
// @termsOfService https://github.com/blog-small-project
func main() {
	if herokuPort := os.Getenv("PORT"); herokuPort != "" {
		global.ServerSetting.HttpPort = herokuPort
	}

	router := router.New()

	s := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20, //1MB
	}

	log.Fatal(s.ListenAndServe())
}

func setupSetting() error {
	s, err := setting.NewSetting()

	if err != nil {
		return err
	}

	err = s.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		fmt.Println("Server err:", err)
		return err
	}

	err = s.ReadSection("DatabaseMysql", &global.DatabaseMysqlSetting)
	if err != nil {
		fmt.Println("DatabaseMysql err:", err)
		return err
	}

	err = s.ReadSection("JWT", &global.JWTSetting)
	if err != nil {
		fmt.Println("JWT err:", err)
		return err
	}

	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second

	return nil
}

func setupDBEngine() error {
	var err error

	global.MysqlEngine, err = database.NewMySQLDBEngine(global.DatabaseMysqlSetting)

	if err != nil {
		return err
	}

	return nil
}

func createMysqlTables() error {
	return database.CreateMysqlTables(global.MysqlEngine)
}
