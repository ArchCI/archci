/* All angular application controllers */
var archciControllers = angular.module("archciControllers", []);

archciControllers.controller("BuildsController", ["$scope", "$routeParams", "$http", "$interval",
  function($scope, $routeParams, $http, $interval) {

    /*
    {"log":"Unable to find image 'golang:1.4' locally","Next":true}
    */
    $http.get("/v1/builds/123/logs/0").success(function(data) {
      $scope.data = data;

      $scope.fullLog = data.log;

      // TODO(tobe): we use "next" but not "Next" in beego api
      next = data.Next;

      index = 1;

      // Refer to https://docs.angularjs.org/api/ng/service/$interval
      get_log_loop = $interval(function() {

        if(next){
          $http.get("/v1/builds/123/logs/" + index).success(function(data) {

            $scope.fullLog = $scope.fullLog + "\n" + data.log;

            next = data.Next;
            index++;
          });
        } else {
          $interval.cancel(get_log_loop);
        }

        console.log("continue to interval")

      }, 500);

    });

}]);


archciControllers.controller('ProjectsController', ['$scope', '$routeParams', '$http',
  function($scope, $routeParams, $http) {

  $scope.Name = "ArchCI"
}]);

archciControllers.controller('WorkersController', ['$scope', '$routeParams', '$http',
  function($scope, $routeParams, $http) {

  $scope.Name = "ArchCI"
}]);
