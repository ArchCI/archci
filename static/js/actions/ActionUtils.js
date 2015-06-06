"use strict";

import AppDispatcher from "../AppDispatcher.js";
import $ from "jquery";

function _ajax(method, url, params, done, fail) {
  let options = {
    type: method,
    data: params,
    dataType: "json"
  };
  return $.ajax(url, options).done(data => {
    if (done) done(data);
  }).fail(xhr => {
    if (fail) fail(xhr);
  });
}

export default {
  dispatch(type, data) {
    AppDispatcher.dispatch({
      type: type,
      data: data
    });
	},

  get(url, params, done, fail) {
    return _ajax("GET", url, params, done, fail);
  },

  post(url, params, done, fail) {
    return _ajax("POST", url, params, done, fail);
  },

  put(url, params, done, fail) {
    return _ajax("PUT", url, params, done, fail);
  },

  delete(url, params, done, fail) {
    return _ajax("DELETE", url, params, done, fail);
  }
};
