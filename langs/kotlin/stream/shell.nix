{ pkgs ? import <nixpkgs> {} }:

pkgs.mkShell {
  buildInputs = [
    pkgs.kotlin
    pkgs.jdk11
    pkgs.maven
    pkgs.gradle
    pkgs.kotlin-language-server
  ];

  shellHook = ''
    export JAVA_HOME=${pkgs.openjdk11}
  '';
}
