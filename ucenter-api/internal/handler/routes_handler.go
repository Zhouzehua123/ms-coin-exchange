// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package handler

import (
	"ucenter-api/internal/svc"
)

func RegisterHandlers(r *Routers, serverCtx *svc.ServiceContext) {
	register := NewRegisterHandler(serverCtx)
	registerRouter := r.Group()
	registerRouter.Get("/uc/register/phone", register.Register)
}
