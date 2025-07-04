// Code generated by Kitex v0.14.1. DO NOT EDIT.

package coin

import "github.com/cloudwego/prutal"

type Order struct {
	Id                 string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Text               string `protobuf:"bytes,2,opt,name=text" json:"text,omitempty"`
	AmendText          string `protobuf:"bytes,3,opt,name=amend_text" json:"amend_text,omitempty"`
	CreateTime         int64  `protobuf:"varint,4,opt,name=create_time" json:"create_time,omitempty"`
	UpdateTime         int64  `protobuf:"varint,5,opt,name=update_time" json:"update_time,omitempty"`
	CreateTimeMs       int64  `protobuf:"varint,6,opt,name=create_time_ms" json:"create_time_ms,omitempty"`
	UpdateTimeMs       int64  `protobuf:"varint,7,opt,name=update_time_ms" json:"update_time_ms,omitempty"`
	Status             string `protobuf:"bytes,8,opt,name=status" json:"status,omitempty"`
	CurrencyPair       string `protobuf:"bytes,9,opt,name=currency_pair" json:"currency_pair,omitempty"`
	Type               string `protobuf:"bytes,10,opt,name=type" json:"type,omitempty"`
	Account            string `protobuf:"bytes,11,opt,name=account" json:"account,omitempty"`
	Side               string `protobuf:"bytes,12,opt,name=side" json:"side,omitempty"`
	Amount             string `protobuf:"bytes,13,opt,name=amount" json:"amount,omitempty"`
	Price              string `protobuf:"bytes,14,opt,name=price" json:"price,omitempty"`
	TimeInForce        string `protobuf:"bytes,15,opt,name=time_in_force" json:"time_in_force,omitempty"`
	Iceberg            string `protobuf:"bytes,16,opt,name=iceberg" json:"iceberg,omitempty"`
	Left               string `protobuf:"bytes,17,opt,name=left" json:"left,omitempty"`
	FilledAmount       string `protobuf:"bytes,18,opt,name=filled_amount" json:"filled_amount,omitempty"`
	FillPrice          string `protobuf:"bytes,19,opt,name=fill_price" json:"fill_price,omitempty"`
	FilledTotal        string `protobuf:"bytes,20,opt,name=filled_total" json:"filled_total,omitempty"`
	AvgDealPrice       string `protobuf:"bytes,21,opt,name=avg_deal_price" json:"avg_deal_price,omitempty"`
	Fee                string `protobuf:"bytes,22,opt,name=fee" json:"fee,omitempty"`
	FeeCurrency        string `protobuf:"bytes,23,opt,name=fee_currency" json:"fee_currency,omitempty"`
	PointFee           string `protobuf:"bytes,24,opt,name=point_fee" json:"point_fee,omitempty"`
	GtFee              string `protobuf:"bytes,25,opt,name=gt_fee" json:"gt_fee,omitempty"`
	GtMakerFee         string `protobuf:"bytes,26,opt,name=gt_maker_fee" json:"gt_maker_fee,omitempty"`
	GtTakerFee         string `protobuf:"bytes,27,opt,name=gt_taker_fee" json:"gt_taker_fee,omitempty"`
	GtDiscount         bool   `protobuf:"varint,28,opt,name=gt_discount" json:"gt_discount,omitempty"`
	RebatedFee         string `protobuf:"bytes,29,opt,name=rebated_fee" json:"rebated_fee,omitempty"`
	RebatedFeeCurrency string `protobuf:"bytes,30,opt,name=rebated_fee_currency" json:"rebated_fee_currency,omitempty"`
	FinishAs           string `protobuf:"bytes,31,opt,name=finish_as" json:"finish_as,omitempty"`
}

func (x *Order) Reset() { *x = Order{} }

func (x *Order) Marshal(in []byte) ([]byte, error) { return prutal.MarshalAppend(in, x) }

func (x *Order) Unmarshal(in []byte) error { return prutal.Unmarshal(in, x) }

func (x *Order) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Order) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *Order) GetAmendText() string {
	if x != nil {
		return x.AmendText
	}
	return ""
}

func (x *Order) GetCreateTime() int64 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

func (x *Order) GetUpdateTime() int64 {
	if x != nil {
		return x.UpdateTime
	}
	return 0
}

func (x *Order) GetCreateTimeMs() int64 {
	if x != nil {
		return x.CreateTimeMs
	}
	return 0
}

func (x *Order) GetUpdateTimeMs() int64 {
	if x != nil {
		return x.UpdateTimeMs
	}
	return 0
}

func (x *Order) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Order) GetCurrencyPair() string {
	if x != nil {
		return x.CurrencyPair
	}
	return ""
}

func (x *Order) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Order) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *Order) GetSide() string {
	if x != nil {
		return x.Side
	}
	return ""
}

func (x *Order) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

func (x *Order) GetPrice() string {
	if x != nil {
		return x.Price
	}
	return ""
}

func (x *Order) GetTimeInForce() string {
	if x != nil {
		return x.TimeInForce
	}
	return ""
}

func (x *Order) GetIceberg() string {
	if x != nil {
		return x.Iceberg
	}
	return ""
}

func (x *Order) GetLeft() string {
	if x != nil {
		return x.Left
	}
	return ""
}

func (x *Order) GetFilledAmount() string {
	if x != nil {
		return x.FilledAmount
	}
	return ""
}

func (x *Order) GetFillPrice() string {
	if x != nil {
		return x.FillPrice
	}
	return ""
}

func (x *Order) GetFilledTotal() string {
	if x != nil {
		return x.FilledTotal
	}
	return ""
}

func (x *Order) GetAvgDealPrice() string {
	if x != nil {
		return x.AvgDealPrice
	}
	return ""
}

func (x *Order) GetFee() string {
	if x != nil {
		return x.Fee
	}
	return ""
}

func (x *Order) GetFeeCurrency() string {
	if x != nil {
		return x.FeeCurrency
	}
	return ""
}

func (x *Order) GetPointFee() string {
	if x != nil {
		return x.PointFee
	}
	return ""
}

func (x *Order) GetGtFee() string {
	if x != nil {
		return x.GtFee
	}
	return ""
}

func (x *Order) GetGtMakerFee() string {
	if x != nil {
		return x.GtMakerFee
	}
	return ""
}

func (x *Order) GetGtTakerFee() string {
	if x != nil {
		return x.GtTakerFee
	}
	return ""
}

func (x *Order) GetGtDiscount() bool {
	if x != nil {
		return x.GtDiscount
	}
	return false
}

func (x *Order) GetRebatedFee() string {
	if x != nil {
		return x.RebatedFee
	}
	return ""
}

func (x *Order) GetRebatedFeeCurrency() string {
	if x != nil {
		return x.RebatedFeeCurrency
	}
	return ""
}

func (x *Order) GetFinishAs() string {
	if x != nil {
		return x.FinishAs
	}
	return ""
}
