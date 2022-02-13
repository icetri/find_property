package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/find_property/internal/domain"
	"github.com/find_property/internal/repository"
	"github.com/find_property/pkg/logger"
	"github.com/find_property/pkg/worker"

	"github.com/lib/pq"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type UsersService struct {
	repo repository.Users
}

func NewUsersService(repo repository.Users, worker *worker.Worker) *UsersService {

	userService := new(UsersService)
	userService.repo = repo

	err := worker.Every(1).Hour().Do(func() {
		if err := userService.parse(); err != nil {
			logger.LogError(err)
		}
	})
	if err != nil {
		logger.LogError(errors.Wrap(err, "err with parse in worker"))
	}

	err = worker.Every(1).Day().At("07:00:00").Do(func() {
		if err := userService.deleteBadAdverts(); err != nil {
			logger.LogError(err)
		}
	})
	if err != nil {
		logger.LogError(errors.Wrap(err, "err with deleteBadAdverts in worker"))
	}

	return userService
}

func (s *UsersService) CheckUser(check *domain.UserSignUpInput) (bool, error) {

	user, err := s.repo.GetUser(check)
	if err != nil && err != gorm.ErrRecordNotFound {
		logger.LogError(errors.Wrap(err, "err with GetUser in CheckUser"))
		return false, domain.ErrorInternalServerError
	}

	isRegister := user != nil

	return isRegister, nil
}

func (s *UsersService) CreateFilter(user *domain.Filter) error {

	if err := s.repo.Create(user); err != nil {
		logger.LogError(errors.Wrap(err, "err with Create in CreateFilter"))
		return domain.ErrorInternalServerError
	}

	return nil
}

func (s *UsersService) CreateView(view *domain.View) error {

	if err := s.repo.AddView(view); err != nil {
		logger.LogError(errors.Wrap(err, "err with AddView in CreateView"))
		return domain.ErrorInternalServerError
	}

	return nil
}

func (s *UsersService) ReverseSweepJK(jk *domain.ReverseJK) error {

	user, err := s.repo.GetFilterUser(strconv.Itoa(jk.UserID))
	if err != nil && err != gorm.ErrRecordNotFound {
		logger.LogError(errors.Wrap(err, "err with GetFilterUser in ReverseSweepJK"))
		return domain.ErrorInternalServerError
	}

	if user == nil {
		return domain.ErrorUserNotExist
	}

	switch len(user.JKIDs) {
	case 5:
		user.JKIDs = s.remove(user.JKIDs, user.JKIDs[0])
		user.JKIDs = append(user.JKIDs, int64(jk.JKID))
	default:
		user.JKIDs = append(user.JKIDs, int64(jk.JKID))
	}

	if err := s.repo.UpdateFilterLines(jk.UserID, user); err != nil {
		logger.LogError(errors.Wrap(err, "err with UpdateFilterLines in ReverseSweepJK"))
		return domain.ErrorInternalServerError
	}

	return nil
}

func (s *UsersService) SubwayLines(query string) ([]domain.SubwayLines, error) {

	query = "%" + query + "%"
	lines, err := s.repo.GetSubwayLines(query)
	if err != nil {
		logger.LogError(errors.Wrap(err, "err with GetSubwayLines in SubwayLines"))
		return nil, domain.ErrorInternalServerError
	}

	return lines, nil
}

func (s *UsersService) GetSubwayLines() ([]domain.SubwayLines, error) {

	lines, err := s.repo.GetSubwaysLines()
	if err != nil {
		logger.LogError(errors.Wrap(err, "err with GetSubwaysLines in GetSubwayLines"))
		return nil, domain.ErrorInternalServerError
	}

	return lines, nil
}

func (s *UsersService) CreateUser(user *domain.Filter) (*domain.Filter, error) {

	user.JKIDs = pq.Int64Array{}

	if err := s.repo.CreateUser(user); err != nil {
		if err == domain.ErrorUserIsExist {
			return nil, domain.ErrorUserIsExist
		}
		logger.LogError(errors.Wrap(err, "err with CreateUser in CreateUser"))
		return nil, domain.ErrorInternalServerError
	}

	return user, nil
}

func (s *UsersService) GetUser(userID string) (*domain.Filter, error) {

	filter, err := s.repo.GetFilterUser(userID)
	if err != nil && err != gorm.ErrRecordNotFound {
		logger.LogError(errors.Wrap(err, "err with GetFilterUser in GetUser"))
		return nil, domain.ErrorInternalServerError
	}

	if filter == nil {
		return nil, domain.ErrorUserNotExist
	}

	return filter, nil
}

func (s *UsersService) parse() error {

	arrayProxy, err := s.proxy()
	if err != nil {
		return err
	}

	i := 1
	for in := range arrayProxy {

		for ; i < 32; i++ {
			urlStr := "https://api.cian.ru/search-offers/v1/search-offers-desktop/"

			var jsonStr = []byte(fmt.Sprintf(`{
		"jsonQuery":{
			"sort":{"type":"term","value":"creation_date_desc"},
			"_type":"flatsale",
			"region": {"type": "terms", "value": [-1]},
			"engine_version":{"type":"term","value":2},
			"house_year":{"type":"range","value":{"gte":2010}},
			"page":{"type":"term","value":%d}
}
}`, i))

			req, err := http.NewRequest("POST", urlStr, bytes.NewBuffer(jsonStr))
			if err != nil {
				logger.LogError(errors.Wrap(err, "err with NewRequest in parse"))
				return domain.ErrorInternalServerError
			}

			req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.212 Safari/537.36")
			req.Header.Set("Content-Type", "application/json")

			proxyUrl := fmt.Sprintf("http://%s", arrayProxy[in])
			proxyUrla, err := url.Parse(proxyUrl)
			if err != nil {
				break
			}

			transport := &http.Transport{
				Proxy: http.ProxyURL(proxyUrla),
			}

			client := &http.Client{
				Transport: transport,
				Timeout:   45 * time.Second,
			}
			resp, err := client.Do(req)
			if err != nil {
				break
			}
			defer resp.Body.Close()

			cian := new(domain.Cian)
			if err = json.NewDecoder(resp.Body).Decode(cian); err != nil {
				break
			}

			cianParses := s.parseCollector(cian)

			if len(cianParses) != 28 {
				return nil
			}

			time.Sleep(45 * time.Second)
		}
	}

	return nil
}

func (s *UsersService) parseCollector(cian *domain.Cian) []domain.CianParse {

	cianParses := make([]domain.CianParse, 0)
	for _, house := range cian.Data.Offersserialized {
		cianParse := new(domain.CianParse)

		cianParse.ID = house.ID

		switch house.Objecttype {
		case 1:
			cianParse.Subcategory = "Вторичка"
		case 2:
			cianParse.Subcategory = "Новостройка"
		}

		switch house.Roomscount {
		case 0:
			cianParse.RoomsCount = "студия"
		case 1:
			cianParse.RoomsCount = "1-комн (евродва)"
		case 2:
			cianParse.RoomsCount = "2-комн (евротри)"
		case 3:
			cianParse.RoomsCount = "3-комн"
		default:
			cianParse.RoomsCount = "4+комн"
		}

		cianParse.Price = house.Bargainterms.Pricerur

		cianParse.Area = house.Totalarea

		if len(house.Geo.Address) != 0 {
			cianParse.Address = house.Geo.Address[0].Name
		}

		for _, addres := range house.Geo.Address {
			if cianParse.Address == addres.Name {
				continue
			}
			cianParse.Address = cianParse.Address + " " + addres.Name
		}

		if len(house.Geo.Undergrounds) != 0 {
			cianParse.Metro = house.Geo.Undergrounds[0].Name
		}

		for _, metro := range house.Geo.Undergrounds {
			if cianParse.Metro == metro.Name {
				continue
			}

			cianParse.Metro = cianParse.Metro + ", " + metro.Name
		}

		for _, metro := range house.Geo.Undergrounds {
			if metro.Transporttype == "walk" {
				cianParse.BeforeMetro = true
				break
			}
		}

		cianParse.RepairType = house.Repairtype

		if cianParse.Subcategory == "Новостройка" {
			cianParse.BuildingYear = house.Building.Deadline.Year
		} else {
			cianParse.BuildingYear = house.Building.Buildyear
		}

		cianParse.IsApartments = house.Isapartments

		cianParse.FloorNumber = house.Floornumber

		cianParse.FloorsCount = house.Building.Floorscount

		cianParse.Balconies = house.Balconiescount != 0

		cianParse.Loggias = house.Loggiascount != 0

		cianParse.URL = house.Fullurl

		cianParse.OfferType = house.Offertype

		if len(house.Phones) != 0 {
			cianParse.Phone = house.Phones[0].Countrycode + house.Phones[0].Number
		}

		if len(house.Photos) != 0 {
			cianParse.Images = house.Photos[0].Fullurl
		}

		for _, photo := range house.Photos {
			if cianParse.Images == photo.Fullurl {
				continue
			}
			cianParse.Images = cianParse.Images + "," + photo.Fullurl
		}

		cianParse.JkFullurl = house.Geo.Jk.Fullurl

		cianParse.JkDisplayname = house.Geo.Jk.Displayname

		cianParse.JKID = house.Geo.Jk.ID

		cianParse.PublishedUserID = house.Publisheduserid

		cianParse.CianUserID = house.Cianuserid

		cianParse.DealType = house.Dealtype

		unixTimeUTC := time.Unix(house.Addedtimestamp, 0)

		cianParse.TimePublish = unixTimeUTC.Format(time.RFC3339)

		switch house.Decoration {
		case "fine":
			cianParse.Decoration = "Чистовая"
		case "without":
			cianParse.Decoration = "Нет"
		case "rough":
			cianParse.Decoration = "Черновая"
		}

		cianParse.Deadlineinfo = house.Newbuilding.Newbuildingfeatures.Deadlineinfo

		cianParse.Websiteurlutm = house.Geo.Jk.Websiteurlutm

		cianParse.Lat = fmt.Sprintf("%f", house.Geo.Coordinates.Lat)
		cianParse.Lng = fmt.Sprintf("%f", house.Geo.Coordinates.Lng)

		cianParse.DeveloperName = house.Geo.Jk.Developer.Name

		if err := s.repo.Parse(cianParse); err != nil {
			break
		}

		for _, metro := range house.Geo.Undergrounds {
			if err := s.repo.CreateMetro(&domain.Metro{
				Metro:         metro.Name,
				TransportType: metro.Transporttype,
				LineColor:     metro.Linecolor,
				Time:          metro.Time,
				LineID:        strconv.Itoa(metro.Lineid),
				IsDefault:     metro.Isdefault,
				CianID:        house.ID,
			}); err != nil {
				logger.LogError(errors.Wrap(err, "err with CreateMetro in Parse"))
				break
			}
		}

		cianParses = append(cianParses, *cianParse)
	}

	return cianParses
}

func (s *UsersService) remove(list pq.Int64Array, item int64) pq.Int64Array {
	for i, v := range list {
		if v == item {
			copy(list[i:], list[i+1:])
			list[len(list)-1] = 0
			list = list[:len(list)-1]
		}
	}

	return list
}

func (s *UsersService) deleteBadAdverts() error {

	adverts, err := s.repo.AllAdverts()
	if err != nil {
		logger.LogError(errors.Wrap(err, "err with AllAdverts in deleteBadAdverts"))
		return domain.ErrorInternalServerError
	}

	arrayProxy, err := s.proxy()
	if err != nil {
		logger.LogError(errors.Wrap(err, "err with proxy in deleteBadAdverts"))
		return domain.ErrorInternalServerError
	}

	i := 0
	for in := range arrayProxy {
		for ; i < len(adverts); i++ {
			req, err := http.NewRequest("GET", adverts[i].URL, nil)
			if err != nil {
				logger.LogError(errors.Wrap(err, "err with NewRequest in deleteBadAdverts"))
				return domain.ErrorInternalServerError
			}

			proxyUrl := fmt.Sprintf("http://%s", arrayProxy[in])
			proxyUrla, err := url.Parse(proxyUrl)
			if err != nil {
				break
			}

			transport := &http.Transport{
				Proxy: http.ProxyURL(proxyUrla),
			}

			client := &http.Client{
				Transport: transport,
				Timeout:   15 * time.Second,
			}

			resp, err := client.Do(req)
			if err != nil {
				break
			}
			defer resp.Body.Close()

			doc, err := goquery.NewDocumentFromReader(resp.Body)
			if err != nil {
				return nil
			}

			var check bool

			doc.Find("title").Each(func(i int, s *goquery.Selection) {
				file := s.Text()
				if file == "Captcha - база объявлений ЦИАН" || file == "Ошибка - Циан" {
					check = true
				}
			})

			if check == true {
				break
			}

			var advertInfo string
			doc.Find(".a10a3f92e9--container--1In69").Each(func(i int, s *goquery.Selection) {
				advertInfo = s.Text()
			})

			if advertInfo == "Объявление снято с публикации" {
				if err := s.repo.DeleteMetros(adverts[i].ID); err != nil {
					logger.LogError(errors.Wrap(err, "err with DeleteMetros in deleteBadAdverts"))
					return domain.ErrorInternalServerError
				}

				if err := s.repo.DeleteAdvert(adverts[i].ID); err != nil {
					logger.LogError(errors.Wrap(err, "err with DeleteAdvert in deleteBadAdverts"))
					return domain.ErrorInternalServerError
				}
			}

			time.Sleep(3 * time.Second)
		}
	}

	return nil
}

func (s *UsersService) proxy() ([]string, error) {

	urlProxy := "https://api.proxyscrape.com/v2/?request=displayproxies&protocol=http&timeout=10000&country=all&ssl=all&anonymity=all"

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, urlProxy, nil)

	if err != nil {
		logger.LogError(errors.Wrap(err, "err with NewRequest in proxy"))
		return nil, domain.ErrorInternalServerError
	}

	res, err := client.Do(req)
	if err != nil {
		logger.LogError(errors.Wrap(err, "err with Do in proxy"))
		return nil, domain.ErrorInternalServerError
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		logger.LogError(errors.Wrap(err, "err with ReadAll in proxy"))
		return nil, domain.ErrorInternalServerError
	}

	arrayW := string(body)
	addresses := strings.Split(arrayW, "\r\n")

	return addresses, nil
}
