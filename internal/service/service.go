package service

import (
	"github.com/find_property/internal/domain"
	"github.com/find_property/internal/repository"
	"github.com/find_property/pkg/worker"
)

type Users interface {
	CreateFilter(user *domain.Filter) error
	CheckUser(check *domain.UserSignUpInput) (bool, error)
	CreateView(view *domain.View) error
	CreateUser(user *domain.Filter) (*domain.Filter, error)
	GetUser(userID string) (*domain.Filter, error)
	SubwayLines(query string) ([]domain.SubwayLines, error)
	parse() error
	ReverseSweepJK(jk *domain.ReverseJK) error
	GetSubwayLines() ([]domain.SubwayLines, error)
	deleteBadAdverts() error
	proxy() ([]string, error)
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
	Search(filter *domain.StartSearch) (*domain.PagList, error)
}

type Service struct {
	Users
	Filters
}

func NewServices(
	db *repository.Repository,
	worker *worker.Worker,
) *Service {
	return &Service{
		Users:   NewUsersService(db.Users, worker),
		Filters: NewFiltersService(db.Filters),
	}
}
