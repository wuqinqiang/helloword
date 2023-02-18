package core

var Default = &Options{
	spec:          "@every 30s",
	selectTimeout: 10, //unit of second
}

type Option func(options *Options)

type Options struct {
	spec          string
	selectTimeout int
}

func WithSpec(spec string) Option {
	return func(options *Options) {
		options.spec = spec
	}
}
