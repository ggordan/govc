/** @jsx React.DOM */

var React = require('react');

// Scope all the components
var StatusList = require('pages/repo/status_list');
var Commits = require('pages/repo/commits');
var Branches = require('pages/repo/branches');

var RepoPage = React.createClass({

    // RENDER

    render: function() {
        return(
            <div id="app" className="repo">
            	<div className="sidebar">
                    <Branches />
                </div>
            	<div className="repository">
	            	<StatusList />
	            	<Commits />
            	</div>
            </div>
        );
    },
});

module.exports = RepoPage;
