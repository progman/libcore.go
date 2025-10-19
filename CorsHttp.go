//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
package libcore
//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
import (
	"net/http"
)
//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
func CorsHttp(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")  // "POST, GET, PUT, DELETE, OPTIONS" // "*"
		w.Header().Set("Access-Control-Allow-Headers", "*")  // "DNT, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Range, Origin, Accept, Total-Count, Authorization, Access-Control-Allow-Origin" // "*"
		w.Header().Set("Access-Control-Expose-Headers", "*") // "Content-Length, Content-Range"


		if r.Method == "OPTIONS" {

			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Max-Age", "1728000")
			w.WriteHeader(http.StatusNoContent) // 204 No Content
			return
		}


		next.ServeHTTP(w, r) // next
	})
}
/*
	router := NewRouter()
	log.Printf("- starting REST API server on \"%s:%s\"...\n", Global.BindHost, Global.BindPort)
	log.Fatal(http.ListenAndServe(Global.BindHost + ":" + Global.BindPort, libcore.CorsHttp(router)))
*/
//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
