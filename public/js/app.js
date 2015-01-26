//app
define(['angularAMD', 'angular-route', 'angularCookies'], function (angularAMD) {


	var  app = angular.module("app",["ngRoute","ngCookies"],function($interpolateProvider){
		$interpolateProvider.startSymbol('[[');
	    $interpolateProvider.endSymbol(']]');
	});

	app.config(function($routeProvider, $locationProvider){

		if(window.history && window.history.pushState){
		      $locationProvider.html5Mode({
				  enabled: true,
				  requireBase: false
				});
		    }

		$routeProvider.
			// when('/',angularAMD.route({
			// 	templateUrl: '/templates/index.tmpl',
			// 	controller: 'indexController',
			// 	controllerUrl: 'controllers/index'
			// })).
			when("/test",angularAMD.route({
				templateUrl: '/public/js/templates/test.tmpl',
				controller: 'testController',
				controllerUrl: 'controllers/test'
			})).
			when("/diceSets",angularAMD.route({
				templateUrl: '/public/js/templates/diceSets.tmpl',
				controller: 'diceSetsController',
				controllerUrl: 'controllers/diceSets'
			})).
			when('/diceSet/:id', angularAMD.route({
		        templateUrl: "/public/js/templates/diceSet.tmpl",
		        controller: "diceSetController",
		        controllerUrl: "controllers/diceSet"
		    })).
		    when('/diceSet', angularAMD.route({
		        templateUrl: "/public/js/templates/diceSet.tmpl",
		        controller: "diceSetController",
		        controllerUrl: "controllers/diceSet"
		    })).
		    when('/registration', angularAMD.route({
		        templateUrl: "/public/js/templates/user.tmpl",
		        controller: "userController",
		        controllerUrl: "controllers/user"
		    })).
		    when('/reset', angularAMD.route({
		        templateUrl: "/public/js/templates/resetPassword.tmpl",
		        controller: "userController",
		        controllerUrl: "controllers/user"
		    })).
		 //     when('/admin/setlistsNew', angularAMD.route({
		 //        templateUrl: "/templates/admin/setlistsNew.tmpl",
		 //        controller: "setlistControllerNew",
		 //        controllerUrl: "controllers/setlistsNew"
		 //    })).
			// when('/signup', angularAMD.route({
		 //        templateUrl: "/templates/signup.tmpl",
		 //        controller: "adminController",
		 //        controllerUrl: "controllers/admin"
		 //    })).
			otherwise({redirectTo: "/"});

		

	});

	app.run(function($rootScope, $cookies){
		var u = $cookies.user;
		$rootScope.user = u;
	});


	//user on rootScope
	// app.run(function($rootScope, mainFactory){
	// 	$rootScope.user={};
	// 	$rootScope.user = mainFactory.getCurrentUser()
	// 		.then(function(data){
	// 			$rootScope.user = data;
	// 			// console.log(data);//log userdata
	// 		},function(data){
	// 			$rootScope.user = {};
	// 			console.log("Not logged in.")
	// 		});
	// });

	// app.factory("mainFactory",function($q, $location, $routeParams, $route, $http){
	// 	var service = {};
	//       service.getCurrentUser = function(){
	//         var deferred = $q.defer();
	//         $http({
	//           method:"GET",
	//           url:"/user"
	//         }).success(function(data){
	//           deferred.resolve(data);
	//         }).error(function(){
	//           deferred.reject("Error");
	//         })
	//         return deferred.promise;
	//       }

	//       return service;
	// });

	// app.directive("ngConfirmClick",function(){
	// 	return {
	// 		restrict: 'A',
	// 		link: function(scope, element, attr){
	// 			var msg = attr.ngConfirmClick || "Are you sure?";
 //                var clickAction = attr.confirmedClick; //proceed with directive confirmedClick
 //                element.bind('click',function (event) {
 //                    if ( window.confirm(msg) ) {
 //                        scope.$eval(clickAction)
 //                    }
 //                });
	// 		}
	// 	}	
	// });





  return angularAMD.bootstrap(app);
});