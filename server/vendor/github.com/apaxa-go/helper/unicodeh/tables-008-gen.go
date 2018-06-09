package unicodeh

import "unicode"

// Unicode property "East_Asian_Width" (known as "ea", "East_Asian_Width").
// Kind of property: "Enumerated".
// Based on file "extracted\DerivedEastAsianWidth.txt".
var (
	EastAsianWidthAmbiguous = eastAsianWidthAmbiguous // Value "Ambiguous" (known as "A", "Ambiguous").
	EastAsianWidthFullwidth = eastAsianWidthFullwidth // Value "Fullwidth" (known as "F", "Fullwidth").
	EastAsianWidthHalfwidth = eastAsianWidthHalfwidth // Value "Halfwidth" (known as "H", "Halfwidth").
	EastAsianWidthNeutral   = eastAsianWidthNeutral   // Value "Neutral" (known as "N", "Neutral").
	EastAsianWidthNarrow    = eastAsianWidthNarrow    // Value "Narrow" (known as "Na", "Narrow").
	EastAsianWidthWide      = eastAsianWidthWide      // Value "Wide" (known as "W", "Wide").
)

var (
	eastAsianWidthAmbiguous = &unicode.RangeTable{[]unicode.Range16{{0xa1, 0xa7, 0x3}, {0xa8, 0xaa, 0x2}, {0xad, 0xae, 0x1}, {0xb0, 0xb4, 0x1}, {0xb6, 0xba, 0x1}, {0xbc, 0xbf, 0x1}, {0xc6, 0xd0, 0xa}, {0xd7, 0xd8, 0x1}, {0xde, 0xe1, 0x1}, {0xe6, 0xe8, 0x2}, {0xe9, 0xea, 0x1}, {0xec, 0xed, 0x1}, {0xf0, 0xf2, 0x2}, {0xf3, 0xf7, 0x4}, {0xf8, 0xfa, 0x1}, {0xfc, 0xfe, 0x2}, {0x101, 0x111, 0x10}, {0x113, 0x11b, 0x8}, {0x126, 0x127, 0x1}, {0x12b, 0x131, 0x6}, {0x132, 0x133, 0x1}, {0x138, 0x13f, 0x7}, {0x140, 0x142, 0x1}, {0x144, 0x148, 0x4}, {0x149, 0x14b, 0x1}, {0x14d, 0x152, 0x5}, {0x153, 0x166, 0x13}, {0x167, 0x16b, 0x4}, {0x1ce, 0x1dc, 0x2}, {0x251, 0x261, 0x10}, {0x2c4, 0x2c7, 0x3}, {0x2c9, 0x2cb, 0x1}, {0x2cd, 0x2d0, 0x3}, {0x2d8, 0x2db, 0x1}, {0x2dd, 0x2df, 0x2}, {0x300, 0x36f, 0x1}, {0x391, 0x3a1, 0x1}, {0x3a3, 0x3a9, 0x1}, {0x3b1, 0x3c1, 0x1}, {0x3c3, 0x3c9, 0x1}, {0x401, 0x410, 0xf}, {0x411, 0x44f, 0x1}, {0x451, 0x2010, 0x1bbf}, {0x2013, 0x2016, 0x1}, {0x2018, 0x2019, 0x1}, {0x201c, 0x201d, 0x1}, {0x2020, 0x2022, 0x1}, {0x2024, 0x2027, 0x1}, {0x2030, 0x2032, 0x2}, {0x2033, 0x2035, 0x2}, {0x203b, 0x203e, 0x3}, {0x2074, 0x207f, 0xb}, {0x2081, 0x2084, 0x1}, {0x20ac, 0x2103, 0x57}, {0x2105, 0x2109, 0x4}, {0x2113, 0x2116, 0x3}, {0x2121, 0x2122, 0x1}, {0x2126, 0x212b, 0x5}, {0x2153, 0x2154, 0x1}, {0x215b, 0x215e, 0x1}, {0x2160, 0x216b, 0x1}, {0x2170, 0x2179, 0x1}, {0x2189, 0x2190, 0x7}, {0x2191, 0x2199, 0x1}, {0x21b8, 0x21b9, 0x1}, {0x21d2, 0x21d4, 0x2}, {0x21e7, 0x2200, 0x19}, {0x2202, 0x2203, 0x1}, {0x2207, 0x2208, 0x1}, {0x220b, 0x220f, 0x4}, {0x2211, 0x2215, 0x4}, {0x221a, 0x221d, 0x3}, {0x221e, 0x2220, 0x1}, {0x2223, 0x2227, 0x2}, {0x2228, 0x222c, 0x1}, {0x222e, 0x2234, 0x6}, {0x2235, 0x2237, 0x1}, {0x223c, 0x223d, 0x1}, {0x2248, 0x224c, 0x4}, {0x2252, 0x2260, 0xe}, {0x2261, 0x2264, 0x3}, {0x2265, 0x2267, 0x1}, {0x226a, 0x226b, 0x1}, {0x226e, 0x226f, 0x1}, {0x2282, 0x2283, 0x1}, {0x2286, 0x2287, 0x1}, {0x2295, 0x2299, 0x4}, {0x22a5, 0x22bf, 0x1a}, {0x2312, 0x2460, 0x14e}, {0x2461, 0x24e9, 0x1}, {0x24eb, 0x254b, 0x1}, {0x2550, 0x2573, 0x1}, {0x2580, 0x258f, 0x1}, {0x2592, 0x2595, 0x1}, {0x25a0, 0x25a1, 0x1}, {0x25a3, 0x25a9, 0x1}, {0x25b2, 0x25b3, 0x1}, {0x25b6, 0x25b7, 0x1}, {0x25bc, 0x25bd, 0x1}, {0x25c0, 0x25c1, 0x1}, {0x25c6, 0x25c8, 0x1}, {0x25cb, 0x25ce, 0x3}, {0x25cf, 0x25d1, 0x1}, {0x25e2, 0x25e5, 0x1}, {0x25ef, 0x2605, 0x16}, {0x2606, 0x2609, 0x3}, {0x260e, 0x260f, 0x1}, {0x261c, 0x261e, 0x2}, {0x2640, 0x2642, 0x2}, {0x2660, 0x2661, 0x1}, {0x2663, 0x2665, 0x1}, {0x2667, 0x266a, 0x1}, {0x266c, 0x266d, 0x1}, {0x266f, 0x269e, 0x2f}, {0x269f, 0x26bf, 0x20}, {0x26c6, 0x26cd, 0x1}, {0x26cf, 0x26d3, 0x1}, {0x26d5, 0x26e1, 0x1}, {0x26e3, 0x26e8, 0x5}, {0x26e9, 0x26eb, 0x2}, {0x26ec, 0x26f1, 0x1}, {0x26f4, 0x26f6, 0x2}, {0x26f7, 0x26f9, 0x1}, {0x26fb, 0x26fc, 0x1}, {0x26fe, 0x26ff, 0x1}, {0x273d, 0x2776, 0x39}, {0x2777, 0x277f, 0x1}, {0x2b56, 0x2b59, 0x1}, {0x3248, 0x324f, 0x1}, {0xe000, 0xf8ff, 0x1}, {0xfe00, 0xfe0f, 0x1}, {0xfffd, 0xfffd, 0x1}}, []unicode.Range32{{0x1f100, 0x1f10a, 0x1}, {0x1f110, 0x1f12d, 0x1}, {0x1f130, 0x1f169, 0x1}, {0x1f170, 0x1f18d, 0x1}, {0x1f18f, 0x1f190, 0x1}, {0x1f19b, 0x1f1ac, 0x1}, {0xe0100, 0xe01ef, 0x1}, {0xf0000, 0xffffd, 0x1}, {0x100000, 0x10fffd, 0x1}}, 16}
	eastAsianWidthFullwidth = &unicode.RangeTable{[]unicode.Range16{{0x3000, 0xff01, 0xcf01}, {0xff02, 0xff60, 0x1}, {0xffe0, 0xffe6, 0x1}}, nil, 0}
	eastAsianWidthHalfwidth = &unicode.RangeTable{[]unicode.Range16{{0x20a9, 0xff61, 0xdeb8}, {0xff62, 0xffbe, 0x1}, {0xffc2, 0xffc7, 0x1}, {0xffca, 0xffcf, 0x1}, {0xffd2, 0xffd7, 0x1}, {0xffda, 0xffdc, 0x1}, {0xffe8, 0xffee, 0x1}}, nil, 0}
	eastAsianWidthNeutral   = &unicode.RangeTable{[]unicode.Range16{{0x0, 0x1f, 0x1}, {0x7f, 0xa0, 0x1}, {0xa9, 0xab, 0x2}, {0xb5, 0xbb, 0x6}, {0xc0, 0xc5, 0x1}, {0xc7, 0xcf, 0x1}, {0xd1, 0xd6, 0x1}, {0xd9, 0xdd, 0x1}, {0xe2, 0xe5, 0x1}, {0xe7, 0xeb, 0x4}, {0xee, 0xef, 0x1}, {0xf1, 0xf4, 0x3}, {0xf5, 0xf6, 0x1}, {0xfb, 0xff, 0x2}, {0x100, 0x102, 0x2}, {0x103, 0x110, 0x1}, {0x112, 0x114, 0x2}, {0x115, 0x11a, 0x1}, {0x11c, 0x125, 0x1}, {0x128, 0x12a, 0x1}, {0x12c, 0x130, 0x1}, {0x134, 0x137, 0x1}, {0x139, 0x13e, 0x1}, {0x143, 0x145, 0x2}, {0x146, 0x147, 0x1}, {0x14c, 0x14e, 0x2}, {0x14f, 0x151, 0x1}, {0x154, 0x165, 0x1}, {0x168, 0x16a, 0x1}, {0x16c, 0x1cd, 0x1}, {0x1cf, 0x1dd, 0x2}, {0x1de, 0x250, 0x1}, {0x252, 0x260, 0x1}, {0x262, 0x2c3, 0x1}, {0x2c5, 0x2c6, 0x1}, {0x2c8, 0x2cc, 0x4}, {0x2ce, 0x2cf, 0x1}, {0x2d1, 0x2d7, 0x1}, {0x2dc, 0x2e0, 0x2}, {0x2e1, 0x2ff, 0x1}, {0x370, 0x390, 0x1}, {0x3a2, 0x3aa, 0x8}, {0x3ab, 0x3b0, 0x1}, {0x3c2, 0x3ca, 0x8}, {0x3cb, 0x400, 0x1}, {0x402, 0x40f, 0x1}, {0x450, 0x452, 0x2}, {0x453, 0x10ff, 0x1}, {0x1160, 0x200f, 0x1}, {0x2011, 0x2012, 0x1}, {0x2017, 0x201a, 0x3}, {0x201b, 0x201e, 0x3}, {0x201f, 0x2023, 0x4}, {0x2028, 0x202f, 0x1}, {0x2031, 0x2034, 0x3}, {0x2036, 0x203a, 0x1}, {0x203c, 0x203d, 0x1}, {0x203f, 0x2073, 0x1}, {0x2075, 0x207e, 0x1}, {0x2080, 0x2085, 0x5}, {0x2086, 0x20a8, 0x1}, {0x20aa, 0x20ab, 0x1}, {0x20ad, 0x2102, 0x1}, {0x2104, 0x2106, 0x2}, {0x2107, 0x2108, 0x1}, {0x210a, 0x2112, 0x1}, {0x2114, 0x2115, 0x1}, {0x2117, 0x2120, 0x1}, {0x2123, 0x2125, 0x1}, {0x2127, 0x212a, 0x1}, {0x212c, 0x2152, 0x1}, {0x2155, 0x215a, 0x1}, {0x215f, 0x216c, 0xd}, {0x216d, 0x216f, 0x1}, {0x217a, 0x2188, 0x1}, {0x218a, 0x218f, 0x1}, {0x219a, 0x21b7, 0x1}, {0x21ba, 0x21d1, 0x1}, {0x21d3, 0x21d5, 0x2}, {0x21d6, 0x21e6, 0x1}, {0x21e8, 0x21ff, 0x1}, {0x2201, 0x2204, 0x3}, {0x2205, 0x2206, 0x1}, {0x2209, 0x220a, 0x1}, {0x220c, 0x220e, 0x1}, {0x2210, 0x2212, 0x2}, {0x2213, 0x2214, 0x1}, {0x2216, 0x2219, 0x1}, {0x221b, 0x221c, 0x1}, {0x2221, 0x2222, 0x1}, {0x2224, 0x2226, 0x2}, {0x222d, 0x222f, 0x2}, {0x2230, 0x2233, 0x1}, {0x2238, 0x223b, 0x1}, {0x223e, 0x2247, 0x1}, {0x2249, 0x224b, 0x1}, {0x224d, 0x2251, 0x1}, {0x2253, 0x225f, 0x1}, {0x2262, 0x2263, 0x1}, {0x2268, 0x2269, 0x1}, {0x226c, 0x226d, 0x1}, {0x2270, 0x2281, 0x1}, {0x2284, 0x2285, 0x1}, {0x2288, 0x2294, 0x1}, {0x2296, 0x2298, 0x1}, {0x229a, 0x22a4, 0x1}, {0x22a6, 0x22be, 0x1}, {0x22c0, 0x2311, 0x1}, {0x2313, 0x2319, 0x1}, {0x231c, 0x2328, 0x1}, {0x232b, 0x23e8, 0x1}, {0x23ed, 0x23ef, 0x1}, {0x23f1, 0x23f2, 0x1}, {0x23f4, 0x245f, 0x1}, {0x24ea, 0x254c, 0x62}, {0x254d, 0x254f, 0x1}, {0x2574, 0x257f, 0x1}, {0x2590, 0x2591, 0x1}, {0x2596, 0x259f, 0x1}, {0x25a2, 0x25aa, 0x8}, {0x25ab, 0x25b1, 0x1}, {0x25b4, 0x25b5, 0x1}, {0x25b8, 0x25bb, 0x1}, {0x25be, 0x25bf, 0x1}, {0x25c2, 0x25c5, 0x1}, {0x25c9, 0x25ca, 0x1}, {0x25cc, 0x25cd, 0x1}, {0x25d2, 0x25e1, 0x1}, {0x25e6, 0x25ee, 0x1}, {0x25f0, 0x25fc, 0x1}, {0x25ff, 0x2604, 0x1}, {0x2607, 0x2608, 0x1}, {0x260a, 0x260d, 0x1}, {0x2610, 0x2613, 0x1}, {0x2616, 0x261b, 0x1}, {0x261d, 0x261f, 0x2}, {0x2620, 0x263f, 0x1}, {0x2641, 0x2643, 0x2}, {0x2644, 0x2647, 0x1}, {0x2654, 0x265f, 0x1}, {0x2662, 0x2666, 0x4}, {0x266b, 0x266e, 0x3}, {0x2670, 0x267e, 0x1}, {0x2680, 0x2692, 0x1}, {0x2694, 0x269d, 0x1}, {0x26a0, 0x26a2, 0x2}, {0x26a3, 0x26a9, 0x1}, {0x26ac, 0x26bc, 0x1}, {0x26c0, 0x26c3, 0x1}, {0x26e2, 0x26e4, 0x2}, {0x26e5, 0x26e7, 0x1}, {0x2700, 0x2704, 0x1}, {0x2706, 0x2709, 0x1}, {0x270c, 0x2727, 0x1}, {0x2729, 0x273c, 0x1}, {0x273e, 0x274b, 0x1}, {0x274d, 0x274f, 0x2}, {0x2750, 0x2752, 0x1}, {0x2756, 0x2758, 0x2}, {0x2759, 0x2775, 0x1}, {0x2780, 0x2794, 0x1}, {0x2798, 0x27af, 0x1}, {0x27b1, 0x27be, 0x1}, {0x27c0, 0x27e5, 0x1}, {0x27ee, 0x2984, 0x1}, {0x2987, 0x2b1a, 0x1}, {0x2b1d, 0x2b4f, 0x1}, {0x2b51, 0x2b54, 0x1}, {0x2b5a, 0x2e7f, 0x1}, {0x2e9a, 0x2ef4, 0x5a}, {0x2ef5, 0x2eff, 0x1}, {0x2fd6, 0x2fef, 0x1}, {0x2ffc, 0x2fff, 0x1}, {0x303f, 0x3040, 0x1}, {0x3097, 0x3098, 0x1}, {0x3100, 0x3104, 0x1}, {0x312f, 0x3130, 0x1}, {0x318f, 0x31bb, 0x2c}, {0x31bc, 0x31bf, 0x1}, {0x31e4, 0x31ef, 0x1}, {0x321f, 0x32ff, 0xe0}, {0x4dc0, 0x4dff, 0x1}, {0xa48d, 0xa48f, 0x1}, {0xa4c7, 0xa95f, 0x1}, {0xa97d, 0xabff, 0x1}, {0xd7a4, 0xdfff, 0x1}, {0xfb00, 0xfdff, 0x1}, {0xfe1a, 0xfe2f, 0x1}, {0xfe53, 0xfe67, 0x14}, {0xfe6c, 0xff00, 0x1}, {0xffbf, 0xffc1, 0x1}, {0xffc8, 0xffc9, 0x1}, {0xffd0, 0xffd1, 0x1}, {0xffd8, 0xffd9, 0x1}, {0xffdd, 0xffdf, 0x1}, {0xffe7, 0xffef, 0x8}, {0xfff0, 0xfffc, 0x1}, {0xfffe, 0xffff, 0x1}}, []unicode.Range32{{0x10000, 0x16fdf, 0x1}, {0x16fe2, 0x16fff, 0x1}, {0x187ed, 0x187ff, 0x1}, {0x18af3, 0x1afff, 0x1}, {0x1b11f, 0x1b16f, 0x1}, {0x1b2fc, 0x1f003, 0x1}, {0x1f005, 0x1f0ce, 0x1}, {0x1f0d0, 0x1f0ff, 0x1}, {0x1f10b, 0x1f10f, 0x1}, {0x1f12e, 0x1f12f, 0x1}, {0x1f16a, 0x1f16f, 0x1}, {0x1f1ad, 0x1f1ff, 0x1}, {0x1f203, 0x1f20f, 0x1}, {0x1f23c, 0x1f23f, 0x1}, {0x1f249, 0x1f24f, 0x1}, {0x1f252, 0x1f25f, 0x1}, {0x1f266, 0x1f2ff, 0x1}, {0x1f321, 0x1f32c, 0x1}, {0x1f336, 0x1f37d, 0x47}, {0x1f394, 0x1f39f, 0x1}, {0x1f3cb, 0x1f3ce, 0x1}, {0x1f3d4, 0x1f3df, 0x1}, {0x1f3f1, 0x1f3f3, 0x1}, {0x1f3f5, 0x1f3f7, 0x1}, {0x1f43f, 0x1f441, 0x2}, {0x1f4fd, 0x1f4fe, 0x1}, {0x1f53e, 0x1f54a, 0x1}, {0x1f54f, 0x1f568, 0x19}, {0x1f569, 0x1f579, 0x1}, {0x1f57b, 0x1f594, 0x1}, {0x1f597, 0x1f5a3, 0x1}, {0x1f5a5, 0x1f5fa, 0x1}, {0x1f650, 0x1f67f, 0x1}, {0x1f6c6, 0x1f6cb, 0x1}, {0x1f6cd, 0x1f6cf, 0x1}, {0x1f6d3, 0x1f6ea, 0x1}, {0x1f6ed, 0x1f6f3, 0x1}, {0x1f6f9, 0x1f90f, 0x1}, {0x1f93f, 0x1f94d, 0xe}, {0x1f94e, 0x1f94f, 0x1}, {0x1f96c, 0x1f97f, 0x1}, {0x1f998, 0x1f9bf, 0x1}, {0x1f9c1, 0x1f9cf, 0x1}, {0x1f9e7, 0x1ffff, 0x1}, {0x2fffe, 0x2ffff, 0x1}, {0x3fffe, 0xe00ff, 0x1}, {0xe01f0, 0xeffff, 0x1}, {0xffffe, 0xfffff, 0x1}, {0x10fffe, 0x10ffff, 0x1}}, 14}
	eastAsianWidthNarrow    = &unicode.RangeTable{[]unicode.Range16{{0x20, 0x7e, 0x1}, {0xa2, 0xa3, 0x1}, {0xa5, 0xa6, 0x1}, {0xac, 0xaf, 0x3}, {0x27e6, 0x27ed, 0x1}, {0x2985, 0x2986, 0x1}}, nil, 4}
	eastAsianWidthWide      = &unicode.RangeTable{[]unicode.Range16{{0x1100, 0x115f, 0x1}, {0x231a, 0x231b, 0x1}, {0x2329, 0x232a, 0x1}, {0x23e9, 0x23ec, 0x1}, {0x23f0, 0x23f3, 0x3}, {0x25fd, 0x25fe, 0x1}, {0x2614, 0x2615, 0x1}, {0x2648, 0x2653, 0x1}, {0x267f, 0x2693, 0x14}, {0x26a1, 0x26aa, 0x9}, {0x26ab, 0x26bd, 0x12}, {0x26be, 0x26c4, 0x6}, {0x26c5, 0x26ce, 0x9}, {0x26d4, 0x26ea, 0x16}, {0x26f2, 0x26f3, 0x1}, {0x26f5, 0x26fa, 0x5}, {0x26fd, 0x2705, 0x8}, {0x270a, 0x270b, 0x1}, {0x2728, 0x274c, 0x24}, {0x274e, 0x2753, 0x5}, {0x2754, 0x2755, 0x1}, {0x2757, 0x2795, 0x3e}, {0x2796, 0x2797, 0x1}, {0x27b0, 0x27bf, 0xf}, {0x2b1b, 0x2b1c, 0x1}, {0x2b50, 0x2b55, 0x5}, {0x2e80, 0x2e99, 0x1}, {0x2e9b, 0x2ef3, 0x1}, {0x2f00, 0x2fd5, 0x1}, {0x2ff0, 0x2ffb, 0x1}, {0x3001, 0x303e, 0x1}, {0x3041, 0x3096, 0x1}, {0x3099, 0x30ff, 0x1}, {0x3105, 0x312e, 0x1}, {0x3131, 0x318e, 0x1}, {0x3190, 0x31ba, 0x1}, {0x31c0, 0x31e3, 0x1}, {0x31f0, 0x321e, 0x1}, {0x3220, 0x3247, 0x1}, {0x3250, 0x32fe, 0x1}, {0x3300, 0x4dbf, 0x1}, {0x4e00, 0xa48c, 0x1}, {0xa490, 0xa4c6, 0x1}, {0xa960, 0xa97c, 0x1}, {0xac00, 0xd7a3, 0x1}, {0xf900, 0xfaff, 0x1}, {0xfe10, 0xfe19, 0x1}, {0xfe30, 0xfe52, 0x1}, {0xfe54, 0xfe66, 0x1}, {0xfe68, 0xfe6b, 0x1}}, []unicode.Range32{{0x16fe0, 0x16fe1, 0x1}, {0x17000, 0x187ec, 0x1}, {0x18800, 0x18af2, 0x1}, {0x1b000, 0x1b11e, 0x1}, {0x1b170, 0x1b2fb, 0x1}, {0x1f004, 0x1f0cf, 0xcb}, {0x1f18e, 0x1f191, 0x3}, {0x1f192, 0x1f19a, 0x1}, {0x1f200, 0x1f202, 0x1}, {0x1f210, 0x1f23b, 0x1}, {0x1f240, 0x1f248, 0x1}, {0x1f250, 0x1f251, 0x1}, {0x1f260, 0x1f265, 0x1}, {0x1f300, 0x1f320, 0x1}, {0x1f32d, 0x1f335, 0x1}, {0x1f337, 0x1f37c, 0x1}, {0x1f37e, 0x1f393, 0x1}, {0x1f3a0, 0x1f3ca, 0x1}, {0x1f3cf, 0x1f3d3, 0x1}, {0x1f3e0, 0x1f3f0, 0x1}, {0x1f3f4, 0x1f3f8, 0x4}, {0x1f3f9, 0x1f43e, 0x1}, {0x1f440, 0x1f442, 0x2}, {0x1f443, 0x1f4fc, 0x1}, {0x1f4ff, 0x1f53d, 0x1}, {0x1f54b, 0x1f54e, 0x1}, {0x1f550, 0x1f567, 0x1}, {0x1f57a, 0x1f595, 0x1b}, {0x1f596, 0x1f5a4, 0xe}, {0x1f5fb, 0x1f64f, 0x1}, {0x1f680, 0x1f6c5, 0x1}, {0x1f6cc, 0x1f6d0, 0x4}, {0x1f6d1, 0x1f6d2, 0x1}, {0x1f6eb, 0x1f6ec, 0x1}, {0x1f6f4, 0x1f6f8, 0x1}, {0x1f910, 0x1f93e, 0x1}, {0x1f940, 0x1f94c, 0x1}, {0x1f950, 0x1f96b, 0x1}, {0x1f980, 0x1f997, 0x1}, {0x1f9c0, 0x1f9d0, 0x10}, {0x1f9d1, 0x1f9e6, 0x1}, {0x20000, 0x2fffd, 0x1}, {0x30000, 0x3fffd, 0x1}}, 0}
)
