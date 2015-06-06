const assign = require("object-assign");
const EventEmitter = require("events").EventEmitter;

const CHANGE_EVENT = "change";

let BaseStore = assign({}, EventEmitter.prototype, {
  addChangeListener(callback) {
    this.on(CHANGE_EVENT, callback);
  },

  removeChangeListener(callback) {
    this.removeListener(CHANGE_EVENT, callback);
  },

  emitChange() {
    this.emit(CHANGE_EVENT);
  }
});

export default BaseStore;
