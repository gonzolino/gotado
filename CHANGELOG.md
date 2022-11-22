# Changelog

## [2.0.2](https://github.com/gonzolino/gotado/compare/v2.0.1...v2.0.2) (2022-11-22)


### Bug Fixes

* Do not set temperature when heating is off ([e31e791](https://github.com/gonzolino/gotado/commit/e31e791a6ba2e1118c27f2f64c4f4187036c4755))

### [2.0.1](https://github.com/gonzolino/gotado/compare/v2.0.0...v2.0.1) (2022-04-13)


### Bug Fixes

* Do not crash in away example if away config has no settings ([4d393d3](https://github.com/gonzolino/gotado/commit/4d393d35f6adf42657d8f93d1f2ef094a94535e3))
* fix import path in examples and client ([2237c40](https://github.com/gonzolino/gotado/commit/2237c400b659c331d1d209fd21b559d4d9c07362))
* Set correct v2 module path ([ffafc4f](https://github.com/gonzolino/gotado/commit/ffafc4f1f69cda16b47f73d0a17656b37479d084))

## [2.0.0](https://github.com/gonzolino/gotado/compare/v1.0.0...v2.0.0) (2022-04-13)


### ⚠ BREAKING CHANGES

* users using this package will no longer need to explicitly use the client struct to call the new package API.
* Give the tado client a central put method that all setter functions can rely on.
* Client method `WithTimeout` removed.

### Features

* add function to get devices ([159df85](https://github.com/gonzolino/gotado/commit/159df8550544f0098f51b979727406f4fdd474a8))
* add function to get home state ([0904a9b](https://github.com/gonzolino/gotado/commit/0904a9b6c3f0b3beca3e62e284a678c97f25b0a5))
* add function to get installations ([4ecbddf](https://github.com/gonzolino/gotado/commit/4ecbddf762e0bbdc5969e765db46aab34a15deea))
* add function to get weather information ([af22c42](https://github.com/gonzolino/gotado/commit/af22c425251f9adcd944dbd4979e053aa8d28964))
* add function to get zone capabilities ([e5f6900](https://github.com/gonzolino/gotado/commit/e5f690048b0854739a8aff6ec111054b50eabb8a))
* add function to list users ([547d46a](https://github.com/gonzolino/gotado/commit/547d46aead79f41c9555ff54845552b303df79f0))
* add function to read mobile devices ([61ea3f0](https://github.com/gonzolino/gotado/commit/61ea3f0d075e4c952ec619cbea41a3b781b5193d))
* add functions to control away config ([08a99b4](https://github.com/gonzolino/gotado/commit/08a99b40dde4f94a91031772730c0a90ef3fe63b))
* add functions to control early start settings ([7a47679](https://github.com/gonzolino/gotado/commit/7a47679bd65116da45148741d97ead18ce30d201))
* add functions to control heating via overlays ([3bb962f](https://github.com/gonzolino/gotado/commit/3bb962f4c3dd398d3434b2ed23251e2151ad4991))
* add functions to control presence lock ([0cf0083](https://github.com/gonzolino/gotado/commit/0cf0083cb8d3d53b6fabc9b93fd222ff08950e6e))
* Add functions to manage schedule ([96c0931](https://github.com/gonzolino/gotado/commit/96c0931c09d89310672deabf01fdee37ff260d31))
* Add interface methods to Device and MobileDevice  object ([87787b5](https://github.com/gonzolino/gotado/commit/87787b5b3c72a0fbcc8c01903c769da696925e76))
* Add interface methods to Home object ([cdd2c69](https://github.com/gonzolino/gotado/commit/cdd2c69c688365e5600510acb4f900fdae472c94))
* Add interface methods to ScheduleTimetable object ([819ef45](https://github.com/gonzolino/gotado/commit/819ef45e33ae797a310ebd543ccc5190664faac9))
* Add interface methods to User object ([eab200d](https://github.com/gonzolino/gotado/commit/eab200ded14d483dc5c75ff500fc39449c0c8a2b))
* Add interface methods to Zone object ([32982d4](https://github.com/gonzolino/gotado/commit/32982d4c0ae333791275d0a3cedd3dd4cf3e7e4a))
* add link state info to zone state ([a9ff31d](https://github.com/gonzolino/gotado/commit/a9ff31d25b92c92be16349e9b4fba753c0fa5a0d))
* Add new Tado object ([42681d5](https://github.com/gonzolino/gotado/commit/42681d540eef2b5250a85a08a580130779795d11))
* add some missing fields to models ([f506072](https://github.com/gonzolino/gotado/commit/f5060721d468528673128433ffba1a00193cdac8))
* Add zone HeatingSchedule object ([21f2617](https://github.com/gonzolino/gotado/commit/21f2617d567e95a4d6caf39f7aeff48e4b74a6a7))
* Allow controling mobile device settings ([7336d47](https://github.com/gonzolino/gotado/commit/7336d47b39c0405a85206da7c50a13c970924cf7))
* allow deleting mobile devices ([20fe381](https://github.com/gonzolino/gotado/commit/20fe3812f36c81b9238df4bba3aa7657c322e09c))
* extend client to allow custom request objects ([4b8eb41](https://github.com/gonzolino/gotado/commit/4b8eb4130822ad1968310333ec13453ffedf9eb0))
* introduce constants for timetable types ([0ede991](https://github.com/gonzolino/gotado/commit/0ede991fb4ff9d1c618b7055d884746d07aff7b6))
* introduce proper types and constants for models ([234d625](https://github.com/gonzolino/gotado/commit/234d6257386414856ab39eeecf5e3aa10c37e59a))
* make http client in gotado configurable ([376ceb9](https://github.com/gonzolino/gotado/commit/376ceb94d8cd5e2f0f871f028aac92fc3c3db10a))


### Bug Fixes

* allow pushNotifications to be empty ([1bd1bc3](https://github.com/gonzolino/gotado/commit/1bd1bc3f0efc288d41a5c36ae90f2189a0f3c8bf))
* properly pass context to http client requests ([a222c80](https://github.com/gonzolino/gotado/commit/a222c80a2261e3c07f6948ad83443a6118c3f5b0))


### Code Refactoring

* client is now a private struct ([9b906ad](https://github.com/gonzolino/gotado/commit/9b906ad249631e8d298f6178bbf583e5cee68320))
* remove duplicated put code ([33aabe1](https://github.com/gonzolino/gotado/commit/33aabe14d08aaae16d2f9ebc690f9bf0d7093226))

## [0.3.0](https://www.github.com/gonzolino/gotado/compare/v0.2.0...v0.3.0) (2021-07-14)


### Features

* add function to get devices ([159df85](https://www.github.com/gonzolino/gotado/commit/159df8550544f0098f51b979727406f4fdd474a8))
* add function to get home state ([0904a9b](https://www.github.com/gonzolino/gotado/commit/0904a9b6c3f0b3beca3e62e284a678c97f25b0a5))
* add function to get installations ([4ecbddf](https://www.github.com/gonzolino/gotado/commit/4ecbddf762e0bbdc5969e765db46aab34a15deea))
* add function to get weather information ([af22c42](https://www.github.com/gonzolino/gotado/commit/af22c425251f9adcd944dbd4979e053aa8d28964))
* add function to get zone capabilities ([e5f6900](https://www.github.com/gonzolino/gotado/commit/e5f690048b0854739a8aff6ec111054b50eabb8a))
* add function to list users ([547d46a](https://www.github.com/gonzolino/gotado/commit/547d46aead79f41c9555ff54845552b303df79f0))
* add function to read mobile devices ([61ea3f0](https://www.github.com/gonzolino/gotado/commit/61ea3f0d075e4c952ec619cbea41a3b781b5193d))
* add functions to control away config ([08a99b4](https://www.github.com/gonzolino/gotado/commit/08a99b40dde4f94a91031772730c0a90ef3fe63b))
* add functions to control early start settings ([7a47679](https://www.github.com/gonzolino/gotado/commit/7a47679bd65116da45148741d97ead18ce30d201))
* add functions to control presence lock ([0cf0083](https://www.github.com/gonzolino/gotado/commit/0cf0083cb8d3d53b6fabc9b93fd222ff08950e6e))
* Allow controling mobile device settings ([7336d47](https://www.github.com/gonzolino/gotado/commit/7336d47b39c0405a85206da7c50a13c970924cf7))
* allow deleting mobile devices ([20fe381](https://www.github.com/gonzolino/gotado/commit/20fe3812f36c81b9238df4bba3aa7657c322e09c))


### Bug Fixes

* allow pushNotifications to be empty ([1bd1bc3](https://www.github.com/gonzolino/gotado/commit/1bd1bc3f0efc288d41a5c36ae90f2189a0f3c8bf))

## [0.2.0](https://www.github.com/gonzolino/gotado/compare/v0.1.0...v0.2.0) (2021-05-14)


### Features

* add functions to control heating via overlays ([3bb962f](https://www.github.com/gonzolino/gotado/commit/3bb962f4c3dd398d3434b2ed23251e2151ad4991))
* Add functions to manage schedule ([96c0931](https://www.github.com/gonzolino/gotado/commit/96c0931c09d89310672deabf01fdee37ff260d31))
* extend client to allow custom request objects ([4b8eb41](https://www.github.com/gonzolino/gotado/commit/4b8eb4130822ad1968310333ec13453ffedf9eb0))
