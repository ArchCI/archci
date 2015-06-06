import {ActionTypes} from "../AppConstants.js";
import ActionUtils from "./ActionUtils.js";

const API_PREFIX = "";

export default {

  receive_workers(params, done, fail) {
    ActionUtils.get(API_PREFIX + "/workers", params, data => {
      if (done) done(data);
    }, fail);
  },

  create_project(params, done, fail) {
    ActionUtils.post(API_PREFIX + "/projects", params, () => {
      ActionUtils.dispatch(ActionTypes.CREATE_PROJECT, params);
      if (done) done();
    }, fail);
  },

  receive_projects(params, done, fail) {
    ActionUtils.get(API_PREFIX + "/projects", params, data => {
      ActionUtils.dispatch(ActionTypes.RECEIVE_PROJECTS, data);
      if (done) done();
    }, fail);
  },

  receive_project_builds(params, done, fail) {
    ActionUtils.get(API_PREFIX + "/projects/" + params.id + "/builds", {}, data => {
      if (done) done(data);
    }, fail);
  },

  receive_builds(params, done, fail) {
    ActionUtils.get(API_PREFIX + "/builds", params, data => {
      if (done) done(data);
    }, fail);
  },

  receive_log(params, done, fail) {
    ActionUtils.get(API_PREFIX + `/builds/${params.id}/logs/${params.last}`, {}, data => {
      if (done) done(data);
    }, fail);
  }

};
