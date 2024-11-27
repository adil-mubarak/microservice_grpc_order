package services

import (
	"context"
	"fmt"
	"microservice_grpc_order/models"
	"microservice_grpc_order/pb/order"

	"gorm.io/gorm"
)

type OrderService struct {
	order.UnimplementedOrderServiceServer
	DB *gorm.DB
}

func (o *OrderService) CreateOrder(ctx context.Context, req *order.CreateOrderRequest) (*order.CreateOrderResponse, error) {

	orders := &models.Order{
		ProductID:      uint(req.ProductID),
		UserID:         uint(req.UserID),
		Quantity:       int64(req.Quantity),
		Order_status:   "Pending",
		Payment_status: "Unpaid",
	}
	if err := o.DB.Create(&orders).Error; err != nil {
		return &order.CreateOrderResponse{
			Status: "Failed to create order",
		}, err
	}

	return &order.CreateOrderResponse{
		OrderID: uint64(orders.ID),
		Status:  "order placed successully",
	}, nil
}

func (o *OrderService) GetOrders(ctx context.Context, req *order.GetOrdersRequest) (*order.GetOrdersResponse, error) {
	var orders []models.Order

	if err := o.DB.Where("user_id = ?", req.UserID).Find(&orders).Error; err != nil {
		return nil, fmt.Errorf("failed to fetch orders: %v", err)
	}

	var orderResponse []*order.Order
	for _, ord := range orders {
		orderResponse = append(orderResponse, &order.Order{
			OrderID:       uint64(ord.ID),
			ProductID:     uint64(ord.ProductID),
			UserID:        uint64(ord.UserID),
			Quantity:      int32(ord.Quantity),
			OrderStatus:   ord.Order_status,
			PaymentStatus: ord.Payment_status,
		})
	}

	return &order.GetOrdersResponse{
		Orders: orderResponse,
	}, nil
}

func (o *OrderService) UpdateOrderStatus(ctx context.Context, req *order.UpdateOrderStatusRequest) (*order.UpdateOrderStatusResponse, error) {
	var orders models.Order
	if err := o.DB.Where("id = ?", req.OrderID).First(&orders).Error; err != nil {
		return &order.UpdateOrderStatusResponse{
			Status: "Order not found",
		}, err
	}

	orders.Order_status = req.OrderStatus
	if err := o.DB.Save(&orders).Error; err != nil {
		return &order.UpdateOrderStatusResponse{
			Status: "Failed to update order status",
		}, err
	}

	return &order.UpdateOrderStatusResponse{
		OrderID: uint64(orders.ID),
		Status:  orders.Order_status,
	}, nil
}

func (o *OrderService) UpdatePaymentStatus(ctx context.Context, req *order.UpdatePaymentStatusRequest) (*order.UpdatePaymentStatusResponse, error) {
	var orders models.Order
	if err := o.DB.Where("id = ?", req.OrderID).First(&orders).Error; err != nil {
		return &order.UpdatePaymentStatusResponse{
			Status: "orderid is not found",
		}, err
	}

	orders.Payment_status = req.PaymentStatus
	if err := o.DB.Save(&orders).Error; err != nil {
		return &order.UpdatePaymentStatusResponse{
			Status: "Failed to update payment status",
		}, err
	}

	return &order.UpdatePaymentStatusResponse{
		OrderID: uint64(orders.ID),
		Status:  orders.Payment_status,
	}, nil
}

func (o *OrderService) CancelOrder(ctx context.Context, req *order.CalcelOrderRequest) (*order.CancelOrderResponse, error) {
	var orders models.Order
	if err := o.DB.Where("id = ?", req.OrderID).First(&orders).Error; err != nil {
		return &order.CancelOrderResponse{
			Status: "order id not found",
		}, err
	}

	orders.Order_status = "cancelled"
	if err := o.DB.Save(&orders).Error; err != nil {
		return &order.CancelOrderResponse{
			Status: "Failed to cancell order",
		}, err
	}

	return &order.CancelOrderResponse{
		OrderID: uint64(orders.ID),
		Status:  orders.Order_status,
	}, nil
}
