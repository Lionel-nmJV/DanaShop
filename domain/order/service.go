package order

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jmoiron/sqlx"
	"strconv"
)

type orderRepository interface {
	readOrderRepo
}

type readOrderRepo interface {
	findAllByMerchantID(ctx *gin.Context, tx *sqlx.Tx, merchantID string, query string, limit interface{}, offset int) ([]orderResponse, error)
}

type orderService struct {
	orderRepo orderRepository
	db        *sqlx.DB
}

func newService(orderRepo orderRepository, db *sqlx.DB) orderService {
	return orderService{
		orderRepo: orderRepo,
		db:        db,
	}
}

func (service orderService) findAllByMerchantID(ctx *gin.Context) (paginateOrdersResponse, error) {
	query := ctx.Query("query")
	pageString := ctx.Query("page")
	page, err := strconv.Atoi(pageString)
	if page < 1 {
		page = 1
	}

	limitString := ctx.Query("limit")
	limit, err := strconv.Atoi(limitString)
	if err != nil || limit < 1 {
		limit = 10
	}
	offset := (page - 1) * limit

	userClaims := ctx.MustGet("user").(jwt.MapClaims)
	merchantID := userClaims["merchant_id"].(string)

	var responsePaginate paginateOrdersResponse
	tx, err := service.db.Beginx()
	if err != nil {
		return responsePaginate, err
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	orders, err := service.orderRepo.findAllByMerchantID(ctx, tx, merchantID, query, limit, offset)

	allOrders, err := service.orderRepo.findAllByMerchantID(ctx, tx, merchantID, "", nil, 0)

	if err := tx.Commit(); err != nil {
		return responsePaginate, err
	}

	totalPages := (len(allOrders) + limit - 1) / limit

	responsePaginate = paginateOrdersResponse{
		Orders: orders,
		Pagination: pagination{
			Page:       page,
			PerPage:    limit,
			TotalPages: totalPages,
			TotalItems: len(allOrders),
		},
	}

	return responsePaginate, nil
}
