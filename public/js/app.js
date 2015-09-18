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
			when("/",angularAMD.route({
				templateUrl: '/public/js/templates/home.tmpl',
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
			otherwise({redirectTo: "/"});

		

	});

	app.run(function($rootScope, $cookies){
		var u = $cookies.user;
		$rootScope.user = u;
	});


  return angularAMD.bootstrap(app);
});