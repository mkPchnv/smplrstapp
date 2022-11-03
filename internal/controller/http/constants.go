package http

import "fmt"

const (
	apiVersion           = "/api/v1"
	idParam              = "id"
	limitQueryParam      = "limit"
	offsetQueryParam     = "offset"
	emptyMessage         = ""
	methodGet            = "GET"
	methodPost           = "POST"
	getUserUrl           = "users/{id}"
	getUsersUrl          = "users"
	addUserUrl           = "users"
	addActivityUrl       = "activities"
	requestWithoutParams = "sorry, request without params, use another url"
	notFoundUser         = "not found user by identificator"
	errInvalidBody       = "not valid body of request"
)

func check(err error) bool {
	if err != nil {
		// todo logging
		fmt.Println(err.Error())
		return true
	}

	return false
}
