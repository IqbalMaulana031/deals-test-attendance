package constant

import (
	"strings"
)

// FindPermission is a service for find permission
func FindPermission(permission, menu string) string {
	menu = strings.ReplaceAll(menu, "-", "_")
	return "can_" + permission + "_" + menu
}

const (
	// Access is access permission
	Access = "access"
	// Create is a create permission
	Create = "create"
	// Update is a update permission
	Update = "update"
	// Delete is a delete permission
	Delete = "delete"
	// Read is a read permission
	Read = "read"
	// Login is a login permission
	Login = "login"
	// MerchantInfo is a merchant info menu
	MerchantInfo = "merchant_info"
	// 	OperationTime is a operation time permission
	OperationTime = "operation_time"
	// Employee is a employee menu
	Employee = "employee"

	Role = "role"
	// Customer is a customer menu
	Customer = "customer"
	// PaymentChannel is a payment channel menu
	PaymentChannel = "payment_channel"
)
