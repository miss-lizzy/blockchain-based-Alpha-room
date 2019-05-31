var express = require('express');
var router = express.Router();
const sign = require('./sign/index') 
const seat = require('./seat/index')

/* GET home page. */
router.get('/', function(req, res, next) {
  res.render('main');
});

router.use('/',sign)
router.use('/',seat)

module.exports = router;
