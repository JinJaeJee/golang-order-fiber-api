package services

import (
	"testing"

	"github.com/JinJaeJee/golang-order-fiber-api/models"
	"github.com/JinJaeJee/golang-order-fiber-api/services"
)

func TestProcessOrders(t *testing.T) {
	tests := []struct {
		name          string
		inputOrders   []models.InputOrder
		expectedCount int
	}{
		{
			name: "Case 1: Only one product",
			inputOrders: []models.InputOrder{
				{
					No:                1,
					PlatformProductId: "FG0A-CLEAR-IPHONE16PROMAX",
					Qty:               2,
					UnitPrice:         50,
					TotalPrice:        100,
				},
			},
			expectedCount: 3,
		},
		{
			name: "Case 2: One product with wrong prefix",
			inputOrders: []models.InputOrder{
				{
					No:                1,
					PlatformProductId: "x2-3&FG0A-CLEAR-IPHONE16PROMAX",
					Qty:               2,
					UnitPrice:         50,
					TotalPrice:        100,
				},
			},
			expectedCount: 3,
		},
		{
			name: "Case 3: One product with wrong prefix and has * symbol that indicates the quantity",
			inputOrders: []models.InputOrder{
				{
					No:                1,
					PlatformProductId: "x2-3&FG0A-MATTE-IPHONE16PROMAX*3",
					Qty:               1,
					UnitPrice:         90,
					TotalPrice:        90,
				},
			},
			expectedCount: 3,
		},
		{
			name: "Case 4: One bundle product split by /",
			inputOrders: []models.InputOrder{
				{
					No:                1,
					PlatformProductId: "FG0A-CLEAR-OPPOA3/%20xFG0A-CLEAR-OPPOA3-B",
					Qty:               1,
					UnitPrice:         80,
					TotalPrice:        80,
				},
			},
			expectedCount: 4,
		},
		{
			name: "Case 5: One bundle product split by / into three products",
			inputOrders: []models.InputOrder{
				{
					No:                1,
					PlatformProductId: "FG0A-CLEAR-OPPOA3/%20xFG0A-CLEAR-OPPOA3-B/FG0A-MATTE-OPPOA3",
					Qty:               1,
					UnitPrice:         120,
					TotalPrice:        120,
				},
			},
			expectedCount: 6,
		},
		{
			name: "Case 6: One bundle product with * and / symbols",
			inputOrders: []models.InputOrder{
				{
					No:                1,
					PlatformProductId: "--FG0A-CLEAR-OPPOA3*2/FG0A-MATTE-OPPOA3",
					Qty:               1,
					UnitPrice:         120,
					TotalPrice:        120,
				},
			},
			expectedCount: 5,
		},
		{
			name: "Case 7: One product and one bundle product",
			inputOrders: []models.InputOrder{
				{
					No:                1,
					PlatformProductId: "--FG0A-CLEAR-OPPOA3*2/FG0A-MATTE-OPPOA3*2",
					Qty:               1,
					UnitPrice:         160,
					TotalPrice:        160,
				},
				{
					No:                2,
					PlatformProductId: "FG0A-PRIVACY-IPHONE16PROMAX",
					Qty:               1,
					UnitPrice:         50,
					TotalPrice:        50,
				},
			},
			expectedCount: 7,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cleanedOrders := services.ProcessOrders(tt.inputOrders)

			if len(cleanedOrders) != tt.expectedCount {
				t.Errorf("%s: Expected %d orders, got %d", tt.name, tt.expectedCount, len(cleanedOrders))
			}
		})
	}
}
