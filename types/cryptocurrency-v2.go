package types

import (
	"time"
)

// GetInfoResp CoinMarketCap ID Map
// [https://coinmarketcap.com/api/documentation/v1/#operation/getV2CryptocurrencyInfo]
type GetInfoResp map[string]struct {
	Urls struct {
		Website      []string `json:"website"`
		TechnicalDoc []string `json:"technical_doc"`
		Twitter      []string `json:"twitter"`
		Reddit       []string `json:"reddit"`
		MessageBoard []string `json:"message_board"`
		Announcement []string `json:"announcement"`
		Chat         []string `json:"chat"`
		Explorer     []string `json:"explorer"`
		SourceCode   []string `json:"source_code"`
	} `json:"urls"`
	Logo                          string      `json:"logo"`
	Id                            int         `json:"id"`
	Name                          string      `json:"name"`
	Symbol                        string      `json:"symbol"`
	Slug                          string      `json:"slug"`
	Description                   string      `json:"description"`
	Notice                        interface{} `json:"notice"`
	DateAdded                     time.Time   `json:"date_added"`
	DateLaunched                  time.Time   `json:"date_launched"`
	Tags                          []string    `json:"tags"`
	Platform                      *Platform   `json:"platform"`
	Category                      string      `json:"category"`
	SelfReportedCirculatingSupply interface{} `json:"self_reported_circulating_supply"`
	SelfReportedMarketCap         interface{} `json:"self_reported_market_cap"`
	SelfReportedTags              interface{} `json:"self_reported_tags"`
	InfiniteSupply                bool        `json:"infinite_supply"`
}

type (
	GetInfoReq struct {
		Id          string `schema:"id,omitempty"`
		Slug        string `schema:"slug,omitempty"`
		Symbol      string `schema:"symbol,omitempty"`
		Address     string `schema:"address,omitempty"`
		SkipInvalid bool   `schema:"skip_invalid,omitempty"`
		Aux         string `schema:"aux,omitempty"`
	}
)

func (p *GetInfoReq) ValidParams() error {
	return nil
}
