let canvas, context;
const particleArray = [];

class Particle {
    constructor(x = 0, y = 0) {
        this.x = x;
        this.y = y;
        this.radius = randomNumberMinMax(10, 20);
        this.dx = (1 + Math.random() * 4) * (Math.floor(Math.random() * 100) > 50) ? 1 : -1;
        this.dy = (1 + Math.random() * 4) * (Math.floor(Math.random() * 100) > 50) ? 1 : -1;
        this.color = 'white';
    }

    draw() {
        context.beginPath();
        context.arc(this.x, this.y, this.radius, 0, 2 * Math.PI);
        context.strokeStyle = this.color;
        context.stroke();

        context.fillStyle = this.color;
        context.fill();
    }

    update() {
        this.x = this.x + this.dx;
        this.y = this.y - this.dy;

        if (this.x < 20 || this.x > (canvas.width - 20)) {
            this.dx *= -1;
        }

        if (this.y < 20 || this.y > (canvas.height - 20)) {
            this.dy *= -1;
        }
    }
}

const resize_primary_div = () => {
    const pd = document.getElementById('primary-div');
    pd.style.backgroundColor = 'white';
    pd.style.height = `${window.innerHeight - 20}px`;
    pd.style.width = `${window.innerWidth - 20}px`;
};


const randomNumberMinMax = (min, max) =>
    Math.floor(min + Math.random() * (max - min + 1))


const drawCircle = (x, y) => {
    context.strokeStyle = 'white';

    context.beginPath();
    context.arc(x, y, randomNumberMinMax(10, 20), 0, 2 * Math.PI);
    context.stroke();
};


const handleDrawCircle = (event) => {
    const particle = new Particle(event.offsetX, event.offsetY);
    particleArray.push(particle);
}

const init_variables = () => {
    canvas = document.getElementById('main-canvas');
    canvas.width = 640;
    canvas.height = 480;
    context = canvas.getContext('2d');

    canvas.style.backgroundColor = '#00b4ff';
    canvas.addEventListener('click', handleDrawCircle);
}

const animate = () => {
    context.clearRect(0, 0, canvas.width, canvas.height);

    particleArray.forEach((particle) => {
        particle?.update();
        particle?.draw();
    });

    requestAnimationFrame(animate);
}

window.onload = function () {
    resize_primary_div();
    init_variables();
    animate();
};
