package unicodeh

import "unicode"

// Unicode property "Diacritic" (known as "Dia", "Diacritic").
// Kind of property: "Binary".
// Based on file "PropList.txt".
var (
	DiacriticNo  = diacriticNo  // Value "No" (known as "N", "No", "F", "False").
	DiacriticYes = diacriticYes // Value "Yes" (known as "Y", "Yes", "T", "True").
)

var (
	diacriticNo  = &unicode.RangeTable{[]unicode.Range16{{0x0, 0x5d, 0x1}, {0x5f, 0x61, 0x2}, {0x62, 0xa7, 0x1}, {0xa9, 0xae, 0x1}, {0xb0, 0xb3, 0x1}, {0xb5, 0xb6, 0x1}, {0xb9, 0x2af, 0x1}, {0x34f, 0x358, 0x9}, {0x359, 0x35c, 0x1}, {0x363, 0x373, 0x1}, {0x376, 0x379, 0x1}, {0x37b, 0x383, 0x1}, {0x386, 0x482, 0x1}, {0x488, 0x558, 0x1}, {0x55a, 0x590, 0x1}, {0x5a2, 0x5be, 0x1c}, {0x5c0, 0x5c3, 0x3}, {0x5c5, 0x64a, 0x1}, {0x653, 0x656, 0x1}, {0x659, 0x6de, 0x1}, {0x6e1, 0x6e4, 0x1}, {0x6e7, 0x6e9, 0x1}, {0x6ed, 0x72f, 0x1}, {0x74b, 0x7a5, 0x1}, {0x7b1, 0x7ea, 0x1}, {0x7f6, 0x817, 0x1}, {0x81a, 0x8e2, 0x1}, {0x8ff, 0x93b, 0x1}, {0x93d, 0x94c, 0x1}, {0x94e, 0x950, 0x1}, {0x955, 0x970, 0x1}, {0x972, 0x9bb, 0x1}, {0x9bd, 0x9cc, 0x1}, {0x9ce, 0xa3b, 0x1}, {0xa3d, 0xa4c, 0x1}, {0xa4e, 0xabb, 0x1}, {0xabd, 0xacc, 0x1}, {0xace, 0xafc, 0x1}, {0xb00, 0xb3b, 0x1}, {0xb3d, 0xb4c, 0x1}, {0xb4e, 0xbcc, 0x1}, {0xbce, 0xc4c, 0x1}, {0xc4e, 0xcbb, 0x1}, {0xcbd, 0xccc, 0x1}, {0xcce, 0xd3a, 0x1}, {0xd3d, 0xd4c, 0x1}, {0xd4e, 0xdc9, 0x1}, {0xdcb, 0xe46, 0x1}, {0xe4d, 0xe4f, 0x2}, {0xe50, 0xec7, 0x1}, {0xecd, 0xf17, 0x1}, {0xf1a, 0xf34, 0x1}, {0xf36, 0xf3a, 0x2}, {0xf3b, 0xf3d, 0x1}, {0xf40, 0xf81, 0x1}, {0xf85, 0xf88, 0x3}, {0xf89, 0xfc5, 0x1}, {0xfc7, 0x1036, 0x1}, {0x1038, 0x103b, 0x3}, {0x103c, 0x1086, 0x1}, {0x108e, 0x1090, 0x2}, {0x1091, 0x1099, 0x1}, {0x109c, 0x17c8, 0x1}, {0x17d4, 0x17dc, 0x1}, {0x17de, 0x1938, 0x1}, {0x193c, 0x1a74, 0x1}, {0x1a7d, 0x1a7e, 0x1}, {0x1a80, 0x1aaf, 0x1}, {0x1abe, 0x1b33, 0x1}, {0x1b35, 0x1b43, 0x1}, {0x1b45, 0x1b6a, 0x1}, {0x1b74, 0x1ba9, 0x1}, {0x1bac, 0x1c35, 0x1}, {0x1c38, 0x1c77, 0x1}, {0x1c7e, 0x1ccf, 0x1}, {0x1ce9, 0x1cec, 0x1}, {0x1cee, 0x1cf3, 0x1}, {0x1cf5, 0x1cf6, 0x1}, {0x1cfa, 0x1d2b, 0x1}, {0x1d6b, 0x1dc3, 0x1}, {0x1dd0, 0x1df4, 0x1}, {0x1dfa, 0x1dfc, 0x1}, {0x1e00, 0x1fbc, 0x1}, {0x1fbe, 0x1fc2, 0x4}, {0x1fc3, 0x1fcc, 0x1}, {0x1fd0, 0x1fdc, 0x1}, {0x1fe0, 0x1fec, 0x1}, {0x1ff0, 0x1ffc, 0x1}, {0x1fff, 0x2cee, 0x1}, {0x2cf2, 0x2e2e, 0x1}, {0x2e30, 0x3029, 0x1}, {0x3030, 0x3098, 0x1}, {0x309d, 0x30fb, 0x1}, {0x30fd, 0xa66e, 0x1}, {0xa670, 0xa67b, 0x1}, {0xa67e, 0xa680, 0x2}, {0xa681, 0xa69b, 0x1}, {0xa69e, 0xa6ef, 0x1}, {0xa6f2, 0xa716, 0x1}, {0xa722, 0xa787, 0x1}, {0xa789, 0xa7f7, 0x1}, {0xa7fa, 0xa8c3, 0x1}, {0xa8c5, 0xa8df, 0x1}, {0xa8f2, 0xa92a, 0x1}, {0xa92f, 0xa952, 0x1}, {0xa954, 0xa9b2, 0x1}, {0xa9b4, 0xa9bf, 0x1}, {0xa9c1, 0xa9e4, 0x1}, {0xa9e6, 0xaa7a, 0x1}, {0xaa7e, 0xaabe, 0x1}, {0xaac3, 0xaaf5, 0x1}, {0xaaf7, 0xab5a, 0x1}, {0xab60, 0xabeb, 0x1}, {0xabee, 0xfb1d, 0x1}, {0xfb1f, 0xfe1f, 0x1}, {0xfe30, 0xff3d, 0x1}, {0xff3f, 0xff41, 0x2}, {0xff42, 0xff6f, 0x1}, {0xff71, 0xff9d, 0x1}, {0xffa0, 0xffe2, 0x1}, {0xffe4, 0xffff, 0x1}}, []unicode.Range32{{0x10000, 0x102df, 0x1}, {0x102e1, 0x10ae4, 0x1}, {0x10ae7, 0x110b8, 0x1}, {0x110bb, 0x11132, 0x1}, {0x11135, 0x11172, 0x1}, {0x11174, 0x111bf, 0x1}, {0x111c1, 0x111c9, 0x1}, {0x111cd, 0x11234, 0x1}, {0x11237, 0x112e8, 0x1}, {0x112eb, 0x1133b, 0x1}, {0x1133d, 0x1134c, 0x1}, {0x1134e, 0x11365, 0x1}, {0x1136d, 0x1136f, 0x1}, {0x11375, 0x11441, 0x1}, {0x11443, 0x11445, 0x1}, {0x11447, 0x114c1, 0x1}, {0x114c4, 0x115be, 0x1}, {0x115c1, 0x1163e, 0x1}, {0x11640, 0x116b5, 0x1}, {0x116b8, 0x1172a, 0x1}, {0x1172c, 0x11a33, 0x1}, {0x11a35, 0x11a46, 0x1}, {0x11a48, 0x11a98, 0x1}, {0x11a9a, 0x11c3e, 0x1}, {0x11c40, 0x11d41, 0x1}, {0x11d43, 0x11d46, 0x3}, {0x11d47, 0x16aef, 0x1}, {0x16af5, 0x16f8e, 0x1}, {0x16fa0, 0x1d166, 0x1}, {0x1d16a, 0x1d16c, 0x1}, {0x1d173, 0x1d17a, 0x1}, {0x1d183, 0x1d184, 0x1}, {0x1d18c, 0x1d1a9, 0x1}, {0x1d1ae, 0x1e8cf, 0x1}, {0x1e8d7, 0x1e943, 0x1}, {0x1e947, 0x1e94b, 0x4}, {0x1e94c, 0x10ffff, 0x1}}, 6}
	diacriticYes = &unicode.RangeTable{[]unicode.Range16{{0x5e, 0x60, 0x2}, {0xa8, 0xaf, 0x7}, {0xb4, 0xb7, 0x3}, {0xb8, 0x2b0, 0x1f8}, {0x2b1, 0x34e, 0x1}, {0x350, 0x357, 0x1}, {0x35d, 0x362, 0x1}, {0x374, 0x375, 0x1}, {0x37a, 0x384, 0xa}, {0x385, 0x483, 0xfe}, {0x484, 0x487, 0x1}, {0x559, 0x591, 0x38}, {0x592, 0x5a1, 0x1}, {0x5a3, 0x5bd, 0x1}, {0x5bf, 0x5c1, 0x2}, {0x5c2, 0x5c4, 0x2}, {0x64b, 0x652, 0x1}, {0x657, 0x658, 0x1}, {0x6df, 0x6e0, 0x1}, {0x6e5, 0x6e6, 0x1}, {0x6ea, 0x6ec, 0x1}, {0x730, 0x74a, 0x1}, {0x7a6, 0x7b0, 0x1}, {0x7eb, 0x7f5, 0x1}, {0x818, 0x819, 0x1}, {0x8e3, 0x8fe, 0x1}, {0x93c, 0x94d, 0x11}, {0x951, 0x954, 0x1}, {0x971, 0x9bc, 0x4b}, {0x9cd, 0xa3c, 0x6f}, {0xa4d, 0xabc, 0x6f}, {0xacd, 0xafd, 0x30}, {0xafe, 0xaff, 0x1}, {0xb3c, 0xb4d, 0x11}, {0xbcd, 0xc4d, 0x80}, {0xcbc, 0xccd, 0x11}, {0xd3b, 0xd3c, 0x1}, {0xd4d, 0xe47, 0x7d}, {0xe48, 0xe4c, 0x1}, {0xe4e, 0xec8, 0x7a}, {0xec9, 0xecc, 0x1}, {0xf18, 0xf19, 0x1}, {0xf35, 0xf39, 0x2}, {0xf3e, 0xf3f, 0x1}, {0xf82, 0xf84, 0x1}, {0xf86, 0xf87, 0x1}, {0xfc6, 0x1037, 0x71}, {0x1039, 0x103a, 0x1}, {0x1087, 0x108d, 0x1}, {0x108f, 0x109a, 0xb}, {0x109b, 0x17c9, 0x72e}, {0x17ca, 0x17d3, 0x1}, {0x17dd, 0x1939, 0x15c}, {0x193a, 0x193b, 0x1}, {0x1a75, 0x1a7c, 0x1}, {0x1a7f, 0x1ab0, 0x31}, {0x1ab1, 0x1abd, 0x1}, {0x1b34, 0x1b44, 0x10}, {0x1b6b, 0x1b73, 0x1}, {0x1baa, 0x1bab, 0x1}, {0x1c36, 0x1c37, 0x1}, {0x1c78, 0x1c7d, 0x1}, {0x1cd0, 0x1ce8, 0x1}, {0x1ced, 0x1cf4, 0x7}, {0x1cf7, 0x1cf9, 0x1}, {0x1d2c, 0x1d6a, 0x1}, {0x1dc4, 0x1dcf, 0x1}, {0x1df5, 0x1df9, 0x1}, {0x1dfd, 0x1dff, 0x1}, {0x1fbd, 0x1fbf, 0x2}, {0x1fc0, 0x1fc1, 0x1}, {0x1fcd, 0x1fcf, 0x1}, {0x1fdd, 0x1fdf, 0x1}, {0x1fed, 0x1fef, 0x1}, {0x1ffd, 0x1ffe, 0x1}, {0x2cef, 0x2cf1, 0x1}, {0x2e2f, 0x302a, 0x1fb}, {0x302b, 0x302f, 0x1}, {0x3099, 0x309c, 0x1}, {0x30fc, 0xa66f, 0x7573}, {0xa67c, 0xa67d, 0x1}, {0xa67f, 0xa69c, 0x1d}, {0xa69d, 0xa6f0, 0x53}, {0xa6f1, 0xa717, 0x26}, {0xa718, 0xa721, 0x1}, {0xa788, 0xa7f8, 0x70}, {0xa7f9, 0xa8c4, 0xcb}, {0xa8e0, 0xa8f1, 0x1}, {0xa92b, 0xa92e, 0x1}, {0xa953, 0xa9b3, 0x60}, {0xa9c0, 0xa9e5, 0x25}, {0xaa7b, 0xaa7d, 0x1}, {0xaabf, 0xaac2, 0x1}, {0xaaf6, 0xab5b, 0x65}, {0xab5c, 0xab5f, 0x1}, {0xabec, 0xabed, 0x1}, {0xfb1e, 0xfe20, 0x302}, {0xfe21, 0xfe2f, 0x1}, {0xff3e, 0xff40, 0x2}, {0xff70, 0xff9e, 0x2e}, {0xff9f, 0xffe3, 0x44}}, []unicode.Range32{{0x102e0, 0x10ae5, 0x805}, {0x10ae6, 0x110b9, 0x5d3}, {0x110ba, 0x11133, 0x79}, {0x11134, 0x11173, 0x3f}, {0x111c0, 0x111ca, 0xa}, {0x111cb, 0x111cc, 0x1}, {0x11235, 0x11236, 0x1}, {0x112e9, 0x112ea, 0x1}, {0x1133c, 0x1134d, 0x11}, {0x11366, 0x1136c, 0x1}, {0x11370, 0x11374, 0x1}, {0x11442, 0x11446, 0x4}, {0x114c2, 0x114c3, 0x1}, {0x115bf, 0x115c0, 0x1}, {0x1163f, 0x116b6, 0x77}, {0x116b7, 0x1172b, 0x74}, {0x11a34, 0x11a47, 0x13}, {0x11a99, 0x11c3f, 0x1a6}, {0x11d42, 0x11d44, 0x2}, {0x11d45, 0x16af0, 0x4dab}, {0x16af1, 0x16af4, 0x1}, {0x16f8f, 0x16f9f, 0x1}, {0x1d167, 0x1d169, 0x1}, {0x1d16d, 0x1d172, 0x1}, {0x1d17b, 0x1d182, 0x1}, {0x1d185, 0x1d18b, 0x1}, {0x1d1aa, 0x1d1ad, 0x1}, {0x1e8d0, 0x1e8d6, 0x1}, {0x1e944, 0x1e946, 0x1}, {0x1e948, 0x1e94a, 0x1}}, 3}
)