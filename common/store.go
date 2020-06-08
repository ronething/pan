// author: ashing
// time: 2020/6/8 1:23 下午
// mail: axingfly@gmail.com
// Less is more.

package common

type StoreType int

const (
	_ StoreType = iota
	// 本地磁盘存储
	StoreLocal
	// Ceph
	StoreCeph
	// 阿里云 oss
	StoreOSS
	// 混合
	StoreMix
	// 所有类型均存储
	StoreAll
)
