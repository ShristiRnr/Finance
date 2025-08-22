package grpc

import (
    "math/big"

    pb "github.com/ShristiRnr/Finance/api/proto/financepb"
    "github.com/ShristiRnr/Finance/internal/core/domain/finance"
)

// Convert proto Money → domain Money
func MoneyFromProto(pm *pb.Money) finance.Money {
    if pm == nil {
        return finance.Money{}
    }

    units := big.NewInt(pm.Units)                // Units = int64
    nanos := big.NewInt(int64(pm.Nanos))        // Nanos = int32, cast to int64

    total := new(big.Int).Mul(units, big.NewInt(100)) // units → minor units
    total.Add(total, new(big.Int).Div(nanos, big.NewInt(1e7)))

    return finance.Money{
        Currency: pm.CurrencyCode,
        Amount:   total.Int64(),
    }
}

// Convert domain Money → proto Money
func MoneyToProto(m finance.Money) *pb.Money {
    units := m.Amount / 100
    nanos := (m.Amount % 100) * 1e7
    return &pb.Money{
        CurrencyCode: m.Currency,
        Units:        units,
        Nanos:        int32(nanos),
    }
}
