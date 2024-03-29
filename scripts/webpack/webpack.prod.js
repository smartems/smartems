'use strict';

const merge = require('webpack-merge');
const TerserPlugin = require('terser-webpack-plugin');
const common = require('./webpack.common.js');
const path = require('path');
const ngAnnotatePlugin = require('ng-annotate-webpack-plugin');
const ForkTsCheckerWebpackPlugin = require('fork-ts-checker-webpack-plugin');
const HtmlWebpackPlugin = require("html-webpack-plugin");
const MiniCssExtractPlugin = require("mini-css-extract-plugin");
const OptimizeCSSAssetsPlugin = require("optimize-css-assets-webpack-plugin");

module.exports = merge(common, {
  mode: 'production',
  devtool: "source-map",

  entry: {
    dark: './public/sass/smartems.dark.scss',
    light: './public/sass/smartems.light.scss',
  },

  module: {
    rules: [{
        test: /\.tsx?$/,
        enforce: 'pre',
        exclude: /node_modules/,
        use: {
          loader: 'tslint-loader',
          options: {
            emitErrors: true,
            typeCheck: false,
          }
        }
      },
      {
        test: /\.tsx?$/,
        exclude: /node_modules/,
        use: {
          loader: 'ts-loader',
          options: {
            transpileOnly: true
          },
        },
      },
      require('./sass.rule.js')({
        sourceMap: false,
        preserveUrl: false
      })
    ]
  },
  optimization: {
    nodeEnv: 'production',
    minimizer: [
      new TerserPlugin({
        cache: false,
        parallel: true,
        sourceMap: true
      }),
      new OptimizeCSSAssetsPlugin({})
    ]
  },
  plugins: [
    new ForkTsCheckerWebpackPlugin({
      checkSyntacticErrors: true,
    }),
    new MiniCssExtractPlugin({
      filename: "smartems.[name].[hash].css"
    }),
    new ngAnnotatePlugin(),
    new HtmlWebpackPlugin({
      filename: path.resolve(__dirname, '../../public/views/error.html'),
      template: path.resolve(__dirname, '../../public/views/error-template.html'),
      inject: false,
      excludeChunks: ['dark', 'light'],
      chunksSortMode: 'none'
    }),
    new HtmlWebpackPlugin({
      filename: path.resolve(__dirname, '../../public/views/index.html'),
      template: path.resolve(__dirname, '../../public/views/index-template.html'),
      inject: 'body',
      excludeChunks: ['manifest', 'dark', 'light'],
      chunksSortMode: 'none'
    }),
    function () {
      this.hooks.done.tap('Done', function (stats) {
        if (stats.compilation.errors && stats.compilation.errors.length) {
          console.log(stats.compilation.errors);
          process.exit(1);
        }
      });
    }
  ]
});
