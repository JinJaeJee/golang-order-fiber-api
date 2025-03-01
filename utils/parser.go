package utils

import (
	"regexp"
	"strconv"
	"strings"
)

func ParseProductId(platformProductId string, Qty int) ([]string, int, int) {

	cleanedId := strings.TrimLeft(platformProductId, "-%20x3&")
	cleanedId = strings.ReplaceAll(cleanedId, "%20x", "")
	cleanedId = strings.ReplaceAll(cleanedId, "--", "")

	productIds := strings.Split(cleanedId, "/")

	extraCount := 0

	for _, product := range productIds {
		parts := strings.Split(product, "*")
		if len(parts) == 2 {
			if count, err := strconv.Atoi(parts[1]); err == nil {
				extraCount += count - 1
			}
		}
	}

	return productIds, len(productIds), extraCount + 1
}

func ExtractMaterialAndModelId(productId string) (string, string) {
	re := regexp.MustCompile(`\*\d+$`)
	cleanedProductId := re.ReplaceAllString(productId, "")
	parts := strings.Split(cleanedProductId, "-")
	materialId := parts[0] + "-" + parts[1]
	modelId := strings.TrimPrefix(cleanedProductId, materialId+"-")
	return materialId, modelId
}

func HandleQuantityMultiplier(productId string, originalQty int, unitPrice float64) (string, int, float64) {
	if strings.Contains(productId, "*") {
		parts := strings.Split(productId, "*")
		if len(parts) == 2 {
			baseProductId := parts[0]
			multiplier := parts[1]
			newQty := originalQty * StringToInt(multiplier)
			newUnitPrice := unitPrice
			return baseProductId, newQty, newUnitPrice
		}
	}
	return productId, originalQty, unitPrice
}

func ExtractTexture(platformProductId string) string {
	parts := strings.Split(platformProductId, "-")
	if len(parts) >= 2 {
		return parts[1]
	}
	return ""
}

func StringToInt(s string) int {
	var result int
	for _, char := range s {
		if char >= '0' && char <= '9' {
			result = result*10 + int(char-'0')
		}
	}
	return result
}
