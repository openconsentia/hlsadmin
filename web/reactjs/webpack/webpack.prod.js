const { merge } = require('webpack-merge')

const commonConfig = require('./webpack.common')

const productionConfig = {
  mode: 'production'
}

module.exports = merge(commonConfig, productionConfig);
