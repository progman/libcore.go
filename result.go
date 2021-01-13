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


func (p *Result) getPlace(stackPosition int) (functionName string, fileName string, line int) {
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


func (p *Result) init() {
	p.flagErr = true
	p.errNum  = 1
	p.errMsg  = "unknown"

	p.parentFunctionName, p.parentFileName, p.parentLine = p.getPlace(3)
	p.curentFunctionName, p.curentFileName, p.curentLine = p.getPlace(2)
}


func (p *Result) getCurentFunctionName() string {
	return p.curentFunctionName
}


func (p *Result) getCurentFileName() string {
	return p.curentFileName
}


func (p *Result) getCurentLine() int {
	return p.curentLine
}


func (p *Result) getParentFunctionName() string {
	return p.parentFunctionName
}


func (p *Result) getParentFileName() string {
	return p.parentFileName
}


func (p *Result) getParentLine() int {
	return p.parentLine
}


func (p *Result) setOk() {
	p.flagErr = false
	p.errNum  = 0
	p.errMsg  = "ok"

	p.curentFunctionName, p.curentFileName, p.curentLine = p.getPlace(2)
}


func (p *Result) setErrMsg(errNum int, errMsg string) {
	p.flagErr = true
	p.errNum  = errNum
	p.errMsg  = errMsg

	p.curentFunctionName, p.curentFileName, p.curentLine = p.getPlace(2)
}


func (p *Result) setErr(err error) {
	p.setErrMsg(1, fmt.Sprintf("%v", err))
}


func (p *Result) isOk() bool {
	if p.flagErr == false {
		return true
	}
	return false
}


func (p *Result) getErrNum() int {
	return p.errNum
}


func (p *Result) getErrMsg() string {
	return p.errMsg
}


func (p *Result) draw1() string {
	return fmt.Sprintf("ERROR[%s()]: %s", p.getCurentFunctionName(), p.getErrMsg())
}


func (p *Result) draw2() string {
	return fmt.Sprintf("ERROR[%s(%s#%d)]: %s", p.getCurentFunctionName(), p.getCurentFileName(), p.getCurentLine(), p.getErrMsg())
}


func (p *Result) draw3() string {
	return fmt.Sprintf("ERROR[%s(%s#%d) >> %s(%s#%d)]: %s", p.getParentFunctionName(), p.getParentFileName(), p.getParentLine(), p.getCurentFunctionName(), p.getCurentFileName(), p.getCurentLine(), p.getErrMsg())
}

/*
func zzz(val int) Result {
	var result Result
	result.init()

	if val == 1 {
		result.setErrMsg(1, "fucked up situation")
	}

	if val == 2 {
		result.setOk()
	}

	return result
}


func main() {
	var rc Result

	rc = zzz(0)
	if rc.isOk() == false {
		fmt.Printf("%s\n", rc.draw1())
	}

	rc = zzz(1)
	if rc.isOk() == false {
		fmt.Printf("%s\n", rc.draw2())
	}

	rc = zzz(1)
	if rc.isOk() == false {
		fmt.Printf("%s\n", rc.draw3())
	}

	rc = zzz(2)
	fmt.Printf("%s\n", rc.draw3())
}
*/
