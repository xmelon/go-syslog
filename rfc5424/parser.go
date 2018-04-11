package rfc5424

import (
	"sync"
)

// Parser represent FSM with mutual exclusion capabilities.
type Parser struct {
	sync.Mutex
	*machine
}

// NewParser creates a new parser and the underlying FSM.
func NewParser() *Parser {
	return &Parser{
		machine: NewMachine(),
	}
}

// Parse parses the input syslog message using its FSM.
func (p *Parser) Parse(input []byte) (*SyslogMessage, error) {
	p.Lock()
	defer p.Unlock()

	msg, err := p.machine.Parse(input)
	if err != nil {
		return nil, err
	}

	return msg, nil
}
