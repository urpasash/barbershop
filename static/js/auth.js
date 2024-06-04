$(document).ready(function() {
    console.log("Script loaded and running");

    $("#login-form").on("submit", function(event) {
        event.preventDefault(); 

        var username = $("#username").val();
        var password = $("#password").val();

        console.log("Sending login request");
        fetch("/login", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify({
                username: username,
                password: password,
            }),
        })
        .then((response) => response.json())
        .then((data) => {
            console.log("Received response:", data); 
            if (data.success) {
                $("#success-alert").fadeIn(); 
                
                if (data.is_admin) {
                    setTimeout(function() {
                        window.location.href = "/admin"; 
                    }, 1000);
                } else {
                    setTimeout(function() {
                        window.location.href = "/user"; 
                    }, 1000);
                }
            } else {
                $("#error-alert").fadeIn();
                setTimeout(function() {
                    $("#error-alert").fadeOut();
                }, 3000);
            }
        })
        .catch((error) => {
            console.error("Error:", error);
            $("#error-alert").fadeIn();
            setTimeout(function() {
                $("#error-alert").fadeOut();
            }, 3000);
        });
    });
});
