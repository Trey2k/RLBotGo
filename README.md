RLBotGo
===========
[![GoDoc](https://img.shields.io/badge/pkg.go.dev-doc-blue)](http://pkg.go.dev/github.com/Trey2k/RLBotGo)


This repository holds a library for making Rocket League bots in Go!

It provides:

  * An easy to use interface for writing bots
  * An [example bot](https://github.com/Trey2k/RLBotGoExmaple/blob/main/main.go) using this library

Table of Contents:

  * [About](#about)
  * [Todo](#todo)
  * [Usage](#usage)
  * [Contributing](#contributing)
  * [License](#license)

About
-----

This project was made to make it easy to write [RL Bots](https://rlbot.org/) in Go. Instead of using flatbuffer's datatypes, this library converts everything to Go types for ease of use.

Todo
-----

Here are some things that could use some work in this repository's current state:

  * ~~Add support for render groups~~
  * ~~Add support for desired game state~~
  * Add game message support
  * Add some (potentially) useful math functions
  * Get #Go channel in [RLBot Discord](https://discord.com/invite/yc643yyd)

Usage
------------

In order to use this library, you'll need to install and configure the following:

  * [Go](https://golang.org) installed and [configured](https://golang.org/doc/install)
  * [Setup](https://www.youtube.com/watch?v=oXkbizklI2U) [RLBot](https://rlbot.org/)
  * A copy of [Rocket League](https://www.rocketleague.com/) installed
  * Port 23234 availible on your local machine for [RLBot](https://rlbot.org/)
  * A little patience :)

The suggested starting point for using this library is using the [RLBotGoExample](https://github.com/Trey2k/RLBotGoExmaple) repository as a template for your bot.

If you don't start with the example repository, start out with a connection to RLBot:
```Go
	rlBot, err := RLBot.InitConnection(23234)
	if err != nil {
		panic(err)
	}
```
After that, send your bot's ready message:
```Go
	err = rlBot.SendReadyMessage(true, true, true)
	if err != nil {
		panic(err)
	}

```
Call SetGetInput with the name of your desired callback function:
```Go
    rlBot.SetGetInput(tick)
```
Finally, write a function to return the player input every tick:
```Go
// getInput takes in a GameState which contains the gameTickPacket, ballPredidctions, fieldInfo and matchSettings
// it also takes in the RLBot object. And returns a PlayerInput
func getInput(gameState *RLBot.GameState, rlBot *RLBot.RLBot) *RLBot.PlayerInput {
	PlayerInput := &RLBot.PlayerInput{}
	PlayerInput.PlayerIndex = 0

	// Count ball touches up to 10 and on 11 clear the messages and jump
	wasjustTouched := false
	if gameState.GameTick.Ball.LatestTouch.GameSeconds != 0 && lastTouch != gameState.GameTick.Ball.LatestTouch.GameSeconds {
		totalTouches++
		lastTouch = gameState.GameTick.Ball.LatestTouch.GameSeconds
		wasjustTouched = true
	}

	if wasjustTouched && totalTouches <= 10 {
    // DebugMessage is a helper function to let you quickly get debug text on screen. it will autmaicly place it so text will not overlap
		rlBot.DebugMessageAdd(fmt.Sprintf("The ball was touched %d times", totalTouches))
		PlayerInput.ControllerState.Jump = false
	} else if wasjustTouched && totalTouches > 10 {
		rlBot.DebugMessageClear()
		totalTouches = 0
		PlayerInput.ControllerState.Jump = true
	}
	return PlayerInput

}
```

After that, you should have a functional bot!

Contributing
------------

Contributions are always welcome. If you're interested in contributing feel free to submit a PR.

License
-------

This project is currently licensed under the permissive MIT license. Please refer to the [license](/LICENSE) file for more information.
