package workforce

import (
	"fmt"
	"golang.org/x/net/html"
	"strings"
)

// GrabDailySub will serve to get the response info and parse it to return our requested info.
// For each item, check if the on sale msg is active. If so, we can return that item title as well as the price.
func GrabDailySub(storeNumber int) (sub LocalSubStruct, err error) {
	response, err := GetPublixResponse(storeNumber)
	if err != nil {
		return LocalSubStruct{}, err
	}

	for _, product := range response.Products {
		if product.Onsalemsg == "On Sale" {
			return LocalSubStruct{
				Name:        html.UnescapeString(product.Title),
				Price:       strings.Split(product.Priceline1, " ")[2],
				Description: product.Shortdescription1,
				ProductID:   product.Productid,
				ItemCode:    product.Itemcode,
				SavingMsg:   product.Savingmsg,
			}, nil
		}
	}

	return LocalSubStruct{}, fmt.Errorf("-> ERROR NO SUB ON SALE FOUND")
}

// GrabSubs will grab all subs and return them as an array.
func GrabSubs(storeNumber int) (subsList []LocalSubStruct, err error) {
	var (
		subs []LocalSubStruct
	)

	response, err := GetPublixResponse(storeNumber)
	if err != nil {
		return nil, err
	}

	for _, product := range response.Products {
		price := "UNAVAILABLE IF 9999."
		if storeNumber != 9999 {
			price = strings.Split(product.Priceline1, " ")[2]
		}

		subs = append(subs, LocalSubStruct{
			Name:        product.Title,
			Price:       price,
			Description: product.Shortdescription1,
			ProductID:   product.Productid,
			ItemCode:    product.Itemcode,
			SavingMsg:   product.Savingmsg,
		})
	}

	return subs, nil
}
