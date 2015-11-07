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
				url:'/register',
				data:user
			}).success(function(data){
				deferred.resolve(data)
			}).error(function(data){
				deferred.reject(data);
			});
			return deferred.promise;
		};

		factory.resetPassword = function(user){
			var deferred = $q.defer();
			$http({
				method:'post',
				headers: {
					'Content-Type': 'application/json',
				},
				url:'/resetPassword',
				data:user
			}).success(function(data){
				deferred.resolve(data)
			}).error(function(data){
				deferred.reject(data);
			});
			return deferred.promise;
		};

		factory.logout = function(user){
			var deferred = $q.defer();
			$http({
				method:'post',
				headers: {
					'Content-Type': 'application/json',
				},
				url:'/logout'
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