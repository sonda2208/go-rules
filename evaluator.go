package rules

import (
	"errors"
	"fmt"
	"reflect"
	"time"
)

func Evaluate(expr Expr, params map[string]interface{}, funcs map[string]Function) (bool, error) {
	if expr == nil {
		return false, errors.New("invalid input expression")
	}

	res, err := evaluate(expr, params, funcs)
	if err != nil {
		return false, err
	}

	l, ok := res.(*boolLiteral)
	if !ok {
		return false, errors.New("invalid result expression")
	}

	return l.Val, nil
}

func evaluate(expr Expr, params map[string]interface{}, funcs map[string]Function) (Expr, error) {
	switch e := expr.(type) {
	case *parentExpr:
		return evaluate(e.Expr, params, funcs)
	case *binaryExpr:
		l, err := evaluate(e.LHS, params, funcs)
		if err != nil {
			return nil, err
		}

		r, err := evaluate(e.RHS, params, funcs)
		if err != nil {
			return nil, err
		}

		return compute(e.Op, l, r)
	case *identifier:
		varVal, ok := params[e.Val]
		if !ok {
			return nil, fmt.Errorf("parameter '%s' not found", e.Val)
		}

		return toLiteral(varVal)
	case *functionExpression:
		fn, ok := funcs[e.Name]
		if !ok {
			return nil, fmt.Errorf("function '%s' not found", e.Name)
		}

		var fnRes interface{}
		var fnErr error
		if e.Arguments != nil && len(e.Arguments) > 0 {
			args := make([]interface{}, len(e.Arguments))
			for i, arg := range e.Arguments {
				res, err := evaluate(arg, params, funcs)
				if err != nil {
					return nil, err
				}

				args[i] = res.(Literal).Value()
			}

			fnRes, fnErr = fn(args...)
		} else {
			fnRes, fnErr = fn()
		}

		if fnErr != nil {
			return nil, fmt.Errorf("error (%v) returned from function '%s'", fnErr, e.Name)
		}

		return toLiteral(fnRes)
	}

	return expr, nil
}

func compute(op Op, lhs, rhs Expr) (*boolLiteral, error) {
	switch op {
	case AND:
		return computeAND(lhs, rhs)
	case OR:
		return computeOR(lhs, rhs)
	case EQ:
		return computeEQ(lhs, rhs)
	case NEQ:
		return computeNEQ(lhs, rhs)
	case LT:
		return computeLT(lhs, rhs)
	case LTE:
		return computeLTE(lhs, rhs)
	case GT:
		return computeGT(lhs, rhs)
	case GTE:
		return computeGTE(lhs, rhs)
	case IN:
		return computeIN(lhs, rhs)
	}

	return nil, errors.New("invalid operator")
}

func computeOR(lhs, rhs Expr) (*boolLiteral, error) {
	switch l := lhs.(type) {
	case *boolLiteral:
		r, ok := rhs.(*boolLiteral)
		if ok {
			return &boolLiteral{Val: (l.Val || r.Val)}, nil
		}
	}

	return nil, errors.New("invalid operator")
}

func computeAND(lhs, rhs Expr) (*boolLiteral, error) {
	switch l := lhs.(type) {
	case *boolLiteral:
		r, ok := rhs.(*boolLiteral)
		if ok {
			return &boolLiteral{Val: (l.Val && r.Val)}, nil
		}
	}

	return nil, errors.New("invalid operator")
}

func computeEQ(lhs, rhs Expr) (*boolLiteral, error) {
	switch l := lhs.(type) {
	case *numberLiteral:
		r, ok := rhs.(*numberLiteral)
		if ok {
			return &boolLiteral{Val: (l.Val == r.Val)}, nil
		}
	case *stringLiteral:
		r, ok := rhs.(*stringLiteral)
		if ok {
			return &boolLiteral{Val: (l.Val == r.Val)}, nil
		}
	case *boolLiteral:
		r, ok := rhs.(*boolLiteral)
		if ok {
			return &boolLiteral{Val: (l.Val == r.Val)}, nil
		}
	case *timeLiteral:
		tr, ok := rhs.(*timeLiteral)
		if ok {
			return &boolLiteral{Val: (l.Val == tr.Val)}, nil
		}

		sr, ok := rhs.(*stringLiteral)
		if ok {
			dt, err := time.Parse(time.RFC3339, sr.Val)
			if err != nil {
				return nil, err
			}

			return &boolLiteral{Val: (l.Val == dt)}, nil
		}
	case *durationLiteral:
		strLit, ok := rhs.(*stringLiteral)
		if ok {
			v, err := time.ParseDuration(strLit.Val)
			if err != nil {
				return nil, err
			}

			return &boolLiteral{Val: l.Val == v}, nil
		}

		durLit, ok := rhs.(*durationLiteral)
		if ok {
			return &boolLiteral{Val: l.Val == durLit.Val}, nil
		}
	}

	return nil, fmt.Errorf(`cannot convert "%s" to %s`, rhs.String(), lhs.Type())
}

func computeNEQ(lhs, rhs Expr) (*boolLiteral, error) {
	switch l := lhs.(type) {
	case *numberLiteral:
		r, ok := rhs.(*numberLiteral)
		if ok {
			return &boolLiteral{Val: (l.Val != r.Val)}, nil
		}
	case *stringLiteral:
		r, ok := rhs.(*stringLiteral)
		if ok {
			return &boolLiteral{Val: (l.Val != r.Val)}, nil
		}
	case *boolLiteral:
		r, ok := rhs.(*boolLiteral)
		if ok {
			return &boolLiteral{Val: (l.Val != r.Val)}, nil
		}
	case *timeLiteral:
		tr, ok := rhs.(*timeLiteral)
		if ok {
			return &boolLiteral{Val: (l.Val != tr.Val)}, nil
		}

		sr, ok := rhs.(*stringLiteral)
		if ok {
			dt, err := time.Parse(time.RFC3339, sr.Val)
			if err != nil {
				return nil, err
			}

			return &boolLiteral{Val: (l.Val != dt)}, nil
		}
	case *durationLiteral:
		strLit, ok := rhs.(*stringLiteral)
		if ok {
			v, err := time.ParseDuration(strLit.Val)
			if err != nil {
				return nil, err
			}

			return &boolLiteral{Val: l.Val != v}, nil
		}

		durLit, ok := rhs.(*durationLiteral)
		if ok {
			return &boolLiteral{Val: l.Val != durLit.Val}, nil
		}
	}

	return nil, fmt.Errorf(`cannot convert "%s" to %s`, rhs.String(), lhs.Type())
}

func computeLT(lhs, rhs Expr) (*boolLiteral, error) {
	switch l := lhs.(type) {
	case *numberLiteral:
		r, ok := rhs.(*numberLiteral)
		if ok {
			return &boolLiteral{Val: (l.Val < r.Val)}, nil
		}
	case *stringLiteral:
		r, ok := rhs.(*stringLiteral)
		if ok {
			return &boolLiteral{Val: (l.Val < r.Val)}, nil
		}
	case *timeLiteral:
		tr, ok := rhs.(*timeLiteral)
		if ok {
			return &boolLiteral{Val: l.Val.Before(tr.Val)}, nil
		}

		sr, ok := rhs.(*stringLiteral)
		if ok {
			dt, err := time.Parse(time.RFC3339, sr.Val)
			if err != nil {
				return nil, err
			}

			return &boolLiteral{Val: l.Val.Before(dt)}, nil
		}
	case *durationLiteral:
		strLit, ok := rhs.(*stringLiteral)
		if ok {
			v, err := time.ParseDuration(strLit.Val)
			if err != nil {
				return nil, err
			}

			return &boolLiteral{Val: l.Val < v}, nil
		}

		durLit, ok := rhs.(*durationLiteral)
		if ok {
			return &boolLiteral{Val: l.Val < durLit.Val}, nil
		}
	}

	return nil, fmt.Errorf(`cannot convert "%s" to %s`, rhs.String(), lhs.Type())
}

func computeLTE(lhs, rhs Expr) (*boolLiteral, error) {
	switch l := lhs.(type) {
	case *numberLiteral:
		r, ok := rhs.(*numberLiteral)
		if ok {
			return &boolLiteral{Val: (l.Val <= r.Val)}, nil
		}
	case *stringLiteral:
		r, ok := rhs.(*stringLiteral)
		if ok {
			return &boolLiteral{Val: (l.Val <= r.Val)}, nil
		}
	case *timeLiteral:
		tr, ok := rhs.(*timeLiteral)
		if ok {
			return &boolLiteral{Val: l.Val.Before(tr.Val) || l.Val == tr.Val}, nil
		}

		sr, ok := rhs.(*stringLiteral)
		if ok {
			dt, err := time.Parse(time.RFC3339, sr.Val)
			if err != nil {
				return nil, err
			}

			return &boolLiteral{Val: l.Val.Before(dt) || l.Val == dt}, nil
		}
	case *durationLiteral:
		strLit, ok := rhs.(*stringLiteral)
		if ok {
			v, err := time.ParseDuration(strLit.Val)
			if err != nil {
				return nil, err
			}

			return &boolLiteral{Val: l.Val <= v}, nil
		}

		durLit, ok := rhs.(*durationLiteral)
		if ok {
			return &boolLiteral{Val: l.Val <= durLit.Val}, nil
		}
	}

	return nil, fmt.Errorf(`cannot convert "%s" to %s`, rhs.String(), lhs.Type())
}

func computeGT(lhs, rhs Expr) (*boolLiteral, error) {
	switch l := lhs.(type) {
	case *numberLiteral:
		r, ok := rhs.(*numberLiteral)
		if ok {
			return &boolLiteral{Val: (l.Val > r.Val)}, nil
		}
	case *stringLiteral:
		r, ok := rhs.(*stringLiteral)
		if ok {
			return &boolLiteral{Val: (l.Val > r.Val)}, nil
		}
	case *timeLiteral:
		tr, ok := rhs.(*timeLiteral)
		if ok {
			return &boolLiteral{Val: l.Val.After(tr.Val)}, nil
		}

		sr, ok := rhs.(*stringLiteral)
		if ok {
			dt, err := time.Parse(time.RFC3339, sr.Val)
			if err != nil {
				return nil, err
			}

			return &boolLiteral{Val: l.Val.After(dt)}, nil
		}
	case *durationLiteral:
		strLit, ok := rhs.(*stringLiteral)
		if ok {
			v, err := time.ParseDuration(strLit.Val)
			if err != nil {
				return nil, err
			}

			return &boolLiteral{Val: l.Val > v}, nil
		}

		durLit, ok := rhs.(*durationLiteral)
		if ok {
			return &boolLiteral{Val: l.Val > durLit.Val}, nil
		}
	}

	return nil, fmt.Errorf(`cannot convert "%s" to %s`, rhs.String(), lhs.Type())
}

func computeGTE(lhs, rhs Expr) (*boolLiteral, error) {
	switch l := lhs.(type) {
	case *numberLiteral:
		r, ok := rhs.(*numberLiteral)
		if ok {
			return &boolLiteral{Val: (l.Val >= r.Val)}, nil
		}
	case *stringLiteral:
		r, ok := rhs.(*stringLiteral)
		if ok {
			return &boolLiteral{Val: (l.Val >= r.Val)}, nil
		}
	case *timeLiteral:
		tr, ok := rhs.(*timeLiteral)
		if ok {
			return &boolLiteral{Val: l.Val.After(tr.Val) || l.Val == tr.Val}, nil
		}

		sr, ok := rhs.(*stringLiteral)
		if ok {
			dt, err := time.Parse(time.RFC3339, sr.Val)
			if err != nil {
				return nil, err
			}

			return &boolLiteral{Val: l.Val.After(dt) || l.Val == dt}, nil
		}
	case *durationLiteral:
		strLit, ok := rhs.(*stringLiteral)
		if ok {
			v, err := time.ParseDuration(strLit.Val)
			if err != nil {
				return nil, err
			}

			return &boolLiteral{Val: l.Val >= v}, nil
		}

		durLit, ok := rhs.(*durationLiteral)
		if ok {
			return &boolLiteral{Val: l.Val >= durLit.Val}, nil
		}
	}

	return nil, fmt.Errorf(`cannot convert "%s" to %s`, rhs.String(), lhs.Type())
}

func computeIN(lhs, rhs Expr) (*boolLiteral, error) {
	switch l := lhs.(type) {
	case *numberLiteral:
		r, ok := rhs.(*numberSliceLiteral)
		if ok {
			res := false
			for _, v := range r.Val {
				if l.Val == v {
					res = true
					break
				}
			}

			return &boolLiteral{Val: res}, nil
		}
	case *stringLiteral:
		r, ok := rhs.(*stringSliceLiteral)
		if ok {
			res := false
			for _, v := range r.Val {
				if l.Val == v {
					res = true
					break
				}
			}

			return &boolLiteral{Val: res}, nil
		}
	}

	return nil, fmt.Errorf(`cannot convert "%s" to %s`, rhs.String(), lhs.Type())
}

func toLiteral(i interface{}) (Expr, error) {
	varType := reflect.TypeOf(i)
	if varType == nil {
		return nil, errors.New("type not supported")
	}

	switch varType.Kind() {
	case reflect.Int:
		return &numberLiteral{Val: float64(i.(int))}, nil
	case reflect.Int32:
		return &numberLiteral{Val: float64(i.(int32))}, nil
	case reflect.Int64:
		dur, isDuration := i.(time.Duration)
		if isDuration {
			return &durationLiteral{Val: dur}, nil
		}

		return &numberLiteral{Val: float64(i.(int64))}, nil
	case reflect.Float32:
		return &numberLiteral{Val: float64(i.(float32))}, nil
	case reflect.Float64:
		return &numberLiteral{Val: float64(i.(float64))}, nil
	case reflect.String:
		return &stringLiteral{Val: i.(string)}, nil
	case reflect.Bool:
		return &boolLiteral{Val: i.(bool)}, nil
	case reflect.Struct:
		switch i.(type) {
		case time.Time:
			return &timeLiteral{i.(time.Time)}, nil
		}
	case reflect.Slice:
		switch i.(type) {
		case []string:
			l := &stringSliceLiteral{}
			for _, v := range i.([]string) {
				l.Val = append(l.Val, v)
			}

			return l, nil
		case []int:
			l := &numberSliceLiteral{}
			for _, v := range i.([]int) {
				l.Val = append(l.Val, float64(v))
			}

			return l, nil
		case []int32:
			l := &numberSliceLiteral{}
			for _, v := range i.([]int32) {
				l.Val = append(l.Val, float64(v))
			}

			return l, nil
		case []int64:
			l := &numberSliceLiteral{}
			for _, v := range i.([]int64) {
				l.Val = append(l.Val, float64(v))
			}

			return l, nil
		case []float32:
			l := &numberSliceLiteral{}
			for _, v := range i.([]float32) {
				l.Val = append(l.Val, float64(v))
			}

			return l, nil
		case []float64:
			l := &numberSliceLiteral{}
			for _, v := range i.([]float64) {
				l.Val = append(l.Val, v)
			}

			return l, nil
		case []interface{}:
			arr := i.([]interface{})
			if len(arr) > 0 {
				switch arr[0].(type) {
				case string:
					l := &stringSliceLiteral{}
					for _, v := range arr {
						tmp, ok := v.(string)
						if !ok {
							return nil, errors.New("unsupported literal")
						}

						l.Val = append(l.Val, tmp)
					}

					return l, nil
				case int, int32, int64, float32, float64:
					l := &numberSliceLiteral{}
					for _, v := range arr {
						switch i := v.(type) {
						case int:
							l.Val = append(l.Val, float64(i))
						case int32:
							l.Val = append(l.Val, float64(i))
						case int64:
							l.Val = append(l.Val, float64(i))
						case float32:
							l.Val = append(l.Val, float64(i))
						case float64:
							l.Val = append(l.Val, float64(i))
						default:
							return nil, errors.New("unsupported literal")
						}
					}

					return l, nil
				}
			}
		}
	}

	return nil, errors.New("unsupported literal")
}
