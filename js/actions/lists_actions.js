
var Dispatcher = require('dispatcher');
var Constants  = require('flux/constants');

var ListsActions = {

  receiveUserLists: function(rawMessage) {
    Dispatcher.handleServerAction({
      type: Constants.ActionTypes.INITIALIZE_LISTS,
      rawMessage: rawMessage
    });
  },

};

module.exports = ListsActions;
