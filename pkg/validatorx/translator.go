package validatorx

import (
	"net/http"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

func TranslateErrorMessage(err error, obj any) (map[string]string, int) {
	errorsMap := make(map[string]string)
	statusCode := http.StatusBadRequest

	if err == nil {
		return nil, http.StatusOK
	}

	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, e := range validationErrors {
			field := getJSONTag(obj, e.StructField())
			target := getJSONTag(obj, e.Param())

			readableField := strings.ReplaceAll(field, "_", " ")
			readableTarget := strings.ReplaceAll(target, "_", " ")

			errorsMap[field] = getErrorMsg(e, readableField, readableTarget)
		}
		return errorsMap, statusCode
	}

	errorsMap["error"] = err.Error()
	return errorsMap, http.StatusInternalServerError
}

func getJSONTag(obj any, structField string) string {
	objType := reflect.TypeOf(obj).Elem()
	if f, ok := objType.FieldByName(structField); ok {
		jsonTag := strings.Split(f.Tag.Get("json"), ",")[0]
		if jsonTag != "" && jsonTag != "-" {
			return jsonTag
		}
	}
	return structField
}

func getErrorMsg(fe validator.FieldError, field, target string) string {
	switch fe.Tag() {
	case "required":
		return field + " is required"
	case "email":
		return field + " must be a valid email address"
	case "min":
		return field + " must be at least " + target + " characters"
	case "max":
		return field + " must be at most " + target + " characters"
	case "eqfield":
		return field + " must be equal to " + target
	default:
		return field + " is invalid"
	}
}
