//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
package libcore
//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
import (
"strings"
)
//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
func UrlAddParam(url string, key string, value string) (urlNew string) {
	urlNew = url

	if strings.Index(url, "?") == -1 {
		urlNew = UrlAddSlash(urlNew)
		urlNew = urlNew + "?" + key + "=" + value
		return
	}

	urlNew = urlNew + "&" + key + "=" + value
	return
}
//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//