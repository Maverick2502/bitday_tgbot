package service

import (
	"database/sql"
	"fmt"
	"log"

	//"time"
)


type Service struct {
	db *sql.DB
}



// мы разве не использовали pgx???

// такс, тут я забыл в аргументах дать бд
func NewService(db *sql.DB) *Service {
	return &Service{db: db}
}



type People struct {
	ID int `json:"id"`
	Name string `json:"name"`
	BDate string `json:"b_date"`	//Date2 time.Time

}



//db.Automigrate(People{})
type Ppp struct {
	Name string
}

// это логика
func (s *Service) List() ([]People, error) {
	fmt.Println("вызвалась логика")

	people := make([]People,0)
	query := fmt.Sprintf("SELECT name, bdate FROM %s order by bdate", "people")
	rows, err := s.db.Query(query)


	if err != nil {
		return nil, err
	}
	for rows.Next(){
		p := People{}

		err := rows.Scan(&p.Name, &p.BDate)
		if err != nil {
			return nil, err
		}



		log.Println("новые данные людей", p)




		people =  append(people, p)

	}

	return people, err

}

func (s *Service) NearBirthday() (People, error)  {
	var p People
	query := fmt.Sprintf("select name,bdate from %s where bdate > current_timestamp  order by bdate  LIMIT 1", "people")

	err := s.db.QueryRow(query).Scan(&p.Name, &p.BDate)
	fmt.Println("ppp", p)
	if err != nil {
		log.Println(err)
		return p, err
	}


	return p, nil
}