//diceSets
define(["app", "services/diceSet"],function(app){

	app.controller("diceSetController", function($scope, diceSetFactory, $routeParams){
		var id  = $routeParams.id;
		id = parseInt(id, 10);
	
		$scope.diceSet = {
			id:id
		}
		$scope.diceSet.id = id;
		$scope.diceSet = angular.toJson($scope.diceSet, false);

		$scope.diceSet = diceSetFactory.getDiceSet($scope.diceSet)
			.then(function(data){
				$scope.diceSet = data;
			},function(err){
				$scope.err = err;
			});

		$scope.roll = function(){
			$scope.results = diceSetFactory.roll($scope.diceSet)
				.then(function(data){
					$scope.results = data;
				},function(err){
					$scope.err = err;
				});
		};

	});
});