package route

import (
	"net"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/metacubex/mihomo/component/mmdb"
)

func geoipRouter() http.Handler {
	r := chi.NewRouter()
	r.Get("/query", queryGeoIP)
	return r
}

func queryGeoIP(w http.ResponseWriter, r *http.Request) {
	ip := net.ParseIP(r.URL.Query().Get("ip"))
	if ip == nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, newError("invalid query ip"))
		return
	}

	responseData := render.M{"code": nil}

	if codes := mmdb.IPInstance().LookupCode(ip); len(codes) > 0 {
		responseData["code"] = codes[0]
	}

	render.JSON(w, r, responseData)
}
