// Machine-translated from SULI.EXE segment 1e9b (unit).

package game

// clearMenuWin  (1e9b:0000)
func (g *Game) clearMenuWin() {
	if g.GetMaxY() == 399 {
		g.SetViewPort(20, 160, 250, 370, 1)
		g.ClearViewPort()
	}
	if g.GetMaxY() != 479 {
		return
	}
	g.SetViewPort(20, 200, 250, 410, 1)
	g.ClearViewPort()
}

// clearTextWin  (1e9b:0047)
func (g *Game) clearTextWin() {
	if g.GetMaxY() == 399 {
		g.SetViewPort(300, 160, 620, 370, 1)
		g.ClearViewPort()
	}
	if g.GetMaxY() != 479 {
		return
	}
	g.SetViewPort(300, 200, 620, 410, 1)
	g.ClearViewPort()
}

// resetViewPort  (1e9b:0090)
func (g *Game) resetViewPort() {
	if g.GetMaxY() == 399 {
		g.SetViewPort(0, 0, g.GetMaxX(), g.GetMaxY(), 1)
	}
	if g.GetMaxY() != 479 {
		return
	}
	g.SetViewPort(0, 39, g.GetMaxX(), 439, 1)
}

// sub1e9b_01c0  (1e9b:01c0)
func (g *Game) sub1e9b_01c0() {
	g.OutTextXY(310, 170, "Lopni nem szép, de hát a szükség nagy")
	g.OutTextXY(310, 185, "úr. Végigtúrod az összes göncöt. Ahá!")
	g.OutTextXY(310, 200, "Az egyik nadrág zsebében egy ötszázast")
	g.OutTextXY(310, 215, "találsz!")
}

// sub1e9b_025b  (1e9b:025b)
func (g *Game) sub1e9b_025b() {
	g.OutTextXY(310, 170, "Itt túl sötét van ahhoz, hogy bármit")
	g.OutTextXY(310, 185, "is láss, de feltehetöleg egy lépcsö")
	g.OutTextXY(310, 200, "elött állsz.")
}

// sub1e9b_02f6  (1e9b:02f6)
func (g *Game) sub1e9b_02f6() {
	g.OutTextXY(310, 170, "Odabent sötét van, de amennyire")
	g.OutTextXY(310, 185, "a beszürödö fényben ki tudod venni,")
	g.OutTextXY(310, 200, "egy lefelé vezetö lépcsö van itt.")
}

// sub1e9b_0396  (1e9b:0396)
func (g *Game) sub1e9b_0396() {
	g.OutTextXY(310, 290, "Most is egy ilyen eset tanúi lehetünk.")
	g.OutTextXY(310, 305, "- Menj ki, fiam! Menj ki a teremböl! ")
	g.OutTextXY(310, 320, " - szólal meg, amint belépsz.")
}

// sub1e9b_0422  (1e9b:0422)
func (g *Game) sub1e9b_0422() {
	g.OutTextXY(310, 290, "Mindeme undorító anyaghalmaz közepette")
	g.OutTextXY(310, 305, "egy fehér ruhás ürge sürgölödik egy")
	g.OutTextXY(310, 320, "fakanállal.")
}

// sub1e9b_04ec  (1e9b:04ec)
func (g *Game) sub1e9b_04ec() {
	g.OutTextXY(310, 170, "Zsolt az Adidas Torsionját füzögeti.")
	g.OutTextXY(310, 185, "Láthatólag még mindig azon")
	g.OutTextXY(310, 200, "bosszankodik, hogy tegnap a ")
	g.OutTextXY(310, 215, "billiárdklubban hagyta a fémellenzös")
	g.OutTextXY(310, 230, "baseball-sapkáját.")
}

// sub1e9b_0597  (1e9b:0597)
func (g *Game) sub1e9b_0597() {
	g.OutTextXY(310, 170, "A kis pöcs Agócs az ajtó felé pislog.")
	g.OutTextXY(310, 185, "Nyilván le akar lépni, beköpni téged")
	g.OutTextXY(310, 200, "a dirinél.")
}

// sub1e9b_063a  (1e9b:063a)
func (g *Game) sub1e9b_063a() {
	g.OutTextXY(310, 170, "- Kisfiam, nem vagy te egy kicsit")
	g.OutTextXY(310, 185, "nagyon neveletlen? - érdeklödik Tordy")
	g.OutTextXY(310, 200, "Erzsébet némi döbbent hallgatás után.")
}

// sub1e9b_06c6  (1e9b:06c6)
func (g *Game) sub1e9b_06c6() {
	g.OutTextXY(310, 170, "A kis pöcs Agócs beköpi Lakit, hogy")
	g.OutTextXY(310, 185, "olvas a pad alatt. Micu bá jól")
	g.OutTextXY(310, 200, "lebassza a rockert.")
}

// sub1e9b_0760  (1e9b:0760)
func (g *Game) sub1e9b_0760() {
	g.OutTextXY(310, 170, "Váév éppen bambán mered a monitorra.")
	g.OutTextXY(310, 185, "Valaha még programozni is tudott, de")
	g.OutTextXY(310, 200, "mára ezt már elfelejtette.")
}

// sub1e9b_0836  (1e9b:0836)
func (g *Game) sub1e9b_0836() {
	g.OutTextXY(310, 170, "\"Házirend. Az épületben tilos a")
	g.OutTextXY(310, 185, "dohányzás. Az épületben tilos az")
	g.OutTextXY(310, 200, "étkezés. Az épületben tilos szeszes")
	g.OutTextXY(310, 215, "italt fogyasztani...\"")
	g.OutTextXY(310, 230, "Szerencsére a kefélést kifelejtették.")
}

// sub1e9b_0916  (1e9b:0916)
func (g *Game) sub1e9b_0916() {
	g.OutTextXY(310, 170, "Kivárod a megfelelö pillanatot, és")
	g.OutTextXY(310, 185, "villámgyors mozdulattal lecsatolod a")
	g.OutTextXY(310, 200, "karórát a nagydarab fazon csuklójáról.")
	g.OutTextXY(310, 215, "Wow, ez egy eredeti Swatch!")
}

// sub1e9b_09f4  (1e9b:09f4)
func (g *Game) sub1e9b_09f4() {
	g.OutTextXY(310, 170, "Egy kicsi hullarablás senkinek sem")
	g.OutTextXY(310, 185, "árt. E filozófiai tantételböl")
	g.OutTextXY(310, 200, "kiindulva leveszed a karórát a")
	g.OutTextXY(310, 215, "nagydarab fazon csuklójáról. Wow, ez")
	g.OutTextXY(310, 230, "egy eredeti Swatch!")
}

// sub1e9b_0afa  (1e9b:0afa)
func (g *Game) sub1e9b_0afa() {
	g.OutTextXY(310, 170, "Egy vastag, kézzel írt könyvet")
	g.OutTextXY(310, 185, "találsz. A címe: \"Mein Kampf\". Az író")
	g.OutTextXY(310, 200, "neve olvashatatlan, és ráadásul az")
	g.OutTextXY(310, 215, "egész kézirat németül van. A franc")
	g.OutTextXY(310, 230, "enné meg. Visszadobod a kacatok közé.")
}

// sub1e9b_0bac  (1e9b:0bac)
func (g *Game) sub1e9b_0bac() {
	g.OutTextXY(310, 170, "\"Gyüjtsd a vasat, és a fémet,")
	g.OutTextXY(310, 185, " Azzal is a békét véded.\"")
	g.OutTextXY(310, 200, "Láthatóan ez egy elég régi felhívás.")
}

// sub1e9b_0c16  (1e9b:0c16)
func (g *Game) sub1e9b_0c16() {
	g.OutTextXY(310, 170, "Ücsörögsz egy kicsit, de semmi")
	g.OutTextXY(310, 185, "érdekes nem történik.")
}

// sub1e9b_0c92  (1e9b:0c92)
func (g *Game) sub1e9b_0c92() {
	g.OutTextXY(310, 170, "BANG !!")
	g.OutTextXY(310, 185, "A hatalmas ajtószárny a rúgástól")
	g.OutTextXY(310, 200, "kifordul a keretéböl, és rád dölve")
	g.OutTextXY(310, 215, "agyonnyom.")
}

// sub1e9b_0d12  (1e9b:0d12)
func (g *Game) sub1e9b_0d12() {
	g.OutTextXY(310, 170, "\"Kolompai Edömér, iskolaigazgató\"")
	g.OutTextXY(310, 185, "Érzed, hogy közel a cél!")
}

// sub1e9b_0dee  (1e9b:0dee)
func (g *Game) sub1e9b_0dee() {
	g.OutTextXY(310, 170, "Hm. Baksa Gábor bakancsa határozottan")
	g.OutTextXY(310, 185, "nem illik össze Marton Erika rózsaszín")
	g.OutTextXY(310, 200, "lastexnadrágjával, és Rabányi Sándor")
	g.OutTextXY(310, 215, "Chicago Bulls dzsekijével. Leveszed")
	g.OutTextXY(310, 230, "hát mindhármat, és visszateszed.")
}

// sub1e9b_0ed8  (1e9b:0ed8)
func (g *Game) sub1e9b_0ed8() {
	g.OutTextXY(310, 170, "Egon magában vihorászva nyomkodja")
	g.OutTextXY(310, 185, "összevissza a gombokat. Láthatóan nagy")
	g.OutTextXY(310, 200, "feladatot állított elé Széchenyi,")
	g.OutTextXY(310, 215, "amikor azt mondta, lépjen be a")
	g.OutTextXY(310, 230, "Windowsba.")
}

// sub1e9b_0ff7  (1e9b:0ff7)
func (g *Game) sub1e9b_0ff7() {
	g.OutTextXY(310, 170, "Széchenyi Gábor tanár úr éppen egy")
	g.OutTextXY(310, 185, "vaskos kézikönyvben keresi, mit is")
	g.OutTextXY(310, 200, "jelent a Bad command or filename")
	g.OutTextXY(310, 215, "hibaüzenet. Amikor megtalálja,")
	g.OutTextXY(310, 230, "értetlenül néz maga elé. Hogyan, hát")
	g.OutTextXY(310, 245, "az XCOPY-t nem I-vel írják?...")
}

// sub1e9b_1120  (1e9b:1120)
func (g *Game) sub1e9b_1120() {
	g.OutTextXY(310, 170, "Okos azon gondolkodik, vajon mi")
	g.OutTextXY(310, 185, "lenne jobb: motorfürésszel alulról")
	g.OutTextXY(310, 200, "kettévágni Tordy Erzsébetet, avagy")
	g.OutTextXY(310, 215, "kiskanállal kimetszeni a csiklóját,")
	g.OutTextXY(310, 230, "esetleg kicsit megismertetni a")
	g.OutTextXY(310, 245, "forrasztópáka örömeivel...")
}

// sub1e9b_11da  (1e9b:11da)
func (g *Game) sub1e9b_11da() {
	g.OutTextXY(310, 170, "Kapiczky Arnold jegyzetel, mint a")
	g.OutTextXY(310, 185, "hülye. Ez azt hiszi, azért jár ide,")
	g.OutTextXY(310, 200, "hogy tanuljon.")
}

// sub1e9b_1267  (1e9b:1267)
func (g *Game) sub1e9b_1267() {
	g.OutTextXY(310, 170, "A folyosókanyar felöl lövések")
	g.OutTextXY(310, 185, "hallatszanak. Egy Kalasnikov, vagy")
	g.OutTextXY(310, 200, "hasonló fegyver lehet.")
}

// sub1e9b_12d7  (1e9b:12d7)
func (g *Game) sub1e9b_12d7() {
	g.OutTextXY(310, 170, "Ücsörögsz egy félórácskát a földön,")
	g.OutTextXY(310, 185, "de ez nem sokat segít.")
}

// sub1e9b_1354  (1e9b:1354)
func (g *Game) sub1e9b_1354() {
	g.OutTextXY(310, 170, "Megint lövéseket hallasz, most egy")
	g.OutTextXY(310, 185, "egész sor géppisztolyét, és egy")
	g.OutTextXY(310, 200, "AKM nehézgéppuskáét.")
}

// sub1e9b_13b6  (1e9b:13b6)
func (g *Game) sub1e9b_13b6() {
	g.OutTextXY(310, 170, "A csend hallgat. (Ah, milyen költöi")
	g.OutTextXY(310, 185, "voltam!)")
}

// sub1e9b_140f  (1e9b:140f)
func (g *Game) sub1e9b_140f() {
	g.OutTextXY(310, 170, "Levél se rezdül. Mellesleg növény")
	g.OutTextXY(310, 185, "sincs a közelben.")
}

// sub1e9b_1476  (1e9b:1476)
func (g *Game) sub1e9b_1476() {
	g.OutTextXY(310, 170, "Nézegetsz ki az ablakon, és vársz,")
	g.OutTextXY(310, 185, "de semmi érdekes nem történik.")
}

// gameOverExit  (1e9b:1b95)
func (g *Game) gameOverExit() {
	g.FlushKeys()
	g.WaitKey()
	g.fadeOutLoadPal(10, "data\\rgb.col")
	g.endSequence()
	g.Delay(1000)
	g.CloseGraph()
	g.writeln(g.Output)
	g.write(g.Output, "  Tomcat Software Incorporated:                          140-05-98, Budapest  ")
	g.writeln(g.Output)
	g.writeln(g.Output)
	g.write(g.Output, "  REVENGE ON SCHOOL  -  BugFixed version!")
	g.writeln(g.Output)
	g.writeln(g.Output)
	g.Halt()
}

// diedExhausted  (1e9b:1d28)
func (g *Game) diedExhausted() {
	g.clearMenuWin()
	g.resetViewPort()
	g.clearTextWin()
	g.resetViewPort()
	g.OutTextXY(310, 170, "Az iskola kegyetlen légköre ismét")
	g.OutTextXY(310, 185, "megtette a magáét. Aléltan fekszel")
	g.OutTextXY(310, 200, "a porban, ahonnan csak délután")
	g.OutTextXY(310, 215, "takarít el a személyzet.")
	g.FlushKeys()
	g.WaitKey()
	g.fadeOutLoadPal(10, "data\\rgb.col")
	g.endSequence()
	g.Delay(1000)
	g.CloseGraph()
	g.writeln(g.Output)
	g.write(g.Output, "  Tomcat Software Incorporated:                          140-05-98, Budapest  ")
	g.writeln(g.Output)
	g.writeln(g.Output)
	g.Halt()
}

// sub1e9b_1f0a  (1e9b:1f0a)
func (g *Game) sub1e9b_1f0a() {
	g.OutTextXY(310, 170, "A 3.A. osztályt úgy is ismerik, mint a")
	g.OutTextXY(310, 185, "suli legrosszabb osztályát. Egy-két")
	g.OutTextXY(310, 200, "tagja még neked is komoly akadályt")
	g.OutTextXY(310, 215, "jelenthet küldetésed teljesítésében.")
	g.OutTextXY(310, 230, "Az osztályfönök egy Lepedös Mihály")
	g.OutTextXY(310, 245, "(alias Micu bá) nevü nevetséges,")
	g.OutTextXY(310, 260, "alkoholista fazon. Kb. a hónaljadig")
	g.OutTextXY(310, 275, "ér, de szereti megjátszani a szigorút.")
}

// sub1e9b_1ff7  (1e9b:1ff7)
func (g *Game) sub1e9b_1ff7() {
	g.OutTextXY(310, 260, "- Megkérlek, drága kisfiam, hogy most")
	g.OutTextXY(310, 275, "azonnal hagyd el a termet... - szól a")
	g.OutTextXY(310, 290, "suli legutáltabb tanárnöje.")
}

// sub1e9b_20f4  (1e9b:20f4)
func (g *Game) sub1e9b_20f4() {
	g.OutTextXY(310, 170, "A számítástechnikát egy Széchenyi")
	g.OutTextXY(310, 185, "Gábor nevü pasas tanítja. Kb. két és")
	g.OutTextXY(310, 200, "fél méter magas, és 62-es inget hord,")
	g.OutTextXY(310, 215, "és igaz, hogy a számítástechnikához")
	g.OutTextXY(310, 230, "nem ért, de a legtöbben azért jó fej")
	g.OutTextXY(310, 245, "tanárnak tartják.")
}

// sub1e9b_2227  (1e9b:2227)
func (g *Game) sub1e9b_2227() {
	g.OutTextXY(310, 170, "A 4.B. terme semmiben sem különbözik")
	g.OutTextXY(310, 185, "az iskola többi osztályainak")
	g.OutTextXY(310, 200, "termeitöl: szürke padok, kemény")
	g.OutTextXY(310, 215, "székek, drága tanári asztal és")
	g.OutTextXY(310, 230, "kényelmes, puha tanári szék. A táblán")
	g.OutTextXY(310, 245, "németóra nyomai: \"ein, zwei, polizei\".")
}

// sub1e9b_232a  (1e9b:232a)
func (g *Game) sub1e9b_232a() {
	g.OutTextXY(310, 260, "Igen, ez valóban németóra: szigorú")
	g.OutTextXY(310, 275, "kussban, fület-farkat behúzva")
	g.OutTextXY(310, 290, "lapítanak a srácok, miközben Tordy")
	g.OutTextXY(310, 305, "Erzsébet tanárnö alakja fenyegetöen")
	g.OutTextXY(310, 320, "tornyosul a katedrán.")
}

// sub1e9b_23d9  (1e9b:23d9)
func (g *Game) sub1e9b_23d9() {
	g.OutTextXY(310, 170, "A folyosókanyarban állsz, a tanári")
	g.OutTextXY(310, 185, "szoba elött. A falra rajzszögezve")
	g.OutTextXY(310, 200, "valami papírt látsz.")
}

// sub1e9b_24cd  (1e9b:24cd)
func (g *Game) sub1e9b_24cd() {
	g.OutTextXY(310, 170, "Az iskolaudvarra nyílik az ajtó. Semmi")
	g.OutTextXY(310, 185, "különöset nem látsz, csak három ajtót:")
	g.OutTextXY(310, 200, "az egyik a pincébe, a másik egy ")
	g.OutTextXY(310, 215, "raktárféle helyiségbe nyílik, a")
	g.OutTextXY(310, 230, "harmadik pedig az udvarra vezetö")
	g.OutTextXY(310, 245, "lépcsö ajtaja.")
}

// sub1e9b_25de  (1e9b:25de)
func (g *Game) sub1e9b_25de() {
	g.OutTextXY(310, 260, "Töprengeni azonban nincs idöd, mert")
	g.OutTextXY(310, 275, "fogvicsorítva ugrik eléd Buksi, a")
	g.OutTextXY(310, 290, "pedellus kutyája. Anyja pit-bull volt,")
	g.OutTextXY(310, 305, "apja pedig doberman, szóval nem sok")
	g.OutTextXY(310, 320, "jóra számíthatsz részéröl.")
}

// sub1e9b_268f  (1e9b:268f)
func (g *Game) sub1e9b_268f() {
	g.OutTextXY(310, 215, "A két izomagy még mindig itt van.")
	g.OutTextXY(310, 230, "- Már megint itt vagy? - veti oda")
	g.OutTextXY(310, 245, "Harami a cigije mellöl.")
}

// sub1e9b_271b  (1e9b:271b)
func (g *Game) sub1e9b_271b() {
	g.OutTextXY(310, 215, "Harami és Békési, a suli két izomagyú")
	g.OutTextXY(310, 230, "tanulója támasztja itt a falat, és")
	g.OutTextXY(310, 245, "beszélgetnek.")
}

// sub1e9b_279b  (1e9b:279b)
func (g *Game) sub1e9b_279b() {
	g.OutTextXY(310, 215, "- Tünjél már el, bazmeg, mert eltöröm")
	g.OutTextXY(310, 230, "az arcodat! - szívélyeskedik Békési.")
}

// sub1e9b_2844  (1e9b:2844)
func (g *Game) sub1e9b_2844() {
	g.OutTextXY(310, 170, "A tornaöltözöben az 1.C. cuccaira")
	g.OutTextXY(310, 185, "bukkansz. Nyilván megint odakint")
	g.OutTextXY(310, 200, "futnak, a suli körül. Három ")
	g.OutTextXY(310, 215, "zuhanyfülke is van itt, az egyikben")
}

// sub1e9b_2910  (1e9b:2910)
func (g *Game) sub1e9b_2910() {
	g.OutTextXY(310, 170, "Az iskola elöterében vagy, a")
	g.OutTextXY(310, 185, "portásfülke mellett állsz, a kapunál.")
	g.OutTextXY(310, 200, "Egy lépcsö vezet az emeletre, egy")
	g.OutTextXY(310, 215, "folyosó pedig jobbkéz felé halad.")
}

// sub1e9b_29f7  (1e9b:29f7)
func (g *Game) sub1e9b_29f7() {
	g.OutTextXY(310, 170, "Az iskolatitkár irodájába nyílik")
	g.OutTextXY(310, 185, "az ajtó. Zsuzsa néni (azaz Zsuzsás),")
	g.OutTextXY(310, 200, "az iskolatitkár, akinek mellesleg a")
	g.OutTextXY(310, 215, "nyaka vastagabb a derekánál, felnéz")
	g.OutTextXY(310, 230, "az íróasztal mögül.")
}

// sub1e9b_2aac  (1e9b:2aac)
func (g *Game) sub1e9b_2aac() {
	g.OutTextXY(310, 170, "A kémiaszertárban több polc")
	g.OutTextXY(310, 185, "vegyszerekkel, edényekkel, és egyéb")
	g.OutTextXY(310, 200, "kísérleti eszközökkel van tele.")
}

// sub1e9b_2b34  (1e9b:2b34)
func (g *Game) sub1e9b_2b34() {
	g.OutTextXY(310, 170, "Akkorát csulázol a tányérba, mint egy")
	g.OutTextXY(310, 185, "veréb tollastul. Sokra persze nem mész")
	g.OutTextXY(310, 200, "vele.")
}

// sub1e9b_2cbf  (1e9b:2cbf)
func (g *Game) sub1e9b_2cbf() {
	g.OutTextXY(310, 170, "A tanári szobát hívod fel.")
	g.OutTextXY(310, 185, "- Halló, izé, itt a Fekete Kéz! Egy")
	g.OutTextXY(310, 200, "  perc múlva, izé, felrobbantom az")
	g.OutTextXY(310, 215, "  iskolát! Menekülni!")
	g.OutTextXY(310, 230, "Kisvártatva két tanárnö rohan ki a")
	g.OutTextXY(310, 245, "kapun, mint a vihar. A többieknek,")
	g.OutTextXY(310, 260, "úgy látszik, nem szóltak a")
	g.OutTextXY(310, 275, "bombariadóról. Kicsit bunkó dolog")
	g.OutTextXY(310, 290, "felrobbanni hagyni 6OO diákot, és egy")
	g.OutTextXY(310, 305, "rakás tanárt, de hát ilyenek a mai")
	g.OutTextXY(310, 320, "magyar erkölcsök.")
}

// sub1e9b_2e02  (1e9b:2e02)
func (g *Game) sub1e9b_2e02() {
	g.OutTextXY(310, 170, "Hm. Nem sok maradt Zsuzsásból. Egy-két")
	g.OutTextXY(310, 185, "égett csont, és szenes íróasztaldarab.")
	g.OutTextXY(310, 200, "Nem kár érte. A kapcsoló a falon")
	g.OutTextXY(310, 215, "azonban szerencsére ép maradt.")
}

// sub1e9b_2f18  (1e9b:2f18)
func (g *Game) sub1e9b_2f18(a6 int) {
	if a6 == 1 {
		g.OutTextXY(310, 170, "\"The password is: Ken sent me.\"")
	}
	if a6 == 2 {
		g.OutTextXY(310, 170, "\"A valóság azoknak való, akik nem")
		g.OutTextXY(310, 185, "bírják a kábítószert.\"")
	}
	if a6 == 3 {
		g.OutTextXY(310, 170, "\"Ne dobd a csikket a WC-be, mert")
		g.OutTextXY(310, 185, "elázik és nehéz meggyújtani.\"")
	}
	if a6 == 4 {
		g.OutTextXY(310, 170, "\"Tomcat szereti Szilvit.\"")
	}
	if a6 != 5 {
		return
	}
	g.OutTextXY(310, 170, "\"Zöld zászló,\"")
	g.OutTextXY(310, 185, "\"Zöld sasok...\"")
}

// sub1e9b_3018  (1e9b:3018)
func (g *Game) sub1e9b_3018() {
	g.OutTextXY(310, 170, "Ahmed mordul egyet, és a kanalával")
	g.OutTextXY(310, 185, "kihajítja a köpést a tányérból.")
	g.OutTextXY(310, 200, "- Maradjál már, bazmeg!")
}

// sub1e9b_30d2  (1e9b:30d2)
func (g *Game) sub1e9b_30d2() {
	g.OutTextXY(310, 170, "A vincsin találsz valami Revenge on")
	g.OutTextXY(310, 185, "School nevü gyenge játékprogramot.")
	g.OutTextXY(310, 200, "Ezzel elszórakozol egy darabig, de")
	g.OutTextXY(310, 215, "azután elunod, és kilépsz.")
}

// sub1e9b_31a6  (1e9b:31a6)
func (g *Game) sub1e9b_31a6() {
	g.OutTextXY(310, 170, "Betöltöd a Turbo Assemblert, és pár")
	g.OutTextXY(310, 185, "perc alatt összedobsz egy vírust, ami")
	g.OutTextXY(310, 200, "minden november 7.-én átírja a C:")
	g.OutTextXY(310, 215, "egység lemezazonosítóját LENIN-re.")
}

// sub1e9b_33ad  (1e9b:33ad)
func (g *Game) sub1e9b_33ad() {
	g.OutTextXY(310, 170, "Körülnézel, hogy mit tehetnél. Ahá!")
	g.OutTextXY(310, 185, "Egy rotyogó olajsütö ötlik a szemedbe.")
	g.OutTextXY(310, 200, "A szakácsra pillantasz, amint épp")
	g.OutTextXY(310, 215, "neked háttal keverget a kondérban.")
	g.OutTextXY(310, 230, "Felröhögsz, és a fejére borítod a")
	g.OutTextXY(310, 245, "fritözt. Borzalmas halálsikolyt")
	g.OutTextXY(310, 260, "hallatva kaparja a sütöt kívülröl, de")
	g.OutTextXY(310, 275, "hamar elunja, és rettenetes kínok")
	g.OutTextXY(310, 290, "közepette pusztul el. Leemeled róla az")
	g.OutTextXY(310, 305, "olajsütöt. A feje megsülve gözölög.")
	g.OutTextXY(310, 320, "A fél arca, és a fejböre a hajával")
	g.OutTextXY(310, 335, "együtt lassan lecsúszik, láthatóvá")
	g.OutTextXY(310, 350, "téve a fehér koponyát.")
}

// sub1e9b_3505  (1e9b:3505)
func (g *Game) sub1e9b_3505() {
	g.OutTextXY(310, 170, "FORMAT C:/U - mondod a gépnek, ami")
	g.OutTextXY(310, 185, "ezt egy pillanat alatt végre is")
	g.OutTextXY(310, 200, "hajtja. Gyorsan átülsz egy másikhoz,")
	g.OutTextXY(310, 215, "mielött valaki észrevenné.")
}

// sub1e9b_3ccc  (1e9b:3ccc)
func (g *Game) sub1e9b_3ccc(a6 int) {
	if a6 == 1 {
		g.OutTextXY(310, 170, "Kizárólag egyesek vannak bennük,")
		g.OutTextXY(310, 185, "meg néhány halotti bizonyítvány,")
		g.OutTextXY(310, 200, "azokról, akik nem bírták a tanítást.")
	}
	if a6 != 2 {
		return
	}
	g.OutTextXY(310, 170, "FUJJ! Mi ez a lónyál!")
}

// sub1e9b_3e08  (1e9b:3e08)
func (g *Game) sub1e9b_3e08() {
	g.OutTextXY(310, 170, "A pince folyosóján vagy, az udvarra")
	g.OutTextXY(310, 185, "nyíló vasajtó elött. Szemben, a")
	g.OutTextXY(310, 200, "folyosó másik végén, egy újabb ajtót")
	g.OutTextXY(310, 215, "látsz, ahol a folyosó balra fordul.")
	g.OutTextXY(310, 230, "Innen is, ahol te állsz, indul")
	g.OutTextXY(310, 245, "balkéz felé egy rövid folyosó.")
	g.OutTextXY(310, 260, "A távolból lövéseket hallasz.")
}

// sub1e9b_3f50  (1e9b:3f50)
func (g *Game) sub1e9b_3f50() {
	g.OutTextXY(310, 170, "Egy dróthálós ajtó áll az utadba.")
	g.OutTextXY(310, 185, "Félresöpröd (nem nehéz, mivel nyitva")
	g.OutTextXY(310, 200, "van), és a pedellus szerszámraktárába")
	g.OutTextXY(310, 215, "jutsz. Fogalmad sincs, minek neki ez")
	g.OutTextXY(310, 230, "a sok szerszám, tekintve hogy soha")
	g.OutTextXY(310, 245, "senki nem látta még dolgozni.")
}

// sub1e9b_4052  (1e9b:4052)
func (g *Game) sub1e9b_4052() {
	g.OutTextXY(310, 170, "A vasajtó közelröl sem erotikusabb,")
	g.OutTextXY(310, 185, "mint a folyosó másik végéböl. Az")
	g.OutTextXY(310, 200, "azonban biztos, hogy a lövések,")
	g.OutTextXY(310, 215, "amiket még mindig hallasz, nem az")
	g.OutTextXY(310, 230, "ajtó mögött dörögnek.")
}

// sub1e9b_41ba  (1e9b:41ba)
func (g *Game) sub1e9b_41ba() {
	g.OutTextXY(310, 170, "Az ajtó mögötti helyiség láthatólag")
	g.OutTextXY(310, 185, "valami fegyverraktár. A falon levö")
	g.OutTextXY(310, 200, "fegyvertartó sajnos üres, de van itt")
	g.OutTextXY(310, 215, "más: egy polc tele mindenféle harci")
	g.OutTextXY(310, 230, "kacattal, egy kisebbfajta repülöbomba")
	g.OutTextXY(310, 245, "a sarokban, és egy ismeretlen")
	g.OutTextXY(310, 260, "rendeltetésü faláda, \"Vigyázat,")
	g.OutTextXY(310, 275, "robban!\" felirattal a tetején.")
}

// sub1e9b_42e1  (1e9b:42e1)
func (g *Game) sub1e9b_42e1() {
	g.OutTextXY(310, 170, "Azt mondja, hogy: \"Szerda. Ebéd:")
	g.OutTextXY(310, 185, "zöldségleves, sült hús, burgonya.\"")
	g.OutTextXY(310, 200, "A kondérokra pillantasz, a")
	g.OutTextXY(310, 215, "konyhaablak mögött. Meg tudod érteni,")
	g.OutTextXY(310, 230, "miért kellett kiírni, mi ez.")
}

// sub1e9b_43e4  (1e9b:43e4)
func (g *Game) sub1e9b_43e4() {
	g.OutTextXY(310, 170, "Lefekszel a földre, és becsukod a")
	g.OutTextXY(310, 185, "szemedet. A nagyobb illúzió kedvéért")
	g.OutTextXY(310, 200, "még a lélegzetedet is visszatartod,")
	g.OutTextXY(310, 215, "de hiába, az istennek sem tudsz")
	g.OutTextXY(310, 230, "meghalni. Dühösen tápászkodsz fel.")
}

// sub1e9b_4448  (1e9b:4448)
func (g *Game) sub1e9b_4448() {
	g.OutTextXY(310, 170, "Félsz egyedül?")
}

// sub1e9b_448a  (1e9b:448a)
func (g *Game) sub1e9b_448a() {
	g.OutTextXY(310, 170, "A folyosókanyar irányából lövéseket")
	g.OutTextXY(310, 185, "hallasz.")
}

// sub1e9b_44ff  (1e9b:44ff)
func (g *Game) sub1e9b_44ff() {
	g.OutTextXY(310, 170, "Ruhástul állsz be a víz alá, és a")
	g.OutTextXY(310, 185, "hatás nem is marad el. Igencsak")
	g.OutTextXY(310, 200, "vizes leszel.")
}

// sub1e9b_45ee  (1e9b:45ee)
func (g *Game) sub1e9b_45ee() {
	g.OutTextXY(310, 170, "Különös íze van. Valahol a friss")
	g.OutTextXY(310, 185, "medúza és a svábbogár zamata között")
	g.OutTextXY(310, 200, "állhat. Hirtelen valami keményet")
	g.OutTextXY(310, 215, "érzel a fogad alatt. Kiveszed, és")
	g.OutTextXY(310, 230, "megnézed. Ah, semmi, csak egy kis")
	g.OutTextXY(310, 245, "patkánykoponya.")
}

// sub1e9b_46c4  (1e9b:46c4)
func (g *Game) sub1e9b_46c4() {
	g.OutTextXY(310, 170, "Ez kissé bajosan menne itt. De hogy")
	g.OutTextXY(310, 185, "így döntöttél, bizonyítja:")
	g.OutTextXY(310, 200, "a program készítöjének hülyesége")
	g.OutTextXY(310, 215, "nem egyedi eset.")
}

// sub1e9b_4768  (1e9b:4768)
func (g *Game) sub1e9b_4768() {
	g.OutTextXY(310, 170, "Miért? Csak nem vagy máris tag?")
	g.OutTextXY(310, 185, "És hol léptél be? Mert én az elözö")
	g.OutTextXY(310, 200, "menüpontban nem engedtelek.")
}

// sub1e9b_47ca  (1e9b:47ca)
func (g *Game) sub1e9b_47ca() {
	g.OutTextXY(310, 170, "Egész nyugodtan. Végül is semmi")
	g.OutTextXY(310, 185, "közöd hozzá.")
}

// sub1e9b_482a  (1e9b:482a)
func (g *Game) sub1e9b_482a() {
	g.OutTextXY(310, 170, "A suli elötti parkot látod. Két")
	g.OutTextXY(310, 185, "cigány éppen biciklit lop.")
}

// sub1e9b_48a2  (1e9b:48a2)
func (g *Game) sub1e9b_48a2() {
	g.OutTextXY(310, 170, "Összegyüjtöd minden nyáladat, és")
	g.OutTextXY(310, 185, "kicsit felhigítod az anyagot.")
	g.OutTextXY(310, 200, "Senkit nem érdekel.")
}

// sub1e9b_49e4  (1e9b:49e4)
func (g *Game) sub1e9b_49e4() {
	g.OutTextXY(310, 170, "Határozott mozdulattal feltéped a")
	g.OutTextXY(310, 185, "lezárt ajtót. A zár reccsenve szakad")
	g.OutTextXY(310, 200, "ki. Sajnos nem találsz semmit.")
	g.OutTextXY(310, 215, "(Elgondolkodtató, hogy 1. mi a fenének")
	g.OutTextXY(310, 230, "zárják be a szekrényt, ha egyszer")
	g.OutTextXY(310, 245, "üres, 2. mi a fenéért zárják be a")
	g.OutTextXY(310, 260, "szekrényt, ha egyszer egy rántással")
	g.OutTextXY(310, 275, "simán ki lehet nyitni.)")
}

// sub1e9b_4b08  (1e9b:4b08)
func (g *Game) sub1e9b_4b08() {
	g.OutTextXY(310, 170, "Jó vicc lenne összespriccelni a")
	g.OutTextXY(310, 185, "folyosót, de sajnos ez a készülék")
	g.OutTextXY(310, 200, "már évek óta bekrepált.")
	g.OutTextXY(310, 215, "(A szervízcédula tanúsága szerint")
	g.OutTextXY(310, 230, "1954-ben ellenörizték utoljára...)")
}

// sub1e9b_4c77  (1e9b:4c77)
func (g *Game) sub1e9b_4c77() {
	g.OutTextXY(310, 170, "Az üvegen át a suli udvarára látsz.")
	g.OutTextXY(310, 185, "Eszedbe jutnak a filmekben látott,")
	g.OutTextXY(310, 200, "vidám nevetéstöl hangos iskolaudvarok,")
	g.OutTextXY(310, 215, "és keserüen elmosolyodsz. Lám, itt")
	g.OutTextXY(310, 230, "az igazság. A diákok négy fal között,")
	g.OutTextXY(310, 245, "szürkeségben tengetik életüket nap")
	g.OutTextXY(310, 260, "mint nap, és soha nem tudhatják meg,")
	g.OutTextXY(310, 275, "mi a nevetés, a napfény...")
}

// sub1e9b_4d8a  (1e9b:4d8a)
func (g *Game) sub1e9b_4d8a() {
	g.OutTextXY(310, 170, "Ücsörögsz egy kicsit, és közben azon")
	g.OutTextXY(310, 185, "morfondírozol, hogy vajon minek volt")
	g.OutTextXY(310, 200, "ez a menüpont a programban, amikor az")
	g.OutTextXY(310, 215, "égvilágon semmi haszna nincs.")
}

// sub1e9b_4e81  (1e9b:4e81)
func (g *Game) sub1e9b_4e81() {
	g.OutTextXY(310, 170, "Valami idióta kulturális adás megy.")
	g.OutTextXY(310, 185, "A kettesen gazdasági vitamüsor, de az")
	g.OutTextXY(310, 200, "MTV-n éppen Beavis & Butt-Head van.")
	g.OutTextXY(310, 215, "Ezt elnézed egy darabig, ám eszedbe")
	g.OutTextXY(310, 230, "jut, hogy küldetésed fontosabb.")
}

// sub1e9b_4f7a  (1e9b:4f7a)
func (g *Game) sub1e9b_4f7a() {
	g.OutTextXY(310, 170, "\"Fiatalok! Gyertek a Papírgyárba!")
	g.OutTextXY(310, 185, "Garantált kereseti lehetöség...\"")
	g.OutTextXY(310, 200, "Miért van az, hogy élö emberrel")
	g.OutTextXY(310, 215, "nem találkoztál még, aki tényleg")
	g.OutTextXY(310, 230, "a Papírgyárba ment volna innen?")
}

// sub1e9b_5062  (1e9b:5062)
func (g *Game) sub1e9b_5062() {
	g.OutTextXY(310, 170, "Mintha csak sovány auschwitzi foglyok")
	g.OutTextXY(310, 185, "fotóit látnád. A felirat:")
	g.OutTextXY(310, 200, "\"4.C. 1986. Találkozunk 1990-ben.\"")
	g.OutTextXY(310, 215, "Elfelejtették odaírni, hogy \"hurrá,")
	g.OutTextXY(310, 230, "túléltük!\".")
}

// sub1e9b_50fe  (1e9b:50fe)
func (g *Game) sub1e9b_50fe() {
	g.OutTextXY(310, 170, "Kis híján kifordul a torkod a szádon,")
	g.OutTextXY(310, 185, "de ez a helyzeten mit sem javít.")
}

// sub1e9b_5180  (1e9b:5180)
func (g *Game) sub1e9b_5180() {
	g.OutTextXY(310, 170, "Az ajtó megremeg, de nem esik ki a")
	g.OutTextXY(310, 185, "keretéböl. Kénytelen leszel a")
	g.OutTextXY(310, 200, "hagyományos módon kinyitni.")
}

// sub1e9b_51f0  (1e9b:51f0)
func (g *Game) sub1e9b_51f0() {
	g.OutTextXY(310, 170, "A lépcsöház remek akusztikája csodás")
	g.OutTextXY(310, 185, "visszhangot produkál.")
}

// sub1e9b_5253  (1e9b:5253)
func (g *Game) sub1e9b_5253() {
	g.OutTextXY(310, 170, "A kalapácsok, reszelök csörömpölnek")
	g.OutTextXY(310, 185, "a polcon a rezonanciától.")
}

// sub1e9b_52b3  (1e9b:52b3)
func (g *Game) sub1e9b_52b3() {
	g.OutTextXY(310, 170, "A lépcsöház remek akusztikája csodás")
	g.OutTextXY(310, 185, "visszhangot produkál.")
}

// sub1e9b_5339  (1e9b:5339)
func (g *Game) sub1e9b_5339() {
	g.OutTextXY(310, 170, "Az üvegek csörömpölve remegnek a")
	g.OutTextXY(310, 185, "polcon, és a vakolat egy darabja")
	g.OutTextXY(310, 200, "leesik. Ez aztán a frekvencia!")
}

// sub1e9b_53a6  (1e9b:53a6)
func (g *Game) sub1e9b_53a6() {
	g.OutTextXY(310, 170, "A folyosót betölti öblös hangod, de")
	g.OutTextXY(310, 185, "ez minden eredmény.")
}

// sub1e9b_53ed  (1e9b:53ed)
func (g *Game) sub1e9b_53ed() {
	g.OutTextXY(310, 170, "A vasajtó belerezonál a hangodba.")
}

// sub1e9b_5414  (1e9b:5414)
func (g *Game) sub1e9b_5414() {
	g.OutTextXY(310, 170, "Emésztés rendben.")
}

// sub1e9b_546a  (1e9b:546a)
func (g *Game) sub1e9b_546a() {
	g.OutTextXY(310, 170, "Micu bát ezúttal nem igazán érdekli")
	g.OutTextXY(310, 185, "a bunkóságod. Már megszokta.")
}

// sub1e9b_54c3  (1e9b:54c3)
func (g *Game) sub1e9b_54c3() {
	g.OutTextXY(310, 170, "Nem jó ötlet. Még valaki azt hinné,")
	g.OutTextXY(310, 185, "ízlett az ebéd.")
}

// sub1e9b_552f  (1e9b:552f)
func (g *Game) sub1e9b_552f() {
	g.OutTextXY(310, 170, "A visszhang végigdübörög a folyosón,")
	g.OutTextXY(310, 185, "és sokhelyütt leveri a vakolatot.")
}

// sub1e9b_55d2  (1e9b:55d2)
func (g *Game) sub1e9b_55d2() {
	g.OutTextXY(310, 170, "A bútorok remegnek a hangodtól, és a")
	g.OutTextXY(310, 185, "feletted levö teremben mindenki")
	g.OutTextXY(310, 200, "hasravágja magát, mert azt hiszik,")
	g.OutTextXY(310, 215, "atomrobbanás történt.")
}

// sub1e9b_5659  (1e9b:5659)
func (g *Game) sub1e9b_5659() {
	g.OutTextXY(310, 170, "Hangod megremegteti a lépcsöházat, de")
	g.OutTextXY(310, 185, "elsöprö sikert nem érsz el.")
}

// sub1e9b_56ed  (1e9b:56ed)
func (g *Game) sub1e9b_56ed() {
	g.OutTextXY(310, 170, "A hangorkán megkopaszt néhány rigót")
	g.OutTextXY(310, 185, "röptében, a szomszédos acélgyár pedig")
	g.OutTextXY(310, 200, "panaszt tesz a rendkívüli zaj miatt.")
}

// sub1e9b_577c  (1e9b:577c)
func (g *Game) sub1e9b_577c() {
	g.OutTextXY(310, 170, "Az ebédlöbe lépve dögletes büz csapja")
	g.OutTextXY(310, 185, "meg az orrodat. Valami érdekeset")
	g.OutTextXY(310, 200, "fözhettek ma is...")
}

// sub1e9b_581d  (1e9b:581d)
func (g *Game) sub1e9b_581d() {
	g.OutTextXY(310, 170, "A földszinti folyosó végében állsz.")
	g.OutTextXY(310, 185, "Két ajtó áll elötted: az egyik a")
	g.OutTextXY(310, 200, "menzára nyílik, a másik a WC-be vezet.")
}

// sub1e9b_5885  (1e9b:5885)
func (g *Game) sub1e9b_5885() {
	g.OutTextXY(310, 170, "A tornaöltözö ajtaja elött állsz.")
	g.OutTextXY(310, 185, "Sehol egy lélek.")
}

// sub1e9b_599c  (1e9b:599c)
func (g *Game) sub1e9b_599c() {
	g.OutTextXY(310, 170, "A folyosó kong az ürességtöl. Sivár,")
	g.OutTextXY(310, 185, "unalmas hely. A falnál ruhásszekrények")
	g.OutTextXY(310, 200, "sorakoznak, mellettük poroltó lóg. Az")
	g.OutTextXY(310, 215, "elektromos elosztódoboz azonban")
	g.OutTextXY(310, 230, "érdeklödésre tarthat számot.")
	g.OutTextXY(310, 245, "Két ajtó is nyílik innen: az egyik a")
	g.OutTextXY(310, 260, "3.A., a másik a 4.B. termébe.")
}

// sub1e9b_5b14  (1e9b:5b14)
func (g *Game) sub1e9b_5b14() {
	g.OutTextXY(310, 170, "A tornaszertárban egyetlen polc kapott")
	g.OutTextXY(310, 185, "helyet a zsámolyok, szekrények, rossz")
	g.OutTextXY(310, 200, "bordásfalak között. Rejtély, hogy az")
	g.OutTextXY(310, 215, "iskolában minek vannak efféle holmik,")
	g.OutTextXY(310, 230, "amikor nincs is tornaterem, és a")
	g.OutTextXY(310, 245, "tornaórák is csak az udvaron vagy a")
	g.OutTextXY(310, 260, "suli körül való rohangálásból állnak.")
}

// sub1e9b_5bf0  (1e9b:5bf0)
func (g *Game) sub1e9b_5bf0() {
	g.OutTextXY(310, 170, "A folyosókanyarhoz érkeztél. Egy nyíl")
	g.OutTextXY(310, 185, "formájú tábla lóg a falon. Balra")
	g.OutTextXY(310, 200, "egy jelzés nélküli ajtó nyílik.")
}

// sub1e9b_5d4f  (1e9b:5d4f)
func (g *Game) sub1e9b_5d4f() {
	g.OutTextXY(310, 170, "Egy raktárhelyiségben vagy, ahol a")
	g.OutTextXY(310, 185, "feleslegessé vált bútorokat, meg egyéb")
	g.OutTextXY(310, 200, "ócska kacatokat tárolják. A kínálat")
	g.OutTextXY(310, 215, "változatos, a törött PVC-csontváztól")
	g.OutTextXY(310, 230, "a rossz írásvetítöig van itt minden.")
	g.OutTextXY(310, 245, "Két ajtó nyílik innen egy-egy kisebb")
	g.OutTextXY(310, 260, "helyiségbe: az egyik a kémia szertár,")
	g.OutTextXY(310, 275, "a másikban meg tornaszereket tárolnak.")
}

// sub1e9b_5eb8  (1e9b:5eb8)
func (g *Game) sub1e9b_5eb8() {
	g.OutTextXY(310, 170, "A portásfülke kicsi, de kényelmes.")
	g.OutTextXY(310, 185, "Eszedbe jut, hogy az a geci, akit az")
	g.OutTextXY(310, 200, "imént levertél, ezreket vágott zsebre")
	g.OutTextXY(310, 215, "azzal, hogy itt ült, és a tévéjét")
	g.OutTextXY(310, 230, "bámulta naphosszat.")
	g.OutTextXY(310, 245, "A falon kulcstartó tábla, sajnos")
	g.OutTextXY(310, 260, "üres. Az asztalon telefon áll.")
}

// sub1e9b_601a  (1e9b:601a)
func (g *Game) sub1e9b_601a() {
	g.OutTextXY(310, 170, "A tanári szoba éppen olyan, mint")
	g.OutTextXY(310, 185, "bármely más suliban. Az asztalon ugyan")
	g.OutTextXY(310, 200, "egy nedves folt csillog, de hát azért")
	g.OutTextXY(310, 215, "EGÉSZEN mégsem lehetnek egyformák a")
	g.OutTextXY(310, 230, "tanári szobák. Egy polcon az")
	g.OutTextXY(310, 245, "osztálynaplók sorakoznak, egy kis")
	g.OutTextXY(310, 260, "asztalon pedig kávégép áll.")
}

// drawFrame  (1e9b:60cf)
func (g *Game) drawFrame() {
	g.ClearDevice()
	g.SetFillStyle(1, 25)
	g.Bar(0, 0, g.GetMaxX(), g.GetMaxY())
	g.SetFillStyle(1, 0)
	g.Bar(20, 20, 179, 119)
	g.SetColor(20)
	g.Line(19, 19, 180, 19)
	g.Line(18, 18, 181, 18)
	g.Line(19, 19, 19, 120)
	g.Line(18, 18, 18, 120)
	g.SetColor(31)
	g.Line(180, 20, 180, 119)
	g.Line(181, 19, 181, 120)
	g.Line(20, 120, 180, 120)
	g.Line(19, 121, 181, 121)
	g.Bar(350, 20, 550, 40)
	g.SetColor(20)
	g.Line(348, 18, 552, 18)
	g.Line(349, 19, 551, 19)
	g.Line(348, 18, 348, 42)
	g.Line(349, 19, 349, 41)
	g.SetColor(31)
	g.Line(349, 42, 552, 42)
	g.Line(350, 41, 551, 41)
	g.Line(551, 42, 551, 20)
	g.Line(552, 42, 552, 19)
	g.SetColor(30)
	g.OutTextXY(250, 29, "Célpont :")
	g.SetColor(20)
	g.OutTextXY(250, 28, "Célpont :")
	g.Bar(350, 70, 550, 90)
	g.SetColor(20)
	g.Line(348, 68, 552, 68)
	g.Line(349, 69, 551, 69)
	g.Line(348, 68, 348, 92)
	g.Line(349, 69, 349, 91)
	g.SetColor(31)
	g.Line(349, 92, 552, 92)
	g.Line(350, 91, 551, 91)
	g.Line(551, 92, 551, 70)
	g.Line(552, 92, 552, 69)
	g.SetColor(30)
	g.OutTextXY(250, 79, "Energia :")
	g.SetColor(20)
	g.OutTextXY(250, 78, "Energia :")
	g.SetColor(30)
	g.OutTextXY(116, 140, "Menü")
	g.SetColor(20)
	g.OutTextXY(116, 139, "Menü")
	g.Bar(20, 160, 250, 370)
	g.SetColor(20)
	g.Line(18, 158, 252, 158)
	g.Line(19, 159, 251, 159)
	g.Line(18, 158, 18, 372)
	g.Line(19, 159, 19, 371)
	g.SetColor(31)
	g.Line(20, 371, 251, 371)
	g.Line(19, 372, 252, 372)
	g.Line(252, 372, 252, 159)
	g.Line(251, 372, 251, 160)
	g.SetColor(30)
	g.OutTextXY(438, 140, "Szöveg")
	g.SetColor(20)
	g.OutTextXY(438, 139, "Szöveg")
	g.Bar(300, 160, 620, 370)
	g.SetColor(20)
	g.Line(299, 158, 622, 158)
	g.Line(299, 159, 621, 159)
	g.Line(298, 158, 298, 372)
	g.Line(299, 159, 299, 371)
	g.SetColor(31)
	g.Line(300, 371, 621, 371)
	g.Line(299, 372, 622, 372)
	g.Line(622, 372, 622, 159)
	g.Line(621, 372, 621, 160)
	g.fadeIn(10)
}

// sub1e9b_6542  (1e9b:6542)
func (g *Game) sub1e9b_6542() {
	g.OutTextXY(310, 170, "Lenyomod a kilincset, és szélesre")
	g.OutTextXY(310, 185, "tárod az ajtót. Odabentröl hirtelen")
	g.OutTextXY(310, 200, "hatalmas sárga felhöben mérges gáz")
	g.OutTextXY(310, 215, "árad ki, és fuldokolni kezdesz.")
}

// sub1e9b_65f4  (1e9b:65f4)
func (g *Game) sub1e9b_65f4() {
	g.OutTextXY(310, 170, "Össztüzet zúdítasz a vén trottyra a")
	g.OutTextXY(310, 185, "maradékokból.")
	g.OutTextXY(310, 200, "- Hé! Fiatalember! Ezt azonnal hagyja")
	g.OutTextXY(310, 215, "  abba! Nem hallja!?")
}

// quitConfirm  (1e9b:66d3)
func (g *Game) quitConfirm() {
	var l1 int
	g.clearTextWin()
	g.resetViewPort()
	g.SetColor(14)
	g.OutTextXY(310, 170, "Biztos? (I/N)")
	g.FlushKeys()
	for {
		if g.KeyPressed() {
			break
		}
	}
	l1 = g.ReadKey()
	if l1 != 73 {
		if l1 != 105 {
			goto L_675b
		}
	}
	g.clearTextWin()
	g.resetViewPort()
	g.SetColor(14)
	g.OutTextXY(310, 170, "Ez a veszélyes iskola tényleg nem")
	g.OutTextXY(310, 185, "neked való. Elöhúzod a késedet, és")
	g.OutTextXY(310, 200, "leugrasz a hídról.")
	g.gameOverExit()
L_675b:
	g.clearTextWin()
	g.resetViewPort()
	g.SetColor(14)
	g.OutTextXY(310, 170, "Akkor meg legközelebb figyelj, hogy")
	g.OutTextXY(310, 185, "hova nyomkodsz!")
}

// sub1e9b_6ea6  (1e9b:6ea6)
func (g *Game) sub1e9b_6ea6() {
	g.OutTextXY(310, 170, "A kékes folyadéknak furcsa, csípös")
	g.OutTextXY(310, 185, "íze van. Nem is csípös, erös. Vagy")
	g.OutTextXY(310, 200, "inkább maró. Vagy inkább... EZ HIPO!")
	g.OutTextXY(310, 215, "Ezt még utoljára megállapítod, mielött")
	g.OutTextXY(310, 230, "rövid exitálás után a másvilágra")
	g.OutTextXY(310, 245, "költöznél.")
	g.diedExhausted()
}

// szamuelyDeath  (1e9b:7851)
func (g *Game) szamuelyDeath() {
	g.clearMenuWin()
	g.resetViewPort()
	g.showRoomPic("suli29.pic")
	g.OutTextXY(310, 170, "A folyosó néhány kanyarja után egy")
	g.OutTextXY(310, 185, "nagy terembe jutsz. Szakadt, kék")
	g.OutTextXY(310, 200, "egyenruhás katonafélék tartanak itt")
	g.OutTextXY(310, 215, "lögyakorlatot. Az uniformisuk")
	g.OutTextXY(310, 230, "valahonnan ismerös. Amíg azon töröd")
	g.OutTextXY(310, 245, "kedvenc fejedet, hogy hol láttál már")
	g.OutTextXY(310, 260, "ilyen katonákat, egy kéz nehezedik")
	g.OutTextXY(310, 275, "a válladra.")
	g.OutTextXY(310, 290, "- Elvtárs! Hol az egyenruhája? Hol a")
	g.OutTextXY(310, 305, "  fegyvere? Kérem az igazolványát!")
	g.OutTextXY(310, 320, "Jaj, ne! Ez az illegális Szamuely")
	g.OutTextXY(310, 335, "Tibor Szocialista Brigád titkos")
	g.OutTextXY(310, 350, "munkásörzászlóalja!")
	g.FlushKeys()
	g.WaitKey()
	g.clearTextWin()
	g.resetViewPort()
	g.clearMenuWin()
	g.resetViewPort()
	g.OutTextXY(310, 170, "- Ööö... izé... - próbálod menteni")
	g.OutTextXY(310, 185, "a helyzetet, de hiába. Elökerül")
	g.OutTextXY(310, 200, "Horváth Adám, a parancsnok, és")
	g.OutTextXY(310, 215, "megvetöen végigmér.")
	g.OutTextXY(310, 230, "- Végezzétek ki ezt az imperialista")
	g.OutTextXY(310, 245, "  kutyát, ezt a kapitalista kémet!")
	g.OutTextXY(310, 260, "Az ítéletet azonnal végre is hajtják.")
	g.OutTextXY(373, 192, ".")
	g.gameOverExit()
}

// sub1e9b_7b4f  (1e9b:7b4f)
func (g *Game) sub1e9b_7b4f(a6 int) {
	if a6 == 1 {
		g.OutTextXY(310, 170, "Két rendör kerget egy cigányt odakinn,")
		g.OutTextXY(310, 185, "a parkban. Ezt elnézegeted egy")
		g.OutTextXY(310, 200, "darabig, de hirtelen az egyik zsaru ")
		g.OutTextXY(310, 215, "elunja a fogócskát, és lelövi a romát.")
		g.OutTextXY(310, 230, "Nagyon látványos.")
	}
	if a6 != 2 {
		return
	}
	g.OutTextXY(310, 170, "Hátrébb húzódsz, majd nekifutva")
	g.OutTextXY(310, 185, "áttöröd testeddel az üveget. Egy")
	g.OutTextXY(310, 200, "emelet magasságból zuhansz Buksinak,")
	g.OutTextXY(310, 215, "a pedellus kutyájának a fejére, aki")
	g.OutTextXY(310, 230, "mindjárt meg is ragadja az alkalmat")
	g.OutTextXY(310, 245, "valamint a torkodat, és ízekre")
	g.OutTextXY(310, 260, "szaggat.")
	g.gameOverExit()
}

// sub1e9b_7c95  (1e9b:7c95)
func (g *Game) sub1e9b_7c95() {
	g.OutTextXY(310, 170, "Egy ajtó zárja le a folyosó végét. Ha")
	g.OutTextXY(310, 185, "nem tévedsz (és nem tévedsz, az")
	g.OutTextXY(310, 200, "biztos), az igazgatói irodához")
	g.OutTextXY(310, 215, "érkeztél.")
}

// sub1e9b_7d3b  (1e9b:7d3b)
func (g *Game) sub1e9b_7d3b() {
	g.OutTextXY(310, 275, "- Idefigyelj, apukám, tüntesd el a")
	g.OutTextXY(310, 290, "  koszos lábnyomaidat a kövemröl, mert")
	g.OutTextXY(310, 305, "  leverem a derekadat!")
}

// sub1e9b_7e10  (1e9b:7e10)
func (g *Game) sub1e9b_7e10() {
	g.OutTextXY(310, 170, "A suli egyetlen WC-je egy kicsit")
	g.OutTextXY(310, 185, "Auschwitz gázkamráit idézi. Ez")
	g.OutTextXY(310, 200, "részben a szarral csurig telt")
	g.OutTextXY(310, 215, "kagylóknak, részint pedig a sok")
	g.OutTextXY(310, 230, "nikotinista tanulónak köszönhetö.")
}

// sub1e9b_7f29  (1e9b:7f29)
func (g *Game) sub1e9b_7f29() {
	g.OutTextXY(310, 170, "Az elsö emeleti folyosón állsz, a")
	g.OutTextXY(310, 185, "lépcsö tetején. A falon érettségi")
	g.OutTextXY(310, 200, "tablók sorakoznak, és egy faliújság")
	g.OutTextXY(310, 215, "díszeleg. Egy ajtó is nyílik")
	g.OutTextXY(310, 230, "balra, mégpedig, ha jól emlékszel,")
	g.OutTextXY(310, 245, "a számítástechnika-terembe.")
}

// sub1e9b_80ad  (1e9b:80ad)
func (g *Game) sub1e9b_80ad() {
	g.OutTextXY(310, 170, "Micsoda disznóól! Kondérokban áll")
	g.OutTextXY(310, 185, "a levesnek elkeresztelt gusztustalan")
	g.OutTextXY(310, 200, "lötty, és a zöld spenót. A tüzhelyen")
	g.OutTextXY(310, 215, "éppen egy újabb adag leves fö. A")
	g.OutTextXY(310, 230, "moslékos vödörben, az ajtó mellett,")
	g.OutTextXY(310, 245, "döglegyek rajai sétálgatnak. Mindenütt")
	g.OutTextXY(310, 260, "ételmaradékok, rothadó húsdarabok")
	g.OutTextXY(310, 275, "hevernek. A szag sem igazán európai.")
}

// sub1e9b_8207  (1e9b:8207)
func (g *Game) sub1e9b_8207() {
	g.OutTextXY(310, 230, "Szerencsére még van annyi idöd, hogy")
	g.OutTextXY(310, 245, "magadra kapd a gázálarcot. A mérgezö")
	g.OutTextXY(310, 260, "anyag hamarosan szétoszlik, és")
	g.OutTextXY(310, 275, "beléphetsz. A kémiaszertárban több")
	g.OutTextXY(310, 290, "polc vegyszerekkel, edényekkel, és")
	g.OutTextXY(310, 305, "egyéb kísérleti eszközökkel van tele.")
}
