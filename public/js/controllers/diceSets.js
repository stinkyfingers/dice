//diceSets
define(["app", "services/diceSets"],function(app){

	app.controller("diceSetsController", function($scope, diceSetsFactory, $rootScope){
		// $scope.test = "Hello";
		// $scope.userId = $rootScope.user;

		$scope.publicDiceSets = diceSetsFactory.getPublicDiceSets()
			.then(function(data){
				$scope.publicDiceSets = data;
			},function(err){
				$scope.err = err;
			});
		$scope.userDiceSets = diceSetsFactory.getUserDiceSets()
			.then(function(data){
					$scope.userDiceSets = data;
				},function(err){
					$scope.err = err;
				});

	});
});