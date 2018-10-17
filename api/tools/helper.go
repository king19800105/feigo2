package tools

import (
	"github.com/gorilla/schema"
	"net/http"
)

func DecodeFormData(r *http.Request, req interface{}) error {
	err := r.ParseForm()

	if nil != err {
		return err
	}

	err = schema.NewDecoder().Decode(req, r.Form)

	return err
}
