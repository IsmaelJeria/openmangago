package resource

import (
	"direst/service"
	"fmt"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

//UserResource ..
type UserResource struct {
	userService *service.UserService
}

//NewUserResource ..
func NewUserResource(s *service.UserService, r *fasthttprouter.Router) *UserResource {
	var rsc = &UserResource{userService: s}
	r.GET("/", rsc.Index)
	r.GET("/api/v1/user/id/:id", rsc.findByID)
	r.POST("/api/v1/user/save", rsc.save)
	return rsc
}

//Index ..
func (rsc *UserResource) Index(ctx *fasthttp.RequestCtx) {
	fmt.Fprint(ctx, "Welcome!\n")
}

//findByID ..
func (rsc *UserResource) findByID(ctx *fasthttp.RequestCtx) {
	id := fmt.Sprintf("%s", ctx.UserValue("id"))
	statusCode, response := rsc.userService.FindByID(id)
	ctx.SetContentType("application/json")
	ctx.SetBody(response)
	ctx.SetStatusCode(statusCode)
}

//save ..
func (rsc *UserResource) save(ctx *fasthttp.RequestCtx) {
	statusCode, response := rsc.userService.Save(ctx.Request.Body())
	ctx.SetContentType("application/json")
	ctx.SetBody(response)
	ctx.SetStatusCode(statusCode)
}
