#!/bin/bash

# Dev dependencies
npm install --save-dev @babel/core
npm install --save-dev @babel/plugin-proposal-class-properties
npm install --save-dev @babel/plugin-proposal-export-default-from
npm install --save-dev @babel/polyfill
npm install --save-dev @babel/preset-env
npm install --save-dev @babel/preset-react
npm install --save-dev babel-eslint
npm install --save-dev babel-loader
npm install --save-dev chai
npm install --save-dev copy-webpack-plugin@6.0.3
npm install --save-dev error-overlay-webpack-plugin
npm install --save-dev eslint
npm install --save-dev eslint-loader
npm install --save-dev eslint-plugin-react
npm install --save-dev html-webpack-plugin
npm install --save-dev http-server
npm install --save-dev mocha
npm install --save-dev webpack
npm install --save-dev webpack-cli
npm install --save-dev webpack-dev-server
npm install --save-dev webpack-merge

# Depedencies
npm install --save @material-ui/core
npm install --save @material-ui/icons
npm install --save axios
npm install --save classnames
npm install --save clsx
npm install --save history
npm install --save prop-types
npm install --save react
npm install --save react-dom
npm install --save react-redux
npm install --save react-router-dom
npm install --save redux
npm install --save redux-logger
npm install --save redux-promise-middleware
npm install --save redux-thunk

npm audit fix 
npm audit
cat ./package.json