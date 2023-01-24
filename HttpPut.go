//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
package libcore
//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
import (
	"io/ioutil"
	"bytes"
	"fmt"
	"net/http"
	"time"
)
//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
/**
 * @author Alexey Potehin <gnuplanet@gmail.com>, http://www.gnuplanet.online/doc/cv
 * @brief  send PUT request to url
 * @param  accessToken access token OR empty string
 * @param  url URL
 * @param  pHeaderMap header map, map[string]string{"Content-Type": "application/json", "Authorization": "Bearer xxx"}
 * @param  timeout timeout like 30 * time.Second
 * @param  dataIn data for send
 * @return httpCode http status code
 * @return dataOut received data
 * @return err error
 */
func HttpPut(url string, pHeaderMap *map[string]string, timeout time.Duration, dataIn []byte) (httpCode int, dataOut []byte, err error) {
//	log.Printf("PUT URL: %s\n", url)
//	log.Printf("PUT BODY: %s\n", string(dataIn))


	r := bytes.NewReader(dataIn)
	req, err := http.NewRequest(http.MethodPut, url, r)
	if err != nil {
		return
	}


	for headerKey, headerValue := range *pHeaderMap {
		req.Header.Add(headerKey, headerValue)
	}


	client := &http.Client{
		Timeout: timeout,
	}
	res, err := client.Do(req)
	if err != nil {
		return
	}
	defer res.Body.Close()


	dataOut, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}
	httpCode = res.StatusCode


	return
}
//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
