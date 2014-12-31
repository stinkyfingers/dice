//test
define(["app", "services/test"],function(app){


	app.controller("testController", function($scope, testFactory, $rootScope){
		$scope.test = "Hello";
		$scope.testGet = testFactory.testGet();

	});
});