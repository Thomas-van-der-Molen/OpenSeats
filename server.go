package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type reservation struct{
	FName string
	LName string
	PNumber string
	date string
	RType string
	quantity string
	price string
	internalNotes string
}


func main(){
	
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)

	//routes
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/reservations", reservationsHandler)
	
	fmt.Println("server started on port 8080")
	err := http.ListenAndServe("127.0.0.1:8080", nil)
	if err != nil{
		log.Fatal(err)
	}
}

func querydb(query string) []string{

	db, _ := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/OpenSeats")
	defer db.Close()

	rows, _ := db.Query(query)
	defer rows.Close()
	var row reservation
	var results []string
	var temp string
	for rows.Next(){
		rows.Scan(&row.FName, &row.LName, &row.PNumber, &row.date, &row.RType, &row.quantity, &row.price, &row.internalNotes)
		temp = row.FName+","+row.LName+","+row.PNumber+","+row.date+","+row.RType+","+row.quantity+","+row.price+","+row.internalNotes
		results = append(results, temp)
	}
	return results
}

func helloHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "This is another message")
}

func reservationsHandler(w http.ResponseWriter, r *http.Request){
	query := "select * from reservations;"
	//fmt.Fprintf(w, "request executed successfully")
	//querydb(query)
	data := querydb(query)
	//w.Header().Set("Content-Type", "application/json")
	fmt.Println(data)
	//fmt.Fprintf(w, data)
	enc := json.NewEncoder(w)
	enc.Encode(data)
	//w.Write(querydb(query))
}
