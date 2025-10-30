//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
package libcore
//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
import (
//	"log"
//	"encoding/json"
)
//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
type ResultStatusType struct {
	Code     uint64 `json:"code"`
	Message  string `json:"message"`
}
type ResultDataAnyType interface {}
//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
/*
{
	"status":
	{
		"code":    0,
		"message": "ok"
	},
	"data": {}
}
*/
type ResultEmpty struct {
	Status  ResultStatusType `json:"status"`
	Data    struct {}        `json:"data"`
}
func (p *ResultEmpty) SetOk() {
	p.Status.Code    = 0
	p.Status.Message = "ok"
}
func (p *ResultEmpty) SetErr(code uint64, msg string) {
	p.Status.Code    = code
	p.Status.Message = msg
}

// {"status":{"code":0,"message":"ok"},"data":{}}
//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
/*
{
	"status":
	{
		"code": 0,
		"message": "ok"
	},
	"data":
	{
		?
	}
}
*/
type ResultAny struct {
	Status  ResultStatusType  `json:"status"`
	Data    ResultDataAnyType `json:"data"`
}
func (p *ResultAny) SetOk() {
	p.Status.Code    = 0
	p.Status.Message = "ok"
//	p.Data           = ResultDataAnyType{}
//	p.Data           = struct{}{}
}
func (p *ResultAny) setErr(code uint64, msg string) {
	p.Status.Code    = code
	p.Status.Message = msg
}

// {"status":{"code":0,"message":"ok"},"data":null}
//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
/*
{
	"status":
	{
		"code": 0,
		"message": "ok"
	},
	"data":
	[
		{
			?
		}
	]
}
*/
type ResultAnyList struct {
	Status  ResultStatusType    `json:"status"`
	Data    []ResultDataAnyType `json:"data"`
}
func (p *ResultAnyList) SetOk() {
	p.Status.Code    = 0
	p.Status.Message = "ok"
	p.Data           = []ResultDataAnyType{}
}
func (p *ResultAnyList) SetErr(code uint64, msg string) {
	p.Status.Code    = code
	p.Status.Message = msg
}

// {"status":{"code":0,"message":"ok"},"data":[]}
//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
/*
{
	"status":
	{
		"code": 0,
		"message": "ok"
	},
	"data":
	[
		{
			?
		}
	],
	"meta":
	{
		"limit": ?,
		"next_key": ""
	}
}
*/
type ResultAnyListMeta struct {
	Status  ResultStatusType    `json:"status"`
	Data    []ResultDataAnyType `json:"data"`
	Meta struct {
		Limit   uint64 `json:"limit"`
		NextKey string `json:"next_key"`
	} `json:"meta"`
}
func (p *ResultAnyListMeta) SetOk() {
	p.Status.Code    = 0
	p.Status.Message = "ok"
	p.Data           = []ResultDataAnyType{}
	p.Meta.Limit     = 0
	p.Meta.NextKey   = ""
}
func (p *ResultAnyListMeta) SetErr(code uint64, msg string) {
	p.Status.Code    = code
	p.Status.Message = msg
}

// {"status":{"code":0,"message":"ok"},"data":[],"meta":{"limit":0,"next_key":""}}
//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
/*
func show(val interface{}) {
	var err error
	var dataOut []byte


	dataOut, err = json.Marshal(val)
	if err != nil {
		log.Printf("ERROR: %s", err)
		return
	}
	log.Printf("dataOut: %s\n", string(dataOut))
}

func main() {
	var x ResultEmpty
	x.SetOk()
	show(x)

	var y ResultAny
	y.SetOk()
	show(y)

	type Obj struct {
		A  string `json:"a"`
		B  int    `json:"b"`
	}
	var obj Obj
	y.Data = obj
	show(y)

	var z ResultAnyList
	z.SetOk()
	show(z)
}
*/
//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
