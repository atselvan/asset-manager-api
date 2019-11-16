package main

import (
	"fmt"
	"github.com/atselvan/go-pgdb-lib"
	"github.com/atselvan/go-utils"
)

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

// TODO : Return custom errors? Log Errors here?
// TODO : Close DB connection

// Init initialises the database
// The method adds the required enums and tables and returns a error if any
func (a *Asset) Init() error {
	var (
		e pgdb.Enum
		t pgdb.Table
	)

	utils.Logger{Message: appInitStr}.Info()

	e.Name = categoryEnumTypeName
	isExists, err := e.Exists()
	if !isExists {
		err = e.Create()
		if err != nil {
			return err
		} else {
			utils.Logger{Message: fmt.Sprintf(enumCreatedStr, categoryEnumName)}.Info()
		}
	} else {
		utils.Logger{Message: fmt.Sprintf(enumExistsStr, categoryEnumName)}.Info()
	}

	e.Name = typeEnumTypeName
	isExists, err = e.Exists()
	if !isExists {
		err = e.Create()
		if err != nil {
			return err
		} else {
			utils.Logger{Message: fmt.Sprintf(enumCreatedStr, typeEnumName)}.Info()
		}
	} else {
		utils.Logger{Message: fmt.Sprintf(enumExistsStr, typeEnumName)}.Info()
	}

	e.Name = brandEnumTypeName
	isExists, err = e.Exists()
	if !isExists {
		err = e.Create()
		if err != nil {
			return err
		} else {
			utils.Logger{Message: fmt.Sprintf(enumCreatedStr, brandEnumName)}.Info()
		}
	} else {
		utils.Logger{Message: fmt.Sprintf(enumExistsStr, brandEnumName)}.Info()
	}

	t = pgdb.Table{
		Name: assetsTableName,
		Columns: []pgdb.TableColumn{
			{
				Name:     "id",
				DataType: "serial",
				Constraints: []string{
					"primary key",
					"not null",
				},
			},
			{
				Name:     "name",
				DataType: "varchar(100)",
				Constraints: []string{
					"not null",
				},
			},
			{
				Name:     "category",
				DataType: categoryEnumTypeName,
				Constraints: []string{
					"not null",
				},
			},
			{
				Name:     "type",
				DataType: typeEnumTypeName,
				Constraints: []string{
					"not null",
				},
			},
			{
				Name:     "model",
				DataType: "varchar(50)",
				Constraints: []string{
					"not null",
				},
			},
			{
				Name:     "serial",
				DataType: "varchar(50)",
				Constraints: []string{
					"unique",
					"not null",
				},
			},
			{
				Name:     "brand",
				DataType: brandEnumTypeName,
				Constraints: []string{
					"not null",
				},
			},
			{
				Name:     "manufactured_year",
				DataType: "int",
				Constraints: []string{
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
		} else {
			utils.Logger{Message: fmt.Sprintf(tableCreatedStr, assetsTableName)}.Info()
		}
	} else {
		utils.Logger{Message: fmt.Sprintf(tableExistsStr, assetsTableName)}.Info()
	}

	utils.Logger{Message: appInitCompletedStr}.Info()

	return nil
}

// GetAssets
func (a *Asset) Get() ([]Asset, error) {
	var (
		dbConn pgdb.DbConn
		assets []Asset
	)
	db, err := dbConn.Connect()
	if err != nil {
		return nil, err
	}
	rows, err := db.Query(fmt.Sprintf("Select * from %s", assetsTableName))
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var a Asset
		err = rows.Scan(&a.Id, &a.Name, &a.Category, &a.Ctype, &a.Model, &a.Serial, &a.Brand, &a.MnfYear, &a.PDate, &a.Price, &a.Status)
		if err != nil {
			return nil, err
		}
		fmt.Println(a)
		assets = append(assets, a)
	}
	err = dbConn.Close(db)
	if err != nil {
		return nil, err
	}
	return assets, nil
}

// GetAssetByID

// GetAssetByType

// GetAssetByCategory

// Exists checks if a asset already exists in the database
// The method returns a boolean value and an error if something goes wrong
func (a *Asset) Exists() (string, error) {
	if err := a.validateSerial(); err != nil {
		return "", err
	}
	var dbConn pgdb.DbConn
	db, err := dbConn.Connect()
	if err != nil {
		return "", err
	}
	defer dbConn.Close(db)
	var id string
	err = db.QueryRow(fmt.Sprintf("select id from %s where serial = '%s'", assetsTableName, a.Serial)).Scan(&id)
	if id == "" {
		return id, err
	} else {
		return id, err
	}
}

func (a *Asset) validateSerial() error {
	if a.Serial == "" {
		return utils.NewError("'serial' is a required parameter")
	} else {
		return nil
	}
}

func (a *Asset) IsValid() error {
	if a.Name == "" || a.Category == "" || a.Ctype == "" || a.Model == "" || a.Serial == "" || a.Brand == "" || a.MnfYear == "" {
		return utils.NewError("name, category, type, model, serial, brand and manufactured_year are required parameters")
	} else {
		return nil
	}
}

// Add adds a asset to the assets table in the database
// The method returns an error if something goes wrong
func (a *Asset) Add() (string, error) {
	var dbConn pgdb.DbConn
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
