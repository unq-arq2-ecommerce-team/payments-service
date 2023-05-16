package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/unq-arq2-ecommerce-team/payments-service/internal/application"
)

func CreatePaymentHandler(createPaymentUsecase *application.CreatePaymentUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		var newPament application.CreatePaymentDto
		if err := c.ShouldBindJSON(&newPament); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": err.Error()})
			return
		}
		payment, err := createPaymentUsecase.Do(&newPament)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"status": "ok", "data": payment})
	}
}

func UpdatePaymentMethodHandler(updatePaymentMethodUsecase *application.UpdatePaymentMethodUsecaseUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		var updatePaymentMethodDto application.UpdatePaymentMethodDto
		if err := c.ShouldBindJSON(&updatePaymentMethodDto); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": err.Error()})
			return
		}
		payment, err := updatePaymentMethodUsecase.Do(&updatePaymentMethodDto)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "ok", "data": payment})
	}
}

func ConfirmPaymentHandler(confirmPaymentUsecase *application.ConfirmPaymentUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		paymentId := c.Param("id")
		if paymentId == "" {
			c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "payment id is required"})
			return
		}
		payment, err := confirmPaymentUsecase.Do(paymentId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "ok", "data": payment})
	}
}

func RejectPaymentHandler(rejectPaymentUsecase *application.RejectPaymentUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		paymentId := c.Param("id")
		if paymentId == "" {
			c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "payment id is required"})
			return
		}
		payment, err := rejectPaymentUsecase.Do(paymentId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "ok", "data": payment})
	}
}
