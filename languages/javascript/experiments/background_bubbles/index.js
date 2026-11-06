let max = 60;
let container = document.querySelector(".bubbles");

for (let i = 0; i < max; i++) {
    const bubble = document.createElement("span");

    let size = 4 + 14 * Math.random();
    let position_x = 100 * Math.random();
    let visibility = 0.08 + 0.22 * Math.random();
    let anim_duration = 16 + 18 * Math.random();
    let anim_delay = -1 * (Math.random() * anim_duration);
    let sway_direction = Math.random() > 0.5 ? 1 : -1;
    let distance = sway_direction * (16 + 28 * Math.random());

    bubble.style.setProperty("--size", size + "px");
    bubble.style.setProperty("--position-x", position_x + "%");
    bubble.style.setProperty("--visibility", visibility);
    bubble.style.setProperty("--anim-duration", anim_duration + "s");
    bubble.style.setProperty("--anim-delay", anim_delay + "s");
    bubble.style.setProperty("--distance", distance + "px");

    container.append(bubble);
}

