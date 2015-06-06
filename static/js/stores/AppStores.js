import {ActionTypes} from "../AppConstants.js";
import AppDispatcher from "../AppDispatcher.js";
import BaseStore from "./BaseStore.js";

const assign = require("object-assign");

let AppStores = assign({}, BaseStore, {
  
  dispatcherIndex: AppDispatcher.register(payload => {
    let changed = true;
    let data = payload.data;

    switch (payload.type) {
      case ActionTypes.NULL:
        data = {};
        break;
      default:
        changed = false;
        break;
    }

    if (changed) AppStores.emitChange();
  })

});

export default AppStores;
