package utils

import (
	"sort"

	"github.com/AlejandroAldana99/yalo-challenge/models"
)

func RankTopProducts(interactions []models.UserInteraction) []string {
	ranks := make(map[string]int)

	for _, interaction := range interactions {
		ranks[interaction.ProductSKU] += interaction.Duration
	}

	type productRank struct {
		ProductSKU string
		Count      int
	}

	ranksList := make([]productRank, 0, len(ranks))
	for sku, count := range ranks {
		ranksList = append(ranksList, productRank{ProductSKU: sku, Count: count})
	}

	sort.Slice(ranksList, func(i, j int) bool {
		return ranksList[i].Count > ranksList[j].Count
	})

	topProducts := make([]string, 0, 3)
	for i := 0; i < len(ranksList) && i < 3; i++ {
		topProducts = append(topProducts, ranksList[i].ProductSKU)
	}

	return topProducts
}
