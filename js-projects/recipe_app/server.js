const http = require("http");
const fs = require("fs").promises;

const host = 'localhost';
const port = 8080;

// going to store the content of our html file that is getting served
let indexFile;

// function will populate the res with the HTML file contents
const requestListener = function (req, res) {
    res.setHeader("Content-Type", "text/html");
    switch (req.url) {
        case "/":    
            res.writeHead(200);
            res.write(indexFile);
            res.end();
            break;
        default:
            res.writeHead(404);
            res.write("<h1>Page Not Found</h1>");
            res.end();
    }


};

// make the instance of the server itself using the requestListener to handle the incoming requests
// createServer takes in a handler to deal with the req/res structure and sets up a server instance
const server = http.createServer(requestListener);

// readFile generates a promise which will be resolved at a different time
// this is asyncronous code and when the promise resolves the then call is made
// if there is a failure we then enter into the catch
// in both cases the function will have a contents or err value that we can then use
fs.readFile(__dirname + "/client/index.html")
    .then(contents => {
        indexFile = contents;
        // listen binds the server to host/port and has a callback function that runs after it is bound
        server.listen(port, host, () => {
            console.log(`Server is running on http://${host}:${port}`);
        });
    })
    .catch(err => {
        console.error(`Could not read index.html file: ${err}`);
        process.exit(1);
    });


