package controllers

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"

	//mysql database driver

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
	// // optionsHandler := cors.Default().Handler(server.Router)
	// optionsHandler := handlers.CORS(
	// 	handlers.AllowedHeaders([]string{"Authorization", "Content-Type", "X-CSRF-Token"}),
	// 	handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"}),
	// 	handlers.AllowedOrigins([]string{"*"}),
	// )(server.Router)

	server.Router.Use(mux.CORSMethodMiddleware(server.Router))

	fmt.Println("Listening to port 8080")
	// log.Fatal(http.ListenAndServe(addr, optionsHandler))
	log.Fatal(http.ListenAndServe(addr, server.Router))
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
