/** @jsx React.DOM */

var React = require('react');

// Scope all the components
var StatusList = require('pages/repo/status_list');
var Commits = require('pages/repo/commits');
var Branches = require('pages/repo/branches');

var RepoPage = React.createClass({

    // RENDER

    render: function() {

        console.log(this.props);

        return(
            <div id="app" className="repo">
            	<div className="sidebar">
                    <Branches pid={this.props.params.pid} />
                </div>
            	<div className="repository">
	            	<StatusList pid={this.props.params.pid} />
	            	<Commits pid={this.props.params.pid} />
            	</div>
            </div>
        );
    },
});

module.exports = RepoPage;
