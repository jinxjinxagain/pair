package pair

import (
	"crypto/rand"
	mrand "math/rand"
	"reflect"
	"strconv"
	"testing"
)

func TestGetMethod1(t *testing.T) {
	var pair = NewPair("1", 1)
	var first = pair.FString()
	var second = pair.SInt()
	expected(t, "1" == first, "1", first)
	expected(t, 1 == second, 1, second)
}

func TestGetMethod2(t *testing.T) {
	var pair = NewPair(1.1, uint64(3))
	var first = pair.FFloat64()
	var second = pair.SUint64()
	var fkind = reflect.ValueOf(first).Kind()
	var skind = reflect.ValueOf(second).Kind()
	expected(t, reflect.Float64 == fkind, reflect.Float64, fkind)
	expected(t, reflect.Uint64 == skind, reflect.Uint64, skind)
}

func TestNewPair1(t *testing.T) {
	var pair = NewPair("123", 123)
	var f, o1 = pair.First.(string)
	assert(t, o1)
	expected(t, "123" == f, "123", f)
	var s, o2 = pair.Second.(int)
	assert(t, o2)
	expected(t, 123 == s, 123, s)
}

func TestNewPair2(t *testing.T) {
	var pair = NewPair(123, 123.4)
	var f, o1 = pair.First.(int)
	assert(t, o1)
	expected(t, 123 == f, "123", f)
	var s, o2 = pair.Second.(float64)
	assert(t, o2)
	expected(t, 123.4 == s, 123.4, s)
}

func TestMakePair(t *testing.T) {
	var arr = []Pair{}
	for i := 0; i < 10; i++ {
		var f = randstring()
		var s = randint()
		var p = NewPair(f, s)
		arr = append(arr, p)
	}
	var pairs = MakePairs(arr)
	var kind = reflect.ValueOf(pairs).Kind()
	expected(t, reflect.Slice == kind, reflect.Slice.String(), kind.String())
}

func TestSortOnStruct(t *testing.T) {
	var data = []Pair{}
	var seq = randsequence(10)
	for _, i := range seq {
		var si = strconv.Itoa(i)
		data = append(data, NewPair(i, si))
	}
	var pairs = MakePairs(data)
	pairs.Sort()
	for i, p := range pairs {
		expected(t, i == p.FInt(), i, p.FInt())
	}
}

func TestSortReverseOnStruct(t *testing.T) {
	var data = []Pair{}
	var seq = randsequence(10)
	for _, i := range seq {
		var si = strconv.Itoa(i)
		data = append(data, NewPair(i, si))
	}
	var pairs = MakePairs(data)
	pairs.SortReserve()
	for i, p := range pairs {
		expected(t, i == len(data)-p.FInt()-1, i, len(data)-p.FInt()-1)
	}
}

func TestSortFunc(t *testing.T) {
	var data = []Pair{}
	var seq = randsequence(10)
	for _, i := range seq {
		var si = strconv.Itoa(i)
		data = append(data, NewPair(i, si))
	}
	Sort(data)
	for i, p := range data {
		expected(t, i == p.FInt(), i, p.FInt())
	}
}

func TestSortReverseFunc(t *testing.T) {
	var data = []Pair{}
	var seq = randsequence(10)
	for _, i := range seq {
		var si = strconv.Itoa(i)
		data = append(data, NewPair(i, si))
	}
	SortReverse(data)
	for i, p := range data {
		expected(t, i == len(data)-p.FInt()-1, i, len(data)-p.FInt()-1)
	}
}

func TestSortByFirstFunc(t *testing.T) {
	var data = []Pair{
		Pair{2, "3"},
		Pair{2, "1"},
		Pair{2, "2"},
		Pair{6, "7"},
		Pair{3, "4"},
		Pair{6, "6"},
	}
	var ans = []int{2, 2, 2, 3, 6, 6}
	SortFirst(data)
	for i := range data {
		expected(t, ans[i] == data[i].FInt(), ans[i], data[i].FInt())
	}
}

func TestSortByFirstReverseFunc(t *testing.T) {
	var data = []Pair{
		Pair{2, "3"},
		Pair{2, "1"},
		Pair{2, "2"},
		Pair{6, "7"},
		Pair{3, "4"},
		Pair{6, "6"},
	}
	var ans = []int{6, 6, 3, 2, 2, 2}
	SortFirstReverse(data)
	for i := range data {
		expected(t, ans[i] == data[i].FInt(), ans[i], data[i].FInt())
	}
}
func TestSortBySecondFunc(t *testing.T) {
	var data = []Pair{
		Pair{1, "6"},
		Pair{2, "5"},
		Pair{3, "4"},
		Pair{4, "3"},
		Pair{5, "2"},
		Pair{6, "1"},
	}
	var ans = []int{6, 5, 4, 3, 2, 1}
	SortSecond(data)
	for i := range data {
		expected(t, ans[i] == data[i].FInt(), ans[i], data[i].FInt())
		expected(t, strconv.Itoa(len(data)-ans[i]+1) == data[i].SString(), strconv.Itoa(len(data)-ans[i]+1), data[i].SString())
	}
}

func TestSortBySecondReverseFunc(t *testing.T) {
	var data = []Pair{
		Pair{6, "1"},
		Pair{5, "2"},
		Pair{4, "3"},
		Pair{3, "4"},
		Pair{2, "5"},
		Pair{1, "6"},
	}
	var ans = []int{1, 2, 3, 4, 5, 6}
	SortSecondReverse(data)
	for i := range data {
		expected(t, ans[i] == data[i].FInt(), ans[i], data[i].FInt())
		expected(t, strconv.Itoa(len(data)-ans[i]+1) == data[i].SString(), strconv.Itoa(len(data)-ans[i]+1), data[i].SString())
	}
}

func assert(t testing.TB, cond bool) {
	if !cond {
		panic("assert failed")
	}
}

func expected(t testing.TB, cond bool, expected, got interface{}) {
	if cond {
		return
	}
	t.Fatalf("expected '%v', but got '%v'", expected, got)
}

func randstring() string {
	var c = 10
	var b = make([]byte, c)
	rand.Read(b)
	return string(b)
}

func randint() int {
	return mrand.Intn(100)
}

func randsequence(n int) []int {
	var a = make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = i
	}
	for i := n - 1; i >= 0; i-- {
		var j = mrand.Intn(i + 1)
		a[i], a[j] = a[j], a[i]
	}
	return a
}
