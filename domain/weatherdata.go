package domain
type WeatherData struct {
	ID        string `json:"id"`
	wind_speed int `json:"wind_speed"`
	heading_direction int `json:"heading_direction"`
	wind_direction int `json:"wind_direction"`
	atmosphere_temp int `json:"atmosphere_temp"`
	atmosphere_pressure int `json:"atmosphere_pressure"`
	humidity int `json:"humidity"`
}
