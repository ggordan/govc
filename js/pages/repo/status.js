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
