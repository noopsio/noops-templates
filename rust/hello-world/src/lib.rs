wit_bindgen::generate!({
    world: "handler",
    exports: {
        world: TestHandler
    }
});

struct TestHandler;

impl Guest for TestHandler {
    fn handle(_: Request) -> Response {
        Response {
            status: 200,
            body: "Hello from Rust!\n".to_string(),
        }
    }
}
