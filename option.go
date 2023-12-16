package loghook

type Option struct {
	types string
}

func (o *Option) Types() string {
	return o.types
}

func NewOption(types string) *Option {
	return &Option{
		types: types,
	}
}
