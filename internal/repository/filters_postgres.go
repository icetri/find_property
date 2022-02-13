package repository

import (
	"fmt"
	"github.com/find_property/internal/domain"

	"gorm.io/gorm"
)

type FiltersRepo struct {
	db *gorm.DB
}

func NewFiltersRepo(db *gorm.DB) *FiltersRepo {
	return &FiltersRepo{
		db: db,
	}
}

func (p *FiltersRepo) UpdateSubcategory(filter *domain.Filter) error {

	if err := p.db.Debug().Table("filters").Select("subcategory").Where("user_id", filter.UserID).Updates(filter).Error; err != nil {
		return err
	}

	return nil
}

func (p *FiltersRepo) UpdatePrice(filter *domain.Filter) error {

	if err := p.db.Debug().Table("filters").Select("price1", "price2").Where("user_id", filter.UserID).Updates(filter).Error; err != nil {
		return err
	}

	return nil
}

func (p *FiltersRepo) UpdateArea(filter *domain.Filter) error {

	if err := p.db.Debug().Table("filters").Select("area1", "area2").Where("user_id", filter.UserID).Updates(filter).Error; err != nil {
		return err
	}

	return nil
}

func (p *FiltersRepo) UpdateJK(filter *domain.Filter) error {

	if err := p.db.Debug().Table("filters").Select("jk_name").Where("user_id", filter.UserID).Updates(filter).Error; err != nil {
		return err
	}

	return nil
}

func (p *FiltersRepo) UpdateCategory(filter *domain.Filter) error {

	if err := p.db.Debug().Table("filters").Select("category").Where("user_id", filter.UserID).Updates(filter).Error; err != nil {
		return err
	}

	return nil
}

func (p *FiltersRepo) UpdateRooms(filter *domain.Filter) error {

	if err := p.db.Debug().Table("filters").Select("rooms_count").Where("user_id", filter.UserID).Updates(filter).Error; err != nil {
		return err
	}

	return nil
}

func (p *FiltersRepo) UpdateMetro(filter *domain.Filter) error {

	if err := p.db.Debug().Table("filters").Select("metro").Where("user_id", filter.UserID).Updates(filter).Error; err != nil {
		return err
	}

	return nil
}

func (p *FiltersRepo) UpdateMetroBefore(filter *domain.Filter) error {

	if err := p.db.Debug().Table("filters").Select("before_metro").Where("user_id", filter.UserID).Updates(filter).Error; err != nil {
		return err
	}

	return nil
}

func (p *FiltersRepo) UpdateBuildingYear(filter *domain.Filter) error {

	if err := p.db.Debug().Table("filters").Select("building_year1", "building_year2").Where("user_id", filter.UserID).Updates(filter).Error; err != nil {
		return err
	}

	return nil
}

func (p *FiltersRepo) UpdateBalcony(filter *domain.Filter) error {

	if err := p.db.Debug().Table("filters").Select("type_balcony").Where("user_id", filter.UserID).Updates(filter).Error; err != nil {
		return err
	}

	return nil
}

func (p *FiltersRepo) UpdateFloors(filter *domain.Filter) error {

	if err := p.db.Debug().Table("filters").Select("floors").Where("user_id", filter.UserID).Updates(filter).Error; err != nil {
		return err
	}

	return nil
}

func (p *FiltersRepo) UpdateRepairType(filter *domain.Filter) error {

	if err := p.db.Debug().Table("filters").Select("repair_type").Where("user_id", filter.UserID).Updates(filter).Error; err != nil {
		return err
	}

	return nil
}

func (p *FiltersRepo) UpdateAddress(filter *domain.Filter) error {

	if err := p.db.Debug().Table("filters").Select("address").Where("user_id", filter.UserID).Updates(filter).Error; err != nil {
		return err
	}

	return nil
}

func (p *FiltersRepo) UpdateDefault(filter *domain.Filter) error {

	if err := p.db.Debug().Table("filters").Select("default").Where("user_id", filter.UserID).Updates(filter).Error; err != nil {
		return err
	}

	return nil
}

func (p *FiltersRepo) UpdateSilence(filter *domain.Filter) error {

	if err := p.db.Debug().Table("filters").Select("silence").Where("user_id", filter.UserID).Updates(filter).Error; err != nil {
		return err
	}

	return nil
}

func (p *FiltersRepo) GetUserByID(userID int) (*domain.Filter, error) {

	filter := new(domain.Filter)
	if err := p.db.Debug().Table("filters").Where("user_id = ?", userID).Take(filter).Error; err != nil {
		return nil, err
	}

	return filter, nil
}

func (p *FiltersRepo) Search(where string, user *domain.Filter, filter *domain.StartSearch) (domain.Adverts, error) {

	adverts := make(domain.Adverts, 0)

	if err := p.db.Debug().Raw(fmt.Sprintf(`SELECT DISTINCT (a.*) FROM adverts a
LEFT JOIN metro m on a.id = m.cian_id WHERE %s 
GROUP BY a.id, m.id ORDER BY a.time_publish DESC LIMIT ? OFFSET ?`, where), user.RoomsCount, user.JKIDs, user.Metro, user.Metro, filter.Limit, filter.Offset).Find(&adverts).Error; err != nil {
		return nil, err
	}

	return adverts, nil
}

func (p *FiltersRepo) SearchCount(where string, user *domain.Filter) (int64, error) {

	var count int64
	if err := p.db.Debug().Raw(fmt.Sprintf(`SELECT count(*) FROM (SELECT DISTINCT a.* FROM adverts a
LEFT JOIN metro m on a.id = m.cian_id WHERE %s 
GROUP BY a.id, m.id) d1`, where), user.RoomsCount, user.JKIDs, user.Metro, user.Metro).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

func (p *FiltersRepo) Underground(advertID int) ([]domain.Metro, error) {

	underground := make([]domain.Metro, 0)
	if err := p.db.Debug().Table("metro").Where("cian_id = ?", advertID).Find(&underground).Error; err != nil {
		return nil, err
	}

	return underground, nil
}
