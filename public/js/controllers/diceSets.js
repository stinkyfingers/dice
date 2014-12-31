//diceSets
define(["app", "services/diceSets"],function(app){

	app.controller("diceSetsController", function($scope, diceSetsFactory, $rootScope){
		$scope.test = "Hello";
		$scope.publicDiceSets = diceSetsFactory.getPublicDiceSets()
			.then(function(data){
				$scope.publicDiceSets = data;
			},function(err){
				$scope.err = err;
			})

	});
});