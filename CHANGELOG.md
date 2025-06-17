# Changelog

## [0.3.0](https://github.com/nobbs/terraform-provider-sops/compare/v0.2.1...v0.3.0) (2025-06-17)


### Features

* implement decrypt functions that ignore MAC mismatches ([#32](https://github.com/nobbs/terraform-provider-sops/issues/32)) ([fc5352f](https://github.com/nobbs/terraform-provider-sops/commit/fc5352f2b93c180a79f86963ebce8d3c630e5612))


### Bug Fixes

* **deps:** update dependencies ([ff115d4](https://github.com/nobbs/terraform-provider-sops/commit/ff115d44b4750e250ac0730e87036bc1f43bd9bd))


### Miscellaneous Chores

* **deps:** bump github.com/cli/go-gh/v2 in /tools ([#29](https://github.com/nobbs/terraform-provider-sops/issues/29)) ([2e7c5de](https://github.com/nobbs/terraform-provider-sops/commit/2e7c5deb24899797f52639ad28d16b2222d9e1ba))
* **deps:** bump github.com/hashicorp/terraform-plugin-go ([#24](https://github.com/nobbs/terraform-provider-sops/issues/24)) ([eba27ac](https://github.com/nobbs/terraform-provider-sops/commit/eba27ac066e651aec98e0b9a9d2b755d6136d6f4))
* **deps:** bump github.com/hashicorp/terraform-plugin-testing ([#25](https://github.com/nobbs/terraform-provider-sops/issues/25)) ([0db9926](https://github.com/nobbs/terraform-provider-sops/commit/0db9926fc970d8784f766754189e4b6b16e115f5))
* **deps:** bump github.com/hashicorp/terraform-plugin-testing ([#30](https://github.com/nobbs/terraform-provider-sops/issues/30)) ([4ffa0f6](https://github.com/nobbs/terraform-provider-sops/commit/4ffa0f6c64b106ec38eec3dbef0eb22a451e2f7a))


### Continuous Integration

* **deps:** add dependabot auto merge workflow ([2e63e73](https://github.com/nobbs/terraform-provider-sops/commit/2e63e73eaf9815f1a4b992b5759f003d0f73eaf7))
* **deps:** change Dependabot auto-merge strategy to squash ([c0e32df](https://github.com/nobbs/terraform-provider-sops/commit/c0e32df7b538641a729275054e15a5218b29803d))
* **deps:** fix dependabot auto merge workflow ([#28](https://github.com/nobbs/terraform-provider-sops/issues/28)) ([d9496cd](https://github.com/nobbs/terraform-provider-sops/commit/d9496cd7b490b18f672038f20856a90145612b17))

## [0.2.1](https://github.com/nobbs/terraform-provider-sops/compare/v0.2.0...v0.2.1) (2025-05-21)


### Bug Fixes

* **deps:** update github.com/hashicorp/terraform-plugin-go to v0.26.0 ([8bcbcb5](https://github.com/nobbs/terraform-provider-sops/commit/8bcbcb5425d0d7e4dd672febdbd64554b4374b8d))
* **deps:** update module github.com/getsops/sops/v3 to v3.9.4 ([5586e7a](https://github.com/nobbs/terraform-provider-sops/commit/5586e7a0c72566f7babb4aa2ed0ad0c5bd3e91bf))
* **provider:** remove obsolete provider model ([fbce531](https://github.com/nobbs/terraform-provider-sops/commit/fbce531aca644bc4ea9a8d2f7ff75adaba43eb54))


### Miscellaneous Chores

* **dependabot:** remove obsolete allow rule for hashicorp dependencies ([83a49e4](https://github.com/nobbs/terraform-provider-sops/commit/83a49e4332be31dd6b6a198661456a0926720398))
* **deps:** bump actions/create-github-app-token from 1 to 2 ([4a01d79](https://github.com/nobbs/terraform-provider-sops/commit/4a01d79eb170fffdd1b4a470faa3429af71afd62))
* **deps:** bump actions/setup-go from 5.1.0 to 5.5.0 ([f79236c](https://github.com/nobbs/terraform-provider-sops/commit/f79236cd11da64267c03e8462687dc296e9418b3))
* **deps:** bump crazy-max/ghaction-import-gpg from 6.2.0 to 6.3.0 ([9eff39a](https://github.com/nobbs/terraform-provider-sops/commit/9eff39a4961b772d59f2ec343061a754f2a964c0))
* **deps:** bump github.com/getsops/sops/v3 from 3.9.1 to 3.9.2 ([#6](https://github.com/nobbs/terraform-provider-sops/issues/6)) ([7f12667](https://github.com/nobbs/terraform-provider-sops/commit/7f12667cde7de5b6198c0344288c70631f13ab75))
* **deps:** bump github.com/getsops/sops/v3 from 3.9.4 to 3.10.2 ([fccabc0](https://github.com/nobbs/terraform-provider-sops/commit/fccabc06016b9f55da6609488b41122e0b8f731a))
* **deps:** bump github.com/hashicorp/terraform-plugin-framework ([ec55c6d](https://github.com/nobbs/terraform-provider-sops/commit/ec55c6d0cb96602d77e9315d8bac943aabe3cbdf))
* **deps:** bump github.com/hashicorp/terraform-plugin-go ([2e061b0](https://github.com/nobbs/terraform-provider-sops/commit/2e061b0a66494a00c39fb1008151fb76149b36dc))
* **deps:** bump github.com/hashicorp/terraform-plugin-testing ([9f89f41](https://github.com/nobbs/terraform-provider-sops/commit/9f89f417083f2ea0ec73e1d88a01a89be362008e))
* **deps:** bump golangci/golangci-lint-action from 6.1.1 to 8.0.0 ([bbd28cf](https://github.com/nobbs/terraform-provider-sops/commit/bbd28cf4ee76084aa59782ee1c0d264f6ee6f77a))
* **deps:** bump goreleaser/goreleaser-action from 6.1.0 to 6.3.0 ([664fef3](https://github.com/nobbs/terraform-provider-sops/commit/664fef3bcb97b181b598d830b13b36ebec4b556a))
* **deps:** update indirect dependencies to latest versions ([72612d9](https://github.com/nobbs/terraform-provider-sops/commit/72612d9bd3faae6870d55b80bfc61f832a80d229))
* **go:** update Go version from 1.22.7 to 1.24.3 in tools and from 1.23.6 to 1.24.3 in main module ([f6fd560](https://github.com/nobbs/terraform-provider-sops/commit/f6fd56094e72f48615f917c4af5d7c95706fde1c))
* **go:** update Go version from 1.23 to 1.24 in mise.toml ([4788573](https://github.com/nobbs/terraform-provider-sops/commit/4788573a4725400cd4a87a7664e2e76373e80560))
* **go:** update Go version from 1.23.3 to 1.23.6 ([a8c7a82](https://github.com/nobbs/terraform-provider-sops/commit/a8c7a82fb5a536bcdf03c8f87c05185e54cd1413))
* **lint:** update golangci-lint configuration to v2 ([9293e7d](https://github.com/nobbs/terraform-provider-sops/commit/9293e7d5b5828a32136fe5a05b75772a83e4e6ef))
* update indirect dependencies and tools ([730d094](https://github.com/nobbs/terraform-provider-sops/commit/730d094c026a58fb64186c70a2be71d83944e927))


### Continuous Integration

* **release:** add changelog sections to release-please configuration ([0b242a9](https://github.com/nobbs/terraform-provider-sops/commit/0b242a9a0ab0167f7015fe0396740459b3c725f0))
* **tests:** add support for Terraform versions 1.11.* and 1.12.* ([95f52a3](https://github.com/nobbs/terraform-provider-sops/commit/95f52a3f35771b0c30dd8a4713ade18050e457e7))

## [0.2.0](https://github.com/nobbs/terraform-provider-sops/compare/v0.1.0...v0.2.0) (2024-11-30)


### Features

* add string function to read and decrypt sops encrypted strings ([#4](https://github.com/nobbs/terraform-provider-sops/issues/4)) ([b4210b7](https://github.com/nobbs/terraform-provider-sops/commit/b4210b7404db349a324f2dd606da62f5c82bcfcd))

## [0.1.0](https://github.com/nobbs/terraform-provider-sops/compare/v0.0.1...v0.1.0) (2024-11-30)


### Features

* add format parameter to SopsFile function and implement dynamic type conversion ([0198bd5](https://github.com/nobbs/terraform-provider-sops/commit/0198bd5cfc1283110a79dd96d6abb17b334fe7e1))
* add support for dotenv format and implement format validation in SopsFile function ([fb95e8f](https://github.com/nobbs/terraform-provider-sops/commit/fb95e8fdd147f2d983d37f9db8bbcfc5a1643780))
* add support for multiple formats in SopsFile function and include test fixtures ([7ae5d06](https://github.com/nobbs/terraform-provider-sops/commit/7ae5d06d2f3cbf2d0002a2ec0985c804213997e6))
* update provider module path and implement SopsFile function ([d6c2feb](https://github.com/nobbs/terraform-provider-sops/commit/d6c2feb95176e49d48d8ce2c4cb1d743301d080d))
