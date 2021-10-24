var express = require('express');
var app = express();


var swaggerUi = require('swagger-ui-express');
var swaggerDocument = require('./swagger.json');


var EmpTrackcontroller = require('./EmpTrackcontroller');

app.use('/api-docs', swaggerUi.serve, swaggerUi.setup(swaggerDocument));
app.use('/emptrack', EmpTrackcontroller);

module.exports = app;
