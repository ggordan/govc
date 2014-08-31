/** @jsx React.DOM */

var React = require('react');
var $ = require('jquery');
var moment = require('moment');

var Diff = React.createClass({

	// METHODS

	_determineClass: function(char) {
		switch (char) {
			case '+':
				return 'added';
			case '-':
				return 'removed';
			default:
				return 'basic';
		}
	},

	_getLines: function() {
		return this.props.data.Patch.split('\n').map(function(l, i) {
			return React.DOM.p({
				className: this._determineClass(l.charAt(0)),
			}, l);
		}.bind(this));
	},

	// RENDER

	render: function() {
		return(
			<div className="diff">
				{this._getLines()}
			</div>
		);
	},
});

module.exports = Diff;