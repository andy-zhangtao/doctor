package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// StartWeb 启动web服务
func StartWeb() (err error) {
	router := mux.NewRouter()
	router.Path("/_ping").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("I AM Doctor!"))
	})
	router.Path("/api").HandlerFunc(api)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"HEAD", "GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
		Debug:            false,
		AllowOriginFunc: func(origin string) bool {
			return true
		},
	})

	handler := c.Handler(router)
	fmt.Println(http.ListenAndServe(":8000", handler))
	return
}
