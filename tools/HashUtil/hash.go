package HashUtil

// Additive Hash（加法哈希）
func Additive(data []byte) uint32 {
	var hash uint32
	for _, b := range data {
		hash += uint32(b)
	}
	return hash
}

// Rotating Hash（旋转哈希）
func Rotating(data []byte) uint32 {
	var hash uint32
	for _, b := range data {
		hash = (hash << 4) + uint32(b)
		hash ^= hash >> 16
	}
	return hash
}

// OneByOne One-by-One Hash（一字节哈希）
func OneByOne(data []byte) uint32 {
	var hash uint32 = 0
	for _, b := range data {
		hash = hash*31 + uint32(b)
	}
	return hash
}

// Bernstein's Hash（Bernstein哈希）
func Bernstein(data []byte) uint32 {
	var hash uint32 = 5381
	for _, b := range data {
		hash = ((hash << 5) + hash) + uint32(b)
	}
	return hash
}

// Universal Hashing（通用哈希）
func Universal(data []byte, a, b, p, m uint32) uint32 {
	var h uint32
	for _, v := range data {
		h = (h + a*uint32(v) + b) % p
	}
	return h % m
}

// Zobrist Hashing（Zobrist哈希）
func Zobrist(data []byte, table [][]uint32) uint32 {
	var hash uint32
	for i, b := range data {
		hash ^= table[i][b]
	}
	return hash
}

// FNV Hash（改进的32位FNV算法）
func FnvHash(data []byte) uint32 {
	const fnvOffsetBasis uint32 = 2166136261
	const fnvPrime uint32 = 16777619

	hash := fnvOffsetBasis
	for _, b := range data {
		hash ^= uint32(b)
		hash *= fnvPrime
	}
	return hash
}

// Thomas Wang’s Hash（Thomas Wang整数哈希）
func IntHash(data uint32) uint32 {
	data = (data ^ (data >> 16)) * 0x45d9f3b
	data = (data ^ (data >> 16)) * 0x45d9f3b
	data = data ^ (data >> 16)
	return data
}

// RS Hash（RS哈希）
func RsHash(data []byte) uint32 {
	var b uint32 = 378551
	var a uint32 = 63689
	var hash uint32

	for _, c := range data {
		hash = hash*a + uint32(c)
		a *= b
	}
	return hash
}

// JsHash JS Hash（JS哈希）
func JsHash(data []byte) uint32 {
	var hash uint32 = 1315423911
	for _, b := range data {
		hash ^= ((hash << 5) + uint32(b) + (hash >> 2))
	}
	return hash
}

// ElfHash ELF Hash（ELF哈希）
func ElfHash(data []byte) uint32 {
	var hash uint32
	var x uint32

	for _, b := range data {
		hash = (hash << 4) + uint32(b)
		if x = hash & 0xF0000000; x != 0 {
			hash ^= (x >> 24)
		}
		hash &= ^x
	}
	return hash
}

// BkdrHash BKDR Hash（BKDR哈希）
func BkdrHash(data []byte) uint32 {
	var hash uint32 = 0
	const seed uint32 = 131

	for _, b := range data {
		hash = (hash * seed) + uint32(b)
	}
	return hash
}

// SdbmHash SDBM Hash（SDBM哈希）
func SdbmHash(data []byte) uint32 {
	var hash uint32
	for _, b := range data {
		hash = uint32(b) + (hash << 6) + (hash << 16) - hash
	}
	return hash
}

// DjbHash DJB Hash（DJB哈希）
func DjbHash(data []byte) uint32 {
	var hash uint32 = 5381
	for _, b := range data {
		hash = ((hash << 5) + hash) + uint32(b)
	}
	return hash
}

// DekHash DEK Hash（DEK哈希）
func DekHash(data []byte) uint32 {
	var hash uint32 = uint32(len(data))
	for _, b := range data {
		hash = ((hash << 5) ^ (hash >> 27)) ^ uint32(b)
	}
	return hash
}

// ApHash AP Hash（AP哈希）
func ApHash(data []byte) uint32 {
	var hash uint32 = 0xAAAAAAAA
	for i, b := range data {
		if i&1 == 0 {
			hash ^= (hash << 7) ^ uint32(b)*(hash>>3)
		} else {
			hash ^= ^((hash << 11) + uint32(b) ^ (hash >> 5))
		}
	}
	return hash
}

// TianlHash TianL Hash（TianL哈希）
func TianlHash(data []byte) uint32 {
	var hash uint32 = 0
	var a uint32 = 0x7FFFFFFF

	for _, b := range data {
		hash = ((hash << 5) - hash) + uint32(b)
		a ^= hash
	}
	return a
}

// PJWHash 计算给定字符串的PJW哈希值
func PJWHash(key string) uint32 {
	var h uint32 = 0
	for i := 0; i < len(key); i++ {
		h = (h << 4) + uint32(key[i])
		if g := h & 0xF0000000; g != 0 {
			h = g >> 24
			h = g
		}
	}
	return h
}
