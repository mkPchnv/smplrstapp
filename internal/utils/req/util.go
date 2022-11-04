package req

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
)

func GetParamFromRequest(request *http.Request, key string, message string) (string, error) {
	params := mux.Vars(request)
	if param, ok := params[key]; ok {
		return param, nil
	}

	return "", errors.New(message)
}

func GetModelFromBodyRequest(request *http.Request, model interface{}) error {
	decoder := json.NewDecoder(request.Body)
	if err := decoder.Decode(&model); err != nil {
		return errors.New(err.Error())
	}

	return nil
}

func GetQueryFromRequest(request *http.Request, key string, message string) (*[]string, error) {
	var result []string
	result, present := request.URL.Query()[key]
	if !present || len(result) == 0 {
		return &result, errors.New(message)
	}

	return &result, nil
}
