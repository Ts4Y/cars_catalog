package usecase

import (
	"cars_catalog/internal/entity/global"
	"cars_catalog/internal/entity/people"
	"cars_catalog/internal/repository"
	"cars_catalog/internal/repository/postgresql"
	"cars_catalog/internal/repository/transaction"
	"database/sql"
	"fmt"
	"log/slog"
	"sort"
)

type PeopleUsecase struct {
	log        *slog.Logger
	peopleRepo repository.People
	sm         transaction.SessionManager
}

func NewPeopleUsecase(log *slog.Logger, sm transaction.SessionManager, peopleRepo repository.People) *PeopleUsecase {
	return &PeopleUsecase{
		log:        log,
		sm:         sm,
		peopleRepo: peopleRepo,
	}
}

func (u *PeopleUsecase) GetPeopleList(sortBy string) (peopleList []people.People, err error) {
	tx := u.sm.CreateSession()
	if err = tx.Start(); err != nil {
		u.log.Error(fmt.Sprintf("не удалось открыть транзакцию, ошибка: %s", err))
		return nil, err
	}

	defer tx.Rollback()

	peopleList, err = u.peopleRepo.GetPeopleList(postgresql.SqlxTx(tx))
	switch err {
	case nil:
	case sql.ErrNoRows:
		u.log.Debug("нет данных")
		return
	default:
		u.log.Error(fmt.Sprintf("не удалось получить список людей, ошибка: %s", err))
		err = global.ErrInternalError
		return
	}

	switch sortBy {
	case "surname":
		u.sortPeopleListBySurname(peopleList)
	case "patronymic":
		u.sortPeopleListByPatronymic(peopleList)
	default:
		u.sortPeopleListByName(peopleList)
	}

	u.log.Info("данные успешно получены")
	return
}

func (u *PeopleUsecase) DeletePeople(peopleID int) (err error) {
	if peopleID <= 0 {
		u.log.Error(fmt.Sprintf("невалидный id человека %d", peopleID))
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

	if err = u.peopleRepo.DeletePeople(postgresql.SqlxTx(tx), peopleID); err != nil {
		u.log.Error(fmt.Sprintf("не удалось удалить человека, ошибка: %s", err))
		err = global.ErrInternalError
		return
	}

	if err = tx.Commit(); err != nil {
		u.log.Error(fmt.Sprintf("не удалось закрыть транзакцию, ошибка: %s", err))
		err = global.ErrInternalError
		return
	}

	u.log.Info("человек успещно обновлен")
	return
}

func (u *PeopleUsecase) UpdatePeople(p people.UpdatePeopleParam) (err error) {
	if !p.Valid() {
		u.log.Error("неверные параметры")
		err = global.ErrIncorectParams
		return
	}

	tx := u.sm.CreateSession()
	if err = tx.Start(); err != nil {
		u.log.Error(fmt.Sprintf("не удалось открыть транзакцию, ошибка: %s", err))
		err = global.ErrInternalError
		return
	}

	if err = u.peopleRepo.UpdatePeople(postgresql.SqlxTx(tx), p); err != nil {
		u.log.Error(fmt.Sprintf("не удалось обновить человека, ошибка: %s", err))
		err = global.ErrInternalError
		return
	}

	if err = tx.Commit(); err != nil {
		u.log.Error(fmt.Sprintf("не удалось закрыть транзакцию, ошибка: %s", err))
		err = global.ErrInternalError
		return
	}

	u.log.Info("человек успешно удален")
	return
}

func (u *PeopleUsecase) sortPeopleListByName(peopleList []people.People) {
	sort.Slice(peopleList, func(i, j int) bool {
		return peopleList[i].Name < peopleList[j].Name
	})
}

func (u *PeopleUsecase) sortPeopleListBySurname(peopleList []people.People) {
	sort.Slice(peopleList, func(i, j int) bool {
		return peopleList[i].Surname < peopleList[j].Surname
	})
}

func (u *PeopleUsecase) sortPeopleListByPatronymic(peopleList []people.People) {
	sort.Slice(peopleList, func(i, j int) bool {
		return peopleList[i].Patronymic < peopleList[j].Patronymic
	})
}
