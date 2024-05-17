package parkrunparser

import (
	"encoding/json"
	"fmt"
)

type JCountry struct {
	Url    string    `json:"url"`
	Bounds []float64 `json:"bounds"`
}

type JGeometry struct {
	Type        string    `json:"type"`
	Coordinates []float64 `json:"coordinates"`
}

type JProperties struct {
	Name              string `json:"eventname"`
	LongName          string `json:"EventLongName"`
	ShortName         string `json:"EventShortName"`
	LocalizedLongName string `json:"LocalisedEventLongName"`
	CountryCode       int    `json:"countrycode"`
	SeriesId          int    `json:"seriesid"`
	Location          string `json:"EventLocation"`
}

type JFeature struct {
	Id         int         `json:"id"`
	Type       string      `json:"type"`
	Geometry   JGeometry   `json:"geometry"`
	Properties JProperties `json:"properties"`
}

type JEvents struct {
	Type     string     `json:"type"`
	Features []JFeature `json:"features"`
}

type JEventsJson struct {
	Counties map[string]JCountry `json:"countries"`
	Events   JEvents             `json:"events"`
}

type LatLng struct {
	Lat, Lng float64
}

type Country struct {
	Url    string
	Bounds []LatLng
	Events []*Event
}

var url2countryname map[string]string = nil

func (c Country) Name() string {
	if url2countryname == nil {
		url2countryname = make(map[string]string)
		url2countryname["www.parkrun.ca"] = "Canada"
		url2countryname["www.parkrun.co.at"] = "Austria"
		url2countryname["www.parkrun.fi"] = "Finland"
		url2countryname["www.parkrun.fr"] = "France"
		url2countryname["www.parkrun.co.nl"] = "Netherlands"
		url2countryname["www.parkrun.no"] = "Norway"
		url2countryname["www.parkrun.pl"] = "Poland"
		url2countryname["www.parkrun.sg"] = "Singapore"
		url2countryname["www.parkrun.dk"] = "Denmark"
		url2countryname["www.parkrun.jp"] = "Japan"
		url2countryname["www.parkrun.us"] = "Unites States"
		url2countryname["www.parkrun.com.au"] = "Australia"
		url2countryname["www.parkrun.com.de"] = "Germany"
		url2countryname["www.parkrun.ie"] = "Ireland"
		url2countryname["www.parkrun.it"] = "Italy"
		url2countryname["www.parkrun.my"] = "Malaysia"
		url2countryname["www.parkrun.co.nz"] = "New Zealand"
		url2countryname["www.parkrun.co.za"] = "South Africa"
		url2countryname["www.parkrun.se"] = "Sweden"
		url2countryname["www.parkrun.org.uk"] = "United Kingdom"
	}

	if name, ok := url2countryname[c.Url]; ok {
		return name
	}
	return "UNKNOWN"
}

type EventType int

const (
	EVENTTYPE_REGULAR EventType = iota
	EVENTTYPE_JUNIOR
)

type Event struct {
	Name        string
	LongName    string
	ShortName   string
	Location    string
	Coordinates LatLng
	Country     *Country
	Type        EventType
}

func (e Event) Url() string {
	return fmt.Sprintf("%s/%s", e.Country.Url, e.Name)
}

type Events struct {
	Countries []*Country
	Events    []*Event
}

func ParseEvents(data []byte) (Events, error) {
	var eventsJson JEventsJson
	if err := json.Unmarshal(data, &eventsJson); err != nil {
		fmt.Printf("error unmarshalling: %v", err)
	}

	events := Events{}
	events.Countries = make([]*Country, 0)
	events.Events = make([]*Event, 0)

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return events, err
	}

	// build countries lookup
	countryLookup := make(map[string]*Country)

	for countryId, countryD := range eventsJson.Counties {
		if len(countryD.Bounds) != 4 {
			return events, fmt.Errorf("invalid size of 'countries/%s/bounds'", countryId)
		}
		bounds := make([]LatLng, 2)
		bounds[0].Lat = countryD.Bounds[0]
		bounds[0].Lng = countryD.Bounds[1]
		bounds[1].Lat = countryD.Bounds[2]
		bounds[1].Lng = countryD.Bounds[3]

		country := &Country{countryD.Url, bounds, make([]*Event, 0)}
		events.Countries = append(events.Countries, country)
		countryLookup[countryId] = country
	}

	for _, featureD := range eventsJson.Events.Features {
		if featureD.Geometry.Type != "Point" {
			return events, fmt.Errorf("'events/features/geometry/type' = '%s' (expected: 'Point')", featureD.Geometry.Type)
		}
		if len(featureD.Geometry.Coordinates) != 2 {
			return events, fmt.Errorf("invalid size of 'events/features/geometry/coordinates'")
		}
		coordinates := LatLng{featureD.Geometry.Coordinates[0], featureD.Geometry.Coordinates[1]}

		countryCode := fmt.Sprintf("%d", featureD.Properties.CountryCode)
		country, ok := countryLookup[countryCode]
		if !ok {
			return events, fmt.Errorf("failed to lookup country '%s' (event '%s')", countryCode, featureD.Properties.Name)
		}

		eventType := EVENTTYPE_REGULAR
		switch featureD.Properties.SeriesId {
		case 1:
			eventType = EVENTTYPE_REGULAR
		case 2:
			eventType = EVENTTYPE_JUNIOR
		default:
			return events, fmt.Errorf("invalid seriesId '%d' (event '%s')", featureD.Properties.SeriesId, featureD.Properties.Name)
		}

		event := &Event{featureD.Properties.Name, featureD.Properties.LongName, featureD.Properties.ShortName, featureD.Properties.Location, coordinates, country, eventType}
		country.Events = append(country.Events, event)
		events.Events = append(events.Events, event)
	}

	return events, nil
}
