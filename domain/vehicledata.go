package domain

import "time"

type VehicleData struct {
	ID        string `json:"id"`
	altitude string `json: "altitude"`
	longitude string `json:"longitude"`
	latitude  string `json:"latitude"`
	time time.Time `json: "time"`


}
