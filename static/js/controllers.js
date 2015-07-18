/* All angular application controllers */
var archciControllers = angular.module("archciControllers", []);

archciControllers.controller("BuildsController", ["$scope", "$routeParams", "$http",
  function($scope, $routeParams, $http) {

    /*
    {"log":"Unable to find image 'golang:1.4' locally","Next":true}
    */
    $http.get("/v1/builds/123/logs/0").success(function(data) {
      $scope.data = data;

      $scope.fullLog = data.log;

      // TODO(tobe): we use "next" but not "Next" in beego api
      next = data.Next;


      // TODO(tobe): how to iterate all the logs, it calls twice
      if(next){

        index = 0;

        $http.get("/v1/builds/123/logs/" + index).success(function(data) {
          $scope.fullLog += "\n";
          $scope.fullLog += data.log;

          console.log(index)
          console.log(data.Next)

          next = data.Next;

          index++;

        });

      }



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
