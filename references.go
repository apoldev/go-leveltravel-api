package leveltravel

import "strings"

type Departures struct {
	Success    bool        `json:"success"`
	Departures []Departure `json:"departures"`
}

type Departure struct {
	Id          int                `json:"id"`
	NameRu      string             `json:"name_ru"`
	NameRuForm1 string             `json:"name_ru_form1"`
	NameRuForm2 string             `json:"name_ru_form2"`
	NameEn      string             `json:"name_en"`
	Iata        string             `json:"iata"`
	Priority    int                `json:"priority"`
	Iso2        string             `json:"iso2"`
	Airports    []DepartureAirport `json:"airports"`
}

type DepartureAirport struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Code     string `json:"code"`
	Lat      string `json:"lat"`
	Long     string `json:"long"`
	Timezone string `json:"timezone"`
}

func (l *LevelTravelApi) Departures() (departures Departures, err error) {
	err = l.getJson("/references/departures", map[string]string{}, &departures)
	return
}

type Destinations struct {
	Success   bool                 `json:"success"`
	Countries []DestinationCountry `json:"countries"`
}

type DestinationCountry struct {
	Id       int               `json:"id"`
	NameRu   string            `json:"name_ru"`
	NameEn   string            `json:"name_en"`
	Iso2     string            `json:"iso2"`
	Priority int               `json:"priority"`
	Cities   []DestinationCity `json:"cities"`
}

type DestinationCity struct {
	Id     int    `json:"id"`
	NameRu string `json:"name_ru"`
	NameEn string `json:"name_en"`
}

func (l *LevelTravelApi) Destinations() (destinations Destinations, err error) {
	err = l.getJson("/references/destinations", map[string]string{}, &destinations)
	return
}

type Operators struct {
	Success   bool       `json:"success"`
	Operators []Operator `json:"operators"`
}

type Operator struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	ShortName string `json:"short_name"`
	CodeName  string `json:"code_name"`
}

func (l *LevelTravelApi) Operators() (operators Operators, err error) {
	err = l.getJson("/references/operators", map[string]string{}, &operators)
	return
}

type Airlines struct {
	Success  bool      `json:"success"`
	Airlines []Airline `json:"airlines"`
}

type Airline struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
}

func (l *LevelTravelApi) Airlines() (airlines Airlines, err error) {
	err = l.getJson("/references/airlines", map[string]string{}, &airlines)
	return
}

type Airports struct {
	Success bool `json:"success"`
}

func (l *LevelTravelApi) Airports() (airports Airports, err error) {
	err = l.getJson("/references/airports", map[string]string{}, &airports)
	return
}

type Hotels struct {
	Success bool    `json:"success"`
	Hotels  []Hotel `json:"hotels"`
}

type Hotel struct {
	Id             int            `json:"id"`
	Name           string         `json:"name"`
	Stars          int            `json:"stars"`
	SiteUrl        string         `json:"site_url"`
	PublicUrl      string         `json:"public_url"`
	Features       HotelFeatures  `json:"features"`
	ShortInfo      string         `json:"short_info"`
	HotelInfo      string         `json:"hotel_info"`
	KidsFacilities []string       `json:"kids_facilities"`
	Territory      []string       `json:"territory"`
	Attractions    []string       `json:"attractions"`
	RoomFacilities []string       `json:"room_facilities"`
	Lat            float64        `json:"lat"`
	Long           float64        `json:"long"`
	City           HotelCity      `json:"city"`
	Country        HotelCountry   `json:"country"`
	Rating         float64        `json:"rating"`
	Images         []HotelImage   `json:"images"`
	RestTypes      HotelRestTypes `json:"rest_types"`
}

type HotelFeatures struct {
	ConstructionYear int    `json:"construction_year"`
	RenovationYear   int    `json:"renovation_year"`
	AirportDistance  int    `json:"airport_distance"`
	BeachDistance    int    `json:"beach_distance"`
	BeachSize        int    `json:"beach_size"`
	BeachType        string `json:"beach_type"`
	BeachSurface     string `json:"beach_surface"`
	Area             int    `json:"area"`
	Line             int    `json:"line"`
	WiFi             string `json:"wi_fi"`
	SkiLiftDistance  int    `json:"ski_lift_distance"`
	SkiIn            string `json:"ski_in"`
	SkiOut           string `json:"ski_out"`
	Fitness          bool   `json:"fitness"`
	Aquapark         bool   `json:"aquapark"`
	Nanny            bool   `json:"nanny"`
	KidsMenu         bool   `json:"kids_menu"`
	WarmPool         bool   `json:"warm_pool"`
	KidsClub         bool   `json:"kids_club"`
	Aircon           bool   `json:"aircon"`
}

type HotelCity struct {
	Id     int    `json:"id"`
	NameEn string `json:"name_en"`
	NameRu string `json:"name_ru"`
	Iata   string `json:"iata"`
}

type HotelCountry struct {
	Id     int    `json:"id"`
	NameEn string `json:"name_en"`
	NameRu string `json:"name_ru"`
	Iso2   string `json:"iso2"`
}

type HotelImage struct {
	Id       int    `json:"id"`
	X90      string `json:"x90"`
	X180     string `json:"x180"`
	X180x132 string `json:"x180x132"`
	X250     string `json:"x250"`
	X500     string `json:"x500"`
	X900     string `json:"x900"`
}

type HotelRestTypes struct {
	Family   int `json:"family"`
	Kids     int `json:"kids"`
	Business int `json:"business"`
	Couples  int `json:"couples"`
	Active   int `json:"active"`
}

func (l *LevelTravelApi) HotelsByIds(hotelIds []string) (hotels Hotels, err error) {
	err = l.getJson("/references/hotels", map[string]string{"hotel_ids": strings.Join(hotelIds, ",")}, &hotels)
	return
}

type ShortHotels struct {
	Success bool         `json:"success"`
	Hotels  []ShortHotel `json:"hotels"`
}

type ShortHotel struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Stars     int    `json:"stars"`
	CityId    int    `json:"city_id"`
	CountryId int    `json:"country_id"`
}

func (l *LevelTravelApi) HotelsByRegions(regionIds []string) (hotels ShortHotels, err error) {
	err = l.getJson("/references/hotels", map[string]string{"region_ids": strings.Join(regionIds, ",")}, &hotels)
	return
}

type HotelsDump struct {
	Success     bool   `json:"success"`
	GeneratedAt string `json:"generated_at"`
	Link        string `json:"link"`
}

func (l *LevelTravelApi) HotelsDump() (hotelsDump HotelsDump, err error) {
	err = l.getJson("/references/hotel_dump", map[string]string{}, &hotelsDump)
	return
}
