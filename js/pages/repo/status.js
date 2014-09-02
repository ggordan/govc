/** @jsx React.DOM */

var React = require('react');

// --------------------------------

var Status = React.createClass({

	getInitialState: function() {
		return {
			commits: [],
		};
	},

    // RENDER

    componentWillMount: function() {
        $.ajax({
            url: "/commits",
            success: function(c) {
                this.setState({
                    commits: JSON.parse(c)
                });
            }.bind(this)
        });
    },

    render: function() {
        return(
            <div className="status">
                <div className="commitBox">
                    <input type="text" />
                </div>
                <div className="filesys">
                    <div className="files">
                    </div>
                    <div className="diff">
                    </div>
                </div>
            </div>
        );
    },
});

module.exports = Status;
