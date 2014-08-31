/** @jsx React.DOM */

var React = require('react');
var Link = require('react-router').Link;

// Scope all the required components
var Search = require('common/header/search');

var Header = React.createClass({
    render: function() {
        return(
            <div id="header">
            	<h1 id="logo" className="tk-fertigo-script"><Link to="install">shoplist</Link></h1>
            	<Search />
        		<ul id="navigation">
		          <li><Link to="dashboard">Dashboard</Link></li>
		          <li><Link to="lists">Lists</Link></li>
		          <li><Link to="lists">Explore</Link></li>
        		</ul>
        		<div id="account"><img src="https://avatars0.githubusercontent.com/u/196237?v=2" height="50" /></div>
            </div>
        );
    },
});

module.exports = Header;
