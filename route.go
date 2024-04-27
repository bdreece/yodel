package yodel

// Route provides the parameters needed to register a given Handler
// in the Router.
type Route struct {
    Method string
    Path string
    Handler Handler
    Middleware []Middleware
}
