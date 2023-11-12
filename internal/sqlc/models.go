// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0

package sqlc

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type TradingAsset struct {
	AssetID     pgtype.UUID
	Name        string
	Ticker      string
	CreatedAt   pgtype.Timestamptz
	UpdatedAt   pgtype.Timestamptz
	DeletedAt   pgtype.Timestamptz
	Description pgtype.Text
}

type TradingPriceHistory struct {
	PriceHistoryID int64
	AssetID        pgtype.UUID
	PriceAssetID   pgtype.UUID
	Price          pgtype.Numeric
	EventTime      pgtype.Timestamptz
	CreatedAt      pgtype.Timestamptz
	UpdatedAt      pgtype.Timestamptz
	DeletedAt      pgtype.Timestamptz
}