package event

// constant event type identifiers
const (
	EventIncr = iota
	EventTiming
	EventAbsolute
	EventTotal
	EventGauge
	EventPrecisionTiming
)

// Event is an interface to a generic StatsD event, used by the buffered client collator
type Event interface {
	Stats() []string
	Type() int
	TypeString() string
	Payload() interface{}
	Update(e2 Event) error
	String() string
	Key() string
}
