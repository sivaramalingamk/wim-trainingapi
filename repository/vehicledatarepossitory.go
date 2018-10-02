package repository

import (
	"fmt"
	"github.com/gocql/gocql"
	"github.com/sivaramalingamk/wim-trainingapi/domain"
)

//ID               string `json:"id"`
//Weight           int    `json:"weight"`
//VehicleSpeed     int    `json:"vehicleSpeed"`
//Acceleration     int    `json:"acceleration"`
//HeadingDirection int    `json:"headingDirection"`
//CoolentTemp      int    `json:"coolentTemp"`
//OilPressure      int    `json:"oilPressure"`
//IntakeAirTemp    int    `json:"intake_airTemp"`
//Rpm              int    `json:"rpm"`
//EngineLoad       int    `json:"engineLoad"`
//ElevationAngle   int    `json:"elevationAngle"`
//O2               int    `json:"o2"`
//FuelFlow         int    `json:"fuelFlow"

var Session *gocql.Session

func getCluster() *gocql.Session {
	var err error
	cluster := gocql.NewCluster("155.238.156.52")
	cluster.Keyspace = "vim"
	Session, err = cluster.CreateSession()
	if err != nil {
		panic(err)
	}
	return Session
}

func AddVehicleData(data domain.VehicleData) {
	fmt.Println(" **** Creating new data ****\n", data)
	defer getCluster().Close()
	if err := Session.Query("INSERT INTO vehicledata(id, weight) VALUES(?,?)", data.ID, data.Weight).Exec(); err != nil {

		fmt.Println("Error while inserting Vehicle Data", err)
		fmt.Println(err)
	}

}

func update(data domain.VehicleData) domain.VehicleData {
	fmt.Println(" **** Updating   data ****\n", data)
	updatedData := domain.VehicleData{
		ID:     data.ID,
		Weight: data.Weight,
	}
	defer getCluster().Close()
	if err := Session.Query("UPDATE vehicledata SET weight = ? WHERE ID = ?",
		updatedData.ID, updatedData.Weight).Exec(); err != nil {
		fmt.Println("Error while updating Emp")
		fmt.Println(err)
	}
	return updatedData

}

func deleteVehicleData(data domain.VehicleData) bool {
	fmt.Println(" **** Deleting Person  ****\n", data)
	defer getCluster().Close()
	if err := Session.Query("DELETE FROM vehicledata WHERE id = ?", data.ID).Exec(); err != nil {
		fmt.Println("Error while deleting Emp")
		fmt.Println(err)
	}
	return true

}

func readVehicleData(id string) domain.VehicleData {
	fmt.Println(" **** Reading a vdata with ID ****\n", id)
	defer getCluster().Close()
	var data domain.VehicleData
	m := map[string]interface{}{}
	iter := getCluster().Query("SELECT * FROM vdata WHERE id=? LIMIT 1").Iter()
	for iter.MapScan(m) {
		data = domain.VehicleData{
			ID:     m["ID"].(string),
			Weight: m["name"].(int),
		}
		m = map[string]interface{}{}
	}
	return data

}

func readAllVehicleData() []domain.VehicleData {
	fmt.Println(" **** Reading All Records ****\n")
	defer getCluster().Close()
	var data []domain.VehicleData
	m := map[string]interface{}{}
	defer getCluster().Close()
	iter := getCluster().Query("SELECT * FROM vdata").Iter()
	for iter.MapScan(m) {
		data = append(data, domain.VehicleData{
			ID:     m["id"].(string),
			Weight: m["name"].(int),
		})
		m = map[string]interface{}{}
	}
	return data
}
