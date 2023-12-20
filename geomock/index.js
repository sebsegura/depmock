const mockoon = require('@mockoon/serverless')
const mockEnv = require('./geomock.json')

const mockoonServerless = new mockoon.MockoonServerless(mockEnv)

exports.handler = mockoonServerless.awsHandler()