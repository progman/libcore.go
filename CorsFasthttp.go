//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
package libcore
//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
import (
	"github.com/valyala/fasthttp"
)
//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
func CorsFasthttp(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		ctx.Response.Header.Set("Access-Control-Allow-Origin", "*")
		ctx.Response.Header.Set("Access-Control-Allow-Methods", "*")  // "POST, GET, PUT, DELETE, OPTIONS" // "*"
		ctx.Response.Header.Set("Access-Control-Allow-Headers", "*")  // "DNT, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Range, Origin, Accept, Total-Count, Authorization, Access-Control-Allow-Origin" // "*"
		ctx.Response.Header.Set("Access-Control-Expose-Headers", "*") // "Content-Length, Content-Range"


		if string(ctx.Method()) == "OPTIONS" {
			ctx.Response.Header.Set("Access-Control-Allow-Credentials", "true")
			ctx.Response.Header.Set("Access-Control-Max-Age", "1728000")
			ctx.SetStatusCode(fasthttp.StatusNoContent)
			return
		}


		next(ctx) // next
	}
}
/*
	r := router.New()
	log.Printf("- starting REST API server on \"%s:%s\"...\n", Global.BindHost, Global.BindPort)
	log.Fatal(fasthttp.ListenAndServe(Global.BindHost + ":" + Global.BindPort, libcore.CorsFasthttp(r.Handler)))
*/
//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
