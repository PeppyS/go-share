package unicodeh

import "unicode"

// Unicode property "Other_Lowercase" (known as "OLower", "Other_Lowercase").
// Kind of property: "Binary".
// Based on file "PropList.txt".
var (
	OtherLowercaseNo  = otherLowercaseNo  // Value "No" (known as "N", "No", "F", "False").
	OtherLowercaseYes = otherLowercaseYes // Value "Yes" (known as "Y", "Yes", "T", "True").
)

var (
	otherLowercaseNo  = &unicode.RangeTable{[]unicode.Range16{{0x0, 0xa9, 0x1}, {0xab, 0xb9, 0x1}, {0xbb, 0x2af, 0x1}, {0x2b9, 0x2bf, 0x1}, {0x2c2, 0x2df, 0x1}, {0x2e5, 0x344, 0x1}, {0x346, 0x379, 0x1}, {0x37b, 0x1d2b, 0x1}, {0x1d6b, 0x1d77, 0x1}, {0x1d79, 0x1d9a, 0x1}, {0x1dc0, 0x2070, 0x1}, {0x2072, 0x207e, 0x1}, {0x2080, 0x208f, 0x1}, {0x209d, 0x216f, 0x1}, {0x2180, 0x24cf, 0x1}, {0x24ea, 0x2c7b, 0x1}, {0x2c7e, 0xa69b, 0x1}, {0xa69e, 0xa76f, 0x1}, {0xa771, 0xa7f7, 0x1}, {0xa7fa, 0xab5b, 0x1}, {0xab60, 0xffff, 0x1}}, []unicode.Range32{{0x10000, 0x10ffff, 0x1}}, 2}
	otherLowercaseYes = &unicode.RangeTable{[]unicode.Range16{{0xaa, 0xba, 0x10}, {0x2b0, 0x2b8, 0x1}, {0x2c0, 0x2c1, 0x1}, {0x2e0, 0x2e4, 0x1}, {0x345, 0x37a, 0x35}, {0x1d2c, 0x1d6a, 0x1}, {0x1d78, 0x1d9b, 0x23}, {0x1d9c, 0x1dbf, 0x1}, {0x2071, 0x207f, 0xe}, {0x2090, 0x209c, 0x1}, {0x2170, 0x217f, 0x1}, {0x24d0, 0x24e9, 0x1}, {0x2c7c, 0x2c7d, 0x1}, {0xa69c, 0xa69d, 0x1}, {0xa770, 0xa7f8, 0x88}, {0xa7f9, 0xab5c, 0x363}, {0xab5d, 0xab5f, 0x1}}, nil, 1}
)
