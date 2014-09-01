/** @jsx React.DOM */

var React = require('react');
var $ = require('jquery');

// --------------------------------

var Branch = React.createClass({

    // RENDER

    render: function() {
        return(
            <li>
            <span className="icon icon-flow-branch" />
            <span className="title"> {this.props.data.Name} </span>
            </li>
        );
    },
});

module.exports = Branch;
