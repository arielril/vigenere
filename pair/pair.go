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

type PairRuneSlice struct {
	Key   int
	Value []rune
}
type PairListRuneSlice []PairRuneSlice

func (p PairListRuneSlice) Len() int           { return len(p) }
func (p PairListRuneSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairListRuneSlice) Less(i, j int) bool { return p[i].Value[0] < p[j].Value[0] }

func NewPairListRuleSliceFromMap(mp map[int][]rune) PairListRuneSlice {
	rt := make(PairListRuneSlice, 0)
	for k, v := range mp {
		rt = append(rt, PairRuneSlice{
			Key:   k,
			Value: v,
		})
	}
	return rt
}

type PairFrequencyScoreAndKey struct {
	Key   string
	Value int
}

type PairFrequencyScoreAndKeyList []PairFrequencyScoreAndKey

func (p PairFrequencyScoreAndKeyList) Len() int           { return len(p) }
func (p PairFrequencyScoreAndKeyList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairFrequencyScoreAndKeyList) Less(i, j int) bool { return p[i].Value < p[j].Value }
