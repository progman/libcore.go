//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
package libcore
//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
import (
	"io/ioutil"
	"bytes"
//	"fmt"
	"net/http"
	"time"
)
//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
/**
 * @author Alexey Potehin <gnuplanet@gmail.com>, http://www.gnuplanet.online/doc/cv
 * @brief  send POST request to url
 * @param  accessToken access token OR empty string
 * @param  url URL
 * @param  pHeaderMap header map, map[string]string{"Content-Type": "application/json", "Authorization": "Bearer xxx"}
 * @param  timeout timeout like 30 * time.Second
 * @param  dataIn data for send
 * @return httpCode http status code
 * @return dataOut received data
 * @return err error
 */
func HttpPost(url string, pHeaderMap *map[string]string, timeout time.Duration, dataIn []byte) (httpCode int, dataOut []byte, cookieList []*http.Cookie, err error) {
	var req *http.Request
//	log.Printf("POST URL: %s\n", url)
//	log.Printf("POST BODY: %s\n", string(dataIn))


	r := bytes.NewReader(dataIn)
	req, err = http.NewRequest(http.MethodPost, url, r)
	if err != nil {
		return
	}


	if pHeaderMap != nil {
		for headerKey, headerValue := range *pHeaderMap {
			req.Header.Add(headerKey, headerValue)
		}
	}


	client := &http.Client{
		Timeout: timeout,
	}
	res, err := client.Do(req)
	if err != nil {
		return
	}
	defer res.Body.Close()


	cookieList = res.Cookies()


	dataOut, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}
	httpCode = res.StatusCode


	return
}
//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
