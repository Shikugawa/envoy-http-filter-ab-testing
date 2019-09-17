const express = require('express');
const path = require('path');
const fallback = require('express-history-api-fallback');
const app = express();

app.use('/', express.static(path.resolve('dist')));
app.use(fallback('index.html', {
  root: path.resolve('dist')
}));

app.listen(3000);