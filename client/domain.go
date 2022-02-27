package godoo

type Domain []any

func (d *Domain) Apply() []any {
	var out []any
	for _, v := range *d {
		out = append(out, v)
	}
	return []any{out}
}

// And add Domain condition r to domain l in place
func (d *Domain) And(l Domain) *Domain {
	if d == nil {
		return &Domain{l}
	}

	if len(*d) == 0 {
		*d = append(*d, l)
		return d
	}

	d = &Domain{"&", (*d)[0]}
	*d = append(*d, l)
	return d
}

// Or add Domain condition r to domain l in place
func (d *Domain) Or(l Domain) *Domain {
	if d == nil {
		return &Domain{l}
	}

	if len(*d) == 0 {
		*d = append(*d, l)
		return d
	}

	d = &Domain{"|", (*d)[0]}
	*d = append(*d, l)
	return d
}
