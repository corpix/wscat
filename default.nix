with import <nixpkgs>{};
{ pkgs ? import <nixpkgs> {} }:

buildGoPackage rec {
  name = "wscat-${version}";
  version = "development";

  buildInputs = with pkgs; [ git glide ];

  src = ./.;
  goPackagePath = "github.com/corpix/wscat";
}
