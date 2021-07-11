const net = require('net')
const { Duplex } = require('stream');

const inoutStream = new Duplex({
    write(chunk, encoding, callback) {
      this.push(chunk)
      callback();
    },
  
    read(size) {
    }
  });

const socket = net.createConnection({
    port: 8765,
    host: "127.0.0.1"
}, () => {
    
});

socket.on('data', (data) => {
    console.log({data})
})

socket.on('connect', () => {
    console.log("connect...")
});

socket.on('end', () => {
    console.log("end....")
    process.exit()
});

socket.on('error', (e) => {
    console.log("e", e)
});

process.stdin.pipe(inoutStream).pipe(socket);