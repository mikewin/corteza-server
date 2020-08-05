package scenario

// Scenario and

import (
	"fmt"
)

type (
	Tester   func(s *Scenario) (bool, error)
	Executor func(s *Scenario) error
	Logger   func(int, string)

	Scenario struct {
		level int
		print Logger
	}
)

func NewScenario(p Logger) *Scenario {
	return &Scenario{
		level: -1,
		print: p,
	}
}

// Log prints with an indentation
func (s *Scenario) Log(msg string, a ...interface{}) {
	if s.print != nil {
		s.print(s.level, fmt.Sprintf(msg, a...))
	}
}

// Play runs provided
func (s *Scenario) Play(ee ...Executor) error {
	return s.play(ee...)
}

func (s *Scenario) play(ee ...Executor) error {
	for _, e := range ee {
		if e == nil {
			continue
		}

		if err := e(s); err != nil {
			return err
		}
	}

	return nil
}

// If is a simplified version of IfElse fn
// and executes onTrue if Tester passes
func If(v Tester, onTrue Executor) Executor {
	return IfElse(v, onTrue, nil)
}

// IfElse executes onTrue if Tester pases, otherwise it executes onFalse
func IfElse(v Tester, onTrue Executor, onFalse Executor) Executor {
	return func(s *Scenario) error {
		if ok, err := v(s); err != nil {
			return err
		} else if ok && onTrue != nil {
			return onTrue(s)
		} else if !ok && onFalse != nil {
			return onFalse(s)
		} else {
			return nil
		}
	}
}

// And returns verifier that returns true if all verifiers return true
func And(vv ...Tester) Tester {
	return func(s *Scenario) (bool, error) {
		for _, v := range vv {
			if v == nil {
				continue
			}

			if ok, err := v(s); err != nil || !ok {
				return false, err
			}
		}

		return true, nil
	}
}

// Or returns verifier that returns first true or error
func Or(vv ...Tester) Tester {
	return func(s *Scenario) (bool, error) {
		for _, v := range vv {
			if v == nil {
				continue
			}

			if ok, err := v(s); err != nil && ok {
				return ok, err
			}
		}

		return false, nil
	}
}

func Log(label string) Executor {
	return func(s *Scenario) error {
		s.Log(label + "\n")
		return nil
	}
}

func Do(ee ...Executor) Executor {
	return func(s *Scenario) error {
		s.level++
		defer func() { s.level-- }()
		return s.play(ee...)
	}
}

func Not(ee Tester) Tester {
	return func(s *Scenario) (bool, error) {
		if r, err := ee(s); err != nil {
			return r, err
		} else {
			return !r, nil
		}
	}
}
