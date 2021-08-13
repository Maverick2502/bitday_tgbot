package app

import (
	"fmt"
	"github.com/AzizRahimov/bitday_tgbot/pkg/service"
	"github.com/yanzay/tbot"
	"log"
	"strconv"
	"time"
)

type Server struct {
	bot   *tbot.Server
	tgBot *service.Service

}


func NewServer(bot *tbot.Server, tgBot *service.Service) *Server {

	return &Server{bot: bot, tgBot: tgBot}
}

func ConvertName(name string) string {
	l := len(name)

	endName := name[len(name)-2:]
	switch endName {
	case "а":
		//	fmt.Println(name[:l-2] + "ы")
		name = name[:l-2] + "ы"
	default:
		name = name + "a"
	}

	return name
}

func (s *Server) timerHandler(m *tbot.Message) {

	// m.Vars contains all variables, parsed during routing
	secondsStr := m.Vars["seconds"]
	// Convert string variable to integer seconds value
	seconds, err := strconv.Atoi(secondsStr)
	if err != nil {
		m.Reply("Invalid number of seconds")
		return
	}
	cId := m.ChatID

	m.Replyf("Timer for %d seconds started, chatID is %d", seconds, cId)
	time.Sleep(time.Duration(seconds) * time.Second)
	m.Reply("Time out!")
}

func (s *Server) listHandler(m *tbot.Message) {
	// вызови логику
	list, err := s.tgBot.List()

	if err != nil {
		log.Println("ошибка при получении данных")
		return
	}

	//m.Reply("Список дни рождений")

	var result string

	result +=  fmt.Sprintf(" %s\n", "Список дни рождений")
	for _, value := range list {

		layout := "2006-01-02T15:04:05Z"
		str := value.BDate

		t, err := time.Parse(layout, str) //ну вот ты же меняешь формат
		//newTime := t.Format("02-01-06")
		//fmt.Println(newTime)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(t)

		result += fmt.Sprintf("%s %d %s\n", value.Name, t.Day(), t.Month())

	}
	m.Reply(result)
}

func (s *Server) NearBDHandler() {
	list, err := s.tgBot.List()
	if err != nil {
		log.Println("ошибка при получении данных")
		return
	}
	for {
		for _, value := range list {
			layout := "2006-01-02T15:04:05Z"
			str := value.BDate
			t, err := time.Parse(layout, str) //ну вот ты же меняешь формат
			if err != nil {
				fmt.Println(err)
			}
			//	newTime := t.Format("02-01-06")
			//	fmt.Println(t.Month())
			//	fmt.Println(t.Year())
			//	fmt.Println(t.Day())
			//	fmt.Println(newTime)
			//
			//	fmt.Println(t)
			today := time.Now().AddDate(0, 0, +6)
			fmt.Println(t.Day(), "сегодня")
			newName := ConvertName(value.Name)

			day := fmt.Sprintf("%d", today.Day())

			if today.Day() == t.Day() && today.Month() == t.Month() {

				//s.bot.Send(-1001267478076, "Через неделю " +  day + "числа у " + newName +" "+"день рождения! Удалите его из группы и начинаем готовиться!")
				s.bot.Send(484843300, "Через неделю " +  day + "числа у " + newName +" "+"день рождения! Удалите его из группы и начинаем готовиться!")

			}

		}
		time.Sleep(time.Hour * 24)
	}
}

func (s *Server) nearBirthDayHandler(m *tbot.Message) {
	// вызови логику
	nearBd, err := s.tgBot.NearBirthday()

	if err != nil {
		log.Println("ошибка ближайщии даты")
		return
	}

	layout := "2006-01-02T15:04:05Z"
	str := nearBd.BDate

	t, err := time.Parse(layout, str)
	if err != nil {
		fmt.Println(err)
	}

	//	newTime := t.Format("02-01-06")
	name := ConvertName(nearBd.Name)
	m.Replyf("Ближайший день рождения у %s %d %s ", name, t.Day(), t.Month())

}

func (s *Server) Init() {
	//	s.bot.Handle("/answer", "42")
	s.bot.HandleFunc("/timer {seconds}", s.timerHandler)
	s.bot.HandleFunc("/list", s.listHandler)
	s.bot.HandleFunc("/near", s.nearBirthDayHandler)
	s.bot.HandleFunc("/check", s.checkHandler)
}

func (s *Server) checkHandler(m *tbot.Message)  {
	now := time.Now()

	m.Replyf("Сегодня: %s", now)

}
