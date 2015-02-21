(function () {
  var appController = ['$http', function ($http) {
    this.foo = 'blah2';
  }];

  var servicesController = ['$scope', '$routeParams', '$http', function ($scope, $routeParams, $http) {
    if ($routeParams.serviceId) {
      $http
        .get('/service/' + $routeParams.serviceId)
        .then(function (response) {
          console.log(response.data);
          $scope.service = response.data;
        }, function (err) {
          console.log("Error: ", err);
        });
    }

    $http
      .get('/services')
      .then(function (response) {
        console.log(response.data);
        $scope.services = response.data;
      }, function (err) {
        console.log("Error: ", err);
      });
  }];

  angular
    .module('app', [
      'ngRoute'
    ])
    .controller('appController', appController)
    .config(['$routeProvider', function ($routeProvider) {
      $routeProvider.when('/services', {
        templateUrl: 'partials/services.html',
        controller: servicesController
      })
      .when('/service/:serviceId', {
        templateUrl: 'partials/modify-service.html',
        controller: servicesController
      });
    }]);
})();
