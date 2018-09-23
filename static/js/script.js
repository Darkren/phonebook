function renderUsers(users) {
    $("#content").html(`
        <table id="users-table">
            <tr>
                <th>Name</th>
                <th>Surname</th>
                <th>Age</th>
                <th></th>
                <th></th>
            </tr>
        </table>`);

    for (let i = 0; i < users.length; i++) {
        $("#users-table").append(`
            <tr data-id="${users[i].id}">
                <td class="name-td"><input type="text" value="${users[i].name}" /></td>
                <td class="surname-td"><input type="text" value="${users[i].surname}" /></td>
                <td class="age-td"><input type="text" value="${users[i].age}" /></td>
                <td><button class="text-white bg-success btn phones-btn">Phones</button></td>
                <td><button class="text-white bg-primary btn save-user-btn">Save</button></td>
                <td><button class="text-white bg-danger btn del-user-btn">Remove</button></td>
            </tr>`);
    }

    $("#panel-div").html(`<button class="text-white bg-success btn" id="add-user-btn">Add</button>`);
}

function renderPhones(phones, userId) {
    $("#content").html(
        `<table id="phones-table">
            <tr>
                <th>Phone</td>
                <th></th>
                <th></th>
            </tr>
        </table>`
    );

    for (let i = 0; i < phones.length; i++) {
        $("#phones-table").append(`
            <tr data-id="${phones[i].id}" data-user-id="${userId}">
                <td class="phone-td"><input type="text" value="${phones[i].phone}" /></td>
                <td><button class="text-white bg-primary btn save-phone-btn">Save</button></td>
                <td><button class="text-white bg-danger btn del-phone-btn">Remove</button></td>
            </tr>
        `);
    }

    $("#panel-div").html(`<button class="text-white bg-success btn" data-user-id="${userId}" id="add-phone-btn">Add</button>`);
    $("#panel-div").append(`<button class="text-white bg-primary btn" id="back-to-users-btn">Back</button>`);
}

function deleteRow(id) {
    $("tr").each(function () {
        if (Number($(this).attr("data-id")) == id) {
            $(this).remove();

            return false;
        }
    });
}

let client = new JSONRPCClient();

$(document).ready(function () {
    client.connect("ws://127.0.0.1:8089/ws").then(() => {
        $("body").on("click", "#add-user-btn", function () {
            $("#content").html(`
                <form>
                    <label>Name:</label>
                    <input type="text" id="name" value="" />
                    <label>Surname:</label>
                    <input type="text" id="surname" value="" />
                    <label>Age:</label>
                    <input type="text" id="age" value="" />
                    <button class="text-white bg-success btn" id="send-add-user-btn">Add</button>
                </form>
            `);

            $("#panel-div").html(`<button class="text-white bg-primary btn" id="add-user-back-btn">Back</button>`);
        });

        $("body").on("click", "#add-user-back-btn", function (e) {
            e.preventDefault();

            client.getUsers().then((usersResp) => renderUsers(usersResp.result));
        });

        $("body").on("click", "#send-add-user-btn", function (e) {
            e.preventDefault();

            let name = $("#name").val().trim();
            let surname = $("#surname").val().trim();
            let age = Number($("#age").val().trim());

            let params = {
                name: name,
                surname: surname,
                age: age
            };

            client.addUser(params).then((addUserResp) => {
                client.getUsers().then((usersResp) => renderUsers(usersResp.result));
            });
        });

        $("body").on("click", ".save-user-btn", function (e) {
            e.preventDefault();

            let id = Number($(this).parent().parent().attr("data-id"));
            let name = $(this).parent().siblings(".name-td").children("input").val().trim();
            let surname = $(this).parent().siblings(".surname-td").children("input").val().trim();
            let age = Number($(this).parent().siblings(".age-td").children("input").val().trim());

            let params = {
                id: id,
                name: name,
                surname: surname,
                age: age
            };

            client.updateUser(params).then((updateUserResp) => { });
        });

        $("body").on("click", ".del-user-btn", function (e) {
            e.preventDefault();

            let id = Number($(this).parent().parent().attr("data-id"));

            client.deleteUser(id).then((deleteUserResp) => deleteRow(id));
        });

        $("body").on("click", ".phones-btn", function (e) {
            e.preventDefault();

            let userId = Number($(this).parent().parent().attr("data-id"));

            client.getPhones(userId).then((phonesResp) => renderPhones(phonesResp.result, userId));
        });

        $("body").on("click", "#back-to-users-btn", function (e) {
            e.preventDefault();

            client.getUsers().then((usersResp) => renderUsers(usersResp.result));
        });

        $("body").on("click", "#add-phone-btn", function (e) {
            e.preventDefault();

            let userId = $(this).attr("data-user-id");

            $("#content").html(`
                <form>
                    <label>Phone:</label>
                    <input type="text" id="phone" value="" />
                    <button class="text-white bg-success btn" data-user-id="${userId}" id="send-add-phone-btn">Add</button>
                </form>
            `);

            $("#panel-div").html(`<button class="text-white bg-primary btn" data-user-id="${userId}" id="add-phone-back-btn">Back</button>`);
        });

        $("body").on("click", "#send-add-phone-btn", function (e) {
            e.preventDefault();

            let phone = $("#phone").val().trim();
            let userId = Number($(this).attr("data-user-id"));

            let params = {
                phone: phone,
                userId: userId
            };

            client.addPhone(params).then((addPhoneResp) => {
                client.getPhones(userId).then((phonesResp) => renderPhones(phonesResp.result, userId));
            });
        });

        $("body").on("click", ".save-phone-btn", function (e) {
            e.preventDefault();

            let phone = $(this).parent().siblings(".phone-td").children("input").val().trim();
            let id = Number($(this).parent().parent().attr("data-id"));
            let userId = Number($(this).parent().parent().attr("data-user-id"));

            let params = {
                phone: phone,
                id: id,
                userId: userId
            };

            client.updatePhone(params).then((updatePhoneResp) => { });
        });

        $("body").on("click", ".del-phone-btn", function (e) {
            e.preventDefault();

            let id = Number($(this).parent().parent().attr("data-id"));
            let userId = Number($(this).parent().parent().attr("user-data-id"));

            client.deletePhone(id).then((deletePhoneResp) => deleteRow(id));
        });

        $("body").on("click", "#add-phone-back-btn", function (e) {
            e.preventDefault();

            let userId = Number($(this).attr("data-user-id"));

            client.getPhones(userId).then((phonesResp) => renderPhones(phonesResp.result, userId));
        })

        client.getUsers().then((usersResp) => renderUsers(usersResp.result));
    });
});