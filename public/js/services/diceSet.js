define(['app'], function (app) {
	app.factory("diceSetFactory",function($q, $http){
		var factory = {};

		factory.getDiceSet = function(diceSet){
			var deferred = $q.defer();
			$http({
				method:'post',
				headers: {
					'Content-Type': 'application/json',
				},
				url:'http://localhost:5000/getDiceSet',
				data:diceSet
			}).success(function(data){
				deferred.resolve(data)
			}).error(function(data){
				deferred.reject(data);
			});
			return deferred.promise;
		};

		factory.roll = function(diceSet){
			var deferred = $q.defer();
			$http({
				method:'post',
				headers: {
					'Content-Type': 'application/json',
				},
				url:'http://localhost:5000/roll',
				data:diceSet
			}).success(function(data){
				deferred.resolve(data)
			}).error(function(data){
				deferred.reject(data);
			});
			return deferred.promise;
		};

		
		return factory;

	});
});