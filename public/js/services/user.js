define(['app'], function (app) {
	app.factory("userFactory",function($q, $http){
		var factory = {};

		factory.register = function(user){
			var deferred = $q.defer();
			$http({
				method:'post',
				headers: {
					'Content-Type': 'application/json',
				},
				url:'http://localhost:5000/register',
				data:user
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