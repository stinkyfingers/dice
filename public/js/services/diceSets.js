define(['app'], function (app) {
	app.factory("diceSetsFactory",function($q, $http){
		var factory = {};
		factory.getPublicDiceSets = function(){
			var deferred = $q.defer();
			$http({
				method:'get',
				url:'http://localhost:5000/getPublicDiceSets'
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