package binding

import (

	"net/http"
	"net/url"


)

// BindQuery bind vars parameters to target.
func BindQuery(vars url.Values, target interface{}) error {

	return mapForm(target, vars)
}

// BindForm bind form parameters to target.
func BindForm(req *http.Request, target interface{}) error {
	if err := req.ParseForm(); err != nil {
		return err
	}

	return mapForm(target, req.Form)
}
