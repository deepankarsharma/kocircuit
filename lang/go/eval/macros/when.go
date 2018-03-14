package macros

import (
	. "github.com/kocircuit/kocircuit/lang/circuit/eval"
	. "github.com/kocircuit/kocircuit/lang/circuit/model"
	. "github.com/kocircuit/kocircuit/lang/go/eval"
	. "github.com/kocircuit/kocircuit/lang/go/eval/symbol"
	. "github.com/kocircuit/kocircuit/lang/go/kit/util"
)

func init() {
	RegisterEvalMacro("When", new(EvalWhenMacro))
}

type EvalWhenMacro struct{}

func (m EvalWhenMacro) MacroID() string { return m.Help() }

func (m EvalWhenMacro) Label() string { return "when" }

func (m EvalWhenMacro) MacroSheathString() *string { return PtrString("When") }

func (m EvalWhenMacro) Help() string { return "When" }

// When(have:█, then:█, else:█)
func (EvalWhenMacro) Invoke(span *Span, arg Arg) (returns Return, effect Effect, err error) {
	a := arg.(*StructSymbol)
	have := a.Walk("have")
	if IsEmptySymbol(have) {
		els := a.Walk("else")
		if IsEmptySymbol(els) {
			return EmptySymbol{}, nil, nil
		} else {
			elsVty, ok := els.(*VarietySymbol)
			if !ok {
				return nil, nil, span.Errorf(nil, "when else is not a variety")
			}
			return elsVty.Invoke(span)
		}
	} else {
		then := a.Walk("then")
		if IsEmptySymbol(then) {
			return EmptySymbol{}, nil, nil
		} else {
			thenVty, ok := then.(*VarietySymbol)
			if !ok {
				return nil, nil, span.Errorf(nil, "when then is not a variety")
			}
			thenKnot := Knot{{Name: "", Shape: have, Effect: nil, Frame: span}}
			return thenVty.Evoke(span, thenKnot)
		}
	}
}
