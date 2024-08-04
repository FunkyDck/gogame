#version 410

uniform double engineTime;
uniform sampler2D tex;

in vec2 fragTexCoord;

out vec4 outputColor;

void main() {
    outputColor = vec4(0.1, 1.0, 0.0, 1.0);
}
