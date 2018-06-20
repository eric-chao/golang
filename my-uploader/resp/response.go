package resp

import (
	"adhoc/adhoc_uploader/model"
	"encoding/json"
	"net/http"
	router "github.com/julienschmidt/httprouter"
)

func NoAuth(w http.ResponseWriter, r *http.Request, p router.Params) {
	RespResult := &model.Result{
		Code: 99999,
		Msg:  "with-no-auth: authentication required",
		Data: "",
	}
	result, _ := json.Marshal(RespResult)
	w.Write(result)
}

func NoAuthKey(w http.ResponseWriter, r *http.Request, p router.Params) {
	RespResult := &model.Result{
		Code: 99999,
		Msg:  "with-no-auth-key: Auth-Key required",
		Data: "",
	}
	result, _ := json.Marshal(RespResult)
	w.Write(result)
}

func ResponseBody(w http.ResponseWriter, code int32, msg string, data interface{}) {
	RespResult := &model.Result{
		Code: code,
		Msg:  msg,
		Data: data,
	}
	result, _ := json.Marshal(RespResult)
	w.Write(result)
}
