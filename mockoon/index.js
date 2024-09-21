const mockoon = require("@mockoon/serverless");
const mockEnv = require("./mockoon.json");

const mockoonServerless = new mockoon.MockoonServerless(mockEnv, {
  enableRandomLatency: true,
});

module.exports.handler = mockoonServerless.awsHandler();
