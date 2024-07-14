#shell.nix
{ pkgs ? import <nixpkgs> {}}:
with pkgs;
let
   # goEnv = pkgs.mkGoEnv { pwd = ./.; };
in
mkShell {
 packages = [
    figlet
    lolcat
    go
    gotools
    gopls
    go-outline
    gopkgs
    gocode-gomod
    bash
 ];
 shellHook = ''
    export GOPATH="/home/stsully/go"
    export GOMODCACHE="$GOPATH/pkg/mod"
    figlet Welcome to Steven\'s go dev env | lolcat
    go mod tidy
 '';
}
