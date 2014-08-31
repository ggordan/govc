// MODULES

var path        = require('path');
var crypto      = require('crypto');

var gulp        = require('gulp');
var webpack     = require('webpack');
var _           = require('underscore');

// PLUGINS
var refHash     = require('gulp-ref-hash');
var concat      = require('gulp-concat');
var uglify      = require('gulp-uglify');
var filter      = require('gulp-filter');
var minifyCSS   = require('gulp-minify-css');
var jshint      = require('gulp-jshint');
var clean       = require('gulp-clean');
var sass        = require('gulp-sass');
var prefix      = require('gulp-autoprefixer');
var ignore      = require('gulp-ignore');
var react       = require('gulp-react');
var tap         = require('gulp-tap');
var gulpIf      = require('gulp-if');
var revReplace  = require('gulp-rev-replace');
var useref      = require('gulp-useref');
var livereload  = require('gulp-livereload');

var webpackConf = require('./webpack.config');

// Define directories here
var devBuild  = path.join(__dirname, 'built', 'dev');
var devStatic = path.join(devBuild, 'assets');
var prodBuild = path.join(__dirname, 'built', 'prod');

gulp.task('index:dev', function() {
	return gulp.src(path.join(__dirname, 'index.html'))
		.pipe(gulp.dest(path.join(devBuild)));
});

gulp.task('sass:dev', function() {
	return gulp.src('sass/**/*.scss')
		.pipe(sass())
		.pipe(prefix('last 5 version'))
		.pipe(gulp.dest(path.join(devStatic, 'css')));
});

gulp.task('img:dev', function() {
	return gulp.src(['img/**/*'])
		.pipe(gulp.dest(path.join(devStatic, 'img')));
});

gulp.task('js:dev', function() {
	return gulp.src('js/**/*.js')
		.pipe(react())
		.pipe(jshint())
		.pipe(jshint.reporter('default'))
		.pipe(gulp.dest(path.join(devStatic, 'js')));
});

gulp.task('webpack', ['js:dev'], function(callback) {
	webpack(_.extend(webpackConf, {
		cache: true,
		watch: true,
		devtool: ['source-map'],
	})).run(function(err, stats) {
		if (err) console.log(err);
		callback();
	});
});

gulp.task('dev', ['sass:dev', 'index:dev', 'img:dev', 'webpack']);

gulp.task('watch', function() {
	livereload.listen();
	webpack(_.extend(webpackConf, {
		cache: true,
		watch: true,
		devtool: ['source-map'],
	})).watch(200, function(err, stats) {
		if (err) console.log(err);
		console.log('Reloading webpack...');
	});

	gulp.watch(['js/**/*.js'],['js:dev', livereload.changed]);
	gulp.watch(['sass/**/*.scss'],['sass:dev', livereload.changed]);
	gulp.watch(['index.html'],['index:dev', livereload.changed]);
});

gulp.task('default', ['dev', 'watch']);