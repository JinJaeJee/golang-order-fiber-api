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
		expectedOrders []models.CleanedOrder
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
			expectedOrders: []models.CleanedOrder{
                {
                    No:          1,
                    ProductId:   "FG0A-CLEAR-IPHONE16PROMAX",
                    MaterialId:  "FG0A-CLEAR",
                    ModelId:     "IPHONE16PROMAX",
                    Qty:         2,
                    UnitPrice:   50,
                    TotalPrice:  100,
                },
                {
                    No:          2,
                    ProductId:   "WIPING-CLOTH",
					MaterialId: "",
					ModelId: "",
                    Qty:         2,
                    UnitPrice:   0,
                    TotalPrice:  0,
                },
                {
                    No:          3,
                    ProductId:   "CLEAR-CLEANNER",
					MaterialId: "",
					ModelId: "",
                    Qty:         2,
                    UnitPrice:   0,
                    TotalPrice:  0,
                },
            },
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
			expectedOrders: []models.CleanedOrder{
                {
                    No:          1,
                    ProductId:   "FG0A-CLEAR-IPHONE16PROMAX",
                    MaterialId:  "FG0A-CLEAR",
                    ModelId:     "IPHONE16PROMAX",
                    Qty:         2,
                    UnitPrice:   50,
                    TotalPrice:  100,
                },
                {
                    No:          2,
                    ProductId:   "WIPING-CLOTH",
					MaterialId: "",
					ModelId: "",
                    Qty:         2,
                    UnitPrice:   0,
                    TotalPrice:  0,
                },
                {
                    No:          3,
                    ProductId:   "CLEAR-CLEANNER",
					MaterialId: "",
					ModelId: "",
                    Qty:         2,
                    UnitPrice:   0,
                    TotalPrice:  0,
                },
            },
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
			expectedOrders: []models.CleanedOrder{
                {
                    No:          1,
                    ProductId:   "FG0A-MATTE-IPHONE16PROMAX",
                    MaterialId:  "FG0A-MATTE",
                    ModelId:     "IPHONE16PROMAX",
                    Qty:         3,
                    UnitPrice:   30,
                    TotalPrice:  90,
                },
                {
                    No:          2,
                    ProductId:   "WIPING-CLOTH",
					MaterialId: "",
					ModelId: "",
                    Qty:         3,
                    UnitPrice:   0,
                    TotalPrice:  0,
                },
                {
                    No:          3,
                    ProductId:   "MATTE-CLEANNER",
					MaterialId: "",
					ModelId: "",
                    Qty:         3,
                    UnitPrice:   0,
                    TotalPrice:  0,
                },
            },
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
			expectedOrders: []models.CleanedOrder{
			{
				No:          1,
				ProductId:   "FG0A-CLEAR-OPPOA3",
				MaterialId:  "FG0A-CLEAR",
				ModelId:     "OPPOA3",
				Qty:         1,
				UnitPrice:   40,
				TotalPrice:  40,
			},
			{
				No:          2,
				ProductId:   "FG0A-CLEAR-OPPOA3-B",
				MaterialId:  "FG0A-CLEAR",
				ModelId:     "OPPOA3-B",
				Qty:         1,
				UnitPrice:   40,
				TotalPrice:  40,
			},
			{
				No:          3,
				ProductId:   "WIPING-CLOTH",
				MaterialId: "",
				ModelId: "",
				Qty:         2,
				UnitPrice:   0,
				TotalPrice:  0,
			},
			{
				No:          4,
				ProductId:   "CLEAR-CLEANNER",
				MaterialId: "",
				ModelId: "",
				Qty:         2,
				UnitPrice:   0,
				TotalPrice:  0,
				},
			},
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
			expectedOrders: []models.CleanedOrder{
                {
                    No:          1,
                    ProductId:   "FG0A-CLEAR-OPPOA3",
                    MaterialId:  "FG0A-CLEAR",
                    ModelId:     "OPPOA3",
                    Qty:         1,
                    UnitPrice:   40,
                    TotalPrice:  40,
                },
                {
                    No:          2,
                    ProductId:   "FG0A-CLEAR-OPPOA3-B",
					MaterialId: "FG0A-CLEAR",
					ModelId: "OPPOA3-B",
                    Qty:         1,
                    UnitPrice:   40,
                    TotalPrice:  40,
                },
                {
                    No:          3,
                    ProductId:   "FG0A-MATTE-OPPOA3",
					MaterialId: "FG0A-MATTE",
					ModelId: "OPPOA3",
                    Qty:         1,
                    UnitPrice:   40,
                    TotalPrice:  40,
                },
				{
                    No:          4,
                    ProductId:   "WIPING-CLOTH",
                    MaterialId:  "",
                    ModelId:     "",
                    Qty:         3,
                    UnitPrice:   0,
                    TotalPrice:  0,
                },
                {
                    No:          5,
                    ProductId:   "CLEAR-CLEANNER",
					MaterialId: "",
					ModelId: "",
                    Qty:         2,
                    UnitPrice:   0,
                    TotalPrice:  0,
                },
                {
                    No:          6,
                    ProductId:   "MATTE-CLEANNER",
					MaterialId: "",
					ModelId: "",
                    Qty:         1,
                    UnitPrice:   0,
                    TotalPrice:  0,
                },
            },
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
			expectedOrders: []models.CleanedOrder{
                {
                    No:          1,
                    ProductId:   "FG0A-CLEAR-OPPOA3",
                    MaterialId:  "FG0A-CLEAR",
                    ModelId:     "OPPOA3",
                    Qty:         2,
                    UnitPrice:   40,
                    TotalPrice:  80,
                },
                {
                    No:          2,
                    ProductId:   "FG0A-MATTE-OPPOA3",
					MaterialId: "FG0A-MATTE",
					ModelId: "OPPOA3",
                    Qty:         1,
                    UnitPrice:   40,
                    TotalPrice:  40,
                },
                {
                    No:          3,
                    ProductId:   "WIPING-CLOTH",
					MaterialId: "",
					ModelId: "",
                    Qty:         3,
                    UnitPrice:   0,
                    TotalPrice:  0,
                },
				{
                    No:          4,
                    ProductId:   "CLEAR-CLEANNER",
                    MaterialId:  "",
                    ModelId:     "",
                    Qty:         2,
                    UnitPrice:   0,
                    TotalPrice:  0,
                },
                {
                    No:          5,
                    ProductId:   "MATTE-CLEANNER",
					MaterialId: "",
					ModelId: "",
                    Qty:         1,
                    UnitPrice:   0,
                    TotalPrice:  0,
                },
            },
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
			expectedOrders: []models.CleanedOrder{
                {
                    No:          1,
                    ProductId:   "FG0A-CLEAR-OPPOA3",
                    MaterialId:  "FG0A-CLEAR",
                    ModelId:     "OPPOA3",
                    Qty:         2,
                    UnitPrice:   40,
                    TotalPrice:  80,
                },
                {
                    No:          2,
                    ProductId:   "FG0A-MATTE-OPPOA3",
					MaterialId: "FG0A-MATTE",
					ModelId: "OPPOA3",
                    Qty:         2,
                    UnitPrice:   40,
                    TotalPrice:  80,
                },
                {
                    No:          3,
                    ProductId:   "FG0A-PRIVACY-IPHONE16PROMAX",
					MaterialId: "FG0A-PRIVACY",
					ModelId: "IPHONE16PROMAX",
                    Qty:         1,
                    UnitPrice:   50,
                    TotalPrice:  50,
                },
				{
                    No:          4,
                    ProductId:   "WIPING-CLOTH",
                    MaterialId:  "",
                    ModelId:     "",
                    Qty:         5,
                    UnitPrice:   0,
                    TotalPrice:  0,
                },
                {
                    No:          5,
                    ProductId:   "CLEAR-CLEANNER",
					MaterialId: "",
					ModelId: "",
                    Qty:         2,
                    UnitPrice:   0,
                    TotalPrice:  0,
                },
				{
                    No:          6,
                    ProductId:   "MATTE-CLEANNER",
					MaterialId: "",
					ModelId: "",
                    Qty:         2,
                    UnitPrice:   0,
                    TotalPrice:  0,
                },
				{
                    No:          7,
                    ProductId:   "PRIVACY-CLEANNER",
					MaterialId: "",
					ModelId: "",
                    Qty:         1,
                    UnitPrice:   0,
                    TotalPrice:  0,
                },
            },
		},
	}

	for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            cleanedOrders := services.ProcessOrders(tt.inputOrders)

            if len(cleanedOrders) != tt.expectedCount {
                t.Errorf("%s: Expected %d orders, got %d", tt.name, tt.expectedCount, len(cleanedOrders))
            }

            if tt.expectedOrders != nil {
                for i, expectedOrder := range tt.expectedOrders {
                    if cleanedOrders[i] != expectedOrder {
                        t.Errorf("%s: Expected order %v, got %v", tt.name, expectedOrder, cleanedOrders[i])
                    }
                }
            }
        })
    }
}