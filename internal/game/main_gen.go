// Machine-translated from SULI.EXE segment 1010 (main program).

package game

// lookupTargetByName  (1010:0048)
func (g *Game) lookupTargetByName(a4 string) {
	var l100 string
	l100 = a4
	g.tgtHP = 0
	g.tgtState = 0
	for g.i = 1; ; g.i++ {
		if g.chars[g.i].name == l100 {
			g.tgtHP = g.chars[g.i].hp
			g.tgtState = g.chars[g.i].state
		}
		if g.i == 30 {
			break
		}
	}
}

// showTextFile  (1010:00b8)
func (g *Game) showTextFile() {
	g.clearTextWin()
	g.resetViewPort()
	g.i = 170
L_00cb:
	g.line = g.txt.ReadLn(255)
	g.SetColor(14)
	if g.line != "***" && g.line != "---" {
		g.OutTextXY(310, g.i, g.line)
	}
	g.i = g.i + 15
	if g.line == "---" {
		g.FlushKeys()
		g.WaitKey()
		g.i = 170
		g.SetFillStyle(1, 0)
		g.Bar(300, 160, 620, 370)
	}
	if g.line == "***" {
		return
	}
	goto L_00cb
}

// drawStatus  (1010:037f)
func (g *Game) drawStatus() {
	g.SetFillStyle(1, 0)
	g.Bar(350, 70, 550, 90)
	g.j = 1
	g.i = 1
	for {
		if g.i <= 25 {
			g.SetColor(4)
		}
		if g.i > 25 && g.i <= 70 {
			g.SetColor(14)
		}
		if g.i > 70 {
			g.SetColor(2)
		}
		g.Line(g.j+351, 74, g.j+351, 86)
		g.i = g.i + 1
		g.j = g.j + 1 + 1
		if g.i >= g.energy {
			break
		}
	}
	g.SetColor(14)
	g.Bar(350, 20, 550, 40)
	if g.target == 0 {
		g.OutTextXY(355, 27, "Nincs")
	}
	if g.target != 0 {
		if g.target == 1 {
			g.tgtName = trunc("A portás", 25)
		}
		if g.target == 2 {
			g.tgtName = trunc("Cigány", 25)
		}
		if g.target == 3 {
			g.tgtName = trunc("Okos, az izomzseni", 25)
		}
		if g.target == 4 {
			g.tgtName = trunc("Brecska, a míveletlen", 25)
		}
		if g.target == 5 {
			g.tgtName = trunc("Kapiczky Arnold", 25)
		}
		if g.target == 6 {
			g.tgtName = trunc("Fejér, a nyelvtudós", 25)
		}
		if g.target == 7 {
			g.tgtName = trunc("Tordy Erzsébet tanárnö", 25)
		}
		if g.target == 8 {
			g.tgtName = trunc("Geri, a skinhead", 25)
		}
		if g.target == 9 {
			g.tgtName = trunc("Agócs, a stréber kisfiú", 25)
		}
		if g.target == 10 {
			g.tgtName = trunc("Laki, a rocker", 25)
		}
		if g.target == 11 {
			g.tgtName = trunc("Zsolt, a nyál", 25)
		}
		if g.target == 12 {
			g.tgtName = trunc("Lepedös Mihály", 25)
		}
		if g.target == 13 {
			g.tgtName = trunc("Ahmed, a szíriai bunkó", 25)
		}
		if g.target == 14 {
			g.tgtName = trunc("Suzy, a suli kurvája", 25)
		}
		if g.target == 15 {
			g.tgtName = trunc("Petra, a bamba", 25)
		}
		if g.target == 16 {
			g.tgtName = trunc("Egy vén fasz tanár", 25)
		}
		if g.target == 17 {
			g.tgtName = trunc("A szakács", 25)
		}
		if g.target == 18 {
			g.tgtName = trunc("A pedellus", 25)
		}
		if g.target == 19 {
			g.tgtName = trunc("Császár, a rejtélyes", 25)
		}
		if g.target == 20 {
			g.tgtName = trunc("Váév, a balfasz", 25)
		}
		if g.target == 21 {
			g.tgtName = trunc("Tibcsi, a jóképü", 25)
		}
		if g.target == 22 {
			g.tgtName = trunc("Egon, a \"fogyatékos\"", 25)
		}
		if g.target == 23 {
			g.tgtName = trunc("Balázs, a nagyon vicces", 25)
		}
		if g.target == 24 {
			g.tgtName = trunc("Széchenyi Gábor", 25)
		}
		if g.target == 25 {
			g.tgtName = trunc("Kutya", 25)
		}
		if g.target == 26 {
			g.tgtName = trunc("Maca, az erotikus", 25)
		}
		if g.target == 27 {
			g.tgtName = trunc("Harami, a melák", 25)
		}
		if g.target == 28 {
			g.tgtName = trunc("Békési, a még melákabb", 25)
		}
		if g.target == 29 {
			g.tgtName = trunc("Az iskolatitkár", 25)
		}
		if g.target == 30 {
			g.tgtName = trunc("A takarító néni", 25)
		}
		g.OutTextXY(355, 27, g.tgtName)
	}
	if g.energy > 0 {
		return
	}
	g.clearMenuWin()
	g.resetViewPort()
	g.SetColor(14)
	g.FlushKeys()
	g.WaitKey()
	g.diedExhausted()
}

// enterRoom  (1010:0e99)
func (g *Game) enterRoom() {
	g.tgtName = ""
	g.target = 0
	g.drawStatus()
	if g.loaded2 == 0 {
		g.tgtName = ""
	}
	if g.loaded2 == 0 {
		g.target = 0
	}
	for g.i = 1; ; g.i++ {
		g.roomPeople[g.i] = ""
		if g.i == 10 {
			break
		}
	}
	g.line = itoa(g.room)
	if g.room < 10 {
		g.line = "0" + g.line
	}
	g.line = "suli" + g.line
	if g.room != 14 && g.room != 25 {
		g.showRoomPic(g.line + ".pic")
	}
	g.txt = g.AssignText("data\\" + g.line + ".dat")
	g.ResetText(g.txt)
	if g.IOResult() != 0 {
		g.DiskError()
	}
	g.i = 1
	g.m2 = 1
	for {
		g.line = g.txt.ReadLn(255)
		if g.line != "Senki" {
			g.roomPeople[g.i] = trunc(g.line, 25)
			g.i = g.i + 1
		}
		g.m2 = g.m2 + 1
		if g.m2 > 10 {
			break
		}
	}
	for g.i = 1; ; g.i++ {
		g.line = g.txt.ReadLn(255)
		g.roomItems[g.i] = ""
		if g.line != "Semmi" {
			g.roomItems[g.i] = trunc(g.line, 25)
		}
		if g.i == 5 {
			break
		}
	}
	for g.i = 1; ; g.i++ {
		g.line = g.txt.ReadLn(255)
		if g.line != "Semmi" {
			g.localMenu[g.i] = trunc(g.line, 25)
		}
		if g.line == "Semmi" {
			g.localMenu[g.i] = ""
		}
		if g.i == 8 {
			break
		}
	}
	g.ResetText(g.txt)
	for {
		g.line = g.txt.ReadLn(255)
		if g.line == "Moving" {
			break
		}
	}
	for g.i = 1; ; g.i++ {
		g.moves[g.i].label = ""
		g.moves[g.i].room = 0
		if g.i == 8 {
			break
		}
	}
	g.i = 1
	for {
		g.line = g.txt.ReadLn(255)
		if g.line != "-----" {
			g.moves[g.i].label = trunc(g.line, 26)
			g.line = g.txt.ReadLn(255)
			g.n = val(g.line, &g.m2)
			g.moves[g.i].room = g.n
			g.i = g.i + 1
		}
		if g.line == "-----" {
			break
		}
	}
	g.CloseText(g.txt)
	g.SetColor(14)
	if g.room == 28 {
		g.sub1e9b_41ba()
	}
	if g.room == 30 {
		g.sub1e9b_3f50()
	}
	if g.room == 14 {
		if g.flagF == 1 {
			g.showRoomPic("suli14.pic")
			g.OutTextXY(310, 170, "Egy lefelé vezetö lépcsö tetején")
			g.OutTextXY(310, 185, "állsz.")
		}
		if g.flagF != 0 {
			goto L_1280
		}
		g.SetFillStyle(1, 0)
		g.Bar(20, 20, 179, 119)
		g.sub1e9b_025b()
	}
L_1280:
	if g.room == 25 {
		g.lookupTargetByName("Az iskolatitkár")
		if g.tgtHP < 5 {
			g.showRoomPic("suli251.pic")
			g.sub1e9b_29f7()
			if g.tgtState == 0 {
				g.OutTextXY(310, 245, "- Az igazgató úr nem fogad. Kifelé!")
			}
			if g.tgtState == 1 {
				g.OutTextXY(310, 245, "- Maga már megint itt van? Kifelé!")
			}
			if g.tgtState == 2 {
				g.OutTextXY(310, 245, "- Azonnal távozzon, vagy kirúgatom!")
			}
			if g.tgtState > 2 {
				g.OutTextXY(310, 245, "- Takarodj kifelé, te taknyos!")
			}
			g.OutTextXY(310, 260, "Egy elektromos kapcsolót látsz a")
			g.OutTextXY(310, 275, "feje fölött, a falon.")
		}
		if g.tgtHP < 5 {
			goto L_133d
		}
		g.showRoomPic("suli252.pic")
		g.sub1e9b_2e02()
	}
L_133d:
	if g.room != 19 {
		goto L_13c6
	}
	if g.flagE == 0 {
		g.sub1e9b_2aac()
	}
	if g.flagE != 1 {
		goto L_13c6
	}
	g.sub1e9b_6542()
	g.i = 0
	for {
		g.i = g.i + 1
		if g.i > 12 {
			goto L_1389
		}
		if g.inv[g.i] == "Gázálarc" {
			break
		}
	}
L_1389:
	if g.i > 12 {
		g.OutTextXY(310, 230, "A gáz hamarosan teljesen szétmarja")
		g.OutTextXY(310, 245, "a tüdödet, és végez veled.")
		g.energy = 0
	}
	if g.i <= 12 {
		g.sub1e9b_8207()
	}
	g.flagE = 0
L_13c6:
	if g.room == 1 {
		g.sub1e9b_2910()
		g.lookupTargetByName("A portás")
		if g.tgtHP >= 5 {
			g.OutTextXY(310, 230, "A halott portás csendesen heverészik")
			g.OutTextXY(310, 245, "a kövön, agyveleje csurdogál.")
		}
		if g.tgtHP >= 5 {
			goto L_142f
		}
		if g.tgtState != 0 {
			goto L_142f
		}
		g.OutTextXY(310, 230, "Hirtelen az iskola portása lép eléd.")
		g.OutTextXY(310, 245, "- Nem mehetsz be tanítás alatt.")
	}
L_142f:
	if g.room == 2 {
		g.sub1e9b_5eb8()
	}
	if g.room == 24 {
		g.sub1e9b_7c95()
		g.lookupTargetByName("A takarító néni")
		if g.tgtHP >= 5 {
			g.OutTextXY(310, 230, "A takarító néni hullája a kiömlött")
			g.OutTextXY(310, 245, "felmosóvízben hever a folyosó közepén.")
		}
		if g.tgtHP >= 5 {
			goto L_1514
		}
		g.OutTextXY(310, 230, "A takarító néni unottan mossa a")
		g.OutTextXY(310, 245, "folyosó kövét. Lépteidet hallva")
		g.OutTextXY(310, 260, "felnéz.")
		if g.tgtState <= 1 {
			g.OutTextXY(310, 275, "- Ne császkáld össze a folyosót,")
			g.OutTextXY(310, 290, "  kisapám, mert eltöröm a lábadat!")
		}
		if g.tgtState > 1 && g.tgtState <= 4 {
			g.sub1e9b_7d3b()
		}
		if g.tgtState < 5 {
			goto L_1514
		}
		g.OutTextXY(310, 275, "- Tünsz el innen! - suhogtatja meg a")
		g.OutTextXY(310, 290, "  partvisnyelet.")
	}
L_1514:
	if g.room == 23 {
		g.sub1e9b_601a()
	}
	if g.room == 15 {
		g.sub1e9b_5885()
	}
	if g.room == 18 {
		g.sub1e9b_5d4f()
	}
	if g.room == 20 {
		g.sub1e9b_5b14()
	}
	if g.room == 13 {
		g.sub1e9b_5bf0()
	}
	if g.room == 17 {
		g.sub1e9b_3e08()
	}
	if g.room == 27 {
		g.sub1e9b_4052()
	}
	if g.room != 3 {
		goto L_1624
	}
	g.OutTextXY(310, 170, "A félemeleti lépcsöfordulóban állsz.")
	g.lookupTargetByName("Cigány")
	if g.tgtHP >= 5 {
		g.OutTextXY(310, 185, "Egy döglött roma hever a padlón.")
	}
	if g.tgtHP >= 5 {
		goto L_1624
	}
	if g.tgtState == 0 {
		g.OutTextXY(310, 185, "Egy cigány ül a lépcsön.")
	}
	if g.tgtState != 1 {
		if g.tgtState != 2 {
			goto L_15ed
		}
	}
	g.OutTextXY(310, 185, "Jöttödre a roma felnéz.")
	g.OutTextXY(310, 200, "- Mán megin itt vagy? - jegyzi meg halkan.")
L_15ed:
	if g.tgtState > 3 {
		g.OutTextXY(310, 185, "A cigány felugrik, amint meglát.")
		g.OutTextXY(310, 200, "- Takaroggyá bazsmeg, mer mindzsá")
		g.OutTextXY(310, 215, "  sétverem a pofádat!")
	}
L_1624:
	if g.room != 22 {
		goto L_1681
	}
	g.sub1e9b_23d9()
	g.lookupTargetByName("Harami, a melák")
	if g.tgtHP >= 5 {
		g.OutTextXY(310, 215, "A két melák békésen oszladozik.")
	}
	if g.tgtHP >= 5 {
		goto L_1681
	}
	if g.tgtState == 0 {
		g.sub1e9b_271b()
	}
	if g.tgtState != 1 {
		if g.tgtState != 2 {
			goto L_1675
		}
	}
	g.sub1e9b_268f()
L_1675:
	if g.tgtState > 3 {
		g.sub1e9b_279b()
	}
L_1681:
	if g.room == 21 {
		g.sub1e9b_2844()
		g.lookupTargetByName("Maca, az erotikus")
		if g.tgtHP >= 5 {
			g.OutTextXY(310, 230, "egy szexis hulla hever.")
		}
		if g.tgtHP >= 5 {
			goto L_16d3
		}
		g.OutTextXY(310, 230, "éppen Maca, az erotikus takarítja")
		g.OutTextXY(310, 245, "habtestét.")
	}
L_16d3:
	if g.room == 4 {
		g.sub1e9b_599c()
	}
	if g.room == 16 {
		g.sub1e9b_24cd()
		g.lookupTargetByName("Kutya")
		if g.tgtHP >= 5 {
			goto L_16ff
		}
		g.sub1e9b_25de()
	}
L_16ff:
	if g.room == 5 {
		g.sub1e9b_2227()
		g.lookupTargetByName("Tordy Erzsébet tanárnö")
		if g.tgtHP >= 5 {
			g.OutTextXY(310, 260, "Szerencsére más nyom a tanárnö ")
			g.OutTextXY(310, 275, "hulláján kívül nincs.")
		}
		if g.tgtHP >= 5 {
			goto L_1759
		}
		if g.tgtState == 0 {
			g.sub1e9b_232a()
		}
		if g.tgtState < 1 {
			goto L_1759
		}
		g.sub1e9b_1ff7()
	}
L_1759:
	if g.room == 12 {
		g.sub1e9b_20f4()
		g.lookupTargetByName("Széchenyi Gábor")
		if g.tgtHP >= 5 {
			g.OutTextXY(310, 260, "Kár, hogy meghalt.")
		}
		if g.tgtHP >= 5 {
			goto L_17ec
		}
		if g.tgtState == 0 {
			g.OutTextXY(310, 260, "Most éppen a 2.B. osztály tanítja öt")
			g.OutTextXY(310, 275, "a Windows alapjaira.")
		}
		if g.tgtState < 1 {
			goto L_17ec
		}
		g.OutTextXY(310, 260, "- Te már megint itt vagy? - néz fel,")
		g.OutTextXY(310, 275, "amint benyitsz. - Már mehetsz is")
		g.OutTextXY(310, 290, "kifelé!")
	}
L_17ec:
	if g.room == 7 {
		g.sub1e9b_1f0a()
		g.lookupTargetByName("Lepedös Mihály")
		if g.tgtHP >= 5 {
			g.OutTextXY(310, 290, "Most azonban nem játszik semmit,")
			g.OutTextXY(310, 305, "élettelenül hever a katedra mellett.")
		}
		if g.tgtHP >= 5 {
			goto L_1859
		}
		g.lookupTargetByName("Lepedös Mihály")
		if g.tgtState == 0 {
			g.OutTextXY(310, 290, "Ilyenkor a legnevetségesebb.")
		}
		if g.tgtState < 1 {
			goto L_1859
		}
		g.sub1e9b_0396()
	}
L_1859:
	if g.room == 8 {
		g.sub1e9b_581d()
	}
	if g.room == 9 {
		g.sub1e9b_577c()
	}
	if g.room == 10 {
		g.sub1e9b_7e10()
		g.lookupTargetByName("Geri, a skinhead")
		if g.tgtHP >= 5 {
			g.OutTextXY(310, 245, "Szerencsére már ezekböl is eggyel")
			g.OutTextXY(310, 260, "kevesebb van.")
		}
		if g.tgtHP >= 5 {
			goto L_18d3
		}
		g.OutTextXY(310, 245, "Most éppen Geri, a suli skinheadje")
		g.OutTextXY(310, 260, "élvezi itt a bagóját.")
	}
L_18d3:
	if g.room == 11 {
		g.sub1e9b_80ad()
		g.lookupTargetByName("A szakács")
		if g.tgtHP >= 5 {
			g.OutTextXY(310, 290, "Szerencsére a szakácsot már kinyírtad.")
		}
		if g.tgtHP >= 5 {
			goto L_190a
		}
		g.sub1e9b_0422()
	}
L_190a:
	if g.room != 6 {
		return
	}
	g.sub1e9b_7f29()
	g.lookupTargetByName("A pedellus")
	if g.tgtHP >= 5 {
		g.OutTextXY(310, 260, "A pedellus holtteste a fal mellett")
		g.OutTextXY(310, 275, "hever.")
	}
	if g.tgtHP >= 5 {
		return
	}
	g.OutTextXY(310, 260, "A falnál a pedellus támaszkodik, és")
	g.OutTextXY(310, 275, "nagy nyugodtan pipázik.")
}

// loadGame  (1010:1982)
func (g *Game) loadGame() {
	g.txt = g.AssignText("data\\savegame.tsi")
	g.ResetText(g.txt)
	g.room = g.txt.ReadInt()
	g.energy = g.txt.ReadInt()
	g.target = g.txt.ReadInt()
	g.tgtName = g.txt.ReadLn(25)
	for g.i = 1; ; g.i++ {
		g.chars[g.i].name = g.txt.ReadLn(25)
		g.chars[g.i].hp = g.txt.ReadInt()
		g.chars[g.i].state = g.txt.ReadInt()
		if g.i == 30 {
			break
		}
	}
	for g.i = 1; ; g.i++ {
		g.inv[g.i] = g.txt.ReadLn(25)
		if g.i == 12 {
			break
		}
	}
	for g.i = 1; ; g.i++ {
		g.taken[g.i] = g.txt.ReadLn(25)
		if g.i == 120 {
			break
		}
	}
	g.flagD = 1
	g.line = g.txt.ReadLn(255)
	if g.line == "0" {
		g.flagD = 0
	}
	g.flagF = 1
	g.line = g.txt.ReadLn(255)
	if g.line == "0" {
		g.flagF = 0
	}
	g.flag0 = 1
	g.line = g.txt.ReadLn(255)
	if g.line == "0" {
		g.flag0 = 0
	}
	g.flagE = 1
	g.line = g.txt.ReadLn(255)
	if g.line == "0" {
		g.flagE = 0
	}
	g.CloseText(g.txt)
	g.loaded = 1
	g.loaded2 = 1
	if g.IOResult() == 0 {
		return
	}
	g.loaded = 0
	g.loaded2 = 0
}

// gameInit  (1010:1dcc)
func (g *Game) gameInit() {
	g.loaded2 = 0
	g.loaded = 0
	g.FindFirst("data\\savegame.tsi")
	if g.DosError == 0 && g.ParamStr(1) == "loadgame" {
		g.loadGame()
	}
	if g.DosError == 0 && g.ParamStr(2) == "loadgame" {
		g.loadGame()
	}
	if g.DosError == 0 && g.ParamStr(3) == "loadgame" {
		g.loadGame()
	}
	if g.loaded == 1 {
		g.drawStatus()
	}
	if g.loaded == 1 {
		g.enterRoom()
	}
	if g.loaded == 1 {
		return
	}
	g.flagD = 0
	g.flagF = 0
	g.flag0 = 0
	g.flagE = 1
	g.energy = 100
	g.target = 0
	g.room = 1
	for g.i = 1; ; g.i++ {
		g.roomPeople[g.i] = ""
		if g.i == 10 {
			break
		}
	}
	for g.i = 1; ; g.i++ {
		g.inv[g.i] = ""
		if g.i == 12 {
			break
		}
	}
	for g.i = 1; ; g.i++ {
		g.taken[g.i] = ""
		if g.i == 120 {
			break
		}
	}
	for g.i = 1; ; g.i++ {
		g.chars[g.i].name = ""
		g.chars[g.i].hp = 0
		g.chars[g.i].state = 0
		if g.i == 30 {
			break
		}
	}
	g.chars[1].name = trunc("A portás", 25)
	g.chars[2].name = trunc("Cigány", 25)
	g.chars[3].name = trunc("Okos, az izomzseni", 25)
	g.chars[4].name = trunc("Brecska, a míveletlen", 25)
	g.chars[5].name = trunc("Kapiczky Arnold", 25)
	g.chars[6].name = trunc("Fejér, a nyelvtudós", 25)
	g.chars[7].name = trunc("Tordy Erzsébet tanárnö", 25)
	g.chars[8].name = trunc("Geri, a skinhead", 25)
	g.chars[9].name = trunc("Agócs, a stréber kisfiú", 25)
	g.chars[10].name = trunc("Laki, a rocker", 25)
	g.chars[11].name = trunc("Zsolt, a nyál", 25)
	g.chars[12].name = trunc("Lepedös Mihály", 25)
	g.chars[13].name = trunc("Ahmed, a szíriai bunkó", 25)
	g.chars[14].name = trunc("Suzy, a suli kurvája", 25)
	g.chars[15].name = trunc("Petra, a bamba", 25)
	g.chars[16].name = trunc("Egy vén fasz tanár", 25)
	g.chars[17].name = trunc("A szakács", 25)
	g.chars[18].name = trunc("A pedellus", 25)
	g.chars[19].name = trunc("Császár, a rejtélyes", 25)
	g.chars[20].name = trunc("Váév, a balfasz", 25)
	g.chars[21].name = trunc("Tibcsi, a jóképü", 25)
	g.chars[22].name = trunc("Egon, a \"fogyatékos\"", 25)
	g.chars[23].name = trunc("Balázs, a nagyon vicces", 25)
	g.chars[24].name = trunc("Széchenyi Gábor", 25)
	g.chars[25].name = trunc("Kutya", 25)
	g.chars[26].name = trunc("Maca, az erotikus", 25)
	g.chars[27].name = trunc("Harami, a melák", 25)
	g.chars[28].name = trunc("Békési, a még melákabb", 25)
	g.chars[29].name = trunc("Az iskolatitkár", 25)
	g.chars[30].name = trunc("A takarító néni", 25)
	g.showRoomPic("suli00.pic")
	g.drawStatus()
	g.txt = g.AssignText("data\\suli01.txt")
	g.ResetText(g.txt)
	if g.IOResult() != 0 {
		g.DiskError()
	}
	g.showTextFile()
	g.CloseText(g.txt)
	g.enterRoom()
}

// selectTarget  (1010:216a)
func (g *Game) selectTarget() {
	g.clearMenuWin()
	g.resetViewPort()
	g.i = 1
L_217d:
	for g.n = 1; ; g.n++ {
		g.SetColor(14)
		for g.m2 = 1; ; g.m2++ {
			if g.roomPeople[g.n] == g.chars[g.m2].name && g.chars[g.m2].hp >= 5 {
				g.SetColor(7)
			}
			if g.m2 == 30 {
				break
			}
		}
		g.OutTextXY(25, g.n*15+150, g.roomPeople[g.n])
		if g.n == 8 {
			break
		}
	}
	g.SetFillStyle(1, 21)
	g.Bar(22, g.i*15+150-1-1, 248, g.i*15+150+10)
	g.SetColor(14)
	for g.m2 = 1; ; g.m2++ {
		if g.roomPeople[g.i] == g.chars[g.m2].name && g.chars[g.m2].hp >= 5 {
			g.SetColor(7)
		}
		if g.m2 == 30 {
			break
		}
	}
	g.OutTextXY(25, g.i*15+150, g.roomPeople[g.i])
	g.FlushKeys()
	for {
		if g.KeyPressed() {
			break
		}
	}
	g.SetFillStyle(1, 0)
	g.FloodFill(22, g.i*15+150, 0)
	g.key = g.ReadKey()
	if g.key == 0 {
		g.key = g.ReadKey()
		if g.key == 80 {
			g.i = g.i + 1
			if g.i > 8 {
				g.i = 1
			}
			if g.roomPeople[g.i] == "" {
				for {
					g.i = g.i + 1
					if g.i > 8 {
						g.i = 1
					}
					if g.roomPeople[g.i] != "" {
						break
					}
				}
			}
			g.key = 0
		}
		if g.key != 72 {
			goto L_235f
		}
		g.i = g.i - 1
		if g.i < 1 {
			g.i = 8
		}
		if g.roomPeople[g.i] == "" {
			for {
				if g.i < 1 {
					g.i = 8
				}
				g.i = g.i - 1
				if g.roomPeople[g.i] != "" {
					break
				}
			}
		}
		g.key = 0
	}
L_235f:
	if g.i < 1 {
		g.i = 8
	}
	if g.i > 8 {
		g.i = 1
	}
	if g.key != 13 && g.key != 27 {
		goto L_217d
	}
	if g.key == 27 {
		return
	}
	if g.room == 1 {
		g.target = 1
	}
	if g.room == 3 {
		g.target = 2
	}
	if g.room == 5 && g.i == 1 {
		g.target = 3
	}
	if g.room == 5 && g.i == 2 {
		g.target = 4
	}
	if g.room == 5 && g.i == 3 {
		g.target = 5
	}
	if g.room == 5 && g.i == 4 {
		g.target = 6
	}
	if g.room == 5 && g.i == 5 {
		g.target = 7
	}
	if g.room == 10 {
		g.target = 8
	}
	if g.room == 7 && g.i == 1 {
		g.target = 9
	}
	if g.room == 7 && g.i == 2 {
		g.target = 10
	}
	if g.room == 7 && g.i == 3 {
		g.target = 11
	}
	if g.room == 7 && g.i == 4 {
		g.target = 12
	}
	if g.room == 9 && g.i == 1 {
		g.target = 13
	}
	if g.room == 9 && g.i == 2 {
		g.target = 14
	}
	if g.room == 9 && g.i == 3 {
		g.target = 15
	}
	if g.room == 9 && g.i == 4 {
		g.target = 16
	}
	if g.room == 11 {
		g.target = 17
	}
	if g.room == 6 {
		g.target = 18
	}
	if g.room == 12 {
		g.target = g.i + 18
	}
	if g.room == 16 {
		g.target = 25
	}
	if g.room == 21 {
		g.target = 26
	}
	if g.room == 22 {
		g.target = g.i + 26
	}
	if g.room == 25 {
		g.target = 29
	}
	if g.room == 24 {
		g.target = 30
	}
	g.drawStatus()
}

// spit  (1010:2548)
func (g *Game) spit() {
	g.line = itoa(g.target)
	if g.target < 10 {
		g.line = "0" + g.line
	}
	g.txt = g.AssignText("data\\trg" + g.line + ".dat")
	g.ResetText(g.txt)
	if g.IOResult() != 0 {
		g.DiskError()
	}
	g.line = itoa(g.i)
	g.line = "*spit"
	for {
		g.line2 = g.txt.ReadLn(255)
		if g.line == g.line2 {
			break
		}
	}
	g.showTextFile()
	g.line2 = g.txt.ReadLn(255)
	g.n = val(g.line2, &g.m2)
	g.m2 = 1
	for {
		if g.chars[g.m2].name == g.tgtName {
			g.chars[g.m2].hp = g.chars[g.m2].hp + g.n
		}
		g.m2 = g.m2 + 1
		if g.m2 > 30 {
			break
		}
	}
	g.line2 = g.txt.ReadLn(255)
	g.n = val(g.line2, &g.m2)
	g.m2 = 1
	for {
		if g.chars[g.m2].name == g.tgtName {
			g.chars[g.m2].state = g.chars[g.m2].state + g.n
		}
		g.m2 = g.m2 + 1
		if g.m2 > 30 {
			break
		}
	}
	g.line2 = g.txt.ReadLn(255)
	g.n = val(g.line2, &g.m2)
	g.energy = g.energy - g.n
	g.CloseText(g.txt)
}

// burp  (1010:2b4c)
func (g *Game) burp() {
	g.clearTextWin()
	g.resetViewPort()
	g.SetColor(14)
	if g.room == 1 {
		g.OutTextXY(310, 170, "Bár néhány versenyen díjnyertes lett")
		g.OutTextXY(310, 185, "volna, amit kiengedtél magadból, de")
		g.lookupTargetByName("A portás")
		if g.tgtHP >= 5 {
			g.OutTextXY(310, 200, "a halott portás nem túl lelkes")
			g.OutTextXY(310, 215, "közönség.")
		}
		if g.tgtHP >= 5 {
			goto L_2bdd
		}
		g.OutTextXY(310, 200, "a portást ezzel nem nagyon tudod")
		g.OutTextXY(310, 215, "felizgatni.")
	}
L_2bdd:
	if g.room == 25 {
		g.OutTextXY(310, 170, "Kieresztesz magadból mindent, ami")
		g.OutTextXY(310, 185, "felül kijöhet.")
		g.lookupTargetByName("Az iskolatitkár")
		if g.tgtHP >= 5 {
			g.OutTextXY(310, 200, "A szenes hullát, ami valaha Zsuzsás")
			g.OutTextXY(310, 215, "volt, nem zavarod nyugalmában.")
		}
		if g.tgtHP >= 5 {
			goto L_2c9b
		}
		g.OutTextXY(310, 200, "- Otthon disznólkodj, te taknyos,")
		g.OutTextXY(310, 215, "  a drága jó anyukádnak!")
		for g.i = 1; ; g.i++ {
			if g.chars[g.i].name == "Az iskolatitkár" {
				g.chars[g.i].hp = g.chars[g.i].hp + 1 + 1
			}
			if g.i == 30 {
				break
			}
		}
	}
L_2c9b:
	if g.room == 6 {
		g.OutTextXY(310, 170, "A falak beleremegnek a hangodba, de")
		g.lookupTargetByName("A pedellus")
		if g.tgtHP >= 5 {
			g.OutTextXY(310, 185, "ha ezzel a halott pedelluson akartál")
			g.OutTextXY(310, 200, "ezzel segíteni, nem sokat értél vele.")
		}
		if g.tgtHP >= 5 {
			goto L_2d08
		}
		g.OutTextXY(310, 185, "a pedellus most sem igazán érdeklödik")
		g.OutTextXY(310, 200, "irántad.")
	}
L_2d08:
	if g.room == 22 {
		g.OutTextXY(310, 170, "- ÖARRGF!")
		g.lookupTargetByName("Harami, a melák")
		if g.tgtHP < 5 {
			g.OutTextXY(310, 185, "- Mi van? Már ennyi az idö? - néz fel")
			g.OutTextXY(310, 200, "Harami.")
		}
		if g.tgtHP < 5 {
			goto L_2d75
		}
		g.OutTextXY(310, 185, "A két hullát nem érdekli az")
		g.OutTextXY(310, 200, "öklendezésed.")
	}
L_2d75:
	if g.room == 21 {
		g.OutTextXY(310, 170, "- BAARF! - jegyzed meg halkan.")
		g.lookupTargetByName("Maca, az erotikus")
		if g.tgtHP >= 5 {
			g.OutTextXY(310, 185, "A halott Macát ez nemigen izgatja.")
		}
		if g.tgtHP >= 5 {
			goto L_2dd2
		}
		g.OutTextXY(310, 185, "- Mi van, apukám, a kultúra kivágta")
		g.OutTextXY(310, 200, "  a biztosítékot?")
	}
L_2dd2:
	if g.room == 2 {
		g.sub1e9b_50fe()
	}
	if g.room == 24 {
		g.sub1e9b_5180()
	}
	if g.room == 20 {
		g.sub1e9b_50fe()
	}
	if g.room == 17 {
		g.sub1e9b_552f()
	}
	if g.room == 27 {
		g.sub1e9b_53ed()
	}
	if g.room == 28 {
		g.sub1e9b_5414()
	}
	if g.room == 30 {
		g.sub1e9b_5253()
	}
	if g.room == 3 {
		g.sub1e9b_51f0()
	}
	if g.room == 15 {
		g.sub1e9b_52b3()
	}
	if g.room == 19 {
		g.sub1e9b_5339()
	}
	if g.room == 23 {
		g.OutTextXY(310, 170, "Semmi hatás.")
	}
	if g.room != 4 && g.room != 8 && g.room != 13 {
		if g.room != 26 {
			goto L_2e82
		}
	}
	g.sub1e9b_53a6()
L_2e82:
	if g.room == 5 {
		g.lookupTargetByName("Tordy Erzsébet tanárnö")
		if g.tgtHP < 5 {
			g.sub1e9b_063a()
			for g.i = 1; ; g.i++ {
				if g.chars[g.i].name == "Tordy Erzsébet tanárnö" {
					g.chars[g.i].state = g.chars[g.i].state + 1 + 1
				}
				if g.i == 30 {
					break
				}
			}
		}
		if g.tgtHP < 5 {
			goto L_2ef2
		}
		g.OutTextXY(310, 170, "Senkit sem érdekel a produkciód.")
	}
L_2ef2:
	if g.room == 12 {
		g.lookupTargetByName("Széchenyi Gábor")
		if g.tgtHP < 5 {
			g.OutTextXY(310, 170, "- Mi van, fiam, nem volt gyerekszobád?")
			g.OutTextXY(310, 185, " - kérdezi Széchenyi tanár úr.")
			for g.i = 1; ; g.i++ {
				if g.chars[g.i].name == "Széchenyi Gábor" {
					g.chars[g.i].state = g.chars[g.i].state + 1
				}
				if g.i == 30 {
					break
				}
			}
		}
		if g.tgtHP < 5 {
			goto L_2f7f
		}
		g.OutTextXY(310, 170, "Semmi hatás.")
	}
L_2f7f:
	if g.room == 10 {
		g.OutTextXY(310, 170, "Nem túl jó a hangzás. Még gyakorold.")
	}
	if g.room == 7 {
		g.sub1e9b_546a()
	}
	if g.room == 9 {
		g.sub1e9b_54c3()
	}
	if g.room == 18 {
		g.sub1e9b_55d2()
	}
	if g.room == 14 {
		g.sub1e9b_5659()
	}
	if g.room == 16 {
		g.sub1e9b_56ed()
	}
	if g.room != 11 {
		return
	}
	g.lookupTargetByName("A szakács")
	if g.tgtHP < 5 {
		g.OutTextXY(310, 170, "- Mi van, fiacskám, ízlett a kaja?")
	}
	if g.tgtHP < 5 {
		return
	}
	g.OutTextXY(310, 170, "Semmi hatás.")
}

// rape  (1010:3027)
func (g *Game) rape() {
	g.line = itoa(g.target)
	if g.target < 10 {
		g.line = "0" + g.line
	}
	g.txt = g.AssignText("data\\trg" + g.line + ".dat")
	g.ResetText(g.txt)
	if g.IOResult() != 0 {
		g.DiskError()
	}
	g.line = itoa(g.i)
	g.line = "*fuck"
	for {
		g.line2 = g.txt.ReadLn(255)
		if g.line == g.line2 {
			break
		}
	}
	g.showTextFile()
	g.line2 = g.txt.ReadLn(255)
	g.n = val(g.line2, &g.m2)
	g.m2 = 1
	for {
		if g.chars[g.m2].name == g.tgtName {
			g.chars[g.m2].hp = g.chars[g.m2].hp + g.n
		}
		g.m2 = g.m2 + 1
		if g.m2 > 30 {
			break
		}
	}
	g.line2 = g.txt.ReadLn(255)
	g.n = val(g.line2, &g.m2)
	g.m2 = 1
	for {
		if g.chars[g.m2].name == g.tgtName {
			g.chars[g.m2].state = g.chars[g.m2].state + g.n
		}
		g.m2 = g.m2 + 1
		if g.m2 > 30 {
			break
		}
	}
	g.line2 = g.txt.ReadLn(255)
	g.n = val(g.line2, &g.m2)
	g.energy = g.energy - g.n
	g.CloseText(g.txt)
}

// waitAround  (1010:3ae4)
func (g *Game) waitAround() {
	g.clearTextWin()
	g.resetViewPort()
	g.Randomize()
	g.j = g.Random(5) + 1
	if g.room == 1 {
		g.lookupTargetByName("A portás")
		if g.tgtHP >= 5 {
			g.OutTextXY(310, 170, "Nem történik semmi.")
		}
		if g.tgtHP >= 5 {
			goto L_3b8d
		}
		if g.tgtState <= 1 {
			g.OutTextXY(310, 170, "A portás várja, hogy távozz.")
		}
		if g.tgtState > 1 && g.tgtState <= 3 {
			g.OutTextXY(310, 170, "A portás nagyon szeretné, hogy végre")
			g.OutTextXY(310, 185, "eltakarodj innen...")
		}
		if g.tgtState <= 3 {
			goto L_3b8d
		}
		g.OutTextXY(310, 170, "- Na mi lesz, eltakarodsz végre?!")
	}
L_3b8d:
	if g.room == 22 {
		g.lookupTargetByName("Harami, a melák")
		if g.tgtHP >= 5 {
			g.OutTextXY(310, 170, "Történik: semmi.")
		}
		if g.tgtHP >= 5 {
			goto L_3c30
		}
		if g.tgtState <= 1 {
			g.OutTextXY(310, 170, "A két meláknak terhére vagy.")
		}
		if g.tgtState > 1 && g.tgtState <= 3 {
			g.OutTextXY(310, 170, "- Te, öreg, mi lenne, ha elmennél")
			g.OutTextXY(310, 185, "  végre a nyomorba?")
		}
		if g.tgtState > 3 {
			g.OutTextXY(310, 170, "- Menjél már innen a rákba!")
		}
		if g.tgtState <= 3 {
			goto L_3c30
		}
		g.OutTextXY(310, 185, "Harami kedves, mint egy atomcsapás.")
	}
L_3c30:
	if g.room == 25 {
		g.lookupTargetByName("Az iskolatitkár")
		if g.tgtHP >= 5 {
			g.OutTextXY(310, 170, "Nem történik semmi.")
		}
		if g.tgtHP >= 5 {
			goto L_3cbc
		}
		if g.tgtState <= 1 {
			g.OutTextXY(310, 170, "- Azt mondtam, menj innen!")
		}
		if g.tgtState > 1 && g.tgtState <= 3 {
			g.OutTextXY(310, 170, "Zsuzsás nagyon szeretné, hogy végre")
			g.OutTextXY(310, 185, "eltakarodj innen...")
		}
		if g.tgtState <= 3 {
			goto L_3cbc
		}
		g.OutTextXY(310, 170, "- És keddre kéretem a szüleidet!")
	}
L_3cbc:
	if g.room != 2 {
		if g.room != 28 {
			goto L_3cda
		}
	}
	g.OutTextXY(310, 170, "Semmi nem történik.")
L_3cda:
	if g.room == 24 {
		g.OutTextXY(310, 170, "Bentröl írógép kopogását hallod.")
	}
	if g.room != 13 {
		if g.room != 23 {
			goto L_3d1f
		}
	}
	g.OutTextXY(310, 170, "Csámcsogást hallasz. Körülnézel, és")
	g.OutTextXY(310, 185, "megnyugszol. Semmi, csak a négy fal.")
L_3d1f:
	if g.room == 27 {
		g.sub1e9b_1354()
	}
	if g.room == 20 {
		g.sub1e9b_13b6()
	}
	if g.room == 30 {
		g.sub1e9b_140f()
	}
	if g.room == 19 {
		g.sub1e9b_12d7()
	}
	if g.room == 17 {
		g.sub1e9b_1267()
	}
	if g.room == 10 {
		g.OutTextXY(310, 170, "A hely is unalmas, meg a szaga is.")
	}
	if g.room == 3 {
		g.lookupTargetByName("Cigány")
		if g.tgtState == 0 && g.tgtHP < 5 {
			g.OutTextXY(310, 170, "Semmi érdekes nem történik.")
		}
		if g.tgtState > 0 && g.tgtHP < 5 {
			g.OutTextXY(310, 170, "A romának terhére vagy.")
		}
		if g.tgtHP < 5 {
			goto L_3de4
		}
		g.OutTextXY(310, 170, "Unalmas dolog egy megmurdelt cigány")
		g.OutTextXY(310, 185, "mellett ácsorogni.")
	}
L_3de4:
	if g.room != 4 && g.room != 8 && g.room != 6 {
		if g.room != 18 {
			goto L_3e10
		}
	}
	g.OutTextXY(310, 170, "Unalmas a várakozás.")
L_3e10:
	if g.room == 14 {
		g.OutTextXY(310, 170, "A lépcsö nem mozdul. Hiába vársz.")
	}
	if g.room == 15 {
		g.sub1e9b_1476()
	}
	if g.room != 5 {
		goto L_3f50
	}
	g.lookupTargetByName("Okos, az izomzseni")
	if g.j == 1 && g.tgtHP >= 5 {
		g.j = 5
	}
	if g.j == 1 {
		g.sub1e9b_1120()
	}
	g.lookupTargetByName("Kapiczky Arnold")
	if g.j == 2 && g.tgtHP >= 5 {
		g.j = 5
	}
	if g.j == 2 {
		g.sub1e9b_11da()
	}
	g.lookupTargetByName("Brecska, a míveletlen")
	if g.j == 3 && g.tgtHP >= 5 {
		g.j = 5
	}
	if g.j == 3 {
		g.lookupTargetByName("Tordy Erzsébet tanárnö")
		if g.tgtHP <= 5 {
			g.OutTextXY(310, 170, "Brecska hatalmasat ásít. Ezért kb.")
			g.OutTextXY(310, 185, "fél órás lebaszásban részesül.")
		}
		if g.tgtHP <= 5 {
			goto L_3ef6
		}
		g.OutTextXY(310, 170, "Brecska hatalmasat ásít.")
	}
L_3ef6:
	g.lookupTargetByName("Tordy Erzsébet tanárnö")
	if g.j == 4 && g.tgtHP >= 5 {
		g.j = 5
	}
	if g.j == 4 {
		g.OutTextXY(310, 170, "- Na, gyerekek, tehát a házi feladat:")
		g.OutTextXY(310, 185, "138/1a, 138/1b, 139/2...")
	}
	if g.j == 5 {
		g.OutTextXY(310, 170, "A legunalmasabb órák egyike...")
	}
L_3f50:
	if g.room == 12 {
		g.lookupTargetByName("Széchenyi Gábor")
		if g.j == 1 && g.tgtHP >= 5 {
			g.j = 5
		}
		if g.j == 1 {
			g.sub1e9b_0ff7()
		}
		g.lookupTargetByName("Váév, az idióta")
		if g.j == 2 && g.tgtHP >= 5 {
			g.j = 5
		}
		if g.j == 2 {
			g.sub1e9b_0760()
		}
		g.lookupTargetByName("Egon, a \"fogyatékos\"")
		if g.j == 3 && g.tgtHP >= 5 {
			g.j = 5
		}
		if g.j == 3 {
			g.sub1e9b_0ed8()
		}
		g.lookupTargetByName("Balázs, a nagyon vicces")
		if g.j == 4 && g.tgtHP >= 5 {
			g.j = 5
		}
		if g.j == 4 {
			g.OutTextXY(310, 170, "Balázs unalmában krétadarabkákkal")
			g.OutTextXY(310, 185, "hajigálja Egont.")
		}
		if g.j != 5 {
			goto L_404c
		}
		g.OutTextXY(310, 170, "Császár a Paintbrushsal szórakozik,")
		g.OutTextXY(310, 185, "szokásos idióta rajzainak egyikét")
		g.OutTextXY(310, 200, "készíti.")
	}
L_404c:
	if g.room == 16 {
		g.lookupTargetByName("Kutya")
		if g.tgtHP < 5 {
			g.OutTextXY(310, 170, "A blöki morog, és az ínyét húzogatja.")
		}
		if g.tgtHP < 5 {
			goto L_4089
		}
		g.OutTextXY(310, 170, "A kerítésen túl lábdobogást hallasz.")
	}
L_4089:
	if g.room == 21 {
		g.lookupTargetByName("Kutya")
		if g.tgtHP < 5 {
			g.OutTextXY(310, 170, "Maca lehajol a szappanért. Jujujj...")
		}
		if g.tgtHP >= 5 {
			g.OutTextXY(310, 170, "Nyomatékosan várakozol, de nem")
		}
		if g.tgtHP < 5 {
			goto L_40dd
		}
		g.OutTextXY(310, 185, "történik semmi.")
	}
L_40dd:
	if g.room != 7 {
		goto L_41ff
	}
	g.lookupTargetByName("Agócs, a stréber kisfiú")
	if g.j == 1 && g.tgtHP >= 5 {
		g.j = 5
	}
	if g.j == 1 {
		g.lookupTargetByName("Lepedös Mihály")
		if g.tgtHP <= 5 {
			g.sub1e9b_06c6()
		}
		if g.tgtHP <= 5 {
			goto L_412a
		}
		g.sub1e9b_0597()
	}
L_412a:
	g.lookupTargetByName("Zsolt, a nyál")
	if g.j == 2 && g.tgtHP >= 5 {
		g.j = 5
	}
	if g.j == 2 {
		g.sub1e9b_04ec()
	}
	g.lookupTargetByName("Laki, a rocker")
	if g.j == 3 && g.tgtHP >= 5 {
		g.j = 5
	}
	if g.j == 3 {
		g.OutTextXY(310, 170, "Laki elunja a kémiát, elöszed egy")
		g.OutTextXY(310, 185, "regényt, és olvasni kezdi.")
	}
	g.lookupTargetByName("Lepedös Mihály")
	if g.j == 4 && g.tgtHP >= 5 {
		g.j = 5
	}
	if g.j == 4 {
		g.OutTextXY(310, 170, "- ...ez az anyag, kérem, a foszfor,")
		g.OutTextXY(310, 185, "a fehérfoszfor. 42°C - on gyullad...")
		g.OutTextXY(310, 200, "Micu bá jól elvan magában.")
	}
	if g.j == 5 {
		g.OutTextXY(310, 170, "Dögunalom ez az óra!")
	}
L_41ff:
	if g.room != 9 {
		goto L_436a
	}
	g.lookupTargetByName("Ahmed, a szíriai bunkó")
	if g.j == 1 && g.tgtHP >= 5 {
		g.j = 5
	}
	if g.j == 1 {
		g.OutTextXY(310, 170, "Ahmed kihajít a tányérjából egy")
		g.OutTextXY(310, 185, "méretes slejmot, és eszik tovább.")
	}
	g.lookupTargetByName("Suzy, a suli kurvája")
	if g.j == 2 && g.tgtHP >= 5 {
		g.j = 5
	}
	if g.j == 2 {
		g.OutTextXY(310, 170, "Suzy megnézegeti, nem kenödött-e")
		g.OutTextXY(310, 185, "el a rúzsa.")
		if g.tgtState >= 1 {
			goto L_42a6
		}
		g.OutTextXY(403, 185, "Közben csábosan néz rád.")
	}
L_42a6:
	g.lookupTargetByName("Petra, a bamba")
	if g.j == 3 && g.tgtHP >= 5 {
		g.j = 5
	}
	if g.j == 3 && g.tgtHP <= 5 {
		g.OutTextXY(310, 170, "Petra bárgyún bámul, és egy")
		g.OutTextXY(310, 185, "csábosnak szánt mosolyt lövell rád.")
	}
	g.lookupTargetByName("Egy vén fasz tanár")
	if g.j == 4 && g.tgtHP >= 5 {
		g.j = 5
	}
	if g.j == 4 {
		g.OutTextXY(310, 170, "Az öreg tanár fuldoklani kezd, és")
		g.OutTextXY(310, 185, "ellilul, de mielött megütné a guta,")
		g.OutTextXY(310, 200, "mindenféle porokat, meg tablettákat")
		g.OutTextXY(310, 215, "kotor elö, és elfalatozgatja öket.")
	}
	if g.j == 5 {
		g.OutTextXY(310, 170, "Nem történik semmi.")
	}
L_436a:
	if g.room != 11 {
		return
	}
	g.lookupTargetByName("A szakács")
	if g.tgtState == 0 && g.tgtHP < 5 {
		g.OutTextXY(310, 170, "A szakács elhaladtában nagyot csuláz")
		g.OutTextXY(310, 185, "a leveses kondérba.")
	}
	if g.tgtState > 0 && g.tgtHP < 5 {
		g.OutTextXY(310, 170, "A szakács ki akar dobni innen.")
	}
	if g.tgtHP < 5 {
		return
	}
	g.OutTextXY(310, 170, "A spenót csendben odakozmál. A francot")
	g.OutTextXY(310, 185, "se érdekli.")
}

// apologize  (1010:47a7)
func (g *Game) apologize() {
	g.clearTextWin()
	g.resetViewPort()
	g.lookupTargetByName(g.tgtName)
	if g.tgtState >= 3 {
		g.OutTextXY(310, 170, "Nem valószínü, hogy megbocsátana")
		g.OutTextXY(310, 185, "ekkora pofátlanságot.")
	}
	if g.tgtState >= 3 {
		return
	}
	g.j = 1
	for {
		if g.chars[g.j].name == g.tgtName {
			g.chars[g.j].state = g.chars[g.j].state - 1
		}
		if g.chars[g.j].state < 0 {
			g.chars[g.j].state = 0
		}
		g.j = g.j + 1
		if g.j > 30 {
			break
		}
	}
	if g.target == 1 {
		g.OutTextXY(310, 170, "- OK, de akkor se mehetsz be.")
	}
	if g.target == 2 {
		g.OutTextXY(310, 170, "- Dikk, mennyé anyádba.")
	}
	if g.target == 3 {
		g.OutTextXY(310, 170, "- Jó... csak kussoljál.")
	}
	if g.target == 4 {
		g.OutTextXY(310, 170, "Nem igazán érdekli a bocsánatkérésed.")
	}
	if g.target == 5 {
		g.OutTextXY(310, 170, "Megalázkodtál elötte. Most jobb?")
	}
	if g.target == 6 {
		g.OutTextXY(310, 170, "- Jetzt entschuldige ich.")
	}
	if g.target == 7 {
		g.OutTextXY(310, 170, "- Rendben van, de figyelni foglak!")
	}
	if g.target == 8 {
		g.OutTextXY(310, 170, "- Jól van, bazmeg, felejts már el!")
	}
	if g.target == 9 {
		g.OutTextXY(310, 170, "- Na, azért, mert megmondtalak volna!")
	}
	if g.target == 10 {
		g.OutTextXY(310, 170, "Megalázkodtál elötte. Most jobb?")
	}
	if g.target == 11 {
		g.OutTextXY(310, 170, "- Kurva anyád.")
	}
	if g.target == 12 {
		g.OutTextXY(310, 170, "- De aztán viselkedjél, fiam, mert")
	}
	if g.target == 12 {
		g.OutTextXY(310, 185, "  hátbaváglak!")
	}
	if g.target == 13 {
		g.OutTextXY(310, 170, "Nem igazán érdekli a bocsánatkérésed.")
	}
	if g.target == 14 {
		g.OutTextXY(310, 170, "Legyint, és kirúzsozza magát.")
	}
	if g.target == 15 {
		g.OutTextXY(310, 170, "- Tessék? Jó, nem haragszom...")
	}
	if g.target == 16 {
		g.OutTextXY(310, 170, "- Rendben van, fiatalember, de aztán")
		g.OutTextXY(310, 185, "  viselkedjen rendesen!")
	}
	if g.target == 17 {
		g.OutTextXY(310, 170, "- Jól van, csak hagyjál dolgozni!")
	}
	if g.target == 18 {
		g.OutTextXY(310, 170, "- Na... azért!")
	}
	if g.target == 19 {
		g.OutTextXY(310, 170, "- Nincs megbocsátás! Ezt én")
	}
	if g.target == 19 {
		g.OutTextXY(310, 185, "  parancsolom!")
	}
	if g.target == 20 {
		g.OutTextXY(310, 170, "- Jól van már, menj a picsába.")
	}
	if g.target == 21 {
		g.OutTextXY(310, 170, "- Hát... na, jó...")
	}
	if g.target == 22 {
		g.OutTextXY(310, 170, "- Bocsánat, hihihi!")
	}
	if g.target == 23 {
		g.OutTextXY(310, 170, "- Pitizz, te rohadék!")
	}
	if g.target == 24 {
		g.OutTextXY(310, 170, "- Viselkedjél fiam, mert kiváglak!")
	}
	if g.target == 25 {
		g.OutTextXY(310, 170, "Semmi hatás.")
	}
	if g.target == 26 {
		g.OutTextXY(310, 170, "Legyint, és zuhanyozik tovább.")
	}
	if g.target == 27 {
		g.OutTextXY(310, 170, "- Menj anyádhoz!")
	}
	if g.target == 28 {
		g.OutTextXY(310, 170, "- Kapd be.")
	}
	if g.target == 29 {
		g.OutTextXY(310, 170, "- Legközelebb beírok az ellenörzödbe!")
	}
	if g.target == 29 {
		g.OutTextXY(310, 185, "  És most menj ki innen!")
	}
	if g.target == 30 {
		g.OutTextXY(310, 170, "- Azért mondom, aranyom, mert én")
	}
	if g.target == 30 {
		g.OutTextXY(310, 185, "  letöröm a derekadat, ha pimasz")
	}
	if g.target != 30 {
		return
	}
	g.OutTextXY(310, 200, "  vagy!")
}

// nothingFound  (1010:4b96)
func (g *Game) nothingFound() {
	g.OutTextXY(310, 170, "Nem találsz semmi érdekeset.")
}

// checkPickup  (1010:4c21)
func (g *Game) checkPickup() {
	g.loaded = 0
	for g.i = 1; ; g.i++ {
		if g.inv[g.i] == "" {
			g.loaded = 1
		}
		if g.i == 12 {
			break
		}
	}
	if g.loaded == 0 {
		g.OutTextXY(310, 170, "Semmit nem vehetsz már fel.")
	}
	if g.loaded == 0 {
		return
	}
	g.i = 0
	for g.n = 1; ; g.n++ {
		if g.roomItems[g.n] != "" {
			g.i = g.i + 1
		}
		if g.n == 5 {
			break
		}
	}
	if g.i > 1 {
		g.m2 = 0
		for g.n = 1; ; g.n++ {
			if g.inv[g.n] == "" {
				g.m2 = g.m2 + 1
			}
			if g.n == 12 {
				break
			}
		}
		if g.i <= g.m2 {
			goto L_4d09
		}
		g.loaded = 0
		g.OutTextXY(310, 170, "Több tárgy is van itt, amit")
		g.OutTextXY(310, 185, "felszedhetnél, de sajnos nincs")
		g.OutTextXY(310, 200, "elég hely a zsebeidben.")
	}
L_4d09:
	if g.loaded == 0 {
		return
	}
	for g.i = 1; ; g.i++ {
		if g.roomItems[g.i] == "" {
			g.roomItems[g.i] = trunc("Lófasz", 25)
		}
		if g.i == 5 {
			break
		}
	}
	for g.i = 1; ; g.i++ {
		for g.n = 1; ; g.n++ {
			if g.inv[g.n] == g.roomItems[g.i] {
				g.loaded = 0
			}
			if g.n == 12 {
				break
			}
		}
		if g.i == 5 {
			break
		}
	}
	for g.i = 1; ; g.i++ {
		for g.n = 1; ; g.n++ {
			if g.taken[g.n] == g.roomItems[g.i] {
				g.loaded = 0
			}
			if g.n == 120 {
				break
			}
		}
		if g.i == 5 {
			break
		}
	}
	for g.i = 1; ; g.i++ {
		if g.roomItems[g.i] == "Lófasz" {
			g.roomItems[g.i] = ""
		}
		if g.i == 5 {
			break
		}
	}
	if g.loaded != 0 {
		return
	}
	g.nothingFound()
}

// addItem  (1010:4e19)
func (g *Game) addItem(a4 string) {
	var l100 string
	l100 = a4
	g.i = 0
	for {
		g.i = g.i + 1
		if g.inv[g.i] == "" {
			break
		}
	}
	g.inv[g.i] = trunc(l100, 25)
}

// search  (1010:553f)
func (g *Game) search() {
	g.clearTextWin()
	g.resetViewPort()
	g.i = 0
	for {
		g.i = g.i + 1
		if g.i > 5 {
			goto L_556b
		}
		if g.roomItems[g.i] != "" {
			break
		}
	}
L_556b:
	if g.i > 5 {
		g.loaded = 0
		if g.room == 2 {
			g.loaded = 1
		}
		if g.room == 20 {
			g.loaded = 1
		}
		if g.room == 18 {
			g.loaded = 1
		}
		if g.room == 19 {
			g.loaded = 1
		}
		if g.room == 13 {
			g.loaded = 1
		}
		if g.room == 12 {
			g.loaded = 1
		}
		if g.room == 11 {
			g.loaded = 1
		}
		if g.room == 21 {
			g.loaded = 1
		}
		if g.room == 5 {
			g.loaded = 1
		}
		if g.room == 7 {
			g.loaded = 1
		}
		if g.room == 6 {
			g.loaded = 1
		}
		if g.room != 22 {
			goto L_560a
		}
		g.loaded = 1
	}
L_560a:
	if g.loaded == 0 {
		g.nothingFound()
	}
	if g.loaded == 0 {
		return
	}
	if g.room == 2 {
		g.checkPickup()
		if g.loaded != 1 {
			goto L_5647
		}
		g.OutTextXY(310, 170, "Egy üres pálinkásüveget találsz.")
		g.addItem("Pálinkásüveg")
	}
L_5647:
	if g.room == 30 {
		g.checkPickup()
		if g.loaded != 1 {
			goto L_5696
		}
		g.OutTextXY(310, 170, "Wow, yeah! Egy motoros fürész!!")
		g.OutTextXY(310, 185, "Kár, hogy alig van benne benzin, így")
		g.OutTextXY(310, 200, "csak néhányszor használhatod.")
		g.addItem("Motorfürész")
		g.v0866 = 5
	}
L_5696:
	if g.room == 28 {
		g.checkPickup()
		if g.loaded != 1 {
			goto L_56bf
		}
		g.OutTextXY(310, 170, "A polcon egy gázálarcot találsz.")
		g.addItem("Gázálarc")
	}
L_56bf:
	if g.room == 23 {
		g.checkPickup()
		if g.loaded != 1 {
			goto L_56e8
		}
		g.OutTextXY(310, 170, "Rábukkansz a pince kulcsára.")
		g.addItem("Pincekulcs")
	}
L_56e8:
	if g.room == 20 {
		g.checkPickup()
		if g.loaded != 1 {
			goto L_5721
		}
		g.OutTextXY(310, 170, "Ahá! Mi van itt! Diadalmasan emelsz")
		g.OutTextXY(310, 185, "a magasba egy baseball-ütöt. Yes!")
		g.addItem("Baseball-ütö")
	}
L_5721:
	if g.room == 18 {
		g.checkPickup()
		if g.loaded != 1 {
			goto L_575a
		}
		g.OutTextXY(310, 170, "A földön egy ócska, olajos rongyra")
		g.OutTextXY(310, 185, "bukkansz. Felveszed és elrakod.")
		g.addItem("Rongy")
	}
L_575a:
	if g.room != 19 {
		goto L_5940
	}
	g.OutTextXY(310, 170, "A sarokban egy marmonkannában találsz")
	g.loaded = 1
	g.i = 0
	for {
		g.i = g.i + 1
		if g.i > 12 {
			goto L_57a3
		}
		if g.inv[g.i] == "Pálinkásüveg" {
			break
		}
	}
L_57a3:
	if g.i > 12 {
		g.loaded = 0
	}
	g.i = 0
	for {
		g.i = g.i + 1
		if g.i > 12 {
			goto L_57d9
		}
		if g.inv[g.i] == "Rongy" {
			break
		}
	}
L_57d9:
	if g.i > 12 {
		g.loaded = 0
	}
	g.i = 0
	for {
		g.i = g.i + 1
		if g.i > 12 {
			goto L_580f
		}
		if g.inv[g.i] == "Tölcsér" {
			break
		}
	}
L_580f:
	if g.i > 12 {
		g.loaded = 0
	}
	if g.loaded == 0 {
		g.OutTextXY(310, 185, "egy kis szuperbenzint. Sajnos nem")
		g.OutTextXY(310, 200, "tudsz vele mit kezdeni.")
	}
	if g.loaded == 1 {
		g.OutTextXY(310, 185, "egy kis szuperbenzint. Támad egy")
		g.OutTextXY(310, 200, "ötleted. Fogod a portásfülkében")
		g.OutTextXY(310, 215, "talált pálinkásüveget, és Micu bá")
		g.OutTextXY(310, 230, "tölcsére segítségével megtöltöd")
		g.OutTextXY(310, 245, "benzinnel. Majd bedugaszolod a")
		g.OutTextXY(310, 260, "palackot egy ronggyal: voilá, egy")
		g.OutTextXY(310, 275, "Molotov-koktél!")
		for g.i = 1; ; g.i++ {
			if g.inv[g.i] == "Pálinkásüveg" {
				g.inv[g.i] = ""
			}
			if g.inv[g.i] == "Rongy" {
				g.inv[g.i] = ""
			}
			if g.i == 12 {
				break
			}
		}
		g.i = 0
		for {
			g.i = g.i + 1
			if g.inv[g.i] == "" {
				break
			}
		}
		g.inv[g.i] = trunc("Molotov-koktél", 25)
	}
L_5940:
	if g.room == 13 {
		g.checkPickup()
		if g.loaded != 1 {
			goto L_5979
		}
		g.OutTextXY(310, 170, "A fal mellett egy döglött csótányra")
		g.OutTextXY(310, 185, "bukkansz. Ki tudja, még hasznos lehet.")
		g.addItem("Döglött csótány")
	}
L_5979:
	if g.room == 12 {
		g.checkPickup()
		if g.loaded != 1 {
			goto L_59a2
		}
		g.OutTextXY(310, 170, "Elemelsz egy floppylemezt.")
		g.addItem("Floppylemez")
	}
L_59a2:
	if g.room == 11 {
		g.checkPickup()
		if g.loaded != 1 {
			goto L_5a0b
		}
		g.OutTextXY(310, 170, "Mindent megvizsgálsz a konyhában,")
		g.OutTextXY(310, 185, "de az egyetlen, amit elvitelre")
		g.OutTextXY(310, 200, "érdemesnek találsz, egy csont a")
		g.OutTextXY(310, 215, "moslékos vödörben. Legyözöd")
		g.OutTextXY(310, 230, "undorodat, és kiemeled.")
		g.addItem("Csont")
	}
L_5a0b:
	if g.room == 21 {
		g.checkPickup()
		if g.loaded != 1 {
			goto L_5a44
		}
		g.OutTextXY(310, 170, "Elemeled Maca törülközöjét. Ki tudja,")
		g.OutTextXY(310, 185, "még jól jöhet.")
		g.addItem("Törülközö")
	}
L_5a44:
	if g.room == 5 {
		g.loaded = 1
		for g.i = 1; ; g.i++ {
			if g.inv[g.i] == "Golyóstoll" {
				g.loaded = 0
			}
			if g.i == 12 {
				break
			}
		}
		for g.i = 1; ; g.i++ {
			if g.taken[g.i] == "Golyóstoll" {
				g.loaded = 0
			}
			if g.i == 120 {
				break
			}
		}
		if g.loaded == 0 {
			g.nothingFound()
		}
		if g.loaded != 1 {
			goto L_5b38
		}
		g.lookupTargetByName("Tordy Erzsébet tanárnö")
		if g.tgtHP < 5 {
			g.OutTextXY(310, 170, "Amikor a tanárnö elfordul egy")
			g.OutTextXY(310, 185, "pillanatra, elemeled a tollát")
			g.OutTextXY(310, 200, "az asztalról.")
			g.addItem("Golyóstoll")
		}
		if g.tgtHP < 5 {
			goto L_5b38
		}
		g.OutTextXY(310, 170, "Elviszed a kifingott tanárnö")
		g.OutTextXY(310, 185, "golyóstollát az asztalról.")
		g.addItem("Golyóstoll")
	}
L_5b38:
	if g.room == 7 {
		g.loaded = 1
		for g.i = 1; ; g.i++ {
			if g.inv[g.i] == "Tölcsér" {
				g.loaded = 0
			}
			if g.i == 12 {
				break
			}
		}
		for g.i = 1; ; g.i++ {
			if g.taken[g.i] == "Tölcsér" {
				g.loaded = 0
			}
			if g.i == 120 {
				break
			}
		}
		if g.loaded == 0 {
			g.nothingFound()
		}
		if g.loaded != 1 {
			goto L_5c2c
		}
		g.lookupTargetByName("Lepedös Mihály")
		if g.tgtHP < 5 {
			g.OutTextXY(310, 170, "Micu bá elfordul, hogy valami")
			g.OutTextXY(310, 185, "baromságot írjon a táblára, te addig")
			g.OutTextXY(310, 200, "is hasznossá teszed magad, és elemelsz")
			g.OutTextXY(310, 215, "az asztalról egy müanyag tölcsért.")
			g.addItem("Tölcsér")
		}
		if g.tgtHP < 5 {
			goto L_5c2c
		}
		g.OutTextXY(310, 170, "Találsz egy tölcsért az asztalon.")
		g.addItem("Tölcsér")
	}
L_5c2c:
	if g.room != 6 {
		goto L_5d84
	}
	g.lookupTargetByName("A pedellus")
	if g.tgtHP < 5 {
		g.nothingFound()
	}
	if g.tgtHP < 5 {
		goto L_5d84
	}
	g.loaded = 0
	g.i = 0
	for {
		g.i = g.i + 1
		if g.i > 120 {
			goto L_5c81
		}
		if g.taken[g.i] == "Raktárkulcs" {
			break
		}
	}
L_5c81:
	if g.taken[g.i] == "Raktárkulcs" {
		g.nothingFound()
	}
	if g.i >= 120 {
		g.loaded = 1
	}
	if g.loaded == 0 {
		return
	}
	g.n = 0
	for {
		g.n = g.n + 1
		if g.n > 12 {
			goto L_5cdb
		}
		if g.inv[g.n] == "Raktárkulcs" {
			break
		}
	}
L_5cdb:
	if g.inv[g.n] == "Raktárkulcs" {
		g.nothingFound()
	}
	if g.n <= 12 {
		goto L_5d84
	}
	g.n = 0
	for {
		g.n = g.n + 1
		if g.n > 12 {
			goto L_5d1e
		}
		if g.inv[g.n] == "" {
			break
		}
	}
L_5d1e:
	if g.n > 12 {
		g.OutTextXY(310, 170, "Már tele vannak a zsebeid.")
	}
	if g.n <= 12 {
		g.OutTextXY(310, 170, "A kinyuvasztott pedellus zsebében egy")
		g.OutTextXY(310, 185, "kulcsot találsz. A kulcstartó felirata")
		g.OutTextXY(310, 200, "szerint ez az iskolaudvarról nyíló")
		g.OutTextXY(310, 215, "raktárhelyiséget nyitja.")
		g.addItem("Raktárkulcs")
	}
L_5d84:
	if g.room != 22 {
		return
	}
	g.lookupTargetByName("Harami, a melák")
	if g.tgtHP < 5 {
		g.nothingFound()
	}
	if g.tgtHP < 5 {
		return
	}
	g.loaded = 0
	g.i = 0
	for {
		g.i = g.i + 1
		if g.i > 120 {
			goto L_5dd9
		}
		if g.taken[g.i] == "Öngyújtó" {
			break
		}
	}
L_5dd9:
	if g.taken[g.i] == "Öngyújtó" {
		g.nothingFound()
	}
	if g.i >= 120 {
		g.loaded = 1
	}
	if g.loaded == 0 {
		return
	}
	g.n = 0
	for {
		g.n = g.n + 1
		if g.n > 12 {
			goto L_5e33
		}
		if g.inv[g.n] == "Öngyújtó" {
			break
		}
	}
L_5e33:
	if g.inv[g.n] == "Öngyújtó" {
		g.nothingFound()
	}
	if g.n <= 12 {
		return
	}
	g.n = 0
	for {
		g.n = g.n + 1
		if g.n > 12 {
			goto L_5e76
		}
		if g.inv[g.n] == "" {
			break
		}
	}
L_5e76:
	if g.n > 12 {
		g.OutTextXY(310, 170, "Már tele vannak a zsebeid.")
	}
	if g.n > 12 {
		return
	}
	g.OutTextXY(310, 170, "Végigkutatod a két izomagy zsebeit.")
	g.OutTextXY(310, 185, "Sok taknyos papírzsebkendöt, használt")
	g.OutTextXY(310, 200, "buszjegyeket találsz, meg egy")
	g.OutTextXY(310, 215, "öngyújtót.")
	g.addItem("Öngyújtó")
}

// otherMenu  (1010:5ede)
func (g *Game) otherMenu() {
	g.clearMenuWin()
	g.resetViewPort()
	g.i = 1
L_5ef1:
	for g.n = 1; ; g.n++ {
		g.SetColor(14)
		g.j = 0
		for g.k = 1; ; g.k++ {
			if g.roomPeople[g.k] != "" {
				g.j = 1
			}
			if g.k == 10 {
				break
			}
		}
		if g.n == 1 && g.j == 0 {
			g.SetColor(7)
		}
		if g.n == 2 && g.target == 0 {
			g.SetColor(7)
		}
		if g.n == 4 && g.target == 0 {
			g.SetColor(7)
		}
		if g.n == 6 && g.target == 0 {
			g.SetColor(7)
		}
		g.OutTextXY(25, g.n*15+150, otherMenu[g.n])
		if g.n == 8 {
			break
		}
	}
	g.SetFillStyle(1, 21)
	g.Bar(22, g.i*15+150-1-1, 248, g.i*15+150+10)
	g.SetColor(14)
	g.j = 0
	for g.k = 1; ; g.k++ {
		if g.roomPeople[g.k] != "" {
			g.j = 1
		}
		if g.k == 10 {
			break
		}
	}
	if g.i == 1 && g.j == 0 {
		g.SetColor(7)
	}
	if g.i == 2 && g.target == 0 {
		g.SetColor(7)
	}
	if g.i == 4 && g.target == 0 {
		g.SetColor(7)
	}
	if g.i == 6 && g.target == 0 {
		g.SetColor(7)
	}
	g.OutTextXY(25, g.i*15+150, otherMenu[g.i])
	g.FlushKeys()
	for {
		if g.KeyPressed() {
			break
		}
	}
	g.SetFillStyle(1, 0)
	g.FloodFill(22, g.i*15+150, 0)
	g.key = g.ReadKey()
	if g.key == 0 {
		g.key = g.ReadKey()
		if g.key == 72 {
			g.i = g.i - 1
		}
		if g.key != 80 {
			goto L_60cd
		}
		g.i = g.i + 1
	}
L_60cd:
	if g.i < 1 {
		g.i = 8
	}
	if g.i > 8 {
		g.i = 1
	}
	if g.key == 13 {
		g.j = 0
		for g.k = 1; ; g.k++ {
			if g.roomPeople[g.k] != "" {
				g.j = 1
			}
			if g.k == 10 {
				break
			}
		}
		if g.i == 1 && g.j == 0 {
			g.key = 0
		}
		if g.i == 2 && g.target == 0 {
			g.key = 0
		}
		if g.i == 4 && g.target == 0 {
			g.key = 0
		}
		if g.i != 6 {
			goto L_6164
		}
		if g.target != 0 {
			goto L_6164
		}
		g.key = 0
	}
L_6164:
	if g.key != 13 && g.key != 27 {
		goto L_5ef1
	}
	if g.key == 27 {
		return
	}
	if g.i == 1 {
		g.selectTarget()
		g.i = 0
	}
	if g.i == 2 {
		g.spit()
		g.i = 0
	}
	if g.i == 3 {
		g.burp()
		g.i = 0
	}
	if g.i == 4 {
		g.rape()
		g.i = 0
	}
	if g.i == 5 {
		g.waitAround()
		g.i = 0
	}
	if g.i == 6 {
		g.apologize()
		g.i = 0
	}
	if g.i == 7 {
		g.search()
		g.i = 0
	}
	if g.i != 8 {
		return
	}
	g.quitConfirm()
	g.i = 0
}

// attackMenu  (1010:621d)
func (g *Game) attackMenu() {
	g.clearMenuWin()
	g.resetViewPort()
	g.i = 1
L_6231:
	for g.n = 1; ; g.n++ {
		g.SetColor(14)
		g.j = 0
		for g.m2 = 1; ; g.m2++ {
			if g.inv[g.m2] == "Baseball-ütö" {
				g.j = 1
			}
			if g.m2 == 12 {
				break
			}
		}
		if g.n == 10 && g.j == 0 {
			g.SetColor(7)
		}
		g.OutTextXY(25, g.n*15+150, attackMenu[g.n])
		if g.n == 10 {
			break
		}
	}
	g.SetFillStyle(1, 21)
	g.Bar(22, g.i*15+150-1-1, 248, g.i*15+150+10)
	g.SetColor(14)
	g.j = 0
	for g.m2 = 1; ; g.m2++ {
		if g.inv[g.m2] == "Baseball-ütö" {
			g.j = 1
		}
		if g.m2 == 12 {
			break
		}
	}
	if g.i == 10 && g.j == 0 {
		g.SetColor(7)
	}
	g.OutTextXY(25, g.i*15+150, attackMenu[g.i])
	g.FlushKeys()
	for {
		if g.KeyPressed() {
			break
		}
	}
	g.SetFillStyle(1, 0)
	g.FloodFill(22, g.i*15+150, 0)
	g.key = g.ReadKey()
	if g.key == 0 {
		g.key = g.ReadKey()
		if g.key == 72 {
			g.i = g.i - 1
		}
		if g.key != 80 {
			goto L_63a2
		}
		g.i = g.i + 1
	}
L_63a2:
	if g.i < 1 {
		g.i = 10
	}
	if g.i > 10 {
		g.i = 1
	}
	if g.key == 13 {
		g.j = 0
		for g.n = 1; ; g.n++ {
			if g.inv[g.n] == "Baseball-ütö" {
				g.j = 1
			}
			if g.n == 12 {
				break
			}
		}
		if g.i != 10 {
			goto L_640b
		}
		if g.j != 0 {
			goto L_640b
		}
		g.key = 0
	}
L_640b:
	if g.key != 13 && g.key != 27 {
		goto L_6231
	}
	if g.key == 27 {
		return
	}
	g.line = itoa(g.target)
	if g.target < 10 {
		g.line = "0" + g.line
	}
	g.txt = g.AssignText("data\\trg" + g.line + ".dat")
	g.ResetText(g.txt)
	if g.IOResult() != 0 {
		g.DiskError()
	}
	g.line = itoa(g.i)
	g.line = "*attack" + g.line
	for {
		g.line2 = g.txt.ReadLn(255)
		if g.line == g.line2 {
			break
		}
	}
	g.showTextFile()
	g.line2 = g.txt.ReadLn(255)
	g.n = val(g.line2, &g.m2)
	g.m2 = 1
	for {
		if g.chars[g.m2].name == g.tgtName {
			g.chars[g.m2].hp = g.chars[g.m2].hp + g.n
		}
		g.m2 = g.m2 + 1
		if g.m2 > 30 {
			break
		}
	}
	g.line2 = g.txt.ReadLn(255)
	g.n = val(g.line2, &g.m2)
	g.m2 = 1
	for {
		if g.chars[g.m2].name == g.tgtName {
			g.chars[g.m2].state = g.chars[g.m2].state + g.n
		}
		g.m2 = g.m2 + 1
		if g.m2 > 30 {
			break
		}
	}
	g.line2 = g.txt.ReadLn(255)
	g.n = val(g.line2, &g.m2)
	g.energy = g.energy - g.n
	g.CloseText(g.txt)
}

// chat  (1010:668f)
func (g *Game) chat() {
	g.clearTextWin()
	g.resetViewPort()
	g.lookupTargetByName("Cigány")
	if g.room == 3 && g.tgtState >= 1 {
		g.OutTextXY(310, 170, "- Helló...")
		g.OutTextXY(310, 185, "- Takaroddzsá, me sétmetsem a sádat!")
	}
	if g.room == 3 && g.tgtState >= 1 {
		return
	}
	g.line = itoa(g.target)
	if g.target < 10 {
		g.line = "0" + g.line
	}
	g.txt = g.AssignText("data\\trg" + g.line + ".dat")
	g.ResetText(g.txt)
	if g.IOResult() != 0 {
		g.DiskError()
	}
	g.line = "*talk"
	for {
		g.line2 = g.txt.ReadLn(255)
		if g.line == g.line2 {
			break
		}
	}
	g.showTextFile()
	g.line2 = g.txt.ReadLn(255)
	g.n = val(g.line2, &g.m2)
	g.m2 = 1
	for {
		if g.chars[g.m2].name == g.tgtName {
			g.chars[g.m2].hp = g.chars[g.m2].hp + g.n
		}
		g.m2 = g.m2 + 1
		if g.m2 > 30 {
			break
		}
	}
	g.line2 = g.txt.ReadLn(255)
	g.n = val(g.line2, &g.m2)
	g.m2 = 1
	for {
		if g.chars[g.m2].name == g.tgtName {
			g.chars[g.m2].state = g.chars[g.m2].state + g.n
		}
		g.m2 = g.m2 + 1
		if g.m2 > 30 {
			break
		}
	}
	g.line2 = g.txt.ReadLn(255)
	g.n = val(g.line2, &g.m2)
	g.energy = g.energy - g.n
	g.CloseText(g.txt)
}

// askMoney  (1010:68f0)
func (g *Game) askMoney() {
	g.clearMenuWin()
	g.resetViewPort()
	g.i = 1
L_6904:
	for g.n = 1; ; g.n++ {
		g.SetColor(14)
		g.OutTextXY(25, g.n*15+150, begLines[g.n])
		if g.n == 8 {
			break
		}
	}
	g.SetFillStyle(1, 21)
	g.Bar(22, g.i*15+150-1-1, 248, g.i*15+150+10)
	g.SetColor(14)
	g.OutTextXY(25, g.i*15+150, begLines[g.i])
	g.FlushKeys()
	for {
		if g.KeyPressed() {
			break
		}
	}
	g.SetFillStyle(1, 0)
	g.FloodFill(22, g.i*15+150, 0)
	g.key = g.ReadKey()
	if g.key == 0 {
		g.key = g.ReadKey()
		if g.key == 72 {
			g.i = g.i - 1
		}
		if g.key != 80 {
			goto L_69e1
		}
		g.i = g.i + 1
	}
L_69e1:
	if g.i < 1 {
		g.i = 8
	}
	if g.i > 8 {
		g.i = 1
	}
	if g.key != 13 && g.key != 27 {
		goto L_6904
	}
	if g.key == 27 {
		return
	}
	g.line = itoa(g.target)
	if g.target < 10 {
		g.line = "0" + g.line
	}
	g.txt = g.AssignText("data\\trg" + g.line + ".dat")
	g.ResetText(g.txt)
	if g.IOResult() != 0 {
		g.DiskError()
	}
	g.line = itoa(g.i)
	g.line = "*askmoney" + g.line
	for {
		g.line2 = g.txt.ReadLn(255)
		if g.line == g.line2 {
			break
		}
	}
	g.showTextFile()
	g.line2 = g.txt.ReadLn(255)
	g.n = val(g.line2, &g.m2)
	g.m2 = 1
	for {
		if g.chars[g.m2].name == g.tgtName {
			g.chars[g.m2].hp = g.chars[g.m2].hp + g.n
		}
		g.m2 = g.m2 + 1
		if g.m2 > 30 {
			break
		}
	}
	g.line2 = g.txt.ReadLn(255)
	g.n = val(g.line2, &g.m2)
	g.m2 = 1
	for {
		if g.chars[g.m2].name == g.tgtName {
			g.chars[g.m2].state = g.chars[g.m2].state + g.n
		}
		g.m2 = g.m2 + 1
		if g.m2 > 30 {
			break
		}
	}
	g.line2 = g.txt.ReadLn(255)
	g.n = val(g.line2, &g.m2)
	g.energy = g.energy - g.n
	g.line = g.txt.ReadLn(255)
	if charAt(g.line, 1) != 42 {
		g.n = val(g.line, &g.m2)
	}
	g.CloseText(g.txt)
}

// drawSwearMenu  (1010:6ca4)
func (g *Game) drawSwearMenu() {
	g.clearMenuWin()
	g.resetViewPort()
	g.SetColor(7)
	g.OutTextXY(25, 335, "PgUp: még káromkodás.")
	g.SetColor(7)
	g.OutTextXY(25, 350, "Space: kiíratás.")
	g.i = 12
L_6ce3:
	for g.n = 12; ; g.n++ {
		g.SetColor(14)
		g.OutTextXY(25, (g.n-11)*15+150, swearWords[g.n])
		if g.n == 22 {
			break
		}
	}
	g.SetFillStyle(1, 21)
	g.Bar(22, (g.i-11)*15+150-1-1, 248, (g.i-11)*15+150+10)
	g.SetColor(14)
	g.OutTextXY(25, (g.i-11)*15+150, swearWords[g.i])
	g.FlushKeys()
	for {
		if g.KeyPressed() {
			break
		}
	}
	g.SetFillStyle(1, 0)
	g.FloodFill(22, (g.i-11)*15+150, 0)
	g.key = g.ReadKey()
	if g.key == 0 {
		g.key = g.ReadKey()
		if g.key == 72 {
			g.i = g.i - 1
		}
		if g.key == 80 {
			g.i = g.i + 1
		}
		if g.key == 73 {
			g.clearMenuWin()
			g.resetViewPort()
			g.SetColor(7)
			g.OutTextXY(25, 335, "PgDn: még káromkodás.")
			g.SetColor(7)
			g.OutTextXY(25, 350, "Space: kiíratás.")
			for g.n = 1; ; g.n++ {
				g.SetColor(14)
				g.OutTextXY(25, g.n*15+150, swearWords[g.n])
				if g.n == 11 {
					break
				}
			}
		}
		if g.key != 73 {
			goto L_6e52
		}
		return
	}
L_6e52:
	if g.i < 12 {
		g.i = 22
	}
	if g.i > 22 {
		g.i = 12
	}
	if g.key == 13 && g.j < 150 {
		g.curse[g.j] = trunc(swearWords[g.i], 20)
		g.j = g.j + 1
	}
	if g.key == 32 {
		return
	}
	if g.key == 27 {
		return
	}
	goto L_6ce3
}

// swear  (1010:6ed9)
func (g *Game) swear() {
	g.j = 1
	for g.i = 1; ; g.i++ {
		g.curse[g.i] = ""
		if g.i == 150 {
			break
		}
	}
	g.clearMenuWin()
	g.resetViewPort()
	g.SetColor(7)
	g.OutTextXY(25, 335, "PgDn: még káromkodás.")
	g.SetColor(7)
	g.OutTextXY(25, 350, "Space: kiíratás.")
	g.j = 1
	g.i = 1
L_6f42:
	for g.n = 1; ; g.n++ {
		g.SetColor(14)
		g.OutTextXY(25, g.n*15+150, swearWords[g.n])
		if g.n == 11 {
			break
		}
	}
	g.SetFillStyle(1, 21)
	g.Bar(22, g.i*15+150-1-1, 248, g.i*15+150+10)
	g.SetColor(14)
	g.OutTextXY(25, g.i*15+150, swearWords[g.i])
	g.FlushKeys()
	for {
		if g.KeyPressed() {
			break
		}
	}
	g.SetFillStyle(1, 0)
	g.FloodFill(22, g.i*15+150, 0)
	g.key = g.ReadKey()
	if g.key == 0 {
		g.key = g.ReadKey()
		if g.key == 72 {
			g.i = g.i - 1
		}
		if g.key == 80 {
			g.i = g.i + 1
		}
		if g.key != 81 {
			goto L_7029
		}
		g.drawSwearMenu()
	}
L_7029:
	if g.i < 1 {
		g.i = 11
	}
	if g.i > 11 {
		g.i = 1
	}
	if g.key == 13 && g.j < 150 {
		g.curse[g.j] = trunc(swearWords[g.i], 20)
		g.j = g.j + 1
	}
	if g.key == 32 {
		return
	}
	if g.key == 27 {
		return
	}
	goto L_6f42
}

// saySomething  (1010:70c8)
func (g *Game) saySomething() {
	g.clearTextWin()
	g.resetViewPort()
	g.j = 1
	g.line = ""
	g.i = 1
	goto L_70ed
L_70e9:
	g.i++
L_70ed:
	if g.curse[g.i] != "" {
		g.line = g.line + g.curse[g.i]
	}
	if g.curse[g.i] == "anyád" && g.curse[g.i+1] != "baszd meg" && g.curse[g.i+1] != "" {
		g.line = g.line + ","
	}
	if g.curse[g.i] == "baszd meg" && g.curse[g.i+1] != "anyád" && g.curse[g.i+1] != "a" && g.curse[g.i+1] != "" {
		g.line = g.line + ","
	}
	if g.curse[g.i] == "szopj le" && g.curse[g.i+1] != "" {
		g.line = g.line + ","
	}
	if g.curse[g.i+1] == "" && g.curse[g.i] != "" {
		g.line = g.line + "!"
	}
	g.line = g.line + " "
	if g.i == 1 {
		g.line = setChar(g.line, 1, upCase(charAt(g.line, 1)))
	}
	if g.TextWidth(g.line+g.curse[g.i+1]) < 300 {
		if g.curse[g.i+1] != "" {
			goto L_732c
		}
	}
	g.OutTextXY(310, g.j*15+155, g.line)
	g.line = ""
	g.j = g.j + 1
L_732c:
	if g.i != 150 {
		goto L_70e9
	}
	g.FlushKeys()
	g.WaitKey()
	g.line = itoa(g.target)
	if g.target < 10 {
		g.line = "0" + g.line
	}
	g.txt = g.AssignText("data\\trg" + g.line + ".dat")
	g.ResetText(g.txt)
	if g.IOResult() != 0 {
		g.DiskError()
	}
	g.line = itoa(g.i)
	g.line = "*saysomething"
	for {
		g.line2 = g.txt.ReadLn(255)
		if g.line == g.line2 {
			break
		}
	}
	g.showTextFile()
	g.line2 = g.txt.ReadLn(255)
	g.n = val(g.line2, &g.m2)
	g.m2 = 1
	for {
		if g.chars[g.m2].name == g.tgtName {
			g.chars[g.m2].hp = g.chars[g.m2].hp + g.n
		}
		g.m2 = g.m2 + 1
		if g.m2 > 30 {
			break
		}
	}
	g.line2 = g.txt.ReadLn(255)
	g.n = val(g.line2, &g.m2)
	g.m2 = 1
	for {
		if g.chars[g.m2].name == g.tgtName {
			g.chars[g.m2].state = g.chars[g.m2].state + g.n
		}
		g.m2 = g.m2 + 1
		if g.m2 > 30 {
			break
		}
	}
	g.line2 = g.txt.ReadLn(255)
	g.n = val(g.line2, &g.m2)
	g.energy = g.energy - g.n
	g.CloseText(g.txt)
}

// talkMenu  (1010:7548)
func (g *Game) talkMenu() {
	g.clearMenuWin()
	g.resetViewPort()
	g.i = 1
L_755b:
	for g.n = 1; ; g.n++ {
		g.SetColor(14)
		g.OutTextXY(25, g.n*15+150, talkMenu[g.n])
		if g.n == 3 {
			break
		}
	}
	g.SetFillStyle(1, 21)
	g.Bar(22, g.i*15+150-1-1, 248, g.i*15+150+10)
	g.SetColor(14)
	g.OutTextXY(25, g.i*15+150, talkMenu[g.i])
	g.FlushKeys()
	for {
		if g.KeyPressed() {
			break
		}
	}
	g.SetFillStyle(1, 0)
	g.FloodFill(22, g.i*15+150, 0)
	g.key = g.ReadKey()
	if g.key == 0 {
		g.key = g.ReadKey()
		if g.key == 72 {
			g.i = g.i - 1
		}
		if g.key != 80 {
			goto L_7638
		}
		g.i = g.i + 1
	}
L_7638:
	if g.i < 1 {
		g.i = 3
	}
	if g.i > 3 {
		g.i = 1
	}
	if g.key != 13 && g.key != 27 {
		goto L_755b
	}
	if g.key == 27 {
		return
	}
	if g.i == 1 {
		g.askMoney()
	}
	if g.i == 2 {
		g.swear()
		if g.key == 27 {
			goto L_768a
		}
		g.saySomething()
	}
L_768a:
	if g.i != 3 {
		return
	}
	g.chat()
}

// drawItemMenu  (1010:76b7)
func (g *Game) drawItemMenu() {
	g.clearMenuWin()
	g.resetViewPort()
	g.SetColor(14)
	g.j = 1
	g.i = 1
	for {
		g.itemMenu[g.j] = ""
		if g.inv[g.j] != "" {
			g.itemMenu[g.i] = trunc(g.inv[g.j], 25)
			g.i = g.i + 1
		}
		g.j = g.j + 1
		if g.j > 12 {
			break
		}
	}
	for g.i = 1; ; g.i++ {
		g.SetColor(14)
		if g.itemMenu[g.i] == "Motorfürész" && g.target == 0 {
			g.SetColor(7)
		}
		if g.itemMenu[g.i] == "Gázspray" && g.target == 0 && g.room != 22 {
			g.SetColor(7)
		}
		g.OutTextXY(25, g.i*15+150, g.itemMenu[g.i])
		if g.i == 12 {
			break
		}
	}
	g.SetColor(7)
	g.OutTextXY(25, 350, "Del: törlés")
}

// wrongItem  (1010:784e)
func (g *Game) wrongItem() {
	g.Randomize()
	g.n = g.Random(5) + 1
	if g.n == 1 {
		g.OutTextXY(310, 170, "Ez nem igazán a legjobb eszköz...")
	}
	if g.n == 2 {
		g.OutTextXY(310, 170, "Nem jó. Próbálkozz mással.")
	}
	if g.n == 3 {
		g.OutTextXY(310, 170, "Jobb ötleted nincs?")
	}
	if g.n == 4 {
		g.OutTextXY(310, 170, "Nem valószínü, hogy ez jó ötlet lenne.")
	}
	if g.n != 5 {
		return
	}
	g.OutTextXY(310, 170, "Sajnos ez itt nem segít.")
}

// useItem  (1010:7cf6)
func (g *Game) useItem() {
	g.drawItemMenu()
	g.i = 0
	for {
		g.i = g.i + 1
		if g.i > 12 {
			goto L_7d1c
		}
		if g.inv[g.i] != "" {
			break
		}
	}
L_7d1c:
	if g.i > 12 {
		return
	}
	g.i = 1
L_7d2c:
	for g.n = 1; ; g.n++ {
		g.SetColor(14)
		if g.itemMenu[g.n] == "Motorfürész" && g.target == 0 {
			g.SetColor(7)
		}
		if g.itemMenu[g.n] == "Gázspray" && g.target == 0 && g.room != 22 {
			g.SetColor(7)
		}
		g.OutTextXY(25, g.n*15+150, g.itemMenu[g.n])
		if g.n == 12 {
			break
		}
	}
	g.SetFillStyle(1, 21)
	g.Bar(22, g.i*15+150-1-1, 248, g.i*15+150+10)
	g.SetColor(14)
	if g.itemMenu[g.i] == "Motorfürész" && g.target == 0 {
		g.SetColor(7)
	}
	if g.itemMenu[g.i] == "Gázspray" && g.target == 0 && g.room != 22 {
		g.SetColor(7)
	}
	g.OutTextXY(25, g.i*15+150, g.itemMenu[g.i])
	g.FlushKeys()
	for {
		if g.KeyPressed() {
			break
		}
	}
	g.SetFillStyle(1, 0)
	g.FloodFill(22, g.i*15+150, 0)
	g.key = g.ReadKey()
	if g.key != 0 {
		goto L_7ee2
	}
	g.key = g.ReadKey()
	if g.key == 72 {
		for {
			g.i = g.i - 1
			if g.i < 1 {
				g.i = 12
			}
			if g.itemMenu[g.i] != "" {
				break
			}
		}
	}
	if g.key == 80 {
		g.i = g.i + 1
		if g.itemMenu[g.i] != "" {
			goto L_7ed6
		}
		g.i = 1
	}
L_7ed6:
	if g.key == 83 {
		g.key = 1
	}
L_7ee2:
	if g.i < 1 {
		g.i = 12
	}
	if g.i > 12 {
		g.i = 1
	}
	if g.key == 13 {
		if g.itemMenu[g.i] == "Motorfürész" && g.tgtName == "" {
			g.key = 0
		}
		if g.itemMenu[g.i] != "Gázspray" {
			goto L_7f50
		}
		if g.tgtName != "" {
			goto L_7f50
		}
		if g.room == 22 {
			goto L_7f50
		}
		g.key = 0
	}
L_7f50:
	if g.key != 13 && g.key != 27 && g.key != 1 {
		goto L_7d2c
	}
	if g.key == 27 {
		return
	}
	if g.key == 13 {
		g.s0a2c = trunc(g.itemMenu[g.i], 25)
	}
	if g.key == 1 {
		g.n = 0
		for {
			g.n = g.n + 1
			if g.inv[g.n] == g.itemMenu[g.i] {
				break
			}
		}
		g.inv[g.n] = ""
		g.n = g.i
		g.i = 0
		for {
			g.i = g.i + 1
			if g.taken[g.i] == "" {
				break
			}
		}
		g.taken[g.i] = trunc(g.itemMenu[g.n], 25)
		for g.i = 1; ; g.i++ {
			if g.roomItems[g.i] == g.itemMenu[g.n] {
				g.roomItems[g.i] = ""
			}
			if g.i == 5 {
				break
			}
		}
	}
	g.itemMenu[g.n] = ""
	if g.key == 1 {
		return
	}
	if g.key != 13 {
		return
	}
	g.clearTextWin()
	g.resetViewPort()
	g.SetColor(14)
	if g.s0a2c == "Baseball-ütö" {
		g.OutTextXY(310, 170, "A baseball-ütö használata a Megtámadom")
		g.OutTextXY(310, 185, "menüpontban található.")
	}
	if g.s0a2c == "Gázspray" && g.target != 0 {
		if g.room == 25 {
			for g.i = 1; ; g.i++ {
				if g.inv[g.i] == "Gázspray" {
					g.inv[g.i] = ""
				}
				if g.i == 12 {
					break
				}
			}
		}
		g.line = itoa(g.target)
		if g.target < 10 {
			g.line = "0" + g.line
		}
		g.txt = g.AssignText("data\\trg" + g.line + ".dat")
		g.ResetText(g.txt)
		if g.IOResult() != 0 {
			g.DiskError()
		}
		g.line = "*gasattack"
		for {
			g.line2 = g.txt.ReadLn(255)
			if g.line == g.line2 {
				break
			}
		}
		g.showTextFile()
		g.line2 = g.txt.ReadLn(255)
		g.n = val(g.line2, &g.m2)
		g.m2 = 1
		for {
			if g.chars[g.m2].name == g.tgtName {
				g.chars[g.m2].hp = g.chars[g.m2].hp + g.n
			}
			g.m2 = g.m2 + 1
			if g.m2 > 30 {
				break
			}
		}
		g.line2 = g.txt.ReadLn(255)
		g.n = val(g.line2, &g.m2)
		g.m2 = 1
		for {
			if g.chars[g.m2].name == g.tgtName {
				g.chars[g.m2].state = g.chars[g.m2].state + g.n
			}
			g.m2 = g.m2 + 1
			if g.m2 > 30 {
				break
			}
		}
		g.line2 = g.txt.ReadLn(255)
		g.n = val(g.line2, &g.m2)
		g.energy = g.energy - g.n
		g.CloseText(g.txt)
		for g.i = 1; ; g.i++ {
			if g.inv[g.i] == "Gázspray" {
				g.inv[g.i] = ""
			}
			if g.i == 12 {
				break
			}
		}
		g.i = 0
		for {
			g.i = g.i + 1
			if g.taken[g.i] == "" {
				break
			}
		}
		g.taken[g.i] = trunc("Gázspray", 25)
	}
	if g.s0a2c == "Motorfürész" && g.target != 0 {
		if g.room == 25 {
			for g.i = 1; ; g.i++ {
				if g.inv[g.i] == "Motorfürész" {
					g.inv[g.i] = ""
				}
				if g.i == 12 {
					break
				}
			}
		}
		g.v0866 = g.v0866 - 1
		g.line = itoa(g.target)
		if g.target < 10 {
			g.line = "0" + g.line
		}
		g.txt = g.AssignText("data\\trg" + g.line + ".dat")
		g.ResetText(g.txt)
		if g.IOResult() != 0 {
			g.DiskError()
		}
		g.line = "*chainsaw"
		for {
			g.line2 = g.txt.ReadLn(255)
			if g.line == g.line2 {
				break
			}
		}
		g.showTextFile()
		g.line2 = g.txt.ReadLn(255)
		g.n = val(g.line2, &g.m2)
		g.m2 = 1
		for {
			if g.chars[g.m2].name == g.tgtName {
				g.chars[g.m2].hp = g.chars[g.m2].hp + g.n
			}
			g.m2 = g.m2 + 1
			if g.m2 > 30 {
				break
			}
		}
		g.line2 = g.txt.ReadLn(255)
		g.n = val(g.line2, &g.m2)
		g.m2 = 1
		for {
			if g.chars[g.m2].name == g.tgtName {
				g.chars[g.m2].state = g.chars[g.m2].state + g.n
			}
			g.m2 = g.m2 + 1
			if g.m2 > 30 {
				break
			}
		}
		g.line2 = g.txt.ReadLn(255)
		g.n = val(g.line2, &g.m2)
		g.energy = g.energy - g.n
		g.CloseText(g.txt)
		if g.v0866 <= 0 {
			for g.i = 1; ; g.i++ {
				if g.inv[g.i] == "Motorfürész" {
					g.inv[g.i] = ""
				}
				if g.i == 12 {
					break
				}
			}
		}
		g.i = 0
		for {
			g.i = g.i + 1
			if g.taken[g.i] == "" {
				break
			}
		}
		g.taken[g.i] = trunc("Motorfürész", 25)
	}
	if g.s0a2c == "Baseball-ütö" {
		return
	}
	if g.s0a2c == "Gázspray" && g.room == 22 && g.target != 0 {
		return
	}
	if g.s0a2c == "Gázspray" && g.room != 22 {
		return
	}
	if g.s0a2c == "Motorfürész" {
		return
	}
	if g.room == 1 {
		g.wrongItem()
	}
	if g.room == 2 {
		g.wrongItem()
	}
	if g.room == 4 {
		g.wrongItem()
	}
	if g.room == 5 {
		g.wrongItem()
	}
	if g.room == 6 {
		g.wrongItem()
	}
	if g.room == 7 {
		g.wrongItem()
	}
	if g.room == 8 {
		g.wrongItem()
	}
	if g.room == 9 {
		g.wrongItem()
	}
	if g.room == 10 {
		g.wrongItem()
	}
	if g.room == 11 {
		g.wrongItem()
	}
	if g.room == 12 {
		g.wrongItem()
	}
	if g.room == 13 {
		g.wrongItem()
	}
	if g.room == 14 {
		g.wrongItem()
	}
	if g.room == 15 {
		g.wrongItem()
	}
	if g.room == 17 {
		g.wrongItem()
	}
	if g.room == 18 {
		g.wrongItem()
	}
	if g.room == 19 {
		g.wrongItem()
	}
	if g.room == 20 {
		g.wrongItem()
	}
	if g.room == 21 {
		g.wrongItem()
	}
	if g.room == 23 {
		g.wrongItem()
	}
	if g.room == 27 {
		g.wrongItem()
	}
	if g.room == 28 {
		g.wrongItem()
	}
	if g.room == 29 {
		g.wrongItem()
	}
	if g.room == 30 {
		g.wrongItem()
	}
	if g.room == 25 {
		g.OutTextXY(310, 170, "- Mi az a kezedben? Azonnal add ide!")
		g.OutTextXY(310, 185, "Ellentmondást nem türve kitekeri a")
		g.OutTextXY(310, 200, "kezedböl, és elrakja az asztalfiókba.")
		g.OutTextXY(310, 215, "- Majd év végén visszakapod! És most")
		g.OutTextXY(310, 230, "  kifelé, de gyorsan!")
		for g.i = 1; ; g.i++ {
			if g.inv[g.i] == g.s0a2c {
				g.inv[g.i] = ""
			}
			if g.i == 12 {
				break
			}
		}
		g.i = 0
		for {
			g.i = g.i + 1
			if g.taken[g.i] == "" {
				break
			}
		}
		g.taken[g.i] = trunc(g.s0a2c, 25)
	}
	if g.room != 24 {
		goto L_899a
	}
	if g.s0a2c != "Molotov-koktél" {
		goto L_8986
	}
	g.loaded = 0
	g.i = 0
	for {
		g.i = g.i + 1
		if g.i > 12 {
			goto L_8867
		}
		if g.inv[g.i] == "Öngyújtó" {
			break
		}
	}
L_8867:
	if g.i <= 12 {
		g.loaded = 1
	}
	if g.loaded == 0 {
		g.OutTextXY(310, 170, "Nincs mivel meggyújtanod.")
	}
	if g.loaded == 0 {
		return
	}
	g.OutTextXY(310, 170, "Meggyújtod a rongyot, és résnyire")
	g.OutTextXY(310, 185, "nyitva az ajtót, behajítod a palackot")
	g.OutTextXY(310, 200, "az irodába. Egy perc alatt a lángok")
	g.OutTextXY(310, 215, "martalékává lesz csaknem minden,")
	g.OutTextXY(310, 230, "beleértve a harcias iskolatitkárt is.")
	for g.i = 1; ; g.i++ {
		if g.inv[g.i] == "Molotov-koktél" {
			g.inv[g.i] = ""
		}
		if g.i == 12 {
			break
		}
	}
	g.i = 0
	for {
		g.i = g.i + 1
		if g.taken[g.i] == "" {
			break
		}
	}
	g.taken[g.i] = trunc("Molotov-koktél", 25)
	for g.i = 1; ; g.i++ {
		if g.chars[g.i].name == "Az iskolatitkár" {
			g.chars[g.i].hp = 15
		}
		if g.i == 30 {
			break
		}
	}
	g.target = 0
	g.tgtName = ""
L_8986:
	if g.s0a2c != "Molotov-koktél" {
		g.wrongItem()
	}
L_899a:
	if g.room == 16 {
		if g.s0a2c == "Csont" {
			g.OutTextXY(310, 170, "- Kutyuli! Mutyuli! - gügyögöd az")
			g.OutTextXY(310, 185, "ebnek, és óvatosan felé nyújtod a")
			g.OutTextXY(310, 200, "csontot. Az kikapja a kezedböl, és")
			g.OutTextXY(310, 215, "farkcsóválva elvonul vele valahova.")
			for g.i = 1; ; g.i++ {
				if g.inv[g.i] == "Csont" {
					g.inv[g.i] = ""
				}
				if g.i == 12 {
					break
				}
			}
			g.i = 0
			for {
				g.i = g.i + 1
				if g.taken[g.i] == "" {
					break
				}
			}
			g.taken[g.i] = trunc("Csont", 25)
			for g.i = 1; ; g.i++ {
				if g.chars[g.i].name == "Kutya" {
					g.chars[g.i].hp = 15
				}
				if g.i == 30 {
					break
				}
			}
			g.target = 0
			g.tgtName = ""
		}
		if g.s0a2c == "Csont" {
			goto L_8aae
		}
		g.wrongItem()
	}
L_8aae:
	if g.room == 22 {
		if g.s0a2c == "Gázspray" {
			g.OutTextXY(310, 170, "Elökapod a Cyklon-B sprayt, és")
			g.OutTextXY(310, 185, "az egész flakont a melákokra nyomod.")
			g.OutTextXY(310, 200, "Ez azért már nekik is sok. Amikor")
			g.OutTextXY(310, 215, "a levegö kitisztul, két tüdöembóliás")
			g.OutTextXY(310, 230, "hullát látsz.")
			for g.i = 1; ; g.i++ {
				if g.inv[g.i] == "Gázspray" {
					g.inv[g.i] = ""
				}
				if g.i == 12 {
					break
				}
			}
			g.i = 0
			for {
				g.i = g.i + 1
				if g.taken[g.i] == "" {
					break
				}
			}
			g.taken[g.i] = trunc("Gázspray", 25)
			for g.i = 1; ; g.i++ {
				if g.chars[g.i].name == "Harami, a melák" {
					g.chars[g.i].hp = 15
				}
				if g.i == 30 {
					break
				}
			}
			for g.i = 1; ; g.i++ {
				if g.chars[g.i].name == "Békési, a még melákabb" {
					g.chars[g.i].hp = 15
				}
				if g.i == 30 {
					break
				}
			}
			g.target = 0
			g.tgtName = ""
		}
		if g.s0a2c == "Gázspray" {
			goto L_8c07
		}
		g.wrongItem()
	}
L_8c07:
	if g.room != 3 {
		return
	}
	if g.s0a2c != "Ötszáz forint" {
		goto L_8d0f
	}
	g.lookupTargetByName("Cigány")
	if g.tgtHP < 1 {
		if g.tgtState < 1 {
			goto L_8c5b
		}
	}
	g.OutTextXY(310, 170, "Azt hiszem, már ugrott az")
	g.OutTextXY(310, 185, "üzlet...")
L_8c5b:
	if g.tgtHP == 0 && g.tgtState == 0 {
		g.OutTextXY(310, 170, "Megveszed a cigány gázsprayjét.")
		g.OutTextXY(310, 185, "Megnézegeted a flakont, az van ráírva:")
		g.OutTextXY(310, 200, "\"Cyklon-B\". Nos, ez tényleg nem semmi.")
		for g.i = 1; ; g.i++ {
			if g.inv[g.i] == "Ötszáz forint" {
				g.inv[g.i] = trunc("Gázspray", 25)
			}
			if g.i == 12 {
				break
			}
		}
		g.i = 0
		for {
			g.i = g.i + 1
			if g.taken[g.i] == "" {
				break
			}
		}
		g.taken[g.i] = trunc("Ötszáz forint", 25)
	}
L_8d0f:
	if g.s0a2c == "Ötszáz forint" {
		return
	}
	g.wrongItem()
}

// moveMenu  (1010:907f)
func (g *Game) moveMenu() {
	g.clearMenuWin()
	g.resetViewPort()
	for g.i = 1; ; g.i++ {
		g.OutTextXY(25, g.i*15+150, g.moves[g.i].label)
		if g.i == 6 {
			break
		}
	}
	g.i = 1
L_90c0:
	for g.n = 1; ; g.n++ {
		g.SetColor(14)
		g.OutTextXY(25, g.n*15+150, g.moves[g.n].label)
		if g.n == 6 {
			break
		}
	}
	g.SetFillStyle(1, 21)
	g.Bar(22, g.i*15+150-1-1, 248, g.i*15+150+10)
	g.SetColor(14)
	g.OutTextXY(25, g.i*15+150, g.moves[g.i].label)
	g.FlushKeys()
	for {
		if g.KeyPressed() {
			break
		}
	}
	g.SetFillStyle(1, 0)
	g.FloodFill(22, g.i*15+150, 0)
	g.key = g.ReadKey()
	if g.key == 0 {
		g.key = g.ReadKey()
		if g.key == 72 {
			for {
				g.i = g.i - 1
				if g.i < 1 {
					g.i = 6
				}
				if g.moves[g.i].label != "" {
					break
				}
			}
		}
		if g.key != 80 {
			goto L_91c8
		}
		g.i = g.i + 1
		if g.moves[g.i].label != "" {
			goto L_91c8
		}
		g.i = 1
	}
L_91c8:
	if g.i < 1 {
		g.i = 6
	}
	if g.i > 6 {
		g.i = 1
	}
	if g.key != 13 && g.key != 27 {
		goto L_90c0
	}
	g.j = g.i
	if g.key == 27 {
		return
	}
	g.loaded = 1
	g.clearTextWin()
	g.resetViewPort()
	if g.room == 1 {
		g.lookupTargetByName("A portás")
		if g.tgtHP >= 5 {
			goto L_9244
		}
		if g.j == 4 {
			goto L_9244
		}
		g.loaded = 0
		g.OutTextXY(310, 170, "A portás nem enged.")
	}
L_9244:
	if g.room == 25 && g.j == 1 {
		g.lookupTargetByName("Az iskolatitkár")
		if g.tgtHP < 5 && g.j == 1 {
			g.loaded = 0
			g.OutTextXY(310, 170, "- Hé, mit képzelsz, az igazgató")
			g.OutTextXY(310, 185, "úrhoz nem mehetsz be csak úgy!")
		}
		if g.loaded != 1 {
			goto L_9328
		}
		g.loaded = 0
		for g.i = 1; ; g.i++ {
			if g.inv[g.i] == "Motorfürész" {
				g.loaded = 1
			}
			if g.i == 12 {
				break
			}
		}
		if g.loaded != 0 {
			goto L_9328
		}
		g.OutTextXY(310, 170, "Nyúlsz a kilincs felé, de hirtelen")
		g.OutTextXY(310, 185, "elbizonytalanodsz. Mi van, ha fegyver")
		g.OutTextXY(310, 200, "van a dirinél? Vagy ha nem tudod csak")
		g.OutTextXY(310, 215, "úgy agyonverni? Nem. Ide komoly eszköz")
		g.OutTextXY(310, 230, "kell. Mégsem lépsz be az ajtón.")
	}
L_9328:
	if g.room != 16 {
		goto L_9472
	}
	g.lookupTargetByName("Kutya")
	if g.tgtHP < 5 {
		g.loaded = 0
		g.OutTextXY(310, 170, "Megpróbálsz óvatosan hátrálni, ami")
		g.OutTextXY(310, 185, "sikerül is. Amikor úgy véled, már elég")
		g.OutTextXY(310, 200, "messze vagy, futásnak eredsz. Sajnos")
		g.OutTextXY(310, 215, "nem elég gyorsan. Közeledö lihegést")
		g.OutTextXY(310, 230, "hallasz, és érzed, ahogy a kutya")
		g.OutTextXY(310, 245, "agyarai darabokat szakítanak ki a")
		g.OutTextXY(310, 260, "testedböl. Majd leharapja a fejedet")
		g.OutTextXY(310, 275, "is, és sötétség borul rád.")
		g.energy = 0
	}
	if g.energy == 0 {
		return
	}
	if g.j != 1 {
		goto L_9425
	}
	g.i = 0
	for {
		g.i = g.i + 1
		if g.i > 12 {
			goto L_9409
		}
		if g.inv[g.i] == "Pincekulcs" {
			break
		}
	}
L_9409:
	if g.i > 12 {
		g.OutTextXY(310, 170, "Zárva van az ajtó.")
		g.loaded = 0
	}
L_9425:
	if g.j != 2 {
		goto L_9472
	}
	g.i = 0
	for {
		g.i = g.i + 1
		if g.i > 12 {
			goto L_9456
		}
		if g.inv[g.i] == "Raktárkulcs" {
			break
		}
	}
L_9456:
	if g.i > 12 {
		g.OutTextXY(310, 170, "Zárva van az ajtó.")
		g.loaded = 0
	}
L_9472:
	if g.room == 22 && g.j == 1 && g.flag0 == 0 {
		g.OutTextXY(310, 170, "Az ajtó zárva van. Különös hangokat")
		g.OutTextXY(310, 185, "hallasz bentröl. Nyögéseket.")
		g.loaded = 0
	}
	if g.room == 14 && g.j == 1 && g.flagF == 0 {
		g.OutTextXY(310, 170, "Sötétben tapogatózva elindulsz")
		g.OutTextXY(310, 185, "lefelé a lépcsön. Az orrodig sem")
		g.OutTextXY(310, 200, "látsz. Hirtelen átesel valamin, és")
		g.OutTextXY(310, 215, "szétzúzott fejjel terülsz el a lépcsö")
		g.OutTextXY(310, 230, "aljában. Kalandod véget ért.")
		g.FlushKeys()
		g.WaitKey()
		g.gameOverExit()
	}
	if g.loaded != 1 {
		return
	}
	g.room = g.moves[g.j].room
	if g.room != 26 {
		g.enterRoom()
	}
	if g.room == 26 {
		g.finalScene(g.flagD)
	}
	if g.room != 29 {
		return
	}
	g.szamuelyDeath()
}

// localMenuAction  (1010:a738)
func (g *Game) localMenuAction() {
	g.clearMenuWin()
	g.resetViewPort()
	for g.i = 1; ; g.i++ {
		g.OutTextXY(25, g.i*15+150, g.localMenu[g.i])
		if g.i == 8 {
			break
		}
	}
	g.i = 1
L_a779:
	for g.n = 1; ; g.n++ {
		g.SetColor(14)
		if g.room == 9 && g.n == 3 && g.target == 0 {
			g.SetColor(7)
		}
		g.OutTextXY(25, g.n*15+150, g.localMenu[g.n])
		if g.n == 8 {
			break
		}
	}
	g.SetFillStyle(1, 21)
	g.Bar(22, g.i*15+150-1-1, 248, g.i*15+150+10)
	g.SetColor(14)
	if g.room == 9 && g.i == 3 && g.target == 0 {
		g.SetColor(7)
	}
	g.OutTextXY(25, g.i*15+150, g.localMenu[g.i])
	g.FlushKeys()
	for {
		if g.KeyPressed() {
			break
		}
	}
	g.SetFillStyle(1, 0)
	g.FloodFill(22, g.i*15+150, 0)
	g.key = g.ReadKey()
	if g.key == 0 {
		g.key = g.ReadKey()
		if g.key == 72 {
			for {
				g.i = g.i - 1
				if g.i < 1 {
					g.i = 8
				}
				if g.localMenu[g.i] != "" {
					break
				}
			}
		}
		if g.key != 80 {
			goto L_a8b9
		}
		g.i = g.i + 1
		if g.localMenu[g.i] != "" {
			goto L_a8b9
		}
		g.i = 1
	}
L_a8b9:
	if g.i < 1 {
		g.i = 8
	}
	if g.i > 8 {
		g.i = 1
	}
	if g.key == 13 && g.room == 9 && g.i == 3 && g.target == 0 {
		g.key = 0
	}
	if g.key != 13 && g.key != 27 {
		goto L_a779
	}
	if g.key == 27 {
		return
	}
	g.loaded = 1
	g.j = g.i
	g.clearTextWin()
	g.resetViewPort()
	if g.room == 1 {
		if g.j == 1 {
			g.sub1e9b_0bac()
		}
		if g.j == 2 {
			g.sub1e9b_0c16()
		}
		if g.j != 3 {
			goto L_a954
		}
		g.sub1e9b_0c92()
		g.energy = 0
	}
L_a954:
	if g.room == 30 {
		if g.j == 1 {
			g.OutTextXY(310, 170, "Kotorászásoddal csak azt éred el,")
			g.OutTextXY(310, 185, "hogy egy százas szög fájdalmasan")
			g.OutTextXY(310, 200, "keresztülszúrja az ujjadat.")
			g.energy = g.energy - 10
		}
		if g.j != 2 {
			goto L_a9de
		}
		g.OutTextXY(310, 170, "Játszadozol vele egy kicsit, majd")
		g.OutTextXY(310, 185, "majd nagyot ordítasz. Belefúrtál a")
		g.OutTextXY(310, 200, "kezedbe. Ez nem a te napod.")
		g.energy = g.energy - 20
	}
L_a9de:
	if g.room == 27 {
		if g.j == 1 {
			g.OutTextXY(310, 170, "Semmi válasz.")
		}
		if g.j != 2 {
			goto L_aa2c
		}
		g.OutTextXY(310, 170, "Nekirohansz a vasajtónak, és nagyot")
		g.OutTextXY(310, 185, "koppansz.")
		g.energy = g.energy - 20
	}
L_aa2c:
	if g.room == 28 {
		if g.j == 1 {
			g.OutTextXY(310, 170, "Löszerrel van tele.")
		}
		if g.j != 2 {
			goto L_aa9a
		}
		g.OutTextXY(310, 170, "Belerúgsz a csapódógyújtóba, és")
		g.OutTextXY(310, 185, "gyors búcsút veszel a világtól.")
		g.FlushKeys()
		g.WaitKey()
		g.showRoomPic("nuke.pic")
		g.OutTextXY(310, 215, "              BUMM!")
		g.gameOverExit()
	}
L_aa9a:
	if g.room == 23 {
		g.sub1e9b_3ccc(g.j)
	}
	if g.room == 24 {
		if g.j == 1 {
			g.sub1e9b_0d12()
		}
		if g.j == 2 {
			g.OutTextXY(310, 170, "Te kis illedelmes!")
		}
		if g.j != 3 {
			goto L_ab21
		}
		g.lookupTargetByName("Az iskolatitkár")
		if g.tgtHP < 5 {
			g.OutTextXY(310, 170, "- Az igazgató úr nem fogad! - kiabál")
			g.OutTextXY(310, 185, "ki bentröl az iskolatitkár.")
		}
		if g.tgtHP < 5 {
			goto L_ab21
		}
		g.OutTextXY(310, 170, "Nincs válasz.")
	}
L_ab21:
	if g.room != 25 {
		goto L_ac03
	}
	if g.j == 1 {
		g.lookupTargetByName("Az iskolatitkár")
		if g.tgtHP < 5 {
			g.OutTextXY(310, 170, "- Hé, mit képzel, hagyja azt a")
			g.OutTextXY(310, 185, "  kapcsolót!")
		}
		if g.tgtHP < 5 {
			goto L_ab9d
		}
		g.OutTextXY(310, 170, "Odalépsz, és elfordítod. Semmi")
		g.OutTextXY(310, 185, "látható hatás, de reméled, hogy nem")
		g.OutTextXY(310, 200, "basztál el semmit.")
		g.flagD = 1
	}
L_ab9d:
	if g.j == 2 {
		g.lookupTargetByName("Az iskolatitkár")
		if g.tgtHP < 5 {
			g.OutTextXY(310, 170, "- Fúj! Te disznó! - Elkapja a")
			g.OutTextXY(310, 185, "füledet, és húzza.")
			g.OutTextXY(310, 200, "- Takarítsd le, de azonnal!")
			g.energy = g.energy - 10
		}
		if g.tgtHP < 5 {
			goto L_ac03
		}
		g.OutTextXY(310, 170, "Rosszul vagy?")
	}
L_ac03:
	if g.room != 21 {
		goto L_ad27
	}
	if g.j == 1 {
		g.sub1e9b_0dee()
	}
	if g.j != 2 {
		goto L_ad1b
	}
	g.loaded = 0
	g.i = 0
	for {
		g.i = g.i + 1
		if g.i > 12 {
			goto L_ac52
		}
		if g.inv[g.i] == "Ötszáz forint" {
			break
		}
	}
L_ac52:
	if g.i > 12 {
		g.loaded = 1
	}
	g.i = 0
	for {
		g.i = g.i + 1
		if g.i > 120 {
			goto L_ac88
		}
		if g.taken[g.i] == "Ötszáz forint" {
			break
		}
	}
L_ac88:
	if g.i <= 120 {
		g.loaded = 0
	}
	if g.loaded == 0 {
		g.OutTextXY(310, 170, "Semmi újat nem találsz a ruhákban.")
	}
	if g.loaded != 1 {
		goto L_ad1b
	}
	g.sub1e9b_01c0()
	g.i = 0
	for {
		g.i = g.i + 1
		if g.i > 12 {
			goto L_acd6
		}
		if g.inv[g.i] == "" {
			break
		}
	}
L_acd6:
	if g.i > 12 {
		g.OutTextXY(310, 230, "Sajnos eltenni már nem tudod, mert")
		g.OutTextXY(310, 245, "tele vannak a zsebeid.")
	}
	if g.i <= 12 {
		g.inv[g.i] = trunc("Ötszáz forint", 25)
	}
L_ad1b:
	if g.j == 3 {
		g.sub1e9b_44ff()
	}
L_ad27:
	if g.room == 17 {
		if g.j == 1 {
			g.sub1e9b_448a()
		}
		if g.j == 2 {
			g.sub1e9b_4448()
		}
		if g.j != 3 {
			goto L_ad52
		}
		g.sub1e9b_43e4()
	}
L_ad52:
	if g.room != 22 {
		goto L_aea4
	}
	if g.j == 1 {
		g.sub1e9b_0836()
	}
	if g.j == 3 {
		if g.flag0 == 0 {
			g.OutTextXY(310, 170, "Ujujj! Odabent nagyon jól elvan a")
			g.OutTextXY(310, 185, "két legfiatalabb tanárnö...")
			g.showRoomPic("olympos.pic")
			g.FlushKeys()
			g.WaitKey()
			g.showRoomPic("suli22.pic")
		}
		if g.flag0 != 1 {
			goto L_adcb
		}
		g.OutTextXY(310, 170, "Senki nincs odabent.")
	}
L_adcb:
	if g.j != 2 {
		goto L_aea4
	}
	g.loaded = 1
	g.i = 0
	for {
		g.i = g.i + 1
		if g.inv[g.i] == "Swatch karóra" {
			goto L_ae04
		}
		if g.i > 12 {
			break
		}
	}
L_ae04:
	if g.i <= 12 {
		g.loaded = 0
	}
	g.i = 0
	for {
		g.i = g.i + 1
		if g.taken[g.i] == "Swatch karóra" {
			goto L_ae3a
		}
		if g.i > 120 {
			break
		}
	}
L_ae3a:
	if g.i <= 120 {
		g.loaded = 0
	}
	if g.loaded == 1 {
		g.lookupTargetByName("Harami, a melák")
		if g.tgtHP < 5 {
			g.sub1e9b_0916()
			g.addItem("Swatch karóra")
		}
		if g.tgtHP < 5 {
			goto L_ae7d
		}
		g.sub1e9b_09f4()
		g.addItem("Swatch karóra")
	}
L_ae7d:
	if g.loaded == 0 {
		g.OutTextXY(310, 170, "Kicsit szenilis vagy, fiam. Már")
		g.OutTextXY(310, 185, "egyszer megtetted.")
	}
L_aea4:
	if g.room != 18 {
		goto L_af81
	}
	if g.j == 1 {
		g.sub1e9b_0afa()
	}
	if g.j != 2 {
		goto L_af81
	}
	g.i = 0
	for {
		g.i = g.i + 1
		if g.inv[g.i] == "Öngyújtó" {
			goto L_aeee
		}
		if g.i > 12 {
			break
		}
	}
L_aeee:
	if g.i > 12 {
		g.OutTextXY(310, 170, "Nincs mivel megtenned.")
	}
	if g.i <= 12 {
		g.OutTextXY(310, 170, "Az öreg bútorok kitünöen égnek. Ez")
		g.OutTextXY(310, 185, "nagyon látványos, de nem veszed észre,")
		g.OutTextXY(310, 200, "hogy egy hordó olajfesték is áll")
		g.OutTextXY(310, 215, "mögöttük, a falnál. A lángoló folyadék")
		g.OutTextXY(310, 230, "rád folyik, mint a napalm, és leéget")
		g.OutTextXY(310, 245, "mindent rólad, csak a csontvázadat")
		g.OutTextXY(310, 260, "találhatják meg másnap a romok között.")
		g.gameOverExit()
	}
L_af81:
	if g.room == 20 {
		if g.j == 1 {
			g.OutTextXY(310, 170, "Egy-kettö! Egy-kettö! Leteszed a")
			g.OutTextXY(310, 185, "súlyzót, és máris kurva erösnek érzed")
			g.OutTextXY(310, 200, "magad. Ez persze nem igaz, de legalább")
			g.OutTextXY(310, 215, "hozzájárultál a KISZ Edzett Ifjúságért")
			g.OutTextXY(310, 230, "mozgalomhoz.")
		}
		if g.j != 2 {
			goto L_b042
		}
		g.OutTextXY(310, 170, "Hatalmasat rúgsz egy földön heverö")
		g.OutTextXY(310, 185, "labdába. Majd felordítasz. Ez bizony")
		g.OutTextXY(310, 200, "a suli saját feltalálású, betonnal")
		g.OutTextXY(310, 215, "töltött, 140 kilós medicinjeinek")
		g.OutTextXY(310, 230, "egyike volt...")
		g.energy = g.energy - 20
	}
L_b042:
	if g.room == 15 {
		g.sub1e9b_7b4f(g.j)
	}
	if g.room == 19 {
		if g.j == 1 {
			g.OutTextXY(310, 170, "Összekeversz egy kicsi ezt, egy")
			g.OutTextXY(310, 185, "kicsi azt, megrázod, és...")
			g.FlushKeys()
			g.WaitKey()
			g.clearMenuWin()
			g.resetViewPort()
			g.showRoomPic("nuke.pic")
			g.OutTextXY(310, 215, "             BUMM !!")
			g.OutTextXY(310, 245, "      A kísérlet sikerült!")
			g.gameOverExit()
		}
		if g.j != 2 {
			goto L_b213
		}
		g.Randomize()
		g.n = g.Random(3) + 1
		if g.n == 1 {
			g.OutTextXY(310, 170, "Egy \"Metil-alkohol\" feliratú flaskát")
			g.OutTextXY(310, 185, "hajtasz fel. Bár kémiaórákon azt")
			g.OutTextXY(310, 200, "tanítják, hogy mérgezö, te az ilyen")
			g.OutTextXY(310, 215, "és ehhez hasonló hülyeségeknek nem")
			g.OutTextXY(310, 230, "dölsz be.")
		}
		if g.n == 2 {
			g.OutTextXY(310, 170, "Hm. Mi ez? \"Szuper Pávör Ital\". Vállat")
			g.OutTextXY(310, 185, "vonsz, és legurítod. Érzed, amint")
			g.OutTextXY(310, 200, "hatalmas energia árad szét a testedben.")
			g.OutTextXY(310, 215, "Ez jó választás volt.")
			g.energy = 100
		}
		if g.n != 3 {
			goto L_b213
		}
		g.OutTextXY(310, 170, "Glugg! Glugg! Tényleg, mi ez a különös")
		g.OutTextXY(310, 185, "ízü folyadék? Azt mondja: \"Ciánhiper-")
		g.OutTextXY(310, 200, "dipóltrioxidtoxikum. Mérgezö.")
		g.OutTextXY(310, 215, "Vigyázat, robban!\".")
		g.OutTextXY(310, 230, "- AAARGH! - közlöd a világgal, és")
		g.OutTextXY(310, 245, "halkan a földre zuhansz. A betonhoz")
		g.OutTextXY(310, 260, "ütödsz, és felrobban benned a ")
		g.OutTextXY(310, 275, "vegyszer. Ennyi!")
		g.gameOverExit()
	}
L_b213:
	if g.room == 16 {
		g.lookupTargetByName("Kutya")
		if g.tgtHP < 5 {
			g.OutTextXY(310, 170, "Próbálod nagy ívben kikerülni a dögöt,")
			g.OutTextXY(310, 185, "de pechedre belebotlasz a lábosába,")
			g.OutTextXY(310, 200, "és ezt igencsak zokon veszi. Közeledö")
			g.OutTextXY(310, 215, "lihegést hallasz, és érzed, ahogy a")
			g.OutTextXY(310, 230, "kutya agyarai darabokat szakítanak ki")
			g.OutTextXY(310, 245, "a testedböl. Majd leharapja a fejedet")
			g.OutTextXY(310, 260, "is, és sötétség borul rád.")
			g.energy = 0
			g.j = 0
		}
		if g.j == 1 {
			g.OutTextXY(310, 170, "Felmászol, és kinézel az utcára.")
			g.OutTextXY(310, 185, "Odakint az 1.C. futja a hetvenegyedik")
			g.OutTextXY(310, 200, "kört a suli körül, tesióra keretében.")
			g.OutTextXY(310, 215, "Mármint azok, akik még élnek. De")
			g.OutTextXY(310, 230, "azok igencsak szedik a lábukat, mert")
			g.OutTextXY(310, 245, "Dirándi István, a tesitanár motorral")
			g.OutTextXY(310, 260, "kíséri öket, kezében hosszú nyárssal.")
			g.OutTextXY(310, 275, "Két-három lemaradót már leszúrt vele,")
			g.OutTextXY(310, 290, "és a többiek nyilván nem akarnak")
			g.OutTextXY(310, 305, "hasonló sorsra jutni.")
		}
		if g.j != 2 {
			goto L_b387
		}
		g.OutTextXY(310, 170, "Célozgatod a lépcsöforduló ablakát, de")
		g.OutTextXY(310, 185, "el nem találod. Ilyen szerencsétlen is")
		g.OutTextXY(310, 200, "csak te lehetsz...")
	}
L_b387:
	if g.room == 14 {
		if g.j == 1 {
			g.OutTextXY(310, 170, "És lön világosság!")
			g.flagF = 1
			g.showRoomPic("suli14.pic")
		}
		if g.j != 2 {
			goto L_b3e7
		}
		g.OutTextXY(310, 170, "Klikk! Fény nuku.")
		g.flagF = 0
		g.SetFillStyle(1, 0)
		g.Bar(20, 20, 179, 119)
	}
L_b3e7:
	if g.room == 13 {
		if g.j == 1 {
			g.OutTextXY(310, 170, "\"Igazgatói iroda\"")
		}
		if g.j != 2 {
			goto L_b42f
		}
		if g.flagF == 0 {
			g.sub1e9b_02f6()
		}
		if g.flagF != 1 {
			goto L_b42f
		}
		g.OutTextXY(310, 170, "Egy lefelé vezetö lépcsöt látsz.")
	}
L_b42f:
	if g.room == 12 {
		if g.j == 1 {
			g.sub1e9b_3505()
		}
		if g.j == 2 {
			g.sub1e9b_30d2()
		}
		if g.j != 3 {
			goto L_b45a
		}
		g.sub1e9b_31a6()
	}
L_b45a:
	if g.room == 6 {
		if g.j == 1 {
			g.sub1e9b_5062()
		}
		if g.j == 2 {
			g.sub1e9b_4f7a()
		}
		if g.j != 3 {
			goto L_b4b9
		}
		g.OutTextXY(310, 170, "BANG !!")
		g.OutTextXY(310, 185, "Kis madárkák csipogását hallod.")
		g.OutTextXY(310, 200, "Érdekes hangeffektus.")
		g.energy = g.energy - 50
	}
L_b4b9:
	if g.room != 2 {
		goto L_b522
	}
	if g.j == 1 {
		g.sub1e9b_4e81()
	}
	if g.j == 2 {
		if g.flag0 == 1 {
			g.OutTextXY(310, 170, "Nem valószínü, hogy nagyobb sikert")
			g.OutTextXY(310, 185, "érnél el egy újabb próbálkozással.")
		}
		if g.flag0 != 0 {
			goto L_b50b
		}
		g.sub1e9b_2cbf()
		g.flag0 = 1
	}
L_b50b:
	if g.j == 3 {
		g.OutTextXY(310, 170, "Üres.")
	}
L_b522:
	if g.room == 3 {
		if g.j == 1 {
			g.sub1e9b_4d8a()
		}
		if g.j != 2 {
			goto L_b541
		}
		g.sub1e9b_4c77()
	}
L_b541:
	if g.room == 4 {
		if g.j == 1 {
			g.sub1e9b_4b08()
		}
		if g.j == 2 {
			g.sub1e9b_49e4()
		}
		if g.j != 3 {
			goto L_b5b0
		}
		g.OutTextXY(310, 170, "Recsegést hallasz, és égett szagot")
		g.OutTextXY(310, 185, "érzel. Ugyanakkor nagyon meleged is")
		g.OutTextXY(310, 200, "lesz. Hm... jobb lett volna nem")
		g.energy = g.energy - 70
		g.OutTextXY(310, 215, "piszkálni...")
	}
L_b5b0:
	if g.room != 11 {
		goto L_b78b
	}
	if g.j == 1 {
		g.sub1e9b_48a2()
	}
	if g.j == 2 {
		g.OutTextXY(310, 170, "Auáúú! Forró, mint a napalm! Jól")
		g.OutTextXY(310, 185, "összeégeted a nyelved, de ez még")
		g.OutTextXY(310, 200, "csak a kisebbik baj. Az íze mint")
		g.OutTextXY(310, 215, "a veszett tehén taknyáé!")
		g.energy = g.energy - 10
	}
	if g.j == 3 {
		g.OutTextXY(310, 170, "Felkapsz egy pár tányért a pultról,")
		g.OutTextXY(310, 185, "és földhöz puhintod öket. Csörömpölve")
		g.OutTextXY(310, 200, "törnek apró szilánkokra.")
		g.lookupTargetByName("A szakács")
		if g.tgtHP >= 5 {
			g.OutTextXY(310, 215, "A szakács azonban csak fekszik, arcra")
			g.OutTextXY(310, 230, "borulva, a kiömlött levesben.")
		}
		if g.tgtHP >= 5 {
			goto L_b6e3
		}
		g.OutTextXY(310, 215, "- Hé! Mit csinálsz! - kapja el a")
		g.OutTextXY(310, 230, "a szakács a kezedet.")
		for g.i = 1; ; g.i++ {
			if g.chars[g.i].name == "A szakács" {
				g.chars[g.i].state = g.chars[g.i].state + 1
			}
			if g.i == 30 {
				break
			}
		}
	}
L_b6e3:
	if g.j == 4 {
		g.lookupTargetByName("A szakács")
		if g.tgtHP >= 5 {
			g.OutTextXY(310, 170, "Kivel?")
		}
		if g.tgtHP >= 5 {
			goto L_b78b
		}
		g.sub1e9b_33ad()
		for g.i = 1; ; g.i++ {
			if g.chars[g.i].name == "A szakács" {
				g.chars[g.i].hp = 5
			}
			if g.i == 30 {
				break
			}
		}
		for g.i = 1; ; g.i++ {
			if g.roomPeople[g.i] == "A szakács" {
				g.roomPeople[g.i] = ""
			}
			if g.i == 10 {
				break
			}
		}
		g.tgtName = ""
		g.target = 0
	}
L_b78b:
	if g.room != 5 {
		goto L_b8f2
	}
	g.lookupTargetByName("Tordy Erzsébet tanárnö")
	if g.j == 1 {
		if g.tgtHP >= 5 {
			g.OutTextXY(310, 170, "Rajzolgatsz ezt-azt, de senkit")
			g.OutTextXY(310, 185, "nem érdekel a müvészeted.")
		}
		if g.tgtHP >= 5 {
			goto L_b834
		}
		g.OutTextXY(310, 170, "A tanárnö lelkébe gázoltál ezzel!")
		g.OutTextXY(310, 185, "(Nagyon helyes.)")
		for g.i = 1; ; g.i++ {
			if g.chars[g.i].name == "Tordy Erzsébet tanárnö" {
				g.chars[g.i].state = g.chars[g.i].state + 4
			}
			if g.i == 30 {
				break
			}
		}
	}
L_b834:
	if g.j == 2 {
		g.OutTextXY(310, 170, "ÖNKÉNT? Te hülye vagy?")
		if g.tgtHP < 5 {
			goto L_b862
		}
		g.OutTextXY(310, 185, "(És mellesleg kinél?)")
	}
L_b862:
	if g.j == 3 {
		if g.tgtHP >= 5 {
			g.OutTextXY(310, 170, "Halottakat ne gyalázz!")
		}
		if g.tgtHP >= 5 {
			goto L_b8f2
		}
		g.OutTextXY(310, 170, "Nagyot csípsz a seggébe. Felsikít,")
		g.OutTextXY(310, 185, "és ad egy pofont.")
		for g.i = 1; ; g.i++ {
			if g.chars[g.i].name == "Tordy Erzsébet tanárnö" {
				g.chars[g.i].state = g.chars[g.i].state + 5
			}
			if g.i == 30 {
				break
			}
		}
		g.energy = g.energy - 10
	}
L_b8f2:
	if g.room != 7 {
		goto L_ba10
	}
	if g.j == 1 {
		g.lookupTargetByName("Lepedös Mihály")
		if g.tgtHP >= 5 {
			g.OutTextXY(310, 170, "Szétvésel egy padot, de ez")
			g.OutTextXY(310, 185, "senkit nem zavar immár.")
		}
		if g.tgtHP >= 5 {
			goto L_b99a
		}
		g.OutTextXY(310, 170, "- Fiam, te firkálsz a padra?!")
		g.OutTextXY(310, 185, "Micu bá ideges lett...")
		for g.i = 1; ; g.i++ {
			if g.chars[g.i].name == "Lepedös Mihály" {
				g.chars[g.i].state = g.chars[g.i].state + 1 + 1
			}
			if g.i == 30 {
				break
			}
		}
	}
L_b99a:
	if g.j == 2 {
		g.sub1e9b_482a()
	}
	if g.j == 3 {
		g.OutTextXY(310, 170, "Összekeversz egy kicsi ezt, egy")
		g.OutTextXY(310, 185, "kicsi azt, megrázod, és...")
		g.FlushKeys()
		g.WaitKey()
		g.clearMenuWin()
		g.resetViewPort()
		g.showRoomPic("nuke.pic")
		g.OutTextXY(310, 215, "             BUMM !!")
		g.OutTextXY(310, 245, "      A kísérlet sikerült!")
		g.gameOverExit()
	}
L_ba10:
	if g.room == 8 {
		if g.j == 1 {
			g.OutTextXY(310, 170, "Miért, eddig mit csináltál?")
		}
		if g.j == 2 {
			g.sub1e9b_46c4()
		}
		if g.j == 3 {
			g.sub1e9b_4768()
		}
		if g.j != 4 {
			goto L_ba52
		}
		g.sub1e9b_47ca()
	}
L_ba52:
	if g.room != 9 {
		goto L_bc6a
	}
	if g.j == 1 {
		g.sub1e9b_45ee()
	}
	if g.j == 2 {
		g.sub1e9b_6ea6()
	}
	if g.j == 3 {
		if g.tgtName == "Ahmed, a szíriai bunkó" {
			for g.i = 1; ; g.i++ {
				if g.chars[g.i].name == g.tgtName {
					g.chars[g.i].state = g.chars[g.i].state + 1 + 1
				}
				if g.i == 30 {
					break
				}
			}
			g.OutTextXY(310, 170, "Pofándobod Ahmedet egy kanál ")
			g.OutTextXY(310, 185, "spenóttal.")
			g.OutTextXY(310, 200, "- Takarodjál már a halál náthás")
			g.OutTextXY(310, 215, "faszára! - mondja, és törölközik.")
		}
		if g.tgtName == "Suzy, a suli kurvája" {
			for g.i = 1; ; g.i++ {
				if g.chars[g.i].name == g.tgtName {
					g.chars[g.i].state = g.chars[g.i].state + 1 + 1
				}
				if g.i == 30 {
					break
				}
			}
			g.OutTextXY(310, 170, "A kis kurva hajába továbbítasz egy")
			g.OutTextXY(310, 185, "jó zsíros húsdarabot.")
			g.OutTextXY(310, 200, "- Hé! Apukám, te mindig ilyen bunkó")
			g.OutTextXY(310, 215, "  vagy?")
		}
		if g.tgtName == "Petra, a bamba" {
			g.OutTextXY(310, 170, "Némi leveszöldség csapódik Petra")
			g.OutTextXY(310, 185, "bamba szemei közé.")
			g.OutTextXY(310, 200, "- Na! Ki dobál? - néz körül.")
		}
		if g.tgtName != "Egy vén fasz tanár" {
			goto L_bc37
		}
		for g.i = 1; ; g.i++ {
			if g.chars[g.i].name == g.tgtName {
				g.chars[g.i].state = g.chars[g.i].state + 1 + 1
			}
			if g.i == 30 {
				break
			}
		}
		g.sub1e9b_65f4()
	}
L_bc37:
	if g.j == 4 {
		g.lookupTargetByName("Ahmed, a szíriai bunkó")
		if g.tgtHP >= 5 {
			g.sub1e9b_2b34()
		}
		if g.tgtHP >= 5 {
			goto L_bc5e
		}
		g.sub1e9b_3018()
	}
L_bc5e:
	if g.j == 5 {
		g.sub1e9b_42e1()
	}
L_bc6a:
	if g.room != 10 {
		return
	}
	g.Randomize()
	g.n = g.Random(5) + 1
	if g.j == 1 {
		g.sub1e9b_2f18(g.n)
	}
	if g.j == 2 {
		g.OutTextXY(310, 170, "Remek élmény a kagyló mellé hugyozni,")
		g.OutTextXY(310, 185, "hát nem?")
	}
	if g.j != 3 {
		return
	}
	g.j = 1
	for g.i = 1; ; g.i++ {
		if g.inv[g.i] == "Golyóstoll" {
			g.j = 0
		}
		if g.i == 12 {
			break
		}
	}
	if g.j == 0 {
		g.OutTextXY(310, 170, "Felírod a kedvenc együttesed a,")
		g.OutTextXY(310, 185, "barátnöd, és néhány focista nevét.")
	}
	if g.j != 1 {
		return
	}
	g.OutTextXY(310, 170, "Sajnos nincs mivel írnod.")
}

// saveGame  (1010:bd5c)
func (g *Game) saveGame() {
	g.txt = g.AssignText("data\\savegame.tsi")
	g.RewriteText(g.txt)
	g.write(g.txt, itoa(g.room))
	g.writeln(g.txt)
	g.write(g.txt, itoa(g.energy))
	g.writeln(g.txt)
	g.write(g.txt, itoa(g.target))
	g.writeln(g.txt)
	g.write(g.txt, g.tgtName)
	g.writeln(g.txt)
	for g.i = 1; ; g.i++ {
		g.write(g.txt, g.chars[g.i].name)
		g.writeln(g.txt)
		g.write(g.txt, itoa(g.chars[g.i].hp))
		g.writeln(g.txt)
		g.write(g.txt, itoa(g.chars[g.i].state))
		g.writeln(g.txt)
		if g.i == 30 {
			break
		}
	}
	for g.i = 1; ; g.i++ {
		g.write(g.txt, g.inv[g.i])
		g.writeln(g.txt)
		if g.i == 12 {
			break
		}
	}
	for g.i = 1; ; g.i++ {
		g.write(g.txt, g.taken[g.i])
		g.writeln(g.txt)
		if g.i == 120 {
			break
		}
	}
	if g.flagD == 0 {
		g.write(g.txt, itoa(0))
		g.writeln(g.txt)
	}
	if g.flagD == 1 {
		g.write(g.txt, itoa(1))
		g.writeln(g.txt)
	}
	if g.flagF == 0 {
		g.write(g.txt, itoa(0))
		g.writeln(g.txt)
	}
	if g.flagF == 1 {
		g.write(g.txt, itoa(1))
		g.writeln(g.txt)
	}
	if g.flag0 == 0 {
		g.write(g.txt, itoa(0))
		g.writeln(g.txt)
	}
	if g.flag0 == 1 {
		g.write(g.txt, itoa(1))
		g.writeln(g.txt)
	}
	if g.flagE == 0 {
		g.write(g.txt, itoa(0))
		g.writeln(g.txt)
	}
	if g.flagE == 1 {
		g.write(g.txt, itoa(1))
		g.writeln(g.txt)
	}
	g.CloseText(g.txt)
	if g.IOResult() == 0 {
		return
	}
	g.clearTextWin()
	g.resetViewPort()
	g.SetColor(14)
	g.OutTextXY(310, 170, "A mentés elbaszva.")
}

// mainMenu  (1010:bfe5)
func (g *Game) mainMenu() {
	g.clearMenuWin()
	g.resetViewPort()
	g.i = 1
L_bff8:
	for g.n = 1; ; g.n++ {
		g.SetColor(14)
		if g.n == 1 && g.target == 0 {
			g.SetColor(7)
		}
		if g.n == 2 && g.target == 0 {
			g.SetColor(7)
		}
		g.j = 0
		for g.m2 = 1; ; g.m2++ {
			if g.inv[g.m2] != "" {
				g.j = 1
			}
			if g.m2 == 12 {
				break
			}
		}
		if g.n == 3 && g.j == 0 {
			g.SetColor(7)
		}
		g.OutTextXY(25, g.n*15+150, mainMenu[g.n])
		if g.n == 6 {
			break
		}
	}
	g.SetFillStyle(1, 21)
	g.Bar(22, g.i*15+150-1-1, 248, g.i*15+150+10)
	g.SetColor(14)
	if g.i == 1 && g.target == 0 {
		g.SetColor(7)
	}
	if g.i == 2 && g.target == 0 {
		g.SetColor(7)
	}
	g.j = 0
	for g.n = 1; ; g.n++ {
		if g.inv[g.n] != "" {
			g.j = 1
		}
		if g.n == 12 {
			break
		}
	}
	if g.i == 3 && g.j == 0 {
		g.SetColor(7)
	}
	g.OutTextXY(25, g.i*15+150, mainMenu[g.i])
	g.FlushKeys()
	for {
		if g.KeyPressed() {
			break
		}
	}
	g.SetFillStyle(1, 0)
	g.FloodFill(22, g.i*15+150, 0)
	g.key = g.ReadKey()
	if g.key == 0 {
		g.key = g.ReadKey()
		if g.key == 30 {
			g.clearTextWin()
			g.resetViewPort()
			for g.i = 1; ; g.i++ {
				g.chars[g.i].state = 0
				g.chars[g.i].hp = 0
				if g.i == 30 {
					break
				}
			}
			g.SetColor(14)
			g.OutTextXY(310, 170, "És lön feltámadás!")
		}
		if g.key == 48 {
			g.clearTextWin()
			g.resetViewPort()
			g.energy = 100
			g.SetColor(14)
			g.OutTextXY(310, 170, "S ekkor Tomcat meggyógyíta téged.")
		}
		if g.key == 31 {
			g.saveGame()
		}
		if g.key == 72 {
			g.i = g.i - 1
		}
		if g.key != 80 {
			goto L_c236
		}
		g.i = g.i + 1
	}
L_c236:
	if g.i < 1 {
		g.i = 6
	}
	if g.i > 6 {
		g.i = 1
	}
	if g.key == 13 {
		if g.i == 1 && g.target == 0 {
			g.key = 0
		}
		if g.i == 2 && g.target == 0 {
			g.key = 0
		}
		g.j = 0
		for g.n = 1; ; g.n++ {
			if g.inv[g.n] != "" {
				g.j = 1
			}
			if g.n == 12 {
				break
			}
		}
		if g.i != 3 {
			goto L_c2ba
		}
		if g.j != 0 {
			goto L_c2ba
		}
		g.key = 0
	}
L_c2ba:
	if g.key != 13 {
		goto L_bff8
	}
	if g.i == 1 {
		g.attackMenu()
		g.i = 0
	}
	if g.i == 2 {
		g.talkMenu()
		g.i = 0
	}
	if g.i == 3 {
		g.useItem()
		g.i = 0
	}
	if g.i == 4 {
		g.moveMenu()
		g.i = 0
	}
	if g.i == 5 {
		g.localMenuAction()
		g.i = 0
	}
	if g.i != 6 {
		return
	}
	g.otherMenu()
	g.i = 0
}

// mainLoopStep  (1010:c320)
func (g *Game) mainLoopStep() {
	g.mainMenu()
	if g.chars[g.i].hp >= 5 {
		g.chars[g.i].state = 0
	}
	g.lookupTargetByName(g.tgtName)
	if g.tgtHP >= 5 {
		g.target = 0
	}
	g.drawStatus()
}
