package items

type Item struct {
	Id                string      `json:"id,omitempty" bson:"-"`
	Seller            int64       `json:"seller,omitempty" bson:"seller"`
	Title             string      `json:"title,omitempty" bson:"title"`
	Description       Description `json:"description,omitempty" bson:"description"`
	Pictures          []Picture   `json:"pictures,omitempty" bson:"pictures"`
	Video             string      `json:"video,omitempty" bson:"video"`
	Price             float32     `json:"price,omitempty" bson:"price"`
	AvailableQuantity int         `json:"available_quantity,omitempty" bson:"available_quantity"`
	SoldQuantity      int         `json:"sold_quantity,omitempty" bson"sold_quantity"`
	Status            string      `json:"status,omitempty" bson:"status"`
}

type Description struct {
	PlainText string `json:"plain_text,omitempty" bson:"plain_text"`
	Html      string `json:"html,omitempty" bson:"html"`
}

type Picture struct {
	Id  int64  `json:"id,omitempty" bson:"id"`
	Url string `json:"url,omitempty" bson:"url"`
}


