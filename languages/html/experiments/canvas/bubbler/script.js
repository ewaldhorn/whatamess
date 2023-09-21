let canvas, context, x, y;

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
    x = event.offsetX;
    y = event.offsetY;

    drawCircle(x, y);
}

const init_variables = () => {
    canvas = document.getElementById('main-canvas');
    canvas.width = 640;
    canvas.height = 480;
    context = canvas.getContext('2d');

    canvas.style.backgroundColor = '#00b4ff';
    canvas.addEventListener('click', handleDrawCircle);
}

window.onload = function () {
    resize_primary_div();
    init_variables();
};
