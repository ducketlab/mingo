package http

type EntryDecorator interface {

}

func NewEntrySet() *EntrySet {
	return &EntrySet{}
}

type EntrySet struct {
	Items []*Entry `json:"items"`
}

func (s *EntrySet) AddEntry(es ...Entry)  {
	for i := range es {
		s.Items = append(s.Items, &es[i])
	}
}

func (s *EntrySet) Merge(target *EntrySet) {
	for i := range target.Items {
		s.Items = append(s.Items, target.Items[i])
	}
}