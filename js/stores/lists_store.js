var merge       = require('react/lib/merge');
var Events      = require('events');
var Dispatcher  = require('dispatcher');

var ListsActions    = require('actions/lists_actions');
var ActionTypes  = require('flux/constants').ActionTypes;

// MAIN -----------------------------------------------------------------------

var CHANGE_EVENT = 'lists change';

// sample data
var sample = [
  {
    "_id": "53e146a79ae3a868c439e1bd",
    "index": 0,
    "guid": "099a1802-a3bd-4774-9aeb-482cdee5592c",
    "title": "House project",
    "budget": "$1,898.46"
  },{
    "_id": "53e146a79ae3a868c439e1bd",
    "index": 0,
    "guid": "099a1802-a3bd-4774-9aeb-482cdee5592c",
    "title": "House project",
    "budget": "$1,898.46"
  },{
    "_id": "53e146a79ae3a868c439e1bd",
    "index": 0,
    "guid": "099a1802-a3bd-4774-9aeb-482cdee5592c",
    "title": "House project",
    "budget": "$1,898.46"
  },
  {
    "_id": "53e146a7a4586a54d74eb190",
    "index": 1,
    "guid": "01cfa209-8494-4893-85ba-234d14ff2a81",
    "title": "Custom PC",
    "budget": "$2,761.51"
  },
  {
    "_id": "53e146a745fdba8a818dda47",
    "index": 2,
    "guid": "d70f3820-18b2-4ac1-887f-9a7b6401d6b1",
    "title": "Living room project",
    "budget": "$2,061.22"
  },
  {
    "_id": "53e146a7f0e4a12f9fcc7d2d",
    "index": 3,
    "guid": "93adb78a-89bb-433a-87b9-9fc58002f519",
    "title": "magna esse minim occaecat occaecat deserunt id",
    "budget": "$3,442.83"
  },
  {
    "_id": "53e146a7422d6273a3929ba4",
    "index": 4,
    "guid": "dd81837a-8dce-43bb-8cbf-71a8dae3eff6",
    "title": "veniam dolor enim qui elit voluptate anim",
    "budget": "$2,055.26"
  },
  {
    "_id": "53e146a778d10315e557e034",
    "index": 5,
    "guid": "6d0e2d5b-03c1-4941-af6a-99b615d91ee8",
    "title": "amet irure quis proident ut adipisicing ea",
    "budget": "$1,408.00"
  }
];


// Stores all the lists
var _lists = [];

var ListsStore = merge(Events.EventEmitter.prototype, {

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

ListsStore.dispatchToken = Dispatcher.register(function(payload) {
	var action = payload.action;
	switch(action.type) {
		case ActionTypes.INITIALIZE_LISTS:
			ListsStore.init(sample);
			console.log('initialized');
      ListsStore.emitChange();
			break;
	}
});

module.exports = ListsStore;
