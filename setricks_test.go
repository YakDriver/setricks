package setricks

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"reflect"
	"regexp"
	"runtime"
	"strconv"
	"testing"
)

func TestUnion(t *testing.T) {
	data1, err := getData(filepath.Join("testdata", "set1.data"))
	check(err)

	data2, err := getData(filepath.Join("testdata", "set2.data"))
	check(err)

	set1 := getSet(data1)
	set2 := getSet(data2)

	unioned := Union(set1, set2)

	answer := make(map[string]int)
	answer["blue"] = 30
	answer["green"] = 58
	answer["orange"] = 88
	answer["red"] = 42
	answer["yellow"] = 18

	equals(t, unioned, answer)
}

func TestIntersect(t *testing.T) {
	data1, err := getData(filepath.Join("testdata", "set1.data"))
	check(err)

	data2, err := getData(filepath.Join("testdata", "set2.data"))
	check(err)

	set1 := getSet(data1)
	set2 := getSet(data2)

	intersection := Intersect(set1, set2)

	answer := make(map[string]int)
	answer["blue"] = 30

	equals(t, intersection, answer)
}

func TestExcept1(t *testing.T) {
	data1, err := getData(filepath.Join("testdata", "set1.data"))
	check(err)

	data2, err := getData(filepath.Join("testdata", "set2.data"))
	check(err)

	set1 := getSet(data1)
	set2 := getSet(data2)

	exception := Except(set1, set2)

	answer := make(map[string]int)
	answer["green"] = 58
	answer["red"] = 42

	equals(t, exception, answer)
}

func TestExcept2(t *testing.T) {
	data1, err := getData(filepath.Join("testdata", "set1.data"))
	check(err)

	data2, err := getData(filepath.Join("testdata", "set2.data"))
	check(err)

	set1 := getSet(data1)
	set2 := getSet(data2)

	exception := Except(set2, set1)

	answer := make(map[string]int)
	answer["orange"] = 88
	answer["yellow"] = 18

	equals(t, exception, answer)
}

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

// ben johnson's helper functions: https://github.com/benbjohnson/testing

// assert fails the test if the condition is false.
func assert(tb testing.TB, condition bool, msg string, v ...interface{}) {
	if !condition {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d: "+msg+"\033[39m\n\n", append([]interface{}{filepath.Base(file), line}, v...)...)
		tb.FailNow()
	}
}

// ok fails the test if an err is not nil.
func ok(tb testing.TB, err error) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d: unexpected error: %s\033[39m\n\n", filepath.Base(file), line, err.Error())
		tb.FailNow()
	}
}

// equals fails the test if exp is not equal to act.
func equals(tb testing.TB, exp, act interface{}) {
	if !reflect.DeepEqual(exp, act) {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d:\n\n\texp: %#v\n\n\tgot: %#v\033[39m\n\n", filepath.Base(file), line, exp, act)
		tb.FailNow()
	}
}
