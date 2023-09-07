var fs = require("fs");
var stdinBuffer = fs.readFileSync(0);
console.log(stdinBuffer.toString());