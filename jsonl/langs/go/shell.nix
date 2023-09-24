{ pkgs ? import <nixpkgs> {} }:

pkgs.mkShell {
  buildInputs = [
    pkgs.go
    pkgs.go-outline
    pkgs.gopls
    pkgs.gopkgs
    pkgs.go-tools
    pkgs.delve
  ];
}
