package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	application "github.com/unq-arq2-ecommerce-team/payments-service/internal/application"
)

type ginApplication struct {
	CreatePaymentUsecase       *application.CreatePaymentUseCase
	UpdatePaymentMethodUsecase *application.UpdatePaymentMethodUsecaseUseCase
	ConfirmPaymentUsecase      *application.ConfirmPaymentUseCase
	RejectPaymentUsecase       *application.RejectPaymentUseCase
}

func NewGinApplication(
	createPaymentUsecase *application.CreatePaymentUseCase,
	updatePaymentMethodUsecase *application.UpdatePaymentMethodUsecaseUseCase,
	confirmPaymentUsecase *application.ConfirmPaymentUseCase,
	rejectPaymentUsecase *application.RejectPaymentUseCase,
) *ginApplication {
	return &ginApplication{
		CreatePaymentUsecase:       createPaymentUsecase,
		UpdatePaymentMethodUsecase: updatePaymentMethodUsecase,
		ConfirmPaymentUsecase:      confirmPaymentUsecase,
		RejectPaymentUsecase:       rejectPaymentUsecase,
	}
}

func (ginApplication *ginApplication) Run() error {

	router := gin.Default()
	router.GET("/", HealthCheck)

	paymentsSubRouter := router.Group("/payments")
	paymentsSubRouter.POST("/", CreatePaymentHandler(ginApplication.CreatePaymentUsecase))
	paymentsSubRouter.PATCH("/", UpdatePaymentMethodHandler(ginApplication.UpdatePaymentMethodUsecase))
	paymentsSubRouter.PATCH("/:id/confirm", ConfirmPaymentHandler(ginApplication.ConfirmPaymentUsecase))
	paymentsSubRouter.PATCH("/:id/reject", RejectPaymentHandler(ginApplication.RejectPaymentUsecase))
	return router.Run()
}

func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, HealthCheckRes{Data: "Server is up and running"})
}

type HealthCheckRes struct {
	Data string `json:"data"`
}
