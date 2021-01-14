package libcore


import "fmt"
import "runtime"


type Result struct {
	flagErr            bool
	errNum             int
	errMsg             string

	curentFunctionName string
	curentFileName     string
	curentLine         int

	parentFunctionName string
	parentFileName     string
	parentLine         int
}


func (p *Result) GetPlace(stackPosition int) (functionName string, fileName string, line int) {
	functionName = "unknown"
	fileName     = "unknown"
	line         = -1

	for {
		pc, file, no, ok := runtime.Caller(stackPosition)
		if ok == false {
			break
		}
		fileName = file
		line     = no


		me := runtime.FuncForPC(pc)
		if me == nil {
			functionName = "unnamed"
			break
		}
		functionName = me.Name()
		break
	}

	return functionName, fileName, line
}


func (p *Result) Init() {
	p.flagErr = true
	p.errNum  = 1
	p.errMsg  = "unknown"

	p.parentFunctionName, p.parentFileName, p.parentLine = p.GetPlace(3)
	p.curentFunctionName, p.curentFileName, p.curentLine = p.GetPlace(2)
}


func (p *Result) GetCurentFunctionName() string {
	return p.curentFunctionName
}


func (p *Result) GetCurentFileName() string {
	return p.curentFileName
}


func (p *Result) GetCurentLine() int {
	return p.curentLine
}


func (p *Result) GetParentFunctionName() string {
	return p.parentFunctionName
}


func (p *Result) GetParentFileName() string {
	return p.parentFileName
}


func (p *Result) GetParentLine() int {
	return p.parentLine
}


func (p *Result) SetOk() {
	p.flagErr = false
	p.errNum  = 0
	p.errMsg  = "ok"

	p.curentFunctionName, p.curentFileName, p.curentLine = p.GetPlace(2)
}


func (p *Result) SetErrMsg(errNum int, errMsg string) {
	p.flagErr = true
	p.errNum  = errNum
	p.errMsg  = errMsg

	p.curentFunctionName, p.curentFileName, p.curentLine = p.GetPlace(2)
}


func (p *Result) SetErr(err error) {
	p.SetErrMsg(1, fmt.Sprintf("%v", err))
}


func (p *Result) IsOk() bool {
	if p.flagErr == false {
		return true
	}
	return false
}


func (p *Result) IsErr() bool {
	return (p.IsOk() == false)
}


func (p *Result) GetErrNum() int {
	return p.errNum
}


func (p *Result) GetErrMsg() string {
	return p.errMsg
}


func (p *Result) Draw1() string {
	return fmt.Sprintf("ERROR[%s()]: %s", p.GetCurentFunctionName(), p.GetErrMsg())
}


func (p *Result) Draw2() string {
	return fmt.Sprintf("ERROR[%s(%s#%d)]: %s", p.GetCurentFunctionName(), p.GetCurentFileName(), p.GetCurentLine(), p.GetErrMsg())
}


func (p *Result) Draw3() string {
	return fmt.Sprintf("ERROR[%s(%s#%d) >> %s(%s#%d)]: %s", p.GetParentFunctionName(), p.GetParentFileName(), p.GetParentLine(), p.GetCurentFunctionName(), p.GetCurentFileName(), p.GetCurentLine(), p.GetErrMsg())
}

/*
func zzz(val int) Result {
	var result Result
	result.Init()

	if val == 1 {
		result.SetErrMsg(1, "fucked up situation")
	}

	if val == 2 {
		result.SetOk()
	}

	return result
}


func main() {
	var rc Result

	rc = zzz(0)
	if rc.IsOk() == false {
		fmt.Printf("%s\n", rc.Draw1())
	}

	rc = zzz(1)
	if rc.IsOk() == false {
		fmt.Printf("%s\n", rc.Draw2())
	}

	rc = zzz(1)
	if rc.IsOk() == false {
		fmt.Printf("%s\n", rc.Draw3())
	}

	rc = zzz(2)
	fmt.Printf("%s\n", rc.Draw3())
}
*/
