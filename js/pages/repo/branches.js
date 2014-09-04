/** @jsx React.DOM */

var React = require('react');
var $ = require('jquery');

var Category = require('pages/repo/branches/category');

// --------------------------------

var Branches = React.createClass({

	getInitialState: function() {
		return {
			tags: [],
            local: [],
            remote: [],
		};
	},

    // RENDER

    componentWillMount: function() {
    	$.ajax({
            url: "/api/"+ this.props.pid +"/branches",
    		success: function(c) {
                var remote = [];
                var local = [];
                var tags =[];
                var data = JSON.parse(c);

                data.Branches.forEach(function(branch, index) {
                    if (branch.Tag) {
                        tags.push(branch);
                    } else {
                        if (branch.Remote) {
                            remote.push(branch);
                        } else {
                            local.push(branch);
                        }
                    }
                });

                this.setState({
                    tags: tags,
                    local: local,
                    remote: remote,
                });

    		}.bind(this)
    	});
    },

    render: function() {
        return(
            <div className="branches">
                <Category title="Local" data={this.state.local} pid={this.props.pid} />
                <Category title="Remote" data={this.state.remote} pid={this.props.pid} />
                <Category title="Tags" data={this.state.tags} pid={this.props.pid} />
            </div>
        );
    },
});

module.exports = Branches;
