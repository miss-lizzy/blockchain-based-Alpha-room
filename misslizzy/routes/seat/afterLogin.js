var express = require('express');
var router = express.Router();

router.get('/', function(req, res){
	res.render('afterLogin');
});

router.post('/', function(req, res){
	console.log("post afterLogin")
	//res.send("post afterLogin")
	// 고유번호 비교해서 매칭되면 그 다음 페이지로 넘어가야지 ,, redirect 같은 거 써서
	res.redirect('/seatState')
});


module.exports = router;