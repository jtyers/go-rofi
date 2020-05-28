// This code is generated by https://github.com/jtyers/slice
// DO NOT EDIT!

// nfn: NewStringSlice

package godash

type chainString struct {
  isPtr bool
	value []string
}

func NewStringSlice(slice []string) *chainString {
	return &chainString{
		value: slice,
		
	}
}

func (c *chainString) Value() []string {
	return c.value
}

func ConcatString(slice []string, slice2 []string) (res []string) {
	res = make([]string, 0, len(slice) + len(slice2))
	for _, entry := range slice {
		res = append(res, entry)
	}
	for _, entry := range slice2 {
		res = append(res, entry)
	}
	return
}

func (c *chainString) Concat(slice2 []string) *chainString {
	return &chainString{value: ConcatString(c.value, slice2)}
}

func DropString(slice []string, n int) (res []string) {
	l := len(slice) - n
	if l < 0 {
		l = 0
	}
	res = make([]string, 0, l)
	for _, entry := range slice[len(slice) - l:] {
		res = append(res, entry)
	}
	return
}

func (c *chainString) Drop(n int) *chainString {
	return &chainString{value: DropString(c.value, n)}
}

func DropRightString(slice []string, n int) (res []string) {
	l := len(slice) - n
	if l < 0 {
		l = 0
	}
	res = make([]string, 0, l)
	for _, entry := range slice[:l] {
		res = append(res, entry)
	}
	return
}

func (c *chainString) DropRight(n int) *chainString {
	return &chainString{value: DropRightString(c.value, n)}
}

func FilterString(slice []string, fn func(string,int)bool) (res []string) {
	res = make([]string, 0, len(slice))
	for index, entry := range slice {
		if fn(entry, index) {
			res = append(res, entry)
		}
	}
	return
}

func (c *chainString) Filter(fn func(string,int)bool) *chainString {
	return &chainString{value: FilterString(c.value, fn)}
}

func FirstString(slice []string) (res string) {
	if len(slice) == 0 {
		return
	}
	res = slice[0]
	return
}

func (c *chainString) First() *chainString {
	return &chainString{value: []string{FirstString(c.value)}}
}

func LastString(slice []string) (res string) {
	if len(slice) == 0 {
		return
	}
	res = slice[len(slice) - 1]
	return
}

func (c *chainString) Last() *chainString {
	return &chainString{value: []string{LastString(c.value)}}
}

func MapString(slice []string, fn func(string,int)string) (res []string) {
	res = make([]string, 0, len(slice))
	for index, entry := range slice {
		res = append(res, fn(entry, index))
	}
	return
}

func (c *chainString) Map(fn func(string,int)string) *chainString {
	return &chainString{value: MapString(c.value, fn)}
}


func ReduceString(slice []string, fn func(string,string,int)string, initial string) (res string) {
	res = initial
	for index, entry := range slice {
		res = fn(res, entry, index)
	}
	return
}

func (c *chainString) Reduce(fn func(string,string,int)string, initial string) *chainString {
	return &chainString{value: []string{ReduceString(c.value, fn, initial)}}
}

func ReverseString(slice []string) (res []string) {
	res = make([]string, len(slice))
	for index, entry := range slice {
		res[len(slice)-1-index] = entry
	}
	return
}

func (c *chainString) Reverse() *chainString {
	return &chainString{value: ReverseString(c.value)}
}

func UniqString(slice []string) (res []string) {
	seen := make(map[string]bool)
	res = []string{}
	for _, entry := range slice {
		if _, found := seen[entry]; !found {
			seen[entry] = true
			res = append(res, entry)
		}
	}
	return
}

func (c *chainString) Uniq() *chainString {
	if c.isPtr {
		panic("Uniq() does not currently support pointers")
	}
	return &chainString{value: UniqString(c.value)}
}
