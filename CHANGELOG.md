# Changelog

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
