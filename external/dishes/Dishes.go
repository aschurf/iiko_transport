package dishes

type Menu struct {
	Id             int          `json:"id"`
	Name           string       `json:"name"`
	Description    string       `json:"description"`
	ItemCategories []Categories `json:"itemCategories"`
}

type Categories struct {
	Id             string      `json:"id"`
	Name           string      `json:"name"`
	Description    string      `json:"description"`
	ButtonImageUrl string      `json:"buttonImageUrl"`
	HeaderImageUrl interface{} `json:"headerImageUrl"`
	IikoGroupId    interface{} `json:"iikoGroupId"`
	Items          []Dish      `json:"items"`
}

type Dish struct {
	Sku            string        `json:"sku"`
	Name           string        `json:"name"`
	Count          int           `json:"count"`
	Description    string        `json:"description"`
	AllergenGroups []interface{} `json:"allergenGroups"`
	Tags           []interface{} `json:"tags"`
	Labels         []interface{} `json:"labels"`
	ItemSizes      []struct {
		Sku                      string        `json:"sku"`
		SizeCode                 interface{}   `json:"sizeCode"`
		SizeName                 interface{}   `json:"sizeName"`
		IsDefault                bool          `json:"isDefault"`
		PortionWeightGrams       float32       `json:"portionWeightGrams"`
		ItemModifierGroups       []interface{} `json:"itemModifierGroups"`
		SizeId                   interface{}   `json:"sizeId"`
		NutritionPerHundredGrams struct {
			Fats          int      `json:"fats"`
			Proteins      int      `json:"proteins"`
			Carbs         int      `json:"carbs"`
			Energy        int      `json:"energy"`
			Organizations []string `json:"organizations"`
		} `json:"nutritionPerHundredGrams"`
		Prices []struct {
			OrganizationId string `json:"organizationId"`
			Price          int    `json:"price"`
		} `json:"prices"`
		Nutritions []struct {
			Fats          int      `json:"fats"`
			Proteins      int      `json:"proteins"`
			Carbs         int      `json:"carbs"`
			Energy        int      `json:"energy"`
			Organizations []string `json:"organizations"`
		} `json:"nutritions"`
		ButtonImageUrl        interface{} `json:"buttonImageUrl"`
		ButtonImageCroppedUrl interface{} `json:"buttonImageCroppedUrl"`
	} `json:"itemSizes"`
	ItemId             string      `json:"itemId"`
	ModifierSchemaId   interface{} `json:"modifierSchemaId"`
	TaxCategory        interface{} `json:"taxCategory"`
	ModifierSchemaName interface{} `json:"modifierSchemaName"`
	OrderItemType      string      `json:"orderItemType"`
}
