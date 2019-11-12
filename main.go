package main

import "fmt"

func init() {

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
