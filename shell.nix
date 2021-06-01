{ pkgs ? import <nixpkgs> {} }:

let
  protobuf = pkgs.protobuf3_8;
in
pkgs.mkShell {
  hardeningDisable = [ "all" ];
  nativeBuildInputs = [ pkgs.cmake ];
  buildInputs = with pkgs;[ go cargo gcc gnumake grpc ] ++ [ protobuf ];
  shellHook = ''
    export PROTOC="${protobuf}/bin/protoc"
  '';
}
