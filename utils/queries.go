package utils

import "encoding/json"

func ParseQueries[T interface{}](queries map[string]string) (T, error) {
	result := new(T)
	jsonByte, err := json.Marshal(queries)
	if err != nil {
		return *result, err
	}
	err = json.Unmarshal(jsonByte, &result)
	if err != nil {
		return *result, err
	}
	return *result, nil
}
