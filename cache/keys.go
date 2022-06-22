package cache

import (
	"fmt"
	"strconv"
)

const (
	// RankKey 每日排行
	RankKey = "rank"
	// ElectricalRank 家电排行
	ElectricalRank = "elecRank"
	// AccessoryRank 配件排行
	AccessoryRank = "acceRank"
)

// ProductViewKey 视频点击数的key
func ProductViewKey(id uint) string {
	return fmt.Sprintf("view:product:%s", strconv.Itoa(int(id)))
}
