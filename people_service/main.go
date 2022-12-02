package main

import (
	"fmt"
	"github.com/RyabovNick/databasecourse_2/golang/tasks/people_service/service/store"
)

func main() {
	conn := "postgresql://kamkin:kamkin@95.217.232.188:7777/kamkin"
	s := store.NewStore(conn)
	fmt.Println(s.GetPeopleByID("400"))
	fmt.Println(s.ListPeople())
}
