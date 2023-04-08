package helper

import "encoding/json"

func DecodeRequest[T any](request any) (T, error) {
	var decodedRequest T
	requestByte, err := json.Marshal(request)
	if err != nil {
		return decodedRequest, err
	}
	err = json.Unmarshal(requestByte, &decodedRequest)
	if err != nil {
		return decodedRequest, err
	}
	return decodedRequest, nil
}
