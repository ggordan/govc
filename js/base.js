/** @jsx React.DOM */

var React = require('react/addons');

var Header = require('common/header');

var Base = React.createClass({
	render: function() {
		return(
			<div>
				<Header />
				{this.props.activeRouteHandler()}
			</div>
		);
	},
});

module.exports = Base;
