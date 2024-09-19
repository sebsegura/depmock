const postmanToOpenApi = require("postman-to-openapi");
const postmanCollection = "../geo.json";
const outputFile = "./specs/geo.yaml";

async function convertCollection() {
  try {
    const result = await postmanToOpenApi(postmanCollection, outputFile, {
      default: "General",
    });
    console.log(`OpenAPI specs: ${result}`);
  } catch (err) {
    console.error(`Conversion failed:`, err);
  }
}

convertCollection();
