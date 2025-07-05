package main

import (
	"mykitex/domain"
	pb "mykitex/kitex_gen/protobuf/coin"
	"mykitex/kitex_gen/prutal/coin"
	"testing"

	"github.com/cloudwego/frugal"
	"github.com/cloudwego/prutal"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/proto"
)

// 創建一個示例 Order 對象
func createPrutalOrder() *coin.Order {
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

func createProtobufOrder() *pb.Order {
	return &pb.Order{
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

func createOrder() *domain.Order {
	return &domain.Order{
		ID:                 "1852454420",
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

func TestFrugal(t *testing.T) {
	order := createOrder()
	buf := make([]byte, frugal.EncodedSize(order))
	_, err := frugal.EncodeObject(buf, nil, order)
	assert.NoError(t, err)

	newOrder := &domain.Order{}
	_, err = frugal.DecodeObject(buf, newOrder)
	assert.NoError(t, err)
	assert.Equal(t, order.ID, newOrder.ID)
}

func TestMsgp(t *testing.T) {
	order := createOrder()

	bs := []byte{}
	var err error
	bs, err = order.MarshalMsg(bs)
	if err != nil {
		t.Fatalf("序列化失敗: %v", err)
	}

	newOrder := &domain.Order{}
	_, err = newOrder.UnmarshalMsg(bs)
	if err != nil {
		t.Fatalf("反序列化失敗: %v", err)
	}

	assert.Equal(t, order.ID, newOrder.ID)
}

// 測試 Protobuf 序列化效能
func BenchmarkProtobufSerialize(b *testing.B) {
	order := createProtobufOrder()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := proto.Marshal(order)
		if err != nil {
			b.Fatalf("序列化失敗: %v", err)
		}
	}
}

// 測試 Protobuf 反序列化效能
func BenchmarkProtobufDeserialize(b *testing.B) {
	order := createProtobufOrder()
	data, err := prutal.Marshal(order)
	if err != nil {
		b.Fatalf("序列化失敗: %v", err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		newOrder := &pb.Order{}
		err := proto.Unmarshal(data, newOrder)
		if err != nil {
			b.Fatalf("反序列化失敗: %v", err)
		}
	}
}

func BenchmarkPrutalSerialize(b *testing.B) {
	order := createPrutalOrder()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := prutal.Marshal(order)
		if err != nil {
			b.Fatalf("序列化失敗: %v", err)
		}
	}
}

// 測試 Protobuf 反序列化效能
func BenchmarkPrutalDeserialize(b *testing.B) {
	order := createPrutalOrder()
	data, err := prutal.Marshal(order)
	if err != nil {
		b.Fatalf("序列化失敗: %v", err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		newOrder := &coin.Order{}
		err := prutal.Unmarshal(data, newOrder)
		if err != nil {
			b.Fatalf("反序列化失敗: %v", err)
		}
	}
}

func BenchmarkMsgpSerialize(b *testing.B) {
	order := createOrder()

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

func BenchmarkMsgpDeserialize(b *testing.B) {
	order := createOrder()

	bs := []byte{}
	var err error
	bs, err = order.MarshalMsg(bs)
	if err != nil {
		b.Fatalf("序列化失敗: %v", err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err = order.UnmarshalMsg(bs)
		if err != nil {
			b.Fatalf("反序列化失敗: %v", err)
		}
	}
}

func BenchmarkFrugalSerialize(b *testing.B) {
	order := createOrder()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		buf := make([]byte, frugal.EncodedSize(order))
		_, _ = frugal.EncodeObject(buf, nil, order)
	}
}

func BenchmarkFrugalDeserialize(b *testing.B) {
	order := createOrder()
	buf := make([]byte, frugal.EncodedSize(order))
	_, _ = frugal.EncodeObject(buf, nil, order)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		newOrder := &domain.Order{}
		_, _ = frugal.DecodeObject(buf, newOrder)
	}
}
