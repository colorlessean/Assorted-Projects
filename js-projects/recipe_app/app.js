const express = require('express');
const path = require('path');

const app = express();

const port = 8080;

// create our recipe model
const testObj = {
    name: "Chicken Parm",
    servings: [2, 4],         // people
    time: [20, 40],           // minutes
    ingredients: [["Chicken Breast", "4 Whole"], ["Param", "300g"]],
    instructions: ["Do this", "then this", "finally this"],
    notes: "extra garbage to note like possibly remember blah blah or something else",
};


// express will serve up our static content to the /static path 
// app.use('/static', express.static(path.join(__dirname, 'public')));

// main page for the app serve up the index.html file 
app.get('/', (req, res) => {
    res.sendFile(path.join(__dirname, 'public/index.html'));
});

// retrieve recipe data
app.get('/recipe', (req, res) => {
    console.log("GET request to recipe endpoint");
    res.json(JSON.stringify(testObj));
});

// retrive specific recipe
app.get('/recipe/:id', (req, res) => {
    console.log(req.params.id);
    res.send(`<h1>Page for Recipe ${req.params.id}</h1>`);
});

// add a new recipe
app.post('/recipe', (req, res) => {
    console.log("POST request to recipe endpoint");
    res.send();
});

// remove a recipe we no longer want
app.delete('/recipe', (req, res) => {
    console.log("DELETE request to recipe endpoint");
    res.send();
});

// going to handle generation of shopping list  
app.post('/list', (req, res) => {
    console.log("POST request to list endpoint");
    res.send();
});

// server binding for the actual website
app.listen(port, () => {
    console.log(`app running on localhost:${port}`);
});

