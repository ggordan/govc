/** @jsx React.DOM */

var React = require('react');
var $ = require('jquery');

var Commit = require('pages/repo/commits/commit');

// --------------------------------

var Commits = React.createClass({

	getInitialState: function() {
		return {
			commits: [],
		};
	},

    // RENDER

    componentWillMount: function() {
    	$.ajax({
    		url: "/1/commits",
    		success: function(c) {
    			this.setState({
    				commits: JSON.parse(c)
    			});
    		}.bind(this)
    	});
    },

    render: function() {

    	var commits = this.state.commits.map(function(commit, index) {
    		return <Commit key={commit + index} data={commit} index={index} />;
    	});

        return(
            <div className="commits">
            	{commits}
            </div>
        );
    },
});

module.exports = Commits;
