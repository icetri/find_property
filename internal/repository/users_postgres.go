package repository

import (
	"fmt"
	"github.com/find_property/internal/domain"
	"gorm.io/gorm"
)

type UsersRepo struct {
	db *gorm.DB
}

func NewUsersRepo(db *gorm.DB) *UsersRepo {
	return &UsersRepo{
		db: db,
	}
}

func (p *UsersRepo) Create(user *domain.Filter) error {

	if err := p.db.Debug().Table("filters").Create(user).Error; err != nil {
		return err
	}

	return nil
}

func (p *UsersRepo) GetUser(check *domain.UserSignUpInput) (*domain.User, error) {

	user := new(domain.User)
	if err := p.db.Debug().Table("filters").Where("user_id = ?", check.UserID).Take(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (p *UsersRepo) AddView(view *domain.View) error {

	if err := p.db.Debug().Table("viewed").Create(view).Error; err != nil {
		return err
	}

	return nil
}

func (p *UsersRepo) Parse(cianParses *domain.CianParse) error {

	if err := p.db.Debug().Table("adverts").Create(cianParses).Error; err != nil {
		return err
	}

	return nil
}

func (p *UsersRepo) GetSubwayLines(query string) ([]domain.SubwayLines, error) {

	lines := make([]domain.SubwayLines, 0)
	if err := p.db.Debug().Table("metro_directory").Where(fmt.Sprintf("line_name ILIKE '%s'", query)).Find(&lines).Error; err != nil {
		return nil, err
	}

	return lines, nil
}

func (p *UsersRepo) GetSubwaysLines() ([]domain.SubwayLines, error) {

	lines := make([]domain.SubwayLines, 0)
	if err := p.db.Debug().Table("metro_directory").Find(&lines).Error; err != nil {
		return nil, err
	}

	return lines, nil
}

func (p *UsersRepo) CreateUser(filter *domain.Filter) error {

	if err := p.db.Transaction(func(tx *gorm.DB) error {

		newUser := new(domain.Filter)
		if p.db.Debug().Table("filters").Where("user_id = ?", filter.UserID).Take(newUser).RowsAffected != 0 {
			return domain.ErrorUserIsExist
		}

		if err := p.db.Debug().Table("filters").Create(filter).Error; err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}

func (p *UsersRepo) GetFilterUser(userID string) (*domain.Filter, error) {

	filter := new(domain.Filter)
	if err := p.db.Debug().Table("filters").Where("user_id = ?", userID).Take(filter).Error; err != nil {
		return nil, err
	}

	return filter, nil
}

func (p *UsersRepo) CreateMetro(metro *domain.Metro) error {

	if err := p.db.Debug().Table("metro").Create(metro).Error; err != nil {
		return err
	}

	return nil
}

func (p *UsersRepo) UpdateFilterLines(userID int, user *domain.Filter) error {

	if err := p.db.Debug().Table("filters").Where("user_id = ?", userID).Update("jk_ids", user.JKIDs).Error; err != nil {
		return err
	}

	return nil
}

func (p *UsersRepo) AllAdverts() (domain.Adverts, error) {

	adverts := make(domain.Adverts, 0)
	if err := p.db.Debug().Table("adverts").Order("time_publish ASC").Find(&adverts).Error; err != nil {
		return nil, err
	}

	return adverts, nil
}

func (p *UsersRepo) DeleteAdvert(advertID int) error {

	if err := p.db.Debug().Table("adverts").Where("id = ?", advertID).Delete(&domain.Advert{}).Error; err != nil {
		return err
	}

	return nil
}

func (p *UsersRepo) DeleteMetros(cianID int) error {

	if err := p.db.Debug().Table("metro").Where("cian_id", cianID).Delete(&domain.Metro{}).Error; err != nil {
		return err
	}

	return nil
}
