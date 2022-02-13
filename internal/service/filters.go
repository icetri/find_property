package service

import (
	"fmt"
	"github.com/find_property/internal/domain"
	"github.com/find_property/internal/repository"
	"github.com/find_property/pkg/logger"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type FiltersService struct {
	repo repository.Filters
}

func NewFiltersService(repo repository.Filters) *FiltersService {
	return &FiltersService{
		repo: repo,
	}
}

func (s *FiltersService) UpdateSubcategory(filter *domain.Filter) error {

	if err := s.repo.UpdateSubcategory(filter); err != nil {
		logger.LogError(errors.Wrap(err, "err with UpdateSubcategory in UpdateSubcategory"))
		return domain.ErrorInternalServerError
	}

	return nil
}

func (s *FiltersService) UpdatePrice(filter *domain.Filter) error {

	if err := s.repo.UpdatePrice(filter); err != nil {
		logger.LogError(errors.Wrap(err, "err with UpdatePrice in UpdatePrice"))
		return domain.ErrorInternalServerError
	}

	return nil
}

func (s *FiltersService) UpdateArea(filter *domain.Filter) error {

	if err := s.repo.UpdateArea(filter); err != nil {
		logger.LogError(errors.Wrap(err, "err with UpdateArea in UpdateArea"))
		return domain.ErrorInternalServerError
	}

	return nil
}

func (s *FiltersService) UpdateJK(filter *domain.Filter) error {

	if err := s.repo.UpdateJK(filter); err != nil {
		logger.LogError(errors.Wrap(err, "err with UpdateJK in UpdateJK"))
		return domain.ErrorInternalServerError
	}

	return nil
}

func (s *FiltersService) UpdateCategory(filter *domain.Filter) error {

	if err := s.repo.UpdateCategory(filter); err != nil {
		logger.LogError(errors.Wrap(err, "err with UpdateJK in UpdateCategory"))
		return domain.ErrorInternalServerError
	}

	return nil
}

func (s *FiltersService) UpdateRooms(filter *domain.Filter) error {

	if err := s.repo.UpdateRooms(filter); err != nil {
		logger.LogError(errors.Wrap(err, "err with UpdateRooms in UpdateRooms"))
		return domain.ErrorInternalServerError
	}

	return nil
}

func (s *FiltersService) UpdateMetro(filter *domain.Filter) error {

	if err := s.repo.UpdateMetro(filter); err != nil {
		logger.LogError(errors.Wrap(err, "err with UpdateMetro in UpdateMetro"))
		return domain.ErrorInternalServerError
	}

	return nil
}

func (s *FiltersService) UpdateMetroBefore(filter *domain.Filter) error {

	if err := s.repo.UpdateMetroBefore(filter); err != nil {
		logger.LogError(errors.Wrap(err, "err with UpdateMetroBefore in UpdateMetroBefore"))
		return domain.ErrorInternalServerError
	}

	return nil
}

func (s *FiltersService) UpdateBuildingYear(filter *domain.Filter) error {

	if err := s.repo.UpdateBuildingYear(filter); err != nil {
		logger.LogError(errors.Wrap(err, "err with UpdateBuildingYear in UpdateBuildingYear"))
		return domain.ErrorInternalServerError
	}

	return nil
}

func (s *FiltersService) UpdateBalcony(filter *domain.Filter) error {

	if err := s.repo.UpdateBalcony(filter); err != nil {
		logger.LogError(errors.Wrap(err, "err with UpdateBalcony in UpdateBalcony"))
		return domain.ErrorInternalServerError
	}

	return nil
}

func (s *FiltersService) UpdateFloors(filter *domain.Filter) error {

	if err := s.repo.UpdateFloors(filter); err != nil {
		logger.LogError(errors.Wrap(err, "err with UpdateFloors in UpdateFloors"))
		return domain.ErrorInternalServerError
	}

	return nil
}

func (s *FiltersService) UpdateRepairType(filter *domain.Filter) error {

	if err := s.repo.UpdateRepairType(filter); err != nil {
		logger.LogError(errors.Wrap(err, "err with UpdateRepairType in UpdateRepairType"))
		return domain.ErrorInternalServerError
	}

	return nil
}

func (s *FiltersService) UpdateAddress(filter *domain.Filter) error {

	if err := s.repo.UpdateAddress(filter); err != nil {
		logger.LogError(errors.Wrap(err, "err with UpdateAddress in UpdateAddress"))
		return domain.ErrorInternalServerError
	}

	return nil
}

func (s *FiltersService) UpdateDefault(filter *domain.Filter) error {

	if err := s.repo.UpdateDefault(filter); err != nil {
		logger.LogError(errors.Wrap(err, "err with UpdateDefault in UpdateDefault"))
		return domain.ErrorInternalServerError
	}

	return nil
}

func (s *FiltersService) UpdateSilence(filter *domain.Filter) error {

	if err := s.repo.UpdateSilence(filter); err != nil {
		logger.LogError(errors.Wrap(err, "err with UpdateDefault in UpdateDefault"))
		return domain.ErrorInternalServerError
	}

	return nil
}

func (s *FiltersService) Search(filter *domain.StartSearch) (*domain.PagList, error) {

	user, err := s.repo.GetUserByID(filter.UserID)
	if err != nil && err != gorm.ErrRecordNotFound {
		logger.LogError(errors.Wrap(err, "err with GetUserByID in Search"))
		return nil, domain.ErrorInternalServerError
	}

	if user == nil {
		return nil, domain.ErrorUserNotExist
	}

	where := s.advertConstructor(user)

	adverts, err := s.repo.Search(where, user, filter)
	if err != nil {
		logger.LogError(errors.Wrap(err, "err with Search in Search"))
		return nil, domain.ErrorInternalServerError
	}

	count, err := s.repo.SearchCount(where, user)
	if err != nil {
		logger.LogError(errors.Wrap(err, "err with SearchCount in Search"))
		return nil, domain.ErrorInternalServerError
	}

	for i := range adverts {
		adverts[i].Underground, err = s.repo.Underground(adverts[i].ID)
		if err != nil {
			logger.LogError(errors.Wrap(err, "err with Underground in Search"))
			return nil, domain.ErrorInternalServerError
		}
	}

	return &domain.PagList{
		Items: adverts,
		Count: count,
	}, nil
}

func (s *FiltersService) advertConstructor(user *domain.Filter) string {

	var where string

	where = fmt.Sprintf(`rooms_count = ANY(?)
AND jk_id != All(?)
AND (m.metro = ANY(?) or m.line_id = ANY(?))
AND price BETWEEN %d AND %d
AND CAST(area AS double precision) BETWEEN CAST('%s' AS double precision) AND CAST('%s' AS double precision)
AND building_year BETWEEN %d AND %d`,
		user.Price1, user.Price2,
		user.Area1, user.Area2,
		user.BuildingYear1, user.BuildingYear2)

	if user.BeforeMetro != "Не важно" {
		where = where + fmt.Sprintf(" AND before_metro = true")
	}

	if user.JKName != "Не важно" {
		user.JKName = "%" + user.JKName + "%"
		where = where + fmt.Sprintf(" AND jk_name ILIKE '%s'", user.JKName)
	}

	if user.Address != "Не важно" {
		user.Address = "%" + user.Address + "%"
		where = where + fmt.Sprintf(" AND address ILIKE '%s'", user.Address)
	}

	//if user.Metro != "Не важно" {
	//	user.Metro = "%" + user.Metro + "%"
	//	where = where + fmt.Sprintf(" AND metro ILIKE '%s'", user.Metro)
	//}

	switch user.Floors {
	case "Только не первый":
		where = where + fmt.Sprintf(" AND floor_number != 1")
	case "Только не последний":
		where = where + fmt.Sprintf(" AND floor_number != floors_count")
	case "All":
		where = where + fmt.Sprintf(" AND floor_number != floors_count AND floor_number != 1")
	}

	switch user.Category {
	case "Только квартиры":
		where = where + fmt.Sprintf(" AND is_apartments = false")
	case "Только апартаменты":
		where = where + fmt.Sprintf(" AND is_apartments = true")
	}

	switch user.TypeBalcony {
	case "Балкон":
		where = where + fmt.Sprintf(" AND balconies = true")
	case "Лоджия":
		where = where + fmt.Sprintf(" AND loggias = true")
	case "All":
		where = where + fmt.Sprintf(" AND loggias = true AND balconies = true")
	}

	if user.RepairType != "Не важно" {
		where = where + fmt.Sprintf(" AND decoration = '%s'", user.RepairType)
	}

	where = where + fmt.Sprintf(" AND a.id NOT IN (SELECT DISTINCT cian_id FROM viewed WHERE cian_id = a.id AND user_id = %d)", user.UserID)
	where = where + " AND true"

	return where
}
