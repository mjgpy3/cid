(function () {
  var appController = [function () {
  }];

  angular
    .module('app', [
      'ngRoute'
    ])
    .controller('appController', appController)
    .config(['$routeProvider', function ($routeProvider) {
      $routeProvider.when('/services', {
        templateUrl: 'js/services/partials/services.html',
        controller: ServicesController
      })
      .when('/service/:serviceId', {
        templateUrl: 'js/services/partials/modify_service.html',
        controller: ServicesController
      });
    }]);
})();
