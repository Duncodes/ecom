package server

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"

	"github.com/Duncodes/ecom/auth"
	"github.com/Duncodes/ecom/database"
	"github.com/gorilla/mux"
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

// PlaceOrder ....
func PlaceOrder(w http.ResponseWriter, r *http.Request) {
	// Read user claims and get more data about the use
	claims := r.Context().Value(auth.ConfigKey).(auth.JwtClaims)
	user, err := database.GetUserByUUID(claims.StandardClaims.Id)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}
	// TODO place an order base on the content of the cart

	// READ the content of the cart
	/*
		{
			"items":[
			{
				"quantity":1,
				"productid":1234,
			}
			]
		}
	*/
	var cart database.Cart
	err = json.NewDecoder(r.Body).Decode(&cart)
	if err != nil {
		log.Println(err)
		//panic(err)
		w.WriteHeader(400)
		return
	}

	// if order is empty
	if len(cart.Items) == 0 {
		log.Println("Provide items to order")
		w.WriteHeader(400)
		return
	}

	order, err := database.AddOrder(user.ID, cart)
	if err != nil {
		log.Println(err)
		w.WriteHeader(400)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	u, _ := json.Marshal(map[string]interface{}{"orders": order})
	w.Write(u)

}

// CategoryHandler ....
func CategoryHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		var category database.Category
		if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
			log.Println(err)
			return
		}
		if err := category.AddCategory(); err != nil {
			log.Println(err)
			return
		}
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	categories, err := database.GetCategories()
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{"categories": categories})
}

func PaymentMethods(w http.ResponseWriter, r *http.Request) {
	payments, err := database.GetPayments()
	if err != nil {
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	json.NewEncoder(w).Encode(map[string]interface{}{"payaments": &payments})
}

func CheckOut(w http.ResponseWriter, r *http.Request) {

}

func GetUserOrders(w http.ResponseWriter, r *http.Request) {
	claims := r.Context().Value(auth.ConfigKey).(auth.JwtClaims)
	user, err := database.GetUserByUUID(claims.StandardClaims.Id)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}

	orders, err := database.GetUserOrders(user.ID)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	json.NewEncoder(w).Encode(map[string]interface{}{"orders": &orders})
}

// StartServer ...
func StartServer(port string) {
	log.Println(port)
	r := mux.NewRouter()

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(dir))))
	r.HandleFunc("/api/products", ProductsHandler).Methods("GET")
	r.HandleFunc("/api", AddProductHandler).Methods("POST")
	r.HandleFunc("/api/product/{productid}", ProductHandler)
	r.HandleFunc("/api/payments", PaymentMethods)
	// Login
	r.HandleFunc("/api/login", LoginHandler)
	r.HandleFunc("/api/register", RegisterHandler)

	// r.Handle("/api/protected", auth.Authenticate(http.HandlerFunc(Protected)))
	//r.Handle("/api/{id}/cart", auth.Authenticate(http.HandlerFunc()))
	r.Handle("/api/order", auth.Authenticate(http.HandlerFunc(PlaceOrder)))
	r.Handle("/api/checkout", auth.Authenticate(http.HandlerFunc(CheckOut)))
	// TODO Create an admin auth middleware
	r.Handle("/api/category", auth.Authenticate(http.HandlerFunc(CategoryHandler)))
	r.Handle("/api/orders", auth.Authenticate(http.HandlerFunc(GetUserOrders)))
	log.Println("Starting server ......")
	srv := &http.Server{
		Handler: r,
		Addr:    "0.0.0.0:" + port,
	}

	log.Fatal(srv.ListenAndServe())
}
