define(['app'], function (app) {
	app.factory("testFactory",function($q, $http, $rootScope){
		var factory = {};
		factory.testGet = function(){
			return "fun";
		}
		return factory;
	});
});