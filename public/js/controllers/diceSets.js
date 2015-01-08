//diceSets
define(["app", "services/diceSets"],function(app){

	app.controller("diceSetsController", function($scope, diceSetsFactory, $rootScope){


		$scope.publicDiceSets = diceSetsFactory.getPublicDiceSets()
			.then(function(data){
				$scope.publicDiceSets = data;
			},function(err){
				$scope.err = err;
			});
		$scope.userDiceSets = diceSetsFactory.getUserDiceSets()
			.then(function(data){
					$scope.userDiceSets = data;
					console.log(data)
				},function(err){
					$scope.err = err;
				});
		$scope.deleteDiceSet = function (diceSet){
			if (confirm("Are you sure?")){
				diceSetsFactory.deleteDiceSet(diceSet)
					.then(function(data){
						angular.forEach($scope.userDiceSets,function(v,k){
							if (diceSet == v){
								$scope.userDiceSets.splice(k,1);
							};
						});
							// $	scope.userDiceSets = data;
						},function(err){
							$scope.err = err;
						});
				};
			};

	});
});