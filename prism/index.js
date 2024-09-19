const {
  getHttpOperationsFromSpec,
} = require("@stoplight/prism-cli/dist/operations");
const { createLogger } = require("@stoplight/prism-core");
const { createServer } = require("@stoplight/prism-http-server");
const { PassThrough } = require("stream");
const path = require("path");

let server;

exports.handler = async (event) => {
  if (!server) {
    const specPath = path.join(__dirname, "openapi.yaml");
    const operations = await getHttpOperationsFromSpec(specPath);
    server = createServer(operations, {
      components: {
        logger: createLogger(),
      },
      cors: true,
      config: {
        mock: { dynamic: true },
      },
    });
  }

  const { req, res } = createMockRequestResponse(event);

  await server.requestListener(req, res);

  return parseLambdaResponse(res);
};

function createMockRequestResponse(event) {
  const req = new PassThrough();
  req.headers = event.headers || {};
  req.method = event.requestContext.http.method;
  req.url =
    event.rawPath + (event.rawQueryString ? "?" + event.rawQueryString : "");

  if (event.body) {
    const isBase64Encoded = event.isBase64Encoded;
    const body = isBase64Encoded
      ? Buffer.from(event.body, "base64")
      : event.body;
    req.write(body);
  }
  req.end();

  const res = new PassThrough();
  res.headers = {};
  res.statusCode = 200;
  res.writeHead = (statusCode, headers) => {
    res.statusCode = statusCode;
    res.headers = { ...res.headers, ...headers };
  };

  let responseBody = "";
  res.on("data", (chunk) => {
    responseBody += chunk;
  });

  res.on("end", () => {
    res.body = responseBody;
  });

  return { req, res };
}

function parseLambdaResponse(res) {
  return {
    statusCode: res.statusCode,
    headers: res.headers,
    body: res.body,
  };
}
