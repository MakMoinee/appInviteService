package response

import (
	"encoding/json"
	"net/http"

	"github.com/MakMoinee/appInviteService/internal/appInviteService/common"
)

type ErrorResponse struct {
	ErrorMessage string `json:"errorMessage"`
	ErrorCode    int    `json:"errorCode"`
}

// Success() - returns success response
func Success(w http.ResponseWriter, payload interface{}) {
	byteResp, err := json.Marshal(payload)

	if err != nil {
		errBuilder := ErrorResponse{}
		errBuilder.ErrorCode = http.StatusInternalServerError
		errBuilder.ErrorMessage = err.Error()
		Error(w, errBuilder)
		return
	}

	w.Header().Set(common.ContentTypeKey, common.ContentTypeValue)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(byteResp))
}

// Error() - returns error response
func Error(w http.ResponseWriter, payload ErrorResponse) {
	result, _ := json.Marshal(payload)
	w.Header().Set(common.ContentTypeKey, common.ContentTypeValue)
	w.WriteHeader(payload.ErrorCode)
	w.Write([]byte(result))
}
