package main

import (
	"tokamak/src/server"
)

func main() {
	/*d := profilegenerator.ProfileData {
	  Name: "blueslimee",
	  AboutMe: "And you push me up to this state of emergency, how beautiful, to be. State of emergency, is where I want to be. State of emergency, how beautiful to be. ",
	  AvatarURL: "https://cdn.discordapp.com/attachments/504668288798949376/771567079538688050/bailey_hat.jpeg",
	  FavColor: "ff0000",
	  Sticker: "kali_uchis_tyrant",
	  Money: "1.78m",
	  Background: "nyc_skyline",
	  Married: true,
	  PartnerName: "Myseljsjsnsksnajamsnkeksf",
	}

	gen.Render()*/
	server.StartServer("1234")
}
