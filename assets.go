package main

import "fmt"

// Asset asset details
type Asset struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Category string `json:"category"`
	Ctype    string `json:"type"`
	Model    string `json:"model"`
	Serial   string `json:"serial"`
	Brand    string `json:"brand"`
	MnfYear  string `json:"manufactured_year"`
	PDate    string `json:"purchased_data"`
	Price    string `json:"price"`
	Status   string `json:"status"`
}

// Init
func (a *Asset) Init() error {
	var (
		e Enum
		t Table
	)

	e.name = categoryEnumTypeName
	isExists, err := e.Exists()
	if !isExists {
		err = e.Create()
		if err != nil {
			return err
		}
	} else {
		fmt.Println("Enum category already exists")
	}

	e.name = typeEnumTypeName
	isExists, err = e.Exists()
	if !isExists {
		err = e.Create()
		if err != nil {
			return err
		}
	} else {
		fmt.Println("Enum type already exists")
	}

	e.name = brandEnumTypeName
	isExists, err = e.Exists()
	if !isExists {
		err = e.Create()
		if err != nil {
			return err
		}
	} else {
		fmt.Println("Enum brand already exists")
	}

	/*
		CREATE TABLE assets (
		        id                      SERIAL          PRIMARY KEY     NOT NULL,
		        name                    VARCHAR(100)    NOT NULL,
		        category                VARCHAR(20)     NOT NULL,
		        type                    VARCHAR(20)     NOT NULL,
		        model                   VARCHAR(20)     NOT NULL,
		        serial                  VARCHAR(20)     UNIQUE     NOT NULL,
		        brand                   VARCHAR(20)     NOT NULL,
		        manufactured_year       INT             NOT NULL,
		        purchased_date          DATE,
		        price                   FLOAT,
		        status                  VARCHAR(10)     NOT NULL
		);
	*/

	t = Table{
		Name: assetsTableName,
		Columns: []TableColumn{
			{
				Name:     "id",
				DataType: "serial",
				constraints: []string{
					"primary key",
					"not null",
				},
			},
			{
				Name:     "name",
				DataType: "varchar(100)",
				constraints: []string{
					"not null",
				},
			},
			{
				Name:     "category",
				DataType: categoryEnumTypeName,
				constraints: []string{
					"not null",
				},
			},
			{
				Name:     "type",
				DataType: typeEnumTypeName,
				constraints: []string{
					"not null",
				},
			},
			{
				Name:     "model",
				DataType: "varchar(50)",
				constraints: []string{
					"not null",
				},
			},
			{
				Name:     "serial",
				DataType: "varchar(50)",
				constraints: []string{
					"not null",
				},
			},
			{
				Name:     "brand",
				DataType: brandEnumTypeName,
				constraints: []string{
					"not null",
				},
			},
			{
				Name:     "manufactured_year",
				DataType: "int",
				constraints: []string{
					"not null",
				},
			},
			{
				Name:     "purchased_date",
				DataType: "date",
			},
			{
				Name:     "price",
				DataType: "float",
			},
			{
				Name:     "status",
				DataType: "varchar(10)",
			},
		},
	}

	isExists, err = t.Exists()
	if err != nil {
		return err
	}
	if !isExists {
		err = t.Create()
		if err != nil {
			return err
		}
	} else {
		fmt.Println("assets table already exists")
	}

	return nil
}

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
	sqlStatement := fmt.Sprintf(`insert into %s (name, category, type, model, serial, brand, manufactured_year, purchased_date, price, status) 
		VALUES ('%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s' ) RETURNING id`,
		assetsTableName, a.Name, a.Category, a.Ctype, a.Model, a.Serial, a.Brand, a.MnfYear, a.PDate, a.Price, a.Status)
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
