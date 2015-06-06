import {ActionTypes} from "../AppConstants.js";
import AppDispatcher from "../AppDispatcher.js";
import BaseStore from "./BaseStore.js";

const assign = require("object-assign");

let _projects = {};

let AppStores = assign({}, BaseStore, {

  getProjects() {
    return _projects;
  },
  
  dispatcherIndex: AppDispatcher.register(payload => {
    let changed = true;
    let data = payload.data;

    switch (payload.type) {
      case ActionTypes.RECEIVE_PROJECTS:
        _projects = {};
        data.map(p => _projects[p.id] = p);
        break;
      case ActionTypes.CREATE_PROJECT:
        _projects[data.id] = data;
        break;
      default:
        changed = false;
        break;
    }

    if (changed) AppStores.emitChange();
  })

});

export default AppStores;
