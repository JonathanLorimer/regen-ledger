package rdf

import "fmt"

func GetOneTerm(it TermIterator) (Term, error) {
	defer it.Close()

	if !it.Next() {
		return nil, nil
	}

	val := it.Term()

	if it.Next() {
		return nil, fmt.Errorf("expected only one value")
	}

	return val, nil
}

type NodeBuilder struct {
	Builder GraphBuilder
	Node    IRIOrBNode
}

func NewNodeBuilder(builder GraphBuilder, props map[IRIOrBNode][]Term) NodeBuilder {
	bNode := builder.NewBNode()
	for p, objs := range props {
		for _, obj := range objs {
			builder.AddTriple(bNode, p, obj)
		}
	}
	return NodeBuilder{
		Builder: builder,
		Node:    bNode,
	}
}

func (builder *NodeBuilder) AddProps(props map[IRIOrBNode][]Term) {
	for p, objs := range props {
		for _, obj := range objs {
			builder.Builder.AddTriple(builder.Node, p, obj)
		}
	}
}

func SubjectIterator(iterator TripleIterator) TermIterator {
	return subIterator{iterator}
}

type subIterator struct{ TripleIterator }

func (s subIterator) Term() Term { return s.Subject() }

func PredicateIterator(iterator TripleIterator) TermIterator {
	return predIterator{iterator}
}

type predIterator struct{ TripleIterator }

func (s predIterator) Term() Term { return s.Predicate() }

func ObjectIterator(iterator TripleIterator) TermIterator {
	return objIterator{iterator}
}

type objIterator struct{ TripleIterator }

func (s objIterator) Term() Term { return s.Object() }

func Filter(iterator TripleIterator, filterFn func(sub IRIOrBNode, pred IRIOrBNode, obj Term) bool) TripleIterator {
	return filterIterator{TripleIterator: iterator, filterFn: filterFn}
}

type filterIterator struct {
	TripleIterator
	filterFn func(sub IRIOrBNode, pred IRIOrBNode, obj Term) bool
}

func (f filterIterator) Count() int {
	panic("implement me")
}

func (f filterIterator) CountGTE(i int) bool {
	panic("implement me")
}

func (f filterIterator) CountLTE(i int) bool {
	panic("implement me")
}

func (f filterIterator) Next() bool {
	for {
		if !f.TripleIterator.Next() {
			return false
		}

		if f.filterFn(f.Subject(), f.Predicate(), f.Object()) {
			return true
		}
	}
}

func AddTriple(builder GraphBuilder, triple Triple) {
	builder.AddTriple(triple.Subject, triple.Predicate, triple.Object)
}

func RemoveTriple(builder GraphBuilder, triple Triple) {
	builder.RemoveTriple(triple.Subject, triple.Predicate, triple.Object)
}

func HasTriple(builder GraphBuilder, triple Triple) bool {
	return builder.HasTriple(triple.Subject, triple.Predicate, triple.Object)
}

func Merge(builder GraphBuilder, graph Graph) {
	bnodeRemapping := map[BNode]BNode{}
	it := graph.FindAll()
	defer it.Close()

	for it.Next() {
		sub := it.Subject()
		if bnode, ok := sub.(BNode); ok {

		}
	}
}