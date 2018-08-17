package types

type CryptocurrencyInfo struct {
	ID       int      `json:"id" bson:"_id"`
	Name     string   `json:"name" bson:"name"`
	Symbol   string   `json:"symbol" bson:"symbol"`
	Categoty string   `json:"category" bson:"category"`
	Slug     string   `json:"slug" bson:"slug"`
	Logo     string   `json:"logo" bson:"logo"`
	Tags     []string `json:"tags" bson:"tags"`
	URLs     URLs     `json:"urls" bson:"urls"`
}

type URLs struct {
	Website      []string `json:"website" bson:"website"`
	Explorer     []string `json:"explorer" bson:"explorer"`
	SourceCode   []string `json:"source_code" bson:"source_code"`
	MessageBoard []string `json:"message_board" bson:"message_board"`
	Chat         []string `json:"chat" bson:"chat"`
	Announcement []string `json:"announcement" bson:"announcement"`
	Reddit       []string `json:"reddit" bson:"reddit"`
	Twitter      []string `json:"twitter" bson:"twitter"`
}
