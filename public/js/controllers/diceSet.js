//diceSets
define(["app", "services/diceSet"],function(app){

	app.controller("diceSetController", function($scope, diceSetFactory, $routeParams){
		$scope.editorEnabled = false;

		var id  = $routeParams.id;
		id = parseInt(id, 10);
	
		$scope.diceSet = {
			id:id,
			dice:[]
		}
		$scope.die = {};
		$scope.diceSet.id = id;
		$scope.diceSet = angular.toJson($scope.diceSet, false);
		console.log($scope.diceSet)

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

		$scope.editor = function(){
			if ($scope.editorEnabled == false){
				$scope.editorEnabled = true;
			}else{
				$scope.editorEnabled = false;
			}
		};

		$scope.saveDiceSet = function(diceSet){
			diceSetFactory.saveDiceSet(diceSet)
				.then(function(data){
					$scope.diceSet = data;
				},function(err){
					$scope.err = err;
				});
		};

		$scope.numDice = [1,2,3,4,5,6,7,8,9];

		//TODO
		$scope.setNumDice = function(){
			$scope.diceSet.dice = [];
			angular.forEach($scope.diceSet.numDice,function(v,k){
				$scope.die = {
					value:"hi"
				};
				$scope.diceSet.dice.push($scope.die);
			});
			console.log($scope.diceSet)
		}


	});
});