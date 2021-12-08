package common

type CharSet map[rune]struct{}

var present = struct{}{}

func MakeCharSet(s string) CharSet {
	cs := make(CharSet)
	for _, r := range s {
		cs.Put(r)
	}
	return cs
}

func (c CharSet) Put(r rune) {
	c[r] = present
}

func (c CharSet) Contains(r rune) bool {
	_, ok := c[r]
	return ok
}

func (c CharSet) Union(other CharSet) CharSet {
	res := make(CharSet)
	for r := range c {
		res.Put(r)
	}
	for r := range other {
		res.Put(r)
	}
	return res
}

func (c CharSet) Intersection(other CharSet) CharSet {
	res := make(CharSet)
	for r := range c {
		if other.Contains(r) {
			res.Put(r)
		}
	}
	return res
}

func (c CharSet) Remove(other CharSet) CharSet {
	res := make(CharSet)
	for r := range c {
		res.Put(r)
	}
	for r := range other {
		delete(res, r)
	}
	return res
}

func (c CharSet) Equals(other CharSet) bool {
	return len(c) == len(other) && len(c.Remove(other)) == 0
}
