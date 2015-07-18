/* The seagull angular application */
var seagull = angular.module('seagull', [
  'ngRoute',
  'seagullControllers',
  'ngCookies', // To save perference of i18n language
  'pascalprecht.translate'
]);

/* Configurate application like router and others*/
seagull.config(['$locationProvider', '$routeProvider',
  function($locationProvider, $routeProvider) {
    /* Remove the # in url from Angular */
    $locationProvider.html5Mode(true);

    /* Set router, all in /js/controllers.js */
    $routeProvider.
      when('/', {
        templateUrl: '/static/html/home.html',
        controller: 'HomeController'
      }).
      when('/containers', {
        templateUrl: '/static/html/containers.html',
        controller: 'ContainersController'
      }).
      when('/containers/:id', {
        templateUrl: '/static/html/container.html',
        controller: 'ContainerController'
      }).
      when('/images', {
        templateUrl: '/static/html/images.html',
        controller: 'ImagesController'
      }).
      when('/images/:id', {
        templateUrl: '/static/html/image.html',
        controller: 'ImageController'
      }).
      when('/images/:user/:repo', {
        templateUrl: '/static/html/image.html',
        controller: 'ImageController'
      }).
      when('/configuration', {
        templateUrl: '/static/html/configuration.html',
        controller: 'ConfigurationController'
      }).
      when('/dockerhub', {
        templateUrl: '/static/html/dockerhub.html',
        controller: 'DockerhubController'
      });
      /* No default page for angular so that beego can process API request
      otherwise({
        redirectTo: '/'
      }); */
  }]
);

/* Refer to http://www.ng-newsletter.com/posts/angular-translate.html for i18n */
seagull.controller('IndexController', function ($scope, $rootScope, $translate, $route, $http) {

  /* Change languages with the language string */
  $scope.changeLanguage = function (key) {
    $translate.use(key);
  };

  /* Determine it is English or not */
  $scope.isEnUs = function () {
     return $translate.use() == "en-us";
  }

  /* Determine it is simplified Chinese or not */
  $scope.isZhCn = function () {
           return $translate.use() == "zh-cn";
  }

  /* Determine it is traditional Chinese or not */
  $scope.isZhHant = function () {
     return $translate.use() == "zh-hant";
  }

    /* Determine it is German or not */
    $scope.isDeDe = function () {
       return $translate.use() == "de-de";
    }

    /* Determine it is French or not */
    $scope.isFrFr = function () {
       return $translate.use() == "fr-fr";
    }

 });


 /* Use angular-translate for i18n and all text should be translated here */
 seagull.config(function ($translateProvider) {
   /* Use cookie to store the perference of i18n language */
   $translateProvider.useCookieStorage();

   /* The default language should be English */
   $translateProvider.preferredLanguage('en-us');

   /* Translate into English */
   $translateProvider.translations('en-us', {
     // Index html
     seagull: 'Seagull',
     containers: 'Containers',
     images: 'Images',
     configuration: 'Configuration',
     dockerhub: 'DockerHub'
  });

});