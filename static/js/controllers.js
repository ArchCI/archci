/* All angular application controllers */
var archciControllers = angular.module('archciControllers', []);

archciControllers.controller('BuildsController', ['$scope', '$routeParams', '$http',
  function($scope, $routeParams, $http) {

    /* Get the version object
    $http.get('/dockerapi/version').success(function(data) {
      $scope.version = data;
      $scope.Os = $scope.version.Os;
      $scope.KernelVersion = $scope.version.KernelVersion;
      $scope.GoVersion = $scope.version.GoVersion;
      $scope.Version = $scope.version.Version;
    });

    /* Get the info object
    $http.get('/dockerapi/info').success(function(data) {
      $scope.info = data;
      $scope.Containers = $scope.info.Containers;
      $scope.Images = $scope.info.Images;
    });
    */

  $scope.Name = "ArchCI"
}]);

archciControllers.controller('ProjectsController', ['$scope', '$routeParams', '$http',
  function($scope, $routeParams, $http) {

  $scope.Name = "ArchCI"
}]);

archciControllers.controller('WorkersController', ['$scope', '$routeParams', '$http',
  function($scope, $routeParams, $http) {

  $scope.Name = "ArchCI"
}]);
