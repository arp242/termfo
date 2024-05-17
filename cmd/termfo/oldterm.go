package main

var oldterm = map[string]int{
	// ANSI, UNIX CONSOLE, AND SPECIAL TYPES
	// 386BSD and BSD/OS Consoles
	"origpc3": 1, "origibmpc3": 1, "oldpc3": 1, "oldibmpc3": 1, "bsdos-pc": 1, "bsdos-pc-nobold": 1, "bsdos-pc-m": 1, "bsdos-pc-mono": 1, "pc3": 1, "ibmpc3": 1, "pc3-bold": 1, "bsdos-sparc": 1, "bsdos-ppc": 1,
	// ANSI.SYS/ISO 6429/ECMA-48 Capabilities
	"ibcs2": 1,
	// ANSI/ECMA-48 terminals and terminal emulators
	"ansi77": 1,
	// Atari ST terminals
	"tw52": 1, "tw52-color": 1, "tw52-m": 1, "tt52": 1, "st52-color": 1, "at-color": 1, "atari-color": 1, "atari_st-color": 1, "st52": 1, "st52-m": 1, "at": 1, "at-m": 1, "atari": 1, "atari-m": 1, "atari_st": 1, "atarist-m": 1, "tw100": 1, "stv52": 1, "stv52pc": 1, "atari-old": 1, "uniterm": 1, "uniterm49": 1, "st52-old": 1,
	// BeOS
	"beterm": 1,
	// DEC VT100 and compatibles
	"vt100+keypad": 1, "vt100+pfkeys": 1, "vt100+fnkeys": 1, "vt100+enq": 1, "vt102+enq": 1, "vt100": 1, "vt100-am": 1, "vt100+4bsd": 1, "vt100nam": 1, "vt100-nam": 1, "vt100-vb": 1, "vt100-w": 1, "vt100-w-am": 1, "vt100-w-nam": 1, "vt100-nam-w": 1, "vt100-nav": 1, "vt100-nav-w": 1, "vt100-w-nav": 1, "vt100-s": 1, "vt100-s-top": 1, "vt100-top-s": 1, "vt100-s-bot": 1, "vt100-bot-s": 1, "vt102": 1, "vt102-w": 1, "vt102-nsgr": 1, "vt125": 1, "vt131": 1, "vt132": 1, "vt220-old": 1, "vt200-old": 1, "vt200": 1, "vt200-w": 1, "vt220-8bit": 1, "vt220-8": 1, "vt200-8bit": 1, "vt200-8": 1, "vt220d": 1, "vt200-js": 1, "vt320nam": 1, "vt320": 1, "vt300": 1, "vt320-nam": 1, "vt300-nam": 1, "vt320-w": 1, "vt300-w": 1, "vt320-w-nam": 1, "vt300-w-nam": 1, "vt340": 1, "dec-vt340": 1, "vt330": 1, "dec-vt330": 1, "vt420+lrmm": 1, "vt400": 1, "vt400-24": 1, "dec-vt400": 1, "vt420": 1, "vt420pc": 1, "vt420pcdos": 1, "vt420f": 1, "vt510": 1, "vt510pc": 1, "vt510pcdos": 1, "vt520": 1, "vt525": 1, "vt520ansi": 1,
	// DEC VT52
	"vt52": 1, "vt52-basic": 1, "vt52+arrows": 1,
	// DOS ANSI.SYS variants
	"ansi.sys-old": 1, "ansi.sys": 1, "ansi.sysk": 1, "ansisysk": 1, "nansi.sys": 1, "nansisys": 1, "nansi.sysk": 1, "nansisysk": 1,
	// Mach
	"mach": 1, "mach-bold": 1, "mach-color": 1, "mach-gnu": 1, "mach-gnu-color": 1, "hurd": 1,
	// NetBSD consoles
	"pcvtXX": 1, "pcvt25": 1, "pcvt28": 1, "pcvt35": 1, "pcvt40": 1, "pcvt43": 1, "pcvt50": 1, "pcvt25w": 1, "pcvt28w": 1, "pcvt35w": 1, "pcvt40w": 1, "pcvt43w": 1, "pcvt50w": 1, "pcvt25-color": 1, "arm100": 1, "arm100-am": 1, "arm100-w": 1, "arm100-wam": 1, "x68k": 1, "x68k-ite": 1, "ofcons": 1, "netbsd6": 1, "rcons": 1, "rcons-color": 1, "mgterm": 1,
	// QNX
	"qnx": 1, "qnx4": 1, "qnxt": 1, "qnxt4": 1, "qnxm": 1, "qnxw": 1, "qnxtmono": 1, "qnxt2": 1, "qansi-g": 1, "qansi": 1, "qansi-t": 1, "qansi-m": 1, "qansi-w": 1,
	// SCO consoles
	"scoansi-old": 1, "scoansi-new": 1, "scoansi": 1,
	// SGI consoles
	"iris-ansi": 1, "iris-ansi-net": 1, "iris-ansi-ap": 1, "iris-color": 1, "xwsh": 1,
	// Specials
	"dumb": 1, "unknown": 1, "lpr": 1, "printer": 1, "glasstty": 1, "vanilla": 1, "9term": 1,
	// VT100 emulations
	"dec-vt100": 1, "dec-vt220": 1, "z340": 1, "z340-nam": 1, "tt": 1, "tkterm": 1,
	// APPLE
	// Terminal.app
	"nsterm+7": 1, "nsterm+acs": 1, "nsterm+mac": 1, "nsterm+s": 1, "nsterm+c": 1, "nsterm+c41": 1, "nsterm-7-m": 1, "nsterm-7-m-s": 1, "nsterm-7": 1, "nsterm-7-c": 1, "nsterm-7-s": 1, "nsterm-7-c-s": 1, "nsterm-old": 1, "nsterm-build309": 1, "nsterm-build326": 1, "nsterm-build343": 1, "nsterm-build361": 1, "nsterm-build400": 1, "nsterm-build440": 1,
	// APPLE
	// iTerm, iTerm2
	"xnuppc+basic": 1, "xnuppc+c": 1, "xnuppc+b": 1, "xnuppc+f": 1, "xnuppc+f2": 1, "xnuppc+80x25": 1, "xnuppc+80x30": 1, "xnuppc+90x30": 1, "xnuppc+100x37": 1, "xnuppc+112x37": 1, "xnuppc+128x40": 1, "xnuppc+128x48": 1, "xnuppc+144x48": 1, "xnuppc+160x64": 1, "xnuppc+200x64": 1, "xnuppc+200x75": 1, "xnuppc+256x96": 1, "xnuppc-m": 1, "darwin-m": 1, "xnuppc": 1, "darwin": 1, "xnuppc-m-b": 1, "darwin-m-b": 1, "xnuppc-b": 1, "darwin-b": 1, "xnuppc-m-f": 1, "darwin-m-f": 1, "xnuppc-f": 1, "darwin-f": 1, "xnuppc-m-f2": 1, "darwin-m-f2": 1, "xnuppc-f2": 1, "darwin-f2": 1, "xnuppc-80x25-m": 1, "darwin-80x25-m": 1, "xnuppc-80x25": 1, "darwin-80x25": 1, "xnuppc-80x30-m": 1, "darwin-80x30-m": 1, "xnuppc-80x30": 1, "darwin-80x30": 1, "xnuppc-90x30-m": 1, "darwin-90x30-m": 1, "xnuppc-90x30": 1, "darwin-90x30": 1, "xnuppc-100x37-m": 1, "darwin-100x37-m": 1, "xnuppc-100x37": 1, "darwin-100x37": 1, "xnuppc-112x37-m": 1, "darwin-112x37-m": 1, "xnuppc-112x37": 1, "darwin-112x37": 1, "xnuppc-128x40-m": 1, "darwin-128x40-m": 1, "xnuppc-128x40": 1, "darwin-128x40": 1, "xnuppc-128x48-m": 1, "darwin-128x48-m": 1, "xnuppc-128x48": 1, "darwin-128x48": 1, "xnuppc-144x48-m": 1, "darwin-144x48-m": 1, "xnuppc-144x48": 1, "darwin-144x48": 1, "xnuppc-160x64-m": 1, "darwin-160x64-m": 1, "xnuppc-160x64": 1, "darwin-160x64": 1, "xnuppc-200x64-m": 1, "darwin-200x64-m": 1, "xnuppc-200x64": 1, "darwin-200x64": 1, "xnuppc-200x75-m": 1, "darwin-200x75-m": 1, "xnuppc-200x75": 1, "darwin-200x75": 1, "xnuppc-256x96-m": 1, "darwin-256x96-m": 1, "xnuppc-256x96": 1, "darwin-256x96": 1,
	// COMMERCIAL WORKSTATION CONSOLES
	// Alpha consoles
	"pccons": 1, "pcconsole": 1,
	// Common Desktop Environment
	"dtterm": 1,
	// Iris consoles
	"wsiris": 1, "iris40": 1,
	// NeWS consoles
	"psterm": 1, "psterm-basic": 1, "psterm-96x48": 1, "psterm-90x28": 1, "psterm-80x24": 1, "psterm-fast": 1,
	// NeXT consoles
	"next": 1, "nextshell": 1,
	// Sony NEWS workstations
	"news-unk": 1, "news-29": 1, "news-29-euc": 1, "news-29-sjis": 1, "news-33": 1, "news-33-euc": 1, "news-33-sjis": 1, "news-42": 1, "news-42-euc": 1, "news-42-sjis": 1, "news-old-unk": 1, "nwp512": 1, "news": 1, "nwp514": 1, "news40": 1, "vt100-bm": 1, "nwp512-o": 1, "nwp514-o": 1, "news-o": 1, "news40-o": 1, "vt100-bm-o": 1, "nwp512-a": 1, "nwp514-a": 1, "news-a": 1, "news42": 1, "news40-a": 1, "nwp513": 1, "nwp518": 1, "nwe501": 1, "newscbm": 1, "news31": 1, "nwp513-o": 1, "nwp518-o": 1, "nwe501-o": 1, "nwp251-o": 1, "newscbm-o": 1, "news31-o": 1, "nwp513-a": 1, "nwp518-a": 1, "nwe501-a": 1, "nwp251-a": 1, "newscbm-a": 1, "news31-a": 1, "newscbm33": 1, "news33": 1, "news28": 1, "news29": 1, "news28-a": 1, "nwp511": 1, "nwp-511": 1, "nwp517": 1, "nwp-517": 1, "nwp517-w": 1, "nwp-517-w": 1,
	// Sun consoles
	"oldsun": 1, "sun-il": 1, "sun-cgsix": 1, "sun-ss5": 1, "sun": 1, "sun1": 1, "sun2": 1, "sun+sl": 1, "sun-s": 1, "sun-e-s": 1, "sun-s-e": 1, "sun-48": 1, "sun-34": 1, "sun-24": 1, "sun-17": 1, "sun-12": 1, "sun-1": 1, "sun-e": 1, "sun-nic": 1, "sune": 1, "sun-c": 1, "sun-cmd": 1, "sun-type4": 1, "sun-color": 1,
	// COMMON TERMINAL TYPES
	// Altos
	"altos2": 1, "alt2": 1, "altos-2": 1, "altos3": 1, "altos5": 1, "alt3": 1, "alt5": 1, "altos-3": 1, "altos-5": 1, "altos4": 1, "alt4": 1, "altos-4": 1, "altos7": 1, "alt7": 1, "altos7pc": 1, "alt7pc": 1,
	// Hewlett-Packard (hp)
	"hpgeneric": 1, "hp": 1, "hp110": 1, "hp+pfk+cr": 1, "hp+pfk-cr": 1, "hp+pfk+arrows": 1, "hp+arrows": 1, "hp262x": 1, "hp2621-ba": 1, "hp2621": 1, "hp2621a": 1, "hp2621A": 1, "2621": 1, "2621a": 1, "2621A": 1, "hp2621-wl": 1, "2621-wl": 1, "hp2621-fl": 1, "hp2621p": 1, "hp2621p-a": 1, "hp2621-k45": 1, "hp2621k45": 1, "k45": 1, "hp2621-48": 1, "hp2621-nl": 1, "hp2621-nt": 1, "hp2624": 1, "hp2624a": 1, "hp2624b": 1, "hp2624b-4p": 1, "hp2626": 1, "hp2626a": 1, "hp2626p": 1, "hp2626-s": 1, "hp2626-ns": 1, "hp2626-12": 1, "hp2626-12x40": 1, "hp2626-x40": 1, "hp2626-12-s": 1, "hp2627a-rev": 1, "hp2627a": 1, "hp2627c": 1, "hp2640a": 1, "hp2640b": 1, "hp2644a": 1, "hp2641a": 1, "hp2645a": 1, "hp2647a": 1, "hp2645": 1, "hp45": 1, "hp2648": 1, "hp2648a": 1, "hp150": 1, "hp2382a": 1, "hp2382": 1, "hp2621-a": 1, "hp2621a-a": 1, "newhpkeyboard": 1, "newhp": 1, "memhp": 1, "scrhp": 1, "hp+labels": 1, "hp+printer": 1, "hp2621b": 1, "hp2621b-p": 1, "hp2621b-kx": 1, "hp2621b-kx-p": 1, "hp2622": 1, "hp2622a": 1, "hp2623": 1, "hp2623a": 1, "hp2624b-p": 1, "hp2624b-4p-p": 1, "hp2624-10p": 1, "hp2624a-10p": 1, "hp2624b-10p": 1, "hp2624b-10p-p": 1, "hp+color": 1, "hp2397a": 1, "hp2397": 1, "hpansi": 1, "hp700": 1, "hp2392": 1, "hpsub": 1, "hpex": 1, "hp2": 1, "hpex2": 1, "hp236": 1, "hp300h": 1, "hp9837": 1, "hp98720": 1, "hp98721": 1, "hp9845": 1, "hp98550": 1, "hp98550a": 1, "hp98550-color": 1, "hp98550a-color": 1, "hp700-wy": 1, "hp70092": 1, "hp70092a": 1, "hp70092A": 1, "bobcat": 1, "sbobcat": 1, "gator-t": 1, "gator": 1, "gator-52": 1, "gator-52t": 1,
	// Honeywell-Bull
	"dku7003-dumb": 1, "dku7003": 1,
	// Kermit terminal emulations
	"kermit": 1, "kermit-am": 1, "pckermit": 1, "pckermit12": 1, "pckermit120": 1, "msk227": 1, "mskermit227": 1, "msk227am": 1, "mskermit227am": 1, "msk22714": 1, "mskermit22714": 1, "vt320-k3": 1, "vt320-k311": 1,
	// Lear-Siegler (LSI adm)
	"adm1a": 1, "adm1": 1, "adm2": 1, "adm3": 1, "adm3a": 1, "adm3a+": 1, "adm5": 1, "adm+sgr": 1, "adm11": 1, "adm12": 1, "adm20": 1, "adm21": 1, "adm22": 1, "adm31": 1, "adm31-old": 1, "o31": 1, "adm36": 1, "adm42": 1, "adm42-ns": 1, "adm1178": 1, "1178": 1,
	// Prime
	"pt100": 1, "pt200": 1, "wren": 1, "fenix": 1, "pt100w": 1, "pt200w": 1, "wrenw": 1, "fenixw": 1, "pt250": 1, "pt250w": 1,
	// Qume (qvt)
	"qvt101": 1, "qvt108": 1, "qvt101+": 1, "qvt101p": 1, "qvt102": 1, "qvt103": 1, "qvt103-w": 1, "qvt119+": 1, "qvt119p": 1, "qvt119": 1, "qvt119+-25": 1, "qvt119p-25": 1, "qvt119+-w": 1, "qvt119p-w": 1, "qvt119-w": 1, "qvt119+-25-w": 1, "qvt119p-25-w": 1, "qvt119-25-w": 1, "qvt203": 1, "qvt203+": 1, "qvt203-w": 1, "qvt203-w-am": 1, "qvt203-25": 1, "qvt203-25-w": 1,
	// TeleVideo (tvi)
	"tvi803": 1, "tvi910": 1, "tvi910+": 1, "tvi912": 1, "tvi914": 1, "tvi920": 1, "tvi912cc": 1, "tvi912b-unk": 1, "tvi912c-unk": 1, "tvi912b+printer": 1, "tvi912b+dim": 1, "tvi912b+mc": 1, "tvi912b+2p": 1, "tvi912b+vb": 1, "tvi920b+fn": 1, "tvi912b-2p-unk": 1, "tvi912c-2p-unk": 1, "tvi912b-unk-2p": 1, "tvi912c-unk-2p": 1, "tvi912b-vb-unk": 1, "tvi912c-vb-unk": 1, "tvi912b-unk-vb": 1, "tvi912c-unk-vb": 1, "tvi912b-p": 1, "tvi912c-p": 1, "tvi912b-2p-p": 1, "tvi912c-2p-p": 1, "tvi912b-p-2p": 1, "tvi912c-p-2p": 1, "tvi912b-vb-p": 1, "tvi912c-vb-p": 1, "tvi912b-p-vb": 1, "tvi912c-p-vb": 1, "tvi912b-2p": 1, "tvi912c-2p": 1, "tvi912b-2p-mc": 1, "tvi912c-2p-mc": 1, "tvi912b-mc-2p": 1, "tvi912c-mc-2p": 1, "tvi912b-vb": 1, "tvi912c-vb": 1, "tvi912b-vb-mc": 1, "tvi912c-vb-mc": 1, "tvi912b-mc-vb": 1, "tvi912c-mc-vb": 1, "tvi912b": 1, "tvi912c": 1, "tvi912b-mc": 1, "tvi912c-mc": 1, "tvi920b-unk": 1, "tvi920c-unk": 1, "tvi920b-2p-unk": 1, "tvi920c-2p-unk": 1, "tvi920b-unk-2p": 1, "tvi920c-unk-2p": 1, "tvi920b-vb-unk": 1, "tvi920c-vb-unk": 1, "tvi920b-unk-vb": 1, "tvi920c-unk-vb": 1, "tvi920b-p": 1, "tvi920c-p": 1, "tvi920b-2p-p": 1, "tvi920c-2p-p": 1, "tvi920b-p-2p": 1, "tvi920c-p-2p": 1, "tvi920b-vb-p": 1, "tvi920c-vb-p": 1, "tvi920b-p-vb": 1, "tvi920c-p-vb": 1, "tvi920b-2p": 1, "tvi920c-2p": 1, "tvi920b-2p-mc": 1, "tvi920c-2p-mc": 1, "tvi920b-mc-2p": 1, "tvi920c-mc-2p": 1, "tvi920b-vb": 1, "tvi920c-vb": 1, "tvi920b-vb-mc": 1, "tvi920c-vb-mc": 1, "tvi920b-mc-vb": 1, "tvi920c-mc-vb": 1, "tvi920b": 1, "tvi920c": 1, "tvi920b-mc": 1, "tvi920c-mc": 1, "tvi921": 1, "tvi92B": 1, "tvi92D": 1, "tvi924": 1, "tvi925": 1, "tvi925-hi": 1, "tvi950": 1, "tvi950-2p": 1, "tvi950-4p": 1, "tvi950-rv": 1, "tvi950-rv-2p": 1, "tvi950-rv-4p": 1, "tvi955": 1, "tvi955-w": 1, "955-w": 1, "tvi955-hb": 1, "955-hb": 1, "tvi970": 1, "tvi970-vb": 1, "tvi970-2p": 1, "tvipt": 1, "tvi9065": 1,
	// Visual (vi)
	"vi50": 1, "vi50adm": 1, "vi55": 1, "vi200": 1, "vi200-f": 1, "vi200-rv": 1, "vi300": 1, "vi300-old": 1, "vi500": 1, "vi550": 1, "vi603": 1, "visual603": 1,
	// Wyse (wy)
	"wy30": 1, "wyse30": 1, "wy30-mc": 1, "wyse30-mc": 1, "wy30-vb": 1, "wyse30-vb": 1, "wy50": 1, "wyse50": 1, "wyse+sl": 1, "wy50-mc": 1, "wyse50-mc": 1, "wy50-vb": 1, "wyse50-vb": 1, "wy50-w": 1, "wyse50-w": 1, "wy50-wvb": 1, "wyse50-wvb": 1, "wy350": 1, "wyse350": 1, "wy350-vb": 1, "wyse350-vb": 1, "wy350-w": 1, "wyse350-w": 1, "wy350-wvb": 1, "wyse350-wvb": 1, "wy100": 1, "wy120": 1, "wyse120": 1, "wy150": 1, "wyse150": 1, "wy120-w": 1, "wyse120-w": 1, "wy150-w": 1, "wyse150-w": 1, "wy120-25": 1, "wyse120-25": 1, "wy150-25": 1, "wyse150-25": 1, "wy120-25-w": 1, "wyse120-25-w": 1, "wy150-25-w": 1, "wyse150-25-w": 1, "wy120-vb": 1, "wyse120-vb": 1, "wy150-vb": 1, "wyse150-vb": 1, "wy120-w-vb": 1, "wy120-wvb": 1, "wyse120-wvb": 1, "wy150-w-vb": 1, "wyse150-w-vb": 1, "wy60": 1, "wyse60": 1, "wy60-w": 1, "wyse60-w": 1, "wy60-25": 1, "wyse60-25": 1, "wy60-25-w": 1, "wyse60-25-w": 1, "wy60-42": 1, "wyse60-42": 1, "wy60-42-w": 1, "wyse60-42-w": 1, "wy60-43": 1, "wyse60-43": 1, "wy60-43-w": 1, "wyse60-43-w": 1, "wy60-vb": 1, "wyse60-vb": 1, "wy60-w-vb": 1, "wy60-wvb": 1, "wyse60-wvb": 1, "wy99gt": 1, "wyse99gt": 1, "wy99gt-w": 1, "wyse99gt-w": 1, "wy99gt-25": 1, "wyse99gt-25": 1, "wy99gt-25-w": 1, "wyse99gt-25-w": 1, "wy99gt-vb": 1, "wyse99gt-vb": 1, "wy99gt-w-vb": 1, "wy99gt-wvb": 1, "wyse99gt-wvb": 1, "wy99-ansi": 1, "wy99a-ansi": 1, "wy99f": 1, "wy99fgt": 1, "wy-99fgt": 1, "wy99fa": 1, "wy99fgta": 1, "wy-99fgta": 1, "wy160": 1, "wyse160": 1, "wy160-w": 1, "wyse160-w": 1, "wy160-25": 1, "wyse160-25": 1, "wy160-25-w": 1, "wyse160-25-w": 1, "wy160-42": 1, "wyse160-42": 1, "wy160-42-w": 1, "wyse160-42-w": 1, "wy160-43": 1, "wyse160-43": 1, "wy160-43-w": 1, "wyse160-43-w": 1, "wy160-vb": 1, "wyse160-vb": 1, "wy160-w-vb": 1, "wy160-wvb": 1, "wyse160-wvb": 1, "wy75": 1, "wyse75": 1, "wy75-mc": 1, "wyse75-mc": 1, "wy75-vb": 1, "wyse75-vb": 1, "wy75-w": 1, "wyse75-w": 1, "wy75-wvb": 1, "wyse75-wvb": 1, "wy85": 1, "wyse85": 1, "wy85-vb": 1, "wyse85-vb": 1, "wy85-w": 1, "wyse85-w": 1, "wy85-wvb": 1, "wyse85-wvb": 1, "wy85-8bit": 1, "wyse85-8bit": 1, "wy185": 1, "wyse185": 1, "wy185-24": 1, "wyse185-24": 1, "wy185-vb": 1, "wyse185-vb": 1, "wy185-w": 1, "wyse185-w": 1, "wy185-wvb": 1, "wyse185-wvb": 1, "wy325": 1, "wyse325": 1, "wy325-vb": 1, "wyse325-vb": 1, "wy325-w": 1, "wyse325-w": 1, "wy325w-24": 1, "wy325-25": 1, "wyse325-25": 1, "wy325-80": 1, "wyse-325": 1, "wy325-25w": 1, "wyse325-25w": 1, "wy325-w-vb": 1, "wy325-wvb": 1, "wyse325-wvb": 1, "wy325-42": 1, "wyse325-42": 1, "wy325-42w": 1, "wyse325-42w": 1, "wy325-42w-vb": 1, "wy325-42wvb": 1, "wy325-43": 1, "wyse325-43": 1, "wy325-43w": 1, "wyse325-43w": 1, "wy325-43w-vb": 1, "wy325-43wvb": 1, "wy370-nk": 1, "wy370": 1, "wyse370": 1, "wy370-101k": 1, "wy370-105k": 1, "wy370-EPC": 1, "wy370-vb": 1, "wy370-w": 1, "wy370-wvb": 1, "wy370-rv": 1, "wy99gt-tek": 1, "wy160-tek": 1, "wy370-tek": 1, "wy520": 1, "wyse520": 1, "wy520-24": 1, "wyse520-24": 1, "wy520-vb": 1, "wyse520-vb": 1, "wy520-w": 1, "wyse520-w": 1, "wy520-wvb": 1, "wyse520-wvb": 1, "wy520-epc": 1, "wyse520-epc": 1, "wy520-epc-24": 1, "wyse520-pc-24": 1, "wy520-epc-vb": 1, "wyse520-pc-vb": 1, "wy520-epc-w": 1, "wyse520-epc-w": 1, "wy520-epc-wvb": 1, "wyse520-p-wvb": 1, "wy520-36": 1, "wyse520-36": 1, "wy520-48": 1, "wyse520-48": 1, "wy520-36w": 1, "wyse520-36w": 1, "wy520-48w": 1, "wyse520-48w": 1, "wy520-36pc": 1, "wyse520-36pc": 1, "wy520-48pc": 1, "wyse520-48pc": 1, "wy520-36wpc": 1, "wyse520-36wpc": 1, "wy520-48wpc": 1, "wyse520-48wpc": 1, "wyse-vp": 1, "wy75ap": 1, "wyse75ap": 1, "wy-75ap": 1, "wyse-75ap": 1, "wy100q": 1,
	// DOS/WINDOWS
	// Command prompt
	"ms-vt100": 1, "ms-vt100-color": 1, "vtnt": 1, "ms-vt100+": 1, "vt100+": 1, "ms-vt-utf8": 1, "vt-utf8": 1, "ms-vt100-16color": 1,
	// PuTTY
	"vt100-putty": 1, "putty-vt100": 1, "putty+fnkeys+vt400": 1, "putty+fnkeys+vt100": 1,
	// CRT is very old outdated version of SecureCRT
	"crt": 1, "crt-vt220": 1,
	// LCD DISPLAYS
	// Matrix Orbital
	"MtxOrb": 1, "MtxOrb204": 1, "MtxOrb162": 1,
	// NON-ANSI TERMINAL EMULATIONS
	// Avatar
	"avatar0": 1, "avatar0+": 1, "avatar": 1, "avatar1": 1,
	// RBcomm
	"rbcomm": 1, "rbcomm-nam": 1, "rbcomm-w": 1,
	// Non-Unix Consoles
	// Cygwin
	"cygwinB19": 1, "cygwin": 1, "cygwinDBG": 1,
	// DJGPP
	"djgpp": 1, "djgpp203": 1, "djgpp204": 1,
	// EMX termcap.dat compatibility modes
	"emx-base": 1, "ansi-emx": 1, "ansi-color-2-emx": 1, "ansi-color-3-emx": 1, "mono-emx": 1,
	// Microsoft (miscellaneous)
	"ansi-nt": 1, "psx_ansi": 1, "pcmw": 1, "interix": 1, "opennt": 1, "opennt-25": 1, "ntconsole": 1, "ntconsole-25": 1, "opennt-35": 1, "ntconsole-35": 1, "opennt-50": 1, "ntconsole-50": 1, "opennt-60": 1, "ntconsole-60": 1, "opennt-100": 1, "ntconsole-100": 1, "opennt-w": 1, "opennt-25-w": 1, "ntconsole-w": 1, "ntconsole-25-w": 1, "opennt-35-w": 1, "ntconsole-35-w": 1, "opennt-50-w": 1, "ntconsole-50-w": 1, "opennt-60-w": 1, "ntconsole-60-w": 1, "opennt-w-vt": 1, "opennt-25-w-vt": 1, "ntconsole-w-vt": 1, "ntconsole-25-w-vt": 1, "interix-nti": 1, "opennt-nti": 1, "opennt-25-nti": 1, "ntconsole-25-nti": 1, "opennt-35-nti": 1, "ntconsole-35-nti": 1, "opennt-50-nti": 1, "ntconsole-50-nti": 1, "opennt-60-nti": 1, "ntconsole-60-nti": 1, "opennt-100-nti": 1, "ntconsole-100-nti": 1,
	// U/Win
	"uwin": 1,
	// OBSOLETE PERSONAL-MICRO CONSOLES AND EMULATIONS
	// Apple II
	"appleIIgs": 1, "appleIIe": 1, "appleIIc": 1, "apple2e": 1, "apple2e-p": 1, "apple-ae": 1, "appleII": 1, "apple-80": 1, "apple-soroc": 1, "apple-videx": 1, "apple-uterm-vb": 1, "apple-uterm": 1, "apple80p": 1, "apple-videx2": 1, "apple-videx3": 1, "vapple": 1, "aepro": 1, "apple-vm80": 1, "ap-vm80": 1,
	// Apple Lisa & Macintosh
	"lisa": 1, "liswb": 1, "lisaterm": 1, "lisaterm-w": 1, "mac": 1, "macintosh": 1, "mac-w": 1, "macterminal-w": 1,
	// Commodore Business Machines
	"amiga": 1, "amiga-h": 1, "amiga-8bit": 1, "amiga-vnc": 1, "morphos": 1, "commodore": 1, "b-128": 1,
	// Console types for obsolete UNIX clones
	"minix": 1, "minix-3.0": 1, "minix-1.7": 1, "minix-old": 1, "minix-1.5": 1, "minix-old-am": 1, "pc-minix": 1, "pc-coherent": 1, "pcz19": 1, "coherent": 1, "pc-venix": 1, "venix": 1,
	// IBM PC and clones
	"pcplot": 1, "kaypro": 1, "kaypro2": 1, "ibm-pc": 1, "ibm5051": 1, "5051": 1, "ibmpc": 1, "wy60-PC": 1, "wyse60-PC": 1,
	// Miscellaneous microcomputer consoles
	"mai": 1, "basic4": 1, "basis": 1, "luna": 1, "luna68k": 1, "megatek": 1, "xerox820": 1, "x820": 1,
	// North Star
	"northstar": 1,
	// Osborne
	"osborne-w": 1, "osborne1-w": 1, "osborne": 1, "osborne1": 1, "osexec": 1,
	// Radio Shack/Tandy
	"coco3": 1, "os9LII": 1, "trs2": 1, "trsII": 1, "trs80II": 1, "trs16": 1,
	// Videotex and teletext
	"m2-nam": 1, "minitel": 1, "minitel-2": 1, "minitel-2-nam": 1, "minitel1": 1, "minitel1b": 1, "minitel1b-80": 1, "minitel1-nb": 1, "minitel1b-nb": 1, "minitel2-80": 1, "minitel12-80": 1, "screen.minitel1": 1, "screen.minitel1b": 1, "screen.minitel1b-80": 1, "screen.minitel2-80": 1, "screen.minitel12-80": 1, "screen.minitel1-nb": 1, "screen.minitel1b-nb": 1, "linux-m1": 1, "linux-m1b": 1, "linux-m2": 1, "linux-s": 1, "screen.linux-m1": 1, "screen.linux-m1b": 1, "screen.linux-m2": 1, "putty-m1": 1, "putty-m1b": 1, "putty-m2": 1, "putty+screen": 1, "putty-screen": 1, "screen.putty-m1": 1, "screen.putty-m1b": 1, "screen.putty-m2": 1, "viewdata": 1, "viewdata-o": 1, "viewdata-rv": 1,
	// OBSOLETE UNIX CONSOLES
	// AT&T consoles
	"att6386": 1, "at386": 1, "386at": 1, "pc6300plus": 1, "att7300": 1, "unixpc": 1, "pc7300": 1, "3b1": 1, "s4": 1,
	// Apollo consoles
	"apollo": 1, "apollo+vt132": 1, "apollo_15P": 1, "apollo_19L": 1, "apollo_color": 1,
	// Convergent Technology
	"aws": 1, "awsc": 1,
	// DEC consoles
	"qdss": 1, "qdcons": 1,
	// Fortune Systems consoles
	"fos": 1, "fortune": 1,
	// Masscomp consoles
	"masscomp": 1, "masscomp1": 1, "masscomp2": 1,
	// OSF Unix
	"pmcons": 1, "pmconsole": 1,
	// Other consoles
	"pcix": 1, "ibmpcx": 1, "xenix": 1, "ibmx": 1,
	// OBSOLETE VDT TYPES
	// Amtek Business Machines
	"abm80": 1,
	// Bell Labs blit terminals
	"blit": 1, "jerq": 1, "cbblit": 1, "fixterm": 1, "oblit": 1, "ojerq": 1,
	// Bolt, Beranek & Newman (bbn)
	"bitgraph": 1, "bg2.0nv": 1, "bg3.10nv": 1, "bg2.0rv": 1, "bg3.10rv": 1, "bg2.0": 1, "bg3.10": 1, "bg1.25rv": 1, "bg1.25nv": 1, "bg1.25": 1,
	// Bull (bq, dku, vip)
	"tws-generic": 1, "dku7102": 1, "tws2102-sna": 1, "dku7102-sna": 1, "tws2103": 1, "xdku": 1, "tws2103-sna": 1, "dku7103-sna": 1, "dku7102-old": 1, "dku7202": 1, "bq300": 1, "bq300-rv": 1, "bq300-w": 1, "bq300-w-rv": 1, "bq300-8": 1, "bq300-8rv": 1, "bq300-8w": 1, "bq300-w-8rv": 1, "bq300-pc": 1, "bq300-pc-rv": 1, "bq300-pc-w": 1, "bq300-pc-w-rv": 1, "bq300-8-pc": 1, "Q306-8-pc": 1, "bq300-8-pc-rv": 1, "bq300-8-pc-w": 1, "bq300-8-pc-w-rv": 1, "vip": 1, "vip-w": 1, "vip7800-w": 1, "Q310-vip-w": 1, "Q310-vip-w-am": 1, "vip-H": 1, "vip7800-H": 1, "Q310-vip-H": 1, "Q310-vip-H-am": 1, "vip-Hw": 1, "vip7800-Hw": 1, "Q310-vip-Hw": 1,
	// Chromatics
	"cg7900": 1, "chromatics": 1,
	// Computer Automation
	"ca22851": 1,
	// Cybernex
	"cyb83": 1, "xl83": 1, "cyb110": 1, "mdl110": 1,
	// DEC terminals (Obsolete types: DECwriter and VT40/42/50)
	"vt52+keypad": 1, "gt40": 1, "gt42": 1, "vt50": 1, "vt50h": 1, "vt61": 1, "vt-61": 1, "vt61.5": 1, "gigi": 1, "vk100": 1, "pro350": 1, "decpro": 1, "dw1": 1, "dw2": 1, "decwriter": 1, "dw": 1, "dw3": 1, "la120": 1, "dw4": 1, "ln03": 1, "ln03-w": 1,
	// Datapoint
	"dp3360": 1, "datapoint": 1, "dp8242": 1,
	// Delta Data (dd)
	"delta": 1, "dd5000": 1,
	// Digital Data Research (ddr)
	"ddr": 1, "rebus3180": 1, "ddr3180": 1,
	// Evans & Sutherland
	"ps300": 1,
	// General Electric (ge)
	"terminet1200": 1, "terminet300": 1, "tn1200": 1, "tn300": 1, "terminet": 1,
	// Heathkit/Zenith
	"h19-a": 1, "h19a": 1, "heath-ansi": 1, "heathkit-a": 1, "h19-bs": 1, "h19-us": 1, "h19us": 1, "h19-smul": 1, "h19": 1, "heath": 1, "h19-b": 1, "heathkit": 1, "heath-19": 1, "z19": 1, "zenith": 1, "h19-u": 1, "h19-g": 1, "h19g": 1, "alto-h19": 1, "altoh19": 1, "altoheath": 1, "alto-heath": 1, "z29": 1, "zenith29": 1, "z29b": 1, "z29a": 1, "z29a-kc-bc": 1, "h29a-kc-bc": 1, "z29a-kc-uc": 1, "h29a-kc-uc": 1, "z29a-nkc-bc": 1, "h29a-nkc-bc": 1, "z29a-nkc-uc": 1, "h29a-nkc-uc": 1, "z39-a": 1, "z39a": 1, "zenith39-a": 1, "zenith39-ansi": 1, "z100": 1, "h100": 1, "z110": 1, "z-100": 1, "h-100": 1, "z100bw": 1, "h100bw": 1, "z110bw": 1, "z-100bw": 1, "h-100bw": 1, "p19": 1, "ztx": 1, "ztx11": 1, "zt-1": 1, "htx11": 1, "ztx-1-a": 1,
	// IMS International (ims)
	"ims950-b": 1, "ims950": 1, "ims950-rv": 1, "ims-ansi": 1, "ultima2": 1, "ultimaII": 1,
	// Intertec Data Systems
	"superbrain": 1, "intertube": 1, "intertec": 1, "intertube2": 1,
	// Ithaca Intersystems
	"graphos": 1, "graphos-30": 1,
	// Modgraph
	"modgraph": 1, "mod24": 1, "modgraph2": 1, "modgraph48": 1, "mod": 1,
	// Morrow Designs
	"mt70": 1, "mt-70": 1,
	// Motorola
	"ex155": 1,
	// Omron
	"omron": 1,
	// RCA
	"rca": 1,
	// Ramtek
	"rt6221": 1, "rt6221-w": 1,
	// Selanar
	"hirez100": 1, "hirez100-w": 1,
	// Signetics
	"vsc": 1,
	// Soroc
	"soroc120": 1, "iq120": 1, "soroc": 1, "soroc140": 1, "iq140": 1,
	// Southwest Technical Products
	"swtp": 1, "ct82": 1,
	// Synertek
	"synertek": 1, "ktm": 1, "synertek380": 1,
	// Tab Office Products
	"tab132": 1, "tab": 1, "tab132-15": 1, "tab132-w": 1, "tab132-rv": 1, "tab132-w-rv": 1,
	// Teleray
	"t3700": 1, "t3800": 1, "t1061": 1, "teleray": 1, "t1061f": 1, "t10": 1, "t16": 1,
	// Texas Instruments (ti)
	"ti700": 1, "ti733": 1, "ti735": 1, "ti745": 1, "ti800": 1, "ti703": 1, "ti707": 1, "ti703-w": 1, "ti707-w": 1, "ti916": 1, "ti916-220-7": 1, "ti916-8": 1, "ti916-220-8": 1, "ti916-132": 1, "ti916-8-132": 1, "ti924": 1, "ti924-8": 1, "ti924w": 1, "ti924-8w": 1, "ti931": 1, "ti926": 1, "ti926-8": 1, "ti_ansi": 1, "ti928": 1, "ti928-8": 1,
	// Zentec (zen)
	"zen30": 1, "z30": 1, "zen50": 1, "z50": 1, "cci": 1, "cci1": 1, "z8001": 1, "zen8001": 1,
	// OLDER TERMINAL TYPES
	// AT&T (att, tty)
	"att2300": 1, "sv80": 1, "att2350": 1, "att5410v1": 1, "att4410v1": 1, "tty5410v1": 1, "att4410v1-w": 1, "att5410v1-w": 1, "tty5410v1-w": 1, "att4410": 1, "att5410": 1, "tty5410": 1, "att5410-w": 1, "att4410-w": 1, "4410-w": 1, "tty5410-w": 1, "5410-w": 1, "v5410": 1, "att4415": 1, "tty5420": 1, "att5420": 1, "att4415-w": 1, "tty5420-w": 1, "att5420-w": 1, "att4415-rv": 1, "tty5420-rv": 1, "att5420-rv": 1, "att4415-w-rv": 1, "tty5420-w-rv": 1, "att5420-w-rv": 1, "att4415+nl": 1, "tty5420+nl": 1, "att5420+nl": 1, "att4415-nl": 1, "tty5420-nl": 1, "att5420-nl": 1, "att4415-rv-nl": 1, "tty5420-rv-nl": 1, "att5420-rv-nl": 1, "att4415-w-nl": 1, "tty5420-w-nl": 1, "att5420-w-nl": 1, "att4415-w-rv-n": 1, "tty5420-w-rv-n": 1, "att5420-w-rv-n": 1, "att5420_2": 1, "att5420_2-w": 1, "att4418": 1, "att5418": 1, "att4418-w": 1, "att5418-w": 1, "att4420": 1, "tty4420": 1, "att4424": 1, "tty4424": 1, "att4424-1": 1, "tty4424-1": 1, "att4424m": 1, "tty4424m": 1, "att5425": 1, "tty5425": 1, "att4425": 1, "att5425-nl": 1, "tty5425-nl": 1, "att4425-nl": 1, "att5425-w": 1, "att4425-w": 1, "tty5425-w": 1, "att4426": 1, "tty4426": 1, "att510a": 1, "bct510a": 1, "att510d": 1, "bct510d": 1, "att500": 1, "att513": 1, "att5310": 1, "att5320": 1, "att5620-1": 1, "tty5620-1": 1, "dmd1": 1, "att5620": 1, "dmd": 1, "tty5620": 1, "ttydmd": 1, "5620": 1, "att5620-24": 1, "tty5620-24": 1, "dmd-24": 1, "att5620-34": 1, "tty5620-34": 1, "dmd-34": 1, "att5620-s": 1, "tty5620-s": 1, "layer": 1, "vitty": 1, "att605": 1, "att605-pc": 1, "att605-w": 1, "att610": 1, "att610-w": 1, "att610-103k": 1, "att610-103k-w": 1, "att615": 1, "att615-w": 1, "att615-103k": 1, "att615-103k-w": 1, "att620": 1, "att620-w": 1, "att620-103k": 1, "att620-103k-w": 1, "att630": 1, "att630-24": 1, "5630-24": 1, "5630DMD-24": 1, "630MTG-24": 1, "att700": 1, "att730": 1, "att730-41": 1, "730MTG-41": 1, "att730-24": 1, "730MTG-24": 1, "att730r": 1, "730MTGr": 1, "att730r-41": 1, "730MTG-41r": 1, "att730r-24": 1, "730MTGr-24": 1, "att505": 1, "pt505": 1, "att5430": 1, "gs5430": 1, "att505-24": 1, "pt505-24": 1, "gs5430-24": 1, "att505-22": 1, "pt505-22": 1, "gs5430-22": 1,
	// Ampex (Dialogue)
	"ampex80": 1, "a80": 1, "d80": 1, "dialogue": 1, "dialogue80": 1, "ampex175": 1, "ampex175-b": 1, "ampex210": 1, "a210": 1, "ampex219": 1, "ampex-219": 1, "amp219": 1, "ampex219w": 1, "ampex-219w": 1, "amp219w": 1, "ampex232": 1, "ampex-232": 1, "ampex232w": 1,
	// Ann Arbor (aa)
	"annarbor4080": 1, "aa4080": 1, "aas1901": 1, "aaa+unk": 1, "aaa-unk": 1, "aaa+rv": 1, "aaa+dec": 1, "aaa-18": 1, "aaa-18-rv": 1, "aaa-20": 1, "aaa-22": 1, "aaa-24": 1, "aaa-24-rv": 1, "aaa-26": 1, "aaa-28": 1, "aaa-30-s": 1, "aaa-s": 1, "aaa-30-s-rv": 1, "aaa-s-rv": 1, "aaa-s-ctxt": 1, "aaa-30-s-ctxt": 1, "aaa-s-rv-ctxt": 1, "aaa-30-s-rv-ct": 1, "aaa": 1, "aaa-30": 1, "ambas": 1, "ambassador": 1, "aaa-30-rv": 1, "aaa-rv": 1, "aaa-30-ctxt": 1, "aaa-ctxt": 1, "aaa-30-rv-ctxt": 1, "aaa-rv-ctxt": 1, "aaa-36": 1, "aaa-36-rv": 1, "aaa-40": 1, "aaa-40-rv": 1, "aaa-48": 1, "aaa-48-rv": 1, "aaa-60-s": 1, "aaa-60-s-rv": 1, "aaa-60-dec-rv": 1, "aaa-60": 1, "aaa-60-rv": 1, "aaa-db": 1, "guru": 1, "guru-33": 1, "guru+unk": 1, "guru+rv": 1, "guru-rv": 1, "guru-33-rv": 1, "guru+s": 1, "guru-nctxt": 1, "guru-s": 1, "guru-33-s": 1, "guru-24": 1, "guru-44": 1, "guru-44-s": 1, "guru-76": 1, "guru-76-s": 1, "guru-76-lp": 1, "guru-lp": 1, "guru-76-w": 1, "guru-76-w-s": 1, "guru-76-wm": 1, "aaa-rv-unk": 1,
	// Applied Digital Data Systems (adds)
	"regent": 1, "regent100": 1, "regent20": 1, "regent25": 1, "regent40": 1, "regent40+": 1, "regent60": 1, "regent200": 1, "adds200": 1, "viewpoint": 1, "addsviewpoint": 1, "screwpoint": 1, "vp3a+": 1, "viewpoint3a+": 1, "vp60": 1, "viewpoint60": 1, "addsvp60": 1, "vp90": 1, "viewpoint90": 1, "adds980": 1, "a980": 1,
	// Beehive Medical Electronics
	"beehive": 1, "bee": 1, "beehive3": 1, "bh3m": 1, "beehiveIIIm": 1, "beehive4": 1, "bh4": 1, "microb": 1, "microbee": 1, "ha8675": 1, "ha8686": 1,
	// C. Itoh Electronics
	"cit80": 1, "cit-80": 1, "cit101": 1, "citc": 1, "cit101e": 1, "cit101e-rv": 1, "cit101e-n": 1, "cit101e-132": 1, "cit101e-n132": 1, "cit500": 1, "citoh": 1, "ci8510": 1, "8510": 1, "citoh-pica": 1, "citoh-elite": 1, "citoh-comp": 1, "citoh-prop": 1, "citoh-ps": 1, "ips": 1, "citoh-6lpi": 1, "citoh-8lpi": 1,
	// Contel Business Systems.
	"contel300": 1, "contel320": 1, "c300": 1, "contel301": 1, "contel321": 1, "c301": 1, "c321": 1,
	// Control Data (cdc)
	"cdc456": 1, "cdc721": 1, "cdc721ll": 1, "cdc752": 1, "cdc756": 1, "cdc721-esc": 1,
	// Data General (dg)
	"dgkeys+8b": 1, "dgkeys+7b": 1, "dgkeys+11": 1, "dgkeys+15": 1, "dgunix+fixed": 1, "dg+fixed": 1, "dg+color8": 1, "dg+color": 1, "dgmode+color8": 1, "dgmode+color": 1, "dgunix+ccc": 1, "dg+ccc": 1, "dg-generic": 1, "dg200": 1, "dg210": 1, "dg-ansi": 1, "dg211": 1, "dg450": 1, "dg6134": 1, "dg460-ansi": 1, "dg6053-old": 1, "dg100": 1, "dg6053": 1, "6053": 1, "6053-dg": 1, "dg605x": 1, "605x": 1, "605x-dg": 1, "d2": 1, "d2-dg": 1, "d200": 1, "d200-dg": 1, "d210": 1, "d214": 1, "d210-dg": 1, "d214-dg": 1, "d211": 1, "d215": 1, "d211-7b": 1, "d215-7b": 1, "d211-dg": 1, "d215-dg": 1, "d216-dg": 1, "d216e-dg": 1, "d216+dg": 1, "d216e+dg": 1, "d217-dg": 1, "d216-unix": 1, "d216e-unix": 1, "d216+": 1, "d216e+": 1, "d216-unix-25": 1, "d216+25": 1, "d217-unix": 1, "d217-unix-25": 1, "d220": 1, "d220-7b": 1, "d220-dg": 1, "d230c": 1, "d230": 1, "d230c-dg": 1, "d230-dg": 1, "d400": 1, "d400-dg": 1, "d450": 1, "d450-dg": 1, "d410": 1, "d411": 1, "d460": 1, "d461": 1, "d410-7b": 1, "d411-7b": 1, "d460-7b": 1, "d461-7b": 1, "d410-dg": 1, "d460-dg": 1, "d411-dg": 1, "d461-dg": 1, "d410-w": 1, "d411-w": 1, "d460-w": 1, "d461-w": 1, "d410-7b-w": 1, "d411-7b-w": 1, "d460-7b-w": 1, "d461-7b-w": 1, "d412-dg": 1, "d462-dg": 1, "d462e-dg": 1, "d412+dg": 1, "d462+dg": 1, "d413-dg": 1, "d463-dg": 1, "d412-unix": 1, "d462-unix": 1, "d412+": 1, "d462+": 1, "d412-unix-w": 1, "d462-unix-w": 1, "d412+w": 1, "d462+w": 1, "d412-unix-25": 1, "d462-unix-25": 1, "d412+25": 1, "d462+25": 1, "d412-unix-s": 1, "d462-unix-s": 1, "d412+s": 1, "d462+s": 1, "d412-unix-sr": 1, "d462-unix-sr": 1, "d412+sr": 1, "d462+sr": 1, "d413-unix": 1, "d463-unix": 1, "d413-unix-w": 1, "d463-unix-w": 1, "d413-unix-25": 1, "d463-unix-25": 1, "d413-unix-s": 1, "d463-unix-s": 1, "d413-unix-sr": 1, "d463-unix-sr": 1, "d414-unix": 1, "d464-unix": 1, "d414-unix-w": 1, "d464-unix-w": 1, "d414-unix-25": 1, "d464-unix-25": 1, "d414-unix-s": 1, "d464-unix-s": 1, "d414-unix-sr": 1, "d464-unix-sr": 1, "d430c-dg": 1, "d430-dg": 1, "d430c-dg-ccc": 1, "d430-dg-ccc": 1, "d430c-unix": 1, "d430-unix": 1, "d430c-unix-w": 1, "d430-unix-w": 1, "d430c-unix-25": 1, "d430-unix-25": 1, "d430c-unix-s": 1, "d430-unix-s": 1, "d430c-unix-sr": 1, "d430-unix-sr": 1, "d430c-unix-ccc": 1, "d430-unix-ccc": 1, "d430c-unix-w-ccc": 1, "d430-unix-w-ccc": 1, "d430c-unix-25-ccc": 1, "d430-unix-25-ccc": 1, "d430c-unix-s-ccc": 1, "d430-unix-s-ccc": 1, "d430c-unix-sr-ccc": 1, "d430-unix-sr-ccc": 1, "d470c": 1, "d470": 1, "d470c-7b": 1, "d470-7b": 1, "d470c-dg": 1, "d470-dg": 1, "d555": 1, "d555-7b": 1, "d555-w": 1, "d555-7b-w": 1, "d555-dg": 1, "d577": 1, "d577-7b": 1, "d577-w": 1, "d577-7b-w": 1, "d577-dg": 1, "d578-dg": 1, "d578": 1, "d578-7b": 1,
	// Datamedia (dm)
	"cs10": 1, "colorscan": 1, "cs10-w": 1, "dm1520": 1, "dm1521": 1, "dm2500": 1, "datamedia2500": 1, "dmchat": 1, "dm3025": 1, "dm3045": 1, "dm80": 1, "dmdt80": 1, "dt80": 1, "dm80w": 1, "dmdt80w": 1, "dt80w": 1, "dt80-sas": 1, "excel62": 1, "excel64": 1, "excel62-w": 1, "excel64-w": 1, "excel62-rv": 1, "excel64-rv": 1,
	// Falco
	"falco": 1, "ts1": 1, "ts-1": 1, "falco-p": 1, "ts1p": 1, "ts-1p": 1, "ts100": 1, "ts100-sp": 1, "ts100-ctxt": 1,
	// Florida Computer Graphics
	"beacon": 1,
	// Fluke
	"f1720": 1, "f1720a": 1,
	// Getronics
	"visa50": 1,
	// GraphOn (go)
	"go140": 1, "go140w": 1, "go225": 1, "go-225": 1,
	// Harris (Beehive)
	"sb1": 1, "sbi": 1, "superbee": 1, "superbee-xsb": 1, "superbeeic": 1, "sb2": 1, "sb3": 1,
	// Hazeltine
	"hz1000": 1, "hz1420": 1, "hz1500": 1, "hz1510": 1, "hz1520": 1, "hz1520-noesc": 1, "hz1552": 1, "hz1552-rv": 1, "hz2000": 1, "esprit": 1, "esprit-am": 1, "hmod1": 1, "hazel": 1, "exec80": 1, "h80": 1, "he80": 1,
	// Human Designed Systems (Concept)
	"c108": 1, "concept108": 1, "c108-8p": 1, "concept108-8p": 1, "c108-4p": 1, "concept108-4p": 1, "c108-rv": 1, "c108-rv-8p": 1, "c108-rv-4p": 1, "concept108rv4p": 1, "c108-w": 1, "c108-w-8p": 1, "concept108-w-8": 1, "concept108-w8p": 1, "c100": 1, "concept100": 1, "concept": 1, "c104": 1, "c100-4p": 1, "c100-rv": 1, "c100-rv-4p": 1, "concept100-rv": 1, "oc100": 1, "oconcept": 1, "c100-1p": 1, "hds200": 1, "avt-ns": 1, "avt-rv-ns": 1, "avt-w-ns": 1, "avt-w-rv-ns": 1, "avt+s": 1, "avt": 1, "avt-s": 1, "concept-avt": 1, "avt-rv": 1, "avt-rv-s": 1, "avt-w": 1, "avt-w-s": 1, "avt-w-rv": 1, "avt-w-rv-s": 1,
	// IBM
	"ibm327x": 1, "ibm3101": 1, "i3101": 1, "ibm3151": 1, "ibm3161": 1, "ibm3163": 1, "wy60-316X": 1, "wyse60-316X": 1, "ibm3161-C": 1, "ibm3162": 1, "ibm3164": 1, "i3164": 1, "ibm5151": 1, "wy60-AT": 1, "wyse60-AT": 1, "ibmaed": 1, "ibm-apl": 1, "apl": 1, "ibmmono": 1, "ibmega": 1, "ibm+color": 1, "ibm+16color": 1, "ibm5154": 1, "ibmega-c": 1, "ibm5154-c": 1, "ibmvga-c": 1, "ibmvga": 1, "rtpc": 1, "ibmapa16": 1, "ibm6155": 1, "ibmapa8c": 1, "ibmapa8": 1, "ibmapa8c-c": 1, "ibm6154-c": 1, "ibm6154": 1, "ibm6153": 1, "ibm6153-90": 1, "ibm6153-40": 1, "ibm8512": 1, "ibm8513": 1, "hft-c": 1, "hft-c-old": 1, "hft-old": 1, "ibm-system1": 1, "system1": 1, "lft": 1, "lft-pc850": 1, "LFT-PC850": 1, "ibm5081": 1, "hft": 1, "ibm5081-c": 1, "ibmmpel-c": 1, "ibm8503": 1, "ibm8507": 1, "ibm8604": 1, "ibm8514": 1, "ibm8514-c": 1, "aixterm": 1, "aixterm+sl": 1, "aixterm-m": 1, "aixterm-m-old": 1, "jaixterm": 1, "jaixterm-m": 1, "aixterm-16color": 1,
	// Infoton/General Terminal Corp.
	"i100": 1, "gt100": 1, "gt100a": 1, "i400": 1, "addrinfo": 1, "infoton2": 1, "infoton": 1, "icl6404": 1, "kds7372": 1, "icl6402": 1, "kds6402": 1, "icl6404-w": 1, "kds7372-w": 1,
	// Interactive Systems Corp
	"intext": 1, "intext2": 1, "intextii": 1,
	// Kimtron (abm, kt)
	"abm85": 1, "abm85h": 1, "abm85e": 1, "abm85h-old": 1, "oabm85h": 1, "o85h": 1, "kt7": 1, "kt7ix": 1,
	// Liberty Electronics (Freedom)
	"f100": 1, "freedom": 1, "freedom100": 1, "f100-rv": 1, "freedom-rv": 1, "f110": 1, "freedom110": 1, "f110-14": 1, "f110-w": 1, "f110-14w": 1, "f200": 1, "freedom200": 1, "f200-w": 1, "f200vi": 1, "f200vi-w": 1,
	// Microdata/MDIS
	"prism2": 1, "prism4": 1, "p4": 1, "P4": 1, "prism5": 1, "p5": 1, "P5": 1, "prism7": 1, "p7": 1, "P7": 1, "prism8": 1, "p8": 1, "P8": 1, "prism8-w": 1, "p8-w": 1, "P8-W": 1, "prism9": 1, "p9": 1, "P9": 1, "prism9-w": 1, "p9-w": 1, "P9-W": 1, "prism9-8": 1, "p9-8": 1, "P9-8": 1, "prism9-8-w": 1, "p9-8-w": 1, "P9-8-W": 1, "prism12": 1, "p12": 1, "P12": 1, "prism12-w": 1, "p12-w": 1, "P12-W": 1, "prism12-m": 1, "p12-m": 1, "P12-M": 1, "prism12-m-w": 1, "p12-m-w": 1, "P12-M-W": 1, "prism14": 1, "p14": 1, "P14": 1, "prism14-w": 1, "p14-w": 1, "P14-W": 1, "prism14-m": 1, "p14-m": 1, "P14-M": 1, "prism14-m-w": 1, "p14-m-w": 1, "P14-M-W": 1, "p8gl": 1, "prism8gl": 1,
	// Microterm (act, mime)
	"act4": 1, "microterm": 1, "act5": 1, "microterm5": 1, "mime-fb": 1, "mime-hb": 1, "mime": 1, "mime1": 1, "mime2": 1, "mimei": 1, "mimeii": 1, "mime2a-s": 1, "mime2a": 1, "mime2a-v": 1, "mime3a": 1, "mime3ax": 1, "mime-3ax": 1, "mime314": 1, "mm314": 1, "mm340": 1, "mime340": 1, "mt4520-rv": 1, "ergo4000": 1,
	// NCR
	"ncr260intan": 1, "ncr260intwan": 1, "ncr260intpp": 1, "ncr260intwpp": 1, "ncr260vppp": 1, "ncr260vp+sl": 1, "ncr260vpwpp": 1, "ncr260vt100an": 1, "ncr260vt+sl": 1, "ncr260vt100wan": 1, "ncr260vt100pp": 1, "ncr260vt100wpp": 1, "ncr260vt200an": 1, "ncr260vt200wan": 1, "ncr260vt200pp": 1, "ncr260vt200wpp": 1, "ncr260vt300an": 1, "ncr260vt300wan": 1, "ncr260vt300pp": 1, "ncr260vt300wpp": 1, "NCR260VT300WPP": 1, "ncr260wy325pp": 1, "ncr260wy325wpp": 1, "ncr260wy350pp": 1, "ncr260wy350wpp": 1, "ncr260wy50+pp": 1, "ncr260wy50+wpp": 1, "ncr260wy60pp": 1, "ncr260wy60wpp": 1, "ncr160vppp": 1, "ncr160vpwpp": 1, "ncr160vt100an": 1, "ncr160vt100pp": 1, "ncr160vt100wan": 1, "ncr160vt100wpp": 1, "ncr160vt200an": 1, "ncr160vt200pp": 1, "ncr160vt200wan": 1, "ncr160vt200wpp": 1, "ncr160vt300an": 1, "ncr160vt300pp": 1, "ncr160vt300wan": 1, "ncr160vt300wpp": 1, "ncr160wy50+pp": 1, "ncr160wy50+wpp": 1, "ncr160wy60pp": 1, "ncr160wy60wpp": 1, "ncrvt100an": 1, "ncrvt100pp": 1, "ncrvt100wan": 1, "NCRVT100WPP": 1, "ncrvt100wpp": 1, "ncr7900i": 1, "ncr7900": 1, "n7900": 1, "ncr7900iv": 1, "ncr7901": 1, "ndr9500": 1, "nd9500": 1, "ndr9500-nl": 1, "ndr9500-25": 1, "ndr9500-25-nl": 1, "ndr9500-mc": 1, "ndr9500-25-mc": 1, "ndr9500-mc-nl": 1, "ndr9500-25-mc-nl": 1,
	// Perkin-Elmer (Owl)
	"bantam": 1, "pe550": 1, "pe6100": 1, "fox": 1, "pe1100": 1, "owl": 1, "pe1200": 1, "pe1251": 1, "pe6300": 1, "pe6312": 1, "pe7000m": 1, "pe7000c": 1,
	// Sperry Univac
	"uts30": 1,
	// Tandem
	"tandem6510": 1, "tandem653": 1, "t653x": 1,
	// Tandy/Radio Shack
	"dmterm": 1, "dt100": 1, "dt-100": 1, "dt100w": 1, "dt-100w": 1, "dt110": 1, "pt210": 1,
	// Tektronix (tek)
	"tek": 1, "tek4012": 1, "tek4013": 1, "tek4014": 1, "tek4015": 1, "tek4014-sm": 1, "tek4015-sm": 1, "tek4023": 1, "tek4024": 1, "tek4025": 1, "tek4027": 1, "tek4025-17": 1, "tek4025-17-ws": 1, "tek4025-ex": 1, "tek4027-ex": 1, "tek4025a": 1, "tek4025-cr": 1, "tek4025ex": 1, "4025ex": 1, "4027ex": 1, "tek4105": 1, "tek4105-30": 1, "tek4105a": 1, "tek4106brl": 1, "tek4107brl": 1, "tek4109brl": 1, "tek4107": 1, "tek4109": 1, "tek4207-s": 1, "otek4112": 1, "o4112-nd": 1, "otek4113": 1, "otek4114": 1, "tek4112": 1, "tek4114": 1, "tek4112-nd": 1, "tek4112-5": 1, "tek4113": 1, "tek4113-34": 1, "tek4113-nd": 1, "otek4115": 1, "tek4115": 1, "tek4125": 1, "tek4207": 1, "tek4404": 1, "ct8500": 1, "tek4205": 1,
	// Teletype (tty)
	"tty33": 1, "tty35": 1, "tty37": 1, "tty40": 1, "ds40": 1, "ds40-2": 1, "dataspeed40": 1, "tty43": 1,
	// Tymshare
	"scanset": 1, "sc410": 1, "sc415": 1,
	// Volker-Craig (vc)
	"vc303": 1, "vc103": 1, "vc203": 1, "vc303a": 1, "vc403a": 1, "vc404": 1, "vc404-s": 1, "vc414": 1, "vc414h": 1, "vc415": 1,
	// OTHER OBSOLETE TYPES
	// Daisy wheel printers
	"diablo1620": 1, "diablo1720": 1, "diablo450": 1, "ipsi": 1, "diablo1620-m8": 1, "diablo1640-m8": 1, "diablo1640": 1, "diablo1730": 1, "diablo1740": 1, "diablo630": 1, "x1700": 1, "diablo": 1, "xerox": 1, "diablo1640-lm": 1, "diablo-lm": 1, "xerox-lm": 1, "diablo1740-lm": 1, "630-lm": 1, "1730-lm": 1, "x1700-lm": 1, "dtc382": 1, "dtc300s": 1, "gsi": 1, "aj830": 1, "aj832": 1, "aj": 1, "aj510": 1, "nec5520": 1, "nec": 1, "spinwriter": 1, "qume5": 1, "qume": 1, "xerox1720": 1, "x1720": 1, "x1750": 1,
	// Miscellaneous obsolete terminals, manufacturers unknown
	"cad68-3": 1, "cgc3": 1, "cad68-2": 1, "cgc2": 1, "cops10": 1, "cops": 1, "cops-10": 1, "d132": 1, "datagraphix": 1, "d800": 1, "digilog": 1, "dwk": 1, "dwk-vt": 1, "env230": 1, "envision230": 1, "ep48": 1, "ep4080": 1, "ep40": 1, "ep4000": 1, "ifmr": 1, "opus3n1+": 1, "teletec": 1, "v3220": 1,
	// Obsolete non-ANSI software emulations
	"ctrm": 1, "gs6300": 1, "emots": 1, "h19k": 1, "h19kermit": 1, "versaterm": 1, "xtalk": 1, "simterm": 1,
	// UNIX VIRTUAL TERMINALS, VIRTUAL CONSOLES, AND TELNET CLIENTS
	// NCSA Telnet
	"ncsa-m": 1, "ncsa-vt220-8": 1, "ncsa": 1, "ncsa-ns": 1, "ncsa-m-ns": 1, "ncsa-vt220": 1,
	// Pilot Pro Palm-Top
	"pilot": 1, "tgtelnet": 1, "elks-glasstty": 1, "elks-vt52": 1, "elks-ansi": 1, "elks": 1, "sibo": 1,
	// Screen
	"screen2": 1, "screen3": 1, "screen4": 1, "screen5": 1,
	// CB UNIX, early 80s
	"cbunix": 1, "vremote": 1, "pty": 1,
	// X TERMINAL EMULATORS
	// EMU
	"emu": 1, "emu-220": 1,
	// GNOME (VTE)
	"gnome-rh62": 1, "gnome-rh72": 1, "gnome-rh80": 1, "gnome-rh90": 1, "gnome-fc5": 1, "vte-2007": 1, "gnome-2007": 1, "vte-2008": 1, "gnome-2008": 1, "vte-2012": 1, "gnome-2012": 1, "gnome+pcfkeys": 1, "gnome": 1, "gnome-256color": 1, "vte-2014": 1, "vte-2017": 1, "vte-2018": 1,
	// HPTERM
	"hpterm": 1, "X-hpterm": 1, "hpterm-color": 1, "hpterm-color2": 1, "X-hpterm-color2": 1,
	// KDE
	"kvt": 1, "konsole-xf3x": 1, "konsole-xf4x": 1,
	// KTERM
	"kterm": 1, "kterm-color": 1, "kterm-co": 1,
	// MGR
	"mgr": 1, "mgr-sun": 1, "mgr-linux": 1,
	// MLTERM
	"mlterm3": 1, "mlterm2": 1,
	// MTERM
	"mterm-ansi": 1, "mterm": 1, "mouse-sun": 1, "decansi": 1,
	// MVTERM
	"mvterm": 1, "vv100": 1,
	// Other GNOME
	"mgt": 1,
	// Other XTERM
	"xtermm": 1, "xtermc": 1, "xterm-pcolor": 1, "color_xterm": 1, "cx": 1, "cx100": 1,
	// RXVT
	"rxvt-basic": 1, "rxvt+pcfkeys": 1, "rxvt": 1, "rxvt-color": 1, "rxvt-256color": 1, "rxvt-88color": 1, "rxvt-xpm": 1, "rxvt-cygwin": 1, "rxvt-cygwin-native": 1, "rxvt-16color": 1,
	// SIMPLETERM
	"st-0.8": 1, "st-0.7": 1, "st-0.6": 1, "simpleterm": 1, "old-st": 1,
	// TERMINOLOGY
	"terminology-0.6.1": 1, "terminology-1.0.0": 1, "terminology-1.8.1": 1,
	// VWM
	"vwmterm": 1,
	// XTERM
	"x10term": 1, "vs100-x10": 1, "x10term+sl": 1, "xterm-r5": 1, "xterm-r6": 1, "xterm-old": 1, "xterm-xf86-v32": 1, "xterm-xf86-v33": 1, "xterm-xf86-v333": 1, "xterm-xf86-v40": 1, "xterm-xf86-v43": 1, "xterm-xf86-v44": 1, "xterm-xfree86": 1, "xterm-new": 1,
	// XTERM Features
	"xterm-8bit": 1, "xterm-vt52": 1, "xterm-nic": 1, "xterm1": 1,
}
