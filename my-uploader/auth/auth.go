package auth

import (
	"net/http"
	"fmt"
	router "github.com/julienschmidt/httprouter"
	. "adhoc/adhoc_uploader/db"
	. "adhoc/adhoc_uploader/resp"
)

var WriteRoles = map[string]string{
	"Admin": "Admin",
	"Owner": "Owner",
}

func ExtractUserId(authKey string) (string, error) {
	redisKey := fmt.Sprintf("adhoc_auth_%s", authKey)
	return Get(redisKey)
}

func ExtractRoles(appId string, userId string) (string, error) {
	if appId == "" {
		appId = "ics"
	}
	redisKey := fmt.Sprintf("adhoc_appuser_%s", appId)
	field := userId
	return HGet(redisKey, field)
}

func Auth(h router.Handle) router.Handle {
	return func(w http.ResponseWriter, r *http.Request, p router.Params) {
		//set response content-type to 'application/json'
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Auth-Key")
		authKey := r.Header.Get("Auth-Key")
		if authKey == "" {
			NoAuthKey(w, r, p)
			return
		}
		appId := p.ByName("appId")
		userId, err := ExtractUserId(authKey)
		if err != nil {
			fmt.Errorf("error: %v", err)
		}
		roles, err := ExtractRoles(appId, userId)
		if err != nil {
			fmt.Errorf("error: %v", err)
		}

		if _, ok := WriteRoles[roles]; ok {
			h(w, r, p)
			return
		}
		//with-no-auth
		NoAuth(w, r, p)
	}
}
