package external

import (
	"cars_catalog/internal/usecase"
	"fmt"
	"log"
	"net/http"
)

type Server struct {
	mux           *http.ServeMux
	peopleUsecase *usecase.PeopleUsecase
	carUsecase    *usecase.CarsUsecase
}

func New(peopleUsecase *usecase.PeopleUsecase, carsUsecase *usecase.CarsUsecase) *Server {
	return &Server{
		mux:           http.NewServeMux(),
		peopleUsecase: peopleUsecase,
		carUsecase:    carsUsecase,
	}
}

func (s *Server) Run() {

	s.mux.HandleFunc("/swagger/", SwaggerHandler)

	s.mux.HandleFunc("/people", s.GetPeopleListHandler)
	s.mux.HandleFunc("/people/delete", s.DeletePeopleHandler)
	s.mux.HandleFunc("/people/update", s.UpdatePeopleHandler)

	s.mux.HandleFunc("/cars/delete", s.DeleteCarHandler)
	s.mux.HandleFunc("/cars/update", s.UpdateCarHandler)
	s.mux.HandleFunc("/cars/create", s.AddCarHandler)
	s.mux.HandleFunc("/cars", s.GetCarsListHandler)

	fmt.Println("сервер успешно запущен на порту :9000")
	if err := http.ListenAndServe(":9000", s.mux); err != nil {
		log.Fatalln("не удалось начать прослушивание, ошибка:", err)
	}
}
