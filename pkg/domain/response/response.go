package response

type Response struct {
	Data   interface{} `json:"data"`
	Errors []string    `json:"errors"`
}
