package user

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"ucenter_api/internal/logic/user"
	"ucenter_api/internal/svc"
	"ucenter_api/internal/types"
)

func FindOneHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FindOneRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewFindOneLogic(r.Context(), svcCtx)
		resp, err := l.FindOne(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
