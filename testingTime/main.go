package main

import (
	"fmt"
	"time"
)

//func main() {
//
//	layout := "2006-01-02T15:04:05Z"
//	//layout := "2021-01-02T00:00:00Z"
//	//str := "2021-09-10T11:45:05Z"
//	str := "2021-09-10T00:00:00Z"
//	  //str := " 2021-09-10T00:00:00"
//
//	t, err := time.Parse(layout, str) //ну вот ты же меняешь формат
//	newTime := t.Format("02-01-06")
//
//	fmt.Println(t.Month())
//	fmt.Println(t.Year())
//	fmt.Println(t.Day())
//	fmt.Println(newTime)
//
//	if err != nil {
//	fmt.Println(err)
//	}
//	fmt.Println(t)
//}


func main() {

	layout := "2006-01-02T15:04:05.000Z"
	// добавим 10 августа
	str := "2021-08-10T11:45:26.371Z"
	// 3 августа он должен
	t, _ := time.Parse(layout, str) //ну вот ты же меняешь формат
	newTime := t.Format("02-01-06")
	fmt.Println(newTime)
	fmt.Println(t.Month())
	today := time.Now().AddDate(0,1,5)
	fmt.Println(today.Month(), "тут должен быть сентябрь ")
	fmt.Println(t.Day(), "сегодня")
	fmt.Println(t.Month(), "сегодня")

	fmt.Println(today.Day(), "тут должно показать 10 число")

	// проверка должна быть не только же по дням, но и по


	if today.Day() == t.Day()  && today.Month() == t.Month() {
		fmt.Println("У этого чела днюха")
	}



	tomorrow := today.Add(24 * time.Hour)
	sameday := tomorrow.Add(-24 * time.Hour)

	if today != tomorrow {
		fmt.Println("today is not tomorrow")
	}

	if sameday == today {
		fmt.Println("sameday is today")
	}

	// using Equal function
	if today.Equal(sameday) {
		fmt.Println("today is sameday")
	}

}
