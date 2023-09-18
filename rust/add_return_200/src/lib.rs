wit_bindgen::generate!({
    world: "handler",
    exports: {
        world: Handler
    }
});

struct Handler;

impl Guest for Handler {
    fn handle(req: Request) -> Response {
        let mut response_body = String::default();
        let mut result = 0;

        for (_, value) in req.query_params {
            result += value.parse::<i32>().expect("NaN");
        }
        response_body = result.to_string();
        response_body.push_str("\n");
        Response {
            status: 200,
            body: response_body,
        }
    }
}
