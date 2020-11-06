package notificationmodal

import "fmt"

func PermitParams(fields []string, params map[string]interface{}) map[string]interface{} {
	queryParams := make(map[string]interface{})
	for _, field := range fields {
		value := params[field]
		if value != nil {
			queryParams[field] = value
		}
	}
	return queryParams
}

func ValidatePresenceOfParams(queryParams map[string]interface{}, paramsToBeValidated ...string) (bool, string, string) {
	res := true
	var errorCode, errorDesc, paramStrValue string

	for _, param := range paramsToBeValidated {
		paramValue := queryParams[param]
		paramStrValue = fmt.Sprintf("%v", paramValue)
		if paramStrValue == "" || paramValue == nil {
			errorCode = "E002"
			errorDesc += "invalid " + param + " | "
			res = false
		}
	}

	return res, errorCode, errorDesc
}
