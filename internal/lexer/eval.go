package lexer

import (
	"fmt"
	"sort"
)

// Iteration struct helps to pass to every iteration the tokens or partial
// results required to execute the iteration operator.
type Iteration[T any] struct {
	tokens [2]string
	values [2]T
}

// MockIteration function returns a mocked Iteration with just the token symbols
// provided.
func MockIteration[T any](a, b string, aT, bT T) *Iteration[T] {
	return &Iteration[T]{
		tokens: [2]string{a, b},
		values: [2]T{aT, bT},
	}
}

// A method return the tag and the current of the first component of the
// operation of the current iteration. It helps to implement Operator's
// providing an API to work with results of previous operations but olso with
// external data source associated with the tag.
func (iter Iteration[T]) A() (string, T) {
	return iter.tokens[0], iter.values[0]
}

// B method return the tag and the current of the second component of the
// operation of the current iteration. It helps to implement Operator's
// providing an API to work with results of previous operations but olso with
// external data source associated with the tag.
func (iter Iteration[T]) B() (string, T) {
	return iter.tokens[1], iter.values[1]
}

// Operator struct helps to define operations with their token tags and
// associated function that implements the type referenced.
type Operator[T any] struct {
	Tag string
	Fn  func(*Iteration[T]) (T, error)
}

// Evaluator struct helps to evaluate a Part using the defined operators, that
// implements the referenced type. It contains a parameter to store the results
// of each level operation, passing it to each iteration that requires it.
type Evaluator[T any] struct {
	operators      []*Operator[T]
	partialResults map[string]T
}

// NewEval function returns an initialized Evaluator with operators implementing
// the type referenced.
func NewEval[T any](ops []*Operator[T]) *Evaluator[T] {
	return &Evaluator[T]{
		operators:      ops,
		partialResults: make(map[string]T),
	}
}

// EvalToken methdo iterates over all the token provided childs executing the
// associated operator passing the tokens or values referenced by each token.
// The child tokens are sorted by descending level ensuring dependencies
// resolution and storing the partial results.
func (e *Evaluator[T]) EvalToken(p *Token) (T, error) {
	// get all childs and sort by level
	childs := p.AllGroups()
	sort.Slice(childs[:], func(i, j int) bool {
		return childs[i].Level > childs[j].Level
	})
	// create a var to store the last iteration result
	var lastResult T
	// iterate over childs
	for _, c := range childs {
		// parse child operator
		op := e.findOperator(c.Operator)
		if op == nil {
			return lastResult, fmt.Errorf("no supported operator '%s'", c.Operator)
		}
		// check the provided group
		if !c.Complete() {
			return lastResult, fmt.Errorf("wrong child group of tokens")
		}
		// call operator function passing the created iteration, if it returns
		// and error, stop the evaluation an return the error
		res, err := op.Fn(e.inflateIteration(c))
		if err != nil {
			return lastResult, fmt.Errorf("error during eval child with id '%d': %w", c.ID, err)
		}
		// if not, store the result as a partial result, uptate the last result
		// and continue
		e.partialResults[fmt.Sprint(c.ID)] = res
		lastResult = res
	}
	// return the last result, from the root part provided
	return lastResult, nil
}

func (e *Evaluator[T]) inflateIteration(c *Group) *Iteration[T] {
	// init a iteration data to inflate it
	iter := new(Iteration[T])
	// get child tokens and parse values, if exist a partial result for any
	// child token, include it, else include the child token
	if value, ok := e.partialResults[c.firstToken]; ok {
		iter.values[0] = value
	} else {
		iter.tokens[0] = c.firstToken
	}
	if value, ok := e.partialResults[c.secondToken]; ok {
		iter.values[1] = value
	} else {
		iter.tokens[1] = c.secondToken
	}
	return iter
}

// findOperator method finds a registered evaluator operator by the tag
// provided, if it is not found, it returns nil
func (e *Evaluator[T]) findOperator(tag string) *Operator[T] {
	for _, op := range e.operators {
		if op.Tag == tag {
			return op
		}
	}
	return nil
}
