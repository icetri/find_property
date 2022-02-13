package domain

import (
	"github.com/lib/pq"
	"time"
)

type UserSignUpInput struct {
	UserID int `json:"user_id" gorm:"column:user_id" binding:"required"`
}

type User struct {
	UserID int `json:"user_id" gorm:"column:user_id"`
}

type PagList struct {
	Items interface{} `json:"items"`
	Count int64       `json:"count"`
}

type View struct {
	ID     string `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	UserID int    `json:"user_id"`
	CianID int    `json:"cian_id"`
}

type ReverseJK struct {
	JKID   int `json:"jk_id" binding:"required`
	UserID int `json:"user_id" binding:"required`
}

type SubwayLines struct {
	LineName string `json:"line_name" gorm:"column:line_name"`
	LineID   string `json:"line_id" gorm:"column:line_id"`
}

type StartSearch struct {
	Limit  int `json:"limit" example:"1" binding:"required"`
	Offset int `json:"offset" example:"0"`
	UserID int `json:"user_id" binding:"required"`
}

type Adverts []Advert
type Advert struct {
	ID            int    `json:"id" gorm:"primaryKey;column:id"`
	Subcategory   string `json:"subcategory" gorm:"column:subcategory"`
	Images        string `json:"images" gorm:"column:images"`
	JkDisplayname string `json:"jk_name" gorm:"column:jk_name"`
	Address       string `json:"address" gorm:"column:address"`
	Metro         string `json:"metro" gorm:"column:metro"`
	RoomsCount    string `json:"rooms_count" gorm:"column:rooms_count"`
	Price         int    `json:"price" gorm:"column:price"`
	Area          string `json:"area" gorm:"column:area"`
	FloorNumber   int    `json:"floor_number" gorm:"column:floor_number"`
	URL           string `json:"url" gorm:"column:url"`
	JkFullurl     string `json:"jk_fullUrl" gorm:"column:jk_fullurl"`
	Decoration    string `json:"decoration" gorm:"column:decoration"`
	Websiteurlutm string `json:"webSiteUrlUtm" gorm:"column:web_site_url_utm"`
	Lat           string `json:"lat" gorm:"column:lat"`
	Lng           string `json:"lng" gorm:"column:lng"`
	DeveloperName string `json:"developer_name" gorm:"column:developer_name"`
	Deadlineinfo  string `json:"deadlineInfo" gorm:"column:deadline_info"`
	JKID          int    `json:"jk_id" gorm:"column:jk_id"`

	Underground []Metro `json:"underground" gorm:"-"`
}

type CianParse struct {
	ID              int    `json:"id" gorm:"primaryKey;column:id"`
	Subcategory     string `json:"subcategory" gorm:"column:subcategory"`
	RoomsCount      string `json:"rooms_count" gorm:"column:rooms_count"`
	BeforeMetro     bool   `json:"before_metro" gorm:"column:before_metro"`
	URL             string `json:"url" gorm:"column:url"`
	Price           int    `json:"price" gorm:"column:price"`
	TimePublish     string `json:"time_publish" gorm:"column:time_publish"`
	Phone           string `json:"phone" gorm:"column:phone"`
	Address         string `json:"address" gorm:"column:address"`
	Metro           string `json:"metro" gorm:"column:metro"`
	OfferType       string `json:"offerType" gorm:"column:offer_type"`
	FloorNumber     int    `json:"floor_number" gorm:"column:floor_number"`
	FloorsCount     int    `json:"floors_count" gorm:"column:floors_count"`
	Area            string `json:"area" gorm:"column:area"`
	BuildingYear    int    `json:"building_year" gorm:"column:building_year"`
	IsApartments    bool   `json:"isApartments" gorm:"column:is_apartments"`
	DealType        string `json:"deal_type" gorm:"column:deal_type"`
	RepairType      string `json:"repair_type" gorm:"column:repair_type"`
	Images          string `json:"images" gorm:"column:images"`
	Balconies       bool   `json:"balconies" gorm:"column:balconies"`
	Loggias         bool   `json:"loggias" gorm:"column:loggias"`
	JkFullurl       string `json:"jk_fullUrl" gorm:"column:jk_fullurl"`
	JkDisplayname   string `json:"jk_name" gorm:"column:jk_name"`
	JKID            int    `json:"jk_id" gorm:"column:jk_id"`
	Decoration      string `json:"decoration" gorm:"column:decoration"`
	Websiteurlutm   string `json:"webSiteUrlUtm" gorm:"column:web_site_url_utm"`
	Lat             string `json:"lat" gorm:"column:lat"`
	Lng             string `json:"lng" gorm:"column:lng"`
	DeveloperName   string `json:"developer_name" gorm:"column:developer_name"`
	Deadlineinfo    string `json:"deadlineInfo" gorm:"column:deadline_info"`
	PublishedUserID int    `json:"published_user_id" gorm:"column:published_user_id"`
	CianUserID      int    `json:"cian_user_id" gorm:"column:cian_user_id"`
}

type Metro struct {
	ID            string    `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	CreatedAt     time.Time `json:"created_at" gorm:"column:created_at"`
	Metro         string    `json:"metro" gorm:"column:metro"`
	TransportType string    `json:"transport_type" gorm:"column:transport_type"`
	LineColor     string    `json:"line_color" gorm:"column:line_color"`
	Time          int       `json:"time" gorm:"column:time"`
	LineID        string    `json:"line_id" gorm:"column:line_id"`
	IsDefault     bool      `json:"is_default" gorm:"column:is_default"`
	CianID        int       `json:"cian_id" gorm:"column:cian_id"`
}

type Filter struct {
	ID            string         `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Subcategory   string         `json:"subcategory" gorm:"column:subcategory;default:Новостройка"`                                                                                                                //example Новостройка, Вторичка
	RoomsCount    pq.StringArray `json:"rooms_count" gorm:"type:text[];column:rooms_count;default:1-комн (евродва)" swaggertype:"array,string"`                                                                    //example [1-комн (евродва),2-комн (евротри),3-комн,4+комн,студия]
	Price1        int            `json:"price1" gorm:"column:price1;default:0"`                                                                                                                                    //example 6459462
	Price2        int            `json:"price2" gorm:"column:price2;default:10000000000"`                                                                                                                          //example 7194195
	Area1         string         `json:"area1" gorm:"column:area1;default:0"`                                                                                                                                      //example 55.9
	Area2         string         `json:"area2" gorm:"column:area2;default:100000000"`                                                                                                                              //example 76.0
	Metro         pq.StringArray `json:"metro" gorm:"type:text[];column:metro;default:1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30 31 32 33 34 35" swaggertype:"array,string"` //example Речной вокзал
	BeforeMetro   string         `json:"before_metro" gorm:"column:before_metro;default:Не важно"`                                                                                                                 //example неважно, только пешком
	RepairType    string         `json:"repair_type" gorm:"column:repair_type;default:Не важно"`                                                                                                                   //example euro
	BuildingYear1 int            `json:"building_year1" gorm:"column:building_year1;default:2010"`                                                                                                                 //example 2010
	BuildingYear2 int            `json:"building_year2" gorm:"column:building_year2;default:2025"`                                                                                                                 //example 2024
	JKName        string         `json:"jk_name" gorm:"column:jk_name;default:Не важно"`                                                                                                                           //example ЖК ... либо пустое
	Address       string         `json:"address" gorm:"column:address;default:Не важно"`                                                                                                                           //example Адрес ... либо пустое
	Category      string         `json:"category" gorm:"column:category;default:Квартиры и апартаменты"`                                                                                                           //example только квартиры, квартиры и апартаменты, только апартаменты
	Floors        string         `json:"floors" gorm:"column:floors;default:Не важно"`                                                                                                                             //example только не первый
	TypeBalcony   string         `json:"type_balcony" gorm:"column:type_balcony;default:Не важно"`                                                                                                                 //example неважно,балкон, лоджия

	JKIDs   pq.Int64Array `json:"jk_ids" gorm:"type:text[];column:jk_ids" swaggertype:"array,string"`
	Default string        `json:"default" gorm:"column:default"`
	Silence string        `json:"silence" gorm:"column:silence"`
	UserID  int           `json:"user_id" gorm:"column:user_id"`
}

type Cian struct {
	Status string `json:"status"`
	Data   struct {
		Offersserialized []struct {
			Newbuilding struct {
				Showjkreliableflag        bool        `json:"showJkReliableFlag"`
				Isfromdeveloper           bool        `json:"isFromDeveloper"`
				Isreliable                bool        `json:"isReliable"`
				Name                      string      `json:"name"`
				Isbuilderreliableinregion interface{} `json:"isBuilderReliableInRegion"`
				Isfrombuilder             bool        `json:"isFromBuilder"`
				House                     struct {
					Isreliable bool `json:"isReliable"`
					Isfinished bool `json:"isFinished"`
					Finishdate struct {
						Quarter int `json:"quarter"`
						Year    int `json:"year"`
					} `json:"finishDate"`
					ID   int    `json:"id"`
					Name string `json:"name"`
				} `json:"house"`
				Newbuildingfeatures struct {
					Reviewscount  interface{} `json:"reviewsCount"`
					Firstimageurl string      `json:"firstImageUrl"`
					Imagescount   int         `json:"imagesCount"`
					Deadlineinfo  string      `json:"deadlineInfo"`
					Videoscount   int         `json:"videosCount"`
				} `json:"newbuildingFeatures"`
				Isfromseller bool   `json:"isFromSeller"`
				ID           string `json:"id"`
			} `json:"newbuilding"`
			Roomtype        string      `json:"roomType"`
			Isapartments    bool        `json:"isApartments"`
			Title           interface{} `json:"title"`
			Decoration      string      `json:"decoration"`
			Addedtimestamp  int64       `json:"addedTimestamp"`
			Publisheduserid int         `json:"publishedUserId"`
			Offertype       string      `json:"offerType"`
			Geo             struct {
				Highways  []interface{} `json:"highways"`
				Districts []struct {
					Geotype    string `json:"geoType"`
					Qs         string `json:"qs"`
					Locationid int    `json:"locationId"`
					Name       string `json:"name"`
					Title      string `json:"title"`
					Shortname  string `json:"shortName"`
					Parentid   int    `json:"parentId"`
					Type       string `json:"type"`
					Fullname   string `json:"fullName"`
					ID         int    `json:"id"`
				} `json:"districts"`
				Userinput    string `json:"userInput"`
				Countryid    int    `json:"countryId"`
				Undergrounds []struct {
					Geotype           string      `json:"geoType"`
					Qs                string      `json:"qs"`
					Name              string      `json:"name"`
					Title             string      `json:"title"`
					Cianid            interface{} `json:"cianId"`
					Releaseyear       interface{} `json:"releaseYear"`
					Transporttype     string      `json:"transportType"`
					ID                int         `json:"id"`
					Underconstruction bool        `json:"underConstruction"`
					Linecolor         string      `json:"lineColor"`
					Time              int         `json:"time"`
					Shortname         string      `json:"shortName"`
					Lineid            int         `json:"lineId"`
					Type              string      `json:"type"`
					Fullname          string      `json:"fullName"`
					Isdefault         bool        `json:"isDefault"`
				} `json:"undergrounds"`
				Jk struct {
					Websiteurl  string `json:"webSiteUrl"`
					Fullurl     string `json:"fullUrl"`
					Displayname string `json:"displayName"`
					Name        string `json:"name"`
					Gageo       struct {
						Cityid int `json:"cityId"`
						Moid   int `json:"moId"`
						Oblid  int `json:"oblId"`
					} `json:"gaGeo"`
					House struct {
						Flat interface{} `json:"flat"`
						Name string      `json:"name"`
						ID   int         `json:"id"`
					} `json:"house"`
					Developer struct {
						Name string      `json:"name"`
						ID   interface{} `json:"id"`
					} `json:"developer"`
					ID            int    `json:"id"`
					Websiteurlutm string `json:"webSiteUrlUtm"`
				} `json:"jk"`
				Address []struct {
					Geotype          string `json:"geoType"`
					Qs               string `json:"qs"`
					Name             string `json:"name"`
					Title            string `json:"title"`
					Locationtypeid   int    `json:"locationTypeId"`
					Shortname        string `json:"shortName"`
					Type             string `json:"type"`
					Fullname         string `json:"fullName"`
					Isformingaddress bool   `json:"isFormingAddress"`
					ID               int    `json:"id"`
				} `json:"address"`
				Buildingaddress interface{} `json:"buildingAddress"`
				Locationpath    struct {
					Childtoparent []int `json:"childToParent"`
					Countryid     int   `json:"countryId"`
				} `json:"locationPath"`
				Coordinates struct {
					Lat float64 `json:"lat"`
					Lng float64 `json:"lng"`
				} `json:"coordinates"`
			} `json:"geo"`
			Fullurl     string      `json:"fullUrl"`
			Objecttype  int         `json:"objectType"`
			Cianid      int         `json:"cianId"`
			Repairtype  string      `json:"repairType"`
			ID          int         `json:"id"`
			Userid      int         `json:"userId"`
			Totalarea   string      `json:"totalArea"`
			Name        interface{} `json:"name"`
			Floornumber int         `json:"floorNumber"`
			Category    string      `json:"category"`
			Roomarea    interface{} `json:"roomArea"`
			Flattype    string      `json:"flatType"`
			Roomscount  int         `json:"roomsCount"`
			Phones      []struct {
				Number      string `json:"number"`
				Countrycode string `json:"countryCode"`
			} `json:"phones"`
			Loggiascount int `json:"loggiasCount"`
			Bargainterms struct {
				Paymentperiod   interface{} `json:"paymentPeriod"`
				Mortgageallowed interface{} `json:"mortgageAllowed"`
				Vatprices       struct {
					Rur interface{} `json:"rur"`
					Usd interface{} `json:"usd"`
					Eur interface{} `json:"eur"`
				} `json:"vatPrices"`
				Agentfee              interface{}   `json:"agentFee"`
				Includedoptions       []interface{} `json:"includedOptions"`
				Securitydepositprices struct {
					Rur interface{} `json:"rur"`
					Usd interface{} `json:"usd"`
					Eur interface{} `json:"eur"`
				} `json:"securityDepositPrices"`
				Price              float64 `json:"price"`
				Bargainpriceprices struct {
					Rur interface{} `json:"rur"`
					Usd interface{} `json:"usd"`
					Eur interface{} `json:"eur"`
				} `json:"bargainPricePrices"`
				Pricerur        int         `json:"priceRur"`
				Priceeur        int         `json:"priceEur"`
				Currency        string      `json:"currency"`
				Tenantstype     interface{} `json:"tenantsType"`
				Securitydeposit interface{} `json:"securityDeposit"`
				Depositprices   struct {
					Rur interface{} `json:"rur"`
					Usd interface{} `json:"usd"`
					Eur interface{} `json:"eur"`
				} `json:"depositPrices"`
				Contracttype         interface{} `json:"contractType"`
				Vatincluded          interface{} `json:"vatIncluded"`
				Utilitiestermsprices interface{} `json:"utilitiesTermsPrices"`
				Vattype              interface{} `json:"vatType"`
				Hasgraceperiod       interface{} `json:"hasGracePeriod"`
				Clientfee            interface{} `json:"clientFee"`
				Vatprice             interface{} `json:"vatPrice"`
				Saletype             string      `json:"saleType"`
				Utilitiesterms       interface{} `json:"utilitiesTerms"`
				Minleaseterm         interface{} `json:"minLeaseTerm"`
				Priceusd             int         `json:"priceUsd"`
				Leasetermtype        interface{} `json:"leaseTermType"`
				Leasetype            interface{} `json:"leaseType"`
				Deposit              interface{} `json:"deposit"`
				Bargainprice         interface{} `json:"bargainPrice"`
				Bargainallowed       interface{} `json:"bargainAllowed"`
				Bargainconditions    interface{} `json:"bargainConditions"`
				Agentbonus           interface{} `json:"agentBonus"`
				Prepaymonths         interface{} `json:"prepayMonths"`
				Pricetype            string      `json:"priceType"`
			} `json:"bargainTerms"`
			Dealtype       string      `json:"dealType"`
			Cianuserid     int         `json:"cianUserId"`
			Jkurl          string      `json:"jkUrl"`
			Conditiontype  interface{} `json:"conditionType"`
			Balconiescount int         `json:"balconiesCount"`
			Building       struct {
				Cargoliftscount interface{} `json:"cargoLiftsCount"`
				Heatingtype     interface{} `json:"heatingType"`
				Type            interface{} `json:"type"`
				Deadline        struct {
					Quarterend string `json:"quarterEnd"`
					Quarter    string `json:"quarter"`
					Year       int    `json:"year"`
					Iscomplete bool   `json:"isComplete"`
				} `json:"deadline"`
				Materialtype        string      `json:"materialType"`
				Passengerliftscount interface{} `json:"passengerLiftsCount"`
				Floorscount         int         `json:"floorsCount"`
				Buildyear           int         `json:"buildYear"`
				Accesstype          interface{} `json:"accessType"`
			} `json:"building"`
			Land       interface{} `json:"land"`
			Livingarea string      `json:"livingArea"`
			Photos     []struct {
				Source        interface{} `json:"source"`
				Fullurl       string      `json:"fullUrl"`
				Rotatedegree  interface{} `json:"rotateDegree"`
				Isdefault     bool        `json:"isDefault"`
				Thumbnailurl  string      `json:"thumbnailUrl"`
				Thumbnail2URL string      `json:"thumbnail2Url"`
				ID            int         `json:"id"`
				Coordinates   interface{} `json:"coordinates"`
				Miniurl       string      `json:"miniUrl"`
			} `json:"photos"`
			User struct {
				Accounttype    string `json:"accountType"`
				Agencyname     string `json:"agencyName"`
				Agentavatarurl string `json:"agentAvatarUrl"`
				Billinginfo    struct {
					Accounttype int `json:"accountType"`
					Packages    []struct {
						Countryid       int    `json:"countryId"`
						Depositrequired bool   `json:"depositRequired"`
						Geoname         string `json:"geoName"`
						Locationid      int    `json:"locationId"`
						Packagetypeid   int    `json:"packageTypeId"`
						Packagetypename string `json:"packageTypeName"`
					} `json:"packages"`
					Removecompetitor  bool `json:"removeCompetitor"`
					Replenishdisabled bool `json:"replenishDisabled"`
				} `json:"billingInfo"`
				Cianprofilestatus  string `json:"cianProfileStatus"`
				Cianuserid         string `json:"cianUserId"`
				Companyname        string `json:"companyName"`
				Isagent            bool   `json:"isAgent"`
				Isbuilder          bool   `json:"isBuilder"`
				Iscallbackuser     bool   `json:"isCallbackUser"`
				Ischatsenabled     bool   `json:"isChatsEnabled"`
				Ishidden           bool   `json:"isHidden"`
				Ispassportverified bool   `json:"isPassportVerified"`
				Isselfemployed     bool   `json:"isSelfEmployed"`
				Issubagent         bool   `json:"isSubAgent"`
				Phonenumbers       []struct {
					Countrycode string `json:"countryCode"`
					Number      string `json:"number"`
				} `json:"phoneNumbers"`
				Usertrustlevel      string `json:"userTrustLevel"`
				Userid              int    `json:"userId"`
				Usertype            string `json:"userType"`
				Iscianpartner       bool   `json:"isCianPartner"`
				Canshowonline       bool   `json:"canShowOnline"`
				Agentmoderationinfo struct {
					Isuseridentified            bool `json:"isUserIdentified"`
					Isuseridentifiedbydocuments bool `json:"isUserIdentifiedByDocuments"`
				} `json:"agentModerationInfo"`
				Personalrating    interface{} `json:"personalRating"`
				Agentavailability struct {
					Available bool `json:"available"`
					Userid    int  `json:"userId"`
				} `json:"agentAvailability"`
			} `json:"user"`
			Ispro                          bool          `json:"isPro"`
			Galabel                        string        `json:"gaLabel"`
			Chatid                         interface{}   `json:"chatId"`
			Chat                           interface{}   `json:"chat"`
			Auction                        interface{}   `json:"auction"`
			Savedsearchlabels              []interface{} `json:"savedSearchLabels"`
			Isnew                          bool          `json:"isNew"`
			Isdealrequestsubstitutionphone bool          `json:"isDealRequestSubstitutionPhone"`
		} `json:"offersSerialized"`
	} `json:"data"`
}
