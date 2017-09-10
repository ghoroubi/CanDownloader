package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"log"
	"strconv"
	"time"
	"os"
	"gopkg.in/telegram-bot-api.v4"
	"strings"
)

var (
	DBHost       string
	DBUser       string
	DBName		 string
	DBPassword	 string
)

func GetConf() {
	LoggerInit()
	conf = viper.New()
	conf.SetConfigName("conf")
	conf.AddConfigPath(".")
	conf.SetConfigType("yml")
	err := conf.ReadInConfig()
	if err != nil {
		log.Println("could not configure app")
		panic("could not configure app")
	}
	token = conf.GetString("Public.Token")
	ios = conf.GetString("File.IOSFile")
	android = conf.GetString("File.AndroidFile")
	//////////// Data Base Definitions ////////////
	DBName = conf.GetString("DB.DBName")
	DBHost = conf.GetString("DB.DBHost")
	DBUser = conf.GetString("DB.DBUser")
	DBPassword = conf.GetString("DB.DBPassword")
}
func LoggerInit(){
	var err error
	LogFile, err = os.OpenFile("logs.txt", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0777)
	log.New(LogFile,time.Now().String(),0)
	log.SetOutput(LogFile)

	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
}
func DBConnect() {
	var schema = `
  	CREATE TABLE IF NOT EXISTS cusers (
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
func Download(chatId int64, text string) {

	msg2Send := tgbotapi.NewMessage(chatId, text)
	bot.Send(msg2Send)

	file := tgbotapi.NewDocumentUpload(chatId, text)
	bot.Send(file)
}
func Normalize(strNum string) string{
	var x string
	for i := 0; i < len(strNum);i++  {
		temp:=charCodeAt(strNum,int(i))
		if temp >=1776 && temp<= 1785 {
			x+=string(temp-1728)

		}else{
			x+=string(temp)

		}

	}
	u:=strings.Replace(x, "\x00", "", -1)
	log.Println("User Verification Code Normalized Succesfully.Code:",u)
	return u

}
func charCodeAt(s string, n int) rune {
	i := 0
	for _, r := range s {
		if i == n {
			return r
		}
		i++
	}
	return 0
}
func MobileIsValid(mobile string) bool {
	var IsMobile bool
	lenFlag := len(mobile) == 11
	fmt.Println("len:", len(mobile))
	firstDigitFlag := string(mobile[0]) == "0"
	fmt.Println("firstDigit:", string(mobile[0]))
	secDigitFlag := string(mobile[1]) == "9"
	fmt.Println("SecondDigit", string(mobile[1]))
	IsMobile = lenFlag && firstDigitFlag && secDigitFlag
	//IsMobile=firstDigitFlag && secDigitFlag
	//fdigit:=
	return IsMobile
}

type fn func() tgbotapi.ReplyKeyboardMarkup

//If User Choose to check his/her active plan
func TimeFormat(t time.Time) string {
	var strTime string
	year := t.Year()
	month := t.Month()
	day := t.Day()
	hour := t.Hour()
	min := t.Minute()
	strTime = fmt.Sprintf("%4d-%2d-%2d __ %2d:%2d", year, month, day, hour, min)

	return strTime
}

func SendError(chatId int64, keyboard fn) {
	msg := tgbotapi.NewMessage(chatId, "")
	msg.ReplyMarkup = keyboard()
	msg.Text = "SystemError"
	bot.Send(msg)
}
func SendTextMessage(chatId int64, text string, keys fn) {
	msg := tgbotapi.NewMessage(chatId, text)
	msg.ReplyMarkup = keys()
	bot.Send(msg)
}

func SendForceReply(chatId int64, text string) {
	fmt.Println("ForceReply: ", text)
	msg := tgbotapi.NewMessage(chatId, text)
	msg.ReplyMarkup = tgbotapi.ForceReply{ForceReply: true}
	bot.Send(msg)
}

func GetMobile(chatId int64, telegramId int) string {
	msg := tgbotapi.NewMessage(chatId, "")
	msg.ReplyMarkup = GetHomeKeys()
	// get mobile
	var strmobile string
	//var id int
	err := db.Get(&strmobile, "select mobile from tgCodes where telegramId="+strconv.Itoa(telegramId))
	fmt.Println(strmobile)
	if err != nil {
		//SendError(chatId, GetHomeKeys)
		strmobile = ""
		return ""
	}
	return strmobile
}





