// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0

package sqlc

import (
	"github.com/gofrs/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/shopspring/decimal"
)

type TradingAsset struct {
	AssetID     uuid.UUID
	Name        string
	Ticker      string
	CreatedAt   pgtype.Timestamptz
	UpdatedAt   pgtype.Timestamptz
	DeletedAt   pgtype.Timestamptz
	Description pgtype.Text
}

type TradingPriceHistory struct {
	PriceHistoryID int64
	AssetID        uuid.UUID
	PriceAssetID   uuid.UUID
	Price          decimal.Decimal
	EventTime      pgtype.Timestamptz
	CreatedAt      pgtype.Timestamptz
	UpdatedAt      pgtype.Timestamptz
	DeletedAt      pgtype.Timestamptz
}
