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

// TODO : Return custom errors? Log Errors here?
// TODO : Close DB connection

// Init initialises the database
// The method adds the required enums and tables and returns a error if any
func (a *Asset) Init() error {
	var (
		e Enum
		t Table
	)

	Logger{Message: appInitStr}.Info()

	e.name = categoryEnumTypeName
	isExists, err := e.Exists()
	if !isExists {
		err = e.Create()
		if err != nil {
			return err
		} else {
			Logger{Message: fmt.Sprintf(enumCreatedStr, categoryEnumName)}.Info()
		}
	} else {
		Logger{Message: fmt.Sprintf(enumExistsStr, categoryEnumName)}.Info()
	}

	e.name = typeEnumTypeName
	isExists, err = e.Exists()
	if !isExists {
		err = e.Create()
		if err != nil {
			return err
		} else {
			Logger{Message: fmt.Sprintf(enumCreatedStr, typeEnumName)}.Info()
		}
	} else {
		Logger{Message: fmt.Sprintf(enumExistsStr, typeEnumName)}.Info()
	}

	e.name = brandEnumTypeName
	isExists, err = e.Exists()
	if !isExists {
		err = e.Create()
		if err != nil {
			return err
		} else {
			Logger{Message: fmt.Sprintf(enumCreatedStr, brandEnumName)}.Info()
		}
	} else {
		Logger{Message: fmt.Sprintf(enumExistsStr, brandEnumName)}.Info()
	}

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
					"unique",
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
		} else {
			Logger{Message: fmt.Sprintf(tableCreatedStr, assetsTableName)}.Info()
		}
	} else {
		Logger{Message: fmt.Sprintf(tableExistsStr, assetsTableName)}.Info()
	}

	Logger{Message: appInitCompletedStr}.Info()

	return nil
}

// GetAssets
func (a *Asset) Get() ([]Asset, error) {
	var (
		dbConn DbConn
		assets []Asset
	)
	db, err := dbConn.Connect()
	defer dbConn.Close(db)
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
	return assets, nil
}

// GetAssetByID

// GetAssetByType

// GetAssetByCategory

// AddAsset
// TODO : Check if a asset already exists
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
