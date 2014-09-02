/** @jsx React.DOM */

var React = require('react');

var $ = require('jquery');


// --------------------------------

var Status = React.createClass({

	getInitialState: function() {
		return {
			status: [],
		};
	},

    // RENDER

    componentWillMount: function() {
        $.ajax({
            url: "/status",
            success: function(c) {
                this.setState({
                    status: JSON.parse(c)
                });
            }.bind(this)
        });
    },

    render: function() {
        var cx = React.addons.classSet;
        var files = this.state.status.map(function(file, index) {
            var fileClasses = {
                even: (index % 2)
            };
            fileClasses[file.Kind] = true;
           return <li className={cx(fileClasses)}>{file.Entry.IndexToWorkdir.NewFile.Path}</li>;
        });


        return(
            <div className="status">
                <div className="commitBox">
                    <input type="text" />
                    <input type="button" value="Commit" />
                </div>
                <div className="filesys">
                    <div className="files">
                        <ul>
                            {files}
                        </ul>
                    </div>
                    <div className="diff">
                        DIFF
                    </div>
                </div>
            </div>
        );
    },
});

module.exports = Status;
