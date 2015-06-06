const $ = global.jQuery = require("jquery");
require("jquery-mockjax");

let _data;

export default (() => {
  $._ajax = $.ajax;
  $.ajax = (url, options) => {
    _data = options.data || {};
    return $._ajax(url, options);
  };
  $.mockjaxSettings.contentType = "application/json";

  $.mockjax({
    url: "/account",
    type: "POST",
    response() {
      this.responseText = {
        id: 1
      };
    }
  });

  $.mockjax({
    url: "/login",
    type: "POST",
    response() {
      this.responseText = {
        id: 1,
        username: _data.username,
        email: _data.username + "@gmail.com"
      };
    }
  });

  $.mockjax({
    url: "/login",
    type: "GET",
    response() {
      this.responseText = {
        id: 1,
        username: "wizawu",
        email: "wizawu@gmail.com"
      };
    }
  });

  $.mockjax({
    url: "/builds/active",
    type: "GET",
    response() {
      let result = [];
      let random = Math.random() * 10;
      for (let i = -1; i < random; i++) {
        result.push({
          id: Date.now() % 1000 + i + 1,
          worker: "192.168.0." + (i + 1),
          dispatched: new Date(Date.now() - Math.random() * 1000000).toJSON(),
          finished: null,
          status: null,
          commit: "f04cdd02",
          committer: "wizawu",
          branch: "master",
          type: "github",
          owner: "wizawu",
          repository: "archci" + (i + 1),
          description: "Say hello"
        });
      }
      this.responseText = result;
    }
  });

  $.mockjax({
    url: "/builds/search",
    type: "GET",
    response() {
      let result = [];
      let random = Math.random() * 10;
      for (let i = -1; i < random; i++) {
        let start = Date.now() - Math.random() * 1000000;
        let end = null, status = null;
        if (Math.random() > 0.5) {
          end = start + Math.random() * 1000000;
          status = Math.random() > 0.5 ? 0 : 1;
        }
        result.push({
          id: Date.now() % 1000 + i + 1,
          worker: "192.168.0." + (i + 1),
          dispatched: new Date(start).toJSON(),
          finished: new Date(end).toJSON(),
          status: status,
          commit: "f04cdd02",
          committer: "wizawu",
          branch: "master",
          type: "github",
          owner: "wizawu",
          repository: "archci" + (i + 1),
          description: "Say hello"
        });
      }
      this.responseText = result;
    }
  });

  $.mockjax({
    url: "/builds/*/log",
    type: "GET",
    response() {
      this.responseText = {
        log: "apt-get install -y golang"
      };
    }
  });

  $.mockjax({
    url: "/builds/*/log/*",
    type: "GET",
    response() {
      this.responseText = {
        log: "apt-get install -y golang " + Date.now(),
        last: Date.now()
      };
    }
  });

  $.mockjax({
    url: "/projects",
    type: "POST",
    status: 201
  });

  $.mockjax({
    url: "/projects",
    type: "GET",
    response() {
      let result = [];
      let random = Math.random() * 10;
      for (let i = 0; i < random; i++) {
        let type = "github", domain= "github.com";
        if (Math.random() > 0.5) {
          type = "gitlab";
          domain = "ci.curio.im";
        }
        result.push({
          id: i + 1,
          type: type,
          owner: "wizawu",
          repository: "archci" + (i + 1),
          description: "Say hello",
          domain: domain
        });
      }
      this.responseText = result;
    }
  });

  $.mockjax({
    url: "/images",
    type: "POST",
    status: 200
  });

  $.mockjax({
    url: "/images/*",
    type: "GET",
    response() {
      if (Math.random() < 0.1) {
        this.responseText = {
          id: Date.now() % 1000,
          name: "debian",
          size: Math.random() * 1000000000,
          dockerfile: "go version",
          created: new Date().toJSON(),
          url: "hub.archci.com/debian"
        };
      } else {
        this.responseText = {};
      }
    }
  });

  $.mockjax({
    url: "/images",
    type: "GET",
    response() {
      let random = Math.random() * 10;
      let result = [];
      for (let i = -1; i < random; i++) {
        result.push({
          id: i + 2,
          name: "debian",
          size: Math.random() < 0.3 ? -1 : Math.random() * 1000000000,
          dockerfile: "go version",
          created: new Date(Date.now() - Math.random() * 1000000).toJSON(),
          url: "hub.archci.com/debian"
        });
      }
      this.responseText = result;
    }
  });

  $.mockjax({
    url: "/workers",
    type: "GET",
    response() {
      let random = Math.random() * 10;
      let result = [];
      for (let i = -1; i < random; i++) {
        result.push({
          ip: "192.168.1." + (i + 2),
          busy: Math.random() > 0.5,
          contacted: new Date(Date.now() - Math.random() * 102400100).toJSON()
        });
      }
      this.responseText = result;
    }
  });

})();
