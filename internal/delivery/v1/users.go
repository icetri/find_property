package v1

import (
	"github.com/find_property/internal/domain"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"net/http"
)

// @Summary CheckUserRegistered
// @Tags user
// @Description Всегда после /start
// @ID register-or-login
// @Accept  json
// @Produce  json
// @Param input body domain.UserSignUpInput true "регистрация по телеграмм юзер айди"
// @Success 200 {bool} bool
// @Failure 400 {object} result
// @Failure 500 {object} result
// @Failure default {object} result
// @Router /login [post]
func (h *Handler) CheckUserRegistered(c *gin.Context) {
	req := new(domain.UserSignUpInput)
	if err := c.BindJSON(req); err != nil {
		apiErrorEncode(c, err)
		return
	}

	isRegistred, err := h.services.Users.CheckUser(req)
	if err != nil {
		apiErrorEncode(c, err)
		return
	}

	c.JSON(http.StatusOK, isRegistred)
}

type createFilter struct {
	Subcategory   string         `json:"subcategory" example:"Новостройка"`
	RoomsCount    pq.StringArray `json:"rooms_count" example:"1-комн (евродва),2-комн (евротри)" swaggertype:"array,string"`                                                         //example [1-комн (евродва),2-комн (евротри),3-комн,4+комн,студия]
	Price1        int            `json:"price1" example:"10000000"`                                                                                                                  //example 6459462
	Price2        int            `json:"price2" example:"100000000"`                                                                                                                 //example 7194195
	Area1         string         `json:"area1" example:"80"`                                                                                                                         //example 55.9
	Area2         string         `json:"area2" example:"150"`                                                                                                                        //example 76.0
	Metro         pq.StringArray `json:"metro" example:"1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30,31,32,33,34,35" swaggertype:"array,string"` //example Речной вокзал
	BeforeMetro   string         `json:"before_metro" example:"Не важно"`                                                                                                            //example неважно, только пешком
	RepairType    string         `json:"repair_type" example:"Не важно"`                                                                                                             //example euro
	BuildingYear1 int            `json:"building_year1" example:"2010"`                                                                                                              //example 2010
	BuildingYear2 int            `json:"building_year2" example:"2024"`                                                                                                              //example 2024
	JKName        string         `json:"jk_name" example:"Не важно"`                                                                                                                 //example ЖК ... либо пустое
	Address       string         `json:"address" example:"Не важно"`                                                                                                                 //example ЖК ... либо пустое
	Category      string         `json:"category" example:"Квартиры и апартаменты"`                                                                                                  //example только квартиры, квартиры и апартаменты, только апартаменты
	Floors        string         `json:"floors" example:"Только не первый"`                                                                                                          //example только не первый
	TypeBalcony   string         `json:"type_balcony" example:"Не важно"`                                                                                                            //example неважно,балкон, лоджия

	UserID int `json:"user_id"`
}

// @Summary RegistrationUser
// @Tags user
// @Description Добавление всех фильтров и регистрация разом
// @ID registration-user
// @Accept  json
// @Produce  json
// @Param input body createFilter true "регистрация все фильтров пользователю"
// @Success 201 {string} string "ok"
// @Failure 400 {object} result
// @Failure 500 {object} result
// @Failure default {object} result
// @Router /create [post]
func (h *Handler) RegistrationUser(c *gin.Context) {
	var req createFilter
	if err := c.BindJSON(&req); err != nil {
		apiErrorEncode(c, err)
		return
	}

	if err := h.services.Users.CreateFilter(&domain.Filter{
		Subcategory:   req.Subcategory,
		RoomsCount:    req.RoomsCount,
		Price1:        req.Price1,
		Price2:        req.Price2,
		Area1:         req.Area1,
		Area2:         req.Area2,
		Metro:         req.Metro,
		BeforeMetro:   req.BeforeMetro,
		RepairType:    req.RepairType,
		BuildingYear1: req.BuildingYear1,
		BuildingYear2: req.BuildingYear2,
		JKName:        req.JKName,
		Address:       req.Address,
		Category:      req.Category,
		Floors:        req.Floors,
		TypeBalcony:   req.TypeBalcony,
		UserID:        req.UserID,
	}); err != nil {
		apiErrorEncode(c, err)
		return
	}

	c.Status(http.StatusCreated)
}

type createView struct {
	UserID int `json:"user_id"`
	CianID int `json:"cian_id"`
}

// @Summary AdViewed
// @Tags user
// @Description Добавление объявлений в просмотренные
// @ID ad-viewed
// @Accept  json
// @Produce  json
// @Param input body createView true "Необходим user_id телеграма и id циан объявления"
// @Success 200 {string} string "ok"
// @Failure 400 {object} result
// @Failure 500 {object} result
// @Failure default {object} result
// @Router /view/add [post]
func (h *Handler) AdViewed(c *gin.Context) {
	var req createView
	if err := c.BindJSON(&req); err != nil {
		apiErrorEncode(c, err)
		return
	}

	if err := h.services.Users.CreateView(&domain.View{
		UserID: req.UserID,
		CianID: req.CianID,
	}); err != nil {
		apiErrorEncode(c, err)
		return
	}

	c.Status(http.StatusOK)
}

type createUser struct {
	UserID int `json:"user_id" binding:"required"`
}

// @Summary CreateUser
// @Tags user
// @Description Добавление юзера если у него нет фильтров
// @ID create-user
// @Accept  json
// @Produce  json
// @Param input body createUser true "Необходим user_id телеграма"
// @Success 201 {object} domain.Filter
// @Failure 400 {object} result
// @Failure 500 {object} result
// @Failure default {object} result
// @Router /user/create [post]
func (h *Handler) CreateUser(c *gin.Context) {
	var req createUser
	if err := c.BindJSON(&req); err != nil {
		apiErrorEncode(c, err)
		return
	}

	user, err := h.services.Users.CreateUser(&domain.Filter{
		UserID: req.UserID,
	})
	if err != nil {
		apiErrorEncode(c, err)
		return
	}

	c.JSON(http.StatusCreated, user)
}

// @Summary GetUser
// @Tags user
// @Description Инфа о юзере
// @ID get-user
// @Accept  json
// @Produce  json
// @Param id path string true "Необходим user_id телеграма"
// @Success 200 {object} domain.Filter
// @Failure 400 {object} result
// @Failure 500 {object} result
// @Failure default {object} result
// @Router /user/{id} [get]
func (h *Handler) GetUser(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		apiErrorEncode(c, domain.ErrorBadRequest)
		return
	}

	user, err := h.services.Users.GetUser(id)
	if err != nil {
		apiErrorEncode(c, err)
		return
	}

	c.JSON(http.StatusOK, user)
}

type findLines struct {
	Query string `json:"query" binding:"required"`
}

// @Summary GetSubwayLines
// @Tags user
// @Description автокомплит для линий метро
// @ID get-subway-lines
// @Accept  json
// @Produce  json
// @Param input body findLines true "кверики"
// @Success 200 {array} domain.SubwayLines
// @Failure 400 {object} result
// @Failure 500 {object} result
// @Failure default {object} result
// @Router /user/subway/lines [post]
func (h *Handler) GetSubwayLines(c *gin.Context) {
	var req findLines
	if err := c.BindJSON(&req); err != nil {
		apiErrorEncode(c, err)
		return
	}

	lines, err := h.services.Users.SubwayLines(req.Query)
	if err != nil {
		apiErrorEncode(c, err)
		return
	}

	c.JSON(http.StatusOK, lines)
}

// @Summary AddReverseSweepJK
// @Tags user
// @Description запись на повтор жк
// @ID add-reverse-sweep-jk
// @Accept  json
// @Produce  json
// @Param input body domain.ReverseJK true "kek"
// @Success 200 {string} string OK
// @Failure 400 {object} result
// @Failure 500 {object} result
// @Failure default {object} result
// @Router /jk/add [post]
func (h *Handler) AddReverseSweepJK(c *gin.Context) {
	req := new(domain.ReverseJK)
	if err := c.BindJSON(req); err != nil {
		apiErrorEncode(c, err)
		return
	}

	err := h.services.Users.ReverseSweepJK(req)
	if err != nil {
		apiErrorEncode(c, err)
		return
	}

	c.Status(http.StatusOK)
}

// @Summary Subways
// @Tags user
// @Description все линий
// @ID subways
// @Accept  json
// @Produce  json
// @Success 200 {array} domain.SubwayLines
// @Failure 500 {object} result
// @Failure default {object} result
// @Router /subways [get]
func (h *Handler) Subways(c *gin.Context) {

	lines, err := h.services.Users.GetSubwayLines()
	if err != nil {
		apiErrorEncode(c, err)
		return
	}

	c.JSON(http.StatusOK, lines)
}
