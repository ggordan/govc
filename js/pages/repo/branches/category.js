/** @jsx React.DOM */

var React = require('react');
var $ = require('jquery');

var Branch = require('pages/repo/branches/branch');

// --------------------------------

var Category = React.createClass({

    // RENDER

    render: function() {

        var data = this.props.data.map(function(item, index) {
            return <Branch data={item} key={index} pid={this.props.pid} />;
        }.bind(this));

        return(
            <div className={"branch category " + this.props.title}>
                <h4> {this.props.title} </h4>
                <ul>
            	   {data}
                </ul>
            </div>
        );
    },
});

module.exports = Category;
