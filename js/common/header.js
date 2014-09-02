/** @jsx React.DOM */

var React = require('react');
var Link = require('react-router').Link;

// Scope all the required components
var Search = require('common/header/search');

var Header = React.createClass({
    render: function() {
        return(
            <div id="header">
            </div>
        );
    },
});

module.exports = Header;
