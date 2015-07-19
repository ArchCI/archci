/* The archci angular application */
var archci = angular.module('archci', [
  'ngRoute',
  'archciControllers',
  'ngCookies', // To save perference of i18n language
  'pascalprecht.translate'
]);

/* Configurate application like router and others*/
archci.config(['$locationProvider', '$routeProvider',
  function($locationProvider, $routeProvider) {
    /* Remove the # in url from Angular */
    $locationProvider.html5Mode(true);

    /* Set router, all in /js/controllers.js */
    $routeProvider.
      when('/', {
        templateUrl: '/static/html/builds.html',
        //controller: 'BuildsController'
      }).
      when('/builds', {
        templateUrl: '/static/html/builds.html',
        //controller: 'BuildsController'
      }).
      when('/builds/:buildId', {
        templateUrl: '/static/html/builds.html',
        //controller: 'BuildsController'
      }).
      when('/projects', {
        templateUrl: '/static/html/projects.html',
        //controller: 'ProjectsController'
      }).
      when('/projects/:projectId', {
        templateUrl: '/static/html/projects.html',
        //controller: 'ProjectsController'
      }).
      when('/workers', {
        templateUrl: '/static/html/workers.html',
        //controller: 'WorkersController'
      });
      /* No default page for angular so that beego can process API request
      otherwise({
        redirectTo: '/'
      }); */
  }]
);

/* Refer to http://www.ng-newsletter.com/posts/angular-translate.html for i18n */
archci.controller('IndexController', function ($scope, $rootScope, $translate, $route, $http) {

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

});


/* Use angular-translate for i18n and all text should be translated here */
archci.config(function ($translateProvider) {
  /* Use cookie to store the perference of i18n language */
  $translateProvider.useCookieStorage();

  /* The default language should be English */
  $translateProvider.preferredLanguage('en-us');

  /* Translate into English */
  $translateProvider.translations('en-us', {
    // Index page
    archci: 'ArchCI',
    builds: 'Builds',
    projects: 'Projects',
    workers: 'Workers',
    more: 'More',
    en_us: 'English',
    zh_cn: '简体中文',
    zh_hant: '繁體中文',

    // Build page
    no_search_result_for: 'No search result for',
    search: 'Search',
  });

  /* Translate into simplified Chinese */
  $translateProvider.translations('zh-cn', {
    // Index page
    archci: 'ArchCI',
    builds: '持续集成',
    projects: '所有项目',
    workers: '所有节点',
    more: '更多',
    en_us: 'English',
    zh_cn: '简体中文',
    zh_hant: '繁體中文',

    // Build page
    no_search_result_for: '找不到相关搜索结果',
    search: '搜索',
  });

  /* Translate into traditional Chinese */
  $translateProvider.translations('zh-hant', {
    // Index page
    archci: 'ArchCI',
    builds: '持續集成',
    projects: '所有項目',
    workers: '所有節點',
    more: '更多',
    en_us: 'English',
    zh_cn: '简体中文',
    zh_hant: '繁體中文',

    // Build page
    no_search_result_for: '找不到相關搜索結果',
    search: '搜索',
  });

});