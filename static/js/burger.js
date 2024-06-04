$(document).ready(function() {
    var burgerButton = $(".navbar-toggler");
    var navbarCollapse = $(".navbar-collapse");

    if (burgerButton.length > 0) {
        burgerButton.on("click", function() {
            navbarCollapse.toggleClass("blur-menu");
        });
    }

    $(window).on("resize", function() {
        if ($(window).width() > 992) { 
            navbarCollapse.removeClass("blur-menu");
        }
    });
});
