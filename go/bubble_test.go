package sort

import (
	"reflect"
	"testing"
)

func TestBubbleInts(t *testing.T) {
	tests := []struct {
		arr      []int
		expected []int
	}{
		{
			arr:      []int{301, 587, 1356, 238, 777, 1095, 352, 662, 1431, 1192, 1833, 950, 915, 687, 1903, 218},
			expected: []int{218, 238, 301, 352, 587, 662, 687, 777, 915, 950, 1095, 1192, 1356, 1431, 1833, 1903},
		},
		{
			arr:      []int{1340, 1153, 71, 10, 710, 560, 298, 1777, 1444, 867, 1484, 846, 510, 1023, 220, 639},
			expected: []int{10, 71, 220, 298, 510, 560, 639, 710, 846, 867, 1023, 1153, 1340, 1444, 1484, 1777},
		},
		{
			arr:      []int{544, 1596, 991, 177, 551, 74, 1496, 450, 978, 1084, 1982, 752, 1329, 541, 988, 1142},
			expected: []int{74, 177, 450, 541, 544, 551, 752, 978, 988, 991, 1084, 1142, 1329, 1496, 1596, 1982},
		},
		{
			arr:      []int{929, 875, 739, 515, 1621, 1079, 1249, 558, 212, 1935, 1688, 129, 1869, 1622, 1999, 340},
			expected: []int{129, 212, 340, 515, 558, 739, 875, 929, 1079, 1249, 1621, 1622, 1688, 1869, 1935, 1999},
		},
		{
			arr:      []int{1143, 1138, 865, 1352, 1901, 360, 1192, 126, 1183, 751, 1992, 388, 1008, 743, 766, 1800},
			expected: []int{126, 360, 388, 743, 751, 766, 865, 1008, 1138, 1143, 1183, 1192, 1352, 1800, 1901, 1992},
		},
		{
			arr:      []int{64, 1622, 1202, 1476, 1906, 1515, 1490, 248, 1784, 824, 454, 1038, 1635, 1354, 754, 1639},
			expected: []int{64, 248, 454, 754, 824, 1038, 1202, 1354, 1476, 1490, 1515, 1622, 1635, 1639, 1784, 1906},
		},
		{
			arr:      []int{162, 1179, 1394, 1297, 763, 161, 1900, 1268, 1629, 517, 1615, 305, 308, 416, 1951, 1110},
			expected: []int{161, 162, 305, 308, 416, 517, 763, 1110, 1179, 1268, 1297, 1394, 1615, 1629, 1900, 1951},
		},
		{
			arr:      []int{412, 381, 1306, 130, 1455, 790, 1877, 661, 1521, 394, 1646, 983, 683, 492, 1509, 1257},
			expected: []int{130, 381, 394, 412, 492, 661, 683, 790, 983, 1257, 1306, 1455, 1509, 1521, 1646, 1877},
		},
		{
			arr:      []int{808, 874, 187, 64, 1732, 998, 1804, 1290, 839, 1575, 1722, 13, 536, 640, 1835, 422},
			expected: []int{13, 64, 187, 422, 536, 640, 808, 839, 874, 998, 1290, 1575, 1722, 1732, 1804, 1835},
		},
		{
			arr:      []int{803, 1060, 1628, 1724, 1080, 1361, 1719, 1010, 1647, 1216, 789, 1179, 655, 947, 1529, 946},
			expected: []int{655, 789, 803, 946, 947, 1010, 1060, 1080, 1179, 1216, 1361, 1529, 1628, 1647, 1719, 1724},
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			Bubble(tt.arr)
			if !reflect.DeepEqual(tt.arr, tt.expected) {
				t.Errorf("Not sorted: Got %v, expected %v", tt.arr, tt.expected)
			}
		})
	}
}

func BenchmarkBubbleInts(b *testing.B) {
	for n := 0; n < b.N; n++ {
		arr := []int{247, 499, 238, 629, 1215, 1877, 1071, 1981, 1369, 597, 1388, 716, 763, 1505, 607, 1342, 179, 909, 1963, 1448, 608, 1125, 1358, 856, 947, 585, 1569, 1991, 244, 375, 850, 1674, 1889, 1570, 871, 1006, 1753, 1247, 1749, 1824, 1035, 746, 1923, 1957, 1747, 1821, 120, 1948, 739, 1381, 775, 1433, 231, 1915, 1888, 489, 1479, 1571, 109, 1068, 1906, 92, 699, 716, 851, 433, 867, 1326, 1288, 371, 1482, 222, 1997, 555, 472, 139, 759, 1976, 496, 30, 888, 383, 1330, 343, 864, 1432, 1733, 1442, 1088, 1118, 923, 223, 1777, 1541, 1856, 408, 711, 1772, 642, 1632, 1262, 1083, 982, 1953, 922, 389, 1248, 1997, 1784, 1333, 1167, 206, 185, 1984, 1278, 1329, 1498, 1655, 1407, 1494, 850, 1337, 1927, 234, 330, 916, 1339, 1411, 612, 1749, 939, 1759, 653, 336, 858, 1389, 888, 1385, 1789, 510, 1861, 431, 220, 1744, 178, 1006, 1229, 1617, 1244, 381, 742, 1669, 1091, 382, 554, 1491, 1215, 1879, 1196, 974}
		Bubble(arr)
	}
}

func TestOptimizedBubbleInt(t *testing.T) {
	tests := []struct {
		arr      []int
		expected []int
	}{
		{
			arr:      []int{301, 587, 1356, 238, 777, 1095, 352, 662, 1431, 1192, 1833, 950, 915, 687, 1903, 218},
			expected: []int{218, 238, 301, 352, 587, 662, 687, 777, 915, 950, 1095, 1192, 1356, 1431, 1833, 1903},
		},
		{
			arr:      []int{1340, 1153, 71, 10, 710, 560, 298, 1777, 1444, 867, 1484, 846, 510, 1023, 220, 639},
			expected: []int{10, 71, 220, 298, 510, 560, 639, 710, 846, 867, 1023, 1153, 1340, 1444, 1484, 1777},
		},
		{
			arr:      []int{544, 1596, 991, 177, 551, 74, 1496, 450, 978, 1084, 1982, 752, 1329, 541, 988, 1142},
			expected: []int{74, 177, 450, 541, 544, 551, 752, 978, 988, 991, 1084, 1142, 1329, 1496, 1596, 1982},
		},
		{
			arr:      []int{929, 875, 739, 515, 1621, 1079, 1249, 558, 212, 1935, 1688, 129, 1869, 1622, 1999, 340},
			expected: []int{129, 212, 340, 515, 558, 739, 875, 929, 1079, 1249, 1621, 1622, 1688, 1869, 1935, 1999},
		},
		{
			arr:      []int{1143, 1138, 865, 1352, 1901, 360, 1192, 126, 1183, 751, 1992, 388, 1008, 743, 766, 1800},
			expected: []int{126, 360, 388, 743, 751, 766, 865, 1008, 1138, 1143, 1183, 1192, 1352, 1800, 1901, 1992},
		},
		{
			arr:      []int{64, 1622, 1202, 1476, 1906, 1515, 1490, 248, 1784, 824, 454, 1038, 1635, 1354, 754, 1639},
			expected: []int{64, 248, 454, 754, 824, 1038, 1202, 1354, 1476, 1490, 1515, 1622, 1635, 1639, 1784, 1906},
		},
		{
			arr:      []int{162, 1179, 1394, 1297, 763, 161, 1900, 1268, 1629, 517, 1615, 305, 308, 416, 1951, 1110},
			expected: []int{161, 162, 305, 308, 416, 517, 763, 1110, 1179, 1268, 1297, 1394, 1615, 1629, 1900, 1951},
		},
		{
			arr:      []int{412, 381, 1306, 130, 1455, 790, 1877, 661, 1521, 394, 1646, 983, 683, 492, 1509, 1257},
			expected: []int{130, 381, 394, 412, 492, 661, 683, 790, 983, 1257, 1306, 1455, 1509, 1521, 1646, 1877},
		},
		{
			arr:      []int{808, 874, 187, 64, 1732, 998, 1804, 1290, 839, 1575, 1722, 13, 536, 640, 1835, 422},
			expected: []int{13, 64, 187, 422, 536, 640, 808, 839, 874, 998, 1290, 1575, 1722, 1732, 1804, 1835},
		},
		{
			arr:      []int{803, 1060, 1628, 1724, 1080, 1361, 1719, 1010, 1647, 1216, 789, 1179, 655, 947, 1529, 946},
			expected: []int{655, 789, 803, 946, 947, 1010, 1060, 1080, 1179, 1216, 1361, 1529, 1628, 1647, 1719, 1724},
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			OptimizedBubble(tt.arr)
			if !reflect.DeepEqual(tt.arr, tt.expected) {
				t.Errorf("Not sorted: Got %v, expected %v", tt.arr, tt.expected)
			}
		})
	}
}

func BenchmarkOptimizedBubbleInts(b *testing.B) {
	for n := 0; n < b.N; n++ {
		arr := []int{247, 499, 238, 629, 1215, 1877, 1071, 1981, 1369, 597, 1388, 716, 763, 1505, 607, 1342, 179, 909, 1963, 1448, 608, 1125, 1358, 856, 947, 585, 1569, 1991, 244, 375, 850, 1674, 1889, 1570, 871, 1006, 1753, 1247, 1749, 1824, 1035, 746, 1923, 1957, 1747, 1821, 120, 1948, 739, 1381, 775, 1433, 231, 1915, 1888, 489, 1479, 1571, 109, 1068, 1906, 92, 699, 716, 851, 433, 867, 1326, 1288, 371, 1482, 222, 1997, 555, 472, 139, 759, 1976, 496, 30, 888, 383, 1330, 343, 864, 1432, 1733, 1442, 1088, 1118, 923, 223, 1777, 1541, 1856, 408, 711, 1772, 642, 1632, 1262, 1083, 982, 1953, 922, 389, 1248, 1997, 1784, 1333, 1167, 206, 185, 1984, 1278, 1329, 1498, 1655, 1407, 1494, 850, 1337, 1927, 234, 330, 916, 1339, 1411, 612, 1749, 939, 1759, 653, 336, 858, 1389, 888, 1385, 1789, 510, 1861, 431, 220, 1744, 178, 1006, 1229, 1617, 1244, 381, 742, 1669, 1091, 382, 554, 1491, 1215, 1879, 1196, 974}
		OptimizedBubble(arr)
	}
}
