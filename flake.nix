{
  inputs.flake-utils.url = "github:numtide/flake-utils";

  outputs = { self,  nixpkgs, flake-utils }:
  (flake-utils.lib.eachDefaultSystem
    (system:
      let
        pkgs = import nixpkgs {
          inherit system;
        };
      in
      rec {
        devShells.default = import ./shell.nix { inherit pkgs; };
        packages.default = pkgs.callPackage ./. {};
        packages.container = pkgs.callPackage ./container.nix { package = packages.default; };
    })
  );
}