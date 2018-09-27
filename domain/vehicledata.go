package domain

type VehicleData struct {
	ID        string `json:"id"`
	altitude string `json:"altitude"`
	longitude string `json:"longitude"`
	latitude  string `json:"latitude"`
	time string `json:"time"`
	weight int `json:"weight"`
	vehicle_speed int `json:"vehicle_speed"`
	acceleration int `json:"acceleration"`
	wind_speed int `json:"wind_speed"`
	heading_direction int `json:"heading_direction"`
	wind_direction int `json:"wind_direction"`
	atmosphere_temp int `json:"atmosphere_temp"`
	coolent_temp int `json:"coolent_temp"`
	intake_air_temp int `json:"intake_air_temp"`
	humidity int `json:"humidity"`
	rpm int `json:"rpm"`
	elevation_angle int `json:"elevation_angle"`
	road_type string `json:"road_type"`
	o2 int `json:"o_2"`
	fuel_flow float32 `json:"fuel_flow"`
	
	








}
