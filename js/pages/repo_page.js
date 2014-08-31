/** @jsx React.DOM */

var React = require('react');

// Scope all the components
var Status = require('pages/repo/status');
var Commits = require('pages/repo/commits');

var RepoPage = React.createClass({

    // RENDER

    render: function() {
        return(
            <div id="app" className="repo">
            	<div className="sidebar" />
            	<div className="repository">
	            	<Status />
	            	<Commits />
            	</div>
            </div>
        );
    },
});

module.exports = RepoPage;
