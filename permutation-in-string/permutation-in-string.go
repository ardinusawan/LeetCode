// build 2 map, s1Map and s2Map
// s1Map contains every char in s1, same with s2Map
// left and right pointer start and end based on length of s1, in s2
// only for first time, it need to compare s1Map and s2Map matches
// after that, relay on var mathes. if matches in the end = 26 then it equal
func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	s1Map, s2Map := make(map[byte]int, 26), make(map[byte]int, 26)
	for i := 0; i < len(s1); i++ {
		s1Map[s1[i]]++
		s2Map[s2[i]]++
	}

	windowSize := len(s1)
	for i := 0; i < len(s2)-windowSize; i++ {
		if mapsMatch(s1Map, s2Map) {
			return true
		}

		s2Map[s2[i]]--
		if s2Map[s2[i]] == 0 {
			delete(s2Map, s2[i])
		}
		s2Map[s2[i+windowSize]]++
	}

	return mapsMatch(s1Map, s2Map)
}

func mapsMatch(m1, m2 map[byte]int) bool {
	if len(m1) != len(m2) {
		return false
	}

	for k, v := range m1 {
		if m2[k] != v {
			return false
		}
	}

	return true
}
