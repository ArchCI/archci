/* All angular application controllers */
var archciControllers = angular.module("archciControllers", []);

archciControllers.controller("BuildsController", ["$scope", "$routeParams", "$http", "$interval",
  function($scope, $routeParams, $http, $interval) {

  /*
  [
    {
      "Id": 1,
      "ProjectName": "tobegit3hub/seagull",
      "Branch": "master",
      "Commit": "a34dbad42",
      "CommitTime": "2015-07-19T07:20:05+08:00",
      "Committer": "tobegit3hub",
      "Status": 0
    },
    {
      "Id": 2,
      "ProjectName": "ArchCI/archci",
      "Branch": "master",
      "Commit": "ba888d42",
      "CommitTime": "2015-07-19T07:21:12+08:00",
      "Committer": "tobegit3hub",
      "Status": 0
    },
    {
      "Id": 3,
      "ProjectName": "ArchCI/simple-worker",
      "Branch": "master",
      "Commit": "ffdbad42",
      "CommitTime": "2015-07-19T07:21:12+08:00",
      "Committer": "tobegit3hub",
      "Status": 0
    }
  ]
  */

  // If access /builds
  if(typeof $routeParams.buildId === 'undefined' || $routeParams.buildId == null){

    // TODO(tobe): check if the length is equal to 0
      $http.get("/v1/builds/all").success(function(data) {
        $scope.builds = data
        $scope.build = data[0]
      });

  }else{
  // If access /builds/:buildId

    $http.get("/v1/builds/all").success(function(data) {
      $scope.builds = data
    });

    // TODO(tobe): check if the id exists or not
    $http.get("/v1/builds/" + $routeParams.buildId).success(function(data) {
      $scope.build = data
    });

  };

  // TODO(tobe): we use "next" but not "Next" in beego api
  next = true;
  index = 0;
  $scope.fullLog = ""

  // More usage in https://docs.angularjs.org/api/ng/service/$interval
  get_log_loop = $interval(function() {

    if(next){
      /*
      {"log":"Unable to find image 'golang:1.4' locally","Next":true}
      */
      $http.get("/v1/builds/" + $scope.build.Id + "/logs/" + index).success(function(data) {

        $scope.fullLog = $scope.fullLog + data.log + "\n";

        next = data.Next;
        index++;
      });
    } else {
      $interval.cancel(get_log_loop);
    }

  }, 500);


  // Change the current build
  $scope.changeBuild = function(build) {
    $scope.build = build;

    // TODO(tobe): duplicated with above, should make it in a function
    next = true;
    index = 0;
    $scope.fullLog = ""

    // More usage in https://docs.angularjs.org/api/ng/service/$interval
    get_log_loop = $interval(function() {

      if(next){
        $http.get("/v1/builds/" + $scope.build.Id + "/logs/" + index).success(function(data) {

          $scope.fullLog = $scope.fullLog + data.log + "\n";

          next = data.Next;
          index++;
        });
      } else {
        $interval.cancel(get_log_loop);
      }

    }, 500);

  };

}]);


archciControllers.controller('ProjectsController', ['$scope', '$routeParams', '$http',
  function($scope, $routeParams, $http) {

  /*
  [
    {
      "Id": 1,
      "ProjectName": "tobegit3hub/seagull",
      "RepoUrl": "https://github.com/tobegit3hub/seagull",
      "Status": 0
    },
    {
      "Id": 2,
      "ProjectName": "tobegit3hub/note",
      "RepoUrl": "https://github.com/tobegit3hub/note",
      "Status": 0
    },
    {
      "Id": 3,
      "ProjectName": "ArchCI/archci",
      "RepoUrl": "https://github.com/ArchCI/archci",
      "Status": 0
    }
  ]
  */

  // If access /projects
  if(typeof $routeParams.projectId === 'undefined' || $routeParams.projectId == null){

    // TODO(tobe): check if the length is equal to 0
    $http.get("/v1/projects/all").success(function(data) {
      $scope.projects = data;
      $scope.project = data[0];
    });

  }else{
  // If access /projects/:projectId

    $http.get("/v1/projects/all").success(function(data) {
      $scope.projects = data;
    });

    // TODO(tobe): check if the id exists or not
    $http.get("/v1/projects/" + $routeParams.projectId).success(function(data) {
      $scope.project = data
    });
  };

  setTimeout(function(){
    //Sleep to get the data of current project

    /*
    [
      {
        "Id": 5,
        "ProjectName": "testproject",
        "RepoUrl": "",
        "Branch": "",
        "Commit": "",
        "CommitTime": "0001-01-01T08:00:00+08:00",
        "Committer": "",
        "BuildTime": "0001-01-01T08:00:00+08:00",
        "Status": 0
      }
    ]
    */

    $http.get("/v1/builds/all/project/" + $scope.project.ProjectName).success(function(data) {
      $scope.builds = data
    });

  }, 1000);

  // Change the current build
  $scope.changeProject = function(project) {
    $scope.project = project;

    $http.get("/v1/builds/all/project/" + $scope.project.ProjectName).success(function(data) {
      $scope.builds = data
    });

  }

  $scope.addProject = function(userName, projectName, repoUrl) {

    var data = {"UserName": userName,
                "ProjectName": projectName,
                "RepoUrl": repoUrl
    }

    $http.post("/v1/projects/new", data).success(function(data, status) {
        // TODO(tobe): add notification if success or fail
        $scope.add_status = "success";
        alert("Success to add project");
    })

    $http.post("/v1/builds/new", data).success(function(data, status) {
        // TODO(tobe): add notification if success or fail
        $scope.add_status = "success";
    })

  /*
  setTimeout(function(){
    $http.get("/v1/projects/all").success(function(data) {
      $scope.projects = data;
    });
  }, 500);
  */

  }

  // Post the data to record new build in database
  $scope.triggerCI = function(project) {

    /*
    var data = $.param({
        json: JSON.stringify({
            ProjectId: project.Id,
            ProjectName: project.ProjectName
        })
    });
    */

    var data = {"Id": project.Id,
                "ProjectName": project.ProjectName,
                "RepoUrl": project.RepoUrl
    }

    $http.post("/v1/builds/new", data).success(function(data, status) {
        // TODO(tobe): add notification if success or fail
        $scope.add_status = "success";
    })

    $http.get("/v1/builds/all/project/" + $scope.project.ProjectName).success(function(data) {
      $scope.builds = data;
    });
  }

}]);

archciControllers.controller('WorkersController', ['$scope', '$routeParams', '$http',
  function($scope, $routeParams, $http) {

  /*
  [
    {
      "Id": 1,
      "Ip": "127.0.0.1",
      "LastUpdate": "0001-01-01T08:00:00+08:00",
      "Status": 0
    },
    {
      "Id": 2,
      "Ip": "192.168.0.255",
      "LastUpdate": "0001-01-01T08:00:00+08:00",
      "Status": 0
    }
  ]
  */
  $http.get("/v1/workers/all/status/0").success(function(data) {
    $scope.busyWorkers = data;
  });

  $http.get("/v1/workers/all/status/1").success(function(data) {
    $scope.idleWorkers = data;
  });

  $http.get("/v1/workers/all/status/2").success(function(data) {
    $scope.dieWorkers = data;
  });
}]);
