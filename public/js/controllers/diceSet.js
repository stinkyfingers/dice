//diceSets
define(["app", "services/diceSet"],function(app){

	app.controller("diceSetController", function($scope, $rootScope, diceSetFactory, $routeParams){
		$scope.editorEnabled = false;

		var id  = $routeParams.id;
		id = parseInt(id, 10);
	
		$scope.diceSet = {
			id:id,
			dice:[]
		}
		$scope.die = {};

		if ($scope.diceSet.id == 0 || isNaN($scope.diceSet.id)){
			$scope.editorEnabled = true;
			dice = {};
			$scope.diceSet = {
				userId: $rootScope.user,
				dice:[]
			}
		}

		if ($scope.diceSet.id > 0){
			$scope.diceSet = diceSetFactory.getDiceSet($scope.diceSet)
				.then(function(data){
					$scope.diceSet = data;
				},function(err){
					$scope.err = err;
				});
		}

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
			if (typeof diceSet.dice == 'undefined'){
				diceSet.dice = [];
			}

			var die = {
				id:0,
				sides:[],
				diceSetId: diceSet.id
			}
			diceSet.dice.push(die);
		}

		$scope.deleteSide = function(side){
			diceSetFactory.deleteSide(side).then(function(data){
				angular.forEach($scope.diceSet.dice,function(die,k){
					angular.forEach($scope.diceSet.dice[k].sides, function(v,key){
						if (side == v){
							$scope.diceSet.dice[k].sides.splice(key,1);
						}
					});
				});
			},function(err){
				$scope.err = err;
			});
		}

		$scope.deleteDie = function(die){
			diceSetFactory.deleteDie(die).then(function(data){
				angular.forEach($scope.diceSet.dice,function(die,k){
					if (die == $scope.diceSet.dice[k]){
						$scope.diceSet.dice.splice(k,1);
					}
				});
			},function(err){
				$scope.err = err;
			});
		}

		$scope.deleteDiceSet = function(diceSet){
			diceSetFactory.deleteDiceSet(diceSet).then(function(data){
				$scope.diceSet = {};
			},function(err){
				$scope.err = err;
			});
		}



	});
});