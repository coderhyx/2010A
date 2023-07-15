package asset

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"ucenter_api/internal/logic/asset"
	"ucenter_api/internal/svc"
	"ucenter_api/internal/types"
)

func WalletHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.WalletRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := asset.NewWalletLogic(r.Context(), svcCtx)
		resp, err := l.Wallet(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
