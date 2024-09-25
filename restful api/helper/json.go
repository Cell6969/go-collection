package helper

import (
	"encoding/json"
	"net/http"
)

func ReadFromRequest(r *http.Request, result interface{}) {
	var decoder *json.Decoder = json.NewDecoder(r.Body)
	var err error = decoder.Decode(result)
	PanicIfError(err)
}

func WriteToResponse(w http.ResponseWriter, response interface{}) {
	w.Header().Set("Content-Type", "application/json")
	var encoder *json.Encoder = json.NewEncoder(w)
	var err error = encoder.Encode(response)
	PanicIfError(err)
}
