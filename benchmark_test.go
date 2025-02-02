package main

import (
	"mykitex/domain"
	"mykitex/kitex_gen/coin"
	"testing"

	"google.golang.org/protobuf/proto"
)

// 創建一個示例 Order 對象
func createOrder() *coin.Order {
	return &coin.Order{
		Id:                 "1852454420",
		Text:               "t-abc123",
		AmendText:          "-",
		CreateTime:         1710488334,
		UpdateTime:         1710488334,
		CreateTimeMs:       1710488334073,
		UpdateTimeMs:       1710488334074,
		Status:             "closed",
		CurrencyPair:       "BTC_USDT",
		Type:               "limit",
		Account:            "unified",
		Side:               "buy",
		Amount:             "0.001",
		Price:              "65000",
		TimeInForce:        "gtc",
		Iceberg:            "0",
		Left:               "0",
		FilledAmount:       "0.001",
		FillPrice:          "63.4693",
		FilledTotal:        "63.4693",
		AvgDealPrice:       "63469.3",
		Fee:                "0.00000022",
		FeeCurrency:        "BTC",
		PointFee:           "0",
		GtFee:              "0",
		GtMakerFee:         "0",
		GtTakerFee:         "0",
		GtDiscount:         false,
		RebatedFee:         "0",
		RebatedFeeCurrency: "USDT",
		FinishAs:           "filled",
	}
}

// 測試 Protobuf 序列化效能
func BenchmarkProtoSerialize(b *testing.B) {
	order := createOrder()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := proto.Marshal(order)
		if err != nil {
			b.Fatalf("序列化失敗: %v", err)
		}
	}
}

// 測試 Protobuf 反序列化效能
func BenchmarkProtoDeserialize(b *testing.B) {
	order := createOrder()
	data, err := proto.Marshal(order)
	if err != nil {
		b.Fatalf("序列化失敗: %v", err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		newOrder := &coin.Order{}
		err := proto.Unmarshal(data, newOrder)
		if err != nil {
			b.Fatalf("反序列化失敗: %v", err)
		}
	}
}

func createMsgpOrder() *domain.Order {
	return &domain.Order{
		Id:                 "1852454420",
		Text:               "t-abc123",
		AmendText:          "-",
		CreateTime:         1710488334,
		UpdateTime:         1710488334,
		CreateTimeMs:       1710488334073,
		UpdateTimeMs:       1710488334074,
		Status:             "closed",
		CurrencyPair:       "BTC_USDT",
		Type:               "limit",
		Account:            "unified",
		Side:               "buy",
		Amount:             "0.001",
		Price:              "65000",
		TimeInForce:        "gtc",
		Iceberg:            "0",
		Left:               "0",
		FilledAmount:       "0.001",
		FillPrice:          "63.4693",
		FilledTotal:        "63.4693",
		AvgDealPrice:       "63469.3",
		Fee:                "0.00000022",
		FeeCurrency:        "BTC",
		PointFee:           "0",
		GtFee:              "0",
		GtMakerFee:         "0",
		GtTakerFee:         "0",
		GtDiscount:         false,
		RebatedFee:         "0",
		RebatedFeeCurrency: "USDT",
		FinishAs:           "filled",
	}
}

func BenchmarkMsgpSerialize(b *testing.B) {
	order := createMsgpOrder()

	bs := []byte{}
	var err error

	b.ResetTimer()
	for i := 0; i < b.N; i++ {

		bs, err = order.MarshalMsg(bs)
		if err != nil {
			b.Fatalf("序列化失敗: %v", err)
		}
	}
}
