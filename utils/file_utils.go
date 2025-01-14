package utils

import (
	"encoding/json"
	"errors"
	"go-react-graphql-orders/model"
	"os"
)

func LoadOrdersFromFile(filePath string) ([]*model.Order, error) {
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	var orders []*model.Order
	if err := json.Unmarshal(fileContent, &orders); err != nil {
		return nil, err
	}
	return orders, nil
}

func SaveOrderToFile(order *model.Order, filePath string) error {
	orders, err := LoadOrdersFromFile(filePath)
	if err != nil {
		return err
	}
	orders = append(orders, order)
	data, err := json.Marshal(orders)
	if err != nil {
		return err
	}
	return os.WriteFile(filePath, data, 0644)
}

func UpdateOrderInFile(id string, newOrder *model.Order, filePath string) (*model.Order, error) {
	// Load all orders from the file
	orders, err := LoadOrdersFromFile(filePath)
	if err != nil {
		return nil, err
	}

	// Find and update the matching order
	for i, order := range orders {
		if order.ID == id {
			// Update only relevant fields; ensure the ID remains unchanged
			newOrder.ID = id
			newOrder.Name = order.Name
			newOrder.CreatedAt = order.CreatedAt
			newOrder.Status = order.Status
			orders[i] = newOrder

			// Serialize and write the updated orders back to the file
			data, err := json.Marshal(orders)
			if err != nil {
				return nil, err
			}
			err = os.WriteFile(filePath, data, 0644)
			if err != nil {
				return nil, err
			}

			return orders[i], nil // Return the updated order
		}
	}

	// Return an error if no matching order was found
	return nil, errors.New("order not found")
}
