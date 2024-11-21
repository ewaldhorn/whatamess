attribute vec3 coordinates;
uniform mat4 modelview;

void main(void) {
    gl_Position = modelview * vec4(coordinates, 1.0);
}
