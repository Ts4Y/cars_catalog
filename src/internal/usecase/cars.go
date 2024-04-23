package usecase

import (
	"cars_catalog/internal/entity/cars"
	"cars_catalog/internal/entity/global"
	"cars_catalog/internal/repository"
	"cars_catalog/internal/repository/postgresql"
	"cars_catalog/internal/repository/transaction"
	"database/sql"
	"fmt"
	"log/slog"
	"sort"
)

type CarsUsecase struct {
	log            *slog.Logger
	sm             transaction.SessionManager
	httpClientRepo repository.HttpCLient
	carsRepo       repository.Car
}

func NewCarsUsecase(log *slog.Logger, sm transaction.SessionManager, carsRepo repository.Car, httpClientRepo repository.HttpCLient) *CarsUsecase {
	return &CarsUsecase{
		log:            log,
		sm:             sm,
		carsRepo:       carsRepo,
		httpClientRepo: httpClientRepo,
	}
}

func (u *CarsUsecase) GetCarList(sortBy string) (carList []cars.Car, err error) {
	tx := u.sm.CreateSession()
	if err = tx.Start(); err != nil {
		u.log.Error(fmt.Sprintf("не удалось открыть транзакцию, ошибка: %s", err))
		err = global.ErrInternalError
		return
	}

	defer tx.Rollback()

	carList, err = u.carsRepo.GetCarsList(postgresql.SqlxTx(tx))
	switch err {
	case nil:
	case sql.ErrNoRows:
		u.log.Debug("нет данных")
		return
	default:
		u.log.Error(fmt.Sprintf("не удалось открыть транзакцию, ошибка: %s", err))
		err = global.ErrInternalError
		return
	}

	switch sortBy {
	case "mark":
		u.sortCarListByMark(carList)
	case "model":
		u.sortCarListByModel(carList)
	case "reg_num":
		u.sortCarListByRegNum(carList)
	default:
		u.sortCarListByYear(carList)
	}

	u.log.Info("данные успешно получены")
	return
}

func (u *CarsUsecase) DeleteCar(carID int) (err error) {
	tx := u.sm.CreateSession()
	if err = tx.Start(); err != nil {
		u.log.Error(fmt.Sprintf("не удалось открыть транзакцию, ошибка: %s", err))
		err = global.ErrInternalError
		return
	}

	defer tx.Rollback()

	if err = u.carsRepo.DeleteCarByID(postgresql.SqlxTx(tx), carID); err != nil {
		u.log.Error(fmt.Sprintf("не удалось удалить автомобиль, ошибка: %s", err))
		err = global.ErrInternalError
		return
	}

	if err = tx.Commit(); err != nil {
		u.log.Error(fmt.Sprintf("не удалось закрыть транзакцию, ошибка: %s", err))
		err = global.ErrInternalError
		return
	}

	u.log.Info(fmt.Sprintf("автомоболиь с id = %d успешно удален", carID))
	return
}

func (u *CarsUsecase) UpdateCar(p cars.UpdateCarParam) (err error) {
	if p.CarID <= 0 {
		u.log.Error("нет автомобиля с нулевым или отрицательным id")
		err = global.ErrIncorectParams
		return
	}

	tx := u.sm.CreateSession()
	if err = tx.Start(); err != nil {
		u.log.Error(fmt.Sprintf("не удалось открыть транзакцию, ошибка: %s", err))
		err = global.ErrInternalError
		return
	}

	defer tx.Rollback()

	if err = u.carsRepo.UpdateCar(postgresql.SqlxTx(tx), p); err != nil {
		u.log.Error(fmt.Sprintf("не удалось обновить данные автомобиля, ошибка: %s", err))
		err = global.ErrInternalError
		return
	}

	if err = tx.Commit(); err != nil {
		u.log.Error(fmt.Sprintf("не удалось закрыть транзакцию, ошибка: %s", err))
		err = global.ErrInternalError
		return
	}

	u.log.Info(fmt.Sprintf("автомоболиь с id = %d успешно обновлен", p.CarID))
	return
}

func (u *CarsUsecase) AddCar(regNumList []string) (err error) {
	tx := u.sm.CreateSession()
	if err = tx.Start(); err != nil {
		u.log.Error(fmt.Sprintf("не удалось открыть транзакцию, ошибка: %s", err))
		err = global.ErrInternalError
		return
	}

	defer tx.Rollback()

	var carList []cars.AddCarParam

	for _, regNum := range regNumList {
		car, err := u.httpClientRepo.GetCarInfo(regNum)
		if err != nil {
			u.log.Error(fmt.Sprintf("не удалось добавить автомобиль, ошибка: %s", err))
			err = global.ErrInternalError
			break
		}

		carList = append(carList, car)
	}

	for _, car := range carList {
		if err = u.carsRepo.AddCar(postgresql.SqlxTx(tx), car); err != nil {
			u.log.Error(fmt.Sprintf("не удалось добавить автомобиль, ошибка: %s", err))
			err = global.ErrInternalError
			break
		}
	}

	if err = tx.Commit(); err != nil {
		u.log.Error(fmt.Sprintf("не удалось закрыть транзакцию, ошибка: %s", err))
		err = global.ErrInternalError
		return
	}

	u.log.Info("автомобиль успешно добавлен")
	return

}

func (u *CarsUsecase) sortCarListByMark(carList []cars.Car) {
	sort.Slice(carList, func(i, j int) bool {
		return carList[i].Mark < carList[j].Mark
	})
}

func (u *CarsUsecase) sortCarListByModel(carList []cars.Car) {
	sort.Slice(carList, func(i, j int) bool {
		return carList[i].Model < carList[j].Model
	})
}

func (u *CarsUsecase) sortCarListByRegNum(carList []cars.Car) {
	sort.Slice(carList, func(i, j int) bool {
		return carList[i].RegNum < carList[j].RegNum
	})
}

func (u *CarsUsecase) sortCarListByYear(carList []cars.Car) {
	sort.Slice(carList, func(i, j int) bool {
		return carList[i].Year < carList[j].Year
	})
}
