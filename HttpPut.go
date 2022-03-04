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
 * @brief  send PUT request to url
 * @param  accessToken access token OR empty string
 * @param  url URL
 * @param  dataIn data for send
 * @param  contentType content type OR empty string
 * @param  timeout timeout like 30 * time.Second
 * @return httpCode http status code
 * @return dataOut received data
 * @return err error
 */
func HttpPut(accessToken string, url string, dataIn []byte, contentType string, timeout time.Duration) (httpCode int, dataOut []byte, err error) {
//	log.Printf("POST URL: %s\n", url)
//	log.Printf("POST BODY: %s\n", string(dataIn))


	r := bytes.NewReader(dataIn)
	req, err := http.NewRequest(http.MethodPost, url, r)
	if err != nil {
		return
	}


	if accessToken != "" {
		req.Header.Add("Authorization", "Bearer " + accessToken)
	}


//log.Printf("contentType: \"%s\"\n", contentType)
	if contentType != "" {
		req.Header.Add("Content-Type", contentType)
	}
	req.Header.Add("Accept", "*/*")


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
