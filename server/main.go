package main

///////////////
//  Imports  //
///////////////

import (
	// Golang
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	// Internal
	"erp/resources/mariadb"
)

func ServeFile(w http.ResponseWriter, r *http.Request, path string) {
	safePath, err := url.QueryUnescape(path)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	http.ServeFile(w, r, safePath)
}

func GetRouter(w http.ResponseWriter, r *http.Request) {
	ServeFile(w, r, "static/"+r.URL.Path)
}

func MainRouter(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		GetRouter(w, r)
	} else {
		w.Write([]byte("Method not permited"))
	}
}

func testDB(){
	// Set DB connection
	db,err := mariadb.ConfigExtended(os.Getenv("MARIADB_USER"),os.Getenv("MARIADB_PASSWORD"),"db","3306",os.Getenv("MARIADB_DATABASE"))
	if err != nil { log.Println("hola");panic(err) }
	
	// Create table
	_,err = db.Create("CREATE TABLE `Persons` ( `id` INT NOT NULL AUTO_INCREMENT, `name` VARCHAR(45) NOT NULL, `birthday` DATE, secret BLOB, PRIMARY KEY (`id`))")
	if err != nil { panic(err) }

	// Insert data
	name := "Paco Gonzalez"
	birthday := "1983-04-11"
	data := []byte{0x34,0x45,0x12,0x05}
	_,err = db.Insert("INSERT INTO `Persons` (`name`,`birthday`,`secret`) VALUES (?,?,?)",name,birthday,data)
	if err != nil { panic(err) }

	name = "Martina Repuello"
	birthday = "1985-02-15"
	data = []byte{0x12,0x13,0x14,0x15}
	_,err = db.Insert("INSERT INTO `Persons` (`name`,`birthday`,`secret`) VALUES (?,?,?)",name,birthday,data)
	if err != nil { panic(err) }

	name = "Lapinta Marquisina"
	birthday = "1835-11-28"
	_,err = db.Insert("INSERT INTO `Persons` (`name`,`birthday`,`secret`) VALUES (?,?,NULL)",name,birthday)
	if err != nil { panic(err) }
	
	// Retrieve data
	output,err := db.Select("SELECT * FROM `Persons`")
	if err != nil { panic(err) }

	fmt.Println(output)
}

func main() {
	// Test database
	testDB()
	// Initialize Server
	fmt.Println("Server initialized on port " + os.Getenv("PORT"))
	// Main Router
	http.HandleFunc("/", MainRouter)
	// Listen And Serve
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		panic(err)
	}
}
