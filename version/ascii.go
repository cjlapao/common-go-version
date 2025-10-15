// Package version provides version parsing, ASCII art generation, and banner printing capabilities.
// It uses the go-figure library which has embedded fonts for portability.
package version

import (
	"strings"

	"github.com/common-nighthawk/go-figure"
)

// FontStyle represents the style of ASCII art font
type FontStyle string

const (
	// FontStyle3D - 3-d font
	FontStyle3D FontStyle = "3-d"
	// FontStyle3x5 - 3x5 font
	FontStyle3x5 FontStyle = "3x5"
	// FontStyle5lineoblique - 5lineoblique font
	FontStyle5lineoblique FontStyle = "5lineoblique"
	// FontStyleAcrobatic - acrobatic font
	FontStyleAcrobatic FontStyle = "acrobatic"
	// FontStyleAlligator - alligator font
	FontStyleAlligator FontStyle = "alligator"
	// FontStyleAlligator2 - alligator2 font
	FontStyleAlligator2 FontStyle = "alligator2"
	// FontStyleAlphabet - alphabet font
	FontStyleAlphabet FontStyle = "alphabet"
	// FontStyleAvatar - avatar font
	FontStyleAvatar FontStyle = "avatar"
	// FontStyleBanner - banner font
	FontStyleBanner FontStyle = "banner"
	// FontStyleBanner3D - banner3-D font
	FontStyleBanner3D FontStyle = "banner3-D"
	// FontStyleBanner3 - banner3 font
	FontStyleBanner3 FontStyle = "banner3"
	// FontStyleBanner4 - banner4 font
	FontStyleBanner4 FontStyle = "banner4"
	// FontStyleBarbwire - barbwire font
	FontStyleBarbwire FontStyle = "barbwire"
	// FontStyleBasic - basic font
	FontStyleBasic FontStyle = "basic"
	// FontStyleBell - bell font
	FontStyleBell FontStyle = "bell"
	// FontStyleBig - big font
	FontStyleBig FontStyle = "big"
	// FontStyleBigchief - bigchief font
	FontStyleBigchief FontStyle = "bigchief"
	// FontStyleBinary - binary font
	FontStyleBinary FontStyle = "binary"
	// FontStyleBlock - block font
	FontStyleBlock FontStyle = "block"
	// FontStyleBubble - bubble font
	FontStyleBubble FontStyle = "bubble"
	// FontStyleBulbhead - bulbhead font
	FontStyleBulbhead FontStyle = "bulbhead"
	// FontStyleCalgphy2 - calgphy2 font
	FontStyleCalgphy2 FontStyle = "calgphy2"
	// FontStyleCaligraphy - caligraphy font
	FontStyleCaligraphy FontStyle = "caligraphy"
	// FontStyleCatwalk - catwalk font
	FontStyleCatwalk FontStyle = "catwalk"
	// FontStyleChunky - chunky font
	FontStyleChunky FontStyle = "chunky"
	// FontStyleCoinstak - coinstak font
	FontStyleCoinstak FontStyle = "coinstak"
	// FontStyleColossal - colossal font
	FontStyleColossal FontStyle = "colossal"
	// FontStyleComputer - computer font
	FontStyleComputer FontStyle = "computer"
	// FontStyleContessa - contessa font
	FontStyleContessa FontStyle = "contessa"
	// FontStyleContrast - contrast font
	FontStyleContrast FontStyle = "contrast"
	// FontStyleCosmic - cosmic font
	FontStyleCosmic FontStyle = "cosmic"
	// FontStyleCosmike - cosmike font
	FontStyleCosmike FontStyle = "cosmike"
	// FontStyleCricket - cricket font
	FontStyleCricket FontStyle = "cricket"
	// FontStyleCursive - cursive font
	FontStyleCursive FontStyle = "cursive"
	// FontStyleCyberlarge - cyberlarge font
	FontStyleCyberlarge FontStyle = "cyberlarge"
	// FontStyleCybermedium - cybermedium font
	FontStyleCybermedium FontStyle = "cybermedium"
	// FontStyleCybersmall - cybersmall font
	FontStyleCybersmall FontStyle = "cybersmall"
	// FontStyleDiamond - diamond font
	FontStyleDiamond FontStyle = "diamond"
	// FontStyleDigital - digital font
	FontStyleDigital FontStyle = "digital"
	// FontStyleDoh - doh font
	FontStyleDoh FontStyle = "doh"
	// FontStyleDoom - doom font
	FontStyleDoom FontStyle = "doom"
	// FontStyleDotmatrix - dotmatrix font
	FontStyleDotmatrix FontStyle = "dotmatrix"
	// FontStyleDrpepper - drpepper font
	FontStyleDrpepper FontStyle = "drpepper"
	// FontStyleEftichess - eftichess font
	FontStyleEftichess FontStyle = "eftichess"
	// FontStyleEftifont - eftifont font
	FontStyleEftifont FontStyle = "eftifont"
	// FontStyleEftipiti - eftipiti font
	FontStyleEftipiti FontStyle = "eftipiti"
	// FontStyleEftirobot - eftirobot font
	FontStyleEftirobot FontStyle = "eftirobot"
	// FontStyleEftitalic - eftitalic font
	FontStyleEftitalic FontStyle = "eftitalic"
	// FontStyleEftiwall - eftiwall font
	FontStyleEftiwall FontStyle = "eftiwall"
	// FontStyleEftiwater - eftiwater font
	FontStyleEftiwater FontStyle = "eftiwater"
	// FontStyleEpic - epic font
	FontStyleEpic FontStyle = "epic"
	// FontStyleFender - fender font
	FontStyleFender FontStyle = "fender"
	// FontStyleFourtops - fourtops font
	FontStyleFourtops FontStyle = "fourtops"
	// FontStyleFuzzy - fuzzy font
	FontStyleFuzzy FontStyle = "fuzzy"
	// FontStyleGoofy - goofy font
	FontStyleGoofy FontStyle = "goofy"
	// FontStyleGothic - gothic font
	FontStyleGothic FontStyle = "gothic"
	// FontStyleGraffiti - graffiti font
	FontStyleGraffiti FontStyle = "graffiti"
	// FontStyleHollywood - hollywood font
	FontStyleHollywood FontStyle = "hollywood"
	// FontStyleInvita - invita font
	FontStyleInvita FontStyle = "invita"
	// FontStyleIsometric1 - isometric1 font
	FontStyleIsometric1 FontStyle = "isometric1"
	// FontStyleIsometric2 - isometric2 font
	FontStyleIsometric2 FontStyle = "isometric2"
	// FontStyleIsometric3 - isometric3 font
	FontStyleIsometric3 FontStyle = "isometric3"
	// FontStyleIsometric4 - isometric4 font
	FontStyleIsometric4 FontStyle = "isometric4"
	// FontStyleItalic - italic font
	FontStyleItalic FontStyle = "italic"
	// FontStyleIvrit - ivrit font
	FontStyleIvrit FontStyle = "ivrit"
	// FontStyleJazmine - jazmine font
	FontStyleJazmine FontStyle = "jazmine"
	// FontStyleJerusalem - jerusalem font
	FontStyleJerusalem FontStyle = "jerusalem"
	// FontStyleKatakana - katakana font
	FontStyleKatakana FontStyle = "katakana"
	// FontStyleKban - kban font
	FontStyleKban FontStyle = "kban"
	// FontStyleLarry3d - larry3d font
	FontStyleLarry3d FontStyle = "larry3d"
	// FontStyleLCD - lcd font
	FontStyleLCD FontStyle = "lcd"
	// FontStyleLean - lean font
	FontStyleLean FontStyle = "lean"
	// FontStyleLetters - letters font
	FontStyleLetters FontStyle = "letters"
	// FontStyleLinux - linux font
	FontStyleLinux FontStyle = "linux"
	// FontStyleLockergnome - lockergnome font
	FontStyleLockergnome FontStyle = "lockergnome"
	// FontStyleMadrid - madrid font
	FontStyleMadrid FontStyle = "madrid"
	// FontStyleMarquee - marquee font
	FontStyleMarquee FontStyle = "marquee"
	// FontStyleMaxfour - maxfour font
	FontStyleMaxfour FontStyle = "maxfour"
	// FontStyleMike - mike font
	FontStyleMike FontStyle = "mike"
	// FontStyleMini - mini font
	FontStyleMini FontStyle = "mini"
	// FontStyleMirror - mirror font
	FontStyleMirror FontStyle = "mirror"
	// FontStyleMnemonic - mnemonic font
	FontStyleMnemonic FontStyle = "mnemonic"
	// FontStyleMorse - morse font
	FontStyleMorse FontStyle = "morse"
	// FontStyleMoscow - moscow font
	FontStyleMoscow FontStyle = "moscow"
	// FontStyleNancyjFancy - nancyj-fancy font
	FontStyleNancyjFancy FontStyle = "nancyj-fancy"
	// FontStyleNancyjUnderlined - nancyj-underlined font
	FontStyleNancyjUnderlined FontStyle = "nancyj-underlined"
	// FontStyleNancyj - nancyj font
	FontStyleNancyj FontStyle = "nancyj"
	// FontStyleNipples - nipples font
	FontStyleNipples FontStyle = "nipples"
	// FontStyleNtgreek - ntgreek font
	FontStyleNtgreek FontStyle = "ntgreek"
	// FontStyleO8 - o8 font
	FontStyleO8 FontStyle = "o8"
	// FontStyleOgre - ogre font
	FontStyleOgre FontStyle = "ogre"
	// FontStylePawp - pawp font
	FontStylePawp FontStyle = "pawp"
	// FontStylePeaks - peaks font
	FontStylePeaks FontStyle = "peaks"
	// FontStylePebbles - pebbles font
	FontStylePebbles FontStyle = "pebbles"
	// FontStylePepper - pepper font
	FontStylePepper FontStyle = "pepper"
	// FontStylePoison - poison font
	FontStylePoison FontStyle = "poison"
	// FontStylePuffy - puffy font
	FontStylePuffy FontStyle = "puffy"
	// FontStylePyramid - pyramid font
	FontStylePyramid FontStyle = "pyramid"
	// FontStyleRectangles - rectangles font
	FontStyleRectangles FontStyle = "rectangles"
	// FontStyleRelief - relief font
	FontStyleRelief FontStyle = "relief"
	// FontStyleRelief2 - relief2 font
	FontStyleRelief2 FontStyle = "relief2"
	// FontStyleRev - rev font
	FontStyleRev FontStyle = "rev"
	// FontStyleRoman - roman font
	FontStyleRoman FontStyle = "roman"
	// FontStyleRot13 - rot13 font
	FontStyleRot13 FontStyle = "rot13"
	// FontStyleRounded - rounded font
	FontStyleRounded FontStyle = "rounded"
	// FontStyleRowancap - rowancap font
	FontStyleRowancap FontStyle = "rowancap"
	// FontStyleRozzo - rozzo font
	FontStyleRozzo FontStyle = "rozzo"
	// FontStyleRunic - runic font
	FontStyleRunic FontStyle = "runic"
	// FontStyleRunyc - runyc font
	FontStyleRunyc FontStyle = "runyc"
	// FontStyleSblood - sblood font
	FontStyleSblood FontStyle = "sblood"
	// FontStyleScript - script font
	FontStyleScript FontStyle = "script"
	// FontStyleSerifcap - serifcap font
	FontStyleSerifcap FontStyle = "serifcap"
	// FontStyleShadow - shadow font
	FontStyleShadow FontStyle = "shadow"
	// FontStyleShort - short font
	FontStyleShort FontStyle = "short"
	// FontStyleSlant - slant font
	FontStyleSlant FontStyle = "slant"
	// FontStyleSlide - slide font
	FontStyleSlide FontStyle = "slide"
	// FontStyleSlscript - slscript font
	FontStyleSlscript FontStyle = "slscript"
	// FontStyleSmall - small font
	FontStyleSmall FontStyle = "small"
	// FontStyleSmisome1 - smisome1 font
	FontStyleSmisome1 FontStyle = "smisome1"
	// FontStyleSmkeyboard - smkeyboard font
	FontStyleSmkeyboard FontStyle = "smkeyboard"
	// FontStyleSmscript - smscript font
	FontStyleSmscript FontStyle = "smscript"
	// FontStyleSmshadow - smshadow font
	FontStyleSmshadow FontStyle = "smshadow"
	// FontStyleSmslant - smslant font
	FontStyleSmslant FontStyle = "smslant"
	// FontStyleSmtengwar - smtengwar font
	FontStyleSmtengwar FontStyle = "smtengwar"
	// FontStyleSpeed - speed font
	FontStyleSpeed FontStyle = "speed"
	// FontStyleStampatello - stampatello font
	FontStyleStampatello FontStyle = "stampatello"
	// FontStyleStandard - standard font (default)
	FontStyleStandard FontStyle = "standard"
	// FontStyleStarwars - starwars font
	FontStyleStarwars FontStyle = "starwars"
	// FontStyleStellar - stellar font
	FontStyleStellar FontStyle = "stellar"
	// FontStyleStop - stop font
	FontStyleStop FontStyle = "stop"
	// FontStyleStraight - straight font
	FontStyleStraight FontStyle = "straight"
	// FontStyleTanja - tanja font
	FontStyleTanja FontStyle = "tanja"
	// FontStyleTengwar - tengwar font
	FontStyleTengwar FontStyle = "tengwar"
	// FontStyleTerm - term font
	FontStyleTerm FontStyle = "term"
	// FontStyleThick - thick font
	FontStyleThick FontStyle = "thick"
	// FontStyleThin - thin font
	FontStyleThin FontStyle = "thin"
	// FontStyleThreepoint - threepoint font
	FontStyleThreepoint FontStyle = "threepoint"
	// FontStyleTicks - ticks font
	FontStyleTicks FontStyle = "ticks"
	// FontStyleTicksslant - ticksslant font
	FontStyleTicksslant FontStyle = "ticksslant"
	// FontStyleTinkerToy - tinker-toy font
	FontStyleTinkerToy FontStyle = "tinker-toy"
	// FontStyleTombstone - tombstone font
	FontStyleTombstone FontStyle = "tombstone"
	// FontStyleTrek - trek font
	FontStyleTrek FontStyle = "trek"
	// FontStyleTsalagi - tsalagi font
	FontStyleTsalagi FontStyle = "tsalagi"
	// FontStyleTwopoint - twopoint font
	FontStyleTwopoint FontStyle = "twopoint"
	// FontStyleUnivers - univers font
	FontStyleUnivers FontStyle = "univers"
	// FontStyleUsaflag - usaflag font
	FontStyleUsaflag FontStyle = "usaflag"
	// FontStyleWavy - wavy font
	FontStyleWavy FontStyle = "wavy"
	// FontStyleWeird - weird font
	FontStyleWeird FontStyle = "weird"
)

// GenerateASCIIArt generates ASCII art for the given text using the default slant font
func GenerateASCIIArt(text string, maxWidth int) []string {
	return GenerateASCIIArtWithStyle(text, maxWidth, FontStyleSlant)
}

// GenerateASCIIArtWithStyle generates ASCII art for the given text with the specified font style
func GenerateASCIIArtWithStyle(text string, maxWidth int, style FontStyle) []string {
	// Handle empty text
	if strings.TrimSpace(text) == "" {
		return []string{}
	}

	// Try rendering the full text first
	fullFigure := figure.NewFigure(text, string(style), true)
	fullRender := fullFigure.String()
	fullWidth := getMaxLineWidth(fullRender)

	// If it fits or no max width specified, return as is
	if maxWidth <= 0 || fullWidth <= maxWidth {
		return splitAndClean(fullRender)
	}

	// Text is too wide - need to break into lines
	words := strings.Fields(text)
	if len(words) <= 1 {
		// Can't break a single word, just return it
		return splitAndClean(fullRender)
	}

	// Find optimal line breaks
	var textLines []string
	var currentWords []string

	for i, word := range words {
		// Try adding this word to current line
		testWords := append([]string{}, currentWords...)
		testWords = append(testWords, word)
		testLine := strings.Join(testWords, " ")

		// Render to check width
		testFigure := figure.NewFigure(testLine, string(style), true)
		testWidth := getMaxLineWidth(testFigure.String())

		if testWidth > maxWidth && len(currentWords) > 0 {
			// This word makes the line too wide, save previous line
			textLines = append(textLines, strings.Join(currentWords, " "))
			currentWords = []string{word}
		} else {
			// Word fits, keep it
			currentWords = testWords
		}

		// Save last line
		if i == len(words)-1 && len(currentWords) > 0 {
			textLines = append(textLines, strings.Join(currentWords, " "))
		}
	}

	// Render each text line and combine
	var result []string
	for i, line := range textLines {
		lineFigure := figure.NewFigure(line, string(style), true)
		lineResult := splitAndClean(lineFigure.String())
		result = append(result, lineResult...)

		// Add spacing between separate text lines
		if i < len(textLines)-1 && len(lineResult) > 0 {
			result = append(result, "")
		}
	}

	return result
}

// splitAndClean splits text into lines and removes trailing/leading whitespace lines
func splitAndClean(text string) []string {
	// Split by newline
	lines := strings.Split(text, "\n")

	// Remove trailing empty lines
	for len(lines) > 0 && strings.TrimSpace(lines[len(lines)-1]) == "" {
		lines = lines[:len(lines)-1]
	}

	// Remove leading empty lines
	for len(lines) > 0 && strings.TrimSpace(lines[0]) == "" {
		lines = lines[1:]
	}

	return lines
}

// getMaxLineWidth returns the maximum width of any line in the text
func getMaxLineWidth(text string) int {
	lines := strings.Split(text, "\n")
	maxWidth := 0
	for _, line := range lines {
		lineLen := len(line)
		if lineLen > maxWidth {
			maxWidth = lineLen
		}
	}
	return maxWidth
}
