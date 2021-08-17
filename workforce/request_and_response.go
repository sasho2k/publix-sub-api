package workforce

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

// GetPublixResponse This function will return the struct containing all the response info from our request to our endpoint.
func GetPublixResponse(storeNumber int) (response PublixResponse, err error) {
	var (
		req            *http.Request
		url            string
		client         http.Client
		publixResponse PublixResponse
	)

	url = "https://services.publix.com/api/v4/ispu/product/Search?storeNumber=" + strconv.Itoa(storeNumber) +
		"&rowCount=60&orderAheadOnly=true&categoryID=d3a5f2da-3002-4c6d-949c-db5304577efb&sort=onsalemsg"
	req, err = http.NewRequest("GET", url, nil)

	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.131 Safari/537.36")

	resp, err := client.Do(req)
	if err != nil {
		return publixResponse, fmt.Errorf("-> ERROR PERFORMING INIT REQUEST")
	}

	if resp.StatusCode != 200 {
		return publixResponse, fmt.Errorf("-> 200 NOT RECEIVED")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return publixResponse, fmt.Errorf("-> ERROR READING BODY")
	}

	err = json.Unmarshal(body, &publixResponse)
	if err != nil {
		return publixResponse, fmt.Errorf("-> ERROR UNMARSHALLING JSON")
	}

	return publixResponse, nil
}
