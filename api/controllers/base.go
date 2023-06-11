package controllers

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/rs/cors"

	//mysql database driver
	"github.com/gorilla/handlers"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/semicolon27/api-e-voting/api/models"
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

func (server *Server) Initialize(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {

	var err error

	DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)
	server.DB, err = gorm.Open(Dbdriver, DBURL)
	if err != nil {
		fmt.Printf("Cannot connect to %s database", Dbdriver)
		log.Fatal("This is the error:", err)
	} else {
		fmt.Printf("We are connected to the %s database", Dbdriver)
	}

	server.DB.Debug().AutoMigrate(&models.Admin{}, &models.Candidate{}) //database migration

	server.Router = mux.NewRouter()

	server.initializeRoutes()

}

func (server *Server) Run(addr string) {
	c := cors.Default().Handler(server.Router)

	credentials := handlers.AllowCredentials()
	methods := handlers.AllowedMethods([]string{"POST", "GET", "PUT", "DELETE", "*"})
	// ttl := handlers.MaxAge(3600)
	origins := handlers.AllowedOrigins([]string{"*", "localhost:4000"})

	fmt.Println("Listening to port 8080")
	log.Fatal(http.ListenAndServe(addr, handlers.CORS(credentials, methods, origins)(server.logRequest(c))))
}

// Middleware untuk logging request
func (server *Server) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		// Menyiapkan recorder response untuk mendapatkan status response
		// recorder := httptest.NewRecorder()

		// // Memanggil handler berikutnya menggunakan recorder response
		// next.ServeHTTP(recorder, r)

		// // Mengambil status response dari recorder
		// status := recorder.Result().StatusCode

		// // Mengirim response yang telah direkam ke response writer asli
		// for k, v := range recorder.Header() {
		// 	w.Header()[k] = v
		// }
		// w.WriteHeader(status)
		// recorder.Body.WriteTo(w)

		// Memanggil handler berikutnya
		next.ServeHTTP(w, r)

		// Logging setelah handler selesai dipanggil
		log.Printf(
			"[%s] %s Duration: %s",
			r.Method,
			r.URL.Path,
			// strconv.Itoa(status),
			time.Since(start),
		)
	})
}

// Middleware to enable CORS
func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*") // Replace * with your desired origin or set it dynamically based on the request origin
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// Handle preflight requests
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Call the next handler
		next.ServeHTTP(w, r)
	})
}
