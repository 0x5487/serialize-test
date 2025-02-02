//go:generate msgp
package domain

type Order struct {
	Id                 string `msg:"id"`
	Text               string `msg:"text"`
	AmendText          string `msg:"amend_text"`
	CreateTime         int64  `msg:"create_time"`
	UpdateTime         int64  `msg:"update_time"`
	CreateTimeMs       int64  `msg:"create_time_ms"`
	UpdateTimeMs       int64  `msg:"update_time_ms"`
	Status             string `msg:"status"`
	CurrencyPair       string `msg:"currency_pair"`
	Type               string `msg:"type"`
	Account            string `msg:"account"`
	Side               string `msg:"side"`
	Amount             string `msg:"amount"`
	Price              string `msg:"price"`
	TimeInForce        string `msg:"time_in_force"`
	Iceberg            string `msg:"iceberg"`
	Left               string `msg:"left"`
	FilledAmount       string `msg:"filled_amount"`
	FillPrice          string `msg:"fill_price"`
	FilledTotal        string `msg:"filled_total"`
	AvgDealPrice       string `msg:"avg_deal_price"`
	Fee                string `msg:"fee"`
	FeeCurrency        string `msg:"fee_currency"`
	PointFee           string `msg:"point_fee"`
	GtFee              string `msg:"gt_fee"`
	GtMakerFee         string `msg:"gt_maker_fee"`
	GtTakerFee         string `msg:"gt_taker_fee"`
	GtDiscount         bool   `msg:"gt_discount"`
	RebatedFee         string `msg:"rebated_fee"`
	RebatedFeeCurrency string `msg:"rebated_fee_currency"`
	FinishAs           string `msg:"finish_as"`
}
