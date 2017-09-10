package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"log"
	"math/rand"
	"strconv"
	"time"
)

var DBEngine, DBName, DBUser, DBHost, DBPassword string

func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}
func GetConf() {
	conf = viper.New()
	conf.SetConfigName("conf")
	conf.AddConfigPath(".")
	conf.SetConfigType("yml")
	err := conf.ReadInConfig()
	if err != nil {
		panic("could not configure app")
	}
	token = conf.GetString("Public.Token")
	ios = conf.GetString("File.IOSFile")
	android = conf.GetString("File.AndroidFile")
	//////////// Data Base Definitions ////////////
	DBEngine = "postgres"
	DBName = conf.GetString("DB.DBName")
	DBHost = conf.GetString("DB.DBHost")
	DBUser = conf.GetString("DB.DBUser")
	DBPassword = conf.GetString("DB.DBPassword")
}
func DBConnect() {
	var schema = `
  	CREATE TABLE IF NOT EXISTS candousers (
    id SERIAL PRIMARY KEY NOT NULL,
 	telegramid int  NOT NULL,
 	mobile text NOT NULL,
 	LoginDate text,
	platform text
	);`
	var err error
	query := fmt.Sprintf("host=%s  user=%s password=%s dbname=%s sslmode=disable", DBHost, DBUser, DBPassword, DBName)
	db, err = sqlx.Connect("postgres", query)
	if err != nil {
		fmt.Println(err.Error())
		log.Panic(err)
	}
	db.MustExec(schema)
}
