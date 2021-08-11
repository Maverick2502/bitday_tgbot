package main

import (
	"fmt"

)


//func main() {
//	// Do jobs without params
//	gocron.Every(1).Second().Do(task)
//	/*
//	gocron.Every(2).Seconds().Do(task)
//	gocron.Every(1).Minute().Do(task)
//	gocron.Every(2).Minutes().Do(task)
//	gocron.Every(1).Hour().Do(task)
//	gocron.Every(2).Hours().Do(task)
//	gocron.Every(1).Day().Do(task)
//	gocron.Every(2).Days().Do(task)
//	gocron.Every(1).Week().Do(task)
//	gocron.Every(2).Weeks().Do(task)
//
//	// Do jobs with params
//	gocron.Every(1).Second().Do(taskWithParams, 1, "hello")
//
//	// Do jobs on specific weekday
//	gocron.Every(1).Monday().Do(task)
//	gocron.Every(1).Thursday().Do(task)
//
//	// Do a job at a specific time - 'hour:min:sec' - seconds optional
//	gocron.Every(1).Day().At("10:30").Do(task)
//	gocron.Every(1).Monday().At("18:30").Do(task)
//	gocron.Every(1).Tuesday().At("18:30:59").Do(task)
//
//	// Begin job immediately upon start
//	gocron.Every(1).Hour().From(gocron.NextTick()).Do(task)
//
//	// Begin job at a specific date/time
//	t := time.Date(2019, time.November, 10, 15, 0, 0, 0, time.Local)
//	gocron.Every(1).Hour().From(&t).Do(task)
//
//	// NextRun gets the next running time
//	_, time := gocron.NextRun()
//	fmt.Println(time)
//
//	// Remove a specific job
//	gocron.Remove(task)
//
//	// Clear all scheduled jobs
//	gocron.Clear()
//*/
//	// Start all the pending jobs
//	<- gocron.Start()
//
//}

type Something struct {
	ID int
	Name string
}

func main() {

	name := ConvertName("Махина")
	fmt.Println(name, "Это конвертируемое имя")

	//
	//name := "Азиз"  // мы к каждой имени говорим с какой буквой заканчивается
	////department := "гласный"
	//l := len(name)
	//fmt.Println(len(name))
	//fmt.Println(name[:l-2] + "ы")
	//endName := name[len(name)-2:]
	//fmt.Println(endName, "выфвфы")
	//
	//
	//// Add SWITCH
	//// и
	//
	//switch  endName{
	//case "а":
	//	fmt.Println(name[:l-2] + "ы")
	//	name = name[:l-2] + "ы"
	//default:
	//	// тут просто м
	//		name = name + "a"
	//	}
	//
	//
	//	fmt.Println(name)

	//if department == "гласный" {
	//fmt.Println(name[:l-2] + "ы")
	//
	//}

	//fmt.Println(strings.Replace(name, name, "a", 14))
	//strings.ReplaceAll(name)


	//s := Something{
	//	ID: 1,
	//	Name: "Aziz",
	//}
	//
	//fmt.Println(s)
	//fmt.Printf("%v\n", s)


	//t := time.Now().AddDate(0,0,0)
	//
	//
	//fmt.Println(t.Format("02-01-06"))

	//input := "2021-08-11" // в таком формате я должен хранить дату?

	// а что если мы просто будем хранить дату и месяц без года? получится думаешьь?
	//input := "2021-08-11" // в таком формате я должен хранить дату?
	//
	//layot := "2006-01-02"
	//
	//
	//t1, _ := time.Parse(layot, input)
	//t2, _ := time.Parse("2006-01-02", "2016-01-01 18:19:20.0")
	//ttt := time.Now().Day() + 7
	//fmt.Println(t1.Format("2006-01-02"))
	//
	//
	//fmt.Println("сегодня", ttt)
	//fmt.Println(t1.Day())
	//fmt.Println(ttt)
	//if t1.Day() == ttt{
	//	fmt.Println("все верно")
	//}
	//
	//if  t1.Year() == t2.Year(){
	//	fmt.Println(true)
	//
	//}
	//

}

func ConvertName(name string) string  {
	l := len(name)
	//fmt.Println(len(name))
	//fmt.Println(name[:l-2] + "ы")
	endName := name[len(name)-2:]
	//fmt.Println(endName, "выфвфы")
	// Add SWITCH
	// и
	switch  endName{
	case "а":
	//	fmt.Println(name[:l-2] + "ы")
		name = name[:l-2] + "ы"
	default:
		// тут просто м
		name = name + "a"
	}


	fmt.Println(name)

	return name

}