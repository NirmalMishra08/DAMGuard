package utils

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"reflect"
)

func ReadJsonAndValidate(w http.ResponseWriter, r *http.Request, data any) error {
	err := readJsonFromBody(w, r, data)
	if err != nil {
		return err
	}

	return nil
}

type jsonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

// readJsonFromBody reads JSON from the request body
func readJsonFromBody(w http.ResponseWriter, r *http.Request, data any) error {
	// TODO - revert to 1 MB
	maxBytes := 10 << 20 // one megabyte

	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	dec := json.NewDecoder(r.Body)
	err := dec.Decode(data)
	if err != nil {
		log.Println(err)
		log.Println(reflect.TypeOf(err))
		return err
	}

	err = dec.Decode(&struct{}{})
	if err != io.EOF {
		return errors.New("body must only have a single JSON value")
	}

	return nil
}

func WriteJson(w http.ResponseWriter, status int, data any, headers ...http.Header) {
	out, err := json.Marshal(data)
	if err != nil {
		log.Println("unable to marshal json:", err)
	}

	if len(headers) > 0 {
		for key, value := range headers[0] {
			w.Header()[key] = value
		}
	}

	w.Header().Set("Content-type", "application-json")
	w.WriteHeader(status)

	_, err = w.Write(out)
	if err != nil {
		log.Println("unable to write json:", err)
	}

}

func ErrJson(w http.ResponseWriter, err error) {
	var payload jsonResponse

	payload.Error = true
	payload.Message = err.Error()

	statusCode, exists := CustomErrorType[err]
	if exists {
		WriteJson(w, statusCode, payload)
		return
	}
	WriteJson(w, http.StatusBadRequest, payload)
}
