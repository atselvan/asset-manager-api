package main

import "fmt"

func init() {
	enum := enum{
		name:  categoryEnumTypeName,
		value: "device",
	}
	err := enum.Add()

	err = enum.Update()

	enum.value = "game"

	err = enum.Update()

	fmt.Println("Add : ", err)

	b, e := enum.Exists()
	fmt.Println("Exists : ", b, e)

	v, _ := enum.Get()
	fmt.Println(v)
}

func main() {
	//r := mux.NewRouter()
	//r.HandleFunc("/", homeHandler)
	//log.Fatal(http.ListenAndServe(":8080", r))

	a := Asset{
		Name:     "MacBook Pro",
		Category: "devices",
		Kind:     "laptop",
		Model:    "MacBook Pro 15-inch SpaceGrey",
		Serial:   "C02VC1TBHTD511",
		Brand:    "Apple",
		MnfYear:  "2017",
		PDate:    "10/31/2017",
		Price:    "2799",
		Status:   "owned",
	}

	fmt.Println(a)
}
