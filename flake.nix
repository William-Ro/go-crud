{
  description = "Dev Shell for Go Development";

  inputs.nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";

  outputs = {
    self,
    nixpkgs,
  }: let
    system = "x86_64-linux"; # Or "aarch64-linux" for ARM
    pkgs = import nixpkgs {
      inherit system;
      #config.allowUnfree = true; # Uncomment to allow unfree packages
    };
  in {
    devShells.${system}.default = pkgs.mkShell {
      buildInputs = with pkgs; [
        go
        air
        golangci-lint
        mongosh
      ];

      shellHook = ''
        echo "Welcome to the Go Development Shell 🚀"
      '';
    };
  };
}
