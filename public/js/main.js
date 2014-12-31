//main
require.config({
    baseUrl: "/public/js", //everything is in the js folder
    
    // alias libraries paths.  Must set 'angular'
    paths: {
        // 'angular': '//ajax.googleapis.com/ajax/libs/angularjs/1.2.16/angular.min',
        // 'angular-route': '//ajax.googleapis.com/ajax/libs/angularjs/1.2.16/angular-route.min',
        // 'angularAMD': '//cdn.jsdelivr.net/angular.amd/0.2.0/angularAMD.min'
        //USE LOCAL COPIES INSTEAD
        'angular': 'bower_components/angular/angular',
        'angular-route': '//ajax.googleapis.com/ajax/libs/angularjs/1.2.16/angular-route.min',
        'angularAMD': 'bower_components/angularAMD/angularAMD',
        "jquery": "//code.jquery.com/jquery-1.11.0.min",
        // "jquery-ui": "lib/jquery-ui-1.11.1.custom/jquery-ui.min"
    },
    
    // Add angular modules that does not support AMD out of the box, put it in a shim
    shim: {
        'angular':['jquery'],
        'angularAMD': ['angular'],
        'angular-route': ['angular'],
         // "jquery-ui": {
         //    exports: "$",
         //    deps: ['jquery']
        // },
    },
    
    // kick start application
    deps: ['app']//was app
});