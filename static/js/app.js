(function () {
  var appController = ['$http', function ($http) {
    this.foo = 'blah';
  }];

  angular
    .module('app', [])
    .controller('appController', appController);
})();
