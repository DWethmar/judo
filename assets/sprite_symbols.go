package assets

import (
	"bytes"
	"image"

	_ "embed"
	"image/png"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	symbolsData, _   = staticSpritesFS.ReadFile("sprites/symbols.png")
	symbolsImg, _    = png.Decode(bytes.NewReader(symbolsData))
	ebitenSymbolsImg = ebiten.NewImageFromImage(symbolsImg)
)

const SymbolsSpriteSize = 16

var (
	symbolsCells = CreateCells(49, 22, SymbolsSpriteSize, SymbolsSpriteSize)
)

func GetSymbolSprite(symbol int) image.Image {
	if symbol < 0 || symbol >= len(symbolsCells) {
		return nil
	}

	return ebitenSymbolsImg.SubImage(symbolsCells[symbol])
}

var (
	Symbol1Sprite  = ebitenSymbolsImg.SubImage(symbolsCells[0])
	Symbol2Sprite  = ebitenSymbolsImg.SubImage(symbolsCells[1])
	Symbol3Sprite  = ebitenSymbolsImg.SubImage(symbolsCells[2])
	Symbol4Sprite  = ebitenSymbolsImg.SubImage(symbolsCells[3])
	Symbol5Sprite  = ebitenSymbolsImg.SubImage(symbolsCells[4])
	Symbol6Sprite  = ebitenSymbolsImg.SubImage(symbolsCells[5])
	Symbol7Sprite  = ebitenSymbolsImg.SubImage(symbolsCells[6])
	Symbol8Sprite  = ebitenSymbolsImg.SubImage(symbolsCells[7])
	Symbol9Sprite  = ebitenSymbolsImg.SubImage(symbolsCells[8])
	Symbol10Sprite = ebitenSymbolsImg.SubImage(symbolsCells[9])
	Symbol11Sprite = ebitenSymbolsImg.SubImage(symbolsCells[10])
	Symbol12Sprite = ebitenSymbolsImg.SubImage(symbolsCells[11])
	Symbol13Sprite = ebitenSymbolsImg.SubImage(symbolsCells[12])
	Symbol14Sprite = ebitenSymbolsImg.SubImage(symbolsCells[13])
	Symbol15Sprite = ebitenSymbolsImg.SubImage(symbolsCells[14])
	Symbol16Sprite = ebitenSymbolsImg.SubImage(symbolsCells[15])
	Symbol17Sprite = ebitenSymbolsImg.SubImage(symbolsCells[16])
	Symbol18Sprite = ebitenSymbolsImg.SubImage(symbolsCells[17])
	Symbol19Sprite = ebitenSymbolsImg.SubImage(symbolsCells[18])
	Symbol20Sprite = ebitenSymbolsImg.SubImage(symbolsCells[19])
	Symbol21Sprite = ebitenSymbolsImg.SubImage(symbolsCells[20])
	Symbol22Sprite = ebitenSymbolsImg.SubImage(symbolsCells[21])
	Symbol23Sprite = ebitenSymbolsImg.SubImage(symbolsCells[22])
	Symbol24Sprite = ebitenSymbolsImg.SubImage(symbolsCells[23])
	Symbol25Sprite = ebitenSymbolsImg.SubImage(symbolsCells[24])
	Symbol26Sprite = ebitenSymbolsImg.SubImage(symbolsCells[25])
	Symbol27Sprite = ebitenSymbolsImg.SubImage(symbolsCells[26])
	Symbol28Sprite = ebitenSymbolsImg.SubImage(symbolsCells[27])
	Symbol29Sprite = ebitenSymbolsImg.SubImage(symbolsCells[28])
	Symbol30Sprite = ebitenSymbolsImg.SubImage(symbolsCells[29])
	Symbol31Sprite = ebitenSymbolsImg.SubImage(symbolsCells[30])
	Symbol32Sprite = ebitenSymbolsImg.SubImage(symbolsCells[31])
	Symbol33Sprite = ebitenSymbolsImg.SubImage(symbolsCells[32])
	Symbol34Sprite = ebitenSymbolsImg.SubImage(symbolsCells[33])
)
