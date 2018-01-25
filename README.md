# Pair
Pair is a simple Go library for combining two variables with any type.

# Install
go get -u github.com/jinxjinxagain/pair

# Example
```golang
var pairs = []Pair{
  Pair{2, "two"},
  Pair{3, "three"},
  Pair{1, "one"},
  Pair{6, "six"},
  Pair{5, "five"},
  Pair{4, "four"},
  Pair{1, "one_again"}
}
pair.Sort(pairs)
for _, p := range pairs {
  fmt.Println(p.SString())
}
```

```golang
one
one_again
two
three
four
five
six
```

# License
Source code is available under the MIT License.