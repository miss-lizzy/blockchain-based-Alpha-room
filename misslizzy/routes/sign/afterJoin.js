var express = require('express');
var router = express.Router();

router.get('/', function(req, res){
	//connection
	res.render('afterJoin');
});


module.exports = router;