{ pkgs ? import <nixpkgs> {} }:
with pkgs; buildGoPackage rec {
  name = "wscat-${version}";
  version = "development";

  src = ./.;
  goDeps = ./deps.nix;
  goPackagePath = "github.com/corpix/wscat";
}
