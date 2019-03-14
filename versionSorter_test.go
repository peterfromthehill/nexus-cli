package main

import ("testing"
		"fmt" 
	)

func Test_SortMixed(t *testing.T) {
	tags := []string{"latest", "1.0.1"}

	Sort(tags)

	if tags[0] != "1.0.1" && tags[1] != "latest" {
		t.Errorf("ordering incorrect when checking mixed tags")
	}
}

func Test_SortAllDigits(t *testing.T) {
	tags := []string{"1.2.1", "1.0.1"}

	Sort(tags)

	if tags[0] != "1.0.1" && tags[1] != "1.2.1" {
		t.Errorf("ordering incorrect in all digits tags")
	}
}

func Test_SortMixedExxetaStyle(t *testing.T) {
	tags := []string{"1.2.9-SNAPSHOT","1.2.10-SNAPSHOT","1.2.7-SNAPSHOT","1.2.9","1.2.4","latest","1.2.8","1.2.5","1.2.6","1.2.8-SNAPSHOT","1.2.7"}


	Sort(tags)
	sortedTags := []string{"1.2.4","1.2.5","1.2.6","1.2.7","1.2.7-SNAPSHOT","1.2.8","1.2.8-SNAPSHOT","1.2.9","1.2.9-SNAPSHOT","1.2.10-SNAPSHOT","latest"}
	
	tagslength := len(tags)
	fmt.Printf("\ntags:       %v\n", tags)
	fmt.Printf("sortedTags: %v\n", sortedTags)

	for i := 0; i < tagslength; i++ {
		if tags[i] != sortedTags[i] {
			t.Errorf("ordering incorrect in all digits tags")
		}
	}
}



