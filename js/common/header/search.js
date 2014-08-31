/** @jsx React.DOM */

var React = require('react');

var SearchBox = React.createClass({

    getDefaultProps: function() {
        return {
            placeholder: 'Search...',
        };
    },

    // EVENT HANDLERS

    _onFocus: function(event) {
        if (event.target.value === this.props.placeholder) {
            event.target.value = '';
        }
    },

    _onBlur: function(event) {
        if (event.target.value === "") {
            event.target.value = this.props.placeholder;
        }
    },

    // RENDER

    render: function() {

        return(
            <div className="search">
                <input type="text" className="searchBox" defaultValue={this.props.placeholder} onFocus={this._onFocus} onBlur={this._onBlur} />
            </div>
        );
    },
});

module.exports = SearchBox;
