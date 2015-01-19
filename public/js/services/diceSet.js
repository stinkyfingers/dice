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
				url:'/getDiceSet',
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
				url:'/roll',
				data:diceSet
			}).success(function(data){
				deferred.resolve(data)
			}).error(function(data){
				deferred.reject(data);
			});
			return deferred.promise;
		};

		factory.saveDiceSet = function(diceSet){
			var deferred = $q.defer();
			$http({
				method:'post',
				headers: {
					'Content-Type': 'application/json',
				},
				url:'/saveDiceSet',
				data:diceSet
			}).success(function(data){
				deferred.resolve(data)
			}).error(function(data){
				deferred.reject(data);
			});
			return deferred.promise;
		};

		factory.deleteSide = function(side){
			var deferred = $q.defer();
			$http({
				method:'post',
				headers: {
					'Content-Type': 'application/json',
				},
				url:'/deleteSide',
				data:side
			}).success(function(data){
				deferred.resolve(data)
			}).error(function(data){
				deferred.reject(data);
			});
			return deferred.promise;
		};

		factory.deleteDie = function(die){
			var deferred = $q.defer();
			$http({
				method:'post',
				headers: {
					'Content-Type': 'application/json',
				},
				url:'/deleteDie',
				data:die
			}).success(function(data){
				deferred.resolve(data)
			}).error(function(data){
				deferred.reject(data);
			});
			return deferred.promise;
		};

		factory.deleteDiceSet = function(diceSet){
			var deferred = $q.defer();
			$http({
				method:'post',
				headers: {
					'Content-Type': 'application/json',
				},
				url:'/deleteDiceSet',
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