package server

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/loggercode/ecom/auth"
	"github.com/loggercode/ecom/database"
)

var (
	dir string
)

func init() {
	flag.StringVar(&dir, "dir", ".", "the directory to serve files from. Defaults to the current dir")
}

// ProductsHandler ...
func ProductsHandler(w http.ResponseWriter, r *http.Request) {
	// TODO introduce pagination
	items, err := database.GetProducts()
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

// AddProductHandler ....
func AddProductHandler(w http.ResponseWriter, r *http.Request) {
	var item database.Product
	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		log.Println("Error decoding request data: ", err)
		return
	}

	err = database.AddProduct(item)

	if err != nil {
		log.Println("Error writing to database: ", err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// ProductHandler ...
func ProductHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	item, err := database.GetProduct(vars["productid"])
	if err != nil {
		return
	}
	json.NewEncoder(w).Encode(item)
}

// LoginHandler ...
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var user database.UserCredential

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fulluser, err := user.VerifyUser()
	if err != nil {
		// TODO handle err with the appropiriate response
		log.Println("Wrong login credentilas")
		return
	}

	token, err := auth.GenerateJWTTokken(user.Username, fulluser.UUID)
	if err != nil {
		// TODO : handle error
		log.Fatal(err)
	}

	response := map[string]interface{}{
		"token": token,
		"uuid":  fulluser.UUID,
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(response)
}

// RegisterHandler ...
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var user database.User
	var err error
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Println(err)
		return
	}

	if err = user.CreateUser(); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	token, err := auth.GenerateJWTTokken(user.Username, user.UUID)

	if err != nil {
		log.Println(err)
		return
	}
	tokenresponse := map[string]string{
		"token": token,
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(tokenresponse)
}

func Protected(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Welcome to super protected endpoint"))
}

// StartServer ...
func StartServer() {

	r := mux.NewRouter()

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(dir))))
	r.HandleFunc("/api", ProductsHandler).Methods("GET")
	r.HandleFunc("/api", AddProductHandler).Methods("POST")
	r.HandleFunc("/api/product/{productid}", ProductHandler)

	// Login
	r.HandleFunc("/api/login", LoginHandler)
	r.HandleFunc("/api/register", RegisterHandler)

	r.Handle("/api/protected", auth.Authenticate(http.HandlerFunc(Protected)))
	log.Println("Starting server ......")
	srv := &http.Server{
		Handler: r,
		Addr:    "0.0.0.0:9200",
	}

	log.Fatal(srv.ListenAndServe())
}
