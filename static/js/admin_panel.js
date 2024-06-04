$(document).ready(function() {
    $.ajax({
        url: "/users",
        method: "GET",
        success: function(response) {
            if (response.error) {
                console.error("Ошибка при получении списка пользователей:", response.error);
            } else {
                displayUsers(response);
            }
        },
        error: function(xhr, status, error) {
            console.error("Ошибка при запросе на сервер:", error);
        }
    });

    function displayUsers(users) {
        var clientsTableBody = document.getElementById("clients-table-body");
        clientsTableBody.innerHTML = "";
    
        users.forEach(function(user) {
            var row = document.createElement("tr");

            var status = user.IsAdmin ? "Администратор" : "Пользователь";

            row.innerHTML = `
                <td>${user.ID}</td>
                <td>${user.Username}</td>
                <td>${user.Password}</td>
                <td>${status}</td> <!-- Добавляем столбец статуса -->
                <td>
                <button class="btn btn-danger btn-sm delete-btn-user">Удалить</button>
                <button class="btn btn-primary btn-sm edit-btn">Редактировать</button>
                </td>
            `;
            
            clientsTableBody.appendChild(row);
        });
    }

    function handleUserDeleteButtonClick(event) {
        var ID = $(event.target).closest("tr").find("td:first").text();
        console.log(ID);
        console.log(JSON.stringify({ ID: ID }));
        $.ajax({
            url: `/users/${ID}`,
            method: "DELETE",
            success: function(response) {
                console.log("Пользователь успешно удален:", response);
                location.reload();
            },
            error: function(xhr, status, error) {
                console.error("Ошибка при удалении пользователя:", error);
            }
        });
    }
    $(document).on("click", ".delete-btn-user", handleUserDeleteButtonClick);
    
});


$(document).ready(function() {
    $("#add-client-btn").on("click", function() {
        $("#addClientForm")[0].reset();
        $("#addClientModal").modal("show");
    });
});

$(document).ready(function() {
    $("#edit-client-btn").on("click", function() {
        $("#editClientForm")[0].reset();
        $("#editClientModal").modal("show");
    });
});

$(document).ready(function() {
    $("#saveClientBtn").on("click", function() {
        var name = $("#clientName").val();
        var password = $("#clientPassword").val();
        var status = $("#clientStatus").val();
        var newClient = {
            username: name,
            password: password,
            isAdmin: status === "admin"
        };

        if (name === "" || password === "") {
            alert("Пожалуйста, заполните все поля");
            return;
        }

        $.ajax({
            url: "/users",
            method: "POST",
            contentType: "application/json",
            data: JSON.stringify(newClient),
            success: function(response) {
                console.log("Новый клиент успешно добавлен:", response);
                location.reload();
            },
            error: function(xhr, status, error) {
                console.error("Ошибка при добавлении клиента:", error);
            }
        });

        $("#addClientModal").modal("hide");
    });

    $(document).on("click", ".edit-btn", function() {
        var row = $(this).closest("tr");
        var ID= row.find("td:eq(0)").text(); 
        var username = row.find("td:eq(1)").text(); 
        var password = row.find("td:eq(2)").text(); 
        var status = row.find("td:eq(3)").text(); 
    
        $("#editClientID").val(ID);
        $("#editClientName").val(username);
        $("#editClientPassword").val(password);
        $("#editClientStatus").val(status);
    
        $("#editClientModal").modal("show");
    
        $("#saveEditedClientBtn").on("click", function() {
        var clientId = $("#editClientID").val(ID);
    
            var name = $("#editClientName").val();
            var password = $("#editClientPassword").val();
            var status = $("#editClientStatus").val();
    
            if (name === "" || password === "") {
                alert("Пожалуйста, заполните все поля");
                return;
            }
    
            var editedClient = {
                username: name,
                password: password,
                isAdmin: status === "admin"
            };
            console.log(JSON.stringify({ ID, name, password, status}));
        
            $.ajax({
                url: `/users/${clientId}`,
                method: "PUT",
                contentType: "application/json",
                data: JSON.stringify(editedClient),
                success: function(response) {
                    console.log("Изменения сохранены:", response);
                    location.reload();
                },
                error: function(xhr, status, error) {
                    console.error("Ошибка при сохранении изменений:", error);
                }
            });
    
            $("#editClientModal").modal("hide");
        });
    });
    
    $(document).on("click", ".edit-btn2", function() {
        var row = $(this).closest("tr");
        var ID = row.find("td:eq(0)").text();
        var username = row.find("td:eq(1)").text();
        var time = row.find("td:eq(2)").text();
        var master = row.find("td:eq(3)").text();

        $("#editAppointmentID").val(ID);
        $("#editAppointmentUsername").val(username);
        $("#editAppointmentTime").val(time);
        $("#editAppointmentMaster").val(master);

        $("#editAppointmentModal").modal("show");
    });

    // Сохранение изменений записи
    $("#saveEditedAppointmentBtn").on("click", function() {
        var appointmentId = $("#editAppointmentID").val();
        var username = $("#editAppointmentUsername").val();
        var time = $("#editAppointmentTime").val();
        var master = $("#editAppointmentMaster").val();

        if (username === "" || time === "" || master === "") {
            alert("Пожалуйста, заполните все поля");
            return;
        }

        var editedAppointment = {
            username: username,
            time: time + ":00",
            master: master
        };

        console.log(JSON.stringify({appointmentId, username, time, master}));

        $.ajax({
            url: `/appointments/${appointmentId}`,
            method: "PUT",
            contentType: "application/json",
            data: JSON.stringify(editedAppointment),
            success: function(response) {
                console.log("Изменения сохранены:", response);
                location.reload();
            },
            error: function(xhr, status, error) {
                console.error("Ошибка при сохранении изменений:", error);
            }
        });

        $("#editAppointmentModal").modal("hide");
    });
});



$(document).ready(function() {
    // Загружаем все записи при загрузке страницы
    $.ajax({
        url: "/appointments",
        method: "GET",
        success: function(response) {
            displayAppointments(response); // Отображаем полученные записи
        },
        error: function(xhr, status, error) {
            console.error("Ошибка при получении списка записей:", error);
        }
    });

    function displayAppointments(appointments) {
        var appointmentsTableBody = $("#appointments-table-body");
        appointmentsTableBody.empty();
        appointments.forEach(function(appointment) {
            var row = $("<tr>");
            row.html(`
                <td>${appointment.ID}</td>
                <td>${appointment.Username}</td>
                <td>${appointment.Time}</td>
                <td>${appointment.Master}</td>
                <td>
                    <button class="btn btn-danger btn-sm delete-btn-appointment" data-id="${appointment.ID}">Удалить</button>
                    <button class="btn btn-primary btn-sm edit-btn2">Редактировать</button>
                </td>
            `);
            appointmentsTableBody.append(row);
        });
    }

    // Обработчик нажатия на кнопку "Удалить"
    $(document).on("click", ".delete-btn", function() {
        var appointmentID = $(this).data("id");

        $.ajax({
            url: `/appointments/${appointmentID}`,
            method: "DELETE",
            success: function(response) {
                console.log("Запись успешно удалена:", response);
                // Обновляем список записей после удаления
                location.reload();
            },
            error: function(xhr, status, error) {
                console.error("Ошибка при удалении записи:", error);
            }
        });
    });

    function handleAppointmentDeleteButtonClick(event) {
        var ID = $(event.target).closest("tr").find("td:first").text();
        console.log(ID);
        console.log(JSON.stringify({ ID: ID }));
        $.ajax({
            url: `/appointments/${ID}`,
            method: "DELETE",
            success: function(response) {
                console.log("Запись успешно удалена:", response);
                location.reload();
            },
            error: function(xhr, status, error) {
                console.error("Ошибка при удалении записи:", error);
            }
        });
    }
    $(document).on("click", ".delete-btn-appointment", handleAppointmentDeleteButtonClick);
});
