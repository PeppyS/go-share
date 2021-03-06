package unicodeh

import "unicode"

// Unicode property "Joining_Group" (known as "jg", "Joining_Group").
// Kind of property: "Enumerated".
// Based on file "extracted\DerivedJoiningGroup.txt".
var (
	JoiningGroupAfricanFeh          = joiningGroupAfricanFeh          // Value "African_Feh".
	JoiningGroupAfricanNoon         = joiningGroupAfricanNoon         // Value "African_Noon".
	JoiningGroupAfricanQaf          = joiningGroupAfricanQaf          // Value "African_Qaf".
	JoiningGroupAin                 = joiningGroupAin                 // Value "Ain".
	JoiningGroupAlaph               = joiningGroupAlaph               // Value "Alaph".
	JoiningGroupAlef                = joiningGroupAlef                // Value "Alef".
	JoiningGroupBeh                 = joiningGroupBeh                 // Value "Beh".
	JoiningGroupBeth                = joiningGroupBeth                // Value "Beth".
	JoiningGroupBurushaskiYehBarree = joiningGroupBurushaskiYehBarree // Value "Burushaski_Yeh_Barree".
	JoiningGroupDal                 = joiningGroupDal                 // Value "Dal".
	JoiningGroupDalathRish          = joiningGroupDalathRish          // Value "Dalath_Rish".
	JoiningGroupE                   = joiningGroupE                   // Value "E".
	JoiningGroupFarsiYeh            = joiningGroupFarsiYeh            // Value "Farsi_Yeh".
	JoiningGroupFe                  = joiningGroupFe                  // Value "Fe".
	JoiningGroupFeh                 = joiningGroupFeh                 // Value "Feh".
	JoiningGroupFinalSemkath        = joiningGroupFinalSemkath        // Value "Final_Semkath".
	JoiningGroupGaf                 = joiningGroupGaf                 // Value "Gaf".
	JoiningGroupGamal               = joiningGroupGamal               // Value "Gamal".
	JoiningGroupHah                 = joiningGroupHah                 // Value "Hah".
	JoiningGroupHe                  = joiningGroupHe                  // Value "He".
	JoiningGroupHeh                 = joiningGroupHeh                 // Value "Heh".
	JoiningGroupHehGoal             = joiningGroupHehGoal             // Value "Heh_Goal".
	JoiningGroupHeth                = joiningGroupHeth                // Value "Heth".
	JoiningGroupKaf                 = joiningGroupKaf                 // Value "Kaf".
	JoiningGroupKaph                = joiningGroupKaph                // Value "Kaph".
	JoiningGroupKhaph               = joiningGroupKhaph               // Value "Khaph".
	JoiningGroupKnottedHeh          = joiningGroupKnottedHeh          // Value "Knotted_Heh".
	JoiningGroupLam                 = joiningGroupLam                 // Value "Lam".
	JoiningGroupLamadh              = joiningGroupLamadh              // Value "Lamadh".
	JoiningGroupMalayalamBha        = joiningGroupMalayalamBha        // Value "Malayalam_Bha".
	JoiningGroupMalayalamJa         = joiningGroupMalayalamJa         // Value "Malayalam_Ja".
	JoiningGroupMalayalamLla        = joiningGroupMalayalamLla        // Value "Malayalam_Lla".
	JoiningGroupMalayalamLlla       = joiningGroupMalayalamLlla       // Value "Malayalam_Llla".
	JoiningGroupMalayalamNga        = joiningGroupMalayalamNga        // Value "Malayalam_Nga".
	JoiningGroupMalayalamNna        = joiningGroupMalayalamNna        // Value "Malayalam_Nna".
	JoiningGroupMalayalamNnna       = joiningGroupMalayalamNnna       // Value "Malayalam_Nnna".
	JoiningGroupMalayalamNya        = joiningGroupMalayalamNya        // Value "Malayalam_Nya".
	JoiningGroupMalayalamRa         = joiningGroupMalayalamRa         // Value "Malayalam_Ra".
	JoiningGroupMalayalamSsa        = joiningGroupMalayalamSsa        // Value "Malayalam_Ssa".
	JoiningGroupMalayalamTta        = joiningGroupMalayalamTta        // Value "Malayalam_Tta".
	JoiningGroupManichaeanAleph     = joiningGroupManichaeanAleph     // Value "Manichaean_Aleph".
	JoiningGroupManichaeanAyin      = joiningGroupManichaeanAyin      // Value "Manichaean_Ayin".
	JoiningGroupManichaeanBeth      = joiningGroupManichaeanBeth      // Value "Manichaean_Beth".
	JoiningGroupManichaeanDaleth    = joiningGroupManichaeanDaleth    // Value "Manichaean_Daleth".
	JoiningGroupManichaeanDhamedh   = joiningGroupManichaeanDhamedh   // Value "Manichaean_Dhamedh".
	JoiningGroupManichaeanFive      = joiningGroupManichaeanFive      // Value "Manichaean_Five".
	JoiningGroupManichaeanGimel     = joiningGroupManichaeanGimel     // Value "Manichaean_Gimel".
	JoiningGroupManichaeanHeth      = joiningGroupManichaeanHeth      // Value "Manichaean_Heth".
	JoiningGroupManichaeanHundred   = joiningGroupManichaeanHundred   // Value "Manichaean_Hundred".
	JoiningGroupManichaeanKaph      = joiningGroupManichaeanKaph      // Value "Manichaean_Kaph".
	JoiningGroupManichaeanLamedh    = joiningGroupManichaeanLamedh    // Value "Manichaean_Lamedh".
	JoiningGroupManichaeanMem       = joiningGroupManichaeanMem       // Value "Manichaean_Mem".
	JoiningGroupManichaeanNun       = joiningGroupManichaeanNun       // Value "Manichaean_Nun".
	JoiningGroupManichaeanOne       = joiningGroupManichaeanOne       // Value "Manichaean_One".
	JoiningGroupManichaeanPe        = joiningGroupManichaeanPe        // Value "Manichaean_Pe".
	JoiningGroupManichaeanQoph      = joiningGroupManichaeanQoph      // Value "Manichaean_Qoph".
	JoiningGroupManichaeanResh      = joiningGroupManichaeanResh      // Value "Manichaean_Resh".
	JoiningGroupManichaeanSadhe     = joiningGroupManichaeanSadhe     // Value "Manichaean_Sadhe".
	JoiningGroupManichaeanSamekh    = joiningGroupManichaeanSamekh    // Value "Manichaean_Samekh".
	JoiningGroupManichaeanTaw       = joiningGroupManichaeanTaw       // Value "Manichaean_Taw".
	JoiningGroupManichaeanTen       = joiningGroupManichaeanTen       // Value "Manichaean_Ten".
	JoiningGroupManichaeanTeth      = joiningGroupManichaeanTeth      // Value "Manichaean_Teth".
	JoiningGroupManichaeanThamedh   = joiningGroupManichaeanThamedh   // Value "Manichaean_Thamedh".
	JoiningGroupManichaeanTwenty    = joiningGroupManichaeanTwenty    // Value "Manichaean_Twenty".
	JoiningGroupManichaeanWaw       = joiningGroupManichaeanWaw       // Value "Manichaean_Waw".
	JoiningGroupManichaeanYodh      = joiningGroupManichaeanYodh      // Value "Manichaean_Yodh".
	JoiningGroupManichaeanZayin     = joiningGroupManichaeanZayin     // Value "Manichaean_Zayin".
	JoiningGroupMeem                = joiningGroupMeem                // Value "Meem".
	JoiningGroupMim                 = joiningGroupMim                 // Value "Mim".
	JoiningGroupNoJoiningGroup      = joiningGroupNoJoiningGroup      // Value "No_Joining_Group".
	JoiningGroupNoon                = joiningGroupNoon                // Value "Noon".
	JoiningGroupNun                 = joiningGroupNun                 // Value "Nun".
	JoiningGroupNya                 = joiningGroupNya                 // Value "Nya".
	JoiningGroupPe                  = joiningGroupPe                  // Value "Pe".
	JoiningGroupQaf                 = joiningGroupQaf                 // Value "Qaf".
	JoiningGroupQaph                = joiningGroupQaph                // Value "Qaph".
	JoiningGroupReh                 = joiningGroupReh                 // Value "Reh".
	JoiningGroupReversedPe          = joiningGroupReversedPe          // Value "Reversed_Pe".
	JoiningGroupRohingyaYeh         = joiningGroupRohingyaYeh         // Value "Rohingya_Yeh".
	JoiningGroupSad                 = joiningGroupSad                 // Value "Sad".
	JoiningGroupSadhe               = joiningGroupSadhe               // Value "Sadhe".
	JoiningGroupSeen                = joiningGroupSeen                // Value "Seen".
	JoiningGroupSemkath             = joiningGroupSemkath             // Value "Semkath".
	JoiningGroupShin                = joiningGroupShin                // Value "Shin".
	JoiningGroupStraightWaw         = joiningGroupStraightWaw         // Value "Straight_Waw".
	JoiningGroupSwashKaf            = joiningGroupSwashKaf            // Value "Swash_Kaf".
	JoiningGroupSyriacWaw           = joiningGroupSyriacWaw           // Value "Syriac_Waw".
	JoiningGroupTah                 = joiningGroupTah                 // Value "Tah".
	JoiningGroupTaw                 = joiningGroupTaw                 // Value "Taw".
	JoiningGroupTehMarbuta          = joiningGroupTehMarbuta          // Value "Teh_Marbuta".
	JoiningGroupHamzaOnHehGoal      = joiningGroupHamzaOnHehGoal      // Value "Hamza_On_Heh_Goal" (known as "Teh_Marbuta_Goal", "Hamza_On_Heh_Goal").
	JoiningGroupTeth                = joiningGroupTeth                // Value "Teth".
	JoiningGroupWaw                 = joiningGroupWaw                 // Value "Waw".
	JoiningGroupYeh                 = joiningGroupYeh                 // Value "Yeh".
	JoiningGroupYehBarree           = joiningGroupYehBarree           // Value "Yeh_Barree".
	JoiningGroupYehWithTail         = joiningGroupYehWithTail         // Value "Yeh_With_Tail".
	JoiningGroupYudh                = joiningGroupYudh                // Value "Yudh".
	JoiningGroupYudhHe              = joiningGroupYudhHe              // Value "Yudh_He".
	JoiningGroupZain                = joiningGroupZain                // Value "Zain".
	JoiningGroupZhain               = joiningGroupZhain               // Value "Zhain".
)

var (
	joiningGroupAfricanFeh          = &unicode.RangeTable{[]unicode.Range16{{0x8bb, 0x8bb, 0x1}}, nil, 0}
	joiningGroupAfricanNoon         = &unicode.RangeTable{[]unicode.Range16{{0x8bd, 0x8bd, 0x1}}, nil, 0}
	joiningGroupAfricanQaf          = &unicode.RangeTable{[]unicode.Range16{{0x8bc, 0x8bc, 0x1}}, nil, 0}
	joiningGroupAin                 = &unicode.RangeTable{[]unicode.Range16{{0x639, 0x63a, 0x1}, {0x6a0, 0x6fc, 0x5c}, {0x75d, 0x75f, 0x1}, {0x8b3, 0x8b3, 0x1}}, nil, 0}
	joiningGroupAlaph               = &unicode.RangeTable{[]unicode.Range16{{0x710, 0x710, 0x1}}, nil, 0}
	joiningGroupAlef                = &unicode.RangeTable{[]unicode.Range16{{0x622, 0x623, 0x1}, {0x625, 0x627, 0x2}, {0x671, 0x673, 0x1}, {0x675, 0x773, 0xfe}, {0x774, 0x774, 0x1}}, nil, 0}
	joiningGroupBeh                 = &unicode.RangeTable{[]unicode.Range16{{0x628, 0x62a, 0x2}, {0x62b, 0x66e, 0x43}, {0x679, 0x680, 0x1}, {0x750, 0x756, 0x1}, {0x8a0, 0x8a1, 0x1}, {0x8b6, 0x8b8, 0x1}}, nil, 0}
	joiningGroupBeth                = &unicode.RangeTable{[]unicode.Range16{{0x712, 0x72d, 0x1b}}, nil, 0}
	joiningGroupBurushaskiYehBarree = &unicode.RangeTable{[]unicode.Range16{{0x77a, 0x77b, 0x1}}, nil, 0}
	joiningGroupDal                 = &unicode.RangeTable{[]unicode.Range16{{0x62f, 0x630, 0x1}, {0x688, 0x690, 0x1}, {0x6ee, 0x759, 0x6b}, {0x75a, 0x8ae, 0x154}}, nil, 0}
	joiningGroupDalathRish          = &unicode.RangeTable{[]unicode.Range16{{0x715, 0x716, 0x1}, {0x72a, 0x72f, 0x5}}, nil, 0}
	joiningGroupE                   = &unicode.RangeTable{[]unicode.Range16{{0x725, 0x725, 0x1}}, nil, 0}
	joiningGroupFarsiYeh            = &unicode.RangeTable{[]unicode.Range16{{0x63d, 0x63f, 0x1}, {0x6cc, 0x6ce, 0x2}, {0x775, 0x776, 0x1}}, nil, 0}
	joiningGroupFe                  = &unicode.RangeTable{[]unicode.Range16{{0x74f, 0x74f, 0x1}}, nil, 0}
	joiningGroupFeh                 = &unicode.RangeTable{[]unicode.Range16{{0x641, 0x6a1, 0x60}, {0x6a2, 0x6a6, 0x1}, {0x760, 0x761, 0x1}, {0x8a4, 0x8a4, 0x1}}, nil, 0}
	joiningGroupFinalSemkath        = &unicode.RangeTable{[]unicode.Range16{{0x724, 0x724, 0x1}}, nil, 0}
	joiningGroupGaf                 = &unicode.RangeTable{[]unicode.Range16{{0x63b, 0x63c, 0x1}, {0x6a9, 0x6ab, 0x2}, {0x6af, 0x6b4, 0x1}, {0x762, 0x764, 0x1}, {0x8b0, 0x8b0, 0x1}}, nil, 0}
	joiningGroupGamal               = &unicode.RangeTable{[]unicode.Range16{{0x713, 0x714, 0x1}, {0x72e, 0x72e, 0x1}}, nil, 0}
	joiningGroupHah                 = &unicode.RangeTable{[]unicode.Range16{{0x62c, 0x62e, 0x1}, {0x681, 0x687, 0x1}, {0x6bf, 0x757, 0x98}, {0x758, 0x76e, 0x16}, {0x76f, 0x772, 0x3}, {0x77c, 0x8a2, 0x126}}, nil, 0}
	joiningGroupHe                  = &unicode.RangeTable{[]unicode.Range16{{0x717, 0x717, 0x1}}, nil, 0}
	joiningGroupHeh                 = &unicode.RangeTable{[]unicode.Range16{{0x647, 0x647, 0x1}}, nil, 0}
	joiningGroupHehGoal             = &unicode.RangeTable{[]unicode.Range16{{0x6c1, 0x6c2, 0x1}}, nil, 0}
	joiningGroupHeth                = &unicode.RangeTable{[]unicode.Range16{{0x71a, 0x71a, 0x1}}, nil, 0}
	joiningGroupKaf                 = &unicode.RangeTable{[]unicode.Range16{{0x643, 0x6ac, 0x69}, {0x6ad, 0x6ae, 0x1}, {0x77f, 0x8b4, 0x135}}, nil, 0}
	joiningGroupKaph                = &unicode.RangeTable{[]unicode.Range16{{0x71f, 0x71f, 0x1}}, nil, 0}
	joiningGroupKhaph               = &unicode.RangeTable{[]unicode.Range16{{0x74e, 0x74e, 0x1}}, nil, 0}
	joiningGroupKnottedHeh          = &unicode.RangeTable{[]unicode.Range16{{0x6be, 0x6ff, 0x41}}, nil, 0}
	joiningGroupLam                 = &unicode.RangeTable{[]unicode.Range16{{0x644, 0x6b5, 0x71}, {0x6b6, 0x6b8, 0x1}, {0x76a, 0x8a6, 0x13c}}, nil, 0}
	joiningGroupLamadh              = &unicode.RangeTable{[]unicode.Range16{{0x720, 0x720, 0x1}}, nil, 0}
	joiningGroupMalayalamBha        = &unicode.RangeTable{[]unicode.Range16{{0x866, 0x866, 0x1}}, nil, 0}
	joiningGroupMalayalamJa         = &unicode.RangeTable{[]unicode.Range16{{0x861, 0x861, 0x1}}, nil, 0}
	joiningGroupMalayalamLla        = &unicode.RangeTable{[]unicode.Range16{{0x868, 0x868, 0x1}}, nil, 0}
	joiningGroupMalayalamLlla       = &unicode.RangeTable{[]unicode.Range16{{0x869, 0x869, 0x1}}, nil, 0}
	joiningGroupMalayalamNga        = &unicode.RangeTable{[]unicode.Range16{{0x860, 0x860, 0x1}}, nil, 0}
	joiningGroupMalayalamNna        = &unicode.RangeTable{[]unicode.Range16{{0x864, 0x864, 0x1}}, nil, 0}
	joiningGroupMalayalamNnna       = &unicode.RangeTable{[]unicode.Range16{{0x865, 0x865, 0x1}}, nil, 0}
	joiningGroupMalayalamNya        = &unicode.RangeTable{[]unicode.Range16{{0x862, 0x862, 0x1}}, nil, 0}
	joiningGroupMalayalamRa         = &unicode.RangeTable{[]unicode.Range16{{0x867, 0x867, 0x1}}, nil, 0}
	joiningGroupMalayalamSsa        = &unicode.RangeTable{[]unicode.Range16{{0x86a, 0x86a, 0x1}}, nil, 0}
	joiningGroupMalayalamTta        = &unicode.RangeTable{[]unicode.Range16{{0x863, 0x863, 0x1}}, nil, 0}
	joiningGroupManichaeanAleph     = &unicode.RangeTable{nil, []unicode.Range32{{0x10ac0, 0x10ac0, 0x1}}, 0}
	joiningGroupManichaeanAyin      = &unicode.RangeTable{nil, []unicode.Range32{{0x10ad9, 0x10ada, 0x1}}, 0}
	joiningGroupManichaeanBeth      = &unicode.RangeTable{nil, []unicode.Range32{{0x10ac1, 0x10ac2, 0x1}}, 0}
	joiningGroupManichaeanDaleth    = &unicode.RangeTable{nil, []unicode.Range32{{0x10ac5, 0x10ac5, 0x1}}, 0}
	joiningGroupManichaeanDhamedh   = &unicode.RangeTable{nil, []unicode.Range32{{0x10ad4, 0x10ad4, 0x1}}, 0}
	joiningGroupManichaeanFive      = &unicode.RangeTable{nil, []unicode.Range32{{0x10aec, 0x10aec, 0x1}}, 0}
	joiningGroupManichaeanGimel     = &unicode.RangeTable{nil, []unicode.Range32{{0x10ac3, 0x10ac4, 0x1}}, 0}
	joiningGroupManichaeanHeth      = &unicode.RangeTable{nil, []unicode.Range32{{0x10acd, 0x10acd, 0x1}}, 0}
	joiningGroupManichaeanHundred   = &unicode.RangeTable{nil, []unicode.Range32{{0x10aef, 0x10aef, 0x1}}, 0}
	joiningGroupManichaeanKaph      = &unicode.RangeTable{nil, []unicode.Range32{{0x10ad0, 0x10ad2, 0x1}}, 0}
	joiningGroupManichaeanLamedh    = &unicode.RangeTable{nil, []unicode.Range32{{0x10ad3, 0x10ad3, 0x1}}, 0}
	joiningGroupManichaeanMem       = &unicode.RangeTable{nil, []unicode.Range32{{0x10ad6, 0x10ad6, 0x1}}, 0}
	joiningGroupManichaeanNun       = &unicode.RangeTable{nil, []unicode.Range32{{0x10ad7, 0x10ad7, 0x1}}, 0}
	joiningGroupManichaeanOne       = &unicode.RangeTable{nil, []unicode.Range32{{0x10aeb, 0x10aeb, 0x1}}, 0}
	joiningGroupManichaeanPe        = &unicode.RangeTable{nil, []unicode.Range32{{0x10adb, 0x10adc, 0x1}}, 0}
	joiningGroupManichaeanQoph      = &unicode.RangeTable{nil, []unicode.Range32{{0x10ade, 0x10ae0, 0x1}}, 0}
	joiningGroupManichaeanResh      = &unicode.RangeTable{nil, []unicode.Range32{{0x10ae1, 0x10ae1, 0x1}}, 0}
	joiningGroupManichaeanSadhe     = &unicode.RangeTable{nil, []unicode.Range32{{0x10add, 0x10add, 0x1}}, 0}
	joiningGroupManichaeanSamekh    = &unicode.RangeTable{nil, []unicode.Range32{{0x10ad8, 0x10ad8, 0x1}}, 0}
	joiningGroupManichaeanTaw       = &unicode.RangeTable{nil, []unicode.Range32{{0x10ae4, 0x10ae4, 0x1}}, 0}
	joiningGroupManichaeanTen       = &unicode.RangeTable{nil, []unicode.Range32{{0x10aed, 0x10aed, 0x1}}, 0}
	joiningGroupManichaeanTeth      = &unicode.RangeTable{nil, []unicode.Range32{{0x10ace, 0x10ace, 0x1}}, 0}
	joiningGroupManichaeanThamedh   = &unicode.RangeTable{nil, []unicode.Range32{{0x10ad5, 0x10ad5, 0x1}}, 0}
	joiningGroupManichaeanTwenty    = &unicode.RangeTable{nil, []unicode.Range32{{0x10aee, 0x10aee, 0x1}}, 0}
	joiningGroupManichaeanWaw       = &unicode.RangeTable{nil, []unicode.Range32{{0x10ac7, 0x10ac7, 0x1}}, 0}
	joiningGroupManichaeanYodh      = &unicode.RangeTable{nil, []unicode.Range32{{0x10acf, 0x10acf, 0x1}}, 0}
	joiningGroupManichaeanZayin     = &unicode.RangeTable{nil, []unicode.Range32{{0x10ac9, 0x10aca, 0x1}}, 0}
	joiningGroupMeem                = &unicode.RangeTable{[]unicode.Range16{{0x645, 0x765, 0x120}, {0x766, 0x8a7, 0x141}}, nil, 0}
	joiningGroupMim                 = &unicode.RangeTable{[]unicode.Range16{{0x721, 0x721, 0x1}}, nil, 0}
	joiningGroupNoJoiningGroup      = &unicode.RangeTable{[]unicode.Range16{{0x0, 0x61f, 0x1}, {0x621, 0x640, 0x1f}, {0x64b, 0x66d, 0x1}, {0x670, 0x674, 0x4}, {0x6d4, 0x6d6, 0x2}, {0x6d7, 0x6ed, 0x1}, {0x6f0, 0x6f9, 0x1}, {0x6fd, 0x6fe, 0x1}, {0x700, 0x70f, 0x1}, {0x711, 0x730, 0x1f}, {0x731, 0x74c, 0x1}, {0x780, 0x85f, 0x1}, {0x86b, 0x89f, 0x1}, {0x8ad, 0x8b5, 0x8}, {0x8be, 0xffff, 0x1}}, []unicode.Range32{{0x10000, 0x10abf, 0x1}, {0x10ac6, 0x10ac8, 0x2}, {0x10acb, 0x10acc, 0x1}, {0x10ae2, 0x10ae3, 0x1}, {0x10ae5, 0x10aea, 0x1}, {0x10af0, 0x10ffff, 0x1}}, 0}
	joiningGroupNoon                = &unicode.RangeTable{[]unicode.Range16{{0x646, 0x6b9, 0x73}, {0x6ba, 0x6bc, 0x1}, {0x767, 0x769, 0x1}}, nil, 0}
	joiningGroupNun                 = &unicode.RangeTable{[]unicode.Range16{{0x722, 0x722, 0x1}}, nil, 0}
	joiningGroupNya                 = &unicode.RangeTable{[]unicode.Range16{{0x6bd, 0x6bd, 0x1}}, nil, 0}
	joiningGroupPe                  = &unicode.RangeTable{[]unicode.Range16{{0x726, 0x726, 0x1}}, nil, 0}
	joiningGroupQaf                 = &unicode.RangeTable{[]unicode.Range16{{0x642, 0x66f, 0x2d}, {0x6a7, 0x6a8, 0x1}, {0x8a5, 0x8a5, 0x1}}, nil, 0}
	joiningGroupQaph                = &unicode.RangeTable{[]unicode.Range16{{0x729, 0x729, 0x1}}, nil, 0}
	joiningGroupReh                 = &unicode.RangeTable{[]unicode.Range16{{0x631, 0x632, 0x1}, {0x691, 0x699, 0x1}, {0x6ef, 0x75b, 0x6c}, {0x76b, 0x76c, 0x1}, {0x771, 0x8aa, 0x139}, {0x8b2, 0x8b9, 0x7}}, nil, 0}
	joiningGroupReversedPe          = &unicode.RangeTable{[]unicode.Range16{{0x727, 0x727, 0x1}}, nil, 0}
	joiningGroupRohingyaYeh         = &unicode.RangeTable{[]unicode.Range16{{0x8ac, 0x8ac, 0x1}}, nil, 0}
	joiningGroupSad                 = &unicode.RangeTable{[]unicode.Range16{{0x635, 0x636, 0x1}, {0x69d, 0x69e, 0x1}, {0x6fb, 0x8af, 0x1b4}}, nil, 0}
	joiningGroupSadhe               = &unicode.RangeTable{[]unicode.Range16{{0x728, 0x728, 0x1}}, nil, 0}
	joiningGroupSeen                = &unicode.RangeTable{[]unicode.Range16{{0x633, 0x634, 0x1}, {0x69a, 0x69c, 0x1}, {0x6fa, 0x75c, 0x62}, {0x76d, 0x770, 0x3}, {0x77d, 0x77e, 0x1}}, nil, 0}
	joiningGroupSemkath             = &unicode.RangeTable{[]unicode.Range16{{0x723, 0x723, 0x1}}, nil, 0}
	joiningGroupShin                = &unicode.RangeTable{[]unicode.Range16{{0x72b, 0x72b, 0x1}}, nil, 0}
	joiningGroupStraightWaw         = &unicode.RangeTable{[]unicode.Range16{{0x8b1, 0x8b1, 0x1}}, nil, 0}
	joiningGroupSwashKaf            = &unicode.RangeTable{[]unicode.Range16{{0x6aa, 0x6aa, 0x1}}, nil, 0}
	joiningGroupSyriacWaw           = &unicode.RangeTable{[]unicode.Range16{{0x718, 0x718, 0x1}}, nil, 0}
	joiningGroupTah                 = &unicode.RangeTable{[]unicode.Range16{{0x637, 0x638, 0x1}, {0x69f, 0x8a3, 0x204}}, nil, 0}
	joiningGroupTaw                 = &unicode.RangeTable{[]unicode.Range16{{0x72c, 0x72c, 0x1}}, nil, 0}
	joiningGroupTehMarbuta          = &unicode.RangeTable{[]unicode.Range16{{0x629, 0x6c0, 0x97}, {0x6d5, 0x6d5, 0x1}}, nil, 0}
	joiningGroupHamzaOnHehGoal      = &unicode.RangeTable{[]unicode.Range16{{0x6c3, 0x6c3, 0x1}}, nil, 0}
	joiningGroupTeth                = &unicode.RangeTable{[]unicode.Range16{{0x71b, 0x71c, 0x1}}, nil, 0}
	joiningGroupWaw                 = &unicode.RangeTable{[]unicode.Range16{{0x624, 0x648, 0x24}, {0x676, 0x677, 0x1}, {0x6c4, 0x6cb, 0x1}, {0x6cf, 0x778, 0xa9}, {0x779, 0x8ab, 0x132}}, nil, 0}
	joiningGroupYeh                 = &unicode.RangeTable{[]unicode.Range16{{0x620, 0x626, 0x6}, {0x649, 0x64a, 0x1}, {0x678, 0x6d0, 0x58}, {0x6d1, 0x777, 0xa6}, {0x8a8, 0x8a9, 0x1}, {0x8ba, 0x8ba, 0x1}}, nil, 0}
	joiningGroupYehBarree           = &unicode.RangeTable{[]unicode.Range16{{0x6d2, 0x6d3, 0x1}}, nil, 0}
	joiningGroupYehWithTail         = &unicode.RangeTable{[]unicode.Range16{{0x6cd, 0x6cd, 0x1}}, nil, 0}
	joiningGroupYudh                = &unicode.RangeTable{[]unicode.Range16{{0x71d, 0x71d, 0x1}}, nil, 0}
	joiningGroupYudhHe              = &unicode.RangeTable{[]unicode.Range16{{0x71e, 0x71e, 0x1}}, nil, 0}
	joiningGroupZain                = &unicode.RangeTable{[]unicode.Range16{{0x719, 0x719, 0x1}}, nil, 0}
	joiningGroupZhain               = &unicode.RangeTable{[]unicode.Range16{{0x74d, 0x74d, 0x1}}, nil, 0}
)
