//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
package libcore
//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
import (
//	"database/sql"
//	_ "github.com/lib/pq"
	"bytes"
	"mime/multipart"
	"strings"
	"encoding/json"
	"strconv"
	"fmt"
//	"log"
//	"time"
//	"os"
	"io"
//	"flag"
	"net/http"
	"io/ioutil"
//	"errors"
//	"sync"
//	"container/list"
//	"context"
//	"math/rand"
	"reflect"
//	"github.com/gorilla/websocket"
//	"github.com/gorilla/mux"
//	"html"
//	"github.com/progman/libcore.go"
)
//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
func GetValueStr(r *http.Request, key string, defaultValue string) (value string) {
//	tmp := r.URL.Query().Get(key)
	tmp := r.URL.Query()[key]
	if (len(tmp) == 0) {
		value = defaultValue
		return
	}
	value = tmp[0]

	return
}
//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
func GetValueStrList(r *http.Request, key string, defaultValue []string) (value []string) {
	tmp := r.URL.Query()[key]
	if (len(tmp) == 0) {
		value = defaultValue
		return
	}

	for i := 0; i < len(tmp); i++ {
		if tmp[i] != "" {
			value = append(value, tmp[i])
		}
	}

	return
}
//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
func GetValueUint(r *http.Request, key string, defaultValue uint64) (value uint64) {
	var err error

	valueStr := GetValueStr(r, key, fmt.Sprintf("%d", defaultValue))

	if IsUint(valueStr) == false {
		value = defaultValue
		return
	}

	valueTmp, err := strconv.ParseUint(valueStr, 10, 64)
	if err != nil {
		value = defaultValue
		return
	}
	value = uint64(valueTmp)

	return
}
//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
func GetValueSint(r *http.Request, key string, defaultValue int64) (value int64) {
	var err error

	valueStr := GetValueStr(r, key, fmt.Sprintf("%d", defaultValue))

	if IsSint(valueStr) == false {
		value = defaultValue
		return
	}

	valueTmp, err := strconv.ParseInt(valueStr, 10, 64)
	if err != nil {
		value = defaultValue
		return
	}
	value = int64(valueTmp)

	return
}
//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
func PostParseMultipart(r *http.Request, maxMemory int64) (err error) {

	if len(r.Form) == 0 {
		err = r.ParseMultipartForm(maxMemory)
		if err != nil {
			return
		}
	}

	return
}
//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
func PostParse(r *http.Request) (err error) {

	if len(r.Form) != 0 {
		return
	}

	if strings.Index(strings.ToLower(r.Header.Get("Content-Type")), "multipart") == -1 {
		err = r.ParseForm()
		if err != nil {
			return
		}
	} else {
		err = PostParseMultipart(r, 1014 * 1024 * 100)
		if err != nil {
			return
		}
	}

	return
}
//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
func PostValueStr(r *http.Request, key string, defaultValue string) (value string) {
	var err error

	err = PostParse(r)
	if err != nil {
		value = defaultValue
		return
	}
	value = r.PostForm.Get(key)

	return
}
//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
func PostValueUint(r *http.Request, key string, defaultValue uint64) (value uint64) {
	var err error

	valueStr := PostValueStr(r, key, fmt.Sprintf("%d", defaultValue))

	if IsUint(valueStr) == false {
		value = defaultValue
		return
	}

	valueTmp, err := strconv.ParseUint(valueStr, 10, 64)
	if err != nil {
		value = defaultValue
		return
	}
	value = uint64(valueTmp)

	return
}
//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
func PostValueSint(r *http.Request, key string, defaultValue int64) (value int64) {
	var err error

	valueStr := PostValueStr(r, key, fmt.Sprintf("%d", defaultValue))

	if IsSint(valueStr) == false {
		value = defaultValue
		return
	}

	valueTmp, err := strconv.ParseInt(valueStr, 10, 64)
	if err != nil {
		value = defaultValue
		return
	}
	value = int64(valueTmp)

	return
}
//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
/*
func MPostValueStr(r *http.Request, key string, defaultValue string) (value string) {
	var err error

	if len(r.Form) == 0 {
		err = r.ParseMultipartForm(1024 * 1024 * 10)
		if err != nil {
			value = defaultValue
			return
		}
	}

//	value = r.PostForm.Get(key)
	value = r.Form.Get(key)

//	tmp := r.URL.Query()[key]
//	if (len(tmp) == 0) {
//		value = defaultValue
//		return
//	}
//	value = tmp[0]

	return
}
*/
//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
/*
func MPostValueUint(r *http.Request, key string, defaultValue uint64) (value uint64) {
	var err error

	valueStr := MPostValueStr(r, key, fmt.Sprintf("%d", defaultValue))

	if IsUint(valueStr) == false {
		value = defaultValue
		return
	}

	valueTmp, err := strconv.ParseUint(valueStr, 10, 64)
	if err != nil {
		value = defaultValue
		return
	}
	value = uint64(valueTmp)

	return
}
*/
//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
/*
func MPostValueSint(r *http.Request, key string, defaultValue int64) (value int64) {
	var err error

	valueStr := MPostValueStr(r, key, fmt.Sprintf("%d", defaultValue))

	if IsSint(valueStr) == false {
		value = defaultValue
		return
	}

	valueTmp, err := strconv.ParseInt(valueStr, 10, 64)
	if err != nil {
		value = defaultValue
		return
	}
	value = int64(valueTmp)

	return
}
*/
//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
func PostValueFile(r *http.Request, key string, defaultValue []byte) (filename string, value []byte) {
	var err error

	err = PostParse(r)
	if err != nil {
		value = defaultValue
		return
	}


	var f multipart.File
	var fHeader *multipart.FileHeader
	f, fHeader, err = r.FormFile(key)
	if err != nil {
		value = defaultValue
		return
	}
	defer f.Close()


	filename = strings.TrimSpace(fHeader.Filename)


//	value, _ = ioutil.ReadAll(f)
	var buf bytes.Buffer
	io.Copy(&buf, f)
	value = buf.Bytes()


	return
}
//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
func PostBodyStr(r *http.Request, defaultValue string) (value string) {
	var err error
	var body []byte


	body, err = ioutil.ReadAll(r.Body)
	if err != nil {
		value = defaultValue
		return
	}


	value = string(body)
	return
}
//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
func PostJsonValueStrList(body string, key string, defaultValue []string) (value []string) {
	var err error
	var result map[string]interface{}


	if body == "" {
		value = defaultValue
		return
	}


	err = json.Unmarshal([]byte(body), &result)
	if err != nil {
		value = defaultValue
		return
	}


	i, ok := result[key]
	if ok == false {
		value = defaultValue
		return
	}
	if i == nil {
		value = defaultValue
		return
	}


	if reflect.TypeOf(i).String() != "[]string" {
		value = defaultValue
		return
	}
	value = i.([]string)


	return
}
//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
/*
func JPostValueUint(body string, key string, defaultValue int) (value int) {
//	var err error


//	valueStr := JPostValueStr(body, key, fmt.Sprintf("%d", defaultValue))


//	if IsUint(valueStr) == false {
//		value = defaultValue
//		return
//	}


//	valueTmp, err := strconv.ParseUint(valueStr, 10, 64)
//	if err != nil {
//		value = defaultValue
//		return
//	}
//	value = int(valueTmp)


//	return


	var err error
	var result map[string]interface{}


	if body == "" {
		value = defaultValue
		return
	}


	err = json.Unmarshal([]byte(body), &result)
	if err != nil {
		value = defaultValue
		return
	}


	i, ok := result[key]
	if ok == false {
		value = defaultValue
		return
	}


	if reflect.TypeOf(i).String() != "int" {
		value = defaultValue
		return
	}
	value = i.(int)


	if value < 0 {
		value = defaultValue
		return
	}


	return
}
*/
//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
/*
func JPostValueSint(body string, key string, defaultValue int) (value int) {
//	var err error


//	valueStr := JPostValueStr(body, key, fmt.Sprintf("%d", defaultValue))


//	if IsSint(valueStr) == false {
//		value = defaultValue
//		return
//	}


//	valueTmp, err := strconv.ParseInt(valueStr, 10, 64)
//	if err != nil {
//		value = defaultValue
//		return
//	}
//	value = int(valueTmp)


//	return


	var err error
	var result map[string]interface{}


	if body == "" {
		value = defaultValue
		return
	}


	err = json.Unmarshal([]byte(body), &result)
	if err != nil {
		value = defaultValue
		return
	}


	i, ok := result[key]
	if ok == false {
		value = defaultValue
		return
	}


	if reflect.TypeOf(result).String() != "int" {
		value = defaultValue
		return
	}
	value = i.(int)


	return
}
*/
//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
