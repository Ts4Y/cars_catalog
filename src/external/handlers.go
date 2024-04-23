package external

import (
	"cars_catalog/internal/entity/cars"
	"cars_catalog/internal/entity/people"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"


)

// @Summary Get people list
// @Tags people
// @Accept json
// @Produce json

func SwaggerHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./docs/index.html")
}


func (s *Server) GetPeopleListHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "incorect method", http.StatusMethodNotAllowed)
		return
	}

	sortBy := r.URL.Query().Get("sort_by")

	peopleList, err := s.peopleUsecase.GetPeopleList(sortBy)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err = json.NewEncoder(w).Encode(peopleList); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
}

func (s *Server) DeletePeopleHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "incorect method", http.StatusMethodNotAllowed)
		return
	}

	peopleID, err := strconv.Atoi(r.URL.Query().Get("people_id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := s.peopleUsecase.DeletePeople(peopleID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "человек успешно удален")
}

func (s *Server) UpdatePeopleHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "incorect method", http.StatusMethodNotAllowed)
		return
	}

	var updatePeopleParams people.UpdatePeopleParam

	if err := json.NewDecoder(r.Body).Decode(&updatePeopleParams); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := s.peopleUsecase.UpdatePeople(updatePeopleParams); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "данные человека были успешно обновлены")
}

func (s *Server) GetCarsListHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "incorect method", http.StatusMethodNotAllowed)
		return
	}

	sortBy := r.URL.Query().Get("sort_by")
	carsList, err := s.carUsecase.GetCarList(sortBy)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err = json.NewEncoder(w).Encode(carsList); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
}

func (s *Server) DeleteCarHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "incorect method", http.StatusMethodNotAllowed)
		return
	}

	carID, err := strconv.Atoi(r.URL.Query().Get("car_id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := s.carUsecase.DeleteCar(carID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "автомобиль успешно удален")
}

func (s *Server) UpdateCarHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "incorect method", http.StatusMethodNotAllowed)
		return
	}

	var updateCarParams cars.UpdateCarParam

	if err := json.NewDecoder(r.Body).Decode(&updateCarParams); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := s.carUsecase.UpdateCar(updateCarParams); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "данные автомобиля были успешно обновлены")
}

func (s *Server) AddCarHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "incorect method", http.StatusMethodNotAllowed)
		return
	}

	var addCarParams struct {
		regNums []string
	}

	if err := json.NewDecoder(r.Body).Decode(&addCarParams); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := s.carUsecase.AddCar(addCarParams.regNums); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "автомобиль успешно добавлен")
}
