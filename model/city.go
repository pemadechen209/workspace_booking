package model

import (
	"context"
	"fmt"
	"time"
	"workspace_booking/migration"
)

// Building struct
type City struct {
	Id        int64     `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Buildings struct
type Cities struct {
	Cities []*City
}

// Building struct
type WorkSpaces struct {
	CityList     []*City
	LocationList []*Location
	BuildingList []*Building
	FloorList    []*Floor
	Purpose      []string
}

func GetAllCities() []*City {
	rows, e := migration.DbPool.Query(context.Background(), "select * from cities")
	defer rows.Close()

	// declare empty post variable
	cities := make([]*City, 0)
	// iterate over rows
	for rows.Next() {
		city := new(City)

		e = rows.Scan(&city.Id, &city.Name, &city.CreatedAt, &city.UpdatedAt)

		if e != nil {
			fmt.Println("Failed to get buildings record :", e)
			return []*City{}
		}
		cities = append(cities, city)
	}
	return cities
}

func (city *City) CreateCity() error {
	dt := time.Now()
	query := "INSERT INTO cities (name, created_at, updated_at) VALUES ($1, $2, $3) RETURNING id, created_at, updated_at"
	d := migration.DbPool.QueryRow(context.Background(), query, &city.Name, dt, dt)
	d.Scan(&city.Id, &city.CreatedAt, &city.UpdatedAt)
	return nil
}

func GetCityByID(cityId int) City {
	city := City{}
	rows := migration.DbPool.QueryRow(context.Background(), "select * from cities where id = $1", cityId)
	err := rows.Scan(&city.Id, &city.Name, &city.CreatedAt, &city.UpdatedAt)
	if err != nil {
		fmt.Println("Failed to get locations record :", err)
		return City{}
	}
	return city
}

func GetAllDetails() WorkSpaces {
	cities := GetAllCities()
	locations := GetAllLocations()
	buildings := GetAllBuildings()
	floors := GetAllFloors()
	return WorkSpaces{
		CityList:     cities,
		LocationList: locations,
		BuildingList: buildings,
		FloorList:    floors,
		Purpose:      []string{"Team meeting", "Client Call", "Monthly Sync-up", "Others"},
	}
}
