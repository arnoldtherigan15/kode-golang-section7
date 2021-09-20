package main

import (
	"fmt"
	"log"
	"os"

	"kode-golang-section7/domain"
	_orderHandler "kode-golang-section7/order/handler"
	_orderRepo "kode-golang-section7/order/repository"
	_orderService "kode-golang-section7/order/service"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(".env file not found\n", err)
	}
}

func main() {
	DB_HOST := os.Getenv("DB_HOST")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DB_USER := os.Getenv("DB_USER")
	DB_NAME := os.Getenv("DB_NAME")
	DB_PORT := os.Getenv("DB_PORT")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", DB_HOST, DB_USER, DB_PASSWORD, DB_NAME, DB_PORT)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Connection Database Error\n", err.Error())
	}
	db.AutoMigrate(&domain.Order{}, &domain.Item{})

	orderRepo := _orderRepo.NewRepository(db)
	orderService := _orderService.NewService(orderRepo)

	e := echo.New()

	_orderHandler.NewHandler(e, orderService)

	PORT := os.Getenv("SERVER_PORT")
	log.Fatal(e.Start(PORT))
}
