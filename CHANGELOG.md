# Changelog

## 0.32.0 (2024-03-27)

Full Changelog: [v0.31.0...v0.32.0](https://github.com/Increase/increase-go/compare/v0.31.0...v0.32.0)

### Features

* **api:** add `funding` parameter to external account update ([#236](https://github.com/Increase/increase-go/issues/236)) ([8ef5546](https://github.com/Increase/increase-go/commit/8ef5546e980a89ecc44b675936a030c76035728f))


### Documentation

* fix typo in docstring for Null() ([#235](https://github.com/Increase/increase-go/issues/235)) ([5d48b33](https://github.com/Increase/increase-go/commit/5d48b33364994831da615cdb6f4d675dd9883e8f))
* **readme:** document file uploads ([#233](https://github.com/Increase/increase-go/issues/233)) ([c68a1ee](https://github.com/Increase/increase-go/commit/c68a1eeef60d394f9298a38275d3270370dfdef4))

## 0.31.0 (2024-03-25)

Full Changelog: [v0.30.0...v0.31.0](https://github.com/Increase/increase-go/compare/v0.30.0...v0.31.0)

### Features

* **api:** updates ([#231](https://github.com/Increase/increase-go/issues/231)) ([857611e](https://github.com/Increase/increase-go/commit/857611e6e9119b7ea39e873072d9c1e06b855785))

## 0.30.0 (2024-03-21)

Full Changelog: [v0.29.1...v0.30.0](https://github.com/Increase/increase-go/compare/v0.29.1...v0.30.0)

### Features

* add IsKnown method to enums ([#228](https://github.com/Increase/increase-go/issues/228)) ([ea1efeb](https://github.com/Increase/increase-go/commit/ea1efebce831e31cb10390e493167eaffc57b1ea))
* **api:** adding `pending_reviewing` wire transfer state ([32ca69a](https://github.com/Increase/increase-go/commit/32ca69a78443609c3506a0c095a73b91319052a6))
* **api:** introduce `network_risk_score` ([#230](https://github.com/Increase/increase-go/issues/230)) ([32ca69a](https://github.com/Increase/increase-go/commit/32ca69a78443609c3506a0c095a73b91319052a6))
* **api:** remove Card Profile ID properties ([32ca69a](https://github.com/Increase/increase-go/commit/32ca69a78443609c3506a0c095a73b91319052a6))


### Chores

* **internal:** update generated pragma comment ([#227](https://github.com/Increase/increase-go/issues/227)) ([f9c18bc](https://github.com/Increase/increase-go/commit/f9c18bc921e01daeabc8c160f45cafe77b58aab1))


### Documentation

* fix typo in CONTRIBUTING.md ([#225](https://github.com/Increase/increase-go/issues/225)) ([ada1f05](https://github.com/Increase/increase-go/commit/ada1f057e4ef47b62d8ee5170ad2ee9509fb88c2))
* **readme:** consistent use of sentence case in headings ([#229](https://github.com/Increase/increase-go/issues/229)) ([c598dce](https://github.com/Increase/increase-go/commit/c598dce5c25d598e98bf0b1745e8068938ea58d2))

## 0.29.1 (2024-03-18)

Full Changelog: [v0.29.0...v0.29.1](https://github.com/Increase/increase-go/compare/v0.29.0...v0.29.1)

### Bug Fixes

* **api:** improve union handling for request bodies ([#223](https://github.com/Increase/increase-go/issues/223)) ([e07dc1f](https://github.com/Increase/increase-go/commit/e07dc1fb77f0e05d7b294ce4410a1694e4659c7b))

## 0.29.0 (2024-03-13)

Full Changelog: [v0.28.1...v0.29.0](https://github.com/Increase/increase-go/compare/v0.28.1...v0.29.0)

### Features

* **api:** confirm entities ([#222](https://github.com/Increase/increase-go/issues/222)) ([94f8ab1](https://github.com/Increase/increase-go/commit/94f8ab1a3c8a93fb595fd8e370f7a11d5e4308d4))
* set user-agent header by default when making requests ([#220](https://github.com/Increase/increase-go/issues/220)) ([36a642a](https://github.com/Increase/increase-go/commit/36a642adc62183132efe2b3fceed9264e31e6cf3))

## 0.28.1 (2024-03-12)

Full Changelog: [v0.28.0...v0.28.1](https://github.com/Increase/increase-go/compare/v0.28.0...v0.28.1)

### Bug Fixes

* **client:** don't include ? in path unless necessary ([#218](https://github.com/Increase/increase-go/issues/218)) ([4028846](https://github.com/Increase/increase-go/commit/40288461e318704c31e0b74ddd74eabd0eb3093b))

## 0.28.0 (2024-03-12)

Full Changelog: [v0.27.0...v0.28.0](https://github.com/Increase/increase-go/compare/v0.27.0...v0.28.0)

### Features

* implement public RawJSON of response structs ([#213](https://github.com/Increase/increase-go/issues/213)) ([098706e](https://github.com/Increase/increase-go/commit/098706e2dca0e9000b6f02684c72373325c80dd3))


### Bug Fixes

* fix String() behavior of param.Field ([#217](https://github.com/Increase/increase-go/issues/217)) ([b01ff25](https://github.com/Increase/increase-go/commit/b01ff255cf1a98aec9f4599636a80437c93b714d))


### Chores

* **internal:** improve union deserialization logic ([#215](https://github.com/Increase/increase-go/issues/215)) ([c345eb6](https://github.com/Increase/increase-go/commit/c345eb6e23a04e348bd6e4046fbefdcd5e9c099a))


### Documentation

* **contributing:** add a CONTRIBUTING.md ([#216](https://github.com/Increase/increase-go/issues/216)) ([253567b](https://github.com/Increase/increase-go/commit/253567beff555cbce99bc4feabffc7837a140660))

## 0.27.0 (2024-03-05)

Full Changelog: [v0.26.0...v0.27.0](https://github.com/Increase/increase-go/compare/v0.26.0...v0.27.0)

### ⚠ BREAKING CHANGES

* rename card_profile_id -> digital_card_profile_id ([#211](https://github.com/Increase/increase-go/issues/211))

### Features

* **api:** add actioner properties ([#209](https://github.com/Increase/increase-go/issues/209)) ([0a14d6d](https://github.com/Increase/increase-go/commit/0a14d6d8a2e648081ccdf90f03f962124f891ec6))


### Bug Fixes

* rename card_profile_id -&gt; digital_card_profile_id ([#211](https://github.com/Increase/increase-go/issues/211)) ([e0e36d1](https://github.com/Increase/increase-go/commit/e0e36d191fbcf3ad9c4954b91189287af69e1faf))


### Documentation

* update docs for digital wallet phone/email ([#212](https://github.com/Increase/increase-go/issues/212)) ([48a8e9d](https://github.com/Increase/increase-go/commit/48a8e9dea06c8171c4c5f16d26ba77b77e3f56a4))

## 0.26.0 (2024-03-01)

Full Changelog: [v0.25.0...v0.26.0](https://github.com/Increase/increase-go/compare/v0.25.0...v0.26.0)

### Features

* **api:** add post /entities/{entity_id}/industry_code endpoint ([#206](https://github.com/Increase/increase-go/issues/206)) ([edf8050](https://github.com/Increase/increase-go/commit/edf805086fcb3cb210f205addf3ec1215ab9461c))


### Documentation

* change industry code docstring ([#208](https://github.com/Increase/increase-go/issues/208)) ([1c50681](https://github.com/Increase/increase-go/commit/1c50681e55e7eb03dc17e91865071c77e35892d6))

## 0.25.0 (2024-02-29)

Full Changelog: [v0.24.0...v0.25.0](https://github.com/Increase/increase-go/compare/v0.24.0...v0.25.0)

### Features

* **api:** add `industry_code` property to methods ([#204](https://github.com/Increase/increase-go/issues/204)) ([d4bcfb1](https://github.com/Increase/increase-go/commit/d4bcfb1c605a051ce184a5d611ff2b8135cb587e))
* **api:** add unusual_activity_report_attachment enum member ([#203](https://github.com/Increase/increase-go/issues/203)) ([d2ccd9e](https://github.com/Increase/increase-go/commit/d2ccd9eca506ac4c694090c3c9ff35565a5e2733))
* **api:** inbound checks ([#200](https://github.com/Increase/increase-go/issues/200)) ([407455b](https://github.com/Increase/increase-go/commit/407455b47858f8455ae42d4d2c8b1156a0418d9c))
* **api:** physical card clone update ([#202](https://github.com/Increase/increase-go/issues/202)) ([ac5df4c](https://github.com/Increase/increase-go/commit/ac5df4cdb24995f9baa825eae9e0d8db08ffe81d))


### Chores

* **internal:** refactor release environment script ([#197](https://github.com/Increase/increase-go/issues/197)) ([ad111a6](https://github.com/Increase/increase-go/commit/ad111a67ac1d00496e80cffa6cd39c4e3e039b91))
* **internal:** update deps ([#201](https://github.com/Increase/increase-go/issues/201)) ([514020f](https://github.com/Increase/increase-go/commit/514020f4f54b367d3adecad5d5a3f59ee31f0846))


### Documentation

* **readme:** improve wording ([#205](https://github.com/Increase/increase-go/issues/205)) ([abc866b](https://github.com/Increase/increase-go/commit/abc866b451a3dfd4fce9b4a4776266ec3228a571))

## 0.24.0 (2024-02-15)

Full Changelog: [v0.23.1...v0.24.0](https://github.com/Increase/increase-go/compare/v0.23.1...v0.24.0)

### ⚠ BREAKING CHANGES

* **api:** split card profile resource into digital and physical card profile resources ([#196](https://github.com/Increase/increase-go/issues/196))

### Features

* **api:** split card profile resource into digital and physical card profile resources ([#196](https://github.com/Increase/increase-go/issues/196)) ([debcde2](https://github.com/Increase/increase-go/commit/debcde24411b678af2074826010f9eec2eb9ffc3))


### Chores

* **ci:** uses Stainless GitHub App for releases ([#195](https://github.com/Increase/increase-go/issues/195)) ([8a3dc3d](https://github.com/Increase/increase-go/commit/8a3dc3db7ed0ec3170972248dc25b2056ce0cd54))
* **internal:** adjust formatting ([#192](https://github.com/Increase/increase-go/issues/192)) ([15bdab6](https://github.com/Increase/increase-go/commit/15bdab6ab62d9f6a6a3d2683991f34ab516d0201))
* **internal:** bump timeout threshold in tests ([#194](https://github.com/Increase/increase-go/issues/194)) ([0283780](https://github.com/Increase/increase-go/commit/0283780d9005f4e36187ebf25c78f436eb60b72c))

## 0.23.1 (2024-02-02)

Full Changelog: [v0.23.0...v0.23.1](https://github.com/increase/increase-go/compare/v0.23.0...v0.23.1)

### Bug Fixes

* **docs:** correct pagination api.md link ([#191](https://github.com/increase/increase-go/issues/191)) ([67ba0cf](https://github.com/increase/increase-go/commit/67ba0cf1d71e065a54c32beba3745d897e169d78))


### Chores

* **interal:** make link to api.md relative ([#190](https://github.com/increase/increase-go/issues/190)) ([06c06c2](https://github.com/increase/increase-go/commit/06c06c2f58ac15d480ae99c8871381dd7b5b2e90))
* **internal:** parse date-time strings more leniently ([#187](https://github.com/increase/increase-go/issues/187)) ([cb474d3](https://github.com/increase/increase-go/commit/cb474d341430b52b43ca5eafba68ce6076dd3b34))
* **internal:** support pre-release versioning ([#189](https://github.com/increase/increase-go/issues/189)) ([012e7ec](https://github.com/increase/increase-go/commit/012e7eca837ded225aa466400482e00aab220c2e))

## 0.23.0 (2024-01-26)

Full Changelog: [v0.22.0...v0.23.0](https://github.com/increase/increase-go/compare/v0.22.0...v0.23.0)

### Features

* **api:** update descriptions ([#185](https://github.com/increase/increase-go/issues/185)) ([da37c8a](https://github.com/increase/increase-go/commit/da37c8a431b8c0f10e2b28ede26ec1246b339e23))

## 0.22.0 (2024-01-24)

Full Changelog: [v0.21.0...v0.22.0](https://github.com/increase/increase-go/compare/v0.21.0...v0.22.0)

### Features

* **api:** list Inbound Wire Transfers, change transfer simulation return types ([#183](https://github.com/increase/increase-go/issues/183)) ([165bb36](https://github.com/increase/increase-go/commit/165bb36e975043986fec63588fe7b796f176ab72))

## 0.21.0 (2024-01-22)

Full Changelog: [v0.20.0...v0.21.0](https://github.com/increase/increase-go/compare/v0.20.0...v0.21.0)

### Features

* **api:** simplify WireDecline and InboundWireTransfer ([#181](https://github.com/increase/increase-go/issues/181)) ([fb83ec9](https://github.com/increase/increase-go/commit/fb83ec90bef468edc969649028b0676255737f7a))

## 0.20.0 (2024-01-19)

Full Changelog: [v0.19.2...v0.20.0](https://github.com/increase/increase-go/compare/v0.19.2...v0.20.0)

### ⚠ BREAKING CHANGES

* fix oauth casing ([#178](https://github.com/increase/increase-go/issues/178))

### Features

* **api:** add oauth token and inbound wire transfer methods ([#180](https://github.com/increase/increase-go/issues/180)) ([cf89aec](https://github.com/increase/increase-go/commit/cf89aec6e0873885cee2411b9acd27d1d8457295))


### Refactors

* fix oauth casing ([#178](https://github.com/increase/increase-go/issues/178)) ([7ba4b27](https://github.com/increase/increase-go/commit/7ba4b2765179e1edf52a53752b0c67fe5a66e02a))

## 0.19.2 (2024-01-18)

Full Changelog: [v0.19.1...v0.19.2](https://github.com/increase/increase-go/compare/v0.19.1...v0.19.2)

### Bug Fixes

* **ci:** ignore stainless-app edits to release PR title ([#176](https://github.com/increase/increase-go/issues/176)) ([d03a9b9](https://github.com/increase/increase-go/commit/d03a9b990c7de97b8c2c93af9ffe8ebd0402ac5a))

## 0.19.1 (2024-01-17)

Full Changelog: [v0.19.0...v0.19.1](https://github.com/increase/increase-go/compare/v0.19.0...v0.19.1)

### Bug Fixes

* **test:** avoid test failures when SKIP_MOCK_TESTS is not set ([#175](https://github.com/increase/increase-go/issues/175)) ([2635e32](https://github.com/increase/increase-go/commit/2635e32ca2899d270f1647585b43e408a342e912))


### Chores

* **internal:** speculative retry-after-ms support ([#174](https://github.com/increase/increase-go/issues/174)) ([13db633](https://github.com/increase/increase-go/commit/13db633a99b4f7869c959deb02f5328cae084cb6))
* remove Alex Rattray from reviewers ([#173](https://github.com/increase/increase-go/issues/173)) ([f708587](https://github.com/increase/increase-go/commit/f708587242a992b8461a0f3756c1f55e3230bf7f))


### Documentation

* **api:** fix typo ([#171](https://github.com/increase/increase-go/issues/171)) ([c9628cf](https://github.com/increase/increase-go/commit/c9628cf92d32b84e599c73dcc6bf9de0ec72df16))

## 0.19.0 (2024-01-12)

Full Changelog: [v0.18.0...v0.19.0](https://github.com/increase/increase-go/compare/v0.18.0...v0.19.0)

### Features

* **api:** add merchant data to simulation api ([#170](https://github.com/increase/increase-go/issues/170)) ([32a6c2f](https://github.com/increase/increase-go/commit/32a6c2f034008d8a6902c5dbbbd59450c78164a4))


### Chores

* add .keep files for examples and custom code directories ([#166](https://github.com/increase/increase-go/issues/166)) ([d8a4de6](https://github.com/increase/increase-go/commit/d8a4de6e1150883ba5653ebd4e9eec446cffe802))
* **api:** update for other platforms ([#168](https://github.com/increase/increase-go/issues/168)) ([887edd7](https://github.com/increase/increase-go/commit/887edd70df65adae9837924744436d8f05928c33))


### Documentation

* **readme:** improve api reference ([#169](https://github.com/increase/increase-go/issues/169)) ([285cddf](https://github.com/increase/increase-go/commit/285cddf009a81975af78998cd5cf16ff8548a57e))

## 0.18.0 (2024-01-08)

Full Changelog: [v0.17.0...v0.18.0](https://github.com/increase/increase-go/compare/v0.17.0...v0.18.0)

### Features

* **api:** add `ach_debit_status` ([#165](https://github.com/increase/increase-go/issues/165)) ([954a0b1](https://github.com/increase/increase-go/commit/954a0b1a9fd6405e83a7182e8f137889985425f9))


### Chores

* **internal:** minor updates to pagination ([#163](https://github.com/increase/increase-go/issues/163)) ([fe89702](https://github.com/increase/increase-go/commit/fe8970229e509e76cf901c534ff71bef80915e8f))

## 0.17.0 (2024-01-02)

Full Changelog: [v0.16.0...v0.17.0](https://github.com/increase/increase-go/compare/v0.16.0...v0.17.0)

### Features

* **api:** add real-time payments request for payment endpoints ([#162](https://github.com/increase/increase-go/issues/162)) ([d5bc6af](https://github.com/increase/increase-go/commit/d5bc6affe973805882432c8f22b689ab099a68c2))


### Chores

* **internal:** bump license ([#161](https://github.com/increase/increase-go/issues/161)) ([6b77d51](https://github.com/increase/increase-go/commit/6b77d5162dd7bcaeffafe7b649ca3309908cfaa7))


### Documentation

* **options:** fix link to readme ([#159](https://github.com/increase/increase-go/issues/159)) ([a487dc4](https://github.com/increase/increase-go/commit/a487dc403bc6707218696c8d5d35d91c2395c5e8))

## 0.16.0 (2023-12-18)

Full Changelog: [v0.15.0...v0.16.0](https://github.com/increase/increase-go/compare/v0.15.0...v0.16.0)

### Features

* **api:** add deposit_submission property to check deposit ([#158](https://github.com/increase/increase-go/issues/158)) ([4f32cd0](https://github.com/increase/increase-go/commit/4f32cd0dcd8fc04775edad840237434e6a0aad03))


### Chores

* **ci:** run release workflow once per day ([#156](https://github.com/increase/increase-go/issues/156)) ([4b116be](https://github.com/increase/increase-go/commit/4b116bedc96bdb3cf0868962dd696de16ea5fc0b))

## 0.15.0 (2023-12-14)

Full Changelog: [v0.14.0...v0.15.0](https://github.com/increase/increase-go/compare/v0.14.0...v0.15.0)

### Features

* **api:** add `suspected_fraud` rejection reason ([#155](https://github.com/increase/increase-go/issues/155)) ([eded3c3](https://github.com/increase/increase-go/commit/eded3c3570d95996d7182c956de38f4239cd2002))
* **internal:** fallback to json serialization if no serialization methods are defined ([#153](https://github.com/increase/increase-go/issues/153)) ([3dfa03a](https://github.com/increase/increase-go/commit/3dfa03ababd7f102d5a1eeac0b0d4cecb74d7d60))

## 0.14.0 (2023-12-08)

Full Changelog: [v0.13.2...v0.14.0](https://github.com/increase/increase-go/compare/v0.13.2...v0.14.0)

### Features

* **api:** updates ([#152](https://github.com/increase/increase-go/issues/152)) ([26817d3](https://github.com/increase/increase-go/commit/26817d3b5dfcda4df46472ccc5db870b4272ed8a))

## 0.13.2 (2023-12-04)

Full Changelog: [v0.13.1...v0.13.2](https://github.com/increase/increase-go/compare/v0.13.1...v0.13.2)

## 0.13.1 (2023-11-17)

Full Changelog: [v0.13.0...v0.13.1](https://github.com/increase/increase-go/compare/v0.13.0...v0.13.1)

### Bug Fixes

* stop sending default idempotency headers with GET requests ([#147](https://github.com/increase/increase-go/issues/147)) ([3d3dc4b](https://github.com/increase/increase-go/commit/3d3dc4b9dfc9fa9f60924106ec196d1ee6c6e3c7))


### Refactors

* do not include `JSON` fields when serialising back to json ([#145](https://github.com/increase/increase-go/issues/145)) ([70fd3ef](https://github.com/increase/increase-go/commit/70fd3ef6b320074dc9b4117efa33a205d8328aa6))

## 0.13.0 (2023-11-08)

Full Changelog: [v0.12.0...v0.13.0](https://github.com/Increase/increase-go/compare/v0.12.0...v0.13.0)

### Features

* **api:** add failed to export status enum ([#139](https://github.com/Increase/increase-go/issues/139)) ([b7c4de4](https://github.com/Increase/increase-go/commit/b7c4de40eb97590f54eb1eba7fb21a2ace11dc27))
* **api:** add fuel confirmation functionality ([#132](https://github.com/Increase/increase-go/issues/132)) ([0dd1aa5](https://github.com/Increase/increase-go/commit/0dd1aa595e40f3202b9b4c2223a4fd324ec2b050))
* **api:** add network identifiers and effective date ([#136](https://github.com/Increase/increase-go/issues/136)) ([da2a8aa](https://github.com/Increase/increase-go/commit/da2a8aa620a6ce780fcac55fbd55c26bf13661aa))
* **api:** restructure balance lookups ([#142](https://github.com/Increase/increase-go/issues/142)) ([beb9f98](https://github.com/Increase/increase-go/commit/beb9f988ab398ecde14a413fb3ca654e55e32964))
* **api:** updates ([#135](https://github.com/Increase/increase-go/issues/135)) ([0df7444](https://github.com/Increase/increase-go/commit/0df7444ffc43b9f34622c1ffa52f443e234ef5f8))
* **client:** adjust retry behavior ([#133](https://github.com/Increase/increase-go/issues/133)) ([cf684e9](https://github.com/Increase/increase-go/commit/cf684e9de7fc7907c7eeaa39dc8a0e827cfd6be5))
* **client:** allow binary returns ([#140](https://github.com/Increase/increase-go/issues/140)) ([f1935de](https://github.com/Increase/increase-go/commit/f1935de01ec39cfb9589400d7fd9106db38193f5))
* **github:** include a devcontainer setup ([#138](https://github.com/Increase/increase-go/issues/138)) ([eb7829a](https://github.com/Increase/increase-go/commit/eb7829a0fdd2a1acb41cd30add79706954f69634))


### Bug Fixes

* make options.WithHeader utils case-insensitive ([#143](https://github.com/Increase/increase-go/issues/143)) ([2fcf014](https://github.com/Increase/increase-go/commit/2fcf01492775bc91196059b3b299aed15e00fbc0))


### Chores

* **internal:** update gitignore ([#137](https://github.com/Increase/increase-go/issues/137)) ([29068e7](https://github.com/Increase/increase-go/commit/29068e72abd23c4455fb48f5c91ccb7b7696c0be))


### Documentation

* **readme:** improve example snippets ([#141](https://github.com/Increase/increase-go/issues/141)) ([b34ca1b](https://github.com/Increase/increase-go/commit/b34ca1b0c72aa1d534e7cf5d45bcd2817a1beeb3))

## 0.12.0 (2023-10-20)

Full Changelog: [v0.11.0...v0.12.0](https://github.com/increase/increase-go/compare/v0.11.0...v0.12.0)

### Features

* **api:** add returned_per_odfi_request enum ([#130](https://github.com/increase/increase-go/issues/130)) ([ca5c55a](https://github.com/increase/increase-go/commit/ca5c55a7fa03cee6f94c9182ec97c0bf3e160a90))

## 0.11.0 (2023-10-20)

Full Changelog: [v0.10.0...v0.11.0](https://github.com/increase/increase-go/compare/v0.10.0...v0.11.0)

### Features

* **api:** updates ([#128](https://github.com/increase/increase-go/issues/128)) ([43d990f](https://github.com/increase/increase-go/commit/43d990fa48dfb8fc32d578872b22e8b32e2db70a))

## 0.10.0 (2023-10-20)

Full Changelog: [v0.9.0...v0.10.0](https://github.com/increase/increase-go/compare/v0.9.0...v0.10.0)

### Features

* **api:** updates ([#126](https://github.com/increase/increase-go/issues/126)) ([df169a7](https://github.com/increase/increase-go/commit/df169a7d42c3ea0e12f0bd3feb575f881c6676a3))

## 0.9.0 (2023-10-19)

Full Changelog: [v0.8.6...v0.9.0](https://github.com/increase/increase-go/compare/v0.8.6...v0.9.0)

### Features

* **api:** add addenda details for ACH transfers ([#124](https://github.com/increase/increase-go/issues/124)) ([62d8fb5](https://github.com/increase/increase-go/commit/62d8fb5b44c8cd43972044782621bc3337e77b8b))
* **api:** updates ([#125](https://github.com/increase/increase-go/issues/125)) ([06f453f](https://github.com/increase/increase-go/commit/06f453fee8f8979903fb7400486dc734ae93baf9))


### Chores

* **internal:** rearrange client arguments ([#121](https://github.com/increase/increase-go/issues/121)) ([6a1fad3](https://github.com/increase/increase-go/commit/6a1fad31bee4f628c95aa1bdea87b7db7f304a50))
* update README ([#119](https://github.com/increase/increase-go/issues/119)) ([3c7b2c8](https://github.com/increase/increase-go/commit/3c7b2c8b89aa36d8938338baf653f6d934130492))


### Documentation

* organisation -&gt; organization (UK to US English) ([#123](https://github.com/increase/increase-go/issues/123)) ([957d4a9](https://github.com/increase/increase-go/commit/957d4a96f3f8ac7344548656f1f59225647e97e0))

## 0.8.6 (2023-10-11)

Full Changelog: [v0.8.5...v0.8.6](https://github.com/increase/increase-go/compare/v0.8.5...v0.8.6)

## 0.8.5 (2023-10-05)

Full Changelog: [v0.8.4...v0.8.5](https://github.com/increase/increase-go/compare/v0.8.4...v0.8.5)

### Features

* **api:** add direction property to CardAuthorization ([#115](https://github.com/increase/increase-go/issues/115)) ([f4e2edf](https://github.com/increase/increase-go/commit/f4e2edf1b196827c110e9c25107e6e89effd7b35))

## 0.8.4 (2023-10-03)

Full Changelog: [v0.8.3...v0.8.4](https://github.com/increase/increase-go/compare/v0.8.3...v0.8.4)

### Features

* **api:** add card payments endpoints ([#109](https://github.com/increase/increase-go/issues/109)) ([baf55ee](https://github.com/increase/increase-go/commit/baf55eeb4994eff992e8f775bdd790fc47a59a54))
* **api:** add physical cards endpoints ([#107](https://github.com/increase/increase-go/issues/107)) ([aa5b28e](https://github.com/increase/increase-go/commit/aa5b28eceabdbb92472edee065c929c8ae5d17c5))
* **api:** expand event categories and Entity status options ([#113](https://github.com/increase/increase-go/issues/113)) ([90b0116](https://github.com/increase/increase-go/commit/90b01162ac64dd5896ef7f27d9bc3b9094f793a2))


### Bug Fixes

* prevent index out of range bug during auto-pagination ([#112](https://github.com/increase/increase-go/issues/112)) ([d86b9cc](https://github.com/increase/increase-go/commit/d86b9ccadc4ff6dde310bfd29b26c225fb532f01))


### Chores

* **tests:** update test examples ([#114](https://github.com/increase/increase-go/issues/114)) ([2d1ba8a](https://github.com/increase/increase-go/commit/2d1ba8afe168713bf4580b03d4e66e98c3822025))

## 0.8.3 (2023-09-25)

Full Changelog: [v0.8.2...v0.8.3](https://github.com/Increase/increase-go/compare/v0.8.2...v0.8.3)

### Features

* **api:** add inbound checks, originating routing number and new event types ([#105](https://github.com/Increase/increase-go/issues/105)) ([7a9328b](https://github.com/Increase/increase-go/commit/7a9328bf3c734fdcdc0a9910d0769733aa1fa274))
* improve retry behavior on context deadline ([#106](https://github.com/Increase/increase-go/issues/106)) ([a2f1c71](https://github.com/Increase/increase-go/commit/a2f1c7186437378f71523c589b6b804bc13a0ce0))


### Bug Fixes

* **docs:** change `Page` package to shared ([#103](https://github.com/Increase/increase-go/issues/103)) ([69077fe](https://github.com/Increase/increase-go/commit/69077fe46795688978b98de73cd52b238f009c95))

## 0.8.2 (2023-09-20)

Full Changelog: [v0.8.1...v0.8.2](https://github.com/Increase/increase-go/compare/v0.8.1...v0.8.2)

### Features

* **api:** add entity_id to Card and make relationship nullable ([#100](https://github.com/Increase/increase-go/issues/100)) ([8a7b176](https://github.com/Increase/increase-go/commit/8a7b176a4d62a822affe1e76cdc2158fbf058601))
* **api:** export account statements in OFX format ([#102](https://github.com/Increase/increase-go/issues/102)) ([2889a31](https://github.com/Increase/increase-go/commit/2889a31818e75fb86daf16dec328f9e3a10a24b3))

## 0.8.1 (2023-09-15)

Full Changelog: [v0.8.0...v0.8.1](https://github.com/increase/increase-go/compare/v0.8.0...v0.8.1)

### Features

* **api:** add card payment ID reference to transaction models ([#94](https://github.com/increase/increase-go/issues/94)) ([c57cd95](https://github.com/increase/increase-go/commit/c57cd951b4b2d8cc91be971a6bfb02674df452b0))
* retry on 408 Request Timeout ([#96](https://github.com/increase/increase-go/issues/96)) ([b70d67d](https://github.com/increase/increase-go/commit/b70d67db84905247ea1b5996c642410b7e9df070))


### Bug Fixes

* **core:** improve retry behavior and related docs ([#97](https://github.com/increase/increase-go/issues/97)) ([739a4b1](https://github.com/increase/increase-go/commit/739a4b1f50e9b20974c54311f7aa0663b143ee63))

## 0.8.0 (2023-09-12)

Full Changelog: [v0.7.3...v0.8.0](https://github.com/increase/increase-go/compare/v0.7.3...v0.8.0)

### ⚠ BREAKING CHANGES

* **api:** remove Limits API, add ACH controls to Account Numbers   ([#92](https://github.com/increase/increase-go/issues/92))

### Features

* **api:** remove Limits API, add ACH controls to Account Numbers   ([#92](https://github.com/increase/increase-go/issues/92)) ([6e7817b](https://github.com/increase/increase-go/commit/6e7817bcde1dca89a0bb4f579a9df61a71f67b80))


### Bug Fixes

* **core:** add null check to prevent segfault when canceling context ([#90](https://github.com/increase/increase-go/issues/90)) ([15bca67](https://github.com/increase/increase-go/commit/15bca675a4b7cbf3fdeb48f78ffbf6a0d551b011))


### Chores

* **internal:** improve reliability of cancel delay test ([#93](https://github.com/increase/increase-go/issues/93)) ([c1efae5](https://github.com/increase/increase-go/commit/c1efae5a55f9fd1bbb2ca6b05602bbfe72709d37))

## 0.7.3 (2023-09-07)

Full Changelog: [v0.7.2...v0.7.3](https://github.com/increase/increase-go/compare/v0.7.2...v0.7.3)

### Features

* **api:** add bank_of_first_deposit_routing_number and transfer_id properties ([#82](https://github.com/increase/increase-go/issues/82)) ([b9407a8](https://github.com/increase/increase-go/commit/b9407a8ea94153ef8f61eea71f1cf46a550dc8db))
* **api:** add Update Address  and Create Notification Change endpoints ([#86](https://github.com/increase/increase-go/issues/86)) ([f3de558](https://github.com/increase/increase-go/commit/f3de558bec0a1bcb0ce9272c81012de12e83860a))
* fixes tests where an array has to have unique enum values ([#87](https://github.com/increase/increase-go/issues/87)) ([ea994ce](https://github.com/increase/increase-go/commit/ea994ceee40286c47dc5e90b8042f9942ad2acb0))


### Chores

* **internal:** implement inline json unmarshalling ([#85](https://github.com/increase/increase-go/issues/85)) ([7862aa5](https://github.com/increase/increase-go/commit/7862aa57236bc7b996d582d22c735cdeb62cb38a))


### Documentation

* **api:** add docstrings and refine enum types ([#89](https://github.com/increase/increase-go/issues/89)) ([d38244b](https://github.com/increase/increase-go/commit/d38244bfe6bb1b43da9c2cc69430ad206183d5f2))
* **readme:** add link to api.md ([#88](https://github.com/increase/increase-go/issues/88)) ([98dcd31](https://github.com/increase/increase-go/commit/98dcd313800a9b122a144d3cb13b876e8735f5df))

## 0.7.2 (2023-08-29)

Full Changelog: [v0.7.1...v0.7.2](https://github.com/Increase/increase-go/compare/v0.7.1...v0.7.2)

### Features

* **api:** remove unused `/inbound_ach_transfer_returns` endpoints ([#81](https://github.com/Increase/increase-go/issues/81)) ([c545a87](https://github.com/Increase/increase-go/commit/c545a8743a69c11450980cd41acb3a36ba85451d))


### Bug Fixes

* **api:** move ACH Return endpoint and add digital wallet properties ([#80](https://github.com/Increase/increase-go/issues/80)) ([a119bd0](https://github.com/Increase/increase-go/commit/a119bd0631538fadefbeac12bf2ef0656d562860))


### Chores

* **ci:** setup workflows to create releases and release PRs ([#77](https://github.com/Increase/increase-go/issues/77)) ([3fd9088](https://github.com/Increase/increase-go/commit/3fd9088965f3ebdfaf7d69021ed796ddce5315f6))

## [0.7.1](https://github.com/Increase/increase-go/compare/v0.7.0...v0.7.1) (2023-08-27)


### Features

* **api:** move inbound ACH transfer returns (⚠️ breaking); add ACH transfer declines ([#75](https://github.com/Increase/increase-go/issues/75)) ([60a3e48](https://github.com/Increase/increase-go/commit/60a3e4810772bba6ee3b9c02aeed5081cd98c13a))
* **api:** updates ([#74](https://github.com/Increase/increase-go/issues/74)) ([7fc28e0](https://github.com/Increase/increase-go/commit/7fc28e044950394cadd73468010a3642b5b4f042))

## [0.7.0](https://github.com/Increase/increase-go/compare/v0.6.3...v0.7.0) (2023-08-17)


### ⚠ BREAKING CHANGES

* **api:** change `physical_cards.status` value, remove `event_subscription` field, add fields ([#70](https://github.com/Increase/increase-go/issues/70))

### Features

* **api:** change `physical_cards.status` value, remove `event_subscription` field, add fields ([#70](https://github.com/Increase/increase-go/issues/70)) ([bda93c7](https://github.com/Increase/increase-go/commit/bda93c7cbea14c011dd699a863c186f00af678a7))


### Chores

* assign default reviewers to release PRs ([#67](https://github.com/Increase/increase-go/issues/67)) ([e832e27](https://github.com/Increase/increase-go/commit/e832e272d6a95c137caf084c5ab8dd32ace49153))

## [0.6.3](https://github.com/Increase/increase-go/compare/v0.6.2...v0.6.3) (2023-08-11)


### Features

* **api:** updates ([#64](https://github.com/Increase/increase-go/issues/64)) ([8152bcf](https://github.com/Increase/increase-go/commit/8152bcfcd24a924ae4306e23ee08cee91488340d))


### Bug Fixes

* **client:** correctly set multipart form data boundary ([#62](https://github.com/Increase/increase-go/issues/62)) ([0eb012e](https://github.com/Increase/increase-go/commit/0eb012e7b813854c249f0a8db7f0dec34afc5f83))


### Documentation

* **api:** change description of various fields ([#65](https://github.com/Increase/increase-go/issues/65)) ([92a0c51](https://github.com/Increase/increase-go/commit/92a0c51da05ac91cc02ba767abc903e40a8ec259))

## [0.6.2](https://github.com/Increase/increase-go/compare/v0.6.1...v0.6.2) (2023-08-08)


### Features

* **api:** updates ([#59](https://github.com/Increase/increase-go/issues/59)) ([3ebf588](https://github.com/Increase/increase-go/commit/3ebf588b69d7cb1a2b304cf6008970a86bf4b7c7))


### Chores

* **internal:** minor reformatting of code ([#56](https://github.com/Increase/increase-go/issues/56)) ([b0cf72f](https://github.com/Increase/increase-go/commit/b0cf72f36e0ae74cace8b05a2b7fbedc09f9df8d))


### Documentation

* **readme:** remove beta status + document versioning policy ([#61](https://github.com/Increase/increase-go/issues/61)) ([4c384c7](https://github.com/Increase/increase-go/commit/4c384c7ebf099370ab11b66a2eb388ab6b02bd8e))

## [0.6.1](https://github.com/Increase/increase-go/compare/v0.6.0...v0.6.1) (2023-07-27)


### Features

* add `Bool` param field helper ([#53](https://github.com/Increase/increase-go/issues/53)) ([5dadbe7](https://github.com/Increase/increase-go/commit/5dadbe7d3a049c05bf971f674504a73357fd74ce))

## [0.6.0](https://github.com/Increase/increase-go/compare/v0.5.1...v0.6.0) (2023-07-22)


### ⚠ BREAKING CHANGES

* **api:** reorganize `check_transfer` and `network fields; add `request_details`; add `unknown` ([#48](https://github.com/Increase/increase-go/issues/48))

### Features

* **api:** add fee_period_start and return_of_erroneous_or_reversing_debit ([#51](https://github.com/Increase/increase-go/issues/51)) ([7f09d65](https://github.com/Increase/increase-go/commit/7f09d6599b40a47e555919713713ecb1379a2268))
* **api:** reorganize `check_transfer` and `network fields; add `request_details`; add `unknown` ([#48](https://github.com/Increase/increase-go/issues/48)) ([e35fba9](https://github.com/Increase/increase-go/commit/e35fba9db8310b4eeea1ee9e137d310a29524f2a))


### Documentation

* **api:** update `model_id` documentation ([#50](https://github.com/Increase/increase-go/issues/50)) ([00ef46d](https://github.com/Increase/increase-go/commit/00ef46d8adb055e823926266c6f4a5986273cc8f))

## [0.5.1](https://github.com/Increase/increase-go/compare/v0.5.0...v0.5.1) (2023-07-17)


### Features

* **api:** add physical_card_id ([#44](https://github.com/Increase/increase-go/issues/44)) ([2b953a5](https://github.com/Increase/increase-go/commit/2b953a5ac6c29129479d247a0ea60fd97f66ab3f))


### Chores

* **internal:** add `codegen.log` to `.gitignore` ([#46](https://github.com/Increase/increase-go/issues/46)) ([720e4ab](https://github.com/Increase/increase-go/commit/720e4ab7c9fab1206120d1e9466f3c8cb9b3d923))

## [0.5.0](https://github.com/Increase/increase-go/compare/v0.4.0...v0.5.0) (2023-07-12)


### ⚠ BREAKING CHANGES

* **api:** add unique_identifier, driver's license backs, inbound funds holds, and more ([#41](https://github.com/Increase/increase-go/issues/41))

### Features

* **api:** add unique_identifier, driver's license backs, inbound funds holds, and more ([#41](https://github.com/Increase/increase-go/issues/41)) ([9809814](https://github.com/Increase/increase-go/commit/980981411a18f16a4de372deabcd8d362b8575d5))

## [0.4.0](https://github.com/Increase/increase-go/compare/v0.3.0...v0.4.0) (2023-07-07)


### ⚠ BREAKING CHANGES

* **api:** add card profiles simulation method ([#39](https://github.com/Increase/increase-go/issues/39))

### Features

* **api:** add card profiles simulation method ([#39](https://github.com/Increase/increase-go/issues/39)) ([0774f0a](https://github.com/Increase/increase-go/commit/0774f0aaae88b27c2d1abb6aa641c4b507651896))


### Documentation

* **types:** add documentation for enum members ([#37](https://github.com/Increase/increase-go/issues/37)) ([7d35ce9](https://github.com/Increase/increase-go/commit/7d35ce9f05e881d8ad701447c54d6d565ff3abfc))

## [0.3.0](https://github.com/Increase/increase-go/compare/v0.2.0...v0.3.0) (2023-06-29)


### ⚠ BREAKING CHANGES

* **types:** singularize array item types ([#36](https://github.com/Increase/increase-go/issues/36))
* rename a response type and remove unnecessary types from paginated endpoints ([#33](https://github.com/Increase/increase-go/issues/33))
* **api:** remove many enum members from document category ([#32](https://github.com/Increase/increase-go/issues/32))

### Features

* **api/types:** mark more check transfer intention properties as nullable ([#31](https://github.com/Increase/increase-go/issues/31)) ([429efca](https://github.com/Increase/increase-go/commit/429efca9994054ed7b33684ebcc9e88d92637895))
* generate `api.md` file ([#29](https://github.com/Increase/increase-go/issues/29)) ([8242f41](https://github.com/Increase/increase-go/commit/8242f41fd1feec60ca3d98ec3bdeac852628b288))
* respect `x-should-retry` header ([#22](https://github.com/Increase/increase-go/issues/22)) ([314b640](https://github.com/Increase/increase-go/commit/314b640c091c895fa3749622c1671df543fab087))


### Chores

* **tests:** minor reformatting of docs and tests ([#25](https://github.com/Increase/increase-go/issues/25)) ([1fb923f](https://github.com/Increase/increase-go/commit/1fb923f31efb17255de1ec19087066fc8b85029b))


### Documentation

* add comments to alias types ([#30](https://github.com/Increase/increase-go/issues/30)) ([7f54628](https://github.com/Increase/increase-go/commit/7f546284c9791d921f988ac3d39df056da8d9c8f))
* add trailing newlines ([#34](https://github.com/Increase/increase-go/issues/34)) ([537501c](https://github.com/Increase/increase-go/commit/537501c031d9fdb0e08f3071a19dff1e13309d78))


### Styles

* minor reordering of types and properties ([#35](https://github.com/Increase/increase-go/issues/35)) ([556169b](https://github.com/Increase/increase-go/commit/556169b6bf09eca01edbdc120d002739c29a798e))


### Refactors

* **api:** remove `other` from reason enum ([#27](https://github.com/Increase/increase-go/issues/27)) ([8136d90](https://github.com/Increase/increase-go/commit/8136d903fffcd1aedcd234f752a82bdee84055ac))
* **api:** remove many enum members from document category ([#32](https://github.com/Increase/increase-go/issues/32)) ([2eb741a](https://github.com/Increase/increase-go/commit/2eb741a136cb2418bf3a34b892ee35bdd1b2f32b))
* **api:** remove unused properties and enum members ([#24](https://github.com/Increase/increase-go/issues/24)) ([6a450bb](https://github.com/Increase/increase-go/commit/6a450bb2600eeaad489b5c363d1adc0d3337e0af))
* improve `time.Time` encoding and decoding ([#20](https://github.com/Increase/increase-go/issues/20)) ([4bde6e9](https://github.com/Increase/increase-go/commit/4bde6e9b7d2b6a6bad2fb4ae074bd8a1b47167b5))
* rename a response type and remove unnecessary types from paginated endpoints ([#33](https://github.com/Increase/increase-go/issues/33)) ([0b8dacd](https://github.com/Increase/increase-go/commit/0b8dacd07acfc3b42c2b1022a5035f1103eace4a))
* **types:** singularize array item types ([#36](https://github.com/Increase/increase-go/issues/36)) ([67ed591](https://github.com/Increase/increase-go/commit/67ed5914941faa47bd4794ea6c0aef1308a1b0d0))

## [0.2.0](https://github.com/Increase/increase-go/compare/v0.1.1...v0.2.0) (2023-06-13)


### ⚠ BREAKING CHANGES

* **api:** rename return reason enum member ([#16](https://github.com/Increase/increase-go/issues/16))

### Features

* implement middleware ([#17](https://github.com/Increase/increase-go/issues/17)) ([a2e8992](https://github.com/Increase/increase-go/commit/a2e8992e0f603921eace035b57673e7f2fe657f8))
* make tests give better error message on missing prism server ([#14](https://github.com/Increase/increase-go/issues/14)) ([092b4e3](https://github.com/Increase/increase-go/commit/092b4e3d5e779662a438ab12fd54caa7de2be553))


### Chores

* add docstrings to field helpers, error struct, enums, and more ([#12](https://github.com/Increase/increase-go/issues/12)) ([89fc73e](https://github.com/Increase/increase-go/commit/89fc73ea340d8e49625cac6f180f57e50e60c242))
* **internal:** add more comments ([#10](https://github.com/Increase/increase-go/issues/10)) ([33a658b](https://github.com/Increase/increase-go/commit/33a658b232891c9ef4af1416a1768942d943ebe9))
* **internal:** remove unused check-test-server script ([#13](https://github.com/Increase/increase-go/issues/13)) ([f67f682](https://github.com/Increase/increase-go/commit/f67f682c55ad09bbd2a7322d9f7b6a96a10cfae6))


### Refactors

* **api:** rename return reason enum member ([#16](https://github.com/Increase/increase-go/issues/16)) ([5fe33f5](https://github.com/Increase/increase-go/commit/5fe33f518020a3a6d22327d443b0507a7ffea412))


### Documentation

* point to github repo instead of email contact ([#18](https://github.com/Increase/increase-go/issues/18)) ([a945477](https://github.com/Increase/increase-go/commit/a9454774268c89988317484fb9b0207675de34e1))

## [0.1.1](https://github.com/Increase/increase-go/compare/v0.1.0...v0.1.1) (2023-06-09)


### Features

* **api:** add new endpoints + properties + enums ([#7](https://github.com/Increase/increase-go/issues/7)) ([52481ec](https://github.com/Increase/increase-go/commit/52481ecdf53674a35b7fbbfaab432e7b74be1468))

## [0.1.0](https://github.com/Increase/increase-go/compare/v0.0.1...v0.1.0) (2023-05-31)


### ⚠ BREAKING CHANGES

* rename `.JSON.Extras` -> `.JSON.ExtraFields`
* rename `field.Field` -> `param.Field`
* make JSON structs private, rename Metadata->Field, improve docs
* correct various identifier names
* **api:** replace notification_of_change with a list, and add merchant_acceptor_id
* rename param types

### Features

* add custom `String()` function on model structs ([5b0d022](https://github.com/Increase/increase-go/commit/5b0d022e7d651f3b28c1c3d37bef152c96213997))
* add custom serialization code to models ([76c9fb8](https://github.com/Increase/increase-go/commit/76c9fb844bce4d18c90d3885723af19df206ba16))
* add deprecated methods ([3747a6a](https://github.com/Increase/increase-go/commit/3747a6a2533ec0394fc653f7c0c7f2407d9efe2f))
* add support for encoding/decoding embedded structs ([fb63bc1](https://github.com/Increase/increase-go/commit/fb63bc1206a778f854ca29d9058dfc99e01ab1ce))
* adjust README examples ([e539be4](https://github.com/Increase/increase-go/commit/e539be4ef818aee89f9a3760ad7e6b2af8af1bd1))
* **api:** add `at_time` property for balance lookups ([301c656](https://github.com/Increase/increase-go/commit/301c6567c72b30ef96da5aca7f37fb93d3588db6))
* **api:** add `collection_receivable` to transaction source category enum ([a8ac578](https://github.com/Increase/increase-go/commit/a8ac5788336aeeeaa7ae1c9214df8b09ff1d5dea))
* **api:** add `expires_at` property ([b4ccd08](https://github.com/Increase/increase-go/commit/b4ccd08a3c1c3a42fa57d73bde08d3917925ef0e))
* **api:** add `Simulations.CheckTransfers.Return()` method ([a5cfea0](https://github.com/Increase/increase-go/commit/a5cfea0d536d5961605238eb4871fd5ef76cedd7))
* **api:** add bookkeeping accounts, entries, and entry sets, and several other changes ([0730e45](https://github.com/Increase/increase-go/commit/0730e453a6e0c793ea5e95a69e5dd64c04b9e460))
* **api:** add new endpoints ([bd59552](https://github.com/Increase/increase-go/commit/bd59552251c7edb9ee834ff9e5af86ea5d3af59a))
* **api:** add new endpoints, several params, fields, enum members, and documentation updates ([b2f9d37](https://github.com/Increase/increase-go/commit/b2f9d37b3a562b148d6f5420d2257616974cf89c))
* **api:** add new enum members ([80dcdaf](https://github.com/Increase/increase-go/commit/80dcdaf5c30530229847134d910b4bfac2b14120))
* **api:** add new fields ([ed68c93](https://github.com/Increase/increase-go/commit/ed68c93cbf4b6a33d2eb3605bc9c0a5a4860183a))
* **api:** add new methods ([28484ce](https://github.com/Increase/increase-go/commit/28484ce6bdf4edc3ce6ac7b83c737c823bd5abd8))
* **api:** add optional `pending_transaction_id` field to pending transaction ([636be67](https://github.com/Increase/increase-go/commit/636be6753839bf0a9d6a287b661af4988ed5d56c))
* **api:** add wire decline object ([3ca48ea](https://github.com/Increase/increase-go/commit/3ca48ea9ab3e2f219a69a0a98d5bf0642b79d83b))
* **api:** enum updates ([8a95d0f](https://github.com/Increase/increase-go/commit/8a95d0ff3f3498429cd708ac93237b92a631f100))
* **api:** make routeType an enum & add ACHTransfer.effectiveDate ([2cbf686](https://github.com/Increase/increase-go/commit/2cbf686956ee22804d47f8de0979fb5d9b4f6683))
* **api:** make routeType an enum & add ACHTransfer.effectiveDate ([2cbf686](https://github.com/Increase/increase-go/commit/2cbf686956ee22804d47f8de0979fb5d9b4f6683))
* **api:** remove CardSettlementTransactionID ([46f3d6a](https://github.com/Increase/increase-go/commit/46f3d6ad39c7c7b07c2a2df00e2f7dc8cab67037))
* **api:** replace notification_of_change with a list, and add merchant_acceptor_id ([8391622](https://github.com/Increase/increase-go/commit/8391622117ac5ecde7cdc235a660d616ea7269e6))
* **api:** updates ([f9c3b7f](https://github.com/Increase/increase-go/commit/f9c3b7f85159a9c2892dbb524542e6c759f54741))
* **api:** updates ([c9a5d96](https://github.com/Increase/increase-go/commit/c9a5d96e7ff4afbbab701d8e0a0a8ba391d2d28b))
* **api:** updates ([95f71fc](https://github.com/Increase/increase-go/commit/95f71fc7fb4ecc71814bb1dd142b9eb491d86cd7))
* **api:** updates ([94fd784](https://github.com/Increase/increase-go/commit/94fd784c80f34df6fbf22653a734eada8caf8c3c))
* better errors ([e12a285](https://github.com/Increase/increase-go/commit/e12a28545d4a17cf84b5812209cfab2f74132072))
* **docs:** include version references in the README ([fc28185](https://github.com/Increase/increase-go/commit/fc28185fbc2c00e516204b073217dc03f6b8838e))
* **docs:** updates ([708e9b0](https://github.com/Increase/increase-go/commit/708e9b092cb5bdb770e73480351dab1d60972817))
* elaborate readme and improve doc comments ([d8b752d](https://github.com/Increase/increase-go/commit/d8b752d74db778ab0b839ea5faaa62fcd5d3fed6))
* elaborate readme and improve doc comments ([d8b752d](https://github.com/Increase/increase-go/commit/d8b752d74db778ab0b839ea5faaa62fcd5d3fed6))
* generate Go tests ([19d4874](https://github.com/Increase/increase-go/commit/19d4874a6a5dcf95330834f472821db359cf7a2a))
* implement auto-retry ([00fcc4f](https://github.com/Increase/increase-go/commit/00fcc4f39f42af7074d84c04813dff0275657347))
* implement bikesheds ([5f57aac](https://github.com/Increase/increase-go/commit/5f57aac2216086f9ad3ef6512144595e3495f13a))
* implement file uploads ([0246753](https://github.com/Increase/increase-go/commit/0246753c8e4584f55bbd0245a4aef982ca25fac8))
* implement improved auto-pagination ([2b8fd25](https://github.com/Increase/increase-go/commit/2b8fd256f22edde7c70b417cdf6adbb0d4fe9622))
* implement unions ([ee3bd0f](https://github.com/Increase/increase-go/commit/ee3bd0ffae121ab811d63ff3bbc6802d6cfb057c))
* implement unions ([ee3bd0f](https://github.com/Increase/increase-go/commit/ee3bd0ffae121ab811d63ff3bbc6802d6cfb057c))
* improve docs and add new property ([3975526](https://github.com/Increase/increase-go/commit/39755268a0c07b5a37cb3aff14d8b7a95ce46370))
* lift fields helpers to main class ([52cd758](https://github.com/Increase/increase-go/commit/52cd758ceb8feb4f546ab9861501b9f10c10b345))
* move all resources to the services package ([88c63b8](https://github.com/Increase/increase-go/commit/88c63b8ecbba2770a00a3fdb4f832980a2555714))
* move to value receivers for getters ([ee23b76](https://github.com/Increase/increase-go/commit/ee23b76979ea4d1e8e01ffd50f7c54d49a354a16))
* prune unused models in requests ([1347220](https://github.com/Increase/increase-go/commit/1347220a93a7b4bb4b716d655824a4d393e034cb))
* read options from environment and use auth headers ([8969b12](https://github.com/Increase/increase-go/commit/8969b1204129a3980fb6640a6c51160d124802be))
* refactor code to use functional options in a separate options package ([47e68b6](https://github.com/Increase/increase-go/commit/47e68b61f9b951ab2ad2d24b851b26db7bff7e61))
* send Idempotency-Key by default for POST requests ([1aa9803](https://github.com/Increase/increase-go/commit/1aa9803777904250a8639cee935861c4a78e821f))
* send package version in X-Stainless-Package-Version ([385b3b1](https://github.com/Increase/increase-go/commit/385b3b10421fe582c21cdca7641fbf5a71dd18a9))
* separate types into params/responses packages ([dedbd12](https://github.com/Increase/increase-go/commit/dedbd124cb65f38f00c8ce358347781b2b588d09))
* serialize request bodies ([2fd8a40](https://github.com/Increase/increase-go/commit/2fd8a40a82f388b4844bc36e2f5292525575b40f))
* serialize/deserialize time as time.Time ([26cc80e](https://github.com/Increase/increase-go/commit/26cc80e61b9c6e0c0caa17ade22b4abf3d36e0f2))
* simplify request bodies ([800252a](https://github.com/Increase/increase-go/commit/800252a284c7d099e55ecb9a0291ccc1c12a0bfb))
* split types file into smaller files based on resources ([c565382](https://github.com/Increase/increase-go/commit/c56538258709a342bab57ab7eaa5ce360611583b))
* update and fix tests ([8932b42](https://github.com/Increase/increase-go/commit/8932b425d776dcddb15f2d7465bd778072ad1f57))
* update docs ([11de847](https://github.com/Increase/increase-go/commit/11de847403ffe5c849d14fa24d1efd35cbca977d))


### Bug Fixes

* change unknown type generation to `interface{}` ([d6204a9](https://github.com/Increase/increase-go/commit/d6204a9efc808da3de2101487189af98f8929553))
* **client:** pass settings to query marshalling ([fd96af5](https://github.com/Increase/increase-go/commit/fd96af5119e4163e2fb716d6868af68ee79c30dc))
* compile errors and failing tests ([f9651ff](https://github.com/Increase/increase-go/commit/f9651ffc6f50062ff2980c7732e5a9caa78f2e35))
* correct various identifier names ([5e11e46](https://github.com/Increase/increase-go/commit/5e11e468fe5ac0e73a53efb6c4594ee99f7bf93a))
* error that can occur during pagination when there are zero items in the response ([7c0f1f5](https://github.com/Increase/increase-go/commit/7c0f1f5c7807ec6c6d50650b8f977f5d1305a70b))
* error that can occur during pagination when there are zero items in the response ([7c0f1f5](https://github.com/Increase/increase-go/commit/7c0f1f5c7807ec6c6d50650b8f977f5d1305a70b))
* fix struct encoding order ([7f182a7](https://github.com/Increase/increase-go/commit/7f182a72920c4e816d3ce430106432a09df40f90))
* pagination return non-nil on error ([4bf8b95](https://github.com/Increase/increase-go/commit/4bf8b95c0895d7c6b47239b29b21e2df51021a6c))
* pagination return non-nil on error ([4bf8b95](https://github.com/Increase/increase-go/commit/4bf8b95c0895d7c6b47239b29b21e2df51021a6c))
* prefix names of types and files for nested resources ([937e0f0](https://github.com/Increase/increase-go/commit/937e0f0d731503cf0a6ba33ae74022c3e6473912))
* resolve accidental duplications and type errors ([169dd19](https://github.com/Increase/increase-go/commit/169dd19438bb55d9cccc8c34157db667be031787))
* segfault when getting next page if request has no body ([647d9b7](https://github.com/Increase/increase-go/commit/647d9b748a7c932fe8767eddec9cd1b0feb87f56))
* segfault when getting next page if request has no body ([647d9b7](https://github.com/Increase/increase-go/commit/647d9b748a7c932fe8767eddec9cd1b0feb87f56))
* segfault when getting next page if request has no body ([8018f9c](https://github.com/Increase/increase-go/commit/8018f9c43a2b3a7672ed35f3a82f673010618b89))
* update outdate docs in README ([0d434eb](https://github.com/Increase/increase-go/commit/0d434eb68bd408ce15e669508b8ff94b59b27680))
* use default baseurl ([a9230ab](https://github.com/Increase/increase-go/commit/a9230ab6b174f5f56731e8ec36b0018635ec74ed))


### Refactors

* change package structure ([4cbf9d4](https://github.com/Increase/increase-go/commit/4cbf9d4a38e00f8d48852488975ba4cfdfad7885))
* make JSON structs private, rename Metadata-&gt;Field, improve docs ([d0b8800](https://github.com/Increase/increase-go/commit/d0b88007c5ca00ceb1443d756508673ceabb7814))
* rename `.JSON.Extras` -&gt; `.JSON.ExtraFields` ([1cb21b3](https://github.com/Increase/increase-go/commit/1cb21b39cced017c7cbef4a1017e0b412c158ae9))
* rename `field.Field` -&gt; `param.Field` ([a823b52](https://github.com/Increase/increase-go/commit/a823b5257e84ab372027070d2c6c7eae92544c69))
* rename param types ([3de39d4](https://github.com/Increase/increase-go/commit/3de39d4ecab7aa82dc61be3265ebf852f1c27e58))
