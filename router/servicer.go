package router

import (
	"fmt"
	"strconv"
)

// CheckParam will check the storeNumber param to see if it is valid.
func CheckParam(param string) (storeNumber int, err error) {
	storeNum, err := strconv.Atoi(param)
	if err != nil {
		return storeNum, fmt.Errorf("-> NO NUMBER DETECTED AFTER /get-sale-sub/#STORE_NUMBER. YOU CAN USE " +
			"/get-sale/ WITH NO PARAMETERS FOR A GLOBAL RESPONSE")
	}
	return storeNum, nil
}
