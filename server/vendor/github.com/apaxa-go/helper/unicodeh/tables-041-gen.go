package unicodeh

import "unicode"

// Unicode property "Default_Ignorable_Code_Point" (known as "DI", "Default_Ignorable_Code_Point").
// Kind of property: "Binary".
// Based on file "DerivedCoreProperties.txt".
var (
	DefaultIgnorableCodePointNo  = defaultIgnorableCodePointNo  // Value "No" (known as "N", "No", "F", "False").
	DefaultIgnorableCodePointYes = defaultIgnorableCodePointYes // Value "Yes" (known as "Y", "Yes", "T", "True").
)

var (
	defaultIgnorableCodePointNo  = &unicode.RangeTable{[]unicode.Range16{{0x0, 0xac, 0x1}, {0xae, 0x34e, 0x1}, {0x350, 0x61b, 0x1}, {0x61d, 0x115e, 0x1}, {0x1161, 0x17b3, 0x1}, {0x17b6, 0x180a, 0x1}, {0x180f, 0x200a, 0x1}, {0x2010, 0x2029, 0x1}, {0x202f, 0x205f, 0x1}, {0x2070, 0x3163, 0x1}, {0x3165, 0xfdff, 0x1}, {0xfe10, 0xfefe, 0x1}, {0xff00, 0xff9f, 0x1}, {0xffa1, 0xffef, 0x1}, {0xfff9, 0xffff, 0x1}}, []unicode.Range32{{0x10000, 0x1bc9f, 0x1}, {0x1bca4, 0x1d172, 0x1}, {0x1d17b, 0xdffff, 0x1}, {0xe1000, 0x10ffff, 0x1}}, 1}
	defaultIgnorableCodePointYes = &unicode.RangeTable{[]unicode.Range16{{0xad, 0x34f, 0x2a2}, {0x61c, 0x115f, 0xb43}, {0x1160, 0x17b4, 0x654}, {0x17b5, 0x180b, 0x56}, {0x180c, 0x180e, 0x1}, {0x200b, 0x200f, 0x1}, {0x202a, 0x202e, 0x1}, {0x2060, 0x206f, 0x1}, {0x3164, 0xfe00, 0xcc9c}, {0xfe01, 0xfe0f, 0x1}, {0xfeff, 0xffa0, 0xa1}, {0xfff0, 0xfff8, 0x1}}, []unicode.Range32{{0x1bca0, 0x1bca3, 0x1}, {0x1d173, 0x1d17a, 0x1}, {0xe0000, 0xe0fff, 0x1}}, 0}
)
