package handler

import (
	"kode-golang-section7/domain"
	"net/http"
	"strconv"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	OrderService domain.OrderService
}

func NewHandler(e *echo.Echo, orderService domain.OrderService) {
	handler := &Handler{orderService}
	g := e.Group("/orders")
	g.POST("", handler.Create)
	g.GET("", handler.FindAll)
	g.PUT("", handler.Update)
	g.DELETE("/:id", handler.Delete)
}

func isRequestValid(order *domain.Order) (bool, error) {
	validate := validator.New()
	if err := validate.Struct(order); err != nil {
		return false, err
	}
	return true, nil
}

func (h *Handler) FindAll(c echo.Context) (err error) {
	orders, err := h.OrderService.FindAll()

	if err != nil {
		errorMsg := map[string]interface{}{
			"errors": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, errorMsg)
	}

	return c.JSON(http.StatusOK, orders)
}

func (h *Handler) Create(c echo.Context) (err error) {
	var order domain.Order
	err = c.Bind(&order)
	if err != nil {
		errorMsg := map[string]interface{}{
			"errors": err.Error(),
		}
		return c.JSON(http.StatusUnprocessableEntity, errorMsg)
	}

	if ok, err := isRequestValid(&order); !ok {
		errorMsg := map[string]interface{}{
			"errors": err.Error(),
		}
		return c.JSON(http.StatusUnprocessableEntity, errorMsg)
	}

	createdOrder, err := h.OrderService.Create(&order)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, createdOrder)
}

func (h *Handler) Update(c echo.Context) (err error) {
	var order domain.Order
	err = c.Bind(&order)
	if err != nil {
		errorMsg := map[string]interface{}{
			"errors": err.Error(),
		}
		return c.JSON(http.StatusUnprocessableEntity, errorMsg)
	}

	if ok, err := isRequestValid(&order); !ok {
		errorMsg := map[string]interface{}{
			"errors": err.Error(),
		}
		return c.JSON(http.StatusBadRequest, errorMsg)
	}

	updatedCar, err := h.OrderService.Update(&order)

	if err != nil {
		errorMsg := map[string]interface{}{
			"errors": err.Error(),
		}
		return c.JSON(http.StatusNotFound, errorMsg)
	}

	return c.JSON(http.StatusOK, updatedCar)
}

func (h *Handler) Delete(c echo.Context) (err error) {
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errorMsg := map[string]interface{}{
			"errors": err.Error(),
		}
		return c.JSON(http.StatusBadRequest, errorMsg)
	}
	isDeleted, err := h.OrderService.Delete(ID)

	if err != nil {
		errorMsg := map[string]interface{}{
			"errors": err.Error(),
		}
		return c.JSON(http.StatusNotFound, errorMsg)
	}

	response := map[string]bool{
		"is_delete": isDeleted,
	}

	return c.JSON(http.StatusOK, response)
}
