var express = require('express');
var router = express.Router();

router.get('/', function(req, res){
	res.render('seatState');
});

router.post('/', function(req, res){
	// connection : 선택한 자리를 올려야함
	res.redirect('/afterLogin')
});


module.exports = router;