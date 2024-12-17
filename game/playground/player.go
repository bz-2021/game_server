package pg

type Player interface {
	Name() string

}

type WebSocketPlayer struct {
	name string
}