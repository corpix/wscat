with import <nixpkgs>{};
{ pkgs ? import <nixpkgs> {} }:

buildGoPackage rec {
  name = "wscat-${version}";
  version = "development";

  src = ./.;
  goDeps = ./deps.nix;
  goPackagePath = "github.com/corpix/wscat";
}
