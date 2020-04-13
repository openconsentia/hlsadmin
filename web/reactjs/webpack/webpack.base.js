const path = require('path');
const HtmlWebpackPlugin = require('html-webpack-plugin');
const CopyPlugin = require('copy-webpack-plugin');

const SRC_DIR = path.resolve(__dirname, '..', 'src');
const BUILD_DIR = path.resolve(__dirname, '..', 'public');

module.exports = {
  entry: [`@babel/polyfill`,`${SRC_DIR}/index.js`],
  output: {
    path: BUILD_DIR,
    publicPath: '/',
    filename: 'bundle.js',
  },
  module: {
    rules: [
      {
        test: /\.js$/,
        exclude: /node_modules/,
        use: 'babel-loader',
      },
    ],
  },
  plugins: [
    new HtmlWebpackPlugin({ 
      template: path.resolve(SRC_DIR,'index.html')
    }),
    new CopyPlugin([
      { from: path.resolve(__dirname,'..','images'), to: path.resolve(BUILD_DIR,'images') },
    ]),
  ]
}
