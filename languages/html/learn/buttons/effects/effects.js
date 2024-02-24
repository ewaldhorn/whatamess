const allButtons = document.querySelectorAll("button");
const turbulence = document.querySelector("feTurbulence");
let verticalFrequency = 0.01;

turbulence.setAttribute("baseFrequency", verticalFrequency + " 0.00001");

const steps = 30, interval = 10;

allButtons.forEach(b => {
    b.addEventListener("mouseover", () => {
        verticalFrequency = 0.01;

        for (let i = 0; i < steps; i++) {
            setTimeout(() => {
                verticalFrequency += 0.002;
                turbulence.setAttribute("baseFrequency", verticalFrequency + " 0.00001");
            }, i * interval);
        }

        // smooth out
        setTimeout(() => {
            verticalFrequency = 0.00001;
            turbulence.setAttribute("baseFrequency", verticalFrequency + " 0.00001");
        }, steps * interval);

    });
});