with import <nixpkgs> {};
stdenv.mkDerivation {
  name = "nix-cage-shell";
  buildInputs = [
    go
  ];
  shellHook = ''
    unset GOPATH
  '';
}
