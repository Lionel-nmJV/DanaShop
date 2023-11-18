package order

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type repoOrder struct {
}

func newRepoOrder() repoOrder {
	return repoOrder{}
}

func (r repoOrder) findAllByMerchantID(ctx *gin.Context, tx *sqlx.Tx, merchantID string, query string, limit interface{}, offset int) ([]orderResponse, error) {
	SQL := `SELECT "id", "product_name", "quantity", "product_price", "user_name", "user_phone_number", "user_address", "total_price", "status" FROM "orders" WHERE "merchant_id"=$1 AND "product_name"  ILIKE CONCAT('%', $2::text, '%') LIMIT $3 OFFSET $4`
	rows, err := tx.QueryContext(ctx, SQL, merchantID, query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []orderResponse
	for rows.Next() {
		order := orderResponse{}
		err := rows.Scan(&order.ID, &order.ProductName, &order.Quantity, &order.ProductPrice, &order.UserName, &order.UserPhoneNumber, &order.UserAddress, &order.TotalPrice, &order.Status)
		if err != nil {
			return nil, err
		}

		orders = append(orders, order)
	}

	return orders, nil
}
