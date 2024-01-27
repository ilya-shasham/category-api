package helpers

type Fault uint8

const (
	ServerFault Fault = iota
	ClientFault
	CoreFault
	UnknownFault
)

func (f Fault) Int() int {
	switch f {
	case ServerFault:
		return 500
	case ClientFault:
		return 400
	case CoreFault:
		return 424
	default:
		return 418
	}
}

func (f Fault) String() string {
	switch f {
	case ServerFault:
		return "server fault"
	case ClientFault:
		return "client fault"
	case CoreFault:
		return "core fault"
	default:
		return "unknown fault"
	}
}
