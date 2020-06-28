const https = require("https");
const fs = require("fs");
const express = require('express');
const { getPosts, formatName, getData, getFills } = require('./loader');

const app = express();

const postFolder = "./public/posts";

const posts = getPosts(postFolder);

const fills = getFills('./data.json');

// app.use(express.static(__dirname + '/public'));
app.use(express.static('public'));

app.set('view engine', 'ejs');

app.get('/', (_, res) => {
  res.render('main.ejs', {
    title: "My site",
    posts,
    formatName,
    fills
  });
});

for (const post of posts) {
  app.get('/post/' + post, (_, res) => {
    res.render('post.ejs', {
      title: formatName(post),
      data: getData(`${postFolder}/${post}`)
    });
  })
}

const c = {
  key: fs.readFileSync(process.env.P_KEYP),
  cert: fs.readFileSync(process.env.F_CHAINP),
};

const port = 443;
https.createServer(c, app).listen(port);

console.log(`Working on port ${port}`);
