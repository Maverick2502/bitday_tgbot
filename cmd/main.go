package main

import (
	"database/sql"
	"fmt"
	"github.com/AzizRahimov/bitday_tgbot/cmd/app"
	"github.com/AzizRahimov/bitday_tgbot/pkg/service"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/yanzay/tbot"
	"log"
	"os"
	"time"
)


func sendingMessage(bot *tbot.Server) {
	//msg := &model.Message{
	//	ChatID: 484843300,
	//	Type:   model.MessageText,
	//	Data:   str,
	//}
	input := "08-11" // в таком формате я должен хранить дату?

	layot := "01-02"

	t1, _ := time.Parse(layot, input)
	t2, _ := time.Parse("2006-01-02", "2016-01-01 18:19:20.0")
	ttt := time.Now().AddDate(0,0,45)
	fmt.Println(t1.Format("2006-01-02"))

	fmt.Println("сегодня", ttt)
	fmt.Println(t1.Day())
	fmt.Println(ttt)

	for {
		if t1.Equal(ttt) {
			fmt.Println("все верно")

			if t1.Year() == t2.Year() {
				fmt.Println(true)

			}

			bot.Send(484843300, "через неделю у этого чела днюха")
			time.Sleep(time.Second * 3)
		}

		//	err := bot.SendMessage(msg)
		//	if err != nil {
		//		fmt.Println(err)
		//		return
		//	}
		//	time.Sleep(1 * time.Second)
		//}

	}
}


func main() {


	//url := "https://some.url"
	//addr := "0.0.0.0:8013"

	//bot, err := tbot.NewServer("1903232530:AAFpzgWZ8j9hXFOs1u9fg2ePVCKgIrqgofo")
	bot, err := tbot.NewServer("1912918677:AAHrMEawvWinS9Exi8OCZIKmIVpeet2nIiM")
	//bot, err := tbot.NewServer("1912918677:AAHrMEawvWinS9Exi8OCZIKmIVpeet2nIiM",tbot.WithWebhook(url, addr) )
	//tbot.NewServer("dasdas", We)
	if err != nil {
		log.Fatal(err)
	}
	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "9999"
	}

	log.Println(port)

	log.Println("The Bot is starting:" )
	//dsn := "postgres://app:pass@192.168.99.100:5424/db"

	dsn := "postgres://omvzgmcaylrqrq:0be0c5fa91fb6ff63b5d1c9872e0f388bcf233964d87d78c9fcd863fa51e9921@ec2-44-196-132-15.compute-1.amazonaws.com:5432/de1e4st8q72m0b"

	db, err := sql.Open("pgx", dsn)

	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	defer func ()  {
		if err := db.Close(); err != nil{
			log.Print(err)
		}

	}()
	tgBot := service.NewService(db)
	server := app.NewServer(bot, tgBot)
	go server.NearBDHandler() // и тут дай бота и все
	server.Init() // ВСЕГДА ВЫЗЫВАЙ ОБРАБОТЧИКОВ:

	bot.ListenAndServe()








}


