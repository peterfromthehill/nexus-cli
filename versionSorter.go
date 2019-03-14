package main

import ("fmt"
		"strings"
		"unicode"
		"math"
		"sort"
	)

func s(s1, s2 string) bool {
	if Filevercmp(s1,s2) < 0 {
		return true
	} else {
		return false
	}
}


// Sort sort something
func Sort(strs []string) {
		strSort := &strSorter{
			strs: strs,
			cmp:  s,
		}
		sort.Sort(strSort)
	}
	
type strSorter struct {
		strs []string
		cmp  func(str1, str2 string) bool
}

func (s *strSorter) Len() int { return len(s.strs) }
	
func (s *strSorter) Swap(i, j int) { s.strs[i], s.strs[j] = s.strs[j], s.strs[i] }
	
func (s *strSorter) Less(i, j int) bool { return s.cmp(s.strs[i], s.strs[j]) }
	


func cIsdigit(c byte) bool {
    return unicode.IsDigit(int32(c));
}

func cIsalpha(c byte) bool {
	return unicode.IsLetter(int32(c))
}

func cIsalnum(c byte) bool {
	fmt.Printf("%c", c)
  return cIsalpha(c) || cIsdigit(c);
}

func strncmp(s1 string, s2 string, l int) int {
    len1 := math.Min(float64(l), float64(len(s1)))
    len2 := math.Min(float64(l), float64(len(s2)))
    return strings.Compare(s1[:int(len1)], s2[:int(len2)]);
  }

func matchSuffix(str string) string {
	match := ""
	readAlpha := false
	for ; len(str) > 0 ; {
		if readAlpha {
			readAlpha = false
			if !cIsalpha(str[0])  && '~' != str[0] {
				match = ""
			} 
		} else if '.' == str[0] {
			readAlpha = true
			if len(match) == 0 {
				match = str
			}
		} else if !cIsalnum(str[0]) && '~' != str[0] {
			match = ""
		}
		str = str[1:len(str)]
	}
	return match
}

func order(c byte) int {
	if cIsdigit(c) {
		return 0
	} else if cIsalpha(c) {
		return int(c)
	} else if c == '~' {
		return -1
	} else {
		return int(c) + 1 + unicode.MaxRune
	}

}

func verrevcmp(s1 string, s2 string) int {
	s1Pos := 0
	s2Pos := 0

	for ; s1Pos < len(s1) || s2Pos < len(s2); {
		firstDiff := 0
		for ; (s1Pos < len(s1) && !cIsdigit(s1[s1Pos])) || (s2Pos < len(s2) && !cIsdigit(s2[s2Pos])); {
			s1C := 0
			if s1Pos < len(s1) {
				s1C = order(s1[s1Pos])
			}
			s2C := 0
			if s2Pos < len(s2) {
				s2C = order(s2[s2Pos])
			}
			if s1C != s2C {
				return s1C - s2C
			}
			s1Pos++
			s2Pos++
		}
		for ; s1Pos < len(s1) && s1[s1Pos] == '0'; {
			s1Pos++
		}
		for ; s2Pos < len(s2) && s2[s2Pos] == '0'; {
			s2Pos++
		}
		for ; (s1Pos < len(s1) && s2Pos < len(s2) && cIsdigit(s1[s1Pos]) && cIsdigit(s2[s2Pos])); {
			if firstDiff == 0 {
				firstDiff = int(s1[s1Pos]) - int(s2[s2Pos])
			}
			s1Pos++
			s2Pos++
		}
		if s1Pos < len(s1) && cIsdigit(s1[s1Pos]) {
			return 1
		}
		if s2Pos < len(s2) && cIsdigit(s2[s2Pos]) {
			return -1
		}
		if firstDiff != 0 {
			return firstDiff
		}
	}
	return 0
}

// Filevercmp compares to strings
func Filevercmp(s1 string, s2 string) int {
	s1Pos := 0
	s2Pos := 0

	result := 0
	simpleCmp := strings.Compare (s1, s2);
	if(simpleCmp == 0) {
		return 0
	}

	if (s1 == "") {
		return -1
	}
	if (s2 == "") {
		return 1
	}
	if (s1 == ".") {
		return -1
	}
	if (s2 == ".") {
		return 1
	}
	if (s1 == "..") {
		return -1
	}
	if (s2 == "..") {
		return 1
	}

	if (s1[s1Pos] == '.' && s2[s2Pos] != '.') {
		return -1
	}
	if (s1[s1Pos] != '.' && s2[s2Pos] == '.') {
		return 1
	}
	if (s1[s1Pos] == '.' && s2[s2Pos] == '.') {
		s1 = s1[1:]
		s2 = s2[1:]
	}

	s1Suffix := matchSuffix(s1)
	s2Suffix := matchSuffix(s2)

	s1Len := len(s1) - len(s1Suffix)
	s2Len := len(s2) - len(s2Suffix)

	if ((len(s1Suffix) > 0 || len(s2Suffix) > 0) && (s1Len == s2Len) && 0 == strings.Compare(s1[:s1Len], s2[:s1Len])) {
		s1Len = len(s1)
		s2Len = len(s2)
	}
	

	result = verrevcmp(s1[:s1Len], s2[:s2Len])
	if result == 0 {
		return simpleCmp 
	}
	return result
}



