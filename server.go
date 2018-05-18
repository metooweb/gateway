package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"github.com/urfave/negroni"
	"github.com/metooweb/gateway/middleware/oauth2"
	"fmt"
	"github.com/metooweb/gateway/xcontext"
	"github.com/metooweb/gateway/storage"
	"github.com/pkg/errors"
	"github.com/metooweb/gateway/xerrors"
	"github.com/go-resty/resty"
	"encoding/json"
	"github.com/metooweb/gateway/backend"
)

func initServer() (err error) {

	router := mux.NewRouter()

	router.HandleFunc("/api", func(writer http.ResponseWriter, request *http.Request) {

		fmt.Fprint(writer, "...........")

	})

	n := negroni.New()

	n.Use(negroni.NewRecovery())
	n.Use(negroni.NewLogger())

	n.Use(func() negroni.HandlerFunc {

		return func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

			ctx := xcontext.InitHydraCodeGenSDK(r.Context(), hydraSDK)
			r = r.WithContext(ctx)
			next(rw, r)

		}
	}())

	n.Use(oauth2.Authorizer())

	//转发
	n.Use(func() negroni.HandlerFunc {

		return func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

			var (
				err  error
				res  = new(backend.Resp)
				resp *resty.Response
			)

			if err = r.ParseForm(); err != nil {
				http.Error(rw, "bad request", http.StatusBadRequest)
			}

			method := r.Form.Get("method")

			st := storage.Storage{}

			api := st.GetApi(method)

			if api == nil {
				err = errors.WithStack(xerrors.NewGateError(xerrors.ErrCodeInvalidMethod))
				panic(err)
			}

			resp, err = resty.R().SetMultiValueFormData(r.PostForm).Post(api.Client.BackendURL)

			if err != nil {
				err = errors.WithStack(xerrors.NewGateError(xerrors.ErrCodeServiceCurrentlyUnavailable).SetError(err))
				panic(err)
			}

			if err = json.Unmarshal(resp.Body(), res); err != nil {
				panic(err)
			}

			if res.Code == "SUCCESS" {

				resBytes, _ := json.Marshal(map[string]interface{}{"code": 0, "result": res.Result})

				rw.Write(resBytes)

			} else {

				panic(errors.WithStack(xerrors.NewGateError(xerrors.ErrCodeRemoteServiceError, "remote service error", res.Code, res.Msg)))
			}

		}

	}())

	n.UseHandler(router)
	n.Run(config.Addr)

	return
}
