package selector

type Option func(srv *Srv)

func WithWordNumber(wordNumer int) Option {
	return func(srv *Srv) {
		srv.wordNumber = wordNumer
	}
}
