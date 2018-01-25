package pair

import (
	"fmt"
	"reflect"
	"sort"
)

// Pair is a struct combines two variables
// variables could be distinct type
// motivation comes from pair in STL utility in C++
type Pair struct {
	First  interface{}
	Second interface{}
}

// Pairs is a wrapper for slice of pair
// has basic sort method
type Pairs []Pair

// NewPair generates a pair
func NewPair(first, second interface{}) Pair {
	return Pair{first, second}
}

// MakePairs generates a Pairs
func MakePairs(pairs []Pair) Pairs {
	return Pairs(pairs)
}

// Sort func sorts the slice of pair
func Sort(pairs []Pair) {
	var ppairs = MakePairs(pairs)
	ppairs.Sort()
}

// SortReverse func sorts the slice of pair, but reverse
func SortReverse(pairs []Pair) {
	var ppairs = MakePairs(pairs)
	ppairs.SortReserve()
}

// SortFirst func sort the slice of pair with the first variable.
func SortFirst(pairs []Pair) {
	var ppairs = MakePairs(pairs)
	ppairs.SortFirst()
}

// SortFirstReverse func sort with first variable, but reverse
func SortFirstReverse(pairs []Pair) {
	var ppairs = MakePairs(pairs)
	ppairs.SortFirstReverse()
}

// SortSecond func sort the slice of pair with second variable.
func SortSecond(pairs []Pair) {
	var ppairs = MakePairs(pairs)
	ppairs.SortSecond()
}

// SortSecondReverse func sort the slice of pair with second variable, but reverse
func SortSecondReverse(pairs []Pair) {
	var ppairs = MakePairs(pairs)
	ppairs.SortSecondReverse()
}

// FString get the first variable as string
func (this *Pair) FString() string {
	return tostring(this.First)
}

// SString get the second variable as string
func (this *Pair) SString() string {
	return tostring(this.Second)
}

// FUint64 get the first variable as uint64
func (this *Pair) FUint64() uint64 {
	return touint64(this.First)
}

// SUint64 get the second variable as uint64
func (this *Pair) SUint64() uint64 {
	return touint64(this.Second)
}

// FInt64 get the first variable as int64
func (this *Pair) FInt64() int64 {
	return toint64(this.First)
}

// SInt64 get the second variable as int64
func (this *Pair) SInt64() int64 {
	return toint64(this.Second)
}

// FUint get the first variable as uint64
func (this *Pair) FUint() uint {
	return touint(this.First)
}

// SUint get the second variable as uint
func (this *Pair) SUint() uint {
	return touint(this.Second)
}

// FInt get the first variable as int
func (this *Pair) FInt() int {
	return toint(this.First)
}

// SInt get the second variable as int
func (this *Pair) SInt() int {
	return toint(this.Second)
}

// FFloat64 get the first variable as float64
func (this *Pair) FFloat64() float64 {
	return tofloat64(this.First)
}

// SInt get the second variable as float64
func (this *Pair) SFloat64() float64 {
	return tofloat64(this.Second)
}

// FFloat64 get the first variable as float32
func (this *Pair) FFloat32() float32 {
	return tofloat32(this.First)
}

// SInt get the second variable as float32
func (this *Pair) SFloat32() float32 {
	return tofloat32(this.Second)
}

// SortFirst is the native SortFirst method on Pairs
func (this *Pairs) SortFirst() {
	sortPairsByFirst(this, false)
}

// SortFirstReverse is the native SortFirstReverse method on Pairs
func (this *Pairs) SortFirstReverse() {
	sortPairsByFirst(this, true)
}

// SortSecond is the native SortSecond method on Pairs
func (this *Pairs) SortSecond() {
	sortPairsBySecond(this, false)
}

// SortSecondReverse is the native SortSecondReverse method on Paris
func (this *Pairs) SortSecondReverse() {
	sortPairsBySecond(this, true)
}

// Sort is the native Sort method on Pairs
func (this *Pairs) Sort() {
	sortPairs(this, false)
}

// SortReverse is the native SortReverse method on Pairs
func (this *Pairs) SortReserve() {
	sortPairs(this, true)
}

func sortPairs(this *Pairs, reverse bool) {
	if len(*this) <= 0 {
		return
	}
	var _1stfunc = lessfunc(this, "First")
	var _2ndfunc = lessfunc(this, "Second")
	var cond1 = (true != reverse)
	var cond2 = (false != reverse)
	sort.Slice(*this, func(i, j int) bool {
		var res1st = _1stfunc(i, j)
		if res1st != 0 {
			return tenary(res1st < 0, cond1, cond2)
		}
		var res2nd = _2ndfunc(i, j)
		return tenary(res2nd < 0, cond1, cond2)
	})
}

func sortPairsByFirst(this *Pairs, reverse bool) {
	if len(*this) <= 0 {
		return
	}
	var _1stfunc = lessfunc(this, "First")
	var cond1 = (true != reverse)
	var cond2 = (false != reverse)
	sort.Slice(*this, func(i, j int) bool {
		var res1st = _1stfunc(i, j)
		return tenary(res1st < 0, cond1, cond2)
	})
}

func sortPairsBySecond(this *Pairs, reverse bool) {
	if len(*this) <= 0 {
		return
	}
	var _2ndfunc = lessfunc(this, "Second")
	var cond1 = (true != reverse)
	var cond2 = (false != reverse)
	sort.Slice(*this, func(i, j int) bool {
		var res2nd = _2ndfunc(i, j)
		return tenary(res2nd < 0, cond1, cond2)
	})
}

func lessfunc(pairs *Pairs, fieldName string) func(i, j int) int {
	var sample = (*pairs)[0]
	var field = reflect.ValueOf(sample).FieldByName(fieldName).Interface()
	var kind = reflect.ValueOf(field).Kind()
	var common = func(i, j int) (interface{}, interface{}) {
		var vi = reflect.ValueOf((*pairs)[i]).FieldByName(fieldName).Interface()
		var vj = reflect.ValueOf((*pairs)[j]).FieldByName(fieldName).Interface()
		return vi, vj
	}

	switch kind {
	case reflect.Float64:
		return func(i, j int) int {
			var vi, vj = common(i, j)
			var vvi, oi = vi.(float64)
			var vvj, oj = vj.(float64)
			if !oi || !oj {
				pn(fieldName)
			}
			if vvi == vvj {
				return 0
			} else if vvi < vvj {
				return -1
			} else {
				return 1
			}
		}
	case reflect.Float32:
		return func(i, j int) int {
			var vi, vj = common(i, j)
			var vvi, oi = vi.(float32)
			var vvj, oj = vj.(float32)
			if !oi || !oj {
				pn(fieldName)
			}
			if vvi == vvj {
				return 0
			} else if vvi < vvj {
				return -1
			} else {
				return 1
			}
		}
	case reflect.Int64:
		return func(i, j int) int {
			var vi, vj = common(i, j)
			var vvi, oi = vi.(int64)
			var vvj, oj = vj.(int64)
			if !oi || !oj {
				pn(fieldName)
			}
			if vvi == vvj {
				return 0
			} else if vvi < vvj {
				return -1
			} else {
				return 1
			}
		}
	case reflect.Int32:
		return func(i, j int) int {
			var vi, vj = common(i, j)
			var vvi, oi = vi.(int32)
			var vvj, oj = vj.(int32)
			if !oi || !oj {
				pn(fieldName)
			}
			if vvi == vvj {
				return 0
			} else if vvi < vvj {
				return -1
			} else {
				return 1
			}
		}
	case reflect.Int:
		return func(i, j int) int {
			var vi, vj = common(i, j)
			var vvi, oi = vi.(int)
			var vvj, oj = vj.(int)
			if !oi || !oj {
				pn(fieldName)
			}
			if vvi == vvj {
				return 0
			} else if vvi < vvj {
				return -1
			} else {
				return 1
			}
		}
	case reflect.String:
		return func(i, j int) int {
			var vi, vj = common(i, j)
			var vvi, oi = vi.(string)
			var vvj, oj = vj.(string)
			if !oi || !oj {
				pn(fieldName)
			}
			if vvi == vvj {
				return 0
			} else if vvi < vvj {
				return -1
			} else {
				return 1
			}
		}
	default:
		panic(fmt.Sprintf("%s kind is not supported yet", kind.String()))
	}
}

func tenary(condition bool, yes, no bool) bool {
	if condition {
		return yes
	}
	return no
}

func tostring(x interface{}) string {
	var s = x.(string)
	return s
}

func touint64(x interface{}) uint64 {
	var i = x.(uint64)
	return i
}

func toint64(x interface{}) int64 {
	var i = x.(int64)
	return i
}

func touint(x interface{}) uint {
	var i = x.(uint)
	return i
}

func toint(x interface{}) int {
	var i = x.(int)
	return i
}

func tofloat64(x interface{}) float64 {
	var i = x.(float64)
	return i
}

func tofloat32(x interface{}) float32 {
	var i = x.(float32)
	return i
}

func pn(fieldname string) {
	var e = fmt.Sprintf("The %s field for sort should always be a same type", fieldname)
	panic(e)
}
