from http.server import HTTPServer, BaseHTTPRequestHandler
import sys
PORT = 8080

cnt = 0
if sys.argv[1:]:
    PORT = int(sys.argv[1])

class MyHandler(BaseHTTPRequestHandler):
    def do_GET(self):
        global cnt
        cnt +=1
        if cnt == 10:
            exit(1)

        self.send_response(200)
        self.send_header("Content-type", "text/text")
        self.end_headers()
        self.wfile.write(f"Hello From {PORT}".encode())

server = HTTPServer(("localhost", PORT), MyHandler)

print(f"Serving on port {PORT}")
server.serve_forever()
