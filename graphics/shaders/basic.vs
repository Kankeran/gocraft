#version 330 core

layout (location = 0) in vec3 position;
uniform mat4 projection, camera, model;

void main()
{
	gl_Position = projection * camera * model * vec4(position, 1.0);
}
