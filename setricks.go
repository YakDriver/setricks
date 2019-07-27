package setricks

// Union takes two "sets", i.e., maps with string keys and int values, and
// creates the union of the two. Rather than modify either of the parameters,
// since maps are passed by reference, a new map is created and returned.
func Union(set1 map[string]int, set2 map[string]int) map[string]int {
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

// Intersect takes two "sets", i.e., maps with string keys and int values, and
// creates the intersection of the two.
func Intersect(set1 map[string]int, set2 map[string]int) map[string]int {
	intersection := make(map[string]int)

	// get everything in set2 that's also in set1
	for k, v := range set2 {
		if _, ok := set1[k]; ok {
			intersection[k] = v
		}
	}

	return intersection
}

// Except takes two "sets", i.e., maps with string keys and int values, and
// creates the exception of the two - everything in set1 that is not in set2.
func Except(set1 map[string]int, set2 map[string]int) map[string]int {
	except := make(map[string]int)

	for k, v := range set1 {
		if _, ok := set2[k]; !ok {
			except[k] = v
		}
	}

	return except
}
