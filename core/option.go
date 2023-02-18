package core

var Default = &Options{
	spec: "@every 20s",
}

type Option func(options *Options)

type Options struct {
	spec string
}

func WithSpec(spec string) Option {
	return func(options *Options) {
		options.spec = spec
	}
}
