const express = require('express');
const { getPosts, formatName, getData } = require('./post_loader');

const app = express();

const postFolder = "./public/posts";

const posts = getPosts(postFolder);

// app.use(express.static(__dirname + '/public'));
app.use(express.static('public'));

app.set('view engine', 'ejs');

app.get('/', (req, res) => {
  res.render('main.ejs', {
    title: "My site",
    posts,
    formatName
  });
});

for (const post of posts) {
  app.get('/post/' + post, (req, res) => {
    res.render('post.ejs', {
      title: formatName(post),
      data: getData(`${postFolder}/${post}`)
    });
  })
}

const port = 80;
app.listen(port);
console.log(`Working on port ${port}`);
