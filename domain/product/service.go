package product

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jmoiron/sqlx"
	"starfish/domain/merchant"
	"strconv"
	"time"
)

type productRepository interface {
	readProductRepo
	writeProductRepo
	findProductByID(ctx *gin.Context, tx *sqlx.Tx, productID string) (Product, error)
}

type readProductRepo interface {
	findAllByMerchantID(ctx *gin.Context, tx *sqlx.Tx, merchantID string, query string, limit interface{}, offset int) ([]productResponses, error)
}

type writeProductRepo interface {
	saveProduct(ctx *gin.Context, tx *sqlx.Tx, product Product) (string, error)
	updateProduct(ctx *gin.Context, tx *sqlx.Tx, productID string, product Product) error
	deleteProduct(ctx *gin.Context, tx *sqlx.Tx, productID string) error
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

func (service productService) updateProduct(ctx *gin.Context, productID string, request updateRequest) error {
	tx, err := service.db.Beginx()
	if err != nil {
		return err
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Retrieve the existing product from the database using the repository's findProductByID method.
	existingProduct, err := service.repoProduct.findProductByID(ctx, tx, productID)
	if err != nil {
		return err
	}

	// Retrieve the merchantFounded ID from the context.
	userClaims := ctx.MustGet("merchantFounded").(jwt.MapClaims)
	merchantFoundedID := userClaims["merchantFounded_ID"].(string)

	if existingProduct.MerchantID != merchantFoundedID {
		return errors.New("Permission denied")
	}

	// Apply the update data from the request to the existing product.
	existingProduct.Name = request.Name
	existingProduct.Category = request.Category
	existingProduct.Price = request.Price
	existingProduct.Stock = request.Stock
	existingProduct.ImageURL = request.ImageURL
	existingProduct.UpdatedAt = time.Now() // Update the "updated_at" timestamp.

	// Update the product in the database using the repository's updateProduct method.
	err = service.repoProduct.updateProduct(ctx, tx, productID, existingProduct)
	if err != nil {
		return err
	}

	// Commit the transaction when the update is successful.
	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (service productService) deleteProduct(ctx *gin.Context, productID string) error {
	// Add logic to delete the product using the repository's deleteProduct method.
	tx, err := service.db.Beginx()
	if err != nil {
		return err
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Call the repository's deleteProduct method to delete the product.
	err = service.repoProduct.deleteProduct(ctx, tx, productID)
	if err != nil {
		tx.Rollback()
		return err
	}

	// Commit the transaction when the deletion is successful.
	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}
