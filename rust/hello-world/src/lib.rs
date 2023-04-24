wit_bindgen::generate!({
    world: "handler",
});

struct TestHandler;

impl Handler for TestHandler {
    fn handle(_: Request) -> Response {
        Response {
            status: 200,
            body: "Hello from Rust!\n".to_string(),
        }
    }
}

export_handler!(TestHandler);
