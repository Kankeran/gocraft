#version 330 core

layout (location = 0) in vec3 position;
layout (location = 1) in vec2 texCoord;
uniform mat4 projection, camera, model;

out vec2 textureCoord;

void main()
{
	gl_Position = projection * camera * model * vec4(position, 1.0);
	textureCoord = texCoord;
}