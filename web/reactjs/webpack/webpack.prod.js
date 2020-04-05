const webpack = require('webpack')
const merge = require('webpack-merge')

const baseConfig = require('./webpack.base')

const productionConfig = {
  mode: 'production'
}

module.exports = merge.smart(baseConfig, productionConfig);
