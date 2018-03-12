with import <nixpkgs>{};
{ pkgs ? import <nixpkgs> {} }:

buildGo19Package rec {
  name = "wscp-unstable-${version}";
  version = "development";

  buildInputs = with pkgs; [ git glide ];

  src = ./.;
  goPackagePath = "github.com/corpix/wscp";
}
