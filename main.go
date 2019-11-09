package main

import "fmt"

func main() {
	//r := mux.NewRouter()
	//r.HandleFunc("/", homeHandler)
	//log.Fatal(http.ListenAndServe(":8080", r))

	var dbConn DbConn

	dbConn.port = "test"
	db := dbConn.GetConn()

	fmt.Println(db)
}
