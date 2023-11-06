const WEBSOCKET = `ws://${config.server.host}:${config.server.port}/bongocat/ws`

function connectWS() {
	let socket = new WebSocket(WEBSOCKET);

	socket.onopen = function (e) {
		console.log(`ws conn open to  ${WEBSOCKET}`);
	};

	socket.onmessage = function (e) {
		msgStruct = JSON.parse(e.data)

		switch (msgStruct.type_of_msg) {
			case "mouse":
				let posX = parseInt(msgStruct.position.x)
				let posY = parseInt(msgStruct.position.y)

				clear();
				drawPaw(posX, posY)

				break;
			case "keyboard":
				console.log(msgStruct.key_board);
				tap(msgStruct.key_board)

				break;
			default:

				break;
		}
	};

	socket.onclose = function (event) {
		if (event.wasClean) {
			console.log(`[close] Соединение закрыто чисто, код=${event.code} причина=${event.reason}`);
		} else {
			// например, сервер убил процесс или сеть недоступна
			// обычно в этом случае event.code 1006
			console.log('[close ws] Соединение прервано');
		}

		setTimeout(function () {
			connectWS();
		}, 1000);
	};

	socket.onerror = function (error) {
		console.log(`[error] ${error}`);
		socket.close();
	};
}
connectWS();