{ 
  pkgs ? import <nixpkgs> {},
  unstable-pkgs ? import <nixos-unstable> { config = { allowUnfree = true; }; } 
}:

pkgs.mkShell {
  buildInputs = [
    unstable-pkgs.nushell
  ];
}
