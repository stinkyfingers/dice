//app
define(['angularAMD', 'angular-route'], function (angularAMD) {


	var  app = angular.module("app",["ngRoute"],function($interpolateProvider){
		$interpolateProvider.startSymbol('[[');
	    $interpolateProvider.endSymbol(']]');
	});

	app.config(function($routeProvider, $locationProvider){

		if(window.history && window.history.pushState){
		      $locationProvider.html5Mode(true);
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
			}))
			// when("/setlists",angularAMD.route({
			// 	templateUrl: '/templates/setlist.tmpl',
			// 	controller: 'setlistController',
			// 	controllerUrl: 'controllers/setlists'
			// })).
			// when('/admin/main', angularAMD.route({
		 //        templateUrl: "/templates/main.tmpl",
		 //        controller: "adminController",
		 //        controllerUrl: "controllers/admin"
		 //    })).
		 //    when('/admin/games', angularAMD.route({
		 //        templateUrl: "/templates/admin/games.tmpl",
		 //        controller: "gamesController",
		 //        controllerUrl: "controllers/games"
		 //    })).
		 //    when('/admin/setlists', angularAMD.route({
		 //        templateUrl: "/templates/admin/setlists.tmpl",
		 //        controller: "setlistController",
		 //        controllerUrl: "controllers/setlists"
		 //    })).
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
			// otherwise({redirectTo: "/"});

		

	});

	// //user on rootScope
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
	//           url:"/api/user"
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