$(document).ready(function() {
    $("#register-form").on("submit", function(event) {
        event.preventDefault();
    
        var username = $("#username_reg").val();
        var password = $("#password_reg").val();

        console.log("Sending registration request");

        fetch("/register", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify({
                username_reg: username,
                password_reg: password,
            }),
        })
        .then(response => {
            if (response.ok) {
                return response.json(); 
            } else {
                if (response.status === 409) { 
                    throw new Error("Username already taken");
                } else {
                    throw new Error("Registration error");
                }
            }
        })
        .then(data => {
            console.log("Received response:", data); 
            if (data.success) { 
                $("#success-alert").fadeIn();
                setTimeout(() => {
                    $("#success-alert").fadeOut();
                }, 1000);
                setTimeout(() => {
                    window.location.href = "/auth"; 
                }, 1000);
            }
        })
        .catch(error => {
            console.error("Error during registration:", error);
            let message = "Ошибка";
            if (error.message === "Username already taken") { 
                message = "Имя уже занято"; 
            }
            $("#register-error-alert").text(message).fadeIn();
            setTimeout(() => {
                $("#register-error-alert").fadeOut();
            }, 1000);
        });
    });
});
