package geocoding

import "github.com/kelvins/geocoder"

type GoogleGeocoder struct {
	ApiKey string
}

func (gg GoogleGeocoder) GetCoordsByName(city string) (latitude, longitude float64, err error) {
	geocoder.ApiKey = gg.ApiKey // TODO

	address := geocoder.Address{
		City: city,
	}

	// Convert address to location (latitude, longitude)
	location, err := geocoder.Geocoding(address)
	return location.Latitude, location.Longitude, err
}
