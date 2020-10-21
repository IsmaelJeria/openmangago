package resource

import (
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

type HealthResource struct {
}

func NewHealthResource(r *fasthttprouter.Router) *HealthResource {
	var rsc = &HealthResource{}
	r.GET("/health", rsc.healthCheck)
	return rsc
}

func (rsc *HealthResource) healthCheck(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("application/json")
	ctx.SetBody([]byte(`{"status":"UP"}`))
	ctx.SetStatusCode(200)
}
