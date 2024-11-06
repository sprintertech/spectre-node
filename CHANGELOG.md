# Changelog

## [1.0.1](https://github.com/sygmaprotocol/spectre-node/compare/v1.0.0...v1.0.1) (2024-11-06)


### Miscellaneous

* Access Secrets based on segregated environments ([#51](https://github.com/sygmaprotocol/spectre-node/issues/51)) ([177b285](https://github.com/sygmaprotocol/spectre-node/commit/177b28548d3a9a33aa89fe6a9663ea9d5d8da4f2))

## 1.0.0 (2024-10-04)


### Features

* add finality threshold check ([#38](https://github.com/sygmaprotocol/spectre-node/issues/38)) ([5c5086d](https://github.com/sygmaprotocol/spectre-node/commit/5c5086d09e744f4c6036607884f67037083354b8))
* add health port ([#24](https://github.com/sygmaprotocol/spectre-node/issues/24)) ([efb6f65](https://github.com/sygmaprotocol/spectre-node/commit/efb6f65e037641f876a49d4edc5521922b51dbff))
* check hashi message ([#48](https://github.com/sygmaprotocol/spectre-node/issues/48)) ([8babdef](https://github.com/sygmaprotocol/spectre-node/commit/8babdef81fecd44e349cc9bcf8b41a2a2d43ff76))
* configuration setup ([#9](https://github.com/sygmaprotocol/spectre-node/issues/9)) ([435cf46](https://github.com/sygmaprotocol/spectre-node/commit/435cf46044cff862c093dd07bcbeb5c4ca760f08))
* deposit event handler ([#10](https://github.com/sygmaprotocol/spectre-node/issues/10)) ([669d420](https://github.com/sygmaprotocol/spectre-node/commit/669d4206bcd05df1e2a57dfaf6cb28405f70e73e))
* enable execute only networks ([#44](https://github.com/sygmaprotocol/spectre-node/issues/44)) ([4b91b55](https://github.com/sygmaprotocol/spectre-node/commit/4b91b551aa6a6a6873e922df4613dd01d93530af))
* force starting period ([#32](https://github.com/sygmaprotocol/spectre-node/issues/32)) ([d260d8c](https://github.com/sygmaprotocol/spectre-node/commit/d260d8cf509fb9ac41e20e20cf2278c98a21b135))
* integrate basic step flow ([#12](https://github.com/sygmaprotocol/spectre-node/issues/12)) ([e47bd3e](https://github.com/sygmaprotocol/spectre-node/commit/e47bd3ed206d2c73bbded2f7b7144a53a6385676))
* integrate rotate basic flow ([#15](https://github.com/sygmaprotocol/spectre-node/issues/15)) ([c656fe8](https://github.com/sygmaprotocol/spectre-node/commit/c656fe8cd0bc8821b5b1e43a4afdbddda3f36c8e))
* proof generation implementation ([#17](https://github.com/sygmaprotocol/spectre-node/issues/17)) ([025a934](https://github.com/sygmaprotocol/spectre-node/commit/025a9344b30c51ed0c28714b8f17d7e9cb6680a1))
* reflect changes in Spectre Prover and contracts ([#37](https://github.com/sygmaprotocol/spectre-node/issues/37)) ([531134c](https://github.com/sygmaprotocol/spectre-node/commit/531134c9ea807bd95a9d5a47d4ae5011042ddd96))
* send step only for filled epochs ([#33](https://github.com/sygmaprotocol/spectre-node/issues/33)) ([bc2f287](https://github.com/sygmaprotocol/spectre-node/commit/bc2f287168486f9da450f9e29150067a02daaf8a))
* spectre proxy implementation ([#30](https://github.com/sygmaprotocol/spectre-node/issues/30)) ([5d6a09c](https://github.com/sygmaprotocol/spectre-node/commit/5d6a09c61799ab5ca9007b56948dae3aa49351b8))
* store rotation period ([#26](https://github.com/sygmaprotocol/spectre-node/issues/26)) ([19a54c0](https://github.com/sygmaprotocol/spectre-node/commit/19a54c034964fff3d00ae327d3f339ecb17569e5))
* use finalized next sync committee branch ([#23](https://github.com/sygmaprotocol/spectre-node/issues/23)) ([b12fc2b](https://github.com/sygmaprotocol/spectre-node/commit/b12fc2b69eebfc1295ec4c97fdfc4bc251480347))


### Bug Fixes

* deneb hard fork  ([#35](https://github.com/sygmaprotocol/spectre-node/issues/35)) ([9dc5782](https://github.com/sygmaprotocol/spectre-node/commit/9dc5782ac35f5b00708d4db9fef599747022c3bf))
* rotate when current period the same as the latest period ([#28](https://github.com/sygmaprotocol/spectre-node/issues/28)) ([3dd7c41](https://github.com/sygmaprotocol/spectre-node/commit/3dd7c416ecb48f627132f5a2cedac7558d6ed5a4))
* setup slot number per network ([#43](https://github.com/sygmaprotocol/spectre-node/issues/43)) ([5c81f6e](https://github.com/sygmaprotocol/spectre-node/commit/5c81f6ebd1970b7d2a0abfdeb5da67c42718e654))
* starting period 1 when period does not exist ([#27](https://github.com/sygmaprotocol/spectre-node/issues/27)) ([8c4efb1](https://github.com/sygmaprotocol/spectre-node/commit/8c4efb155fb52b3011daf563d1294a5649d94fda))
* switch block to deneb ([#36](https://github.com/sygmaprotocol/spectre-node/issues/36)) ([93e8d19](https://github.com/sygmaprotocol/spectre-node/commit/93e8d1999b841462555bbff4c93593da0c5fd19a))
* update log so it says for what period are we rotating ([#29](https://github.com/sygmaprotocol/spectre-node/issues/29)) ([980d6b5](https://github.com/sygmaprotocol/spectre-node/commit/980d6b57f8a2d119a85fb6684425bd394a2bad04))
* wait for db lock ([#31](https://github.com/sygmaprotocol/spectre-node/issues/31)) ([9f9e048](https://github.com/sygmaprotocol/spectre-node/commit/9f9e048d35e2e3fd950657ef20343afda5f1778c))


### Miscellaneous

* add message dispatched log ([#49](https://github.com/sygmaprotocol/spectre-node/issues/49)) ([e8701fa](https://github.com/sygmaprotocol/spectre-node/commit/e8701fac791fb67a88773e2a56b64d31681979dc))
* Added deployment pipeline for mainnet ([#45](https://github.com/sygmaprotocol/spectre-node/issues/45)) ([7b6e21b](https://github.com/sygmaprotocol/spectre-node/commit/7b6e21b3fc674186d4b352f0fdbe30c020efae5e))
* added pipeline upgrade & image version visibility ([#40](https://github.com/sygmaprotocol/spectre-node/issues/40)) ([99770e5](https://github.com/sygmaprotocol/spectre-node/commit/99770e551a7a92bee37a126f8aa3d74e6a67d700))
* bump eth2 client dependecy ([#41](https://github.com/sygmaprotocol/spectre-node/issues/41)) ([1c1a02b](https://github.com/sygmaprotocol/spectre-node/commit/1c1a02be0667b2be41f80f36eec1c7f53a9e8a84))
* CICD Pipelines ([#20](https://github.com/sygmaprotocol/spectre-node/issues/20)) ([2119b87](https://github.com/sygmaprotocol/spectre-node/commit/2119b871586fad57ab53bb70afdee1ef15d67a03))
* fix mainnet service name ([#50](https://github.com/sygmaprotocol/spectre-node/issues/50)) ([0f1bb8c](https://github.com/sygmaprotocol/spectre-node/commit/0f1bb8cead13569f5a6b2a0d0f36925a9f6de373))
* Initial setup ([#8](https://github.com/sygmaprotocol/spectre-node/issues/8)) ([eb0acb4](https://github.com/sygmaprotocol/spectre-node/commit/eb0acb4d808de402611ac55197ed41a35defaeb6))
