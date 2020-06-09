const fs = require('fs')

function formatName(name) {
  name = name.slice(0, name.indexOf(".")).replace(
    /[`~!@#$%^&*()_|+\-=?;:'",.<>\{\}\[\]\\\/]/gi, ' '
  );
  return name.charAt(0).toUpperCase() + name.slice(1);
}

function getPosts(path) {
  return fs.readdirSync(path);
}

function getData(path) {
  return fs.readFileSync(path);
}

module.exports = { getPosts, formatName, getData };

