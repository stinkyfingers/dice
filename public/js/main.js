//main
require.config({
    baseUrl: "/public/js", //everything is in the js folder
    
    // alias libraries paths.  Must set 'angular'
    paths: {
        //USE LOCAL COPIES INSTEAD
        'angular': 'bower_components/angular/angular',
        'angular-route': '//ajax.googleapis.com/ajax/libs/angularjs/1.2.16/angular-route.min',
        'angularAMD': 'bower_components/angularAMD/angularAMD',
        "jquery": "//code.jquery.com/jquery-1.11.0.min",
        'angularCookies':'node_modules/angular-cookies/angular-cookies',
    },
    
    // Add angular modules that does not support AMD out of the box, put it in a shim
    shim: {
        'angular':['jquery'],
        'angularAMD': ['angular'],
        'angular-route': ['angular'],
         'angularCookies'  :{
            deps: ['angular']
        },
    },
    
    // kick start application
    deps: ['app']//was app
});