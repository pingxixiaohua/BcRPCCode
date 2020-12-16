package entity

/**
 * {
	"id":,
	"method":
	"jsonrpc":2.0,
	"params":
*/
type RPCRequest struct {
	Id      int64         `json:"id"`
	Method  string        `json:"method"`
	Jsonrpc string        `json:"jsonrpc"`
	Params  []interface{} `json:"params"`
}


