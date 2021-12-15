package handler

type Item struct {
	Key   string
	Value int
}

type FormRequest struct {
	Paragraf string `form:"paragraf"`
}

type SortList []Item

func (s SortList) Len() int           { return len(s) }
func (s SortList) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s SortList) Less(i, j int) bool { return s[i].Value > s[j].Value }
