//go:generate msgp
package domain

type Order struct {
	ID                 string `msg:"id" frugal:"1,required"`
	Text               string `msg:"text" frugal:"2,required"`
	AmendText          string `msg:"amend_text" frugal:"3,required"`
	Status             string `msg:"status" frugal:"8,required"`
	CurrencyPair       string `msg:"currency_pair" frugal:"9,required"`
	Type               string `msg:"type" frugal:"10,required"`
	Account            string `msg:"account" frugal:"11,required"`
	Side               string `msg:"side" frugal:"12,required"`
	Amount             string `msg:"amount" frugal:"13,required"`
	Price              string `msg:"price" frugal:"14,required"`
	TimeInForce        string `msg:"time_in_force" frugal:"15,required"`
	Iceberg            string `msg:"iceberg" frugal:"16,required"`
	Left               string `msg:"left" frugal:"17,required"`
	FilledAmount       string `msg:"filled_amount" frugal:"18,required"`
	FillPrice          string `msg:"fill_price" frugal:"19,required"`
	FilledTotal        string `msg:"filled_total" frugal:"20,required"`
	AvgDealPrice       string `msg:"avg_deal_price" frugal:"21,required"`
	Fee                string `msg:"fee" frugal:"22,required"`
	FeeCurrency        string `msg:"fee_currency" frugal:"23,required"`
	PointFee           string `msg:"point_fee" frugal:"24,required"`
	GtFee              string `msg:"gt_fee" frugal:"25,required"`
	GtMakerFee         string `msg:"gt_maker_fee" frugal:"26,required"`
	GtTakerFee         string `msg:"gt_taker_fee" frugal:"27,required"`
	RebatedFee         string `msg:"rebated_fee" frugal:"29,required"`
	RebatedFeeCurrency string `msg:"rebated_fee_currency" frugal:"30,required"`
	FinishAs           string `msg:"finish_as" frugal:"31,required"`
	CreateTime         int64  `msg:"create_time" frugal:"4,required"`
	UpdateTime         int64  `msg:"update_time" frugal:"5,required"`
	CreateTimeMs       int64  `msg:"create_time_ms" frugal:"6,required"`
	UpdateTimeMs       int64  `msg:"update_time_ms" frugal:"7,required"`
	GtDiscount         bool   `msg:"gt_discount" frugal:"28,required"`
}
