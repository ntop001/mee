package mee

// ParseTmpl parse template into args
// eg: "(int256,string,[](int256,int16),(uint256, string))"
func ParseTmpl(tmpl string) []*Arg {
	arg, _ := getTuple(tmpl)
	return arg.Meta
}

func getTuple(tmpl string) (*Arg, int) {
	arg := &Arg{ Type: "tuple" }
	sz := 0
	for i := 1; i < len(tmpl); i++ {
		if tmpl[i] == ' ' || tmpl[i] == ',' {
			continue
		}

		// tuple
		if tmpl[i] == '(' {
			a, j := getTuple(tmpl[i:])
			i += j
			arg.Meta = append(arg.Meta, a)
		}

		// array
		if tmpl[i] == '[' {
			a, j := getArray(tmpl[i:])
			i += j
			arg.Meta = append(arg.Meta, a)
		}

		// primitive
		if tmpl[i] != ' ' && tmpl[i] != ',' && tmpl[i] != ')' {
			a, j := getType(tmpl[i:])
			i += j
			arg.Meta = append(arg.Meta, a)
		}

		// end
		if tmpl[i] == ')' {
			sz = i+1; break
		}
	}
	return arg, sz
}

func getArray(tmpl string) (*Arg, int) {
	arg := &Arg{ Type: "array" }
	i := 2
	for ; i < len(tmpl); i++ {
		if tmpl[i] == '(' {
			a, j := getTuple(tmpl[i:])
			i += j
			arg.Meta = append(arg.Meta, a)
			break
		}
		if tmpl[i] == '[' {
			a, j := getArray(tmpl[i:])
			i += j
			arg.Meta = append(arg.Meta, a)
			break
		}
		if tmpl[i] != ' ' {
			a, j := getType(tmpl[i:])
			i += j
			arg.Meta = append(arg.Meta, a)
			break
		}
		if tmpl[i] == ',' || tmpl[i] == ')' {
			break
		}
	}
	return arg, i
}

func getType(tmpl string) (*Arg, int) {
	sz := 0
	for i := 0; i < len(tmpl); i++ {
		if tmpl[i] == ',' || tmpl[i] == ')' {
			sz = i; break
		}
	}
	arg := &Arg{
		Type: tmpl[:sz],
	}
	return arg, sz
}