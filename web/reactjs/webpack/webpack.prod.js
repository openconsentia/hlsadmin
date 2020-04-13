const baseConfig = require('./webpack.base');
const merge = require('webpack-merge');

const productionConfig = {
  mode: 'production',
  output: {
    filename: '[name].bundle.js',
  },
  optimization: {
    splitChunks: {
      chunks: 'all',
    },
  },
}

module.exports = merge.smart(baseConfig, productionConfig);
