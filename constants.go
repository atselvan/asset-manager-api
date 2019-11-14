package main

import "fmt"

const (
	assetsTableName     = "assets"
	categoryEnumName    = "category"
	typeEnumName        = "type"
	brandEnumName       = "brand"
	appInitStr          = "Begin application initialisation"
	appInitCompletedStr = "Application initialisation Completed"
	appInitErrorStr     = "There was an error in the initialisation"
	enumNameFormat      = "%s_%s"
	enumExistsStr       = "Enum '%s' already exists"
	enumCreatedStr      = "Enum '%s' is created in the database"
	tableExistsStr      = "Table '%s' already exists"
	tableCreatedStr     = "Table '%s' is created"
)

var (
	categoryEnumTypeName = fmt.Sprintf(enumNameFormat, assetsTableName, categoryEnumName)
	typeEnumTypeName     = fmt.Sprintf(enumNameFormat, assetsTableName, typeEnumName)
	brandEnumTypeName    = fmt.Sprintf(enumNameFormat, assetsTableName, brandEnumName)
)
