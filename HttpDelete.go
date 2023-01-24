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
 * @brief  send DELETE request to url
 * @param  url URL
 * @param  pHeaderMap header map, map[string]string{"Content-Type": "application/json", "Authorization": "Bearer xxx"}
 * @param  timeout timeout like 30 * time.Second
 * @return httpCode http status code
 * @return dataOut received data
 * @return err error
 */
func HttpDelete(url string, pHeaderMap *map[string]string, timeout time.Duration) (httpCode int, dataOut []byte, err error) {
//	log.Printf("DELETE URL: %s\n", url)


	r := bytes.NewReader([]byte{})
	req, err := http.NewRequest(http.MethodDelete, url, r)
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
