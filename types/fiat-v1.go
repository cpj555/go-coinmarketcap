package types

type GetFiatMapReq struct {
	Start         int  `schema:"start,omitempty"`
	Limit         int  `schema:"limit,omitempty"`
	Sort          Sort `schema:"sort,omitempty"`
	IncludeMetals bool `schema:"include_metals,omitempty"`
}

type GetFiatResp []GetFiatItem

type GetFiatItem struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Sign   string `json:"sign"`
	Symbol string `json:"symbol"`
}
