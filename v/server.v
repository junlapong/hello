module main

import vweb

const (
	port = 8080
)

struct App {
pub mut:
	vweb vweb.Context
}

pub fn (mut app App) init() {
}

pub fn (mut app App) init_once() {
}

pub fn (mut app App) index() {
	app.vweb.text("Hello, World!\n")
}

[get]
['/ping']
pub fn (mut app App) json() vweb.Result {
	app.vweb.set_content_type('application/json')
	return app.vweb.ok('{"message":"pong"}')
}

fn main() {
	println('Running on http://localhost:$port')
	vweb.run<App>(port)
}
