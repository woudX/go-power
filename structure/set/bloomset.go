package set

import "github.com/Workiva/go-datastructures/bitarray"

//	Bloom set
// k = 0.7 *(l / n)
// f = 0.6185^(l/n)

// n - element num predict
// l - bit array len
// k - hash function num
// f - fault rate

// Input : f, n --> l, k
type BloomSet struct {
	//	fault rate
	f float64

	//	element num
	n int64

	//	hash num
	k int

	//	bit array length
	l int64

	bitArray bitarray.BitArray
}

func NewBloomSet(n int64, f float64) {

}

//	Insert element to set, O(1)
func (s *BloomSet) Insert(interface{}) {

}

//	Clear all elements in set, O(n)?
func (s *BloomSet) Clear() {

}

//	Remove one element from set
func (s *BloomSet) Remove(interface{}) {

}

//
func (s *BloomSet) Contain(interface{}) {

}
