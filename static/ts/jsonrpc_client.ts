class JSONRPCClient {
    private requestId: number
    private socket: WebSocket
    private responses: any[]
    private client;

    constructor() {
        this.requestId = 1;
        this.responses = [];
    }

    public connect(url: string): Promise<void> {
        this.client = this;
        this.socket = new WebSocket(url);
        let client = this;

        this.socket.onmessage = function (event) {
            let data = JSON.parse(event.data);

            let requestId = data.id;

            client.responses[requestId] = data;
        }

        this.socket.onclose = function (event) {
            if (event.wasClean) {
                console.log('Clean connection close');
            } else {
                console.log('Dicsonnection'); // например, "убит" процесс сервера
            }
            console.log('Code: ' + event.code + ' cause: ' + event.reason);
        };

        return new Promise((resolve, reject) => {
            this.socket.onerror = (error) => reject(error);
            this.socket.onopen = () => {
                this.socket.onerror = function (error) {
                    console.log("Error " + error);
                };
                
                resolve();
            };
        });
    }

    public getUsers(): Promise<any> {
        return this.sendRequest("UserController.List", JSON.stringify({}));
    }

    public addUser(user: any): Promise<any> {
        return this.sendRequest("UserController.Add", JSON.stringify(user));
    }

    public updateUser(user: any): Promise<any> {
        return this.sendRequest("UserController.Update", JSON.stringify(user));
    }

    public deleteUser(id: number): Promise<any> {
        return this.sendRequest("UserController.Delete", "" + id);
    }

    public getPhones(userId: number): Promise<any> {
        return this.sendRequest("PhoneController.ListByUser", "" + userId);
    }

    public addPhone(phone: any): Promise<any> {
        return this.sendRequest("PhoneController.Add", JSON.stringify(phone));
    }

    public updatePhone(phone: any): Promise<any> {
        return this.sendRequest("PhoneController.Update", JSON.stringify(phone));
    }

    public deletePhone(id: number): Promise<any> {
        return this.sendRequest("PhoneController.Delete", "" + id);
    }

    private sendRequest(method: string, params: string): Promise<any> {
        let reqId = this.requestId++;
        this.socket.send(`{"jsonrpc": "2.0", "method": "${method}", "params": [${params}], "id": ${reqId}}`);

        return this.pollResponse(reqId);
    }

    private pollResponse(requestId: number): Promise<any> {
        let client = this;

        return new Promise((resolve, reject) => {
            (function checkResponse() {
                if (client.responses[requestId] != undefined && client.responses[requestId] != null) {
                    resolve(client.responses[requestId]);
                }

                setTimeout(checkResponse, 100);
            })();
        });
    }
}