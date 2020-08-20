#[macro_use]
extern crate nickel;

use nickel::{HttpRouter, Nickel};

fn main() {
    let mut server = Nickel::new();

    server.get("/", middleware!("Hello, World!"));
    server.get("/pong", middleware!("pong"));

    server.listen("0.0.0.0:3000");
}
