#version 410

uniform double engineTime;

out vec4 outputColor;

in GS_OUT
{
    vec3 v_FragPos;
} fs_in;

void main() {
    outputColor = vec4(0.0, 0.0, 0.0, 1.0);

    vec3 p = fs_in.v_FragPos.xyz;
    vec3 normal = normalize(
        cross(
            p - dFdx(fs_in.v_FragPos).xyz,
            p - dFdy(fs_in.v_FragPos).xyz
        )
    );
    outputColor.rgb += normal;
    // outputColor.g += 0.3;

    // outputColor.rgb = fs_in.v_FragPos.xyz / 2.0;
    // outputColor.rgb = vec3(1.0);
}
