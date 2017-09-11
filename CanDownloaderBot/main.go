package main

import (
	_ "database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"gopkg.in/telegram-bot-api.v4"
	"log"
)

var db *sqlx.DB
var err error
var conf *viper.Viper
var bot *tgbotapi.BotAPI
var token string
var Ios, Android string
var CUser CandoUser

type CandoUser struct {
	TelegramID   int    `json:"telegram_id"`
	MobileNumber string `json:"mobile_number"`
}

func main() {

	GetConf()
	DBConnect()
	bot, err = tgbotapi.NewBotAPI(token)
	if err != nil {
		fmt.Println(err)
	}
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Panic(err)
	}
	for update := range updates {
		if update.Message != nil {
			CUser = CandoUser{}
			fmt.Println("Normal Message", update.Message.Text)
			CUser.TelegramID = update.Message.From.ID
			CUser.MobileNumber = GetMobile(update.Message.Chat.ID, CUser.TelegramID)
			fmt.Printf("tid:%d ___ mobile:%s", CUser.TelegramID, CUser.MobileNumber)
			switch update.Message.Text {

			case "/start":
				CUser.TelegramID = update.Message.From.ID
				CUser.MobileNumber = GetMobile(update.Message.Chat.ID, CUser.TelegramID)
				fmt.Println(CUser.MobileNumber)
				SendTextMessage(update.Message.Chat.ID, welcome+"\n", GetHomeKeys)
				fmt.Println("mobile: ", CUser.MobileNumber, "\n", "TelID: ", update.Message.From.ID, "\n ChatID", update.Message.Chat.ID)

			case iOSDownload:
				CUser.TelegramID = update.Message.From.ID
				CUser.MobileNumber = GetMobile(update.Message.Chat.ID, CUser.TelegramID)
				Download(update.Message.Chat.ID, Ios)

			case AndroidDownload:
				CUser.TelegramID = update.Message.From.ID
				CUser.MobileNumber = GetMobile(update.Message.Chat.ID, CUser.TelegramID)
				Download(update.Message.Chat.ID, Android)

			}
			/*	if update.Message.ReplyToMessage != nil {
					fmt.Println("Replay Message:", update.Message.ReplyToMessage.Text)
					switch update.Message.ReplyToMessage.Text {

				case EnterNumberForCheck_Buy:
					m:=Normalize(update.Message.Text)
					if MobileIsValid(m) {
						CUser.MobileNumber = update.Message.Text
						CUser.TelegramID = update.Message.From.ID
						NewUserQuery := fmt.Sprintf("UPDATE tgcodes SET (mobile='%s',LoginDate='%s' WHERE telegramid=%d", CUser.MobileNumber, TimeFormat(time.Now()), update.Message.From.ID)
						db.Exec(NewUserQuery)
						CheckMobileNumber(update.Message.Text, update.Message.Chat.ID, &flag)
						fmt.Println(flag, "FlagStatus")
						if flag == true {
							ShowInvoice(update.Message.Chat.ID, PID, CUser.MobileNumber) //CUser wa update.message.text
							//CUser.MobileNumber=update.Message.Text
						}
					} else {
						//SendTextMessage(update.Message.Chat.ID,MobileIsWrong,nil)
			*/ /*msg:=tgbotapi.NewMessage(update.Message.Chat.ID,MobileIsWrong)
			bot.Send(msg)*/ /*
						SendForceReply(update.Message.Chat.ID, MobileIsWrong)
					}
				case MobileIsWrong:
					flag := true

					if MobileIsValid(update.Message.Text) {
						CUser.MobileNumber = update.Message.Text
						CUser.TelegramID = update.Message.From.ID
						NewUserQuery := fmt.Sprintf("UPDATE tgcodes SET (mobile='%s',LoginDate='%s') WHERE telegramid=%d", CUser.MobileNumber, TimeFormat(time.Now()), update.Message.From.ID)
						db.Exec(NewUserQuery)
						CheckMobileNumber(update.Message.Text, update.Message.Chat.ID, &flag)
						//CUser.MobileNumber=update.Message.Text
						fmt.Println(flag, "FlagStatus")
						if flag == true {
							ShowInvoice(update.Message.Chat.ID, PID, update.Message.Text)
							//CUser.MobileNumber=update.Message.Text
						}
					} else {
						//SendTextMessage(update.Message.Chat.ID,MobileIsWrong,nil)
			*/ /*msg:=tgbotapi.NewMessage(update.Message.Chat.ID,MobileIsWrong)
			bot.Send(msg)*/ /*
						SendForceReply(update.Message.Chat.ID, MobileIsWrong)
					}
				case MobileIsWrong_CheckingCase:
					if MobileIsValid(update.Message.Text) {
						CUser.MobileNumber = update.Message.Text
						CUser.TelegramID = update.Message.From.ID
						NewUserQuery := fmt.Sprintf("UPDATE tgcodes SET (mobile='%s',LoginDate='%s') WHERE telegramid=%d", CUser.MobileNumber, TimeFormat(time.Now()), update.Message.From.ID)
						db.Exec(NewUserQuery)
						SendSecCode(update.Message.Chat.ID, update.Message.Text)
						//CUser.MobileNumber = update.Message.Text
						SendForceReply(update.Message.Chat.ID, SendCodeNotification)
						//CheckMobileNumber(update.Message.Text, update.Message.Chat.ID)
					} else {
						//SendTextMessage(update.Message.Chat.ID,MobileIsWrong,nil)
			*/ /*msg:=tgbotapi.NewMessage(update.Message.Chat.ID,MobileIsWrong)
			bot.Send(msg)*/ /*
						SendForceReply(update.Message.Chat.ID, MobileIsWrong_CheckingCase)
					}
				case EnterNumberForCheck:
					if MobileIsValid(update.Message.Text) {
						CUser.MobileNumber = update.Message.Text
						CUser.TelegramID = update.Message.From.ID
						NewUserQuery := fmt.Sprintf("UPDATE tgcodes SET (mobile='%s',LoginDate='%s') WHERE telegramid=%d", CUser.MobileNumber, TimeFormat(time.Now()), update.Message.From.ID)
						db.Exec(NewUserQuery)
						SendSecCode(update.Message.Chat.ID, update.Message.Text)
						//	CUser.MobileNumber = update.Message.Text
						SendForceReply(update.Message.Chat.ID, SendCodeNotification)
						//CheckMobileNumber(update.Message.Text, update.Message.Chat.ID)
					} else {
						//SendTextMessage(update.Message.Chat.ID,MobileIsWrong,nil)
			*/ /*msg:=tgbotapi.NewMessage(update.Message.Chat.ID,MobileIsWrong)
			bot.Send(msg)*/ /*
						SendForceReply(update.Message.Chat.ID, MobileIsWrong_CheckingCase)
					}
				case SendCodeNotification:
					SecCodeReview(update.Message.Chat.ID, update.Message.Text)

				case WrongCode:
					SecCodeReview(update.Message.Chat.ID, update.Message.Text)
			*/ /*SendSecCode(update.Message.Chat.ID, update.Message.Text)
			SendForceReply(update.Message.Chat.ID, SendCodeNotification)*/ /*
				}
			}*/
		}
	}
}
