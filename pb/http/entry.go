package http

import "fmt"

type EntryDecorator interface {

}

func (e *Entry) UniquePath() string {
	return fmt.Sprintf("%s:%s", e.Method, e.Path)
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