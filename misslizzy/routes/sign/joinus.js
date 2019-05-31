var express = require('express');
var router = express.Router();

router.get('/', function(req, res){
	res.render('joinus');
});

router.post('/', function(req, res){
	console.log("joinus post")
	// connection
	res.redirect("/afterJoin")
});

module.exports = router;