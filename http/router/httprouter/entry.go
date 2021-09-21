package httprouter

import (
	"fmt"
	httppb "github.com/ducketlab/mingo/pb/http"
	"net/http"
)

type entry struct {
	*httppb.Entry
	h http.HandlerFunc
}

func newEntrySet() *entrySet {
	return &entrySet{
		items: map[string]*entry{},
	}
}

type entrySet struct {
	order []string
	items map[string]*entry
}

func (s *entrySet) AddEntry(es ...*entry) error {
	for _, e := range es {
		key := s.ID(e.Path, e.Method)
		if _, ok := s.items[key]; ok {
			return fmt.Errorf("router entry %s has exist", key)
		}
		s.items[key] = e
		s.order = append(s.order, key)
	}
	return nil
}

func (s *entrySet) FindEntry(path string, method string) *entry {
	id := s.ID(path, method)

	entry, ok := s.items[id]

	if ok {
		return entry
	}

	return nil
}

func (s *entrySet) ID(path string, method string) string {
	return path + "." + method
}
