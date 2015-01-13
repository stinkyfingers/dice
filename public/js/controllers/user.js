//diceSets
define(["app", "services/user"],function(app){

	app.controller("userController", function($scope, userFactory, $location){

		$scope.register = function(user){
			if (user.password != $scope.passwordConfirm){
				alert("Passwords do not match.");
				$scope.passwordConfirm = "";
				$scope.user.password = "";
				return 
			}
			userFactory.register(user).then(function(data){
				$scope.user = data;
				alert("Success, "+user.email+".");
				$location.path("/");
			},function(err){
				$scope.err = err;
			});
		}

		

	});
});