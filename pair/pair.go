package pair

type Pair struct {
	Key   int
	Value int
}

type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }

func NewPairListFromMap(mp map[int]int) PairList {
	res := make(PairList, 0)

	for k, v := range mp {
		res = append(res, Pair{
			Key:   k,
			Value: v,
		})
	}

	return res
}
