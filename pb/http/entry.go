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