package main

import (
	"github.com/spf13/viper"
	"math/rand"
	"strconv"
	"time"
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
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
	vcf = conf.GetString("file.VCFile")
	//////////// Data Base Definitions ////////////
	DBEngine = "postgres"
	DBName = conf.GetString("DB.DBName")
	DBHost = conf.GetString("DB.DBHost")
	DBUser = conf.GetString("DB.DBUser")
	DBPassword = conf.GetString("DB.DBPassword")
}
func DBConnect() {
	var schema = `
  	CREATE TABLE IF NOT EXISTS tgcodes (
    id SERIAL PRIMARY KEY NOT NULL,
 	telegramid int  NOT NULL,
 	mobile text NOT NULL,
 	LoginDate text
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
func TodayWitZeroStr() string {
	year, m, d := time.Now().Date()
	m_ := strconv.Itoa(int(m))
	if len(m_) == 1 {
		m_ = "0" + m_
	}
	d_ := strconv.Itoa(d)

	if len(d_) == 1 {
		d_ = "0" + d_
	}
	strDate := strconv.Itoa(year) + "/" + m_ + "/" + d_
	return strDate

}
