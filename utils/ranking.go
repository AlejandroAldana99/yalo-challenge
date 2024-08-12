package utils

import "sort"

func RankTopProducts(productRanks map[string]int) []string {
	type productRank struct {
		ProductSKU string
		Count      int
	}

	ranks := make([]productRank, 0, len(productRanks))
	for productSKU, count := range productRanks {
		ranks = append(ranks, productRank{ProductSKU: productSKU, Count: count})
	}

	sort.Slice(ranks, func(i, j int) bool {
		return ranks[i].Count > ranks[j].Count
	})

	topProducts := []string{}
	for i := 0; i < len(ranks) && i < 3; i++ {
		topProducts = append(topProducts, ranks[i].ProductSKU)
	}

	return topProducts
}
