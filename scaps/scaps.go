// Code generated by term.h.zsh; DO NOT EDIT.

// Package scaps contains a list of all terminfo capabilities.
package scaps

import "zgo.at/termfo/caps"

// CursesVersion is the version of curses this data was generated with, as [implementation]-[version].
const CursesVersion = `ncurses-6.3.20221105`

var (
	Bw       = caps.AutoLeftMargin
	Am       = caps.AutoRightMargin
	Xsb      = caps.NoEscCtlc
	Xhp      = caps.CeolStandoutGlitch
	Xenl     = caps.EatNewlineGlitch
	Eo       = caps.EraseOverstrike
	Gn       = caps.GenericType
	Hc       = caps.HardCopy
	Km       = caps.HasMetaKey
	Hs       = caps.HasStatusLine
	In       = caps.InsertNullGlitch
	Da       = caps.MemoryAbove
	Db       = caps.MemoryBelow
	Mir      = caps.MoveInsertMode
	Msgr     = caps.MoveStandoutMode
	Os       = caps.OverStrike
	Eslok    = caps.StatusLineEscOk
	Xt       = caps.DestTabsMagicSmso
	Hz       = caps.TildeGlitch
	Ul       = caps.TransparentUnderline
	Xon      = caps.XonXoff
	Nxon     = caps.NeedsXonXoff
	Mc5i     = caps.PrtrSilent
	Chts     = caps.HardCursor
	Nrrmc    = caps.NonRevRmcup
	Npc      = caps.NoPadChar
	Ndscr    = caps.NonDestScrollRegion
	Ccc      = caps.CanChange
	Bce      = caps.BackColorErase
	Hls      = caps.HueLightnessSaturation
	Xhpa     = caps.ColAddrGlitch
	Crxm     = caps.CrCancelsMicroMode
	Daisy    = caps.HasPrintWheel
	Xvpa     = caps.RowAddrGlitch
	Sam      = caps.SemiAutoRightMargin
	Cpix     = caps.CpiChangesRes
	Lpix     = caps.LpiChangesRes
	Cols     = caps.Columns
	It       = caps.InitTabs
	Lines    = caps.Lines
	Lm       = caps.LinesOfMemory
	Xmc      = caps.MagicCookieGlitch
	Pb       = caps.PaddingBaudRate
	Vt       = caps.VirtualTerminal
	Wsl      = caps.WidthStatusLine
	Nlab     = caps.NumLabels
	Lh       = caps.LabelHeight
	Lw       = caps.LabelWidth
	Ma       = caps.MaxAttributes
	Wnum     = caps.MaximumWindows
	Colors   = caps.MaxColors
	Pairs    = caps.MaxPairs
	Ncv      = caps.NoColorVideo
	Bufsz    = caps.BufferCapacity
	Spinv    = caps.DotVertSpacing
	Spinh    = caps.DotHorzSpacing
	Maddr    = caps.MaxMicroAddress
	Mjump    = caps.MaxMicroJump
	Mcs      = caps.MicroColSize
	Mls      = caps.MicroLineSize
	Npins    = caps.NumberOfPins
	Orc      = caps.OutputResChar
	Orl      = caps.OutputResLine
	Orhi     = caps.OutputResHorzInch
	Orvi     = caps.OutputResVertInch
	Cps      = caps.PrintRate
	Widcs    = caps.WideCharSize
	Btns     = caps.Buttons
	Bitwin   = caps.BitImageEntwining
	Bitype   = caps.BitImageType
	Cbt      = caps.BackTab
	Bel      = caps.Bell
	Cr       = caps.CarriageReturn
	Csr      = caps.ChangeScrollRegion
	Tbc      = caps.ClearAllTabs
	Clear    = caps.ClearScreen
	El       = caps.ClrEol
	Ed       = caps.ClrEos
	Hpa      = caps.ColumnAddress
	Cmdch    = caps.CommandCharacter
	Cup      = caps.CursorAddress
	Cud1     = caps.CursorDown
	Home     = caps.CursorHome
	Civis    = caps.CursorInvisible
	Cub1     = caps.CursorLeft
	Mrcup    = caps.CursorMemAddress
	Cnorm    = caps.CursorNormal
	Cuf1     = caps.CursorRight
	Ll       = caps.CursorToLl
	Cuu1     = caps.CursorUp
	Cvvis    = caps.CursorVisible
	Dch1     = caps.DeleteCharacter
	Dl1      = caps.DeleteLine
	Dsl      = caps.DisStatusLine
	Hd       = caps.DownHalfLine
	Smacs    = caps.EnterAltCharsetMode
	Blink    = caps.EnterBlinkMode
	Bold     = caps.EnterBoldMode
	Smcup    = caps.EnterCaMode
	Smdc     = caps.EnterDeleteMode
	Dim      = caps.EnterDimMode
	Smir     = caps.EnterInsertMode
	Invis    = caps.EnterSecureMode
	Prot     = caps.EnterProtectedMode
	Rev      = caps.EnterReverseMode
	Smso     = caps.EnterStandoutMode
	Smul     = caps.EnterUnderlineMode
	Ech      = caps.EraseChars
	Rmacs    = caps.ExitAltCharsetMode
	Sgr0     = caps.ExitAttributeMode
	Rmcup    = caps.ExitCaMode
	Rmdc     = caps.ExitDeleteMode
	Rmir     = caps.ExitInsertMode
	Rmso     = caps.ExitStandoutMode
	Rmul     = caps.ExitUnderlineMode
	Flash    = caps.FlashScreen
	Ff       = caps.FormFeed
	Fsl      = caps.FromStatusLine
	Is1      = caps.Init1string
	Is2      = caps.Init2string
	Is3      = caps.Init3string
	If       = caps.InitFile
	Ich1     = caps.InsertCharacter
	Il1      = caps.InsertLine
	Ip       = caps.InsertPadding
	Kbs      = caps.KeyBackspace
	Ktbc     = caps.KeyCatab
	Kclr     = caps.KeyClear
	Kctab    = caps.KeyCtab
	Kdch1    = caps.KeyDc
	Kdl1     = caps.KeyDl
	Kcud1    = caps.KeyDown
	Krmir    = caps.KeyEic
	Kel      = caps.KeyEol
	Ked      = caps.KeyEos
	Kf0      = caps.KeyF0
	Kf1      = caps.KeyF1
	Kf10     = caps.KeyF10
	Kf2      = caps.KeyF2
	Kf3      = caps.KeyF3
	Kf4      = caps.KeyF4
	Kf5      = caps.KeyF5
	Kf6      = caps.KeyF6
	Kf7      = caps.KeyF7
	Kf8      = caps.KeyF8
	Kf9      = caps.KeyF9
	Khome    = caps.KeyHome
	Kich1    = caps.KeyIc
	Kil1     = caps.KeyIl
	Kcub1    = caps.KeyLeft
	Kll      = caps.KeyLl
	Knp      = caps.KeyNpage
	Kpp      = caps.KeyPpage
	Kcuf1    = caps.KeyRight
	Kind     = caps.KeySf
	Kri      = caps.KeySr
	Khts     = caps.KeyStab
	Kcuu1    = caps.KeyUp
	Rmkx     = caps.KeypadLocal
	Smkx     = caps.KeypadXmit
	Lf0      = caps.LabF0
	Lf1      = caps.LabF1
	Lf10     = caps.LabF10
	Lf2      = caps.LabF2
	Lf3      = caps.LabF3
	Lf4      = caps.LabF4
	Lf5      = caps.LabF5
	Lf6      = caps.LabF6
	Lf7      = caps.LabF7
	Lf8      = caps.LabF8
	Lf9      = caps.LabF9
	Rmm      = caps.MetaOff
	Smm      = caps.MetaOn
	Nel      = caps.Newline
	Pad      = caps.PadChar
	Dch      = caps.ParmDch
	Dl       = caps.ParmDeleteLine
	Cud      = caps.ParmDownCursor
	Ich      = caps.ParmIch
	Indn     = caps.ParmIndex
	Il       = caps.ParmInsertLine
	Cub      = caps.ParmLeftCursor
	Cuf      = caps.ParmRightCursor
	Rin      = caps.ParmRindex
	Cuu      = caps.ParmUpCursor
	Pfkey    = caps.PkeyKey
	Pfloc    = caps.PkeyLocal
	Pfx      = caps.PkeyXmit
	Mc0      = caps.PrintScreen
	Mc4      = caps.PrtrOff
	Mc5      = caps.PrtrOn
	Rep      = caps.RepeatChar
	Rs1      = caps.Reset1string
	Rs2      = caps.Reset2string
	Rs3      = caps.Reset3string
	Rf       = caps.ResetFile
	Rc       = caps.RestoreCursor
	Vpa      = caps.RowAddress
	Sc       = caps.SaveCursor
	Ind      = caps.ScrollForward
	Ri       = caps.ScrollReverse
	Sgr      = caps.SetAttributes
	Hts      = caps.SetTab
	Wind     = caps.SetWindow
	Ht       = caps.Tab
	Tsl      = caps.ToStatusLine
	Uc       = caps.UnderlineChar
	Hu       = caps.UpHalfLine
	Iprog    = caps.InitProg
	Ka1      = caps.KeyA1
	Ka3      = caps.KeyA3
	Kb2      = caps.KeyB2
	Kc1      = caps.KeyC1
	Kc3      = caps.KeyC3
	Mc5p     = caps.PrtrNon
	Rmp      = caps.CharPadding
	Acsc     = caps.AcsChars
	Pln      = caps.PlabNorm
	Kcbt     = caps.KeyBtab
	Smxon    = caps.EnterXonMode
	Rmxon    = caps.ExitXonMode
	Smam     = caps.EnterAmMode
	Rmam     = caps.ExitAmMode
	Xonc     = caps.XonCharacter
	Xoffc    = caps.XoffCharacter
	Enacs    = caps.EnaAcs
	Smln     = caps.LabelOn
	Rmln     = caps.LabelOff
	Kbeg     = caps.KeyBeg
	Kcan     = caps.KeyCancel
	Kclo     = caps.KeyClose
	Kcmd     = caps.KeyCommand
	Kcpy     = caps.KeyCopy
	Kcrt     = caps.KeyCreate
	Kend     = caps.KeyEnd
	Kent     = caps.KeyEnter
	Kext     = caps.KeyExit
	Kfnd     = caps.KeyFind
	Khlp     = caps.KeyHelp
	Kmrk     = caps.KeyMark
	Kmsg     = caps.KeyMessage
	Kmov     = caps.KeyMove
	Knxt     = caps.KeyNext
	Kopn     = caps.KeyOpen
	Kopt     = caps.KeyOptions
	Kprv     = caps.KeyPrevious
	Kprt     = caps.KeyPrint
	Krdo     = caps.KeyRedo
	Kref     = caps.KeyReference
	Krfr     = caps.KeyRefresh
	Krpl     = caps.KeyReplace
	Krst     = caps.KeyRestart
	Kres     = caps.KeyResume
	Ksav     = caps.KeySave
	Kspd     = caps.KeySuspend
	Kund     = caps.KeyUndo
	KBEG     = caps.KeySbeg
	KCAN     = caps.KeyScancel
	KCMD     = caps.KeyScommand
	KCPY     = caps.KeyScopy
	KCRT     = caps.KeyScreate
	KDC      = caps.KeySdc
	KDL      = caps.KeySdl
	Kslt     = caps.KeySelect
	KEND     = caps.KeySend
	KEOL     = caps.KeySeol
	KEXT     = caps.KeySexit
	KFND     = caps.KeySfind
	KHLP     = caps.KeyShelp
	KHOM     = caps.KeyShome
	KIC      = caps.KeySic
	KLFT     = caps.KeySleft
	KMSG     = caps.KeySmessage
	KMOV     = caps.KeySmove
	KNXT     = caps.KeySnext
	KOPT     = caps.KeySoptions
	KPRV     = caps.KeySprevious
	KPRT     = caps.KeySprint
	KRDO     = caps.KeySredo
	KRPL     = caps.KeySreplace
	KRIT     = caps.KeySright
	KRES     = caps.KeySrsume
	KSAV     = caps.KeySsave
	KSPD     = caps.KeySsuspend
	KUND     = caps.KeySundo
	Rfi      = caps.ReqForInput
	Kf11     = caps.KeyF11
	Kf12     = caps.KeyF12
	Kf13     = caps.KeyF13
	Kf14     = caps.KeyF14
	Kf15     = caps.KeyF15
	Kf16     = caps.KeyF16
	Kf17     = caps.KeyF17
	Kf18     = caps.KeyF18
	Kf19     = caps.KeyF19
	Kf20     = caps.KeyF20
	Kf21     = caps.KeyF21
	Kf22     = caps.KeyF22
	Kf23     = caps.KeyF23
	Kf24     = caps.KeyF24
	Kf25     = caps.KeyF25
	Kf26     = caps.KeyF26
	Kf27     = caps.KeyF27
	Kf28     = caps.KeyF28
	Kf29     = caps.KeyF29
	Kf30     = caps.KeyF30
	Kf31     = caps.KeyF31
	Kf32     = caps.KeyF32
	Kf33     = caps.KeyF33
	Kf34     = caps.KeyF34
	Kf35     = caps.KeyF35
	Kf36     = caps.KeyF36
	Kf37     = caps.KeyF37
	Kf38     = caps.KeyF38
	Kf39     = caps.KeyF39
	Kf40     = caps.KeyF40
	Kf41     = caps.KeyF41
	Kf42     = caps.KeyF42
	Kf43     = caps.KeyF43
	Kf44     = caps.KeyF44
	Kf45     = caps.KeyF45
	Kf46     = caps.KeyF46
	Kf47     = caps.KeyF47
	Kf48     = caps.KeyF48
	Kf49     = caps.KeyF49
	Kf50     = caps.KeyF50
	Kf51     = caps.KeyF51
	Kf52     = caps.KeyF52
	Kf53     = caps.KeyF53
	Kf54     = caps.KeyF54
	Kf55     = caps.KeyF55
	Kf56     = caps.KeyF56
	Kf57     = caps.KeyF57
	Kf58     = caps.KeyF58
	Kf59     = caps.KeyF59
	Kf60     = caps.KeyF60
	Kf61     = caps.KeyF61
	Kf62     = caps.KeyF62
	Kf63     = caps.KeyF63
	El1      = caps.ClrBol
	Mgc      = caps.ClearMargins
	Smgl     = caps.SetLeftMargin
	Smgr     = caps.SetRightMargin
	Fln      = caps.LabelFormat
	Sclk     = caps.SetClock
	Dclk     = caps.DisplayClock
	Rmclk    = caps.RemoveClock
	Cwin     = caps.CreateWindow
	Wingo    = caps.GotoWindow
	Hup      = caps.Hangup
	Dial     = caps.DialPhone
	Qdial    = caps.QuickDial
	Tone     = caps.Tone
	Pulse    = caps.Pulse
	Hook     = caps.FlashHook
	Pause    = caps.FixedPause
	Wait     = caps.WaitTone
	U0       = caps.User0
	U1       = caps.User1
	U2       = caps.User2
	U3       = caps.User3
	U4       = caps.User4
	U5       = caps.User5
	U6       = caps.User6
	U7       = caps.User7
	U8       = caps.User8
	U9       = caps.User9
	Op       = caps.OrigPair
	Oc       = caps.OrigColors
	Initc    = caps.InitializeColor
	Initp    = caps.InitializePair
	Scp      = caps.SetColorPair
	Setf     = caps.SetForeground
	Setb     = caps.SetBackground
	Cpi      = caps.ChangeCharPitch
	Lpi      = caps.ChangeLinePitch
	Chr      = caps.ChangeResHorz
	Cvr      = caps.ChangeResVert
	Defc     = caps.DefineChar
	Swidm    = caps.EnterDoublewideMode
	Sdrfq    = caps.EnterDraftQuality
	Sitm     = caps.EnterItalicsMode
	Slm      = caps.EnterLeftwardMode
	Smicm    = caps.EnterMicroMode
	Snlq     = caps.EnterNearLetterQuality
	Snrmq    = caps.EnterNormalQuality
	Sshm     = caps.EnterShadowMode
	Ssubm    = caps.EnterSubscriptMode
	Ssupm    = caps.EnterSuperscriptMode
	Sum      = caps.EnterUpwardMode
	Rwidm    = caps.ExitDoublewideMode
	Ritm     = caps.ExitItalicsMode
	Rlm      = caps.ExitLeftwardMode
	Rmicm    = caps.ExitMicroMode
	Rshm     = caps.ExitShadowMode
	Rsubm    = caps.ExitSubscriptMode
	Rsupm    = caps.ExitSuperscriptMode
	Rum      = caps.ExitUpwardMode
	Mhpa     = caps.MicroColumnAddress
	Mcud1    = caps.MicroDown
	Mcub1    = caps.MicroLeft
	Mcuf1    = caps.MicroRight
	Mvpa     = caps.MicroRowAddress
	Mcuu1    = caps.MicroUp
	Porder   = caps.OrderOfPins
	Mcud     = caps.ParmDownMicro
	Mcub     = caps.ParmLeftMicro
	Mcuf     = caps.ParmRightMicro
	Mcuu     = caps.ParmUpMicro
	Scs      = caps.SelectCharSet
	Smgb     = caps.SetBottomMargin
	Smgbp    = caps.SetBottomMarginParm
	Smglp    = caps.SetLeftMarginParm
	Smgrp    = caps.SetRightMarginParm
	Smgt     = caps.SetTopMargin
	Smgtp    = caps.SetTopMarginParm
	Sbim     = caps.StartBitImage
	Scsd     = caps.StartCharSetDef
	Rbim     = caps.StopBitImage
	Rcsd     = caps.StopCharSetDef
	Subcs    = caps.SubscriptCharacters
	Supcs    = caps.SuperscriptCharacters
	Docr     = caps.TheseCauseCr
	Zerom    = caps.ZeroMotion
	Csnm     = caps.CharSetNames
	Kmous    = caps.KeyMouse
	Minfo    = caps.MouseInfo
	Reqmp    = caps.ReqMousePos
	Getm     = caps.GetMouse
	Setaf    = caps.SetAForeground
	Setab    = caps.SetABackground
	Pfxl     = caps.PkeyPlab
	Devt     = caps.DeviceType
	Csin     = caps.CodeSetInit
	S0ds     = caps.Set0DesSeq
	S1ds     = caps.Set1DesSeq
	S2ds     = caps.Set2DesSeq
	S3ds     = caps.Set3DesSeq
	Smglr    = caps.SetLrMargin
	Smgtb    = caps.SetTbMargin
	Birep    = caps.BitImageRepeat
	Binel    = caps.BitImageNewline
	Bicr     = caps.BitImageCarriageReturn
	Colornm  = caps.ColorNames
	Defbi    = caps.DefineBitImageRegion
	Endbi    = caps.EndBitImageRegion
	Setcolor = caps.SetColorBand
	Slines   = caps.SetPageLength
	Dispc    = caps.DisplayPcChar
	Smpch    = caps.EnterPcCharsetMode
	Rmpch    = caps.ExitPcCharsetMode
	Smsc     = caps.EnterScancodeMode
	Rmsc     = caps.ExitScancodeMode
	Pctrm    = caps.PcTermOptions
	Scesc    = caps.ScancodeEscape
	Scesa    = caps.AltScancodeEsc
	Ehhlm    = caps.EnterHorizontalHlMode
	Elhlm    = caps.EnterLeftHlMode
	Elohlm   = caps.EnterLowHlMode
	Erhlm    = caps.EnterRightHlMode
	Ethlm    = caps.EnterTopHlMode
	Evhlm    = caps.EnterVerticalHlMode
	Sgr1     = caps.SetAAttributes
	Slength  = caps.SetPglenInch
	OTi2     = caps.TermcapInit2
	OTrs     = caps.TermcapReset
	OTug     = caps.MagicCookieGlitchUl
	OTbs     = caps.BackspacesWithBs
	OTns     = caps.CrtNoScrolling
	OTnc     = caps.NoCorrectlyWorkingCr
	OTdC     = caps.CarriageReturnDelay
	OTdN     = caps.NewLineDelay
	OTnl     = caps.LinefeedIfNotLf
	OTbc     = caps.BackspaceIfNotBs
	OTMT     = caps.GnuHasMetaKey
	OTNL     = caps.LinefeedIsNewline
	OTdB     = caps.BackspaceDelay
	OTdT     = caps.HorizontalTabDelay
	OTkn     = caps.NumberOfFunctionKeys
	OTko     = caps.OtherNonFunctionKeys
	OTma     = caps.ArrowKeyMap
	OTpt     = caps.HasHardwareTabs
	OTxr     = caps.ReturnDoesClrEol
	OTG2     = caps.AcsUlcorner
	OTG3     = caps.AcsLlcorner
	OTG1     = caps.AcsUrcorner
	OTG4     = caps.AcsLrcorner
	OTGR     = caps.AcsLtee
	OTGL     = caps.AcsRtee
	OTGU     = caps.AcsBtee
	OTGD     = caps.AcsTtee
	OTGH     = caps.AcsHline
	OTGV     = caps.AcsVline
	OTGC     = caps.AcsPlus
	Meml     = caps.MemoryLock
	Memu     = caps.MemoryUnlock
	Box1     = caps.BoxChars1

	// Extentions
	CO     = caps.CO
	E3     = caps.E3
	NQ     = caps.NQ
	RGB    = caps.RGB
	TS     = caps.TS
	U8     = caps.U8
	XM     = caps.XM
	Grbom  = caps.Grbom
	Gsbom  = caps.Gsbom
	Xm     = caps.Xm
	Rmol   = caps.Rmol
	Smol   = caps.Smol
	Blink2 = caps.Blink2
	Norm   = caps.Norm
	Opaq   = caps.Opaq
	Setal  = caps.Setal
	Smul2  = caps.Smul2
	AN     = caps.AN
	AX     = caps.AX
	C0     = caps.C0
	C8     = caps.C8
	CE     = caps.CE
	CS     = caps.CS
	E0     = caps.E0
	G0     = caps.G0
	KJ     = caps.KJ
	OL     = caps.OL
	S0     = caps.S0
	TF     = caps.TF
	WS     = caps.WS
	XC     = caps.XC
	XT     = caps.XT
	Z0     = caps.Z0
	Z1     = caps.Z1
	Cr     = caps.Cr
	Cs     = caps.Cs
	Csr    = caps.Csr
	Ms     = caps.Ms
	Se     = caps.Se
	Smulx  = caps.Smulx
	Ss     = caps.Ss
	Rmxx   = caps.Rmxx
	Smxx   = caps.Smxx
	Csl    = caps.Csl
	KDC3   = caps.KDC3
	KDC4   = caps.KDC4
	KDC5   = caps.KDC5
	KDC6   = caps.KDC6
	KDC7   = caps.KDC7
	KDN    = caps.KDN
	KDN3   = caps.KDN3
	KDN4   = caps.KDN4
	KDN5   = caps.KDN5
	KDN6   = caps.KDN6
	KDN7   = caps.KDN7
	KEND3  = caps.KEND3
	KEND4  = caps.KEND4
	KEND5  = caps.KEND5
	KEND6  = caps.KEND6
	KEND7  = caps.KEND7
	KHOM3  = caps.KHOM3
	KHOM4  = caps.KHOM4
	KHOM5  = caps.KHOM5
	KHOM6  = caps.KHOM6
	KHOM7  = caps.KHOM7
	KIC3   = caps.KIC3
	KIC4   = caps.KIC4
	KIC5   = caps.KIC5
	KIC6   = caps.KIC6
	KIC7   = caps.KIC7
	KLFT3  = caps.KLFT3
	KLFT4  = caps.KLFT4
	KLFT5  = caps.KLFT5
	KLFT6  = caps.KLFT6
	KLFT7  = caps.KLFT7
	KNXT3  = caps.KNXT3
	KNXT4  = caps.KNXT4
	KNXT5  = caps.KNXT5
	KNXT6  = caps.KNXT6
	KNXT7  = caps.KNXT7
	KPRV3  = caps.KPRV3
	KPRV4  = caps.KPRV4
	KPRV5  = caps.KPRV5
	KPRV6  = caps.KPRV6
	KPRV7  = caps.KPRV7
	KRIT3  = caps.KRIT3
	KRIT4  = caps.KRIT4
	KRIT5  = caps.KRIT5
	KRIT6  = caps.KRIT6
	KRIT7  = caps.KRIT7
	KUP    = caps.KUP
	KUP3   = caps.KUP3
	KUP4   = caps.KUP4
	KUP5   = caps.KUP5
	KUP6   = caps.KUP6
	KUP7   = caps.KUP7
	Ka2    = caps.Ka2
	Kb1    = caps.Kb1
	Kb3    = caps.Kb3
	Kc2    = caps.Kc2
)
