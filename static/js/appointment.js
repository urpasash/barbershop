$(document).ready(function() {
    console.log("Script loaded and running");
var timeIntervals = [
    "10:00", "11:00", "12:00", "13:00", "14:00", "15:00", "16:00", "17:00", "18:00", "19:00", "20:00"
];

function createTimeOptions() {
    var selectTime = document.getElementById("time");
    
    selectTime.innerHTML = "";

    timeIntervals.forEach(function(interval) {
        var option = document.createElement("option");
        option.value = interval;
        option.text = interval;
        selectTime.appendChild(option);
    });
}

createTimeOptions();

});

$(document).ready(function() {
    var Masters = [
        "Никита", "Коля", "Антон", "Кирилл"
    ];
    console.log("Script");
    function createMasterOptions() {
        var selectMaster = document.getElementById("master");
        
        selectMaster.innerHTML = "";
    
        Masters.forEach(function(master) {
            var option = document.createElement("option");
            option.value = master;
            option.text = master;
            selectMaster.appendChild(option);
        });
    }
    createMasterOptions();
});



$(document).ready(function() {
    $("#appointment-form").submit(function(event) {
        event.preventDefault();

        var username = $("#username").val();
        var date = $("#date").val();
        var time = $("#time").val();
        var master = $("#master").val();

        var appointmentData = {
            username: username,
            time: date + " "+ time +":00",
            master: master
        };

        $.ajax({
            url: "/appointment",
            method: "POST",
            contentType: "application/json",
            data: JSON.stringify(appointmentData),
            success: function(response) {
                    location.reload();
                    alert("Запись добавлена");
            },
            error: function(xhr, status, error) {
                console.error("Ошибка при создании записи:", error);
                alert("Запись уже есть");
            }
        });
    });
});


