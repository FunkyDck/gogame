#version 410

uniform mat4 projection;
uniform mat4 camera;
uniform mat4 model;

layout (points) in;
layout (triangle_strip, max_vertices = 14) out;

out GS_OUT
{
    vec3 v_FragPos;
} gs_out;

const float cubeVertices[] = float[](
    +0.5, +0.5, -0.5, // Back-top-right
    -0.5, +0.5, -0.5, // Back-top-left
    +0.5, -0.5, -0.5, // Back-bottom-right
    -0.5, -0.5, -0.5, // Back-bottom-left
    -0.5, -0.5, +0.5, // Front-bottom-left
    -0.5, +0.5, -0.5, // Back-top-left
    -0.5, +0.5, +0.5, // Front-top-left
    +0.5, +0.5, -0.5, // Back-top-right
    +0.5, +0.5, +0.5, // Front-top-right
    +0.5, -0.5, -0.5, // Back-bottom-right
    +0.5, -0.5, +0.5, // Front-bottom-right
    -0.5, -0.5, +0.5, // Front-bottom-left
    +0.5, +0.5, +0.5, // Front-top-right
    -0.5, +0.5, +0.5  // Front-top-left
);

void main() {
    int N = cubeVertices.length();
    for (int i = 0; i <= N; i += 3) {
        vec4 vertex = vec4(
            cubeVertices[i + 0],
            cubeVertices[i + 1],
            cubeVertices[i + 2],
            1.0
        );
        vertex.xyz += gl_in[0].gl_Position.xyz;
        gl_Position = projection * camera * model * vertex;
        gs_out.v_FragPos = (vertex).xyz;
        EmitVertex();
    }
    EndPrimitive();
}

