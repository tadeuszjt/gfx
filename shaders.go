package gfx

import (
	"github.com/faiface/glhf"
)

type GLShader struct {
	vertexSrc, fragmentSrc string
	vertexFmt, uniformFmt  glhf.AttrFormat
}

func newShader(s *GLShader) (*glhf.Shader, error) {
	return glhf.NewShader(s.vertexFmt, s.uniformFmt, s.vertexSrc, s.fragmentSrc)
}

var shader2D = GLShader{
	`#version 330 core
uniform mat3 matrix;
in vec2 position;
in vec2 texCoord;
in vec4 colour;
out vec2 TexCoord;
out vec4 Colour;
void main() {
	gl_Position = vec4(matrix * vec3(position, 1), 1.0);
	TexCoord = texCoord;
	Colour = colour;
}`,
	`#version 330 core
uniform sampler2D tex;
in vec2 TexCoord;
in vec4 Colour;
out vec4 outColor;
void main() {
	outColor = Colour * texture(tex, TexCoord);
}`,
	glhf.AttrFormat{
		{Name: "position", Type: glhf.Vec2},
		{Name: "texCoord", Type: glhf.Vec2},
		{Name: "colour", Type: glhf.Vec4},
	},
	glhf.AttrFormat{{Name: "matrix", Type: glhf.Mat3}},
}
