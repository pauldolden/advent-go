package config

type Options struct {
	Test        bool
	SplitInputs bool
	TestPart    int
}

func NewDefaultOptions() Options {
	return Options{
		Test:        false,
		SplitInputs: false,
	}
}

func NewTestOptions() Options {
	return Options{
		Test:        true,
		SplitInputs: false,
	}
}
