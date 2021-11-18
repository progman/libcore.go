//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
package libcore
//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
import (
//	"database/sql"
//	_ "github.com/lib/pq"
//	"bytes"
//	"mime/multipart"
//	"strings"
	"encoding/json"
	"strconv"
	"fmt"
//	"log"
//	"time"
//	"os"
//	"flag"
	"net/http"
	"io/ioutil"
//	"errors"
//	"sync"
//	"container/list"
//	"context"
//	"math/rand"
//	"github.com/gorilla/websocket"
//	"github.com/gorilla/mux"
//	"html"
//	"github.com/progman/libcore.go"
)
//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
// r.URL.Query().Get("param1")
func GetValueStr(r *http.Request, key string, defaultValue string) (value string) {
	tmp := r.URL.Query()[key]
	if (len(tmp) == 0) {
		value = defaultValue
		return
	}
	value = tmp[0]

	return
}
//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
func GetValueUint(r *http.Request, key string, defaultValue int) (value int) {
	var err error

	valueStr := GetValueStr(r, key, fmt.Sprintf("%d", defaultValue))
//	tmp := r.URL.Query()[key]
//	if (len(tmp) == 0) {
//		value = defaultValue
//		return
//	}
//	valueStr := tmp[0]

	if IsUint(valueStr) == false {
		value = defaultValue
		return
	}

	valueTmp, err := strconv.ParseUint(valueStr, 10, 64)
	if err != nil {
		value = defaultValue
		return
	}
	value = int(valueTmp)

	return
}
//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
func GetValueSint(r *http.Request, key string, defaultValue int) (value int) {
	var err error

	valueStr := GetValueStr(r, key, fmt.Sprintf("%d", defaultValue))
//	tmp := r.URL.Query()[key]
//	if (len(tmp) == 0) {
//		value = defaultValue
//		return
//	}
//	valueStr := tmp[0]

	if IsSint(valueStr) == false {
		value = defaultValue
		return
	}

	valueTmp, err := strconv.ParseInt(valueStr, 10, 64)
	if err != nil {
		value = defaultValue
		return
	}
	value = int(valueTmp)

	return
}
//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
// r.URL.Query().Get("param1")
func PostValueStr(r *http.Request, key string, defaultValue string) (value string) {
	var err error

	err = r.ParseForm()
	if err != nil {
		value = defaultValue
		return
	}

	value = r.PostForm.Get(key)


//	tmp := r.URL.Query()[key]
//	if (len(tmp) == 0) {
//		value = defaultValue
//		return
//	}
//	value = tmp[0]

	return
}
//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
func PostValueUint(r *http.Request, key string, defaultValue int) (value int) {
	var err error

	valueStr := PostValueStr(r, key, fmt.Sprintf("%d", defaultValue))
//	tmp := r.URL.Query()[key]
//	if (len(tmp) == 0) {
//		value = defaultValue
//		return
//	}
//	valueStr := tmp[0]

	if IsUint(valueStr) == false {
		value = defaultValue
		return
	}

	valueTmp, err := strconv.ParseUint(valueStr, 10, 64)
	if err != nil {
		value = defaultValue
		return
	}
	value = int(valueTmp)

	return
}
//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
func PostValueSint(r *http.Request, key string, defaultValue int) (value int) {
	var err error

	valueStr := PostValueStr(r, key, fmt.Sprintf("%d", defaultValue))
//	tmp := r.URL.Query()[key]
//	if (len(tmp) == 0) {
//		value = defaultValue
//		return
//	}
//	valueStr := tmp[0]

	if IsSint(valueStr) == false {
		value = defaultValue
		return
	}

	valueTmp, err := strconv.ParseInt(valueStr, 10, 64)
	if err != nil {
		value = defaultValue
		return
	}
	value = int(valueTmp)

	return
}
//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
// r.URL.Query().Get("param1")
func MPostValueStr(r *http.Request, key string, defaultValue string) (value string) {
	var err error

//	err = r.ParseForm()
	err = r.ParseMultipartForm(1024 * 1024)
	if err != nil {
		value = defaultValue
		return
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
//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
func MPostValueUint(r *http.Request, key string, defaultValue int) (value int) {
	var err error

	valueStr := MPostValueStr(r, key, fmt.Sprintf("%d", defaultValue))
//	tmp := r.URL.Query()[key]
//	if (len(tmp) == 0) {
//		value = defaultValue
//		return
//	}
//	valueStr := tmp[0]

	if IsUint(valueStr) == false {
		value = defaultValue
		return
	}

	valueTmp, err := strconv.ParseUint(valueStr, 10, 64)
	if err != nil {
		value = defaultValue
		return
	}
	value = int(valueTmp)

	return
}
//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
func MPostValueSint(r *http.Request, key string, defaultValue int) (value int) {
	var err error

	valueStr := MPostValueStr(r, key, fmt.Sprintf("%d", defaultValue))
//	tmp := r.URL.Query()[key]
//	if (len(tmp) == 0) {
//		value = defaultValue
//		return
//	}
//	valueStr := tmp[0]

	if IsSint(valueStr) == false {
		value = defaultValue
		return
	}

	valueTmp, err := strconv.ParseInt(valueStr, 10, 64)
	if err != nil {
		value = defaultValue
		return
	}
	value = int(valueTmp)

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
func JPostValueStr(r *http.Request, key string, defaultValue string) (value string) {
	var err error
	var body
	var result map[string]interface{}


	body = PostBodyStr(r, "")
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


	value = i.(string)


	return
}
//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
func JPostValueUint(r *http.Request, key string, defaultValue int) (value int) {
	var err error


	valueStr := JPostValueStr(r, key, fmt.Sprintf("%d", defaultValue))


	if IsUint(valueStr) == false {
		value = defaultValue
		return
	}


	valueTmp, err := strconv.ParseUint(valueStr, 10, 64)
	if err != nil {
		value = defaultValue
		return
	}
	value = int(valueTmp)


	return
}
//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
func JPostValueSint(r *http.Request, key string, defaultValue int) (value int) {
	var err error


	valueStr := JPostValueStr(r, key, fmt.Sprintf("%d", defaultValue))


	if IsSint(valueStr) == false {
		value = defaultValue
		return
	}


	valueTmp, err := strconv.ParseInt(valueStr, 10, 64)
	if err != nil {
		value = defaultValue
		return
	}
	value = int(valueTmp)


	return
}
//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
