// This file is autogenerated by cmd/gen DO NOT EDIT
package flags

// Uint64Value flag type.
type Uint64Value struct {
	Value
	V *uint64
}

// Uint64 creates new Uint64 flag.
// Accepts a list of additional resolvers that are evaluated in sequence and
// the first one to yield a valid value is chosen.
// If no resolver yileds a valid value the default flag value is used.
// If flag is provided as a cli arg it will take precedance over all resolvers and the default value.
func (fs *FlagSet) Uint64(name, usage string, val uint64, r ...ResolverFunc) *uint64 {
	fs.initFlagSet()

	v := Uint64Value{
		Value: Value{
			name:      name,
			resolvers: r,
		},
		V: fs.fs.Uint64(name, val, usage),
	}

	fs.Values = append(fs.Values, v)

	return v.V
}

func (fs *FlagSet) parseUint64Vals() {
	for i, val := range fs.Values {
		uint64Val, ok := val.(Uint64Value)
		if !ok {
			continue
		}

		if fs.hasArg(uint64Val.name) {
			continue
		}

		for _, r := range uint64Val.resolvers {
			if r(fs, uint64Val.name, uint64(0), i) {
				break
			}
		}
	}
}