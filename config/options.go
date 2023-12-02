package config

type Options struct {
	Test bool
}

func NewDefaultOptions() Options {
	return Options{
		Test: false,
	}
}

func NewTestOptions() Options {
	return Options{
		Test: true,
	}
}
