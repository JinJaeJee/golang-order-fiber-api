package services

import (
	"strings"

	"github.com/JinJaeJee/golang-order-fiber-api/models"
	"github.com/JinJaeJee/golang-order-fiber-api/utils"
)

func ProcessOrders(inputOrders []models.InputOrder) []models.CleanedOrder {
	var cleanedOrders []models.CleanedOrder
	orderNo := 1

	for _, inputOrder := range inputOrders {
		productIds, count, extraCount := utils.ParseProductId(inputOrder.PlatformProductId)
		for _, productId := range productIds {
			materialId, modelId := utils.ExtractMaterialAndModelId(productId)

			baseProductId, quantity, unitPrice := utils.HandleQuantityMultiplier(productId, inputOrder.Qty, inputOrder.UnitPrice)

			cleanedOrder := models.CleanedOrder{
				No:         orderNo,
				ProductId:  baseProductId,
				MaterialId: materialId,
				ModelId:    modelId,
				Qty:        quantity,
				UnitPrice:  unitPrice / float64(count+extraCount),
				TotalPrice: (unitPrice * float64(quantity)) / float64(count+extraCount),
			}
			cleanedOrders = append(cleanedOrders, cleanedOrder)
			orderNo++
		}
		cleanedOrders = append(cleanedOrders, addComplementaryItems(count, extraCount, orderNo)...)
		orderNo++
		cleanedOrders = append(cleanedOrders, addCleanerItem(inputOrder.PlatformProductId, orderNo)...)
		orderNo++
	}

	return cleanedOrders
}

func addComplementaryItems(qty int, extraQty int, orderNo int) []models.CleanedOrder {
	var complementaryItems []models.CleanedOrder

	complementaryItems = append(complementaryItems, models.CleanedOrder{
		No:         orderNo,
		ProductId:  "WIPING-CLOTH",
		Qty:        qty + extraQty,
		UnitPrice:  0.00,
		TotalPrice: 0.00,
	})

	return complementaryItems
}

func addCleanerItem(platformProductId string, orderNo int) []models.CleanedOrder {
	var cleanerItems []models.CleanedOrder
	textureCount := make(map[string]int)

	// Split product IDs
	productIds := strings.Split(platformProductId, "/")
	for _, productId := range productIds {
		parts := strings.Split(productId, "-")
		if len(parts) >= 2 {
			texture := strings.ToUpper(parts[1])
			textureCount[texture]++
		}
	}

	for texture, count := range textureCount {
		cleanerId := texture + "-CLEANER"
		cleanerItems = append(cleanerItems, models.CleanedOrder{
			No:         orderNo,
			ProductId:  cleanerId,
			Qty:        count,
			UnitPrice:  0.00,
			TotalPrice: 0.00,
		})
		orderNo++
	}

	return cleanerItems
}
