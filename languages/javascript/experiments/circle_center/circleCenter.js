class CircleCenterFinder {
    constructor(canvasId) {
        this.canvas = document.getElementById(canvasId);
        this.ctx = this.canvas.getContext("2d");

        this.canvas.width = 500;
        this.canvas.height = 500;

        this.radius = 100;
        this.centerX = this.canvas.width / 2;
        this.centerY = this.canvas.height / 2;
        this.angleStep = Math.PI / 6;
        this.currentAngle = 0;
        this.triangleSize = 80;
        this.animationRunning = false;

        this.canvas.addEventListener("click", (e) => this.setNewCircle(e));
        this.init();
    }

    init() {
        this.clearCanvas();
        this.drawCircle();
        this.animateTriangle();
    }

    clearCanvas() {
        this.ctx.clearRect(0, 0, this.canvas.width, this.canvas.height);
    }

    drawCircle() {
        this.ctx.beginPath();
        this.ctx.arc(this.centerX, this.centerY, this.radius, 0, Math.PI * 2);
        this.ctx.strokeStyle = "#000";
        this.ctx.lineWidth = 2;
        this.ctx.stroke();
    }

    drawTriangle(position) {
        const { x, y, angle } = position;

        const x1 = x;
        const y1 = y;
        const x2 = x + this.triangleSize * Math.cos(angle);
        const y2 = y + this.triangleSize * Math.sin(angle);
        const x3 = x + this.triangleSize * Math.cos(angle + Math.PI / 2);
        const y3 = y + this.triangleSize * Math.sin(angle + Math.PI / 2);

        this.ctx.beginPath();
        this.ctx.moveTo(x1, y1);
        this.ctx.lineTo(x2, y2);
        this.ctx.lineTo(x3, y3);
        this.ctx.closePath();

        this.ctx.strokeStyle = "#ff0000";
        this.ctx.lineWidth = 2;
        this.ctx.stroke();
    }

    animateTriangle() {
        if (!this.animationRunning) {
            this.animationRunning = true;
            this.currentAngle = 0;
            this.animationLoop();
        }
    }

    animationLoop() {
        if (this.currentAngle < Math.PI * 2) {
            this.clearCanvas();
            this.drawCircle();

            let x = this.centerX + this.radius * Math.cos(this.currentAngle);
            let y = this.centerY + this.radius * Math.sin(this.currentAngle);

            this.drawTriangle({ x, y, angle: this.currentAngle });

            this.currentAngle += this.angleStep;
            requestAnimationFrame(() => this.animationLoop());
        } else {
            this.drawCenterPoint();
            this.animationRunning = false;
        }
    }

    drawCenterPoint() {
        this.ctx.beginPath();
        this.ctx.arc(this.centerX, this.centerY, 5, 0, Math.PI * 2);
        this.ctx.fillStyle = "blue";
        this.ctx.fill();
    }

    setNewCircle(event) {
        if (!this.animationRunning) {
            const rect = this.canvas.getBoundingClientRect();
            this.centerX = event.clientX - rect.left;
            this.centerY = event.clientY - rect.top;
            this.radius = Math.random() * 60 + 50;
            this.init();
        }
    }
}

// Function to reset the animation
function resetCircle() {
    new CircleCenterFinder("circleCanvas");
}

// Run when the page loads
document.addEventListener("DOMContentLoaded", () => {
    resetCircle();
});
