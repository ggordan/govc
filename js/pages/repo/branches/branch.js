/** @jsx React.DOM */

var React = require('react');
var $ = require('jquery');

// --------------------------------

var Branch = React.createClass({

    // EVENT HANDLERS

    _changeBranch: function(event) {
        $.ajax({
            url: "/api/"+ this.props.pid +"/branches/checkout?b=" + this.props.data.Name,
            success: function(c) {
                console.log('done');
            }.bind(this)
        });
    },

    // RENDER

    render: function() {

        if (!this.props.data.Remote) {
            console.log(this.props.data);
        }

        var head = (this.props.data.Head) ? <span className="head">(head)</span> : null;

        return(
            <li onClick={this._changeBranch}>
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
