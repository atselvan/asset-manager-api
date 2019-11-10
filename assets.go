package main

import "fmt"

// Asset asset details
type Asset struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Category string `json:"category"`
	Kind     string `json:"kind"`
	Model    string `json:"model"`
	Serial   string `json:"serial"`
	Brand    string `json:"brand"`
	MnfYear  string `json:"manufactured_year"`
	PDate    string `json:"purchased_data"`
	Price    string `json:"price"`
	Status   string `json:"status"`
}

const (
	assetsTableName = "assets"
)

// GetAssets

// GetAssetByID

// GetAssetByType

// GetAssetByCategory

// AddAsset
func (a *Asset) Add() (string, error) {
	var dbConn DbConn
	db, err := dbConn.Connect()
	if err != nil {
		return "", err
	}
	sqlStatement := fmt.Sprintf(`insert into %s (name, category, kind, model, serial, brand, manufactured_year, purchased_date, price, status) 
		VALUES ('%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s' ) RETURNING id`,
		assetsTableName, a.Name, a.Category, a.Kind, a.Model, a.Serial, a.Brand, a.MnfYear, a.PDate, a.Price, a.Status)
	var id string
	err = db.QueryRow(sqlStatement).Scan(&id)
	if err != nil {
		return "", err
	}
	err = dbConn.Close(db)
	if err != nil {
		return "", err
	}
	return id, nil
}

// UpdateAsset

// DeleteAsset
