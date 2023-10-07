package user

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

var svc UserService

var dbSql *sqlx.DB

func ConnectPostgre(host, port, user, pass, dbname string) (db *sqlx.DB, err error) {
	dsn := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=disable",
		host, port, user, pass, dbname,
	)

	db, err = sqlx.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return
}

func TestConnectPostgres(t *testing.T) {
	db, err := ConnectPostgre(
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME"),
	)

	if err != nil {
		t.Errorf("expected no error, but got : %v", err.Error())
	}

	if db == nil {
		t.Error("expect not nil, but got nil ")
	}
}

func init() {
	err := godotenv.Load("../../.env")
	if err != nil {
		panic(err)
	}

	db, err := ConnectPostgre(
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME"),
	)

	if err != nil {
		panic(err)
	}

	dbSql = db
	repo := NewPostgres()
	svc = NewService(repo, db)
}

func TestService_Register(t *testing.T) {
	// Skip the test if database connection details are not provided
	if os.Getenv("DB_HOST") == "" || os.Getenv("DB_PORT") == "" || os.Getenv("DB_USER") == "" || os.Getenv("DB_PASS") == "" || os.Getenv("DB_NAME") == "" {
		t.Skip("Skipping test: Database connection details not provided.")
	}

	// Create a test user and merchant
	testUser := User{
		Email:    "test@example.com",
		Password: "testpassword",
	}

	testMerchant := Merchant{
		Name: "Test Merchant",
	}

	// Perform registration using the actual service and database
	err := svc.Register(context.TODO(), testUser, testMerchant)
	if err != nil {
		t.Errorf("Expected no error, but got an error: %v", err)
	}

	// Clean up the testing database (rollback any changes made by the test)
	// This ensures that the database is in a clean state for the next test.
	dbSql.Exec("DELETE FROM merchants WHERE name = 'Test Merchant'")
	dbSql.Exec("DELETE FROM users WHERE email = 'test@example.com'")

}

func createTestUser() (User, error) {
	// Create a test user with known credentials
	testUser := User{
		Email:    "test@example.com",
		Password: "testpassword", // Replace with the actual password or hash
	}

	// Create a test merchant
	testMerchant := Merchant{
		Name: "Test Merchant",
	}

	// Use the Register service to create the test user and merchant
	err := svc.Register(context.TODO(), testUser, testMerchant)
	if err != nil {
		return User{}, err
	}

	return testUser, nil
}

func TestUserService_Login(t *testing.T) {
	// Skip the test if database connection details are not provided
	if os.Getenv("DB_HOST") == "" || os.Getenv("DB_PORT") == "" || os.Getenv("DB_USER") == "" || os.Getenv("DB_PASS") == "" || os.Getenv("DB_NAME") == "" {
		t.Skip("Skipping test: Database connection details not provided.")
	}

	// Create a test user and merchant using the createTestUser function
	testUser, err := createTestUser()
	if err != nil {
		t.Fatalf("Failed to create a test user: %v", err)
	}

	// Define a UserLogin with valid credentials
	userLogin := User{
		Email:    testUser.Email,
		Password: "testpassword", // Replace with the actual plain text password
	}

	// Perform the login using the actual service and database
	loginResponse, err := svc.login(context.TODO(), userLogin)
	if err != nil {
		t.Errorf("Expected no error, but got an error: %v", err)
	}

	// Add assertions to check the result of login
	if len(loginResponse.AccessToken) == 0 {
		t.Error("Expected a non-empty access token, but got an empty token.")
	}

	// Clean up the testing database (remove the test user)
	if _, err := dbSql.Exec("DELETE FROM merchants WHERE name = 'Test Merchant'"); err != nil {
		t.Errorf("Failed to delete test merchant: %v", err)
	}

	if _, err = dbSql.Exec("DELETE FROM users WHERE email = 'test@example.com'"); err != nil {
		t.Errorf("Failed to delete test user: %v", err)
	}

}
