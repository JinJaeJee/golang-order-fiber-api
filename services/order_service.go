package services

import (
	"strconv"
	"strings"
	"fmt"

	"github.com/JinJaeJee/golang-order-fiber-api/models"
	"github.com/JinJaeJee/golang-order-fiber-api/utils"
)

func ProcessOrders(inputOrders []models.InputOrder) []models.CleanedOrder {
	var cleanedOrders []models.CleanedOrder
	orderNo := 1

	for _, inputOrder := range inputOrders {
		productIds, count, extraCount := utils.ParseProductId(inputOrder.PlatformProductId, inputOrder.Qty)
		for _, productId := range productIds {
			materialId, modelId := utils.ExtractMaterialAndModelId(productId)

			baseProductId, quantity, unitPrice := utils.HandleQuantityMultiplier(productId, inputOrder.Qty, inputOrder.UnitPrice)

			cleanedOrder := models.CleanedOrder{
				No:         orderNo,
				ProductId:  baseProductId,
				MaterialId: materialId,
				ModelId:    modelId,
				Qty:        quantity,
				UnitPrice:  unitPrice / float64(count+extraCount-1),
				TotalPrice: (unitPrice * float64(quantity)) / float64(count+extraCount-1),
			}
			cleanedOrders = append(cleanedOrders, cleanedOrder)
			orderNo++
		}
		cleanedOrders = append(cleanedOrders, addComplementaryItems(count*inputOrder.Qty, extraCount, orderNo)...)
		orderNo++
		cleanedOrders = append(cleanedOrders, addCleanerItem(productIds, inputOrder.Qty, extraCount, orderNo)...)
		orderNo++
	}

	reOrderCleanedOrders := mergeProducts(cleanedOrders)

	return reOrderCleanedOrders
}

func addComplementaryItems(qty int, extraQty int, orderNo int) []models.CleanedOrder {
	var complementaryItems []models.CleanedOrder

	complementaryItems = append(complementaryItems, models.CleanedOrder{
		No:         orderNo,
		ProductId:  "WIPING-CLOTH",
		Qty:        qty + extraQty - 1,
		UnitPrice:  0.00,
		TotalPrice: 0.00,
	})

	return complementaryItems
}

func addCleanerItem(platformProductIds []string, qty int, extraQty int, orderNo int) []models.CleanedOrder {
	cleanerCount := make(map[string]int)

	for _, productId := range platformProductIds {

		mainParts := strings.Split(productId, "*")
		baseId := mainParts[0]
		itemQty := 1

		if len(mainParts) == 2 {
			parsedQty, err := strconv.Atoi(mainParts[1])
			if err == nil {
				itemQty = parsedQty
			}
		}

		parts := strings.Split(baseId, "-")
		if len(parts) < 3 {
			continue
		}

		var cleanerId string
		switch {
		case strings.Contains(baseId, "CLEAR"):
			cleanerId = "CLEAR-CLEANNER"
		case strings.Contains(baseId, "MATTE"):
			cleanerId = "MATTE-CLEANNER"
		case strings.Contains(baseId, "PRIVACY"):
			cleanerId = "PRIVACY-CLEANNER"
		default:
			continue
		}
		cleanerCount[cleanerId] += itemQty
	}

	var cleanerItems []models.CleanedOrder
	for cleanerId, count := range cleanerCount {
		cleanerItems = append(cleanerItems, models.CleanedOrder{
			No:         orderNo,
			ProductId:  cleanerId,
			MaterialId: "",
			ModelId:    "",
			Qty:        count * qty,
			UnitPrice:  0.00,
			TotalPrice: 0.00,
		})
		orderNo++
	}

	return cleanerItems
}

func mergeProducts(products []models.CleanedOrder) []models.CleanedOrder {
	productMap := make(map[string]models.CleanedOrder)
	var mergedProducts []models.CleanedOrder

	for _, p := range products {
		key := fmt.Sprintf("%s-%s-%s", p.ProductId, p.MaterialId, p.ModelId)

		if existingProduct, ok := productMap[key]; ok {
			existingProduct.Qty += p.Qty
			existingProduct.TotalPrice += p.TotalPrice
			productMap[key] = existingProduct
		} else {
			productMap[key] = p
		}
	}

	for _, p := range productMap {
		mergedProducts = append(mergedProducts, p)
	}


	for i := range mergedProducts{
		mergedProducts[i].No = i+1
	}

	return mergedProducts
}

