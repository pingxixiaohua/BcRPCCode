package entity

type RPCResult struct {
	Id int			`json:"id"`
	Error string	`json:"error"`
	Result interface{}	`json:"result"`
}


