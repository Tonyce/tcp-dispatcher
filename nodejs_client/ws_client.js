const WebSocket = require('ws');
const { Duplex } = require('stream');

const ws = new WebSocket('ws://localhost:7654/ws');

ws.on('open', function open() {
  ws.send('something');
});

ws.on('message', function incoming(data) {
  console.log(data);
});

const inoutStream = new Duplex({
    write(chunk, encoding, callback) {
      this.push(chunk)
      callback();
    },
  
    read(size) {
    }
  });
inoutStream.on('data', (data) => {
    console.log({ data })
    ws.send(data)
})
  process.stdin.pipe(inoutStream);