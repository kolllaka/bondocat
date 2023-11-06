const canvas = document.getElementById('canvas'),
	canvasX = 612,
	canvasY = 354;

// radius of paw
const radius = 30;
// coordinates of 1 and 2 point of arm (x1, y1) and (x2, y2)
const x1 = 177,
	y1 = 105,
	x2 = 228,
	y2 = 169;

let ctx;
if (canvas.getContext) {
	ctx = canvas.getContext('2d');
	canvas.width = canvasX;
	canvas.height = canvasY;
	ctx.lineWidth = 6;
}

function clear() {
	ctx.clearRect(0, 0, canvasX, canvasY);
}

function drawPaw(x, y) {
	clear();
	ctx.drawImage(img, x - 26, y - 17, 104, 68); //рисуем картинку в канвас

	const kasatel1 = (xt, yt, brnch) => {
		const L2 = dist((xt - x), (yt - y))
		const S2 = Math.sqrt((Math.pow(L2, 2) - Math.pow(radius, 2)))

		sin1 = radius / L2
		atg2 = (Math.atan2((y - yt), (x - xt)))
		deg = brnch * Math.asin(sin1) + atg2
		sin = Math.sin(deg)
		cos = Math.cos(deg)

		return { x: (xt + (cos * S2)), y: (yt + (sin * S2)) }
	}

	const dist = (x, y) => {
		return Math.sqrt(Math.pow(x, 2) + Math.pow(y, 2))
	}

	let leftX = kasatel1(x1, y1, 1).x
	let leftY = kasatel1(x1, y1, 1).y

	let leftControlDist = dist((x1 - leftX), (y1 - leftY))

	let leftControlX1 = x1 - 1.2 * leftControlDist / 4
	let leftControlY1 = y1 + leftControlDist / 4

	let leftControlX2 = x1 - .7 * leftControlDist * 1 / 2
	let leftControlY2 = y1 + .9 * leftControlDist * 1 / 2

	let rightX = kasatel1(x2, y2, -1).x
	let rightY = kasatel1(x2, y2, -1).y

	ctx.beginPath();

	ctx.moveTo(x1, y1);
	ctx.bezierCurveTo(leftControlX1, leftControlY1, leftControlX2, leftControlY2, leftX, leftY);

	angle1 = (Math.atan2(leftY - y, leftX - x))
	angle2 = (Math.atan2(rightY - y, rightX - x))
	ctx.arc(x, y, radius, angle1, angle2, true)

	ctx.lineTo(x2, y2);
	ctx.stroke();

	ctx.moveTo(x1, y1)
	ctx.fillStyle = config.paw.color;
	ctx.fill();

	ctx.stroke()
}


