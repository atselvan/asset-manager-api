package main

import (
	"fmt"
	"github.com/atselvan/go-pgdb-lib"
	"github.com/atselvan/go-utils"
	"strings"
)

// Asset asset details
type Asset struct {
	Id       int     `json:"id"`
	Name     string  `json:"name"`
	Category string  `json:"category"`
	Ctype    string  `json:"type"`
	Brand    string  `json:"brand"`
	Model    string  `json:"model"`
	Colour   string  `json:"colour"`
	Serial   string  `json:"serial"`
	MnfYear  int  `json:"manufactured_year"`
	PDate    string  `json:"purchased_data"`
	Price    float64 `json:"price"`
	Status   string  `json:"status"`
}

// isValidAssetName check if asset name is valid
func (a *Asset) isValidAssetName() error {
	if a.Name == "" {
		return utils.Error{ErrMsg: assetNameReqStr}.NewError()
	}
	return nil
}

// isValidAssetCategory checks if asset category is valid
func (a *Asset) isValidAssetCategory() error {
	a.Category = strings.TrimSpace(a.Category)
	if a.Category == "" {
		return utils.Error{ErrMsg: assetCategoryReqStr}.NewError()
	}
	e := pgdb.Enum{Name: categoryEnumTypeName}
	if err := e.Get(); err != nil {
		return err
	}
	if !utils.EntryExists(e.Values, a.Category) {
		return utils.Error{ErrMsg: fmt.Sprintf(assetCategoryNotFoundStr, a.Category, utils.GetSliceAsCommaSeparatedString(e.Values))}.NewError()
	}
	return nil
}

// isValidAssetType checks if asset type is valid
func (a *Asset) isValidAssetType() error {
	a.Ctype = strings.TrimSpace(a.Ctype)
	if a.Ctype == "" {
		return utils.Error{ErrMsg: assetTypeReqStr}.NewError()
	}
	e := pgdb.Enum{Name: typeEnumTypeName}
	if err := e.Get(); err != nil {
		return err
	}
	if !utils.EntryExists(e.Values, a.Ctype) {
		return utils.Error{ErrMsg: fmt.Sprintf(assetTypeNotFoundStr, a.Category, utils.GetSliceAsCommaSeparatedString(e.Values))}.NewError()
	}
	return nil
}

// isValidAssetBrand checks if asset brand is valid
func (a *Asset) isValidAssetBrand() error {
	a.Brand = strings.TrimSpace(a.Brand)
	if a.Brand == "" {
		return utils.Error{ErrMsg: assetBrandReqStr}.NewError()
	}
	e := pgdb.Enum{Name: brandEnumTypeName}
	if err := e.Get(); err != nil {
		return err
	}
	if !utils.EntryExists(e.Values, a.Brand) {
		return utils.Error{ErrMsg: fmt.Sprintf(assetBrandNotFoundStr, a.Category, utils.GetSliceAsCommaSeparatedString(e.Values))}.NewError()
	}
	return nil
}

// isValidAssetModel checks if asset model is valid
func (a *Asset) isValidAssetModel() error {
	if a.Model == "" {
		return utils.Error{ErrMsg: assetModelReqStr}.NewError()
	}
	return nil
}

// isValidAssetColour checks if asset colour is valid
func (a *Asset) isValidAssetColour() error {
	if a.Colour == "" {
		return utils.Error{ErrMsg: assetColourReqStr}.NewError()
	}
	return nil
}

// isValidAssetSerial checks if asset serial is valid
func (a *Asset) isValidAssetSerial() error {
	a.Serial = strings.TrimSpace(a.Serial)
	if a.Serial == "" {
		return utils.Error{ErrMsg: assetSerialReqStr}.NewError()
	}
	if id, err := a.Exists(); err != nil {
		return utils.Error{ErrMsg: fmt.Sprintf(assetSerialExistsStr, a.Serial, id)}.NewError()
	}
	return nil
}

// isValidMnfYear check if the manufactured year is valid
func (a *Asset) isValidMnfYear() error {
	if a.MnfYear == 0 {
		return utils.Error{ErrMsg: assetMnfYearReqStr}.NewError()
	}
	if err := utils.IsValidYear(a.MnfYear); err != nil {
		return err
	}
	return nil
}

// isValidAssetPDate checks if the asset purchase date is valid
func (a *Asset) isValidAssetPDate() error {
	a.PDate = strings.TrimSpace(a.PDate)
	if a.PDate == "" {
		return utils.Error{ErrMsg: assetPDateReqStr}.NewError()
	}
	if err := utils.IsValidDate(a.PDate); err != nil {
		return err
	}
	return nil
}

// isValidAssetPrice trims spaces and returns the value of the asset price
func (a *Asset) isValidAssetPrice() error {
	if a.Price == 0 {
		return utils.Error{ErrMsg: assetPriceReqStr}.NewError()
	}
	return nil
}

// isValidAssetStatus checks if asset status is valid
func (a *Asset) isValidAssetStatus() error {
	a.Status = strings.TrimSpace(a.Status)
	if a.Status == "" {
		return utils.Error{ErrMsg: assetStatusReqStr}.NewError()
	}
	e := pgdb.Enum{Name: statusEnumTypeName}
	if err := e.Get(); err != nil {
		return err
	}
	if !utils.EntryExists(e.Values, a.Status) {
		return utils.Error{ErrMsg: fmt.Sprintf(assetStatusNotFoundStr, a.Category, utils.GetSliceAsCommaSeparatedString(e.Values))}.NewError()
	}
	return nil
}

func (a *Asset) IsValidAssetInfo() error {

	return nil
}

func (a *Asset) IsNotEmptyAssetInfo() error {
	if a.Name == "" || a.Category == "" || a.Ctype == "" || a.Model == "" || a.Serial == "" || a.Brand == "" || a.MnfYear == "" {
		return utils.NewError("name, category, type, model, serial, brand and manufactured_year are required parameters")
	} else {
		return nil
	}
}

// Init initialises the database
// The method adds the required enums and tables and returns a error if any
func (a *Asset) Init() error {
	var (
		e pgdb.Enum
		t pgdb.Table
	)
	utils.Logger{Message: appInitStr}.Info()
	// Check/Initialise Category enum type
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
	// Check/Initialise Type enum type
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
	// Check/Initialise brand enum type
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
	// Check/Initialise status enum type
	e.Name = statusEnumTypeName
	isExists, err = e.Exists()
	if !isExists {
		err = e.Create()
		if err != nil {
			return err
		} else {
			utils.Logger{Message: fmt.Sprintf(enumCreatedStr, statusEnumName)}.Info()
		}
	} else {
		utils.Logger{Message: fmt.Sprintf(enumExistsStr, statusEnumName)}.Info()
	}
	// Check/Initialize assets table
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
				Name:     "brand",
				DataType: brandEnumTypeName,
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
				Name:     "colour",
				DataType: "varchar(20)",
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
				DataType: statusEnumTypeName,
				Constraints: []string{
					"not null",
				},
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
		return nil, dbConn.ConnectionError(err)
	}
	rows, err := db.Query(fmt.Sprintf("Select * from %s", assetsTableName))
	if err != nil {
		return nil, dbConn.QueryExecError(err)
	}
	for rows.Next() {
		var a Asset
		err = rows.Scan(&a.Id, &a.Name, &a.Category, &a.Ctype, &a.Model, &a.Serial, &a.Brand, &a.MnfYear, &a.PDate, &a.Price, &a.Status)
		if err != nil {
			return nil, dbConn.RowScanError(err)
		}
		fmt.Println(a)
		assets = append(assets, a)
	}
	if err = dbConn.Close(db); err != nil {
		return nil, dbConn.ClosureError(err)
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
