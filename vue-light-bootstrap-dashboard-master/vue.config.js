const path = require('path');
const webpack = require('webpack');

function resolveSrc(_path) {
  return path.join(__dirname, _path);
}

module.exports = {
  lintOnSave: false,
  configureWebpack: {
    // Set up all the aliases we use in our app.
    resolve: {
      alias: {
        src: resolveSrc('src'),
        'chart.js': 'chart.js/dist/Chart.js'
      }
    },
    plugins: [
      new webpack.LoaderOptionsPlugin({
        // test: /\.xxx$/, // may apply this only for some modules
        options: {
          loaders: [
            {
              test: /\.vue$/,
              loader: 'vue'
            },
            {
              test: /\.s[a|c]ss$/,
              loader: 'style!css!sass'
            }
          ]
      }
      }),

new webpack.optimize.LimitChunkCountPlugin({
        maxChunks: 6
      })
    ],

    // loaders: [
    //   {
    //     test: /\.vue$/,
    //     loader: 'vue'
    //   },
    //   {
    //     test: /\.s[a|c]ss$/,
    //     loader: 'style!css!sass'
    //   }
    // ]
  },
  // vue: {
  //   loaders: {
  //     scss: 'style!css!sass'
  //   }
  // },
  pwa: {
    name: 'Vue Light Bootstrap Dashboard',
    themeColor: '#344675',
    msTileColor: '#344675',
    appleMobileWebAppCapable: 'yes',
    appleMobileWebAppStatusBarStyle: '#344675'
  },
  css: {
    // Enable CSS source maps.
    sourceMap: process.env.NODE_ENV !== 'production'
  },

};
