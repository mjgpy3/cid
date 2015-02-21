var ServicesController;

(function () {
  ServicesController = [
    '$scope',
    '$routeParams',
    '$http',
    '$location', function ($scope, $routeParams, $http, $location) {

    $scope.update = function (service) {
      $http.put('/service', service);

      $location.path('/services');
    };

    if ($routeParams.serviceId) {
      $http
        .get('/service/' + $routeParams.serviceId)
        .then(function (response) {
          console.log("Found Services: ", response.data);
          $scope.service = response.data;
        }, function (err) {
          console.log("Error: ", err);
        });
    } else {
      $http
        .get('/services')
        .then(function (response) {
          console.log(response.data);
          $scope.services = response.data;
        }, function (err) {
          console.log("Error: ", err);
        });
    }
  }];
})();
