package encoding

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

func WriteJSONResult(
	w http.ResponseWriter,
	statusCode int,
	data interface{},
) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error":"could not encode response"}`))
		return err
	}

	w.WriteHeader(statusCode)
	_, err = io.Copy(
		w,
		bytes.NewReader(jsonData),
	)
	return err
}
