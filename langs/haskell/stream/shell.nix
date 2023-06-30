with import <nixpkgs> { };

let
  pkgs = import <nixpkgs> { };
  haskellPackages = pkgs.haskellPackages;


in
haskell.lib.buildStackProject {
  name = "myenv";

  buildInputs = [
    zlib
    ghc
    haskellPackages.haskell-language-server
    cabal-install
    gmp
    # avoid on circleci: commitBuffer: invalid argument (invalid character)
    glibcLocales
  ];

  # avoid on circleci: commitBuffer: invalid argument (invalid character)
  LANG = "en_US.utf8";
}
