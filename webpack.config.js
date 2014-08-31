var _       = require('underscore');
var webpack = require('webpack');
var path    = require('path');
var glob    = require('glob');

var entries = {};
// Automatically create webpack entry points based on the url
var entryPaths = path.join(__dirname, 'js', '**') + '/*.' + 'entry.js';
glob.sync(entryPaths).forEach(function (file) {
    var location =  file.replace(path.join(__dirname, 'js'), '');
    var name = location.substring(1).replace('.js', '');
    entries[name] = '.' + location;
});

// Base webpack configuration
module.exports = {
    target: "web",
    context: path.join(__dirname, 'js'),
    entry: entries,
    output: {
        path: './built/dev/assets/bundles/',
        filename: "[name].js",
        chunkFilename: "[id].js",
        sourceMapFilename: '[file].map',
    },
    resolve : {
        root: __dirname,
        modulesDirectories: [
            "bower_components",
            "node_modules",
            path.join(__dirname, 'js'),
        ],
    },
    module: {
        loaders: [
            { test: /\.js$/, loader: 'jsx-loader' },
        ],
    },
    plugins: [
        new webpack.ResolverPlugin([
            new webpack.ResolverPlugin.DirectoryDescriptionFilePlugin("bower.json", ["main"]),
            new webpack.ResolverPlugin.DirectoryDescriptionFilePlugin(".bower.json", ["main"]),
            new webpack.ResolverPlugin.DirectoryDescriptionFilePlugin("package.json", ["main"])
        ])
    ]
};
