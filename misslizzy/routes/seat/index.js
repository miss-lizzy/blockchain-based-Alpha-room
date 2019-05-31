var express = require('express');
var router = express.Router();
const afterLogin = require('./afterLogin')
const seatState = require('./seatState')

router.use('/afterLogin',afterLogin)
router.use('/seatState',seatState)

module.exports = router;