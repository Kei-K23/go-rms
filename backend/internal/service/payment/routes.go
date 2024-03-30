package payment

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Kei-K23/go-rms/backend/internal/types"
	"github.com/Kei-K23/go-rms/backend/internal/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/checkout/session"
)

type Handler struct {
	orderItemStore types.OrderItemStore
	paymentStore   types.PaymentStore
}

func NewHandler(orderItemStore types.OrderItemStore, paymentStore types.PaymentStore) *Handler {
	return &Handler{orderItemStore: orderItemStore, paymentStore: paymentStore}
}

// TODO: make repo pattern
func (h *Handler) RegisterRoute(router fiber.Router) {
	router.Post("/checkout", h.createCheckoutSession)
	router.Get("/success", h.checkoutSuccess)
	router.Post("/cancel", h.checkoutCancel)
}

func (h *Handler) checkoutSuccess(c *fiber.Ctx) error {
	oID := c.Query("order_id")
	amount := c.Query("amount")
	paymentType := c.Query("payment_type")

	if oID == "" || amount == "" || paymentType == "" {
		return utils.WriteError(c, http.StatusBadRequest, fmt.Errorf("missing query parameters"))
	}

	intOId, err := strconv.Atoi(oID)
	if err != nil {
		return utils.WriteError(c, http.StatusBadRequest, err)
	}

	intAmount, err := strconv.Atoi(amount)
	if err != nil {
		return utils.WriteError(c, http.StatusBadRequest, err)
	}

	err = h.paymentStore.CreatePayment(types.CreatePayment{
		OrderID:     intOId,
		Amount:      intAmount,
		PaymentType: paymentType,
	})
	if err != nil {
		return utils.WriteError(c, http.StatusInternalServerError, err)
	}
	return c.SendString("Payment successful. Thank you!")
}

func (h *Handler) checkoutCancel(c *fiber.Ctx) error {
	return c.SendString("cancel")
}

func (h *Handler) createCheckoutSession(c *fiber.Ctx) error {
	var payload types.CreatePayment
	var stripeLineItems []*stripe.CheckoutSessionLineItemParams
	if err := utils.ParseJson(c, &payload); err != nil {
		return utils.WriteError(c, http.StatusBadRequest, err)
	}
	if err := utils.ValidatePayload(payload); err != nil {
		return utils.WriteError(c, http.StatusBadRequest, err)
	}

	orderItems, err := h.orderItemStore.GetOrderItemByOrderID(payload.OrderID)
	if err != nil {
		return utils.WriteError(c, http.StatusInternalServerError, err)
	}
	fmt.Println(orderItems)
	for _, oi := range orderItems {
		orderItem := oi
		fmt.Println(orderItem.Price)
		stripeLineItem := &stripe.CheckoutSessionLineItemParams{
			PriceData: &stripe.CheckoutSessionLineItemPriceDataParams{
				Currency: stripe.String("usd"),
				ProductData: &stripe.CheckoutSessionLineItemPriceDataProductDataParams{
					Name: stripe.String(string(orderItem.MenuID)),
				},
				UnitAmount: stripe.Int64(int64(orderItem.Price)),
			},
			Quantity: stripe.Int64(int64(orderItem.Quantity)),
		}

		fmt.Printf("Created line item: %+v\n", stripeLineItem)
		stripeLineItems = append(stripeLineItems, stripeLineItem)
	}

	params := &stripe.CheckoutSessionParams{
		Mode:       stripe.String(string(stripe.CheckoutSessionModePayment)),
		LineItems:  stripeLineItems,
		SuccessURL: stripe.String(fmt.Sprintf("http://localhost:4000/api/v1/success/?order_id=%d&amount=%d&payment_type=%s", payload.OrderID, payload.Amount, payload.PaymentType)),
		CancelURL:  stripe.String("http://localhost:4000/api/v1/cancel"),
	}

	s, err := session.New(params)
	if err != nil {
		return err
	}
	fmt.Println(s.URL)
	return c.Redirect(s.URL, http.StatusSeeOther)
}
