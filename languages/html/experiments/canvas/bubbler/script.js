/** @type {HTMLCanvasElement} */
let canvas;

/** @type {CanvasRenderingContext2D} */
let context;

/** @type {HTMLParagraphElement} */
let info;

/** @type {Array<Particle>} */
let particleArray = [];

class Particle {
    /**
     * 
     * @param {number} x 
     * @param {number} y 
     */
    constructor(x = 0, y = 0) {
        this.x = x;
        this.y = y;
        this.radius = randomNumberMinMax(8, 25);
        this.ttl = randomNumberMinMax(50, 2000);
        this.dx = (1 + Math.random() * 4) * ((Math.floor(Math.random() * 100) > 50) ? 1 : -1);
        this.dy = (1 + Math.random() * 4) * ((Math.floor(Math.random() * 100) > 50) ? 1 : -1);
        this.color = 'white';
    }

    draw() {
        if (this.radius > 0) {
            context.beginPath();
            context.arc(this.x, this.y, this.radius, 0, 2 * Math.PI);
            // context.strokeStyle = `hsl(${this.hue} 100% 50%)`;
            context.stroke();

            const gradient = context.createRadialGradient(
                this.x,
                this.y,
                1,
                this.x + 0.5,
                this.y + 0.5,
                this.radius
            );

            gradient.addColorStop(0.3, "rgba(255, 255, 255, 0.3)");
            gradient.addColorStop(0.95, "#e7feff");

            context.fillStyle = gradient;
            context.fill();
        }
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

        if (this.ttl > 0) {
            this.ttl -= 1;
        } else {
            this.radius -= 1;
        }
    }
}

const resize_primary_div = () => {
    const pd = document.getElementById('primary-div');
    if (pd) {
        pd.style.backgroundColor = 'white';
        pd.style.height = `${window.innerHeight - 20}px`;
        pd.style.width = `${window.innerWidth - 20}px`;
    }
};

/**
 * 
 * @param {number} min - Minimum number required
 * @param {number} max - Maximum number allowed
 * @returns {number} - Between min and max
 */
const randomNumberMinMax = (min, max) =>
    Math.floor(min + Math.random() * (max - min + 1))

const handleDrawCircle = (/** @type {{ offsetX: number | undefined; offsetY: number | undefined; }} */ event) => {
    const particle = new Particle(event.offsetX, event.offsetY);
    particleArray.push(particle);
}

const init_variables = () => {
    // @ts-ignore
    canvas = document.getElementById('main-canvas');
    canvas.width = window.innerWidth / 2;
    canvas.height = window.innerHeight / 2;
    canvas.style.backgroundColor = '#0074aa';
    canvas.addEventListener('click', handleDrawCircle);

    // @ts-ignore
    context = canvas.getContext('2d');

    // @ts-ignore
    info = document.getElementById('info');
}

const animate = () => {
    context.clearRect(0, 0, canvas.width, canvas.height);

    particleArray.forEach((particle) => {
        particle?.update();
        particle?.draw();
    });

    particleArray = particleArray.filter((p) => { return p.radius > 0; })
    info.innerHTML = `Active bubbles: ${particleArray.length}`;
    requestAnimationFrame(animate);
}

window.onload = function () {
    resize_primary_div();
    init_variables();
    animate();
};
