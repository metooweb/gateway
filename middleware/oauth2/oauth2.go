package oauth2

import (
	"github.com/urfave/negroni"
	"net/http"
	"metooweb-gateway/xcontext"
	"github.com/ory/hydra/sdk/go/hydra/swagger"
	"github.com/pkg/errors"
	"metooweb-gateway/xerrors"
	"fmt"
)

//token验证
func Authorizer() negroni.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

		var (
			err         error
			res         *swagger.OAuth2TokenIntrospection
			resp        *swagger.APIResponse
			accessToken = BearerTokenFromRequest(r)
		)

		fmt.Println("run")

		if accessToken == "" {
			err = errors.WithStack(xerrors.NewGateError(xerrors.ErrCodeMissingAccessToken, "missing access_token"))
		} else if res, resp, err = xcontext.GetHydraCodeGenSDK(r.Context()).IntrospectOAuth2Token(accessToken, ""); err != nil {
			//网络请求错误
			err = errors.WithStack(xerrors.NewGateError(xerrors.ErrCodeRemoteServiceError, "connect oauth sever error", "isp.remote-service-error").SetError(err))
		} else if resp.StatusCode != http.StatusOK {
			err = errors.WithStack(xerrors.NewGateError(xerrors.ErrCodeRemoteServiceError, "oauth server invalid status code", "isp.remote-service-error"))
		} else if !res.Active {
			err = errors.WithStack(xerrors.NewGateError(xerrors.ErrCodeInvalidAccessToken, "invalid access_token"))
		}

		if err != nil {
			panic(err)
		}

		fmt.Println(res)

		next(w, r)

	}

}
