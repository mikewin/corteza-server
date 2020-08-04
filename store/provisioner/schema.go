package provisioner

// Provisioner and

import (
	"fmt"
)

type (
	Tester   func(s *Provisioner) (bool, error)
	Executor func(s *Provisioner) error
	Printer  func(int, string)

	Provisioner struct {
		level int
		print Printer
	}
)

func NewProvisioner(p Printer) *Provisioner {
	return &Provisioner{
		level: -1,
		print: p,
	}
}

// Log prints with an indentation
func (s *Provisioner) Log(msg string, a ...interface{}) {
	if s.print != nil {
		s.print(s.level, fmt.Sprintf(msg, a...))
	}
}

func (s *Provisioner) Run(ee ...Executor) error {
	return s.run(ee...)
}

func (s *Provisioner) run(ee ...Executor) error {
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
	return func(s *Provisioner) error {
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
	return func(s *Provisioner) (bool, error) {
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
	return func(s *Provisioner) (bool, error) {
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

func Label(label string) Executor {
	return func(s *Provisioner) error {
		s.Log(label + "\n")
		return nil
	}
}

func Do(ee ...Executor) Executor {
	return func(s *Provisioner) error {
		s.level++
		defer func() { s.level-- }()
		return s.run(ee...)
	}
}

func Not(ee Tester) Tester {
	return func(s *Provisioner) (bool, error) {
		if r, err := ee(s); err != nil {
			return r, err
		} else {
			return !r, nil
		}
	}
}
