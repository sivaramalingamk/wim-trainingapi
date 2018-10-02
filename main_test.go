package main

import (
	"fmt"
	"github.com/sivaramalingamk/wim-trainingapi/domain"
	"github.com/stretchr/testify/assert"
	"gopkg.in/resty.v1"
	"testing"
)

func TestHomePage(t *testing.T) {
	resp, _ := resty.R().Get("http://localhost:8080")
	assert.Equal(t, 200, resp.StatusCode())
}

func TestIncomingRawData(t *testing.T) {
	data := domain.RawInputData{
		ID:               "123",
		Time:             "2017",
		Latitude:         "12234",
		Longitude:        "12345",
		Weight:           200,
		VehicleSpeed:     20,
		Acceleration:     20,
		HeadingDirection: 40,
		CoolentTemp:      50,
		OilPressure:      69,
		IntakeAirTemp:    70,
		Rpm:              80,
		EngineLoad:       90,
		ElevationAngle:   100,
		O2:               20,
		FuelFlow:         30,
	}
	resp, _ := resty.R().
		SetHeader("Content-Type", "application/json").
		SetBody(data).
		Post("http://localhost:8080/ecudata")
	fmt.Println("The Result For Single Data Point is ", resp)
	fmt.Println("The Result For Single Data Point  ", resp.Status())
	assert.Equal(t, "200 OK", resp.Status())
}

func TestIncomingRawDataCollection(t *testing.T) {
	data1 := domain.RawInputData{
		ID:               "123",
		Time:             "2017",
		Latitude:         "12234",
		Longitude:        "12345",
		Weight:           200,
		VehicleSpeed:     20,
		Acceleration:     20,
		HeadingDirection: 40,
		CoolentTemp:      50,
		OilPressure:      69,
		IntakeAirTemp:    70,
		Rpm:              80,
		EngineLoad:       90,
		ElevationAngle:   100,
		O2:               20,
		FuelFlow:         30,
	}

	data2 := domain.RawInputData{
		ID:               "567",
		Time:             "2017",
		Latitude:         "12234",
		Longitude:        "12345",
		Weight:           200,
		VehicleSpeed:     20,
		Acceleration:     20,
		HeadingDirection: 40,
		CoolentTemp:      50,
		OilPressure:      69,
		IntakeAirTemp:    70,
		Rpm:              80,
		EngineLoad:       90,
		ElevationAngle:   100,
		O2:               20,
		FuelFlow:         30,
	}

	data3 := domain.RawInputData{
		ID:               "891",
		Time:             "2017",
		Latitude:         "12234",
		Longitude:        "12345",
		Weight:           200,
		VehicleSpeed:     20,
		Acceleration:     20,
		HeadingDirection: 40,
		CoolentTemp:      50,
		OilPressure:      69,
		IntakeAirTemp:    70,
		Rpm:              80,
		EngineLoad:       90,
		ElevationAngle:   100,
		O2:               20,
		FuelFlow:         30,
	}
	data := []domain.RawInputData{data1, data2, data3}
	resp, _ := resty.R().
		SetHeader("Content-Type", "application/json").
		SetBody(data).
		Post("http://localhost:8080/bulkecudata")
	fmt.Println("The Result For Collections is ", resp)
	fmt.Println("The Result Collections is ", resp.Status())
	assert.Equal(t, "200 OK", resp.Status())
}

/*
var Session *gocql.Session


func TestDBConnection(t *testing.T) {

	var err error
	cluster := gocql.NewCluster("155.238.156.52")
	cluster.Keyspace="wim"
	Session, err = cluster.CreateSession()
	if err != nil {
		t.Fatalf("0x%x: NewCluster: %v",  err)
	}

}

func TestAddVehicleData(t *testing.T) {
	var _ error
	cluster := gocql.NewCluster("155.238.156.52")
	cluster.Keyspace="wim"
	Session, _ = cluster.CreateSession()
	if err := Session.Query("INSERT INTO vehicledata(id, weight) VALUES(?,?)","21212",1122).Exec(); err != nil {
		fmt.Println("Error while inserting Vehicle Data")
		fmt.Println(err)
	}
	Session.Close()
}
*/
