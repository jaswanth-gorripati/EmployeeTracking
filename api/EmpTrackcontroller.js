var express = require("express");
var router = express.Router();
var bodyParser = require("body-parser");

router.use(bodyParser.urlencoded({ extended: true }));
router.use(bodyParser.json());

const txSubmit = require("./invoke");
const txFetch = require("./query");


router.post("/addOrganisation", async function (req, res) {
  try {
    let result = await txSubmit.invoke("AddOrganisation", JSON.stringify(req.body));
    res.send(result);
  } catch (err) {
    res.status(500).send(err);
  }
});

router.post("/allOrganisation", async function (req, res) {
  try {
    let result = await txFetch.query("ListAllOrganisation", " ");
    let jsonResult = await JSON.parse(result)
    res.send(jsonResult);
  } catch (err) {
    res.status(500).send(err);
  }
});

router.post("/addEmployee", async function (req, res) {
  try {
    let result = await txSubmit.invoke("AddEmployee", JSON.stringify(req.body));
    res.send(result);
  } catch (err) {
    res.status(500).send(err);
  }
});

router.post("/getAllEmployees", async function (req, res) {
  //TFBC.getLC(req, res); req.body.lcId
  try {
    let result = await txFetch.query("ListAllEmployees", " ");
    res.send(JSON.parse(result));
  } catch (err) {
    res.status(500).send(err);
  }
});

router.post("/transferEmployee", async function (req, res) {
  try {
    let result = await txSubmit.invoke("TransferEmployee", JSON.stringify(req.body));
    res.send(result);
  } catch (err) {
    res.status(500).send(err);
  }
});

router.post("/employeeHistory", async function (req, res) {
  try {
    let result = await txFetch.query("GetEmployeeTransferHistory", req.body.empID);
    res.send(JSON.parse(result));
  } catch (err) {
    res.status(500).send(err);
  }
});

module.exports = router;
