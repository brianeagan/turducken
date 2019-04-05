package main

import (
	"encoding/json"
	"github.com/davecgh/go-spew/spew"
	"syscall/js"
)


type ethJSONRequest struct {
	JsonRPC string "2.0"
	Method string  `json:"method"`
	Params [2]string `json:"params"`
	Id		int `json:"id"`


}

var EthNode = "mainnet.infura.io/v3/0d3c8b422cbc4044a2c4bbcd1bef6b5f"
var Secret = "9e00254e641d403e94768bf19cbf011e"
var ProjectID = "0d3c8b422cbc4044a2c4bbcd1bef6b5f"

func findBalance(address string) ( string ){
	params := [2]string {address, "latest"}

	rpcReq := ethJSONRequest{
		Method: "eth_getBalance",
		Params: params,
		Id: 1,
	}
	spew.Dump(rpcReq)

	jsonReq, err := json.Marshal(rpcReq)
	if err != nil {
		println(err)
		return "err"
	}
	return string(jsonReq)
	//request, err := http.NewRequest("POST", EthNode, bytes.NewBuffer(jsonReq) )
	//spew.Dump(request)

}



func main(){
	spew.Dump("I'm a turtle, hi.")

	//make a wrapper for js to call findBalance
	//FuncOf is a go 1.12 breaking change from 1.11
	callBack := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		addr := js.Global().Get("document").Call("getElementById", "ethAddr").Get("value").String()
		loljson := findBalance(addr)
		js.Global().Get("document").Call("getElementById", "output").Set("value", loljson)
		return nil
	})

	//register the callback to a button push js function

	js.Global().Get("document").Call("getElementById", "runButton").
		Call("addEventListener", "click", callBack)

	c := make(chan struct{}, 0)
	spew.Dump("I'm a turtle, hi.")
	<- c


}
