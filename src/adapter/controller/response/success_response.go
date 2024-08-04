package response

import "time"

type SuccessResponse struct {
	RequestID string `json:"request_id"`
	Timestamp string `json:"timestamp"`
	Result    any    `json:"result"`
}

func NewSuccessReponse(requestID string, result any) *SuccessResponse {
	return &SuccessResponse{
		RequestID: requestID,
		Timestamp: time.Now().Format("2006-01-02T15:04:05+09:00"),
		Result:    result,
	}
}
