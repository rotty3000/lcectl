package flags

import "strconv"

type Output int

func (s *Output) Set(val string) error {
	*s = Output(strconv.Atoi(val))
	return nil
}
func (s *Output) Type() string {
	return "int8"
}

func (s *Output) String() string { return strconv.Itoa(*s) }

var Outputs = struct {
	Silent      Output
	Verbose     Output
	VeryVerbose Output
}{
	Silent:      0,
	Verbose:     1,
	VeryVerbose: 2,
}

func ToOutput(val int, p *Output) *Output {
	if val == int(Outputs.Verbose) {
		*p = Outputs.Verbose
	} else if val == int(Outputs.VeryVerbose) {
		*p = Outputs.VeryVerbose
	} else {
		*p = Outputs.Silent
	}
	return (*Output)(p)
}

var (
	ClientExtensionDir string
	ConfigFile         string
	Verbose            Output
)
