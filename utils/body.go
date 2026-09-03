package utils

import "encoding/json"

func ParseBody[T interface{}](bodyBytes []byte) (T, error) {
	result := new(T)
	err := json.Unmarshal(bodyBytes, &result)
	if err != nil {
		return *result, err
	}
	return *result, nil
}
