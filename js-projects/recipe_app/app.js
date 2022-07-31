const express = require('express');
const path = require('path');

const app = express();

const port = 8080;

// express will serve up our static content to the /static path 
// app.use('/static', express.static(path.join(__dirname, 'public')));

// main page for the app serve up the index.html file 
app.get('/', (req, res) => {
    res.sendFile(path.join(__dirname, 'public/index.html'));
});

// retrieve recipe data
app.get('/recipe', (req, res) => {
    console.log("GET request to recipe endpoint");
});

// add a new recipe
app.post('/recipe', (req, res) => {
    console.log("POST request to recipe endpoint");
});

// remove a recipe we no longer want
app.delete('/recipe', (req, res) => {
    console.log("DELETE request to recipe endpoint");
});

// going to handle generation of shopping list  
app.post('/list', (req, res) => {
    console.log("POST request to list endpoint");
});

// server binding for the actual website
app.listen(port, () => {
    console.log(`app running on localhost:${port}`);
});
