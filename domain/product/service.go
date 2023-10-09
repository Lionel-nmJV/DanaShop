package product

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jmoiron/sqlx"
	"starfish/domain/merchant"
	"strconv"
)

type productRepository interface {
	readProductRepo
	writeProductRepo
}

type readProductRepo interface {
	findAllByMerchantID(ctx *gin.Context, tx *sqlx.Tx, merchantID string, query string, limit interface{}, offset int) ([]productResponses, error)
}

type writeProductRepo interface {
	saveProduct(ctx *gin.Context, tx *sqlx.Tx, product Product) (string, error)
}

type MerchantRepository interface {
	ReadMerchantRepo
}

type ReadMerchantRepo interface {
	FindByUserID(ctx *gin.Context, tx *sqlx.Tx, merchantID string) (merchant.Merchant, error)
}

type productService struct {
	repoProduct  productRepository
	RepoMerchant MerchantRepository
	db           *sqlx.DB
}

func newService(repoProduct productRepository, repoMerchant MerchantRepository, db *sqlx.DB) productService {
	return productService{
		repoProduct:  repoProduct,
		RepoMerchant: repoMerchant,
		db:           db,
	}
}

func (service productService) findAllByMerchantID(ctx *gin.Context) (paginateProductsResponse, error) {
	query := ctx.Query("query")
	pageString := ctx.Query("page")
	page, err := strconv.Atoi(pageString)
	limit := 10
	if page < 1 {
		page = 1
	}
	offset := (page - 1) * limit

	userClaims := ctx.MustGet("user").(jwt.MapClaims)
	userID := userClaims["user_id"].(string)

	var responsePaginate paginateProductsResponse
	tx, err := service.db.Beginx()
	if err != nil {
		return responsePaginate, err
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	merchantFounded, err := service.RepoMerchant.FindByUserID(ctx, tx, userID)
	if err != nil {
		return responsePaginate, err
	}

	products, err := service.repoProduct.findAllByMerchantID(ctx, tx, merchantFounded.ID, query, limit, offset)

	allProducts, err := service.repoProduct.findAllByMerchantID(ctx, tx, merchantFounded.ID, "", nil, 0)

	if err := tx.Commit(); err != nil {
		return responsePaginate, err
	}

	totalPages := (len(allProducts) + limit - 1) / limit

	responsePaginate = paginateProductsResponse{
		Products: products,
		Pagination: pagination{
			Page:       page,
			PerPage:    10,
			TotalPages: totalPages,
			TotalItems: len(allProducts),
		},
	}

	return responsePaginate, nil
}

func (service productService) addProduct(ctx *gin.Context, request Product) error {
	userClaims := ctx.MustGet("user").(jwt.MapClaims)
	userID := userClaims["user_id"].(string)

	tx, err := service.db.Beginx()
	if err != nil {
		return err
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	merchantFounded, err := service.RepoMerchant.FindByUserID(ctx, tx, userID)
	if err != nil {
		return err
	}

	request.MerchantID = merchantFounded.ID

	_, err = service.repoProduct.saveProduct(ctx, tx, request)
	if err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}