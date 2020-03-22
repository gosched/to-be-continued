const express = require('express');
const bodyParser = require('body-parser');
var { graphql, buildSchema } = require('graphql');
const ejs = require('ejs');

var schema = buildSchema(`
  type Query {
    hello: String
  }
`);

let root = {
  hello: () => {
    return 'Hello world!';
  },
};

var app = express();
app.set('view engine', 'ejs');

app.use(bodyParser.urlencoded({ extended: false }))
app.use(bodyParser.json())

app.get('/', function (req, res) {
  res.render('index')
})

app.post('/graphql', (req, res) => {
  console.log('req.body', req.body);
  // graphql().then((result) => {
  //   res.send(JSON.stringify(result));
  // });
});

app.listen(8080);