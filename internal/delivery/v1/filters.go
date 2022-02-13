package v1

import (
	"github.com/find_property/internal/domain"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"net/http"
)

type updateSubcategory struct {
	Subcategory string `json:"subcategory" example:"Новостройка" binding:"required"`
	UserID      int    `json:"user_id" binding:"required"`
}

// @Summary Subcategory
// @Tags filters
// @Description обнавление субкатегории
// @ID subcategory
// @Accept  json
// @Produce  json
// @Param input body updateSubcategory true "только Новостройка или Вторичка слова начинаются с большой буквы"
// @Success 200 {string} string "ok"
// @Failure 400 {object} result
// @Failure 500 {object} result
// @Failure default {object} result
// @Router /subcategory [post]
func (h *Handler) Subcategory(c *gin.Context) {
	var req updateSubcategory
	if err := c.BindJSON(&req); err != nil {
		apiErrorEncode(c, err)
		return
	}

	if err := h.services.Filters.UpdateSubcategory(&domain.Filter{
		Subcategory: req.Subcategory,
		UserID:      req.UserID,
	}); err != nil {
		apiErrorEncode(c, err)
		return
	}

	c.Status(http.StatusOK)
}

type updatePrice struct {
	Price1 int `json:"price1" example:"7500000" binding:"required"`
	Price2 int `json:"price2" example:"9000000" binding:"required"`
	UserID int `json:"user_id" binding:"required"`
}

// @Summary Price
// @Tags filters
// @Description обнавление цены
// @ID price
// @Accept  json
// @Produce  json
// @Param input body updatePrice true "json price"
// @Success 200 {string} string "ok"
// @Failure 400 {object} result
// @Failure 500 {object} result
// @Failure default {object} result
// @Router /price [post]
func (h *Handler) Price(c *gin.Context) {
	var req updatePrice
	if err := c.BindJSON(&req); err != nil {
		apiErrorEncode(c, err)
		return
	}

	if err := h.services.Filters.UpdatePrice(&domain.Filter{
		Price1: req.Price1,
		Price2: req.Price2,
		UserID: req.UserID,
	}); err != nil {
		apiErrorEncode(c, err)
		return
	}

	c.Status(http.StatusOK)
}

type updateArea struct {
	Area1  string `json:"area1" example:"40" binding:"required"`
	Area2  string `json:"area2" example:"60" binding:"required"`
	UserID int    `json:"user_id" binding:"required"`
}

// @Summary Area
// @Tags filters
// @Description обнавление площади
// @ID area
// @Accept  json
// @Produce  json
// @Param input body updateArea true "json area"
// @Success 200 {string} string "ok"
// @Failure 400 {object} result
// @Failure 500 {object} result
// @Failure default {object} result
// @Router /area [post]
func (h *Handler) Area(c *gin.Context) {
	var req updateArea
	if err := c.BindJSON(&req); err != nil {
		apiErrorEncode(c, err)
		return
	}

	if err := h.services.Filters.UpdateArea(&domain.Filter{
		Area1:  req.Area1,
		Area2:  req.Area2,
		UserID: req.UserID,
	}); err != nil {
		apiErrorEncode(c, err)
		return
	}

	c.Status(http.StatusOK)
}

type updateJK struct {
	JKName string `json:"jk_name" example:"Не важно" binding:"required"`
	UserID int    `json:"user_id" binding:"required"`
}

// @Summary JK
// @Tags filters
// @Description обнавление ЖК
// @ID jk
// @Accept  json
// @Produce  json
// @Param input body updateJK true "Ввести название жк или Не важно"
// @Success 200 {string} string "ok"
// @Failure 400 {object} result
// @Failure 500 {object} result
// @Failure default {object} result
// @Router /jk [post]
func (h *Handler) JK(c *gin.Context) {
	var req updateJK
	if err := c.BindJSON(&req); err != nil {
		apiErrorEncode(c, err)
		return
	}

	if err := h.services.Filters.UpdateJK(&domain.Filter{
		JKName: req.JKName,
		UserID: req.UserID,
	}); err != nil {
		apiErrorEncode(c, err)
		return
	}

	c.Status(http.StatusOK)
}

type updateCategory struct {
	Category string `json:"category" example:"Квартиры и апартаменты" binding:"required"`
	UserID   int    `json:"user_id" binding:"required"`
}

// @Summary Category
// @Tags filters
// @Description обнавление категории дома
// @ID category
// @Accept  json
// @Produce  json
// @Param input body updateCategory true "Вводить 1 из 3: Только квартиры или Квартиры и апартаменты или Только апартаменты все с большой буквы!!"
// @Success 200 {string} string "ok"
// @Failure 400 {object} result
// @Failure 500 {object} result
// @Failure default {object} result
// @Router /category [post]
func (h *Handler) Category(c *gin.Context) {
	var req updateCategory
	if err := c.BindJSON(&req); err != nil {
		apiErrorEncode(c, err)
		return
	}

	if err := h.services.Filters.UpdateCategory(&domain.Filter{
		Category: req.Category,
		UserID:   req.UserID,
	}); err != nil {
		apiErrorEncode(c, err)
		return
	}

	c.Status(http.StatusOK)
}

type updateRooms struct {
	RoomsCount pq.StringArray `json:"rooms_count" example:"1-комн (евродва),2-комн (евротри)" swaggertype:"array,string" binding:"required"`
	UserID     int            `json:"user_id" binding:"required"`
}

// @Summary Rooms
// @Tags filters
// @Description обнавление комнат
// @ID rooms
// @Accept  json
// @Produce  json
// @Param input body updateRooms true "Название вводить строго такие: 1-комн (евродва), 2-комн (евротри), 3-комн, 4+комн, студия"
// @Success 200 {string} string "ok"
// @Failure 400 {object} result
// @Failure 500 {object} result
// @Failure default {object} result
// @Router /rooms [post]
func (h *Handler) Rooms(c *gin.Context) {
	var req updateRooms
	if err := c.BindJSON(&req); err != nil {
		apiErrorEncode(c, err)
		return
	}

	if err := h.services.Filters.UpdateRooms(&domain.Filter{
		RoomsCount: req.RoomsCount,
		UserID:     req.UserID,
	}); err != nil {
		apiErrorEncode(c, err)
		return
	}

	c.Status(http.StatusOK)
}

type updateMetro struct {
	Metro  pq.StringArray `json:"metro" example:"1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30,31,32,33,34,35" binding:"required" swaggertype:"array,string"`
	UserID int            `json:"user_id" binding:"required"`
}

// @Summary Metro
// @Tags filters
// @Description обнавление метро
// @ID metro
// @Accept  json
// @Produce  json
// @Param input body updateMetro true "Вводить или название метро или Не важно"
// @Success 200 {string} string "ok"
// @Failure 400 {object} result
// @Failure 500 {object} result
// @Failure default {object} result
// @Router /metro [post]
func (h *Handler) Metro(c *gin.Context) {
	var req updateMetro
	if err := c.BindJSON(&req); err != nil {
		apiErrorEncode(c, err)
		return
	}

	if err := h.services.Filters.UpdateMetro(&domain.Filter{
		Metro:  req.Metro,
		UserID: req.UserID,
	}); err != nil {
		apiErrorEncode(c, err)
		return
	}

	c.Status(http.StatusOK)
}

type updateMetroBefore struct {
	BeforeMetro string `json:"before_metro" example:"Не важно" binding:"required"`
	UserID      int    `json:"user_id" binding:"required"`
}

// @Summary MetroBefore
// @Tags filters
// @Description обнавление до метро
// @ID metro-before
// @Accept  json
// @Produce  json
// @Param input body updateMetroBefore true "Вводить Только пешком с большой буквы или Не важно"
// @Success 200 {string} string "ok"
// @Failure 400 {object} result
// @Failure 500 {object} result
// @Failure default {object} result
// @Router /metro/before [post]
func (h *Handler) MetroBefore(c *gin.Context) {
	var req updateMetroBefore
	if err := c.BindJSON(&req); err != nil {
		apiErrorEncode(c, err)
		return
	}

	if err := h.services.Filters.UpdateMetroBefore(&domain.Filter{
		BeforeMetro: req.BeforeMetro,
		UserID:      req.UserID,
	}); err != nil {
		apiErrorEncode(c, err)
		return
	}

	c.Status(http.StatusOK)
}

type updateBuildingYear struct {
	BuildingYear1 int `json:"building_year1" example:"2010" binding:"required"`
	BuildingYear2 int `json:"building_year2" example:"2024" binding:"required"`
	UserID        int `json:"user_id" binding:"required"`
}

// @Summary BuildingYear
// @Tags filters
// @Description обнавление года постройки
// @ID building-year
// @Accept  json
// @Produce  json
// @Param input body updateBuildingYear true "building_year1 строго с 2010, building_year2 любое"
// @Success 200 {string} string "ok"
// @Failure 400 {object} result
// @Failure 500 {object} result
// @Failure default {object} result
// @Router /buildingYear [post]
func (h *Handler) BuildingYear(c *gin.Context) {
	var req updateBuildingYear
	if err := c.BindJSON(&req); err != nil {
		apiErrorEncode(c, err)
		return
	}

	if err := h.services.Filters.UpdateBuildingYear(&domain.Filter{
		BuildingYear1: req.BuildingYear1,
		BuildingYear2: req.BuildingYear2,
		UserID:        req.UserID,
	}); err != nil {
		apiErrorEncode(c, err)
		return
	}

	c.Status(http.StatusOK)
}

type updateTypeBalcony struct {
	TypeBalcony string `json:"type_balcony" example:"Не важно" binding:"required"`
	UserID      int    `json:"user_id" binding:"required"`
}

// @Summary Balcony
// @Tags filters
// @Description обнавление ложия или балкон
// @ID balcony
// @Accept  json
// @Produce  json
// @Param input body updateTypeBalcony true "Типы балконо: Не важно или Лоджия или Балкон, All все с большой буквы"
// @Success 200 {string} string "ok"
// @Failure 400 {object} result
// @Failure 500 {object} result
// @Failure default {object} result
// @Router /balcony [post]
func (h *Handler) Balcony(c *gin.Context) {
	var req updateTypeBalcony
	if err := c.BindJSON(&req); err != nil {
		apiErrorEncode(c, err)
		return
	}

	if err := h.services.Filters.UpdateBalcony(&domain.Filter{
		TypeBalcony: req.TypeBalcony,
		UserID:      req.UserID,
	}); err != nil {
		apiErrorEncode(c, err)
		return
	}

	c.Status(http.StatusOK)
}

type updateFloors struct {
	Floors string `json:"floors" example:"Только не последний" binding:"required"`
	UserID int    `json:"user_id" binding:"required"`
}

// @Summary Floors
// @Tags filters
// @Description обнавление этажей
// @ID floors
// @Accept  json
// @Produce  json
// @Param input body updateFloors true " или Только не первый или Только не последний или All"
// @Success 200 {string} string "ok"
// @Failure 400 {object} result
// @Failure 500 {object} result
// @Failure default {object} result
// @Router /floors [post]
func (h *Handler) Floors(c *gin.Context) {
	var req updateFloors
	if err := c.BindJSON(&req); err != nil {
		apiErrorEncode(c, err)
		return
	}

	if err := h.services.Filters.UpdateFloors(&domain.Filter{
		Floors: req.Floors,
		UserID: req.UserID,
	}); err != nil {
		apiErrorEncode(c, err)
		return
	}

	c.Status(http.StatusOK)
}

type updateRepairType struct {
	RepairType string `json:"repair_type" example:"Не важно" binding:"required"`
	UserID     int    `json:"user_id" binding:"required"`
}

// @Summary RepairType
// @Tags filters
// @Description обнавление ремонта
// @ID repair-type
// @Accept  json
// @Produce  json
// @Param input body updateRepairType true "Типы ремонта строго такие: Чистовая, Черновая, Нет, Не важно"
// @Success 200 {string} string "ok"
// @Failure 400 {object} result
// @Failure 500 {object} result
// @Failure default {object} result
// @Router /repair [post]
func (h *Handler) RepairType(c *gin.Context) {
	var req updateRepairType
	if err := c.BindJSON(&req); err != nil {
		apiErrorEncode(c, err)
		return
	}

	if err := h.services.Filters.UpdateRepairType(&domain.Filter{
		RepairType: req.RepairType,
		UserID:     req.UserID,
	}); err != nil {
		apiErrorEncode(c, err)
		return
	}

	c.Status(http.StatusOK)
}

// @Summary Поиск
// @Tags search
// @Description поиск
// @ID search
// @Accept  json
// @Produce  json
// @Param input body domain.StartSearch true "поиск по юзер айди фильтру с пагинацией"
// @Success 200 {array} domain.Advert
// @Failure 400 {object} result
// @Failure 500 {object} result
// @Failure default {object} result
// @Router /search [post]
func (h *Handler) Search(c *gin.Context) {
	req := new(domain.StartSearch)
	if err := c.BindJSON(req); err != nil {
		apiErrorEncode(c, err)
		return
	}

	adverts, err := h.services.Filters.Search(req)
	if err != nil {
		apiErrorEncode(c, err)
		return
	}

	c.JSON(http.StatusOK, adverts)
}

type updateAddress struct {
	Address string `json:"address" example:"Не важно" binding:"required"`
	UserID  int    `json:"user_id" binding:"required"`
}

// @Summary Address
// @Tags filters
// @Description обнавление адреса
// @ID address
// @Accept  json
// @Produce  json
// @Param input body updateAddress true "Ввести название например улицы или Не важно"
// @Success 200 {string} string "ok"
// @Failure 400 {object} result
// @Failure 500 {object} result
// @Failure default {object} result
// @Router /address [post]
func (h *Handler) Address(c *gin.Context) {
	var req updateAddress
	if err := c.BindJSON(&req); err != nil {
		apiErrorEncode(c, err)
		return
	}

	if err := h.services.Filters.UpdateAddress(&domain.Filter{
		Address: req.Address,
		UserID:  req.UserID,
	}); err != nil {
		apiErrorEncode(c, err)
		return
	}

	c.Status(http.StatusOK)
}

type updateDefault struct {
	Default string `json:"default"`
	UserID  int    `json:"user_id" binding:"required"`
}

// @Summary Default
// @Tags filters
// @Description обнавление дефолта
// @ID default
// @Accept  json
// @Produce  json
// @Param input body updateDefault true "для дефолта"
// @Success 200 {string} string "ok"
// @Failure 400 {object} result
// @Failure 500 {object} result
// @Failure default {object} result
// @Router /default [post]
func (h *Handler) Default(c *gin.Context) {
	var req updateDefault
	if err := c.BindJSON(&req); err != nil {
		apiErrorEncode(c, err)
		return
	}

	if err := h.services.Filters.UpdateDefault(&domain.Filter{
		Default: req.Default,
		UserID:  req.UserID,
	}); err != nil {
		apiErrorEncode(c, err)
		return
	}

	c.Status(http.StatusOK)
}

type updateSilence struct {
	Silence string `json:"silence"`
	UserID  int    `json:"user_id" binding:"required"`
}

// @Summary Silence
// @Tags filters
// @Description обнавление дефолта1
// @ID silence
// @Accept  json
// @Produce  json
// @Param input body updateSilence true "для дефолта1"
// @Success 200 {string} string "ok"
// @Failure 400 {object} result
// @Failure 500 {object} result
// @Failure default {object} result
// @Router /silence [post]
func (h *Handler) Silence(c *gin.Context) {
	var req updateSilence
	if err := c.BindJSON(&req); err != nil {
		apiErrorEncode(c, err)
		return
	}

	if err := h.services.Filters.UpdateSilence(&domain.Filter{
		Silence: req.Silence,
		UserID:  req.UserID,
	}); err != nil {
		apiErrorEncode(c, err)
		return
	}

	c.Status(http.StatusOK)
}
