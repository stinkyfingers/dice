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
		console.log($scope.diceSet.id)
		if ($scope.diceSet.id == 0 || isNaN($scope.diceSet.id)){
			$scope.editorEnabled = true;
			$scope.diceSet = {
				dice:[]
			}
		}
		// $scope.diceSet.id = id;
		// $scope.diceSet = angular.toJson($scope.diceSet, false);

		//TODO - add die on new diceSet or no-die ones

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

		$scope.addSide = function(die){
			var side = {
				id:0,
				dieId: die.id,
				value:''
			}
			die.sides.push(side);
		};
		$scope.addDie = function(diceSet){
			var die = {
				id:0,
				sides:[],
				diceSetId: diceSet.id
			}
			diceSet.dice.push(die);
		}



	});
});