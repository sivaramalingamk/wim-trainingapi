package domain
//this is the output from training/ data merger app
type TrainingData struct {
	ID        string `json:"id"`
	weight int `json:"weight"`
	vehicleSpeed int `json:"vehicleSpeed"`
	acceleration int `json:"acceleration"`
	windSpeed int `json:"windSpeed"`
	headingDirection int `json:"headingDirection"`
	windDirection int `json:"windDirection"`
	atmosphereTemp int `json:"atmosphereTemp"`
	atmospherePressure int `json:"atmospherePressure"`
	coolentTemp int `json:"coolentTemp"`
	oilPressure int `json:"oilPressure"`
	intakeAirTemp int `json:"intakeAirTemp"`
	humidity int `json:"humidity"`
	rpm int `json:"rpm"`
	engineLoad int `json:"engineLoad"`
	elevationAngle int `json:"elevationAngle"`
	o2 int `json:"o2"`
	fuelFlow int `json:"fuelFlow"`

}
