var JSONRPCClient = (function () {
    function JSONRPCClient() {
        this.requestId = 1;
        this.responses = [];
    }
    JSONRPCClient.prototype.connect = function (url) {
        var _this = this;
        this.client = this;
        this.socket = new WebSocket(url);
        var client = this;
        this.socket.onmessage = function (event) {
            var data = JSON.parse(event.data);
            var requestId = data.id;
            client.responses[requestId] = data;
        };
        this.socket.onclose = function (event) {
            if (event.wasClean) {
                console.log('Clean connection close');
            }
            else {
                console.log('Dicsonnection');
            }
            console.log('Code: ' + event.code + ' cause: ' + event.reason);
        };
        return new Promise(function (resolve, reject) {
            _this.socket.onerror = function (error) { return reject(error); };
            _this.socket.onopen = function () {
                _this.socket.onerror = function (error) {
                    console.log("Error " + error);
                };
                resolve();
            };
        });
    };
    JSONRPCClient.prototype.getUsers = function () {
        return this.sendRequest("UserController.List", JSON.stringify({}));
    };
    JSONRPCClient.prototype.addUser = function (user) {
        return this.sendRequest("UserController.Add", JSON.stringify(user));
    };
    JSONRPCClient.prototype.updateUser = function (user) {
        return this.sendRequest("UserController.Update", JSON.stringify(user));
    };
    JSONRPCClient.prototype.deleteUser = function (id) {
        return this.sendRequest("UserController.Delete", "" + id);
    };
    JSONRPCClient.prototype.getPhones = function (userId) {
        return this.sendRequest("PhoneController.ListByUser", "" + userId);
    };
    JSONRPCClient.prototype.addPhone = function (phone) {
        return this.sendRequest("PhoneController.Add", JSON.stringify(phone));
    };
    JSONRPCClient.prototype.updatePhone = function (phone) {
        return this.sendRequest("PhoneController.Update", JSON.stringify(phone));
    };
    JSONRPCClient.prototype.deletePhone = function (id) {
        return this.sendRequest("PhoneController.Delete", "" + id);
    };
    JSONRPCClient.prototype.sendRequest = function (method, params) {
        var reqId = this.requestId++;
        this.socket.send("{\"jsonrpc\": \"2.0\", \"method\": \"" + method + "\", \"params\": [" + params + "], \"id\": " + reqId + "}");
        return this.pollResponse(reqId);
    };
    JSONRPCClient.prototype.pollResponse = function (requestId) {
        var client = this;
        return new Promise(function (resolve, reject) {
            (function checkResponse() {
                if (client.responses[requestId] != undefined && client.responses[requestId] != null) {
                    resolve(client.responses[requestId]);
                }
                setTimeout(checkResponse, 100);
            })();
        });
    };
    return JSONRPCClient;
}());
//# sourceMappingURL=jsonrpc_client.js.map