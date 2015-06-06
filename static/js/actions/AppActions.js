import ActionUtils from "./ActionUtils.js";
import {ActionTypes} from "../AppConstants.js";

const API_PREFIX = "";

export default {

  receive_workers(params, done, fail) {
    ActionUtils.get(API_PREFIX + "/workers", params, data => {
      if (done) done(data);
    }, fail);
  }

};
