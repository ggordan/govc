/** @jsx React.DOM */

var React = require('react');
var $ = require('jquery');

// --------------------------------

var Branch = React.createClass({

    // RENDER

    render: function() {

        var head = (this.props.data.Head) ? <span className="head">(head)</span> : null;

        return(
            <li>
            <span className="icon icon-flow-branch" />
            <div className="title">
                <span>{this.props.data.Name}</span>
                {head}
            </div>
            </li>
        );
    },
});

module.exports = Branch;
