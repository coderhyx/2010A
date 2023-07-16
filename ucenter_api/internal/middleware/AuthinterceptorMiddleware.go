package middleware

import (
	common "2010A/common/ctl"
	"2010A/common/feature"
	"context"
	"net/http"
	"ucenter_api/internal/config"

	"github.com/zeromicro/go-zero/rest/httpx"
)

type AuthInterceptorMiddleware struct {
	Config config.Config
}

func NewAuthInterceptorMiddleware(c config.Config) *AuthInterceptorMiddleware {
	return &AuthInterceptorMiddleware{
		Config: c,
	}
}

func (m *AuthInterceptorMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		result := common.NewResult()
		token := r.Header.Get("x-auth-token")
		if token == "" {
			result.Code = 10001
			result.Message = "token为空,请在接口的header填写"
			httpx.WriteJson(w, 200, result)
			return
		}

		userId, err := feature.ParseToken(token, m.Config.JWT.AccessSecret)
		if err != nil {
			result.Code = 10002
			result.Message = "token没有校验通过"
			httpx.WriteJson(w, 200, result)
			return
		}
		ctx := r.Context()
		ctx = context.WithValue(ctx, "userId", userId)
		r = r.WithContext(ctx)
		next(w, r)
	}
}
