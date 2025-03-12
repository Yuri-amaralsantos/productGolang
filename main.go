package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"

    "github.com/gorilla/mux"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

var db *gorm.DB

type Product struct {
    ID          uint    `json:"id" gorm:"primaryKey"`
    Name        string  `json:"name"`
    Description string  `json:"description"`
    Price       float64 `json:"price"`
}

func initDB() {
    dsn := "host=localhost user=postgres password=admin dbname=productdb port=5432 sslmode=disable"
    var err error
    db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }

    db.AutoMigrate(&Product{})
}

// CORS middleware function
func enableCORS(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "*") // Allow all origins
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

        if r.Method == "OPTIONS" {
            w.WriteHeader(http.StatusOK)
            return
        }

        next.ServeHTTP(w, r)
    })
}

func createProduct(w http.ResponseWriter, r *http.Request) {
    var product Product
    json.NewDecoder(r.Body).Decode(&product)
    db.Create(&product)
    json.NewEncoder(w).Encode(product)
}

func getProducts(w http.ResponseWriter, r *http.Request) {
    var products []Product
    db.Find(&products)
    json.NewEncoder(w).Encode(products)
}

func main() {
    initDB()

    r := mux.NewRouter()
    r.HandleFunc("/products", createProduct).Methods("POST")
    r.HandleFunc("/products", getProducts).Methods("GET")

    // Wrap the router with the CORS middleware
    handler := enableCORS(r)

    fmt.Println("Server running on port 8080")
    log.Fatal(http.ListenAndServe(":8080", handler))
}
