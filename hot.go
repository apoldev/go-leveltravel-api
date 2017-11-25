package leveltravel

import (
	"fmt"
	"strings"
)

type HotTours struct {
	Success  bool         `json:"success"`
	HotTours []HotTour    `json:"hot_tours"`
	Pages    HotTourPages `json:"pages"`
}

type HotTour struct {
	Id                 int          `json:"id"`
	Link               string       `json:"link"`
	Date               string       `json:"date"`
	Nights             int          `json:"nights"`
	Adults             int          `json:"adults"`
	Region             string       `json:"region"`
	Country            string       `json:"country"`
	Discount           int          `json:"discount"`
	Transfer           bool         `json:"transfer"`
	MedicalInsurance   bool         `json:"medical_insurance"`
	PansionName        string       `json:"pansion_name"`
	PansionDescription string       `json:"pansion_description"`
	Hotel              HotTourHotel `json:"hotel"`
}

type HotTourHotel struct {
	Id      int     `json:"id"`
	Name    string  `json:"name"`
	Lat     float64 `json:"lat"`
	Long    float64 `json:"long"`
	Picture string  `json:"picture"`
}

type HotTourPages struct {
	Count   int `json:"count"`
	Current int `json:"current"`
	Total   int `json:"total"`
}

type HotToursFilter struct {
	Countries  []string
	NightsFrom int
	NightsTo   int
	StarsFrom  int
	StarsTo    int
	Adults     int
	Pansions   []string
	MinPrice   int
	MaxPrice   int
	PerPage    int
	Page       int
}

func (l *LevelTravelApi) HotTours(startDate, endDate, sortBy string, filter HotToursFilter) (hotTours *HotTours, err error) {
	params := map[string]string{
		"start_date": startDate,
		"end_date":   endDate,
		"sort_by":    sortBy,
	}

	if len(filter.Countries) > 0 {
		params["countries"] = strings.Join(filter.Countries, ",")
	}

	if filter.NightsFrom > 0 && filter.NightsTo >= filter.NightsFrom {
		params["nights"] = fmt.Sprintf("%d,%d", filter.NightsFrom, filter.NightsTo)
	}

	if filter.Adults > 0 {
		params["adults"] = fmt.Sprintf("%d", filter.Adults)
	}

	if len(filter.Pansions) > 0 {
		params["pansions"] = strings.Join(filter.Pansions, ",")
	}

	if filter.MinPrice > 0 {
		params["min_price"] = fmt.Sprintf("%d", filter.MinPrice)
	}

	if filter.MaxPrice > 0 {
		params["max_price"] = fmt.Sprintf("%d", filter.MaxPrice)
	}

	if filter.PerPage > 0 {
		params["per_page"] = fmt.Sprintf("%d", filter.PerPage)
	}

	if filter.Page > 0 {
		params["page"] = fmt.Sprintf("%d", filter.Page)
	}

	err = l.getJson("/hot/tours", params, hotTours)
	return
}
