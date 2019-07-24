package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getData(filename string) ([]byte, error) {
	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return dat, nil
}

func getSet(rawSet []byte) map[string]int {
	re := regexp.MustCompile(`\s+(.*)\s(\d+)`)
	found := re.FindAllSubmatch(rawSet, -1)

	set := make(map[string]int)
	for _, sub := range found {
		i, _ := strconv.Atoi(string(sub[2]))
		set[string(sub[1])] = i
	}

	return set
}

func union(set1 map[string]int, set2 map[string]int) map[string]int {
	// this could be done more simply by using the parameters but
	// since maps are passed by reference that would change the map
	// in the calling scope, which is not what we want here

	// just copying set1
	union := make(map[string]int)
	for k, v := range set1 {
		union[k] = v
	}

	for k, v := range set2 {
		union[k] = v
	}

	return union
}

func intersect(set1 map[string]int, set2 map[string]int) map[string]int {
	intersection := make(map[string]int)

	// get everything in set2 that's also in set1
	for k, v := range set2 {
		if _, ok := set1[k]; ok {
			intersection[k] = v
		}
	}

	return intersection
}

func except(set1 map[string]int, set2 map[string]int) map[string]int {
	except := make(map[string]int)

	for k, v := range set1 {
		if _, ok := set2[k]; !ok {
			except[k] = v
		}
	}

	return except
}

func main() {
	data1, err := getData("set1.data")
	check(err)

	data2, err := getData("set2.data")
	check(err)

	set1 := getSet(data1)
	set2 := getSet(data2)

	unioned := union(set1, set2)
	intersection := intersect(set1, set2)
	exception1 := except(set1, set2)
	exception2 := except(set2, set1)

	fmt.Printf("Union: %v\n\n", unioned)
	fmt.Printf("Intersection: %v\n\n", intersection)
	fmt.Printf("Except 1: %v\n\n", exception1)
	fmt.Printf("Except 2: %v\n\n", exception2)
}
