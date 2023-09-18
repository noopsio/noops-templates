wit_bindgen::generate!({
    world: "handler",
    exports: {
        world: Handler
    }
});

struct Handler;

impl Guest for Handler {
    fn handle(_: Request) -> Response {
        Response {
            status: 200,
            body: "Hello from Rust!\n".to_string(),
        }
    }
}
