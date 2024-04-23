package httpclient

import (
	"cars_catalog/internal/entity/cars"
	"cars_catalog/internal/repository"
	"net/http"
	"time"

	"github.com/brianvoe/gofakeit/v7"
)

type httpClientRepo struct {
	client *http.Client
}

func NewHttpClientRepo() repository.HttpCLient {
	return &httpClientRepo{
		client: &http.Client{
			Transport: http.DefaultTransport,
			Timeout:   time.Second * 10,
		},
	}
}
func (r *httpClientRepo) GetCarInfo(regNum string) (cars.AddCarParam, error) {


	// logic here
	// url := fmt.Sprintf("example.com/%s", regNum)

	// req, err := http.NewRequest("GET", url, nil)
	// if err != nil {
	// 	return cars.AddCarParam{}, err
	// }
	// _, err = r.client.Do(req)
	// if err != nil {
	// 	return cars.AddCarParam{}, err
	// }

	var car cars.AddCarParam
	car.RegNum.String = regNum
	car.Mark.String = gofakeit.Word()
	car.Model.String = gofakeit.Word()
	car.Year.Int64 = gofakeit.Int64()
	car.PeopleID.Int64 = gofakeit.Int64()

	return car, nil
}
