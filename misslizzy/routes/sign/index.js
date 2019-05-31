var express = require('express');
var router = express.Router();
const login = require('./login')
const joinus = require('./joinus')
const afterJoin = require('./afterJoin')
router.use('/login',login)
router.use('/joinus',joinus)
router.use('/afterJoin',afterJoin)


module.exports = router;
