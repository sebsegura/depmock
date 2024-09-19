const path = require("path");
const CopyWebpackPlugin = require("copy-webpack-plugin");

module.exports = {
  target: "node",
  mode: "production",
  entry: "./index.js",
  node: {
    __dirname: false,
    __filename: false,
  },
  output: {
    filename: "index.js",
    path: path.resolve(__dirname, "dist"),
    libraryTarget: "commonjs2",
  },
  plugins: [
    new CopyWebpackPlugin({
      patterns: [{ from: "openapi.yaml", to: "." }],
    }),
  ],
};
