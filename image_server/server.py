import json
import tornado.ioloop
import tornado.web
from image_analyzer import get_image_from_text
from utils.filter import filter_by_list

class MainHandler(tornado.web.RequestHandler):
    def post(self):
        request_body = self.request.body
        request_data = json.loads(request_body)
        print(request_data)
        lines = get_image_from_text(request_data['path'])
        filtered_lines = filter_by_list(['you win', 'click', 'get it now'], lines)
        self.set_header("Content-Type", "application/json")
        print(filtered_lines)
        self.write(json.dumps(filtered_lines))

def make_app():
    return tornado.web.Application([
        (r"/image", MainHandler),
    ])

if __name__ == "__main__":
    app = make_app()
    app.listen(8080)
    print("Serving on port 8080")
    tornado.ioloop.IOLoop.current().start()
