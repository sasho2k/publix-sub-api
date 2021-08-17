package workforce

// PublixResponse Data structure for the Publix response from the API. We will be looking at the Products array mostly.
type PublixResponse struct {
	TotalCount int         `json:"TotalCount"`
	NoImageURL interface{} `json:"NoImageUrl"`
	Products   []struct {
		IsPOS                  bool        `json:"isPOS"`
		Productid              string      `json:"Productid"`
		Itemcode               string      `json:"itemcode"`
		Title                  string      `json:"title"`
		Shortdescription1      string      `json:"shortdescription1"`
		Shortdescription2      interface{} `json:"shortdescription2"`
		Limitedquantitymsg     interface{} `json:"limitedquantitymsg"`
		Priceline1             string      `json:"priceline1"`
		Priceline2             interface{} `json:"priceline2"`
		Displaytype            string      `json:"displaytype"`
		Savingmsg              string      `json:"savingmsg"`
		Validthrumsg           interface{} `json:"validthrumsg"`
		Onsalemsg              string      `json:"onsalemsg"`
		Rssinfo                string      `json:"rssinfo"`
		Rsslocation            interface{} `json:"rsslocation"`
		Productimages          string      `json:"productimages"`
		Productimagesxl        string      `json:"productimagesxl"`
		Activationstatus       string      `json:"activationstatus"`
		Productmoreinfo        interface{} `json:"productmoreinfo"`
		Specialpromotion       interface{} `json:"specialpromotion"`
		Age                    int         `json:"age"`
		SizeDescription        interface{} `json:"sizeDescription"`
		NutritionalDescription interface{} `json:"nutritionalDescription"`
		Advertising            bool        `json:"advertising"`
		MarketingImages        []struct {
			ImageURL string      `json:"ImageURL"`
			Title    string      `json:"Title"`
			URL      interface{} `json:"Url"`
		} `json:"MarketingImages"`
		NoImageURL        string      `json:"NoImageUrl"`
		FauxTaxonomy      []string    `json:"fauxTaxonomy"`
		FavoriteProductID interface{} `json:"favoriteProductId"`
		GroceryListItemID string      `json:"groceryListItemId"`
	} `json:"Products"`
	Facets     []interface{} `json:"Facets"`
	Categories interface{}   `json:"Categories"`
	TopParent  interface{}   `json:"TopParent"`
	Parent     interface{}   `json:"Parent"`
}

// LocalSubStruct is the local struct for sub information. This is what we transfer to from initial response struct.
type LocalSubStruct struct {
	Name        string `json:"Name"`
	Price       string `json:"Price"`
	Description string `json:"Description"`
	ProductID   string `json:"ProductID"`
	ItemCode    string `json:"ItemCode"`
	SavingMsg   string `json:"SavingMsg"`
}
