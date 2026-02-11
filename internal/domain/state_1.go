package domain

type SessionState int

const (
	StateMain SessionState = iota
	StateOpts
)