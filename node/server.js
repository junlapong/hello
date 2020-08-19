const http = require('http');

const requestListener = function (req, res) {
  res.writeHead(200);
  res.end("Hello, World!\n");
}

const server = http.createServer(requestListener);

console.log("start http://localhost:8080");
server.listen(8080);
