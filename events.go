package parkrunparser

import (
	"encoding/json"
	"fmt"
)

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
	EVENTTYPE_JUNIOR  EventType = iota
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
	events := Events{}
	events.Countries = make([]*Country, 0)
	events.Events = make([]*Event, 0)

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return events, err
	}

	// build countries lookup
	countryLookup := make(map[string]*Country)

	countriesI, ok := result["countries"]
	if !ok {
		return events, fmt.Errorf("cannot get 'countries'")
	}
	countriesD := countriesI.(map[string]interface{})

	for countryId, countryI := range countriesD {
		if countryId == "0" {
			continue
		}
		countryD := countryI.(map[string]interface{})

		urlI, ok := countryD["url"]
		if !ok {
			return events, fmt.Errorf("cannot get 'countries/%s/url'", countryId)
		}

		boundsI, ok := countryD["bounds"]
		if !ok {
			return events, fmt.Errorf("cannot get 'countries/%s/bounds'", countryId)
		}
		boundsD := boundsI.([]interface{})
		if len(boundsD) != 4 {
			return events, fmt.Errorf("invalid size of 'countries/%s/bounds'", countryId)
		}
		bounds := make([]LatLng, 2)
		bounds[0].Lat = boundsD[0].(float64)
		bounds[0].Lng = boundsD[1].(float64)
		bounds[1].Lat = boundsD[2].(float64)
		bounds[1].Lng = boundsD[3].(float64)

		country := &Country{urlI.(string), bounds, make([]*Event, 0)}
		events.Countries = append(events.Countries, country)
		countryLookup[countryId] = country
	}

	eventsI, ok := result["events"]
	if !ok {
		return events, fmt.Errorf("cannot get 'events'")
	}
	eventsD := eventsI.(map[string]interface{})

	featuresI, ok := eventsD["features"]
	if !ok {
		return events, fmt.Errorf("cannot get 'events/features'")
	}

	featuresD := featuresI.([]interface{})
	for _, featureI := range featuresD {
		featureD := featureI.(map[string]interface{})

		geometryI, ok := featureD["geometry"]
		if !ok {
			return events, fmt.Errorf("cannot get 'events/features/geometry'")
		}
		geometryD := geometryI.(map[string]interface{})
		geoTypeI, ok := geometryD["type"]
		if !ok {
			return events, fmt.Errorf("cannot get 'events/features/geometry/type'")
		}
		if geoTypeI.(string) != "Point" {
			return events, fmt.Errorf("'events/features/geometry/type' = '%s' (expected: 'Point')", geoTypeI.(string))
		}
		geoCoordinatesI, ok := geometryD["coordinates"]
		if !ok {
			return events, fmt.Errorf("cannot get 'events/features/geometry/coordinates'")
		}
		geoCoordinatesD := geoCoordinatesI.([]interface{})
		if len(geoCoordinatesD) != 2 {
			return events, fmt.Errorf("invalid size of 'events/features/geometry/coordinates'")
		}
		coordinates := LatLng{geoCoordinatesD[0].(float64), geoCoordinatesD[1].(float64)}

		propertiesI, ok := featureD["properties"]
		if !ok {
			return events, fmt.Errorf("cannot get 'events/features/properties'")
		}

		propertiesD := propertiesI.(map[string]interface{})
		nameI, ok := propertiesD["eventname"]
		if !ok {
			return events, fmt.Errorf("cannot get 'events/features/properties/eventname'")
		}
		longNameI, ok := propertiesD["EventLongName"]
		if !ok {
			return events, fmt.Errorf("cannot get 'events/features/properties/EventLongName'")
		}
		shortNameI, ok := propertiesD["EventShortName"]
		if !ok {
			return events, fmt.Errorf("cannot get 'events/features/properties/EventShortName'")
		}
		locationI, ok := propertiesD["EventLocation"]
		if !ok {
			return events, fmt.Errorf("cannot get 'events/features/properties/EventLocation'")
		}
		countryCodeI, ok := propertiesD["countrycode"]
		if !ok {
			return events, fmt.Errorf("cannot get 'events/features/properties/countrycode'")
		}

		name := nameI.(string)
		countryCode := fmt.Sprintf("%.0f", countryCodeI.(float64))

		country, ok := countryLookup[countryCode]
		if !ok {
			return events, fmt.Errorf("failed to lookup country '%s' (event '%s')", countryCode, name)
		}

		eventType := EVENTTYPE_REGULAR

		event := &Event{name, longNameI.(string), shortNameI.(string), locationI.(string), coordinates, country, eventType}
		country.Events = append(country.Events, event)
		events.Events = append(events.Events, event)
	}

	return events, nil
}
