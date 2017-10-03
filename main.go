package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/loggercode/ecom/database"
)

var (
	dir string
)

func init() {
	flag.StringVar(&dir, "dir", ".", "the directory to serve files from. Defaults to the current dir")
}

func ItemsHandler(w http.ResponseWriter, r *http.Request) {
	// TODO introduce pagination
	items, err := database.GetItems()
	if err != nil {
		log.Println(err)
		return
	}
	responseitems := map[string]interface{}{
		"items": items,
		"count": len(items),
	}
	json.NewEncoder(w).Encode(responseitems)
}

func AddItemHandler(w http.ResponseWriter, r *http.Request) {
	var item database.Item
	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		log.Println("Error decoding request data: ", err)
		return
	}

	err = database.AddItem(item)

	if err != nil {
		log.Println("Error writing to database: ", err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func ItemHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	item, err := database.GetItem(vars["itemid"])
	if err != nil {
		return
	}
	json.NewEncoder(w).Encode(item)
}
func main() {
	flag.Parse()

	database.InitDB()
	r := mux.NewRouter()

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(dir))))
	r.HandleFunc("/api", ItemsHandler).Methods("GET")
	r.HandleFunc("/api", AddItemHandler).Methods("POST")
	r.HandleFunc("/api/item/{itemid}", ItemHandler)
	srv := &http.Server{
		Handler: r,
		Addr:    "0.0.0.0:9200",
	}

	log.Fatal(srv.ListenAndServe())
}
