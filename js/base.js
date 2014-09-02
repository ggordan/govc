/** @jsx React.DOM */

var React = require('react/addons');

var Base = React.createClass({
	render: function() {
		return(
			<div>
				{this.props.activeRouteHandler()}
			</div>
		);
	},
});

module.exports = Base;
