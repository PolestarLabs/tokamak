package main

import (
  "tokamak/src/generator"
  "tokamak/src/generator/profile"
)

func main() {
	p := generator.NewGenerator()
	gen := profilegenerator.DefaultProfile {
	  Generator: p,
	  Name: "blueslimee",
	  AboutMe: "And you push me up to this state of emergency, how beautiful to be. State of emergency, is where I want to be. State of emergency, how beautiful to be. ",
	  BackgroundURL: "https://cdn.discordapp.com/attachments/504668288798949376/771450739737755748/MARINA-6912.png",
	  AvatarURL: "https://cdn.discordapp.com/attachments/504668288798949376/771567079538688050/bailey_hat.jpeg",
	  FavColor: "ff0085",
	  HighestRole: "Jesus",
	}
	
	gen.Render()
}