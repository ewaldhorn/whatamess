const canvas = document.getElementById('canvas')
const ctx = canvas.getContext('2d')

canvas.width = innerWidth
canvas.height = innerHeight

let currentStep = 0.01
let update = 0.001

function render() {
    requestAnimationFrame(render)

    ctx.clearRect(0, 0, canvas.width, canvas.height)

    ctx.beginPath()

    ctx.moveTo(0, canvas.height / 2)

    for (let i = 0; i < canvas.width; i++) {
        ctx.lineTo(i, canvas.height / 2 + Math.sin(i * currentStep) * 100)
    }

    ctx.stroke()

    if (currentStep > 0.08 || currentStep < 0.0001) {
        update *= -1
    }
    currentStep += update
}

render()