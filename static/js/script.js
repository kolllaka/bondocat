const bongocat = document.querySelector('.bongocat'),
	bongocatLayers = bongocat.querySelectorAll('.bongocat__layer'),
	bongocatUp = bongocat.querySelector('.bongocat__layer-up');

const firstPosPow = {
	x: 130,
	y: 210,
}

const img = new Image();   // Создаём новый объект Image
img.src = './static/img/mouse.png'; // Устанавливаем путь к источнику

img.onload = () => {
	drawPaw(firstPosPow.x, firstPosPow.y)
}

const tap = (button) => {
	const bongocatLayer = bongocat.querySelector(`.bongocat__layer-${button}`)

	if (bongocatLayer) {
		tapFrame(bongocatLayer, "block");
		setTimeout(tapFrame, 100, bongocatLayer, "none");
	}
}

const tapFrame = (node, display) => {
	node.style.display = display;
	bongocatUp.style.display = toggle(display);
}

const toggle = (display = "none") => {
	if (display == "block") {
		return "none"
	}

	return "block"
}