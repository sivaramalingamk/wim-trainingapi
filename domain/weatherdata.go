package domain
type WeatherData struct {
	ID        string `json:"id"`
	windSpeed int `json:"windSpeed"`
	headingDirection int `json:"headingDirection"`
	windDirection int `json:"windDirection"`
	atmosphereTemp int `json:"atmosphereTemp"`
	atmospherePressure int `json:"atmospherePressure"`
	humidity int `json:"humidity"`
}
