let config;

const initConfig = () => {
	const url = "./static/config/config.json";

	fetch(url).then((resp) => {
		resp.json().then((json) => {
			console.log("[json] ", json);
			config = json
			const colorPaw = "#010101";

			loadScript("./static/js/script.js");
			loadScript("./static/js/draw.js");
			loadScript("./static/js/ws.js");
		})
	})
}
initConfig()


const loadScript = (url) => {
	let tag = document.createElement('script');
	tag.src = url;
	document.body.insertAdjacentElement("afterEnd", tag);
}