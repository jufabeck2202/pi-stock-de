package types

type Platform int

const (
	Undefined Platform = iota
	PushHover
	Mail
	Webhook
)
