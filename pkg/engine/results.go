package engine

import (
	"github.com/korrel8r/korrel8r/pkg/korrel8r"
	"github.com/korrel8r/korrel8r/pkg/unique"
)

// Result accumulates queries and objects for a single class in a correlation chain.
type Result struct {
	Class   korrel8r.Class
	Queries unique.JSONList[korrel8r.Query]
	Rules   []korrel8r.Rule
	Objects korrel8r.Result
	Errors  unique.List[string]
}

func NewResult(class korrel8r.Class) Result {
	return Result{
		Class:   class,
		Objects: korrel8r.NewResult(class),
		Queries: unique.NewJSONList[korrel8r.Query](),
		Errors:  unique.NewList[string](),
	}
}

// Results is a sequence of results generated by following correlation paths.
type Results []Result

func (rs Results) find(class korrel8r.Class) (int, bool) {
	for i := range rs {
		if (rs)[i].Class == class {
			return i, true
		}
	}
	return len(rs), false
}

// Get a pointer to the Result for class, one is added if necessary.
func (rs *Results) Get(class korrel8r.Class) *Result {
	i, ok := rs.find(class)
	if !ok {
		*rs = append(*rs, NewResult(class))
	}
	return &(*rs)[i]
}

func (rs *Results) Merge(other Result) {
	self := rs.Get(other.Class)
	self.Queries.Append(other.Queries.List...)
	self.Rules = append(self.Rules, other.Rules...)
	self.Objects.Append(other.Objects.List()...)
	self.Errors.Append(other.Errors.List...)
}
