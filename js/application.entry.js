/** @jsx React.DOM */

var React = require('react');

var Routes = require('react-router').Routes;
var Route = require('react-router').Route;

// Require the base app
var App = require('base');

var InstallPage = require('pages/install_page');
var RepoPage = require('pages/repo_page');

	var io = require('socket.io-client');
  var socket = io('http://localhost:8090');
  socket.on('news', function (data) {
    console.log(data);
    console.log("news data");
    socket.emit('my other event', { my: 'data' });
  });

React.renderComponent((
  <Routes>
		<Route handler={App}>
			<Route name="repo" path="/repo" handler={RepoPage} />
	        <Route name="install" path="/install" handler={InstallPage} />
		</Route>
  </Routes>
), document.body);
