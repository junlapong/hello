from http.server import HTTPServer, BaseHTTPRequestHandler


class SimpleHTTPRequestHandler(BaseHTTPRequestHandler):

    def do_GET(self):
        self.send_response(200)
        self.end_headers()
        self.wfile.write(b'Hello, World!')


httpd = HTTPServer(('localhost', 8080), SimpleHTTPRequestHandler)

print("start http://localhost:8080")
httpd.serve_forever()

