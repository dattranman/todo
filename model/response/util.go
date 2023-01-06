package response

const (
	StatusSuccess = "success"
	StatusFail    = "fail"
)

type UtilResponse struct {
	Data   any    `json:"data"`
	Status string `json:"status"`
	Msg    string `json:"msg"`
	Code   int16  `json:"code,omitempty"`
}
