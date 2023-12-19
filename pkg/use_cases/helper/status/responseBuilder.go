package status

import (
	"encoding/json"
	"github.com/Lorahmir2024/lorahmir_library/pkg/domain/response"
)

func ResponseBuilder(status int, message string, data interface{}) ([]byte, error) {
	res := response.Response{
		Status:  status,
		Message: message,
		Data:    data,
	}

	marshalResponse, err := json.Marshal(res)
	if err != nil {
		return nil, err
	}
	return marshalResponse, nil

}
