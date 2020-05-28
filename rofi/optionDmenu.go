package rofi

import (
	"io"
	"strings"
)

// DmenuOption puts Rofi in dmenu mode, displaying stdin and outputting selected options to stdout.
type DmenuOption struct {
	choices []string
}

func (o *DmenuOption) ProvideArguments(existingArgs Arguments) (Arguments, error) {
	return NewArguments("-dmenu"), nil
}

func (o *DmenuOption) ProvideStdin() io.Reader {
	return strings.NewReader(strings.Join(o.choices, "\n"))
}

// WithDmenu puts rofi in Dmenu mode (via -dmenu), where stdin is presented to the user and selected items are output to stdout.
func (f *Factory) WithDmenu(choices ...string) *Factory {
	return f.withOption(&DmenuOption{choices})
}

// DDmenuPromptOption customises the prompt displayed
type DmenuPromptOption struct {
	prompt string
}

func (o *DmenuPromptOption) ProvideArguments(existingArgs Arguments) (Arguments, error) {
	return NewArguments("-prompt", o.prompt), nil
}

// WithDmenu puts rofi in Dmenu mode (via -dmenu), where stdin is presented to the user and selected items are output to stdout.
func (f *Factory) WithDmenuPrompt(prompt string) *Factory {
	return f.withOption(&DmenuPromptOption{prompt})
}

// DmenuOnlyMatch forces the user to select an existing item (i.e. not enter input other than the options shown)
type DmenuOnlyMatch struct {
}

func (o *DmenuOnlyMatch) ProvideArguments(existingArgs Arguments) (Arguments, error) {
	return NewArguments("-only-match"), nil
}

// WithDmenu puts rofi in Dmenu mode (via -dmenu), where stdin is presented to the user and selected items are output to stdout.
func (f *Factory) WithDmenuOnlyMatch() *Factory {
	return f.withOption(&DmenuOnlyMatch{})
}
