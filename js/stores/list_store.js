var merge       = require('react/lib/merge');
var Events      = require('events');
var Dispatcher  = require('dispatcher');

var ListsActions    = require('actions/lists_actions');
var ActionTypes  = require('flux/constants').ActionTypes;

// MAIN -----------------------------------------------------------------------

var CHANGE_EVENT = 'lists change';

var _list = {};

var ListStore = merge(Events.EventEmitter.prototype, {

	// BASE

	emitChange: function() {
		this.emit(CHANGE_EVENT);
	},

	addChangeListener: function(callback) {
		this.on(CHANGE_EVENT, callback);
	},

	removeChangeListener: function(callback) {
		this.removeListener(CHANGE_EVENT, callback);
	},

	// SETTERS

	init: function(decks) {
		_lists = decks;
	},

	// GETTERS

	getAll: function () {
		return _lists;
	},

});

ListStore.dispatchToken = Dispatcher.register(function(payload) {
	var action = payload.action;
	switch(action.type) {
		case ActionTypes.INITIALIZE_LISTS:
			ListStore.init(sample);
			console.log('initialized');
      ListStore.emitChange();
			break;
	}
});

module.exports = ListStore;
