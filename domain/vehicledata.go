package domain
type VehicleData struct {
	ID        string `json:"id"`
	weight int `json:"weight"`
	vehicleSpeed int `json:"vehicleSpeed"`
	acceleration int `json:"acceleration"`
	headingDirection int `json:"headingDirection"`
	coolentTemp int `json:"coolentTemp"`
	oilPressure int `json:"oilPressure"`
	intakeAirTemp int `json:"intake_airTemp"`
	rpm int `json:"rpm"`
	engineLoad int `json:"engineLoad"`
	elevationAngle int `json:"elevationAngle"`
	o2 int `json:"o2"`
	fuelFlow int `json:"fuelFlow"`
}
