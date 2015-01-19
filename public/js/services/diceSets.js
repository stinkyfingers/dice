define(['app'], function (app) {
	app.factory("diceSetsFactory",function($q, $http){
		var factory = {};
		factory.getPublicDiceSets = function(){
			var deferred = $q.defer();
			$http({
				method:'get',
				url:'/getPublicDiceSets'
			}).success(function(data){
				deferred.resolve(data)
			}).error(function(data){
				deferred.reject(data);
			});
			return deferred.promise;
		};

		factory.getUserDiceSets = function(){
			var deferred = $q.defer();
			$http({
				method:'get',
				url:'/getUserDiceSets'
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