package response

type APIResponse struct {
	Success bool		`json:"success"`
	Massage string		`json:"message"`
	Data interface{} 	`json:"data"`
}