package repository

import (
	"github.com/find_property/internal/domain"
	"gorm.io/gorm"
)

type Users interface {
	Create(user *domain.Filter) error
	GetUser(check *domain.UserSignUpInput) (*domain.User, error)
	GetFilterUser(userID string) (*domain.Filter, error)
	AddView(view *domain.View) error
	CreateUser(filter *domain.Filter) error
	Parse(cianParses *domain.CianParse) error
	CreateMetro(metro *domain.Metro) error
	GetSubwayLines(query string) ([]domain.SubwayLines, error)
	GetSubwaysLines() ([]domain.SubwayLines, error)
	UpdateFilterLines(userID int, user *domain.Filter) error
	AllAdverts() (domain.Adverts, error)
	DeleteAdvert(advertID int) error
	DeleteMetros(cianID int) error
}

type Filters interface {
	UpdateSubcategory(filter *domain.Filter) error
	UpdatePrice(filter *domain.Filter) error
	UpdateArea(filter *domain.Filter) error
	UpdateJK(filter *domain.Filter) error
	UpdateCategory(filter *domain.Filter) error
	UpdateRooms(filter *domain.Filter) error
	UpdateMetro(filter *domain.Filter) error
	UpdateMetroBefore(filter *domain.Filter) error
	UpdateBuildingYear(filter *domain.Filter) error
	UpdateBalcony(filter *domain.Filter) error
	UpdateFloors(filter *domain.Filter) error
	UpdateRepairType(filter *domain.Filter) error
	UpdateAddress(filter *domain.Filter) error
	UpdateDefault(filter *domain.Filter) error
	UpdateSilence(filter *domain.Filter) error
	GetUserByID(userID int) (*domain.Filter, error)
	Search(where string, user *domain.Filter, filter *domain.StartSearch) (domain.Adverts, error)
	SearchCount(where string, user *domain.Filter) (int64, error)
	Underground(advertID int) ([]domain.Metro, error)
}

type Repository struct {
	Users
	Filters
}

func NewRepositories(db *gorm.DB) *Repository {
	return &Repository{
		Users:   NewUsersRepo(db),
		Filters: NewFiltersRepo(db),
	}
}
