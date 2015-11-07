//diceSets
define(["app", "services/user"],function(app){

	app.controller("mainController", function($scope, $rootScope, userFactory, $location){
		console.log('mainnn')
		userFactory.logout().then(function(){
			$location.path("/");
			
		}, function(err){
			$scope.err = err;
		});

	});
});