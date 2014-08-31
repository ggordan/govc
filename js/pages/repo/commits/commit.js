/** @jsx React.DOM */

var React = require('react/addons');
var $ = require('jquery');
var moment = require('moment');
var Diff = require('pages/repo/commits/diff');

var Commit = React.createClass({

	getInitialState: function() {
		return {
			diffVisible: false,
		};
	},

	// METHODS

	_getGravatarURL: function() {
		return "//www.gravatar.com/avatar/" + this.props.data.Author.MD5 + "?s=50&d=mm";
	},

	_formatDate: function() {
		return moment(Date.parse(this.props.data.Timestamp)).fromNow();
	},

	// EVENT HANDLERS

	_toggleDiff: function(event) {
		this.setState({
			diffVisible: !this.state.diffVisible
		});
	},

	// RENDER

	render: function() {

		var diffContainerClasses = React.addons.classSet({
			diffContainer: true,
			visible: this.state.diffVisible,
		});

		var diffs = this.props.data.Diff.FilesChanged.map(function(d, i) {
			return <Diff data={d} />;
		});

		return(
			<div className="commit">
				<div className="meta">
					<img src={this._getGravatarURL()} height="50" width="50" />
					<div className="data">
						<h3> {this.props.data.Message} </h3>
						<p className="sha">SHA1: {this.props.data.SHA1}</p>
						<p>
							by <a href={"mailto:" + this.props.data.Author.Email}> {this.props.data.Author.Name} </a>
							authored <span title={this.props.data.Timestamp}>{this._formatDate()}</span>
						</p>
					</div>
				</div>
				<span className="diffToggle" onClick={this._toggleDiff}> Show diff </span>
				<div className={diffContainerClasses}>
					{diffs}
				</div>
			</div>
		);
	},
});

module.exports = Commit;