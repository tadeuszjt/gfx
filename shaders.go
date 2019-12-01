package gfx

var shader2D = struct {
	vertex, fragment string
}{
`#version 330 core
uniform mat3 matrix;
in vec2 position;
void main() {
	gl_Position = vec4(matrix * vec3(position, 1), 1.0);
}`,

`#version 330 core
out vec4 outColor;
void main() {
	outColor = vec4(1, 0, 0, 1);
}`,
}
