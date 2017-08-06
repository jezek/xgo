package xgo

import "github.com/BurntSushi/xgb/xproto"

var keysym map[string]xproto.Keysym = map[string]xproto.Keysym{
	"VoidSymbol":                  0xffffff,
	"BackSpace":                   0xff08,
	"Tab":                         0xff09,
	"Linefeed":                    0xff0a,
	"Clear":                       0xff0b,
	"Return":                      0xff0d,
	"Pause":                       0xff13,
	"Scroll_Lock":                 0xff14,
	"Sys_Req":                     0xff15,
	"Escape":                      0xff1b,
	"Delete":                      0xffff,
	"Multi_key":                   0xff20,
	"Codeinput":                   0xff37,
	"SingleCandidate":             0xff3c,
	"MultipleCandidate":           0xff3d,
	"PreviousCandidate":           0xff3e,
	"Kanji":                       0xff21,
	"Muhenkan":                    0xff22,
	"Henkan_Mode":                 0xff23,
	"Henkan":                      0xff23,
	"Romaji":                      0xff24,
	"Hiragana":                    0xff25,
	"Katakana":                    0xff26,
	"Hiragana_Katakana":           0xff27,
	"Zenkaku":                     0xff28,
	"Hankaku":                     0xff29,
	"Zenkaku_Hankaku":             0xff2a,
	"Touroku":                     0xff2b,
	"Massyo":                      0xff2c,
	"Kana_Lock":                   0xff2d,
	"Kana_Shift":                  0xff2e,
	"Eisu_Shift":                  0xff2f,
	"Eisu_toggle":                 0xff30,
	"Kanji_Bangou":                0xff37,
	"Zen_Koho":                    0xff3d,
	"Mae_Koho":                    0xff3e,
	"Home":                        0xff50,
	"Left":                        0xff51,
	"Up":                          0xff52,
	"Right":                       0xff53,
	"Down":                        0xff54,
	"Prior":                       0xff55,
	"Page_Up":                     0xff55,
	"Next":                        0xff56,
	"Page_Down":                   0xff56,
	"End":                         0xff57,
	"Begin":                       0xff58,
	"Select":                      0xff60,
	"Print":                       0xff61,
	"Execute":                     0xff62,
	"Insert":                      0xff63,
	"Undo":                        0xff65,
	"Redo":                        0xff66,
	"Menu":                        0xff67,
	"Find":                        0xff68,
	"Cancel":                      0xff69,
	"Help":                        0xff6a,
	"Break":                       0xff6b,
	"Mode_switch":                 0xff7e,
	"script_switch":               0xff7e,
	"Num_Lock":                    0xff7f,
	"KP_Space":                    0xff80,
	"KP_Tab":                      0xff89,
	"KP_Enter":                    0xff8d,
	"KP_F1":                       0xff91,
	"KP_F2":                       0xff92,
	"KP_F3":                       0xff93,
	"KP_F4":                       0xff94,
	"KP_Home":                     0xff95,
	"KP_Left":                     0xff96,
	"KP_Up":                       0xff97,
	"KP_Right":                    0xff98,
	"KP_Down":                     0xff99,
	"KP_Prior":                    0xff9a,
	"KP_Page_Up":                  0xff9a,
	"KP_Next":                     0xff9b,
	"KP_Page_Down":                0xff9b,
	"KP_End":                      0xff9c,
	"KP_Begin":                    0xff9d,
	"KP_Insert":                   0xff9e,
	"KP_Delete":                   0xff9f,
	"KP_Equal":                    0xffbd,
	"KP_Multiply":                 0xffaa,
	"KP_Add":                      0xffab,
	"KP_Separator":                0xffac,
	"KP_Subtract":                 0xffad,
	"KP_Decimal":                  0xffae,
	"KP_Divide":                   0xffaf,
	"KP_0":                        0xffb0,
	"KP_1":                        0xffb1,
	"KP_2":                        0xffb2,
	"KP_3":                        0xffb3,
	"KP_4":                        0xffb4,
	"KP_5":                        0xffb5,
	"KP_6":                        0xffb6,
	"KP_7":                        0xffb7,
	"KP_8":                        0xffb8,
	"KP_9":                        0xffb9,
	"F1":                          0xffbe,
	"F2":                          0xffbf,
	"F3":                          0xffc0,
	"F4":                          0xffc1,
	"F5":                          0xffc2,
	"F6":                          0xffc3,
	"F7":                          0xffc4,
	"F8":                          0xffc5,
	"F9":                          0xffc6,
	"F10":                         0xffc7,
	"F11":                         0xffc8,
	"L1":                          0xffc8,
	"F12":                         0xffc9,
	"L2":                          0xffc9,
	"F13":                         0xffca,
	"L3":                          0xffca,
	"F14":                         0xffcb,
	"L4":                          0xffcb,
	"F15":                         0xffcc,
	"L5":                          0xffcc,
	"F16":                         0xffcd,
	"L6":                          0xffcd,
	"F17":                         0xffce,
	"L7":                          0xffce,
	"F18":                         0xffcf,
	"L8":                          0xffcf,
	"F19":                         0xffd0,
	"L9":                          0xffd0,
	"F20":                         0xffd1,
	"L10":                         0xffd1,
	"F21":                         0xffd2,
	"R1":                          0xffd2,
	"F22":                         0xffd3,
	"R2":                          0xffd3,
	"F23":                         0xffd4,
	"R3":                          0xffd4,
	"F24":                         0xffd5,
	"R4":                          0xffd5,
	"F25":                         0xffd6,
	"R5":                          0xffd6,
	"F26":                         0xffd7,
	"R6":                          0xffd7,
	"F27":                         0xffd8,
	"R7":                          0xffd8,
	"F28":                         0xffd9,
	"R8":                          0xffd9,
	"F29":                         0xffda,
	"R9":                          0xffda,
	"F30":                         0xffdb,
	"R10":                         0xffdb,
	"F31":                         0xffdc,
	"R11":                         0xffdc,
	"F32":                         0xffdd,
	"R12":                         0xffdd,
	"F33":                         0xffde,
	"R13":                         0xffde,
	"F34":                         0xffdf,
	"R14":                         0xffdf,
	"F35":                         0xffe0,
	"R15":                         0xffe0,
	"Shift_L":                     0xffe1,
	"Shift_R":                     0xffe2,
	"Control_L":                   0xffe3,
	"Control_R":                   0xffe4,
	"Caps_Lock":                   0xffe5,
	"Shift_Lock":                  0xffe6,
	"Meta_L":                      0xffe7,
	"Meta_R":                      0xffe8,
	"Alt_L":                       0xffe9,
	"Alt_R":                       0xffea,
	"Super_L":                     0xffeb,
	"Super_R":                     0xffec,
	"Hyper_L":                     0xffed,
	"Hyper_R":                     0xffee,
	"ISO_Lock":                    0xfe01,
	"ISO_Level2_Latch":            0xfe02,
	"ISO_Level3_Shift":            0xfe03,
	"ISO_Level3_Latch":            0xfe04,
	"ISO_Level3_Lock":             0xfe05,
	"ISO_Level5_Shift":            0xfe11,
	"ISO_Level5_Latch":            0xfe12,
	"ISO_Level5_Lock":             0xfe13,
	"ISO_Group_Shift":             0xff7e,
	"ISO_Group_Latch":             0xfe06,
	"ISO_Group_Lock":              0xfe07,
	"ISO_Next_Group":              0xfe08,
	"ISO_Next_Group_Lock":         0xfe09,
	"ISO_Prev_Group":              0xfe0a,
	"ISO_Prev_Group_Lock":         0xfe0b,
	"ISO_First_Group":             0xfe0c,
	"ISO_First_Group_Lock":        0xfe0d,
	"ISO_Last_Group":              0xfe0e,
	"ISO_Last_Group_Lock":         0xfe0f,
	"ISO_Left_Tab":                0xfe20,
	"ISO_Move_Line_Up":            0xfe21,
	"ISO_Move_Line_Down":          0xfe22,
	"ISO_Partial_Line_Up":         0xfe23,
	"ISO_Partial_Line_Down":       0xfe24,
	"ISO_Partial_Space_Left":      0xfe25,
	"ISO_Partial_Space_Right":     0xfe26,
	"ISO_Set_Margin_Left":         0xfe27,
	"ISO_Set_Margin_Right":        0xfe28,
	"ISO_Release_Margin_Left":     0xfe29,
	"ISO_Release_Margin_Right":    0xfe2a,
	"ISO_Release_Both_Margins":    0xfe2b,
	"ISO_Fast_Cursor_Left":        0xfe2c,
	"ISO_Fast_Cursor_Right":       0xfe2d,
	"ISO_Fast_Cursor_Up":          0xfe2e,
	"ISO_Fast_Cursor_Down":        0xfe2f,
	"ISO_Continuous_Underline":    0xfe30,
	"ISO_Discontinuous_Underline": 0xfe31,
	"ISO_Emphasize":               0xfe32,
	"ISO_Center_Object":           0xfe33,
	"ISO_Enter":                   0xfe34,
	"dead_grave":                  0xfe50,
	"dead_acute":                  0xfe51,
	"dead_circumflex":             0xfe52,
	"dead_tilde":                  0xfe53,
	"dead_perispomeni":            0xfe53,
	"dead_macron":                 0xfe54,
	"dead_breve":                  0xfe55,
	"dead_abovedot":               0xfe56,
	"dead_diaeresis":              0xfe57,
	"dead_abovering":              0xfe58,
	"dead_doubleacute":            0xfe59,
	"dead_caron":                  0xfe5a,
	"dead_cedilla":                0xfe5b,
	"dead_ogonek":                 0xfe5c,
	"dead_iota":                   0xfe5d,
	"dead_voiced_sound":           0xfe5e,
	"dead_semivoiced_sound":       0xfe5f,
	"dead_belowdot":               0xfe60,
	"dead_hook":                   0xfe61,
	"dead_horn":                   0xfe62,
	"dead_stroke":                 0xfe63,
	"dead_abovecomma":             0xfe64,
	"dead_psili":                  0xfe64,
	"dead_abovereversedcomma":     0xfe65,
	"dead_dasia":                  0xfe65,
	"dead_doublegrave":            0xfe66,
	"dead_belowring":              0xfe67,
	"dead_belowmacron":            0xfe68,
	"dead_belowcircumflex":        0xfe69,
	"dead_belowtilde":             0xfe6a,
	"dead_belowbreve":             0xfe6b,
	"dead_belowdiaeresis":         0xfe6c,
	"dead_invertedbreve":          0xfe6d,
	"dead_belowcomma":             0xfe6e,
	"dead_currency":               0xfe6f,
	"dead_a":                      0xfe80,
	"dead_A":                      0xfe81,
	"dead_e":                      0xfe82,
	"dead_E":                      0xfe83,
	"dead_i":                      0xfe84,
	"dead_I":                      0xfe85,
	"dead_o":                      0xfe86,
	"dead_O":                      0xfe87,
	"dead_u":                      0xfe88,
	"dead_U":                      0xfe89,
	"dead_small_schwa":            0xfe8a,
	"dead_capital_schwa":          0xfe8b,
	"First_Virtual_Screen":        0xfed0,
	"Prev_Virtual_Screen":         0xfed1,
	"Next_Virtual_Screen":         0xfed2,
	"Last_Virtual_Screen":         0xfed4,
	"Terminate_Server":            0xfed5,
	"AccessX_Enable":              0xfe70,
	"AccessX_Feedback_Enable":     0xfe71,
	"RepeatKeys_Enable":           0xfe72,
	"SlowKeys_Enable":             0xfe73,
	"BounceKeys_Enable":           0xfe74,
	"StickyKeys_Enable":           0xfe75,
	"MouseKeys_Enable":            0xfe76,
	"MouseKeys_Accel_Enable":      0xfe77,
	"Overlay1_Enable":             0xfe78,
	"Overlay2_Enable":             0xfe79,
	"AudibleBell_Enable":          0xfe7a,
	"Pointer_Left":                0xfee0,
	"Pointer_Right":               0xfee1,
	"Pointer_Up":                  0xfee2,
	"Pointer_Down":                0xfee3,
	"Pointer_UpLeft":              0xfee4,
	"Pointer_UpRight":             0xfee5,
	"Pointer_DownLeft":            0xfee6,
	"Pointer_DownRight":           0xfee7,
	"Pointer_Button_Dflt":         0xfee8,
	"Pointer_Button1":             0xfee9,
	"Pointer_Button2":             0xfeea,
	"Pointer_Button3":             0xfeeb,
	"Pointer_Button4":             0xfeec,
	"Pointer_Button5":             0xfeed,
	"Pointer_DblClick_Dflt":       0xfeee,
	"Pointer_DblClick1":           0xfeef,
	"Pointer_DblClick2":           0xfef0,
	"Pointer_DblClick3":           0xfef1,
	"Pointer_DblClick4":           0xfef2,
	"Pointer_DblClick5":           0xfef3,
	"Pointer_Drag_Dflt":           0xfef4,
	"Pointer_Drag1":               0xfef5,
	"Pointer_Drag2":               0xfef6,
	"Pointer_Drag3":               0xfef7,
	"Pointer_Drag4":               0xfef8,
	"Pointer_Drag5":               0xfefd,
	"Pointer_EnableKeys":          0xfef9,
	"Pointer_Accelerate":          0xfefa,
	"Pointer_DfltBtnNext":         0xfefb,
	"Pointer_DfltBtnPrev":         0xfefc,
	"3270_Duplicate":              0xfd01,
	"3270_FieldMark":              0xfd02,
	"3270_Right2":                 0xfd03,
	"3270_Left2":                  0xfd04,
	"3270_BackTab":                0xfd05,
	"3270_EraseEOF":               0xfd06,
	"3270_EraseInput":             0xfd07,
	"3270_Reset":                  0xfd08,
	"3270_Quit":                   0xfd09,
	"3270_PA1":                    0xfd0a,
	"3270_PA2":                    0xfd0b,
	"3270_PA3":                    0xfd0c,
	"3270_Test":                   0xfd0d,
	"3270_Attn":                   0xfd0e,
	"3270_CursorBlink":            0xfd0f,
	"3270_AltCursor":              0xfd10,
	"3270_KeyClick":               0xfd11,
	"3270_Jump":                   0xfd12,
	"3270_Ident":                  0xfd13,
	"3270_Rule":                   0xfd14,
	"3270_Copy":                   0xfd15,
	"3270_Play":                   0xfd16,
	"3270_Setup":                  0xfd17,
	"3270_Record":                 0xfd18,
	"3270_ChangeScreen":           0xfd19,
	"3270_DeleteWord":             0xfd1a,
	"3270_ExSelect":               0xfd1b,
	"3270_CursorSelect":           0xfd1c,
	"3270_PrintScreen":            0xfd1d,
	"3270_Enter":                  0xfd1e,
	"space":                       0x0020,
	"exclam":                      0x0021,
	"quotedbl":                    0x0022,
	"numbersign":                  0x0023,
	"dollar":                      0x0024,
	"percent":                     0x0025,
	"ampersand":                   0x0026,
	"apostrophe":                  0x0027,
	"quoteright":                  0x0027,
	"parenleft":                   0x0028,
	"parenright":                  0x0029,
	"asterisk":                    0x002a,
	"plus":                        0x002b,
	"comma":                       0x002c,
	"minus":                       0x002d,
	"period":                      0x002e,
	"slash":                       0x002f,
	"0":                           0x0030,
	"1":                           0x0031,
	"2":                           0x0032,
	"3":                           0x0033,
	"4":                           0x0034,
	"5":                           0x0035,
	"6":                           0x0036,
	"7":                           0x0037,
	"8":                           0x0038,
	"9":                           0x0039,
	"colon":                       0x003a,
	"semicolon":                   0x003b,
	"less":                        0x003c,
	"equal":                       0x003d,
	"greater":                     0x003e,
	"question":                    0x003f,
	"at":                          0x0040,
	"A":                           0x0041,
	"B":                           0x0042,
	"C":                           0x0043,
	"D":                           0x0044,
	"E":                           0x0045,
	"F":                           0x0046,
	"G":                           0x0047,
	"H":                           0x0048,
	"I":                           0x0049,
	"J":                           0x004a,
	"K":                           0x004b,
	"L":                           0x004c,
	"M":                           0x004d,
	"N":                           0x004e,
	"O":                           0x004f,
	"P":                           0x0050,
	"Q":                           0x0051,
	"R":                           0x0052,
	"S":                           0x0053,
	"T":                           0x0054,
	"U":                           0x0055,
	"V":                           0x0056,
	"W":                           0x0057,
	"X":                           0x0058,
	"Y":                           0x0059,
	"Z":                           0x005a,
	"bracketleft":                 0x005b,
	"backslash":                   0x005c,
	"bracketright":                0x005d,
	"asciicircum":                 0x005e,
	"underscore":                  0x005f,
	"grave":                       0x0060,
	"quoteleft":                   0x0060,
	"a":                           0x0061,
	"b":                           0x0062,
	"c":                           0x0063,
	"d":                           0x0064,
	"e":                           0x0065,
	"f":                           0x0066,
	"g":                           0x0067,
	"h":                           0x0068,
	"i":                           0x0069,
	"j":                           0x006a,
	"k":                           0x006b,
	"l":                           0x006c,
	"m":                           0x006d,
	"n":                           0x006e,
	"o":                           0x006f,
	"p":                           0x0070,
	"q":                           0x0071,
	"r":                           0x0072,
	"s":                           0x0073,
	"t":                           0x0074,
	"u":                           0x0075,
	"v":                           0x0076,
	"w":                           0x0077,
	"x":                           0x0078,
	"y":                           0x0079,
	"z":                           0x007a,
	"braceleft":                   0x007b,
	"bar":                         0x007c,
	"braceright":                  0x007d,
	"asciitilde":                  0x007e,
	"nobreakspace":                0x00a0,
	"exclamdown":                  0x00a1,
	"cent":                        0x00a2,
	"sterling":                    0x00a3,
	"currency":                    0x00a4,
	"yen":                         0x00a5,
	"brokenbar":                   0x00a6,
	"section":                     0x00a7,
	"diaeresis":                   0x00a8,
	"copyright":                   0x00a9,
	"ordfeminine":                 0x00aa,
	"guillemotleft":               0x00ab,
	"notsign":                     0x00ac,
	"hyphen":                      0x00ad,
	"registered":                  0x00ae,
	"macron":                      0x00af,
	"degree":                      0x00b0,
	"plusminus":                   0x00b1,
	"twosuperior":                 0x00b2,
	"threesuperior":               0x00b3,
	"acute":                       0x00b4,
	"mu":                          0x00b5,
	"paragraph":                   0x00b6,
	"periodcentered":              0x00b7,
	"cedilla":                     0x00b8,
	"onesuperior":                 0x00b9,
	"masculine":                   0x00ba,
	"guillemotright":              0x00bb,
	"onequarter":                  0x00bc,
	"onehalf":                     0x00bd,
	"threequarters":               0x00be,
	"questiondown":                0x00bf,
	"Agrave":                      0x00c0,
	"Aacute":                      0x00c1,
	"Acircumflex":                 0x00c2,
	"Atilde":                      0x00c3,
	"Adiaeresis":                  0x00c4,
	"Aring":                       0x00c5,
	"AE":                          0x00c6,
	"Ccedilla":                    0x00c7,
	"Egrave":                      0x00c8,
	"Eacute":                      0x00c9,
	"Ecircumflex":                 0x00ca,
	"Ediaeresis":                  0x00cb,
	"Igrave":                      0x00cc,
	"Iacute":                      0x00cd,
	"Icircumflex":                 0x00ce,
	"Idiaeresis":                  0x00cf,
	"ETH":                         0x00d0,
	"Eth":                         0x00d0,
	"Ntilde":                      0x00d1,
	"Ograve":                      0x00d2,
	"Oacute":                      0x00d3,
	"Ocircumflex":                 0x00d4,
	"Otilde":                      0x00d5,
	"Odiaeresis":                  0x00d6,
	"multiply":                    0x00d7,
	"Oslash":                      0x00d8,
	"Ooblique":                    0x00d8,
	"Ugrave":                      0x00d9,
	"Uacute":                      0x00da,
	"Ucircumflex":                 0x00db,
	"Udiaeresis":                  0x00dc,
	"Yacute":                      0x00dd,
	"THORN":                       0x00de,
	"Thorn":                       0x00de,
	"ssharp":                      0x00df,
	"agrave":                      0x00e0,
	"aacute":                      0x00e1,
	"acircumflex":                 0x00e2,
	"atilde":                      0x00e3,
	"adiaeresis":                  0x00e4,
	"aring":                       0x00e5,
	"ae":                          0x00e6,
	"ccedilla":                    0x00e7,
	"egrave":                      0x00e8,
	"eacute":                      0x00e9,
	"ecircumflex":                 0x00ea,
	"ediaeresis":                  0x00eb,
	"igrave":                      0x00ec,
	"iacute":                      0x00ed,
	"icircumflex":                 0x00ee,
	"idiaeresis":                  0x00ef,
	"eth":                         0x00f0,
	"ntilde":                      0x00f1,
	"ograve":                      0x00f2,
	"oacute":                      0x00f3,
	"ocircumflex":                 0x00f4,
	"otilde":                      0x00f5,
	"odiaeresis":                  0x00f6,
	"division":                    0x00f7,
	"oslash":                      0x00f8,
	"ooblique":                    0x00f8,
	"ugrave":                      0x00f9,
	"uacute":                      0x00fa,
	"ucircumflex":                 0x00fb,
	"udiaeresis":                  0x00fc,
	"yacute":                      0x00fd,
	"thorn":                       0x00fe,
	"ydiaeresis":                  0x00ff,
	"Aogonek":                     0x01a1,
	"breve":                       0x01a2,
	"Lstroke":                     0x01a3,
	"Lcaron":                      0x01a5,
	"Sacute":                      0x01a6,
	"Scaron":                      0x01a9,
	"Scedilla":                    0x01aa,
	"Tcaron":                      0x01ab,
	"Zacute":                      0x01ac,
	"Zcaron":                      0x01ae,
	"Zabovedot":                   0x01af,
	"aogonek":                     0x01b1,
	"ogonek":                      0x01b2,
	"lstroke":                     0x01b3,
	"lcaron":                      0x01b5,
	"sacute":                      0x01b6,
	"caron":                       0x01b7,
	"scaron":                      0x01b9,
	"scedilla":                    0x01ba,
	"tcaron":                      0x01bb,
	"zacute":                      0x01bc,
	"doubleacute":                 0x01bd,
	"zcaron":                      0x01be,
	"zabovedot":                   0x01bf,
	"Racute":                      0x01c0,
	"Abreve":                      0x01c3,
	"Lacute":                      0x01c5,
	"Cacute":                      0x01c6,
	"Ccaron":                      0x01c8,
	"Eogonek":                     0x01ca,
	"Ecaron":                      0x01cc,
	"Dcaron":                      0x01cf,
	"Dstroke":                     0x01d0,
	"Nacute":                      0x01d1,
	"Ncaron":                      0x01d2,
	"Odoubleacute":                0x01d5,
	"Rcaron":                      0x01d8,
	"Uring":                       0x01d9,
	"Udoubleacute":                0x01db,
	"Tcedilla":                    0x01de,
	"racute":                      0x01e0,
	"abreve":                      0x01e3,
	"lacute":                      0x01e5,
	"cacute":                      0x01e6,
	"ccaron":                      0x01e8,
	"eogonek":                     0x01ea,
	"ecaron":                      0x01ec,
	"dcaron":                      0x01ef,
	"dstroke":                     0x01f0,
	"nacute":                      0x01f1,
	"ncaron":                      0x01f2,
	"odoubleacute":                0x01f5,
	"udoubleacute":                0x01fb,
	"rcaron":                      0x01f8,
	"uring":                       0x01f9,
	"tcedilla":                    0x01fe,
	"abovedot":                    0x01ff,
	"Hstroke":                     0x02a1,
	"Hcircumflex":                 0x02a6,
	"Iabovedot":                   0x02a9,
	"Gbreve":                      0x02ab,
	"Jcircumflex":                 0x02ac,
	"hstroke":                     0x02b1,
	"hcircumflex":                 0x02b6,
	"idotless":                    0x02b9,
	"gbreve":                      0x02bb,
	"jcircumflex":                 0x02bc,
	"Cabovedot":                   0x02c5,
	"Ccircumflex":                 0x02c6,
	"Gabovedot":                   0x02d5,
	"Gcircumflex":                 0x02d8,
	"Ubreve":                      0x02dd,
	"Scircumflex":                 0x02de,
	"cabovedot":                   0x02e5,
	"ccircumflex":                 0x02e6,
	"gabovedot":                   0x02f5,
	"gcircumflex":                 0x02f8,
	"ubreve":                      0x02fd,
	"scircumflex":                 0x02fe,
	"kra":                         0x03a2,
	"kappa":                       0x03a2,
	"Rcedilla":                    0x03a3,
	"Itilde":                      0x03a5,
	"Lcedilla":                    0x03a6,
	"Emacron":                     0x03aa,
	"Gcedilla":                    0x03ab,
	"Tslash":                      0x03ac,
	"rcedilla":                    0x03b3,
	"itilde":                      0x03b5,
	"lcedilla":                    0x03b6,
	"emacron":                     0x03ba,
	"gcedilla":                    0x03bb,
	"tslash":                      0x03bc,
	"ENG":                         0x03bd,
	"eng":                         0x03bf,
	"Amacron":                     0x03c0,
	"Iogonek":                     0x03c7,
	"Eabovedot":                   0x03cc,
	"Imacron":                     0x03cf,
	"Ncedilla":                    0x03d1,
	"Omacron":                     0x03d2,
	"Kcedilla":                    0x03d3,
	"Uogonek":                     0x03d9,
	"Utilde":                      0x03dd,
	"Umacron":                     0x03de,
	"amacron":                     0x03e0,
	"iogonek":                     0x03e7,
	"eabovedot":                   0x03ec,
	"imacron":                     0x03ef,
	"ncedilla":                    0x03f1,
	"omacron":                     0x03f2,
	"kcedilla":                    0x03f3,
	"uogonek":                     0x03f9,
	"utilde":                      0x03fd,
	"umacron":                     0x03fe,
	"Babovedot":                   0x1001e02,
	"babovedot":                   0x1001e03,
	"Dabovedot":                   0x1001e0a,
	"Wgrave":                      0x1001e80,
	"Wacute":                      0x1001e82,
	"dabovedot":                   0x1001e0b,
	"Ygrave":                      0x1001ef2,
	"Fabovedot":                   0x1001e1e,
	"fabovedot":                   0x1001e1f,
	"Mabovedot":                   0x1001e40,
	"mabovedot":                   0x1001e41,
	"Pabovedot":                   0x1001e56,
	"wgrave":                      0x1001e81,
	"pabovedot":                   0x1001e57,
	"wacute":                      0x1001e83,
	"Sabovedot":                   0x1001e60,
	"ygrave":                      0x1001ef3,
	"Wdiaeresis":                  0x1001e84,
	"wdiaeresis":                  0x1001e85,
	"sabovedot":                   0x1001e61,
	"Wcircumflex":                 0x1000174,
	"Tabovedot":                   0x1001e6a,
	"Ycircumflex":                 0x1000176,
	"wcircumflex":                 0x1000175,
	"tabovedot":                   0x1001e6b,
	"ycircumflex":                 0x1000177,
	"OE":                          0x13bc,
	"oe":                          0x13bd,
	"Ydiaeresis":                  0x13be,
	"overline":                    0x047e,
	"kana_fullstop":               0x04a1,
	"kana_openingbracket":         0x04a2,
	"kana_closingbracket":         0x04a3,
	"kana_comma":                  0x04a4,
	"kana_conjunctive":            0x04a5,
	"kana_middledot":              0x04a5,
	"kana_WO":                     0x04a6,
	"kana_a":                      0x04a7,
	"kana_i":                      0x04a8,
	"kana_u":                      0x04a9,
	"kana_e":                      0x04aa,
	"kana_o":                      0x04ab,
	"kana_ya":                     0x04ac,
	"kana_yu":                     0x04ad,
	"kana_yo":                     0x04ae,
	"kana_tsu":                    0x04af,
	"kana_tu":                     0x04af,
	"prolongedsound":              0x04b0,
	"kana_A":                      0x04b1,
	"kana_I":                      0x04b2,
	"kana_U":                      0x04b3,
	"kana_E":                      0x04b4,
	"kana_O":                      0x04b5,
	"kana_KA":                     0x04b6,
	"kana_KI":                     0x04b7,
	"kana_KU":                     0x04b8,
	"kana_KE":                     0x04b9,
	"kana_KO":                     0x04ba,
	"kana_SA":                     0x04bb,
	"kana_SHI":                    0x04bc,
	"kana_SU":                     0x04bd,
	"kana_SE":                     0x04be,
	"kana_SO":                     0x04bf,
	"kana_TA":                     0x04c0,
	"kana_CHI":                    0x04c1,
	"kana_TI":                     0x04c1,
	"kana_TSU":                    0x04c2,
	"kana_TU":                     0x04c2,
	"kana_TE":                     0x04c3,
	"kana_TO":                     0x04c4,
	"kana_NA":                     0x04c5,
	"kana_NI":                     0x04c6,
	"kana_NU":                     0x04c7,
	"kana_NE":                     0x04c8,
	"kana_NO":                     0x04c9,
	"kana_HA":                     0x04ca,
	"kana_HI":                     0x04cb,
	"kana_FU":                     0x04cc,
	"kana_HU":                     0x04cc,
	"kana_HE":                     0x04cd,
	"kana_HO":                     0x04ce,
	"kana_MA":                     0x04cf,
	"kana_MI":                     0x04d0,
	"kana_MU":                     0x04d1,
	"kana_ME":                     0x04d2,
	"kana_MO":                     0x04d3,
	"kana_YA":                     0x04d4,
	"kana_YU":                     0x04d5,
	"kana_YO":                     0x04d6,
	"kana_RA":                     0x04d7,
	"kana_RI":                     0x04d8,
	"kana_RU":                     0x04d9,
	"kana_RE":                     0x04da,
	"kana_RO":                     0x04db,
	"kana_WA":                     0x04dc,
	"kana_N":                      0x04dd,
	"voicedsound":                 0x04de,
	"semivoicedsound":             0x04df,
	"kana_switch":                 0xff7e,
	"Farsi_0":                     0x10006f0,
	"Farsi_1":                     0x10006f1,
	"Farsi_2":                     0x10006f2,
	"Farsi_3":                     0x10006f3,
	"Farsi_4":                     0x10006f4,
	"Farsi_5":                     0x10006f5,
	"Farsi_6":                     0x10006f6,
	"Farsi_7":                     0x10006f7,
	"Farsi_8":                     0x10006f8,
	"Farsi_9":                     0x10006f9,
	"Arabic_percent":              0x100066a,
	"Arabic_superscript_alef":     0x1000670,
	"Arabic_tteh":                 0x1000679,
	"Arabic_peh":                  0x100067e,
	"Arabic_tcheh":                0x1000686,
	"Arabic_ddal":                 0x1000688,
	"Arabic_rreh":                 0x1000691,
	"Arabic_comma":                0x05ac,
	"Arabic_fullstop":             0x10006d4,
	"Arabic_0":                    0x1000660,
	"Arabic_1":                    0x1000661,
	"Arabic_2":                    0x1000662,
	"Arabic_3":                    0x1000663,
	"Arabic_4":                    0x1000664,
	"Arabic_5":                    0x1000665,
	"Arabic_6":                    0x1000666,
	"Arabic_7":                    0x1000667,
	"Arabic_8":                    0x1000668,
	"Arabic_9":                    0x1000669,
	"Arabic_semicolon":            0x05bb,
	"Arabic_question_mark":        0x05bf,
	"Arabic_hamza":                0x05c1,
	"Arabic_maddaonalef":          0x05c2,
	"Arabic_hamzaonalef":          0x05c3,
	"Arabic_hamzaonwaw":           0x05c4,
	"Arabic_hamzaunderalef":       0x05c5,
	"Arabic_hamzaonyeh":           0x05c6,
	"Arabic_alef":                 0x05c7,
	"Arabic_beh":                  0x05c8,
	"Arabic_tehmarbuta":           0x05c9,
	"Arabic_teh":                  0x05ca,
	"Arabic_theh":                 0x05cb,
	"Arabic_jeem":                 0x05cc,
	"Arabic_hah":                  0x05cd,
	"Arabic_khah":                 0x05ce,
	"Arabic_dal":                  0x05cf,
	"Arabic_thal":                 0x05d0,
	"Arabic_ra":                   0x05d1,
	"Arabic_zain":                 0x05d2,
	"Arabic_seen":                 0x05d3,
	"Arabic_sheen":                0x05d4,
	"Arabic_sad":                  0x05d5,
	"Arabic_dad":                  0x05d6,
	"Arabic_tah":                  0x05d7,
	"Arabic_zah":                  0x05d8,
	"Arabic_ain":                  0x05d9,
	"Arabic_ghain":                0x05da,
	"Arabic_tatweel":              0x05e0,
	"Arabic_feh":                  0x05e1,
	"Arabic_qaf":                  0x05e2,
	"Arabic_kaf":                  0x05e3,
	"Arabic_lam":                  0x05e4,
	"Arabic_meem":                 0x05e5,
	"Arabic_noon":                 0x05e6,
	"Arabic_ha":                   0x05e7,
	"Arabic_heh":                  0x05e7,
	"Arabic_waw":                  0x05e8,
	"Arabic_alefmaksura":          0x05e9,
	"Arabic_yeh":                  0x05ea,
	"Arabic_fathatan":             0x05eb,
	"Arabic_dammatan":             0x05ec,
	"Arabic_kasratan":             0x05ed,
	"Arabic_fatha":                0x05ee,
	"Arabic_damma":                0x05ef,
	"Arabic_kasra":                0x05f0,
	"Arabic_shadda":               0x05f1,
	"Arabic_sukun":                0x05f2,
	"Arabic_madda_above":          0x1000653,
	"Arabic_hamza_above":          0x1000654,
	"Arabic_hamza_below":          0x1000655,
	"Arabic_jeh":                  0x1000698,
	"Arabic_veh":                  0x10006a4,
	"Arabic_keheh":                0x10006a9,
	"Arabic_gaf":                  0x10006af,
	"Arabic_noon_ghunna":          0x10006ba,
	"Arabic_heh_doachashmee":      0x10006be,
	"Farsi_yeh":                   0x10006cc,
	"Arabic_farsi_yeh":            0x10006cc,
	"Arabic_yeh_baree":            0x10006d2,
	"Arabic_heh_goal":             0x10006c1,
	"Arabic_switch":               0xff7e,
	"Cyrillic_GHE_bar":            0x1000492,
	"Cyrillic_ghe_bar":            0x1000493,
	"Cyrillic_ZHE_descender":      0x1000496,
	"Cyrillic_zhe_descender":      0x1000497,
	"Cyrillic_KA_descender":       0x100049a,
	"Cyrillic_ka_descender":       0x100049b,
	"Cyrillic_KA_vertstroke":      0x100049c,
	"Cyrillic_ka_vertstroke":      0x100049d,
	"Cyrillic_EN_descender":       0x10004a2,
	"Cyrillic_en_descender":       0x10004a3,
	"Cyrillic_U_straight":         0x10004ae,
	"Cyrillic_u_straight":         0x10004af,
	"Cyrillic_U_straight_bar":     0x10004b0,
	"Cyrillic_u_straight_bar":     0x10004b1,
	"Cyrillic_HA_descender":       0x10004b2,
	"Cyrillic_ha_descender":       0x10004b3,
	"Cyrillic_CHE_descender":      0x10004b6,
	"Cyrillic_che_descender":      0x10004b7,
	"Cyrillic_CHE_vertstroke":     0x10004b8,
	"Cyrillic_che_vertstroke":     0x10004b9,
	"Cyrillic_SHHA":               0x10004ba,
	"Cyrillic_shha":               0x10004bb,
	"Cyrillic_SCHWA":              0x10004d8,
	"Cyrillic_schwa":              0x10004d9,
	"Cyrillic_I_macron":           0x10004e2,
	"Cyrillic_i_macron":           0x10004e3,
	"Cyrillic_O_bar":              0x10004e8,
	"Cyrillic_o_bar":              0x10004e9,
	"Cyrillic_U_macron":           0x10004ee,
	"Cyrillic_u_macron":           0x10004ef,
	"Serbian_dje":                 0x06a1,
	"Macedonia_gje":               0x06a2,
	"Cyrillic_io":                 0x06a3,
	"Ukrainian_ie":                0x06a4,
	"Ukranian_je":                 0x06a4,
	"Macedonia_dse":               0x06a5,
	"Ukrainian_i":                 0x06a6,
	"Ukranian_i":                  0x06a6,
	"Ukrainian_yi":                0x06a7,
	"Ukranian_yi":                 0x06a7,
	"Cyrillic_je":                 0x06a8,
	"Serbian_je":                  0x06a8,
	"Cyrillic_lje":                0x06a9,
	"Serbian_lje":                 0x06a9,
	"Cyrillic_nje":                0x06aa,
	"Serbian_nje":                 0x06aa,
	"Serbian_tshe":                0x06ab,
	"Macedonia_kje":               0x06ac,
	"Ukrainian_ghe_with_upturn":   0x06ad,
	"Byelorussian_shortu":         0x06ae,
	"Cyrillic_dzhe":               0x06af,
	"Serbian_dze":                 0x06af,
	"numerosign":                  0x06b0,
	"Serbian_DJE":                 0x06b1,
	"Macedonia_GJE":               0x06b2,
	"Cyrillic_IO":                 0x06b3,
	"Ukrainian_IE":                0x06b4,
	"Ukranian_JE":                 0x06b4,
	"Macedonia_DSE":               0x06b5,
	"Ukrainian_I":                 0x06b6,
	"Ukranian_I":                  0x06b6,
	"Ukrainian_YI":                0x06b7,
	"Ukranian_YI":                 0x06b7,
	"Cyrillic_JE":                 0x06b8,
	"Serbian_JE":                  0x06b8,
	"Cyrillic_LJE":                0x06b9,
	"Serbian_LJE":                 0x06b9,
	"Cyrillic_NJE":                0x06ba,
	"Serbian_NJE":                 0x06ba,
	"Serbian_TSHE":                0x06bb,
	"Macedonia_KJE":               0x06bc,
	"Ukrainian_GHE_WITH_UPTURN":   0x06bd,
	"Byelorussian_SHORTU":         0x06be,
	"Cyrillic_DZHE":               0x06bf,
	"Serbian_DZE":                 0x06bf,
	"Cyrillic_yu":                 0x06c0,
	"Cyrillic_a":                  0x06c1,
	"Cyrillic_be":                 0x06c2,
	"Cyrillic_tse":                0x06c3,
	"Cyrillic_de":                 0x06c4,
	"Cyrillic_ie":                 0x06c5,
	"Cyrillic_ef":                 0x06c6,
	"Cyrillic_ghe":                0x06c7,
	"Cyrillic_ha":                 0x06c8,
	"Cyrillic_i":                  0x06c9,
	"Cyrillic_shorti":             0x06ca,
	"Cyrillic_ka":                 0x06cb,
	"Cyrillic_el":                 0x06cc,
	"Cyrillic_em":                 0x06cd,
	"Cyrillic_en":                 0x06ce,
	"Cyrillic_o":                  0x06cf,
	"Cyrillic_pe":                 0x06d0,
	"Cyrillic_ya":                 0x06d1,
	"Cyrillic_er":                 0x06d2,
	"Cyrillic_es":                 0x06d3,
	"Cyrillic_te":                 0x06d4,
	"Cyrillic_u":                  0x06d5,
	"Cyrillic_zhe":                0x06d6,
	"Cyrillic_ve":                 0x06d7,
	"Cyrillic_softsign":           0x06d8,
	"Cyrillic_yeru":               0x06d9,
	"Cyrillic_ze":                 0x06da,
	"Cyrillic_sha":                0x06db,
	"Cyrillic_e":                  0x06dc,
	"Cyrillic_shcha":              0x06dd,
	"Cyrillic_che":                0x06de,
	"Cyrillic_hardsign":           0x06df,
	"Cyrillic_YU":                 0x06e0,
	"Cyrillic_A":                  0x06e1,
	"Cyrillic_BE":                 0x06e2,
	"Cyrillic_TSE":                0x06e3,
	"Cyrillic_DE":                 0x06e4,
	"Cyrillic_IE":                 0x06e5,
	"Cyrillic_EF":                 0x06e6,
	"Cyrillic_GHE":                0x06e7,
	"Cyrillic_HA":                 0x06e8,
	"Cyrillic_I":                  0x06e9,
	"Cyrillic_SHORTI":             0x06ea,
	"Cyrillic_KA":                 0x06eb,
	"Cyrillic_EL":                 0x06ec,
	"Cyrillic_EM":                 0x06ed,
	"Cyrillic_EN":                 0x06ee,
	"Cyrillic_O":                  0x06ef,
	"Cyrillic_PE":                 0x06f0,
	"Cyrillic_YA":                 0x06f1,
	"Cyrillic_ER":                 0x06f2,
	"Cyrillic_ES":                 0x06f3,
	"Cyrillic_TE":                 0x06f4,
	"Cyrillic_U":                  0x06f5,
	"Cyrillic_ZHE":                0x06f6,
	"Cyrillic_VE":                 0x06f7,
	"Cyrillic_SOFTSIGN":           0x06f8,
	"Cyrillic_YERU":               0x06f9,
	"Cyrillic_ZE":                 0x06fa,
	"Cyrillic_SHA":                0x06fb,
	"Cyrillic_E":                  0x06fc,
	"Cyrillic_SHCHA":              0x06fd,
	"Cyrillic_CHE":                0x06fe,
	"Cyrillic_HARDSIGN":           0x06ff,
	"Greek_ALPHAaccent":           0x07a1,
	"Greek_EPSILONaccent":         0x07a2,
	"Greek_ETAaccent":             0x07a3,
	"Greek_IOTAaccent":            0x07a4,
	"Greek_IOTAdieresis":          0x07a5,
	"Greek_IOTAdiaeresis":         0x07a5,
	"Greek_OMICRONaccent":         0x07a7,
	"Greek_UPSILONaccent":         0x07a8,
	"Greek_UPSILONdieresis":       0x07a9,
	"Greek_OMEGAaccent":           0x07ab,
	"Greek_accentdieresis":        0x07ae,
	"Greek_horizbar":              0x07af,
	"Greek_alphaaccent":           0x07b1,
	"Greek_epsilonaccent":         0x07b2,
	"Greek_etaaccent":             0x07b3,
	"Greek_iotaaccent":            0x07b4,
	"Greek_iotadieresis":          0x07b5,
	"Greek_iotaaccentdieresis":    0x07b6,
	"Greek_omicronaccent":         0x07b7,
	"Greek_upsilonaccent":         0x07b8,
	"Greek_upsilondieresis":       0x07b9,
	"Greek_upsilonaccentdieresis": 0x07ba,
	"Greek_omegaaccent":           0x07bb,
	"Greek_ALPHA":                 0x07c1,
	"Greek_BETA":                  0x07c2,
	"Greek_GAMMA":                 0x07c3,
	"Greek_DELTA":                 0x07c4,
	"Greek_EPSILON":               0x07c5,
	"Greek_ZETA":                  0x07c6,
	"Greek_ETA":                   0x07c7,
	"Greek_THETA":                 0x07c8,
	"Greek_IOTA":                  0x07c9,
	"Greek_KAPPA":                 0x07ca,
	"Greek_LAMDA":                 0x07cb,
	"Greek_LAMBDA":                0x07cb,
	"Greek_MU":                    0x07cc,
	"Greek_NU":                    0x07cd,
	"Greek_XI":                    0x07ce,
	"Greek_OMICRON":               0x07cf,
	"Greek_PI":                    0x07d0,
	"Greek_RHO":                   0x07d1,
	"Greek_SIGMA":                 0x07d2,
	"Greek_TAU":                   0x07d4,
	"Greek_UPSILON":               0x07d5,
	"Greek_PHI":                   0x07d6,
	"Greek_CHI":                   0x07d7,
	"Greek_PSI":                   0x07d8,
	"Greek_OMEGA":                 0x07d9,
	"Greek_alpha":                 0x07e1,
	"Greek_beta":                  0x07e2,
	"Greek_gamma":                 0x07e3,
	"Greek_delta":                 0x07e4,
	"Greek_epsilon":               0x07e5,
	"Greek_zeta":                  0x07e6,
	"Greek_eta":                   0x07e7,
	"Greek_theta":                 0x07e8,
	"Greek_iota":                  0x07e9,
	"Greek_kappa":                 0x07ea,
	"Greek_lamda":                 0x07eb,
	"Greek_lambda":                0x07eb,
	"Greek_mu":                    0x07ec,
	"Greek_nu":                    0x07ed,
	"Greek_xi":                    0x07ee,
	"Greek_omicron":               0x07ef,
	"Greek_pi":                    0x07f0,
	"Greek_rho":                   0x07f1,
	"Greek_sigma":                 0x07f2,
	"Greek_finalsmallsigma":       0x07f3,
	"Greek_tau":                   0x07f4,
	"Greek_upsilon":               0x07f5,
	"Greek_phi":                   0x07f6,
	"Greek_chi":                   0x07f7,
	"Greek_psi":                   0x07f8,
	"Greek_omega":                 0x07f9,
	"Greek_switch":                0xff7e,
	"leftradical":                 0x08a1,
	"topleftradical":              0x08a2,
	"horizconnector":              0x08a3,
	"topintegral":                 0x08a4,
	"botintegral":                 0x08a5,
	"vertconnector":               0x08a6,
	"topleftsqbracket":            0x08a7,
	"botleftsqbracket":            0x08a8,
	"toprightsqbracket":           0x08a9,
	"botrightsqbracket":           0x08aa,
	"topleftparens":               0x08ab,
	"botleftparens":               0x08ac,
	"toprightparens":              0x08ad,
	"botrightparens":              0x08ae,
	"leftmiddlecurlybrace":        0x08af,
	"rightmiddlecurlybrace":       0x08b0,
	"topleftsummation":            0x08b1,
	"botleftsummation":            0x08b2,
	"topvertsummationconnector":   0x08b3,
	"botvertsummationconnector":   0x08b4,
	"toprightsummation":           0x08b5,
	"botrightsummation":           0x08b6,
	"rightmiddlesummation":        0x08b7,
	"lessthanequal":               0x08bc,
	"notequal":                    0x08bd,
	"greaterthanequal":            0x08be,
	"integral":                    0x08bf,
	"therefore":                   0x08c0,
	"variation":                   0x08c1,
	"infinity":                    0x08c2,
	"nabla":                       0x08c5,
	"approximate":                 0x08c8,
	"similarequal":                0x08c9,
	"ifonlyif":                    0x08cd,
	"implies":                     0x08ce,
	"identical":                   0x08cf,
	"radical":                     0x08d6,
	"includedin":                  0x08da,
	"includes":                    0x08db,
	"intersection":                0x08dc,
	"union":                       0x08dd,
	"logicaland":                  0x08de,
	"logicalor":                   0x08df,
	"partialderivative":           0x08ef,
	"function":                    0x08f6,
	"leftarrow":                   0x08fb,
	"uparrow":                     0x08fc,
	"rightarrow":                  0x08fd,
	"downarrow":                   0x08fe,
	"blank":                       0x09df,
	"soliddiamond":                0x09e0,
	"checkerboard":                0x09e1,
	"ht":                          0x09e2,
	"ff":                          0x09e3,
	"cr":                          0x09e4,
	"lf":                          0x09e5,
	"nl":                          0x09e8,
	"vt":                          0x09e9,
	"lowrightcorner":              0x09ea,
	"uprightcorner":               0x09eb,
	"upleftcorner":                0x09ec,
	"lowleftcorner":               0x09ed,
	"crossinglines":               0x09ee,
	"horizlinescan1":              0x09ef,
	"horizlinescan3":              0x09f0,
	"horizlinescan5":              0x09f1,
	"horizlinescan7":              0x09f2,
	"horizlinescan9":              0x09f3,
	"leftt":                       0x09f4,
	"rightt":                      0x09f5,
	"bott":                        0x09f6,
	"topt":                        0x09f7,
	"vertbar":                     0x09f8,
	"emspace":                     0x0aa1,
	"enspace":                     0x0aa2,
	"em3space":                    0x0aa3,
	"em4space":                    0x0aa4,
	"digitspace":                  0x0aa5,
	"punctspace":                  0x0aa6,
	"thinspace":                   0x0aa7,
	"hairspace":                   0x0aa8,
	"emdash":                      0x0aa9,
	"endash":                      0x0aaa,
	"signifblank":                 0x0aac,
	"ellipsis":                    0x0aae,
	"doubbaselinedot":             0x0aaf,
	"onethird":                    0x0ab0,
	"twothirds":                   0x0ab1,
	"onefifth":                    0x0ab2,
	"twofifths":                   0x0ab3,
	"threefifths":                 0x0ab4,
	"fourfifths":                  0x0ab5,
	"onesixth":                    0x0ab6,
	"fivesixths":                  0x0ab7,
	"careof":                      0x0ab8,
	"figdash":                     0x0abb,
	"leftanglebracket":            0x0abc,
	"decimalpoint":                0x0abd,
	"rightanglebracket":           0x0abe,
	"marker":                      0x0abf,
	"oneeighth":                   0x0ac3,
	"threeeighths":                0x0ac4,
	"fiveeighths":                 0x0ac5,
	"seveneighths":                0x0ac6,
	"trademark":                   0x0ac9,
	"signaturemark":               0x0aca,
	"trademarkincircle":           0x0acb,
	"leftopentriangle":            0x0acc,
	"rightopentriangle":           0x0acd,
	"emopencircle":                0x0ace,
	"emopenrectangle":             0x0acf,
	"leftsinglequotemark":         0x0ad0,
	"rightsinglequotemark":        0x0ad1,
	"leftdoublequotemark":         0x0ad2,
	"rightdoublequotemark":        0x0ad3,
	"prescription":                0x0ad4,
	"minutes":                     0x0ad6,
	"seconds":                     0x0ad7,
	"latincross":                  0x0ad9,
	"hexagram":                    0x0ada,
	"filledrectbullet":            0x0adb,
	"filledlefttribullet":         0x0adc,
	"filledrighttribullet":        0x0add,
	"emfilledcircle":              0x0ade,
	"emfilledrect":                0x0adf,
	"enopencircbullet":            0x0ae0,
	"enopensquarebullet":          0x0ae1,
	"openrectbullet":              0x0ae2,
	"opentribulletup":             0x0ae3,
	"opentribulletdown":           0x0ae4,
	"openstar":                    0x0ae5,
	"enfilledcircbullet":          0x0ae6,
	"enfilledsqbullet":            0x0ae7,
	"filledtribulletup":           0x0ae8,
	"filledtribulletdown":         0x0ae9,
	"leftpointer":                 0x0aea,
	"rightpointer":                0x0aeb,
	"club":                        0x0aec,
	"diamond":                     0x0aed,
	"heart":                       0x0aee,
	"maltesecross":                0x0af0,
	"dagger":                      0x0af1,
	"doubledagger":                0x0af2,
	"checkmark":                   0x0af3,
	"ballotcross":                 0x0af4,
	"musicalsharp":                0x0af5,
	"musicalflat":                 0x0af6,
	"malesymbol":                  0x0af7,
	"femalesymbol":                0x0af8,
	"telephone":                   0x0af9,
	"telephonerecorder":           0x0afa,
	"phonographcopyright":         0x0afb,
	"caret":                       0x0afc,
	"singlelowquotemark":          0x0afd,
	"doublelowquotemark":          0x0afe,
	"cursor":                      0x0aff,
	"leftcaret":                   0x0ba3,
	"rightcaret":                  0x0ba6,
	"downcaret":                   0x0ba8,
	"upcaret":                     0x0ba9,
	"overbar":                     0x0bc0,
	"downtack":                    0x0bc2,
	"upshoe":                      0x0bc3,
	"downstile":                   0x0bc4,
	"underbar":                    0x0bc6,
	"jot":                         0x0bca,
	"quad":                        0x0bcc,
	"uptack":                      0x0bce,
	"circle":                      0x0bcf,
	"upstile":                     0x0bd3,
	"downshoe":                    0x0bd6,
	"rightshoe":                   0x0bd8,
	"leftshoe":                    0x0bda,
	"lefttack":                    0x0bdc,
	"righttack":                   0x0bfc,
	"hebrew_doublelowline":        0x0cdf,
	"hebrew_aleph":                0x0ce0,
	"hebrew_bet":                  0x0ce1,
	"hebrew_beth":                 0x0ce1,
	"hebrew_gimel":                0x0ce2,
	"hebrew_gimmel":               0x0ce2,
	"hebrew_dalet":                0x0ce3,
	"hebrew_daleth":               0x0ce3,
	"hebrew_he":                   0x0ce4,
	"hebrew_waw":                  0x0ce5,
	"hebrew_zain":                 0x0ce6,
	"hebrew_zayin":                0x0ce6,
	"hebrew_chet":                 0x0ce7,
	"hebrew_het":                  0x0ce7,
	"hebrew_tet":                  0x0ce8,
	"hebrew_teth":                 0x0ce8,
	"hebrew_yod":                  0x0ce9,
	"hebrew_finalkaph":            0x0cea,
	"hebrew_kaph":                 0x0ceb,
	"hebrew_lamed":                0x0cec,
	"hebrew_finalmem":             0x0ced,
	"hebrew_mem":                  0x0cee,
	"hebrew_finalnun":             0x0cef,
	"hebrew_nun":                  0x0cf0,
	"hebrew_samech":               0x0cf1,
	"hebrew_samekh":               0x0cf1,
	"hebrew_ayin":                 0x0cf2,
	"hebrew_finalpe":              0x0cf3,
	"hebrew_pe":                   0x0cf4,
	"hebrew_finalzade":            0x0cf5,
	"hebrew_finalzadi":            0x0cf5,
	"hebrew_zade":                 0x0cf6,
	"hebrew_zadi":                 0x0cf6,
	"hebrew_qoph":                 0x0cf7,
	"hebrew_kuf":                  0x0cf7,
	"hebrew_resh":                 0x0cf8,
	"hebrew_shin":                 0x0cf9,
	"hebrew_taw":                  0x0cfa,
	"hebrew_taf":                  0x0cfa,
	"Hebrew_switch":               0xff7e,
	"Thai_kokai":                  0x0da1,
	"Thai_khokhai":                0x0da2,
	"Thai_khokhuat":               0x0da3,
	"Thai_khokhwai":               0x0da4,
	"Thai_khokhon":                0x0da5,
	"Thai_khorakhang":             0x0da6,
	"Thai_ngongu":                 0x0da7,
	"Thai_chochan":                0x0da8,
	"Thai_choching":               0x0da9,
	"Thai_chochang":               0x0daa,
	"Thai_soso":                   0x0dab,
	"Thai_chochoe":                0x0dac,
	"Thai_yoying":                 0x0dad,
	"Thai_dochada":                0x0dae,
	"Thai_topatak":                0x0daf,
	"Thai_thothan":                0x0db0,
	"Thai_thonangmontho":          0x0db1,
	"Thai_thophuthao":             0x0db2,
	"Thai_nonen":                  0x0db3,
	"Thai_dodek":                  0x0db4,
	"Thai_totao":                  0x0db5,
	"Thai_thothung":               0x0db6,
	"Thai_thothahan":              0x0db7,
	"Thai_thothong":               0x0db8,
	"Thai_nonu":                   0x0db9,
	"Thai_bobaimai":               0x0dba,
	"Thai_popla":                  0x0dbb,
	"Thai_phophung":               0x0dbc,
	"Thai_fofa":                   0x0dbd,
	"Thai_phophan":                0x0dbe,
	"Thai_fofan":                  0x0dbf,
	"Thai_phosamphao":             0x0dc0,
	"Thai_moma":                   0x0dc1,
	"Thai_yoyak":                  0x0dc2,
	"Thai_rorua":                  0x0dc3,
	"Thai_ru":                     0x0dc4,
	"Thai_loling":                 0x0dc5,
	"Thai_lu":                     0x0dc6,
	"Thai_wowaen":                 0x0dc7,
	"Thai_sosala":                 0x0dc8,
	"Thai_sorusi":                 0x0dc9,
	"Thai_sosua":                  0x0dca,
	"Thai_hohip":                  0x0dcb,
	"Thai_lochula":                0x0dcc,
	"Thai_oang":                   0x0dcd,
	"Thai_honokhuk":               0x0dce,
	"Thai_paiyannoi":              0x0dcf,
	"Thai_saraa":                  0x0dd0,
	"Thai_maihanakat":             0x0dd1,
	"Thai_saraaa":                 0x0dd2,
	"Thai_saraam":                 0x0dd3,
	"Thai_sarai":                  0x0dd4,
	"Thai_saraii":                 0x0dd5,
	"Thai_saraue":                 0x0dd6,
	"Thai_sarauee":                0x0dd7,
	"Thai_sarau":                  0x0dd8,
	"Thai_sarauu":                 0x0dd9,
	"Thai_phinthu":                0x0dda,
	"Thai_maihanakat_maitho":      0x0dde,
	"Thai_baht":                   0x0ddf,
	"Thai_sarae":                  0x0de0,
	"Thai_saraae":                 0x0de1,
	"Thai_sarao":                  0x0de2,
	"Thai_saraaimaimuan":          0x0de3,
	"Thai_saraaimaimalai":         0x0de4,
	"Thai_lakkhangyao":            0x0de5,
	"Thai_maiyamok":               0x0de6,
	"Thai_maitaikhu":              0x0de7,
	"Thai_maiek":                  0x0de8,
	"Thai_maitho":                 0x0de9,
	"Thai_maitri":                 0x0dea,
	"Thai_maichattawa":            0x0deb,
	"Thai_thanthakhat":            0x0dec,
	"Thai_nikhahit":               0x0ded,
	"Thai_leksun":                 0x0df0,
	"Thai_leknung":                0x0df1,
	"Thai_leksong":                0x0df2,
	"Thai_leksam":                 0x0df3,
	"Thai_leksi":                  0x0df4,
	"Thai_lekha":                  0x0df5,
	"Thai_lekhok":                 0x0df6,
	"Thai_lekchet":                0x0df7,
	"Thai_lekpaet":                0x0df8,
	"Thai_lekkao":                 0x0df9,
	"Hangul":                      0xff31,
	"Hangul_Start":                0xff32,
	"Hangul_End":                  0xff33,
	"Hangul_Hanja":                0xff34,
	"Hangul_Jamo":                 0xff35,
	"Hangul_Romaja":               0xff36,
	"Hangul_Codeinput":            0xff37,
	"Hangul_Jeonja":               0xff38,
	"Hangul_Banja":                0xff39,
	"Hangul_PreHanja":             0xff3a,
	"Hangul_PostHanja":            0xff3b,
	"Hangul_SingleCandidate":      0xff3c,
	"Hangul_MultipleCandidate":    0xff3d,
	"Hangul_PreviousCandidate":    0xff3e,
	"Hangul_Special":              0xff3f,
	"Hangul_switch":               0xff7e,
	"Hangul_Kiyeog":               0x0ea1,
	"Hangul_SsangKiyeog":          0x0ea2,
	"Hangul_KiyeogSios":           0x0ea3,
	"Hangul_Nieun":                0x0ea4,
	"Hangul_NieunJieuj":           0x0ea5,
	"Hangul_NieunHieuh":           0x0ea6,
	"Hangul_Dikeud":               0x0ea7,
	"Hangul_SsangDikeud":          0x0ea8,
	"Hangul_Rieul":                0x0ea9,
	"Hangul_RieulKiyeog":          0x0eaa,
	"Hangul_RieulMieum":           0x0eab,
	"Hangul_RieulPieub":           0x0eac,
	"Hangul_RieulSios":            0x0ead,
	"Hangul_RieulTieut":           0x0eae,
	"Hangul_RieulPhieuf":          0x0eaf,
	"Hangul_RieulHieuh":           0x0eb0,
	"Hangul_Mieum":                0x0eb1,
	"Hangul_Pieub":                0x0eb2,
	"Hangul_SsangPieub":           0x0eb3,
	"Hangul_PieubSios":            0x0eb4,
	"Hangul_Sios":                 0x0eb5,
	"Hangul_SsangSios":            0x0eb6,
	"Hangul_Ieung":                0x0eb7,
	"Hangul_Jieuj":                0x0eb8,
	"Hangul_SsangJieuj":           0x0eb9,
	"Hangul_Cieuc":                0x0eba,
	"Hangul_Khieuq":               0x0ebb,
	"Hangul_Tieut":                0x0ebc,
	"Hangul_Phieuf":               0x0ebd,
	"Hangul_Hieuh":                0x0ebe,
	"Hangul_A":                    0x0ebf,
	"Hangul_AE":                   0x0ec0,
	"Hangul_YA":                   0x0ec1,
	"Hangul_YAE":                  0x0ec2,
	"Hangul_EO":                   0x0ec3,
	"Hangul_E":                    0x0ec4,
	"Hangul_YEO":                  0x0ec5,
	"Hangul_YE":                   0x0ec6,
	"Hangul_O":                    0x0ec7,
	"Hangul_WA":                   0x0ec8,
	"Hangul_WAE":                  0x0ec9,
	"Hangul_OE":                   0x0eca,
	"Hangul_YO":                   0x0ecb,
	"Hangul_U":                    0x0ecc,
	"Hangul_WEO":                  0x0ecd,
	"Hangul_WE":                   0x0ece,
	"Hangul_WI":                   0x0ecf,
	"Hangul_YU":                   0x0ed0,
	"Hangul_EU":                   0x0ed1,
	"Hangul_YI":                   0x0ed2,
	"Hangul_I":                    0x0ed3,
	"Hangul_J_Kiyeog":             0x0ed4,
	"Hangul_J_SsangKiyeog":        0x0ed5,
	"Hangul_J_KiyeogSios":         0x0ed6,
	"Hangul_J_Nieun":              0x0ed7,
	"Hangul_J_NieunJieuj":         0x0ed8,
	"Hangul_J_NieunHieuh":         0x0ed9,
	"Hangul_J_Dikeud":             0x0eda,
	"Hangul_J_Rieul":              0x0edb,
	"Hangul_J_RieulKiyeog":        0x0edc,
	"Hangul_J_RieulMieum":         0x0edd,
	"Hangul_J_RieulPieub":         0x0ede,
	"Hangul_J_RieulSios":          0x0edf,
	"Hangul_J_RieulTieut":         0x0ee0,
	"Hangul_J_RieulPhieuf":        0x0ee1,
	"Hangul_J_RieulHieuh":         0x0ee2,
	"Hangul_J_Mieum":              0x0ee3,
	"Hangul_J_Pieub":              0x0ee4,
	"Hangul_J_PieubSios":          0x0ee5,
	"Hangul_J_Sios":               0x0ee6,
	"Hangul_J_SsangSios":          0x0ee7,
	"Hangul_J_Ieung":              0x0ee8,
	"Hangul_J_Jieuj":              0x0ee9,
	"Hangul_J_Cieuc":              0x0eea,
	"Hangul_J_Khieuq":             0x0eeb,
	"Hangul_J_Tieut":              0x0eec,
	"Hangul_J_Phieuf":             0x0eed,
	"Hangul_J_Hieuh":              0x0eee,
	"Hangul_RieulYeorinHieuh":     0x0eef,
	"Hangul_SunkyeongeumMieum":    0x0ef0,
	"Hangul_SunkyeongeumPieub":    0x0ef1,
	"Hangul_PanSios":              0x0ef2,
	"Hangul_KkogjiDalrinIeung":    0x0ef3,
	"Hangul_SunkyeongeumPhieuf":   0x0ef4,
	"Hangul_YeorinHieuh":          0x0ef5,
	"Hangul_AraeA":                0x0ef6,
	"Hangul_AraeAE":               0x0ef7,
	"Hangul_J_PanSios":            0x0ef8,
	"Hangul_J_KkogjiDalrinIeung":  0x0ef9,
	"Hangul_J_YeorinHieuh":        0x0efa,
	"Korean_Won":                  0x0eff,
	"Armenian_ligature_ew":        0x1000587,
	"Armenian_full_stop":          0x1000589,
	"Armenian_verjaket":           0x1000589,
	"Armenian_separation_mark":    0x100055d,
	"Armenian_but":                0x100055d,
	"Armenian_hyphen":             0x100058a,
	"Armenian_yentamna":           0x100058a,
	"Armenian_exclam":             0x100055c,
	"Armenian_amanak":             0x100055c,
	"Armenian_accent":             0x100055b,
	"Armenian_shesht":             0x100055b,
	"Armenian_question":           0x100055e,
	"Armenian_paruyk":             0x100055e,
	"Armenian_AYB":                0x1000531,
	"Armenian_ayb":                0x1000561,
	"Armenian_BEN":                0x1000532,
	"Armenian_ben":                0x1000562,
	"Armenian_GIM":                0x1000533,
	"Armenian_gim":                0x1000563,
	"Armenian_DA":                 0x1000534,
	"Armenian_da":                 0x1000564,
	"Armenian_YECH":               0x1000535,
	"Armenian_yech":               0x1000565,
	"Armenian_ZA":                 0x1000536,
	"Armenian_za":                 0x1000566,
	"Armenian_E":                  0x1000537,
	"Armenian_e":                  0x1000567,
	"Armenian_AT":                 0x1000538,
	"Armenian_at":                 0x1000568,
	"Armenian_TO":                 0x1000539,
	"Armenian_to":                 0x1000569,
	"Armenian_ZHE":                0x100053a,
	"Armenian_zhe":                0x100056a,
	"Armenian_INI":                0x100053b,
	"Armenian_ini":                0x100056b,
	"Armenian_LYUN":               0x100053c,
	"Armenian_lyun":               0x100056c,
	"Armenian_KHE":                0x100053d,
	"Armenian_khe":                0x100056d,
	"Armenian_TSA":                0x100053e,
	"Armenian_tsa":                0x100056e,
	"Armenian_KEN":                0x100053f,
	"Armenian_ken":                0x100056f,
	"Armenian_HO":                 0x1000540,
	"Armenian_ho":                 0x1000570,
	"Armenian_DZA":                0x1000541,
	"Armenian_dza":                0x1000571,
	"Armenian_GHAT":               0x1000542,
	"Armenian_ghat":               0x1000572,
	"Armenian_TCHE":               0x1000543,
	"Armenian_tche":               0x1000573,
	"Armenian_MEN":                0x1000544,
	"Armenian_men":                0x1000574,
	"Armenian_HI":                 0x1000545,
	"Armenian_hi":                 0x1000575,
	"Armenian_NU":                 0x1000546,
	"Armenian_nu":                 0x1000576,
	"Armenian_SHA":                0x1000547,
	"Armenian_sha":                0x1000577,
	"Armenian_VO":                 0x1000548,
	"Armenian_vo":                 0x1000578,
	"Armenian_CHA":                0x1000549,
	"Armenian_cha":                0x1000579,
	"Armenian_PE":                 0x100054a,
	"Armenian_pe":                 0x100057a,
	"Armenian_JE":                 0x100054b,
	"Armenian_je":                 0x100057b,
	"Armenian_RA":                 0x100054c,
	"Armenian_ra":                 0x100057c,
	"Armenian_SE":                 0x100054d,
	"Armenian_se":                 0x100057d,
	"Armenian_VEV":                0x100054e,
	"Armenian_vev":                0x100057e,
	"Armenian_TYUN":               0x100054f,
	"Armenian_tyun":               0x100057f,
	"Armenian_RE":                 0x1000550,
	"Armenian_re":                 0x1000580,
	"Armenian_TSO":                0x1000551,
	"Armenian_tso":                0x1000581,
	"Armenian_VYUN":               0x1000552,
	"Armenian_vyun":               0x1000582,
	"Armenian_PYUR":               0x1000553,
	"Armenian_pyur":               0x1000583,
	"Armenian_KE":                 0x1000554,
	"Armenian_ke":                 0x1000584,
	"Armenian_O":                  0x1000555,
	"Armenian_o":                  0x1000585,
	"Armenian_FE":                 0x1000556,
	"Armenian_fe":                 0x1000586,
	"Armenian_apostrophe":         0x100055a,
	"Georgian_an":                 0x10010d0,
	"Georgian_ban":                0x10010d1,
	"Georgian_gan":                0x10010d2,
	"Georgian_don":                0x10010d3,
	"Georgian_en":                 0x10010d4,
	"Georgian_vin":                0x10010d5,
	"Georgian_zen":                0x10010d6,
	"Georgian_tan":                0x10010d7,
	"Georgian_in":                 0x10010d8,
	"Georgian_kan":                0x10010d9,
	"Georgian_las":                0x10010da,
	"Georgian_man":                0x10010db,
	"Georgian_nar":                0x10010dc,
	"Georgian_on":                 0x10010dd,
	"Georgian_par":                0x10010de,
	"Georgian_zhar":               0x10010df,
	"Georgian_rae":                0x10010e0,
	"Georgian_san":                0x10010e1,
	"Georgian_tar":                0x10010e2,
	"Georgian_un":                 0x10010e3,
	"Georgian_phar":               0x10010e4,
	"Georgian_khar":               0x10010e5,
	"Georgian_ghan":               0x10010e6,
	"Georgian_qar":                0x10010e7,
	"Georgian_shin":               0x10010e8,
	"Georgian_chin":               0x10010e9,
	"Georgian_can":                0x10010ea,
	"Georgian_jil":                0x10010eb,
	"Georgian_cil":                0x10010ec,
	"Georgian_char":               0x10010ed,
	"Georgian_xan":                0x10010ee,
	"Georgian_jhan":               0x10010ef,
	"Georgian_hae":                0x10010f0,
	"Georgian_he":                 0x10010f1,
	"Georgian_hie":                0x10010f2,
	"Georgian_we":                 0x10010f3,
	"Georgian_har":                0x10010f4,
	"Georgian_hoe":                0x10010f5,
	"Georgian_fi":                 0x10010f6,
	"Xabovedot":                   0x1001e8a,
	"Ibreve":                      0x100012c,
	"Zstroke":                     0x10001b5,
	"Gcaron":                      0x10001e6,
	"Ocaron":                      0x10001d1,
	"Obarred":                     0x100019f,
	"xabovedot":                   0x1001e8b,
	"ibreve":                      0x100012d,
	"zstroke":                     0x10001b6,
	"gcaron":                      0x10001e7,
	"ocaron":                      0x10001d2,
	"obarred":                     0x1000275,
	"SCHWA":                       0x100018f,
	"schwa":                       0x1000259,
	"Lbelowdot":                   0x1001e36,
	"lbelowdot":                   0x1001e37,
	"Abelowdot":                   0x1001ea0,
	"abelowdot":                   0x1001ea1,
	"Ahook":                       0x1001ea2,
	"ahook":                       0x1001ea3,
	"Acircumflexacute":            0x1001ea4,
	"acircumflexacute":            0x1001ea5,
	"Acircumflexgrave":            0x1001ea6,
	"acircumflexgrave":            0x1001ea7,
	"Acircumflexhook":             0x1001ea8,
	"acircumflexhook":             0x1001ea9,
	"Acircumflextilde":            0x1001eaa,
	"acircumflextilde":            0x1001eab,
	"Acircumflexbelowdot":         0x1001eac,
	"acircumflexbelowdot":         0x1001ead,
	"Abreveacute":                 0x1001eae,
	"abreveacute":                 0x1001eaf,
	"Abrevegrave":                 0x1001eb0,
	"abrevegrave":                 0x1001eb1,
	"Abrevehook":                  0x1001eb2,
	"abrevehook":                  0x1001eb3,
	"Abrevetilde":                 0x1001eb4,
	"abrevetilde":                 0x1001eb5,
	"Abrevebelowdot":              0x1001eb6,
	"abrevebelowdot":              0x1001eb7,
	"Ebelowdot":                   0x1001eb8,
	"ebelowdot":                   0x1001eb9,
	"Ehook":                       0x1001eba,
	"ehook":                       0x1001ebb,
	"Etilde":                      0x1001ebc,
	"etilde":                      0x1001ebd,
	"Ecircumflexacute":            0x1001ebe,
	"ecircumflexacute":            0x1001ebf,
	"Ecircumflexgrave":            0x1001ec0,
	"ecircumflexgrave":            0x1001ec1,
	"Ecircumflexhook":             0x1001ec2,
	"ecircumflexhook":             0x1001ec3,
	"Ecircumflextilde":            0x1001ec4,
	"ecircumflextilde":            0x1001ec5,
	"Ecircumflexbelowdot":         0x1001ec6,
	"ecircumflexbelowdot":         0x1001ec7,
	"Ihook":                       0x1001ec8,
	"ihook":                       0x1001ec9,
	"Ibelowdot":                   0x1001eca,
	"ibelowdot":                   0x1001ecb,
	"Obelowdot":                   0x1001ecc,
	"obelowdot":                   0x1001ecd,
	"Ohook":                       0x1001ece,
	"ohook":                       0x1001ecf,
	"Ocircumflexacute":            0x1001ed0,
	"ocircumflexacute":            0x1001ed1,
	"Ocircumflexgrave":            0x1001ed2,
	"ocircumflexgrave":            0x1001ed3,
	"Ocircumflexhook":             0x1001ed4,
	"ocircumflexhook":             0x1001ed5,
	"Ocircumflextilde":            0x1001ed6,
	"ocircumflextilde":            0x1001ed7,
	"Ocircumflexbelowdot":         0x1001ed8,
	"ocircumflexbelowdot":         0x1001ed9,
	"Ohornacute":                  0x1001eda,
	"ohornacute":                  0x1001edb,
	"Ohorngrave":                  0x1001edc,
	"ohorngrave":                  0x1001edd,
	"Ohornhook":                   0x1001ede,
	"ohornhook":                   0x1001edf,
	"Ohorntilde":                  0x1001ee0,
	"ohorntilde":                  0x1001ee1,
	"Ohornbelowdot":               0x1001ee2,
	"ohornbelowdot":               0x1001ee3,
	"Ubelowdot":                   0x1001ee4,
	"ubelowdot":                   0x1001ee5,
	"Uhook":                       0x1001ee6,
	"uhook":                       0x1001ee7,
	"Uhornacute":                  0x1001ee8,
	"uhornacute":                  0x1001ee9,
	"Uhorngrave":                  0x1001eea,
	"uhorngrave":                  0x1001eeb,
	"Uhornhook":                   0x1001eec,
	"uhornhook":                   0x1001eed,
	"Uhorntilde":                  0x1001eee,
	"uhorntilde":                  0x1001eef,
	"Uhornbelowdot":               0x1001ef0,
	"uhornbelowdot":               0x1001ef1,
	"Ybelowdot":                   0x1001ef4,
	"ybelowdot":                   0x1001ef5,
	"Yhook":                       0x1001ef6,
	"yhook":                       0x1001ef7,
	"Ytilde":                      0x1001ef8,
	"ytilde":                      0x1001ef9,
	"Ohorn":                       0x10001a0,
	"ohorn":                       0x10001a1,
	"Uhorn":                       0x10001af,
	"uhorn":                       0x10001b0,
	"EcuSign":                     0x10020a0,
	"ColonSign":                   0x10020a1,
	"CruzeiroSign":                0x10020a2,
	"FFrancSign":                  0x10020a3,
	"LiraSign":                    0x10020a4,
	"MillSign":                    0x10020a5,
	"NairaSign":                   0x10020a6,
	"PesetaSign":                  0x10020a7,
	"RupeeSign":                   0x10020a8,
	"WonSign":                     0x10020a9,
	"NewSheqelSign":               0x10020aa,
	"DongSign":                    0x10020ab,
	"EuroSign":                    0x20ac,
	"zerosuperior":                0x1002070,
	"foursuperior":                0x1002074,
	"fivesuperior":                0x1002075,
	"sixsuperior":                 0x1002076,
	"sevensuperior":               0x1002077,
	"eightsuperior":               0x1002078,
	"ninesuperior":                0x1002079,
	"zerosubscript":               0x1002080,
	"onesubscript":                0x1002081,
	"twosubscript":                0x1002082,
	"threesubscript":              0x1002083,
	"foursubscript":               0x1002084,
	"fivesubscript":               0x1002085,
	"sixsubscript":                0x1002086,
	"sevensubscript":              0x1002087,
	"eightsubscript":              0x1002088,
	"ninesubscript":               0x1002089,
	"partdifferential":            0x1002202,
	"emptyset":                    0x1002205,
	"elementof":                   0x1002208,
	"notelementof":                0x1002209,
	"containsas":                  0x100220B,
	"squareroot":                  0x100221A,
	"cuberoot":                    0x100221B,
	"fourthroot":                  0x100221C,
	"dintegral":                   0x100222C,
	"tintegral":                   0x100222D,
	"because":                     0x1002235,
	"approxeq":                    0x1002248,
	"notapproxeq":                 0x1002247,
	"notidentical":                0x1002262,
	"stricteq":                    0x1002263,
	"braille_dot_1":               0xfff1,
	"braille_dot_2":               0xfff2,
	"braille_dot_3":               0xfff3,
	"braille_dot_4":               0xfff4,
	"braille_dot_5":               0xfff5,
	"braille_dot_6":               0xfff6,
	"braille_dot_7":               0xfff7,
	"braille_dot_8":               0xfff8,
	"braille_dot_9":               0xfff9,
	"braille_dot_10":              0xfffa,
	"braille_blank":               0x1002800,
	"braille_dots_1":              0x1002801,
	"braille_dots_2":              0x1002802,
	"braille_dots_12":             0x1002803,
	"braille_dots_3":              0x1002804,
	"braille_dots_13":             0x1002805,
	"braille_dots_23":             0x1002806,
	"braille_dots_123":            0x1002807,
	"braille_dots_4":              0x1002808,
	"braille_dots_14":             0x1002809,
	"braille_dots_24":             0x100280a,
	"braille_dots_124":            0x100280b,
	"braille_dots_34":             0x100280c,
	"braille_dots_134":            0x100280d,
	"braille_dots_234":            0x100280e,
	"braille_dots_1234":           0x100280f,
	"braille_dots_5":              0x1002810,
	"braille_dots_15":             0x1002811,
	"braille_dots_25":             0x1002812,
	"braille_dots_125":            0x1002813,
	"braille_dots_35":             0x1002814,
	"braille_dots_135":            0x1002815,
	"braille_dots_235":            0x1002816,
	"braille_dots_1235":           0x1002817,
	"braille_dots_45":             0x1002818,
	"braille_dots_145":            0x1002819,
	"braille_dots_245":            0x100281a,
	"braille_dots_1245":           0x100281b,
	"braille_dots_345":            0x100281c,
	"braille_dots_1345":           0x100281d,
	"braille_dots_2345":           0x100281e,
	"braille_dots_12345":          0x100281f,
	"braille_dots_6":              0x1002820,
	"braille_dots_16":             0x1002821,
	"braille_dots_26":             0x1002822,
	"braille_dots_126":            0x1002823,
	"braille_dots_36":             0x1002824,
	"braille_dots_136":            0x1002825,
	"braille_dots_236":            0x1002826,
	"braille_dots_1236":           0x1002827,
	"braille_dots_46":             0x1002828,
	"braille_dots_146":            0x1002829,
	"braille_dots_246":            0x100282a,
	"braille_dots_1246":           0x100282b,
	"braille_dots_346":            0x100282c,
	"braille_dots_1346":           0x100282d,
	"braille_dots_2346":           0x100282e,
	"braille_dots_12346":          0x100282f,
	"braille_dots_56":             0x1002830,
	"braille_dots_156":            0x1002831,
	"braille_dots_256":            0x1002832,
	"braille_dots_1256":           0x1002833,
	"braille_dots_356":            0x1002834,
	"braille_dots_1356":           0x1002835,
	"braille_dots_2356":           0x1002836,
	"braille_dots_12356":          0x1002837,
	"braille_dots_456":            0x1002838,
	"braille_dots_1456":           0x1002839,
	"braille_dots_2456":           0x100283a,
	"braille_dots_12456":          0x100283b,
	"braille_dots_3456":           0x100283c,
	"braille_dots_13456":          0x100283d,
	"braille_dots_23456":          0x100283e,
	"braille_dots_123456":         0x100283f,
	"braille_dots_7":              0x1002840,
	"braille_dots_17":             0x1002841,
	"braille_dots_27":             0x1002842,
	"braille_dots_127":            0x1002843,
	"braille_dots_37":             0x1002844,
	"braille_dots_137":            0x1002845,
	"braille_dots_237":            0x1002846,
	"braille_dots_1237":           0x1002847,
	"braille_dots_47":             0x1002848,
	"braille_dots_147":            0x1002849,
	"braille_dots_247":            0x100284a,
	"braille_dots_1247":           0x100284b,
	"braille_dots_347":            0x100284c,
	"braille_dots_1347":           0x100284d,
	"braille_dots_2347":           0x100284e,
	"braille_dots_12347":          0x100284f,
	"braille_dots_57":             0x1002850,
	"braille_dots_157":            0x1002851,
	"braille_dots_257":            0x1002852,
	"braille_dots_1257":           0x1002853,
	"braille_dots_357":            0x1002854,
	"braille_dots_1357":           0x1002855,
	"braille_dots_2357":           0x1002856,
	"braille_dots_12357":          0x1002857,
	"braille_dots_457":            0x1002858,
	"braille_dots_1457":           0x1002859,
	"braille_dots_2457":           0x100285a,
	"braille_dots_12457":          0x100285b,
	"braille_dots_3457":           0x100285c,
	"braille_dots_13457":          0x100285d,
	"braille_dots_23457":          0x100285e,
	"braille_dots_123457":         0x100285f,
	"braille_dots_67":             0x1002860,
	"braille_dots_167":            0x1002861,
	"braille_dots_267":            0x1002862,
	"braille_dots_1267":           0x1002863,
	"braille_dots_367":            0x1002864,
	"braille_dots_1367":           0x1002865,
	"braille_dots_2367":           0x1002866,
	"braille_dots_12367":          0x1002867,
	"braille_dots_467":            0x1002868,
	"braille_dots_1467":           0x1002869,
	"braille_dots_2467":           0x100286a,
	"braille_dots_12467":          0x100286b,
	"braille_dots_3467":           0x100286c,
	"braille_dots_13467":          0x100286d,
	"braille_dots_23467":          0x100286e,
	"braille_dots_123467":         0x100286f,
	"braille_dots_567":            0x1002870,
	"braille_dots_1567":           0x1002871,
	"braille_dots_2567":           0x1002872,
	"braille_dots_12567":          0x1002873,
	"braille_dots_3567":           0x1002874,
	"braille_dots_13567":          0x1002875,
	"braille_dots_23567":          0x1002876,
	"braille_dots_123567":         0x1002877,
	"braille_dots_4567":           0x1002878,
	"braille_dots_14567":          0x1002879,
	"braille_dots_24567":          0x100287a,
	"braille_dots_124567":         0x100287b,
	"braille_dots_34567":          0x100287c,
	"braille_dots_134567":         0x100287d,
	"braille_dots_234567":         0x100287e,
	"braille_dots_1234567":        0x100287f,
	"braille_dots_8":              0x1002880,
	"braille_dots_18":             0x1002881,
	"braille_dots_28":             0x1002882,
	"braille_dots_128":            0x1002883,
	"braille_dots_38":             0x1002884,
	"braille_dots_138":            0x1002885,
	"braille_dots_238":            0x1002886,
	"braille_dots_1238":           0x1002887,
	"braille_dots_48":             0x1002888,
	"braille_dots_148":            0x1002889,
	"braille_dots_248":            0x100288a,
	"braille_dots_1248":           0x100288b,
	"braille_dots_348":            0x100288c,
	"braille_dots_1348":           0x100288d,
	"braille_dots_2348":           0x100288e,
	"braille_dots_12348":          0x100288f,
	"braille_dots_58":             0x1002890,
	"braille_dots_158":            0x1002891,
	"braille_dots_258":            0x1002892,
	"braille_dots_1258":           0x1002893,
	"braille_dots_358":            0x1002894,
	"braille_dots_1358":           0x1002895,
	"braille_dots_2358":           0x1002896,
	"braille_dots_12358":          0x1002897,
	"braille_dots_458":            0x1002898,
	"braille_dots_1458":           0x1002899,
	"braille_dots_2458":           0x100289a,
	"braille_dots_12458":          0x100289b,
	"braille_dots_3458":           0x100289c,
	"braille_dots_13458":          0x100289d,
	"braille_dots_23458":          0x100289e,
	"braille_dots_123458":         0x100289f,
	"braille_dots_68":             0x10028a0,
	"braille_dots_168":            0x10028a1,
	"braille_dots_268":            0x10028a2,
	"braille_dots_1268":           0x10028a3,
	"braille_dots_368":            0x10028a4,
	"braille_dots_1368":           0x10028a5,
	"braille_dots_2368":           0x10028a6,
	"braille_dots_12368":          0x10028a7,
	"braille_dots_468":            0x10028a8,
	"braille_dots_1468":           0x10028a9,
	"braille_dots_2468":           0x10028aa,
	"braille_dots_12468":          0x10028ab,
	"braille_dots_3468":           0x10028ac,
	"braille_dots_13468":          0x10028ad,
	"braille_dots_23468":          0x10028ae,
	"braille_dots_123468":         0x10028af,
	"braille_dots_568":            0x10028b0,
	"braille_dots_1568":           0x10028b1,
	"braille_dots_2568":           0x10028b2,
	"braille_dots_12568":          0x10028b3,
	"braille_dots_3568":           0x10028b4,
	"braille_dots_13568":          0x10028b5,
	"braille_dots_23568":          0x10028b6,
	"braille_dots_123568":         0x10028b7,
	"braille_dots_4568":           0x10028b8,
	"braille_dots_14568":          0x10028b9,
	"braille_dots_24568":          0x10028ba,
	"braille_dots_124568":         0x10028bb,
	"braille_dots_34568":          0x10028bc,
	"braille_dots_134568":         0x10028bd,
	"braille_dots_234568":         0x10028be,
	"braille_dots_1234568":        0x10028bf,
	"braille_dots_78":             0x10028c0,
	"braille_dots_178":            0x10028c1,
	"braille_dots_278":            0x10028c2,
	"braille_dots_1278":           0x10028c3,
	"braille_dots_378":            0x10028c4,
	"braille_dots_1378":           0x10028c5,
	"braille_dots_2378":           0x10028c6,
	"braille_dots_12378":          0x10028c7,
	"braille_dots_478":            0x10028c8,
	"braille_dots_1478":           0x10028c9,
	"braille_dots_2478":           0x10028ca,
	"braille_dots_12478":          0x10028cb,
	"braille_dots_3478":           0x10028cc,
	"braille_dots_13478":          0x10028cd,
	"braille_dots_23478":          0x10028ce,
	"braille_dots_123478":         0x10028cf,
	"braille_dots_578":            0x10028d0,
	"braille_dots_1578":           0x10028d1,
	"braille_dots_2578":           0x10028d2,
	"braille_dots_12578":          0x10028d3,
	"braille_dots_3578":           0x10028d4,
	"braille_dots_13578":          0x10028d5,
	"braille_dots_23578":          0x10028d6,
	"braille_dots_123578":         0x10028d7,
	"braille_dots_4578":           0x10028d8,
	"braille_dots_14578":          0x10028d9,
	"braille_dots_24578":          0x10028da,
	"braille_dots_124578":         0x10028db,
	"braille_dots_34578":          0x10028dc,
	"braille_dots_134578":         0x10028dd,
	"braille_dots_234578":         0x10028de,
	"braille_dots_1234578":        0x10028df,
	"braille_dots_678":            0x10028e0,
	"braille_dots_1678":           0x10028e1,
	"braille_dots_2678":           0x10028e2,
	"braille_dots_12678":          0x10028e3,
	"braille_dots_3678":           0x10028e4,
	"braille_dots_13678":          0x10028e5,
	"braille_dots_23678":          0x10028e6,
	"braille_dots_123678":         0x10028e7,
	"braille_dots_4678":           0x10028e8,
	"braille_dots_14678":          0x10028e9,
	"braille_dots_24678":          0x10028ea,
	"braille_dots_124678":         0x10028eb,
	"braille_dots_34678":          0x10028ec,
	"braille_dots_134678":         0x10028ed,
	"braille_dots_234678":         0x10028ee,
	"braille_dots_1234678":        0x10028ef,
	"braille_dots_5678":           0x10028f0,
	"braille_dots_15678":          0x10028f1,
	"braille_dots_25678":          0x10028f2,
	"braille_dots_125678":         0x10028f3,
	"braille_dots_35678":          0x10028f4,
	"braille_dots_135678":         0x10028f5,
	"braille_dots_235678":         0x10028f6,
	"braille_dots_1235678":        0x10028f7,
	"braille_dots_45678":          0x10028f8,
	"braille_dots_145678":         0x10028f9,
	"braille_dots_245678":         0x10028fa,
	"braille_dots_1245678":        0x10028fb,
	"braille_dots_345678":         0x10028fc,
	"braille_dots_1345678":        0x10028fd,
	"braille_dots_2345678":        0x10028fe,
	"braille_dots_12345678":       0x10028ff,

	"XF86ModeLock":          0x1008FF01,
	"XF86MonBrightnessUp":   0x1008FF02,
	"XF86MonBrightnessDown": 0x1008FF03,
	"XF86KbdLightOnOff":     0x1008FF04,
	"XF86KbdBrightnessUp":   0x1008FF05,
	"XF86KbdBrightnessDown": 0x1008FF06,
	"XF86Standby":           0x1008FF10,
	"XF86AudioLowerVolume":  0x1008FF11,
	"XF86AudioMute":         0x1008FF12,
	"XF86AudioRaiseVolume":  0x1008FF13,
	"XF86AudioPlay":         0x1008FF14,
	"XF86AudioStop":         0x1008FF15,
	"XF86AudioPrev":         0x1008FF16,
	"XF86AudioNext":         0x1008FF17,
	"XF86HomePage":          0x1008FF18,
	"XF86Mail":              0x1008FF19,
	"XF86Start":             0x1008FF1A,
	"XF86Search":            0x1008FF1B,
	"XF86AudioRecord":       0x1008FF1C,
	"XF86Calculator":        0x1008FF1D,
	"XF86Memo":              0x1008FF1E,
	"XF86ToDoList":          0x1008FF1F,
	"XF86Calendar":          0x1008FF20,
	"XF86PowerDown":         0x1008FF21,
	"XF86ContrastAdjust":    0x1008FF22,
	"XF86RockerUp":          0x1008FF23,
	"XF86RockerDown":        0x1008FF24,
	"XF86RockerEnter":       0x1008FF25,
	"XF86Back":              0x1008FF26,
	"XF86Forward":           0x1008FF27,
	"XF86Stop":              0x1008FF28,
	"XF86Refresh":           0x1008FF29,
	"XF86PowerOff":          0x1008FF2A,
	"XF86WakeUp":            0x1008FF2B,
	"XF86Eject":             0x1008FF2C,
	"XF86ScreenSaver":       0x1008FF2D,
	"XF86WWW":               0x1008FF2E,
	"XF86Sleep":             0x1008FF2F,
	"XF86Favorites":         0x1008FF30,
	"XF86AudioPause":        0x1008FF31,
	"XF86AudioMedia":        0x1008FF32,
	"XF86MyComputer":        0x1008FF33,
	"XF86VendorHome":        0x1008FF34,
	"XF86LightBulb":         0x1008FF35,
	"XF86Shop":              0x1008FF36,
	"XF86History":           0x1008FF37,
	"XF86OpenURL":           0x1008FF38,
	"XF86AddFavorite":       0x1008FF39,
	"XF86HotLinks":          0x1008FF3A,
	"XF86BrightnessAdjust":  0x1008FF3B,
	"XF86Finance":           0x1008FF3C,
	"XF86Community":         0x1008FF3D,
	"XF86AudioRewind":       0x1008FF3E,
	"XF86BackForward":       0x1008FF3F,
	"XF86Launch0":           0x1008FF40,
	"XF86Launch1":           0x1008FF41,
	"XF86Launch2":           0x1008FF42,
	"XF86Launch3":           0x1008FF43,
	"XF86Launch4":           0x1008FF44,
	"XF86Launch5":           0x1008FF45,
	"XF86Launch6":           0x1008FF46,
	"XF86Launch7":           0x1008FF47,
	"XF86Launch8":           0x1008FF48,
	"XF86Launch9":           0x1008FF49,
	"XF86LaunchA":           0x1008FF4A,
	"XF86LaunchB":           0x1008FF4B,
	"XF86LaunchC":           0x1008FF4C,
	"XF86LaunchD":           0x1008FF4D,
	"XF86LaunchE":           0x1008FF4E,
	"XF86LaunchF":           0x1008FF4F,
	"XF86ApplicationLeft":   0x1008FF50,
	"XF86ApplicationRight":  0x1008FF51,
	"XF86Book":              0x1008FF52,
	"XF86CD":                0x1008FF53,
	"XF86Calculater":        0x1008FF54,
	"XF86Clear":             0x1008FF55,
	"XF86Close":             0x1008FF56,
	"XF86Copy":              0x1008FF57,
	"XF86Cut":               0x1008FF58,
	"XF86Display":           0x1008FF59,
	"XF86DOS":               0x1008FF5A,
	"XF86Documents":         0x1008FF5B,
	"XF86Excel":             0x1008FF5C,
	"XF86Explorer":          0x1008FF5D,
	"XF86Game":              0x1008FF5E,
	"XF86Go":                0x1008FF5F,
	"XF86iTouch":            0x1008FF60,
	"XF86LogOff":            0x1008FF61,
	"XF86Market":            0x1008FF62,
	"XF86Meeting":           0x1008FF63,
	"XF86MenuKB":            0x1008FF65,
	"XF86MenuPB":            0x1008FF66,
	"XF86MySites":           0x1008FF67,
	"XF86New":               0x1008FF68,
	"XF86News":              0x1008FF69,
	"XF86OfficeHome":        0x1008FF6A,
	"XF86Open":              0x1008FF6B,
	"XF86Option":            0x1008FF6C,
	"XF86Paste":             0x1008FF6D,
	"XF86Phone":             0x1008FF6E,
	"XF86Q":                 0x1008FF70,
	"XF86Reply":             0x1008FF72,
	"XF86Reload":            0x1008FF73,
	"XF86RotateWindows":     0x1008FF74,
	"XF86RotationPB":        0x1008FF75,
	"XF86RotationKB":        0x1008FF76,
	"XF86Save":              0x1008FF77,
	"XF86ScrollUp":          0x1008FF78,
	"XF86ScrollDown":        0x1008FF79,
	"XF86ScrollClick":       0x1008FF7A,
	"XF86Send":              0x1008FF7B,
	"XF86Spell":             0x1008FF7C,
	"XF86SplitScreen":       0x1008FF7D,
	"XF86Support":           0x1008FF7E,
	"XF86TaskPane":          0x1008FF7F,
	"XF86Terminal":          0x1008FF80,
	"XF86Tools":             0x1008FF81,
	"XF86Travel":            0x1008FF82,
	"XF86UserPB":            0x1008FF84,
	"XF86User1KB":           0x1008FF85,
	"XF86User2KB":           0x1008FF86,
	"XF86Video":             0x1008FF87,
	"XF86WheelButton":       0x1008FF88,
	"XF86Word":              0x1008FF89,
	"XF86Xfer":              0x1008FF8A,
	"XF86ZoomIn":            0x1008FF8B,
	"XF86ZoomOut":           0x1008FF8C,
	"XF86Away":              0x1008FF8D,
	"XF86Messenger":         0x1008FF8E,
	"XF86WebCam":            0x1008FF8F,
	"XF86MailForward":       0x1008FF90,
	"XF86Pictures":          0x1008FF91,
	"XF86Music":             0x1008FF92,
	"XF86Battery":           0x1008FF93,
	"XF86Bluetooth":         0x1008FF94,
	"XF86WLAN":              0x1008FF95,
	"XF86UWB":               0x1008FF96,
	"XF86AudioForward":      0x1008FF97,
	"XF86AudioRepeat":       0x1008FF98,
	"XF86AudioRandomPlay":   0x1008FF99,
	"XF86Subtitle":          0x1008FF9A,
	"XF86AudioCycleTrack":   0x1008FF9B,
	"XF86CycleAngle":        0x1008FF9C,
	"XF86FrameBack":         0x1008FF9D,
	"XF86FrameForward":      0x1008FF9E,
	"XF86Time":              0x1008FF9F,
	"XF86Select":            0x1008FFA0,
	"XF86View":              0x1008FFA1,
	"XF86TopMenu":           0x1008FFA2,
	"XF86Red":               0x1008FFA3,
	"XF86Green":             0x1008FFA4,
	"XF86Yellow":            0x1008FFA5,
	"XF86Blue":              0x1008FFA6,
	"XF86Suspend":           0x1008FFA7,
	"XF86Hibernate":         0x1008FFA8,
	"XF86TouchpadToggle":    0x1008FFA9,
	"XF86TouchpadOn":        0x1008FFB0,
	"XF86TouchpadOff":       0x1008FFB1,
	"XF86AudioMicMute":      0x1008FFB2,
	"XF86Switch_VT_1":       0x1008FE01,
	"XF86Switch_VT_2":       0x1008FE02,
	"XF86Switch_VT_3":       0x1008FE03,
	"XF86Switch_VT_4":       0x1008FE04,
	"XF86Switch_VT_5":       0x1008FE05,
	"XF86Switch_VT_6":       0x1008FE06,
	"XF86Switch_VT_7":       0x1008FE07,
	"XF86Switch_VT_8":       0x1008FE08,
	"XF86Switch_VT_9":       0x1008FE09,
	"XF86Switch_VT_10":      0x1008FE0A,
	"XF86Switch_VT_11":      0x1008FE0B,
	"XF86Switch_VT_12":      0x1008FE0C,
	"XF86Ungrab":            0x1008FE20,
	"XF86ClearGrab":         0x1008FE21,
	"XF86Next_VMode":        0x1008FE22,
	"XF86Prev_VMode":        0x1008FE23,
	"XF86LogWindowTree":     0x1008FE24,
	"XF86LogGrabInfo":       0x1008FE25,
}

/*
# Mapping of X11 keysyms to ISO 10646 / Unicode
#
# The "X11 Window System Protocol" standard (Release 6.4) defines in
# Appendix A the keysym codes. These 29-bit integer values identify
# characters or functions associated with each key (e.g., via the
# visible engraving) of a keyboard layout. In addition, mnemonic macro
# names are provided for the keysyms in the C header file
# <X11/keysymdef.h>. These are compiled (by xc/lib/X11/util/
# makekeys.c) into a hash table that can be accessed with X11 library
# functions such as XStringToKeysym() and XKeysymToString().
#
# The creation of the keysym codes predates ISO 10646 / Unicode, but
# they represent a similar attempt to merge several existing coded
# character sets (mostly early drafts of ISO 8859, as well as some --
# long since forgotten -- DEC font encodings). X.Org and XFree86 have
# agreed that for any future extension of the keysyms with characters
# already found in ISO 10646 / Unicode, the following algorithm will
# be used. The new keysym code position will simply be the character's
# Unicode number plus 0x01000000. The keysym codes in the range
# 0x01000100 0x0110ffff are now reserved to represent Unicode
# characters in the range U0100 to U10FFFF. (Note that the ISO 8859-1
# characters that make up Unicode positions below U0100 are excluded
# from this rule, as they are already covered by the keysyms of the
# same value.)
#
# While most newer Unicode-based X11 clients do already accept
# Unicode-mapped keysyms in the range 0x01000100 to 0x0110ffff, it
# will remain necessary for clients -- in the interest of
# compatibility with existing servers -- to also understand the
# existing keysym values. Clients can use the table below to map the
# pre-Unicode keysym values (0x0100 to 0x20ff) to the corresponding
# Unicode characters for further processing.
#
# The following fields are used in this mapping table:
#
# 1    The hexadecimal X11 keysym number (as defined in Appendix A of
#      the X11 protocol specification and as listed in <X11/keysymdef.h>)
#
# 2    The corresponding Unicode position
#      (U0000 means that there is no equivalent Unicode character)
#
# 3    Status of this keysym and its Unicode mapping
#
#         .  regular -- This is a regular well-established keysym with
#            a straightforward Unicode equivalent (e.g., any keysym
#            derived from ISO 8859). There can be at most one regular
#            keysym associated with each Unicode character.
#
#         d  duplicate -- This keysym has the same Unicode mapping as
#            another one with status 'regular'. It represents a case
#            where keysyms distinguish between several characters that
#            Unicode has unified into a single one (examples are
#            several APL symbols)
#
#         o  obsolete -- While it may be possible to find a Unicode of
#            similar name, the exact semantics of this keysym are
#            unclear, because the font or character set from which it
#            came has never been widely used. Examples are various
#            symbols from the DEC Publishing character set, which may
#            have been used in a special font shipped with the
#            DECwrite product. Where no similar Unicode character
#            can be identified, U0000 is used in column 2.
#
#         f  function -- While it may be possible to find a Unicode
#            of similar name, this keysym differs semantically
#            substantially from the corresponding Unicode character,
#            because it describes a particular function key or will
#            first have to be processed by an input method that will
#            translate it into a proper stream of Unicode characters.
#
#         r  remove -- This is a bogus keysym that was added in error,
#            is not used in any known keyboard layout, and should be
#            removed from both <X11/keysymdef.h> and the standard.
#
#         u  unicode-remap -- This keysym was added rather recently to
#            the <X11/keysymdef.h> of XFree86, but into a number range
#            reserved for future extensions of the standard by
#            X.Org. It is not widely used at present, but its name
#            appears to be sufficiently useful and it should therefore
#            be directly mapped to Unicode in the 0x1xxxxxx range in
#            future versions of <X11/keysymdef.h>. This way, the macro
#            name will be preserved, but the standard will not have to
#            be extended.
#
#      Recommendations for using the keysym status:
#
#        - All keysyms with status regular, duplicate, obsolete and
#          function should be listed in Appendix A of the X11 protocol
#          spec.
#
#        - All keysyms except for those with status remove should be
#          listed in <X11/keysymdef.h>.
#
#        - Keysyms with status duplicate, obsolete, and remove should
#          not be used in future keyboard layouts, as there are other
#          keysyms with status regular, function and unicode-remap
#          that give access to the same Unicode characters.
#
#        - Keysym to Unicode conversion tables in clients should include
#          all mappings except those with status function and those
#          with U0000.
#
# #    comment marker
#
# 4    the name of the X11 keysym macro without the leading XK_,
#      as defined in <X11/keysymdef.h>
#
# The last columns may be followed by comments copied from <X11/keysymdef.h>.
# A keysym may be listed several times, if there are several macro names
# associated with it in <X11/keysymdef.h>.
#
# Author: Markus Kuhn <http://www.cl.cam.ac.uk/~mgk25/>
# Date:   2004-08-08
#
# This table evolved out of an earlier one by Richard Verhoeven, TU Eindhoven.
#
# We begin with the original keysyms found in X11R6.4
#
*/

var utf map[rune]xproto.Keysym = map[rune]xproto.Keysym{
	'\u0020': 0x0020, //   .   # space
	'\u0021': 0x0021, //   .   # exclam
	'\u0022': 0x0022, //   .   # quotedbl
	'\u0023': 0x0023, //   .   # numbersign
	'\u0024': 0x0024, //   .   # dollar
	'\u0025': 0x0025, //   .   # percent
	'\u0026': 0x0026, //   .   # ampersand
	'\u0027': 0x0027, //   .   # apostrophe
	//	'\u0027': 0x0027,   //   .   # quoteright	/* deprecated */
	'\u0028': 0x0028, //   .   # parenleft
	'\u0029': 0x0029, //   .   # parenright
	'\u002a': 0x002a, //   .   # asterisk
	'\u002b': 0x002b, //   .   # plus
	'\u002c': 0x002c, //   .   # comma
	'\u002d': 0x002d, //   .   # minus
	'\u002e': 0x002e, //   .   # period
	'\u002f': 0x002f, //   .   # slash
	'\u0030': 0x0030, //   .   # 0
	'\u0031': 0x0031, //   .   # 1
	'\u0032': 0x0032, //   .   # 2
	'\u0033': 0x0033, //   .   # 3
	'\u0034': 0x0034, //   .   # 4
	'\u0035': 0x0035, //   .   # 5
	'\u0036': 0x0036, //   .   # 6
	'\u0037': 0x0037, //   .   # 7
	'\u0038': 0x0038, //   .   # 8
	'\u0039': 0x0039, //   .   # 9
	'\u003a': 0x003a, //   .   # colon
	'\u003b': 0x003b, //   .   # semicolon
	'\u003c': 0x003c, //   .   # less
	'\u003d': 0x003d, //   .   # equal
	'\u003e': 0x003e, //   .   # greater
	'\u003f': 0x003f, //   .   # question
	'\u0040': 0x0040, //   .   # at
	'\u0041': 0x0041, //   .   # A
	'\u0042': 0x0042, //   .   # B
	'\u0043': 0x0043, //   .   # C
	'\u0044': 0x0044, //   .   # D
	'\u0045': 0x0045, //   .   # E
	'\u0046': 0x0046, //   .   # F
	'\u0047': 0x0047, //   .   # G
	'\u0048': 0x0048, //   .   # H
	'\u0049': 0x0049, //   .   # I
	'\u004a': 0x004a, //   .   # J
	'\u004b': 0x004b, //   .   # K
	'\u004c': 0x004c, //   .   # L
	'\u004d': 0x004d, //   .   # M
	'\u004e': 0x004e, //   .   # N
	'\u004f': 0x004f, //   .   # O
	'\u0050': 0x0050, //   .   # P
	'\u0051': 0x0051, //   .   # Q
	'\u0052': 0x0052, //   .   # R
	'\u0053': 0x0053, //   .   # S
	'\u0054': 0x0054, //   .   # T
	'\u0055': 0x0055, //   .   # U
	'\u0056': 0x0056, //   .   # V
	'\u0057': 0x0057, //   .   # W
	'\u0058': 0x0058, //   .   # X
	'\u0059': 0x0059, //   .   # Y
	'\u005a': 0x005a, //   .   # Z
	'\u005b': 0x005b, //   .   # bracketleft
	'\u005c': 0x005c, //   .   # backslash
	'\u005d': 0x005d, //   .   # bracketright
	'\u005e': 0x005e, //   .   # asciicircum
	'\u005f': 0x005f, //   .   # underscore
	'\u0060': 0x0060, //   .   # grave
	//'\u0060': 0x0060,   //   .   # quoteleft	/* deprecated */
	'\u0061': 0x0061, //   .   # a
	'\u0062': 0x0062, //   .   # b
	'\u0063': 0x0063, //   .   # c
	'\u0064': 0x0064, //   .   # d
	'\u0065': 0x0065, //   .   # e
	'\u0066': 0x0066, //   .   # f
	'\u0067': 0x0067, //   .   # g
	'\u0068': 0x0068, //   .   # h
	'\u0069': 0x0069, //   .   # i
	'\u006a': 0x006a, //   .   # j
	'\u006b': 0x006b, //   .   # k
	'\u006c': 0x006c, //   .   # l
	'\u006d': 0x006d, //   .   # m
	'\u006e': 0x006e, //   .   # n
	'\u006f': 0x006f, //   .   # o
	'\u0070': 0x0070, //   .   # p
	'\u0071': 0x0071, //   .   # q
	'\u0072': 0x0072, //   .   # r
	'\u0073': 0x0073, //   .   # s
	'\u0074': 0x0074, //   .   # t
	'\u0075': 0x0075, //   .   # u
	'\u0076': 0x0076, //   .   # v
	'\u0077': 0x0077, //   .   # w
	'\u0078': 0x0078, //   .   # x
	'\u0079': 0x0079, //   .   # y
	'\u007a': 0x007a, //   .   # z
	'\u007b': 0x007b, //   .   # braceleft
	'\u007c': 0x007c, //   .   # bar
	'\u007d': 0x007d, //   .   # braceright
	'\u007e': 0x007e, //   .   # asciitilde
	'\u00a0': 0x00a0, //   .   # nobreakspace
	'\u00a1': 0x00a1, //   .   # exclamdown
	'\u00a2': 0x00a2, //   .   # cent
	'\u00a3': 0x00a3, //   .   # sterling
	'\u00a4': 0x00a4, //   .   # currency
	'\u00a5': 0x00a5, //   .   # yen
	'\u00a6': 0x00a6, //   .   # brokenbar
	'\u00a7': 0x00a7, //   .   # section
	'\u00a8': 0x00a8, //   .   # diaeresis
	'\u00a9': 0x00a9, //   .   # copyright
	'\u00aa': 0x00aa, //   .   # ordfeminine
	'\u00ab': 0x00ab, //   .   # guillemotleft	/* left angle quotation mark */
	'\u00ac': 0x00ac, //   .   # notsign
	'\u00ad': 0x00ad, //   .   # hyphen
	'\u00ae': 0x00ae, //   .   # registered
	'\u00af': 0x00af, //   .   # macron
	'\u00b0': 0x00b0, //   .   # degree
	'\u00b1': 0x00b1, //   .   # plusminus
	'\u00b2': 0x00b2, //   .   # twosuperior
	'\u00b3': 0x00b3, //   .   # threesuperior
	'\u00b4': 0x00b4, //   .   # acute
	'\u00b5': 0x00b5, //   .   # mu
	'\u00b6': 0x00b6, //   .   # paragraph
	'\u00b7': 0x00b7, //   .   # periodcentered
	'\u00b8': 0x00b8, //   .   # cedilla
	'\u00b9': 0x00b9, //   .   # onesuperior
	'\u00ba': 0x00ba, //   .   # masculine
	'\u00bb': 0x00bb, //   .   # guillemotright	/* right angle quotation mark */
	'\u00bc': 0x00bc, //   .   # onequarter
	'\u00bd': 0x00bd, //   .   # onehalf
	'\u00be': 0x00be, //   .   # threequarters
	'\u00bf': 0x00bf, //   .   # questiondown
	'\u00c0': 0x00c0, //   .   # Agrave
	'\u00c1': 0x00c1, //   .   # Aacute
	'\u00c2': 0x00c2, //   .   # Acircumflex
	'\u00c3': 0x00c3, //   .   # Atilde
	'\u00c4': 0x00c4, //   .   # Adiaeresis
	'\u00c5': 0x00c5, //   .   # Aring
	'\u00c6': 0x00c6, //   .   # AE
	'\u00c7': 0x00c7, //   .   # Ccedilla
	'\u00c8': 0x00c8, //   .   # Egrave
	'\u00c9': 0x00c9, //   .   # Eacute
	'\u00ca': 0x00ca, //   .   # Ecircumflex
	'\u00cb': 0x00cb, //   .   # Ediaeresis
	'\u00cc': 0x00cc, //   .   # Igrave
	'\u00cd': 0x00cd, //   .   # Iacute
	'\u00ce': 0x00ce, //   .   # Icircumflex
	'\u00cf': 0x00cf, //   .   # Idiaeresis
	'\u00d0': 0x00d0, //   .   # ETH
	//'\u00d0': 0x00d0,   //   .   # Eth	/* deprecated */
	'\u00d1': 0x00d1, //   .   # Ntilde
	'\u00d2': 0x00d2, //   .   # Ograve
	'\u00d3': 0x00d3, //   .   # Oacute
	'\u00d4': 0x00d4, //   .   # Ocircumflex
	'\u00d5': 0x00d5, //   .   # Otilde
	'\u00d6': 0x00d6, //   .   # Odiaeresis
	'\u00d7': 0x00d7, //   .   # multiply
	'\u00d8': 0x00d8, //   .   # Ooblique
	'\u00d9': 0x00d9, //   .   # Ugrave
	'\u00da': 0x00da, //   .   # Uacute
	'\u00db': 0x00db, //   .   # Ucircumflex
	'\u00dc': 0x00dc, //   .   # Udiaeresis
	'\u00dd': 0x00dd, //   .   # Yacute
	'\u00de': 0x00de, //   .   # THORN
	//'\u00de': 0x00de,   //   .   # Thorn	/* deprecated */
	'\u00df': 0x00df, //   .   # ssharp
	'\u00e0': 0x00e0, //   .   # agrave
	'\u00e1': 0x00e1, //   .   # aacute
	'\u00e2': 0x00e2, //   .   # acircumflex
	'\u00e3': 0x00e3, //   .   # atilde
	'\u00e4': 0x00e4, //   .   # adiaeresis
	'\u00e5': 0x00e5, //   .   # aring
	'\u00e6': 0x00e6, //   .   # ae
	'\u00e7': 0x00e7, //   .   # ccedilla
	'\u00e8': 0x00e8, //   .   # egrave
	'\u00e9': 0x00e9, //   .   # eacute
	'\u00ea': 0x00ea, //   .   # ecircumflex
	'\u00eb': 0x00eb, //   .   # ediaeresis
	'\u00ec': 0x00ec, //   .   # igrave
	'\u00ed': 0x00ed, //   .   # iacute
	'\u00ee': 0x00ee, //   .   # icircumflex
	'\u00ef': 0x00ef, //   .   # idiaeresis
	'\u00f0': 0x00f0, //   .   # eth
	'\u00f1': 0x00f1, //   .   # ntilde
	'\u00f2': 0x00f2, //   .   # ograve
	'\u00f3': 0x00f3, //   .   # oacute
	'\u00f4': 0x00f4, //   .   # ocircumflex
	'\u00f5': 0x00f5, //   .   # otilde
	'\u00f6': 0x00f6, //   .   # odiaeresis
	'\u00f7': 0x00f7, //   .   # division
	'\u00f8': 0x00f8, //   .   # oslash
	'\u00f9': 0x00f9, //   .   # ugrave
	'\u00fa': 0x00fa, //   .   # uacute
	'\u00fb': 0x00fb, //   .   # ucircumflex
	'\u00fc': 0x00fc, //   .   # udiaeresis
	'\u00fd': 0x00fd, //   .   # yacute
	'\u00fe': 0x00fe, //   .   # thorn
	'\u00ff': 0x00ff, //   .   # ydiaeresis
	'\u0104': 0x01a1, //   .   # Aogonek
	'\u02d8': 0x01a2, //   .   # breve
	'\u0141': 0x01a3, //   .   # Lstroke
	'\u013d': 0x01a5, //   .   # Lcaron
	'\u015a': 0x01a6, //   .   # Sacute
	'\u0160': 0x01a9, //   .   # Scaron
	'\u015e': 0x01aa, //   .   # Scedilla
	'\u0164': 0x01ab, //   .   # Tcaron
	'\u0179': 0x01ac, //   .   # Zacute
	'\u017d': 0x01ae, //   .   # Zcaron
	'\u017b': 0x01af, //   .   # Zabovedot
	'\u0105': 0x01b1, //   .   # aogonek
	'\u02db': 0x01b2, //   .   # ogonek
	'\u0142': 0x01b3, //   .   # lstroke
	'\u013e': 0x01b5, //   .   # lcaron
	'\u015b': 0x01b6, //   .   # sacute
	'\u02c7': 0x01b7, //   .   # caron
	'\u0161': 0x01b9, //   .   # scaron
	'\u015f': 0x01ba, //   .   # scedilla
	'\u0165': 0x01bb, //   .   # tcaron
	'\u017a': 0x01bc, //   .   # zacute
	'\u02dd': 0x01bd, //   .   # doubleacute
	'\u017e': 0x01be, //   .   # zcaron
	'\u017c': 0x01bf, //   .   # zabovedot
	'\u0154': 0x01c0, //   .   # Racute
	'\u0102': 0x01c3, //   .   # Abreve
	'\u0139': 0x01c5, //   .   # Lacute
	'\u0106': 0x01c6, //   .   # Cacute
	'\u010c': 0x01c8, //   .   # Ccaron
	'\u0118': 0x01ca, //   .   # Eogonek
	'\u011a': 0x01cc, //   .   # Ecaron
	'\u010e': 0x01cf, //   .   # Dcaron
	'\u0110': 0x01d0, //   .   # Dstroke
	'\u0143': 0x01d1, //   .   # Nacute
	'\u0147': 0x01d2, //   .   # Ncaron
	'\u0150': 0x01d5, //   .   # Odoubleacute
	'\u0158': 0x01d8, //   .   # Rcaron
	'\u016e': 0x01d9, //   .   # Uring
	'\u0170': 0x01db, //   .   # Udoubleacute
	'\u0162': 0x01de, //   .   # Tcedilla
	'\u0155': 0x01e0, //   .   # racute
	'\u0103': 0x01e3, //   .   # abreve
	'\u013a': 0x01e5, //   .   # lacute
	'\u0107': 0x01e6, //   .   # cacute
	'\u010d': 0x01e8, //   .   # ccaron
	'\u0119': 0x01ea, //   .   # eogonek
	'\u011b': 0x01ec, //   .   # ecaron
	'\u010f': 0x01ef, //   .   # dcaron
	'\u0111': 0x01f0, //   .   # dstroke
	'\u0144': 0x01f1, //   .   # nacute
	'\u0148': 0x01f2, //   .   # ncaron
	'\u0151': 0x01f5, //   .   # odoubleacute
	'\u0159': 0x01f8, //   .   # rcaron
	'\u016f': 0x01f9, //   .   # uring
	'\u0171': 0x01fb, //   .   # udoubleacute
	'\u0163': 0x01fe, //   .   # tcedilla
	'\u02d9': 0x01ff, //   .   # abovedot
	'\u0126': 0x02a1, //   .   # Hstroke
	'\u0124': 0x02a6, //   .   # Hcircumflex
	'\u0130': 0x02a9, //   .   # Iabovedot
	'\u011e': 0x02ab, //   .   # Gbreve
	'\u0134': 0x02ac, //   .   # Jcircumflex
	'\u0127': 0x02b1, //   .   # hstroke
	'\u0125': 0x02b6, //   .   # hcircumflex
	'\u0131': 0x02b9, //   .   # idotless
	'\u011f': 0x02bb, //   .   # gbreve
	'\u0135': 0x02bc, //   .   # jcircumflex
	'\u010a': 0x02c5, //   .   # Cabovedot
	'\u0108': 0x02c6, //   .   # Ccircumflex
	'\u0120': 0x02d5, //   .   # Gabovedot
	'\u011c': 0x02d8, //   .   # Gcircumflex
	'\u016c': 0x02dd, //   .   # Ubreve
	'\u015c': 0x02de, //   .   # Scircumflex
	'\u010b': 0x02e5, //   .   # cabovedot
	'\u0109': 0x02e6, //   .   # ccircumflex
	'\u0121': 0x02f5, //   .   # gabovedot
	'\u011d': 0x02f8, //   .   # gcircumflex
	'\u016d': 0x02fd, //   .   # ubreve
	'\u015d': 0x02fe, //   .   # scircumflex
	'\u0138': 0x03a2, //   .   # kra
	'\u0156': 0x03a3, //   .   # Rcedilla
	'\u0128': 0x03a5, //   .   # Itilde
	'\u013b': 0x03a6, //   .   # Lcedilla
	'\u0112': 0x03aa, //   .   # Emacron
	'\u0122': 0x03ab, //   .   # Gcedilla
	'\u0166': 0x03ac, //   .   # Tslash
	'\u0157': 0x03b3, //   .   # rcedilla
	'\u0129': 0x03b5, //   .   # itilde
	'\u013c': 0x03b6, //   .   # lcedilla
	'\u0113': 0x03ba, //   .   # emacron
	'\u0123': 0x03bb, //   .   # gcedilla
	'\u0167': 0x03bc, //   .   # tslash
	'\u014a': 0x03bd, //   .   # ENG
	'\u014b': 0x03bf, //   .   # eng
	'\u0100': 0x03c0, //   .   # Amacron
	'\u012e': 0x03c7, //   .   # Iogonek
	'\u0116': 0x03cc, //   .   # Eabovedot
	'\u012a': 0x03cf, //   .   # Imacron
	'\u0145': 0x03d1, //   .   # Ncedilla
	'\u014c': 0x03d2, //   .   # Omacron
	'\u0136': 0x03d3, //   .   # Kcedilla
	'\u0172': 0x03d9, //   .   # Uogonek
	'\u0168': 0x03dd, //   .   # Utilde
	'\u016a': 0x03de, //   .   # Umacron
	'\u0101': 0x03e0, //   .   # amacron
	'\u012f': 0x03e7, //   .   # iogonek
	'\u0117': 0x03ec, //   .   # eabovedot
	'\u012b': 0x03ef, //   .   # imacron
	'\u0146': 0x03f1, //   .   # ncedilla
	'\u014d': 0x03f2, //   .   # omacron
	'\u0137': 0x03f3, //   .   # kcedilla
	'\u0173': 0x03f9, //   .   # uogonek
	'\u0169': 0x03fd, //   .   # utilde
	'\u016b': 0x03fe, //   .   # umacron
	'\u203e': 0x047e, //   .   # overline
	'\u3002': 0x04a1, //   .   # kana_fullstop
	'\u300c': 0x04a2, //   .   # kana_openingbracket
	'\u300d': 0x04a3, //   .   # kana_closingbracket
	'\u3001': 0x04a4, //   .   # kana_comma
	'\u30fb': 0x04a5, //   .   # kana_conjunctive
	'\u30f2': 0x04a6, //   .   # kana_WO
	'\u30a1': 0x04a7, //   .   # kana_a
	'\u30a3': 0x04a8, //   .   # kana_i
	'\u30a5': 0x04a9, //   .   # kana_u
	'\u30a7': 0x04aa, //   .   # kana_e
	'\u30a9': 0x04ab, //   .   # kana_o
	'\u30e3': 0x04ac, //   .   # kana_ya
	'\u30e5': 0x04ad, //   .   # kana_yu
	'\u30e7': 0x04ae, //   .   # kana_yo
	'\u30c3': 0x04af, //   .   # kana_tsu
	'\u30fc': 0x04b0, //   .   # prolongedsound
	'\u30a2': 0x04b1, //   .   # kana_A
	'\u30a4': 0x04b2, //   .   # kana_I
	'\u30a6': 0x04b3, //   .   # kana_U
	'\u30a8': 0x04b4, //   .   # kana_E
	'\u30aa': 0x04b5, //   .   # kana_O
	'\u30ab': 0x04b6, //   .   # kana_KA
	'\u30ad': 0x04b7, //   .   # kana_KI
	'\u30af': 0x04b8, //   .   # kana_KU
	'\u30b1': 0x04b9, //   .   # kana_KE
	'\u30b3': 0x04ba, //   .   # kana_KO
	'\u30b5': 0x04bb, //   .   # kana_SA
	'\u30b7': 0x04bc, //   .   # kana_SHI
	'\u30b9': 0x04bd, //   .   # kana_SU
	'\u30bb': 0x04be, //   .   # kana_SE
	'\u30bd': 0x04bf, //   .   # kana_SO
	'\u30bf': 0x04c0, //   .   # kana_TA
	'\u30c1': 0x04c1, //   .   # kana_CHI
	'\u30c4': 0x04c2, //   .   # kana_TSU
	'\u30c6': 0x04c3, //   .   # kana_TE
	'\u30c8': 0x04c4, //   .   # kana_TO
	'\u30ca': 0x04c5, //   .   # kana_NA
	'\u30cb': 0x04c6, //   .   # kana_NI
	'\u30cc': 0x04c7, //   .   # kana_NU
	'\u30cd': 0x04c8, //   .   # kana_NE
	'\u30ce': 0x04c9, //   .   # kana_NO
	'\u30cf': 0x04ca, //   .   # kana_HA
	'\u30d2': 0x04cb, //   .   # kana_HI
	'\u30d5': 0x04cc, //   .   # kana_FU
	'\u30d8': 0x04cd, //   .   # kana_HE
	'\u30db': 0x04ce, //   .   # kana_HO
	'\u30de': 0x04cf, //   .   # kana_MA
	'\u30df': 0x04d0, //   .   # kana_MI
	'\u30e0': 0x04d1, //   .   # kana_MU
	'\u30e1': 0x04d2, //   .   # kana_ME
	'\u30e2': 0x04d3, //   .   # kana_MO
	'\u30e4': 0x04d4, //   .   # kana_YA
	'\u30e6': 0x04d5, //   .   # kana_YU
	'\u30e8': 0x04d6, //   .   # kana_YO
	'\u30e9': 0x04d7, //   .   # kana_RA
	'\u30ea': 0x04d8, //   .   # kana_RI
	'\u30eb': 0x04d9, //   .   # kana_RU
	'\u30ec': 0x04da, //   .   # kana_RE
	'\u30ed': 0x04db, //   .   # kana_RO
	'\u30ef': 0x04dc, //   .   # kana_WA
	'\u30f3': 0x04dd, //   .   # kana_N
	'\u309b': 0x04de, //   .   # voicedsound
	'\u309c': 0x04df, //   .   # semivoicedsound
	'\u060c': 0x05ac, //   .   # Arabic_comma
	'\u061b': 0x05bb, //   .   # Arabic_semicolon
	'\u061f': 0x05bf, //   .   # Arabic_question_mark
	'\u0621': 0x05c1, //   .   # Arabic_hamza
	'\u0622': 0x05c2, //   .   # Arabic_maddaonalef
	'\u0623': 0x05c3, //   .   # Arabic_hamzaonalef
	'\u0624': 0x05c4, //   .   # Arabic_hamzaonwaw
	'\u0625': 0x05c5, //   .   # Arabic_hamzaunderalef
	'\u0626': 0x05c6, //   .   # Arabic_hamzaonyeh
	'\u0627': 0x05c7, //   .   # Arabic_alef
	'\u0628': 0x05c8, //   .   # Arabic_beh
	'\u0629': 0x05c9, //   .   # Arabic_tehmarbuta
	'\u062a': 0x05ca, //   .   # Arabic_teh
	'\u062b': 0x05cb, //   .   # Arabic_theh
	'\u062c': 0x05cc, //   .   # Arabic_jeem
	'\u062d': 0x05cd, //   .   # Arabic_hah
	'\u062e': 0x05ce, //   .   # Arabic_khah
	'\u062f': 0x05cf, //   .   # Arabic_dal
	'\u0630': 0x05d0, //   .   # Arabic_thal
	'\u0631': 0x05d1, //   .   # Arabic_ra
	'\u0632': 0x05d2, //   .   # Arabic_zain
	'\u0633': 0x05d3, //   .   # Arabic_seen
	'\u0634': 0x05d4, //   .   # Arabic_sheen
	'\u0635': 0x05d5, //   .   # Arabic_sad
	'\u0636': 0x05d6, //   .   # Arabic_dad
	'\u0637': 0x05d7, //   .   # Arabic_tah
	'\u0638': 0x05d8, //   .   # Arabic_zah
	'\u0639': 0x05d9, //   .   # Arabic_ain
	'\u063a': 0x05da, //   .   # Arabic_ghain
	'\u0640': 0x05e0, //   .   # Arabic_tatweel
	'\u0641': 0x05e1, //   .   # Arabic_feh
	'\u0642': 0x05e2, //   .   # Arabic_qaf
	'\u0643': 0x05e3, //   .   # Arabic_kaf
	'\u0644': 0x05e4, //   .   # Arabic_lam
	'\u0645': 0x05e5, //   .   # Arabic_meem
	'\u0646': 0x05e6, //   .   # Arabic_noon
	'\u0647': 0x05e7, //   .   # Arabic_ha
	'\u0648': 0x05e8, //   .   # Arabic_waw
	'\u0649': 0x05e9, //   .   # Arabic_alefmaksura
	'\u064a': 0x05ea, //   .   # Arabic_yeh
	'\u064b': 0x05eb, //   .   # Arabic_fathatan
	'\u064c': 0x05ec, //   .   # Arabic_dammatan
	'\u064d': 0x05ed, //   .   # Arabic_kasratan
	'\u064e': 0x05ee, //   .   # Arabic_fatha
	'\u064f': 0x05ef, //   .   # Arabic_damma
	'\u0650': 0x05f0, //   .   # Arabic_kasra
	'\u0651': 0x05f1, //   .   # Arabic_shadda
	'\u0652': 0x05f2, //   .   # Arabic_sukun
	'\u0452': 0x06a1, //   .   # Serbian_dje
	'\u0453': 0x06a2, //   .   # Macedonia_gje
	'\u0451': 0x06a3, //   .   # Cyrillic_io
	'\u0454': 0x06a4, //   .   # Ukrainian_ie
	'\u0455': 0x06a5, //   .   # Macedonia_dse
	'\u0456': 0x06a6, //   .   # Ukrainian_i
	'\u0457': 0x06a7, //   .   # Ukrainian_yi
	'\u0458': 0x06a8, //   .   # Cyrillic_je
	'\u0459': 0x06a9, //   .   # Cyrillic_lje
	'\u045a': 0x06aa, //   .   # Cyrillic_nje
	'\u045b': 0x06ab, //   .   # Serbian_tshe
	'\u045c': 0x06ac, //   .   # Macedonia_kje
	'\u045e': 0x06ae, //   .   # Byelorussian_shortu
	'\u045f': 0x06af, //   .   # Cyrillic_dzhe
	'\u2116': 0x06b0, //   .   # numerosign
	'\u0402': 0x06b1, //   .   # Serbian_DJE
	'\u0403': 0x06b2, //   .   # Macedonia_GJE
	'\u0401': 0x06b3, //   .   # Cyrillic_IO
	'\u0404': 0x06b4, //   .   # Ukrainian_IE
	'\u0405': 0x06b5, //   .   # Macedonia_DSE
	'\u0406': 0x06b6, //   .   # Ukrainian_I
	'\u0407': 0x06b7, //   .   # Ukrainian_YI
	'\u0408': 0x06b8, //   .   # Cyrillic_JE
	'\u0409': 0x06b9, //   .   # Cyrillic_LJE
	'\u040a': 0x06ba, //   .   # Cyrillic_NJE
	'\u040b': 0x06bb, //   .   # Serbian_TSHE
	'\u040c': 0x06bc, //   .   # Macedonia_KJE
	'\u040e': 0x06be, //   .   # Byelorussian_SHORTU
	'\u040f': 0x06bf, //   .   # Cyrillic_DZHE
	'\u044e': 0x06c0, //   .   # Cyrillic_yu
	'\u0430': 0x06c1, //   .   # Cyrillic_a
	'\u0431': 0x06c2, //   .   # Cyrillic_be
	'\u0446': 0x06c3, //   .   # Cyrillic_tse
	'\u0434': 0x06c4, //   .   # Cyrillic_de
	'\u0435': 0x06c5, //   .   # Cyrillic_ie
	'\u0444': 0x06c6, //   .   # Cyrillic_ef
	'\u0433': 0x06c7, //   .   # Cyrillic_ghe
	'\u0445': 0x06c8, //   .   # Cyrillic_ha
	'\u0438': 0x06c9, //   .   # Cyrillic_i
	'\u0439': 0x06ca, //   .   # Cyrillic_shorti
	'\u043a': 0x06cb, //   .   # Cyrillic_ka
	'\u043b': 0x06cc, //   .   # Cyrillic_el
	'\u043c': 0x06cd, //   .   # Cyrillic_em
	'\u043d': 0x06ce, //   .   # Cyrillic_en
	'\u043e': 0x06cf, //   .   # Cyrillic_o
	'\u043f': 0x06d0, //   .   # Cyrillic_pe
	'\u044f': 0x06d1, //   .   # Cyrillic_ya
	'\u0440': 0x06d2, //   .   # Cyrillic_er
	'\u0441': 0x06d3, //   .   # Cyrillic_es
	'\u0442': 0x06d4, //   .   # Cyrillic_te
	'\u0443': 0x06d5, //   .   # Cyrillic_u
	'\u0436': 0x06d6, //   .   # Cyrillic_zhe
	'\u0432': 0x06d7, //   .   # Cyrillic_ve
	'\u044c': 0x06d8, //   .   # Cyrillic_softsign
	'\u044b': 0x06d9, //   .   # Cyrillic_yeru
	'\u0437': 0x06da, //   .   # Cyrillic_ze
	'\u0448': 0x06db, //   .   # Cyrillic_sha
	'\u044d': 0x06dc, //   .   # Cyrillic_e
	'\u0449': 0x06dd, //   .   # Cyrillic_shcha
	'\u0447': 0x06de, //   .   # Cyrillic_che
	'\u044a': 0x06df, //   .   # Cyrillic_hardsign
	'\u042e': 0x06e0, //   .   # Cyrillic_YU
	'\u0410': 0x06e1, //   .   # Cyrillic_A
	'\u0411': 0x06e2, //   .   # Cyrillic_BE
	'\u0426': 0x06e3, //   .   # Cyrillic_TSE
	'\u0414': 0x06e4, //   .   # Cyrillic_DE
	'\u0415': 0x06e5, //   .   # Cyrillic_IE
	'\u0424': 0x06e6, //   .   # Cyrillic_EF
	'\u0413': 0x06e7, //   .   # Cyrillic_GHE
	'\u0425': 0x06e8, //   .   # Cyrillic_HA
	'\u0418': 0x06e9, //   .   # Cyrillic_I
	'\u0419': 0x06ea, //   .   # Cyrillic_SHORTI
	'\u041a': 0x06eb, //   .   # Cyrillic_KA
	'\u041b': 0x06ec, //   .   # Cyrillic_EL
	'\u041c': 0x06ed, //   .   # Cyrillic_EM
	'\u041d': 0x06ee, //   .   # Cyrillic_EN
	'\u041e': 0x06ef, //   .   # Cyrillic_O
	'\u041f': 0x06f0, //   .   # Cyrillic_PE
	'\u042f': 0x06f1, //   .   # Cyrillic_YA
	'\u0420': 0x06f2, //   .   # Cyrillic_ER
	'\u0421': 0x06f3, //   .   # Cyrillic_ES
	'\u0422': 0x06f4, //   .   # Cyrillic_TE
	'\u0423': 0x06f5, //   .   # Cyrillic_U
	'\u0416': 0x06f6, //   .   # Cyrillic_ZHE
	'\u0412': 0x06f7, //   .   # Cyrillic_VE
	'\u042c': 0x06f8, //   .   # Cyrillic_SOFTSIGN
	'\u042b': 0x06f9, //   .   # Cyrillic_YERU
	'\u0417': 0x06fa, //   .   # Cyrillic_ZE
	'\u0428': 0x06fb, //   .   # Cyrillic_SHA
	'\u042d': 0x06fc, //   .   # Cyrillic_E
	'\u0429': 0x06fd, //   .   # Cyrillic_SHCHA
	'\u0427': 0x06fe, //   .   # Cyrillic_CHE
	'\u042a': 0x06ff, //   .   # Cyrillic_HARDSIGN
	'\u0386': 0x07a1, //   .   # Greek_ALPHAaccent
	'\u0388': 0x07a2, //   .   # Greek_EPSILONaccent
	'\u0389': 0x07a3, //   .   # Greek_ETAaccent
	'\u038a': 0x07a4, //   .   # Greek_IOTAaccent
	'\u03aa': 0x07a5, //   .   # Greek_IOTAdiaeresis
	'\u038c': 0x07a7, //   .   # Greek_OMICRONaccent
	'\u038e': 0x07a8, //   .   # Greek_UPSILONaccent
	'\u03ab': 0x07a9, //   .   # Greek_UPSILONdieresis
	'\u038f': 0x07ab, //   .   # Greek_OMEGAaccent
	'\u0385': 0x07ae, //   .   # Greek_accentdieresis
	'\u2015': 0x07af, //   .   # Greek_horizbar
	'\u03ac': 0x07b1, //   .   # Greek_alphaaccent
	'\u03ad': 0x07b2, //   .   # Greek_epsilonaccent
	'\u03ae': 0x07b3, //   .   # Greek_etaaccent
	'\u03af': 0x07b4, //   .   # Greek_iotaaccent
	'\u03ca': 0x07b5, //   .   # Greek_iotadieresis
	'\u0390': 0x07b6, //   .   # Greek_iotaaccentdieresis
	'\u03cc': 0x07b7, //   .   # Greek_omicronaccent
	'\u03cd': 0x07b8, //   .   # Greek_upsilonaccent
	'\u03cb': 0x07b9, //   .   # Greek_upsilondieresis
	'\u03b0': 0x07ba, //   .   # Greek_upsilonaccentdieresis
	'\u03ce': 0x07bb, //   .   # Greek_omegaaccent
	'\u0391': 0x07c1, //   .   # Greek_ALPHA
	'\u0392': 0x07c2, //   .   # Greek_BETA
	'\u0393': 0x07c3, //   .   # Greek_GAMMA
	'\u0394': 0x07c4, //   .   # Greek_DELTA
	'\u0395': 0x07c5, //   .   # Greek_EPSILON
	'\u0396': 0x07c6, //   .   # Greek_ZETA
	'\u0397': 0x07c7, //   .   # Greek_ETA
	'\u0398': 0x07c8, //   .   # Greek_THETA
	'\u0399': 0x07c9, //   .   # Greek_IOTA
	'\u039a': 0x07ca, //   .   # Greek_KAPPA
	'\u039b': 0x07cb, //   .   # Greek_LAMBDA
	//'\u039b': 0x07cb,   //   .   # Greek_LAMDA
	'\u039c': 0x07cc, //   .   # Greek_MU
	'\u039d': 0x07cd, //   .   # Greek_NU
	'\u039e': 0x07ce, //   .   # Greek_XI
	'\u039f': 0x07cf, //   .   # Greek_OMICRON
	'\u03a0': 0x07d0, //   .   # Greek_PI
	'\u03a1': 0x07d1, //   .   # Greek_RHO
	'\u03a3': 0x07d2, //   .   # Greek_SIGMA
	'\u03a4': 0x07d4, //   .   # Greek_TAU
	'\u03a5': 0x07d5, //   .   # Greek_UPSILON
	'\u03a6': 0x07d6, //   .   # Greek_PHI
	'\u03a7': 0x07d7, //   .   # Greek_CHI
	'\u03a8': 0x07d8, //   .   # Greek_PSI
	'\u03a9': 0x07d9, //   .   # Greek_OMEGA
	'\u03b1': 0x07e1, //   .   # Greek_alpha
	'\u03b2': 0x07e2, //   .   # Greek_beta
	'\u03b3': 0x07e3, //   .   # Greek_gamma
	'\u03b4': 0x07e4, //   .   # Greek_delta
	'\u03b5': 0x07e5, //   .   # Greek_epsilon
	'\u03b6': 0x07e6, //   .   # Greek_zeta
	'\u03b7': 0x07e7, //   .   # Greek_eta
	'\u03b8': 0x07e8, //   .   # Greek_theta
	'\u03b9': 0x07e9, //   .   # Greek_iota
	'\u03ba': 0x07ea, //   .   # Greek_kappa
	'\u03bb': 0x07eb, //   .   # Greek_lambda
	'\u03bc': 0x07ec, //   .   # Greek_mu
	'\u03bd': 0x07ed, //   .   # Greek_nu
	'\u03be': 0x07ee, //   .   # Greek_xi
	'\u03bf': 0x07ef, //   .   # Greek_omicron
	'\u03c0': 0x07f0, //   .   # Greek_pi
	'\u03c1': 0x07f1, //   .   # Greek_rho
	'\u03c3': 0x07f2, //   .   # Greek_sigma
	'\u03c2': 0x07f3, //   .   # Greek_finalsmallsigma
	'\u03c4': 0x07f4, //   .   # Greek_tau
	'\u03c5': 0x07f5, //   .   # Greek_upsilon
	'\u03c6': 0x07f6, //   .   # Greek_phi
	'\u03c7': 0x07f7, //   .   # Greek_chi
	'\u03c8': 0x07f8, //   .   # Greek_psi
	'\u03c9': 0x07f9, //   .   # Greek_omega
	'\u23b7': 0x08a1, //   .   # leftradical
	'\u250c': 0x08a2, //   d   # topleftradical
	'\u2500': 0x08a3, //   d   # horizconnector
	'\u2320': 0x08a4, //   .   # topintegral
	'\u2321': 0x08a5, //   .   # botintegral
	'\u2502': 0x08a6, //   d   # vertconnector
	'\u23a1': 0x08a7, //   .   # topleftsqbracket
	'\u23a3': 0x08a8, //   .   # botleftsqbracket
	'\u23a4': 0x08a9, //   .   # toprightsqbracket
	'\u23a6': 0x08aa, //   .   # botrightsqbracket
	'\u239b': 0x08ab, //   .   # topleftparens
	'\u239d': 0x08ac, //   .   # botleftparens
	'\u239e': 0x08ad, //   .   # toprightparens
	'\u23a0': 0x08ae, //   .   # botrightparens
	'\u23a8': 0x08af, //   .   # leftmiddlecurlybrace
	'\u23ac': 0x08b0, //   .   # rightmiddlecurlybrace
	'\u0000': 0x08b1, //   o   # topleftsummation
	//'\u0000': 0x08b2,   //   o   # botleftsummation
	//'\u0000': 0x08b3,   //   o   # topvertsummationconnector
	//'\u0000': 0x08b4,   //   o   # botvertsummationconnector
	//'\u0000': 0x08b5,   //   o   # toprightsummation
	//'\u0000': 0x08b6,   //   o   # botrightsummation
	//'\u0000': 0x08b7,   //   o   # rightmiddlesummation
	'\u2264': 0x08bc, //   .   # lessthanequal
	'\u2260': 0x08bd, //   .   # notequal
	'\u2265': 0x08be, //   .   # greaterthanequal
	'\u222b': 0x08bf, //   .   # integral
	'\u2234': 0x08c0, //   .   # therefore
	'\u221d': 0x08c1, //   .   # variation
	'\u221e': 0x08c2, //   .   # infinity
	'\u2207': 0x08c5, //   .   # nabla
	'\u223c': 0x08c8, //   .   # approximate
	'\u2243': 0x08c9, //   .   # similarequal
	'\u21d4': 0x08cd, //   .   # ifonlyif
	'\u21d2': 0x08ce, //   .   # implies
	'\u2261': 0x08cf, //   .   # identical
	'\u221a': 0x08d6, //   .   # radical
	'\u2282': 0x08da, //   .   # includedin
	'\u2283': 0x08db, //   .   # includes
	'\u2229': 0x08dc, //   .   # intersection
	'\u222a': 0x08dd, //   .   # union
	'\u2227': 0x08de, //   .   # logicaland
	'\u2228': 0x08df, //   .   # logicalor
	'\u2202': 0x08ef, //   .   # partialderivative
	'\u0192': 0x08f6, //   .   # function
	'\u2190': 0x08fb, //   .   # leftarrow
	'\u2191': 0x08fc, //   .   # uparrow
	'\u2192': 0x08fd, //   .   # rightarrow
	'\u2193': 0x08fe, //   .   # downarrow
	//'\u0000': 0x09df,   //   o   # blank
	'\u25c6': 0x09e0, //   .   # soliddiamond
	'\u2592': 0x09e1, //   .   # checkerboard
	'\u2409': 0x09e2, //   .   # ht
	'\u240c': 0x09e3, //   .   # ff
	'\u240d': 0x09e4, //   .   # cr
	'\u240a': 0x09e5, //   .   # lf
	'\u2424': 0x09e8, //   .   # nl
	'\u240b': 0x09e9, //   .   # vt
	'\u2518': 0x09ea, //   .   # lowrightcorner
	'\u2510': 0x09eb, //   .   # uprightcorner
	//'\u250c': 0x09ec,   //   .   # upleftcorner
	'\u2514': 0x09ed, //   .   # lowleftcorner
	'\u253c': 0x09ee, //   .   # crossinglines
	'\u23ba': 0x09ef, //   .   # horizlinescan1
	'\u23bb': 0x09f0, //   .   # horizlinescan3
	//'\u2500': 0x09f1,   //   .   # horizlinescan5
	'\u23bc': 0x09f2, //   .   # horizlinescan7
	'\u23bd': 0x09f3, //   .   # horizlinescan9
	'\u251c': 0x09f4, //   .   # leftt
	'\u2524': 0x09f5, //   .   # rightt
	'\u2534': 0x09f6, //   .   # bott
	'\u252c': 0x09f7, //   .   # topt
	//'\u2502': 0x09f8,   //   .   # vertbar
	'\u2003': 0x0aa1, //   .   # emspace
	'\u2002': 0x0aa2, //   .   # enspace
	'\u2004': 0x0aa3, //   .   # em3space
	'\u2005': 0x0aa4, //   .   # em4space
	'\u2007': 0x0aa5, //   .   # digitspace
	'\u2008': 0x0aa6, //   .   # punctspace
	'\u2009': 0x0aa7, //   .   # thinspace
	'\u200a': 0x0aa8, //   .   # hairspace
	'\u2014': 0x0aa9, //   .   # emdash
	'\u2013': 0x0aaa, //   .   # endash
	'\u2423': 0x0aac, //   o   # signifblank
	'\u2026': 0x0aae, //   .   # ellipsis
	'\u2025': 0x0aaf, //   .   # doubbaselinedot
	'\u2153': 0x0ab0, //   .   # onethird
	'\u2154': 0x0ab1, //   .   # twothirds
	'\u2155': 0x0ab2, //   .   # onefifth
	'\u2156': 0x0ab3, //   .   # twofifths
	'\u2157': 0x0ab4, //   .   # threefifths
	'\u2158': 0x0ab5, //   .   # fourfifths
	'\u2159': 0x0ab6, //   .   # onesixth
	'\u215a': 0x0ab7, //   .   # fivesixths
	'\u2105': 0x0ab8, //   .   # careof
	'\u2012': 0x0abb, //   .   # figdash
	'\u27e8': 0x0abc, //   o   # leftanglebracket
	//'\u002e': 0x0abd,   //   o   # decimalpoint
	'\u27e9': 0x0abe, //   o   # rightanglebracket
	//'\u0000': 0x0abf,   //   o   # marker
	'\u215b': 0x0ac3, //   .   # oneeighth
	'\u215c': 0x0ac4, //   .   # threeeighths
	'\u215d': 0x0ac5, //   .   # fiveeighths
	'\u215e': 0x0ac6, //   .   # seveneighths
	'\u2122': 0x0ac9, //   .   # trademark
	'\u2613': 0x0aca, //   o   # signaturemark
	//'\u0000': 0x0acb,   //   o   # trademarkincircle
	'\u25c1': 0x0acc, //   o   # leftopentriangle
	'\u25b7': 0x0acd, //   o   # rightopentriangle
	'\u25cb': 0x0ace, //   o   # emopencircle
	'\u25af': 0x0acf, //   o   # emopenrectangle
	'\u2018': 0x0ad0, //   .   # leftsinglequotemark
	'\u2019': 0x0ad1, //   .   # rightsinglequotemark
	'\u201c': 0x0ad2, //   .   # leftdoublequotemark
	'\u201d': 0x0ad3, //   .   # rightdoublequotemark
	'\u211e': 0x0ad4, //   .   # prescription
	'\u2032': 0x0ad6, //   .   # minutes
	'\u2033': 0x0ad7, //   .   # seconds
	'\u271d': 0x0ad9, //   .   # latincross
	//'\u0000': 0x0ada,   //   o   # hexagram
	'\u25ac': 0x0adb, //   o   # filledrectbullet
	'\u25c0': 0x0adc, //   o   # filledlefttribullet
	'\u25b6': 0x0add, //   o   # filledrighttribullet
	'\u25cf': 0x0ade, //   o   # emfilledcircle
	'\u25ae': 0x0adf, //   o   # emfilledrect
	'\u25e6': 0x0ae0, //   o   # enopencircbullet
	'\u25ab': 0x0ae1, //   o   # enopensquarebullet
	'\u25ad': 0x0ae2, //   o   # openrectbullet
	'\u25b3': 0x0ae3, //   o   # opentribulletup
	'\u25bd': 0x0ae4, //   o   # opentribulletdown
	'\u2606': 0x0ae5, //   o   # openstar
	'\u2022': 0x0ae6, //   o   # enfilledcircbullet
	'\u25aa': 0x0ae7, //   o   # enfilledsqbullet
	'\u25b2': 0x0ae8, //   o   # filledtribulletup
	'\u25bc': 0x0ae9, //   o   # filledtribulletdown
	'\u261c': 0x0aea, //   o   # leftpointer
	'\u261e': 0x0aeb, //   o   # rightpointer
	'\u2663': 0x0aec, //   .   # club
	'\u2666': 0x0aed, //   .   # diamond
	'\u2665': 0x0aee, //   .   # heart
	'\u2720': 0x0af0, //   .   # maltesecross
	'\u2020': 0x0af1, //   .   # dagger
	'\u2021': 0x0af2, //   .   # doubledagger
	'\u2713': 0x0af3, //   .   # checkmark
	'\u2717': 0x0af4, //   .   # ballotcross
	'\u266f': 0x0af5, //   .   # musicalsharp
	'\u266d': 0x0af6, //   .   # musicalflat
	'\u2642': 0x0af7, //   .   # malesymbol
	'\u2640': 0x0af8, //   .   # femalesymbol
	'\u260e': 0x0af9, //   .   # telephone
	'\u2315': 0x0afa, //   .   # telephonerecorder
	'\u2117': 0x0afb, //   .   # phonographcopyright
	'\u2038': 0x0afc, //   .   # caret
	'\u201a': 0x0afd, //   .   # singlelowquotemark
	'\u201e': 0x0afe, //   .   # doublelowquotemark
	//'\u0000': 0x0aff,   //   o   # cursor
	//'\u003c': 0x0ba3,   //   d   # leftcaret
	//'\u003e': 0x0ba6,   //   d   # rightcaret
	//'\u2228': 0x0ba8,   //   d   # downcaret
	//'\u2227': 0x0ba9,   //   d   # upcaret
	//'\u00af': 0x0bc0,   //   d   # overbar
	'\u22a5': 0x0bc2, //   .   # downtack
	//'\u2229': 0x0bc3,   //   d   # upshoe
	'\u230a': 0x0bc4, //   .   # downstile
	//'\u005f': 0x0bc6,   //   d   # underbar
	'\u2218': 0x0bca, //   .   # jot
	'\u2395': 0x0bcc, //   .   # quad
	'\u22a4': 0x0bce, //   .   # uptack
	//'\u25cb': 0x0bcf,   //   .   # circle
	'\u2308': 0x0bd3, //   .   # upstile
	//'\u222a': 0x0bd6,   //   d   # downshoe
	//'\u2283': 0x0bd8,   //   d   # rightshoe
	//'\u2282': 0x0bda,   //   d   # leftshoe
	'\u22a2': 0x0bdc, //   .   # lefttack
	'\u22a3': 0x0bfc, //   .   # righttack
	'\u2017': 0x0cdf, //   .   # hebrew_doublelowline
	'\u05d0': 0x0ce0, //   .   # hebrew_aleph
	'\u05d1': 0x0ce1, //   .   # hebrew_bet
	//'\u05d1': 0x0ce1,   //   .   # hebrew_beth  /* deprecated */
	'\u05d2': 0x0ce2, //   .   # hebrew_gimel
	//'\u05d2': 0x0ce2,   //   .   # hebrew_gimmel  /* deprecated */
	'\u05d3': 0x0ce3, //   .   # hebrew_dalet
	//'\u05d3': 0x0ce3,   //   .   # hebrew_daleth  /* deprecated */
	'\u05d4': 0x0ce4, //   .   # hebrew_he
	'\u05d5': 0x0ce5, //   .   # hebrew_waw
	'\u05d6': 0x0ce6, //   .   # hebrew_zain
	//'\u05d6': 0x0ce6,   //   .   # hebrew_zayin  /* deprecated */
	'\u05d7': 0x0ce7, //   .   # hebrew_chet
	//'\u05d7': 0x0ce7,   //   .   # hebrew_het  /* deprecated */
	'\u05d8': 0x0ce8, //   .   # hebrew_tet
	//'\u05d8': 0x0ce8,   //   .   # hebrew_teth  /* deprecated */
	'\u05d9': 0x0ce9, //   .   # hebrew_yod
	'\u05da': 0x0cea, //   .   # hebrew_finalkaph
	'\u05db': 0x0ceb, //   .   # hebrew_kaph
	'\u05dc': 0x0cec, //   .   # hebrew_lamed
	'\u05dd': 0x0ced, //   .   # hebrew_finalmem
	'\u05de': 0x0cee, //   .   # hebrew_mem
	'\u05df': 0x0cef, //   .   # hebrew_finalnun
	'\u05e0': 0x0cf0, //   .   # hebrew_nun
	'\u05e1': 0x0cf1, //   .   # hebrew_samech
	//'\u05e1': 0x0cf1,   //   .   # hebrew_samekh  /* deprecated */
	'\u05e2': 0x0cf2, //   .   # hebrew_ayin
	'\u05e3': 0x0cf3, //   .   # hebrew_finalpe
	'\u05e4': 0x0cf4, //   .   # hebrew_pe
	'\u05e5': 0x0cf5, //   .   # hebrew_finalzade
	//'\u05e5': 0x0cf5,   //   .   # hebrew_finalzadi  /* deprecated */
	'\u05e6': 0x0cf6, //   .   # hebrew_zade
	//'\u05e6': 0x0cf6,   //   .   # hebrew_zadi  /* deprecated */
	'\u05e7': 0x0cf7, //   .   # hebrew_kuf  /* deprecated */
	//'\u05e7': 0x0cf7,   //   .   # hebrew_qoph
	'\u05e8': 0x0cf8, //   .   # hebrew_resh
	'\u05e9': 0x0cf9, //   .   # hebrew_shin
	'\u05ea': 0x0cfa, //   .   # hebrew_taf  /* deprecated */
	//'\u05ea': 0x0cfa,   //   .   # hebrew_taw
	'\u0e01': 0x0da1, //   .   # Thai_kokai
	'\u0e02': 0x0da2, //   .   # Thai_khokhai
	'\u0e03': 0x0da3, //   .   # Thai_khokhuat
	'\u0e04': 0x0da4, //   .   # Thai_khokhwai
	'\u0e05': 0x0da5, //   .   # Thai_khokhon
	'\u0e06': 0x0da6, //   .   # Thai_khorakhang
	'\u0e07': 0x0da7, //   .   # Thai_ngongu
	'\u0e08': 0x0da8, //   .   # Thai_chochan
	'\u0e09': 0x0da9, //   .   # Thai_choching
	'\u0e0a': 0x0daa, //   .   # Thai_chochang
	'\u0e0b': 0x0dab, //   .   # Thai_soso
	'\u0e0c': 0x0dac, //   .   # Thai_chochoe
	'\u0e0d': 0x0dad, //   .   # Thai_yoying
	'\u0e0e': 0x0dae, //   .   # Thai_dochada
	'\u0e0f': 0x0daf, //   .   # Thai_topatak
	'\u0e10': 0x0db0, //   .   # Thai_thothan
	'\u0e11': 0x0db1, //   .   # Thai_thonangmontho
	'\u0e12': 0x0db2, //   .   # Thai_thophuthao
	'\u0e13': 0x0db3, //   .   # Thai_nonen
	'\u0e14': 0x0db4, //   .   # Thai_dodek
	'\u0e15': 0x0db5, //   .   # Thai_totao
	'\u0e16': 0x0db6, //   .   # Thai_thothung
	'\u0e17': 0x0db7, //   .   # Thai_thothahan
	'\u0e18': 0x0db8, //   .   # Thai_thothong
	'\u0e19': 0x0db9, //   .   # Thai_nonu
	'\u0e1a': 0x0dba, //   .   # Thai_bobaimai
	'\u0e1b': 0x0dbb, //   .   # Thai_popla
	'\u0e1c': 0x0dbc, //   .   # Thai_phophung
	'\u0e1d': 0x0dbd, //   .   # Thai_fofa
	'\u0e1e': 0x0dbe, //   .   # Thai_phophan
	'\u0e1f': 0x0dbf, //   .   # Thai_fofan
	'\u0e20': 0x0dc0, //   .   # Thai_phosamphao
	'\u0e21': 0x0dc1, //   .   # Thai_moma
	'\u0e22': 0x0dc2, //   .   # Thai_yoyak
	'\u0e23': 0x0dc3, //   .   # Thai_rorua
	'\u0e24': 0x0dc4, //   .   # Thai_ru
	'\u0e25': 0x0dc5, //   .   # Thai_loling
	'\u0e26': 0x0dc6, //   .   # Thai_lu
	'\u0e27': 0x0dc7, //   .   # Thai_wowaen
	'\u0e28': 0x0dc8, //   .   # Thai_sosala
	'\u0e29': 0x0dc9, //   .   # Thai_sorusi
	'\u0e2a': 0x0dca, //   .   # Thai_sosua
	'\u0e2b': 0x0dcb, //   .   # Thai_hohip
	'\u0e2c': 0x0dcc, //   .   # Thai_lochula
	'\u0e2d': 0x0dcd, //   .   # Thai_oang
	'\u0e2e': 0x0dce, //   .   # Thai_honokhuk
	'\u0e2f': 0x0dcf, //   .   # Thai_paiyannoi
	'\u0e30': 0x0dd0, //   .   # Thai_saraa
	'\u0e31': 0x0dd1, //   .   # Thai_maihanakat
	'\u0e32': 0x0dd2, //   .   # Thai_saraaa
	'\u0e33': 0x0dd3, //   .   # Thai_saraam
	'\u0e34': 0x0dd4, //   .   # Thai_sarai
	'\u0e35': 0x0dd5, //   .   # Thai_saraii
	'\u0e36': 0x0dd6, //   .   # Thai_saraue
	'\u0e37': 0x0dd7, //   .   # Thai_sarauee
	'\u0e38': 0x0dd8, //   .   # Thai_sarau
	'\u0e39': 0x0dd9, //   .   # Thai_sarauu
	'\u0e3a': 0x0dda, //   .   # Thai_phinthu
	//'\u0000': 0x0dde,   //   o   # Thai_maihanakat_maitho
	'\u0e3f': 0x0ddf, //   .   # Thai_baht
	'\u0e40': 0x0de0, //   .   # Thai_sarae
	'\u0e41': 0x0de1, //   .   # Thai_saraae
	'\u0e42': 0x0de2, //   .   # Thai_sarao
	'\u0e43': 0x0de3, //   .   # Thai_saraaimaimuan
	'\u0e44': 0x0de4, //   .   # Thai_saraaimaimalai
	'\u0e45': 0x0de5, //   .   # Thai_lakkhangyao
	'\u0e46': 0x0de6, //   .   # Thai_maiyamok
	'\u0e47': 0x0de7, //   .   # Thai_maitaikhu
	'\u0e48': 0x0de8, //   .   # Thai_maiek
	'\u0e49': 0x0de9, //   .   # Thai_maitho
	'\u0e4a': 0x0dea, //   .   # Thai_maitri
	'\u0e4b': 0x0deb, //   .   # Thai_maichattawa
	'\u0e4c': 0x0dec, //   .   # Thai_thanthakhat
	'\u0e4d': 0x0ded, //   .   # Thai_nikhahit
	'\u0e50': 0x0df0, //   .   # Thai_leksun
	'\u0e51': 0x0df1, //   .   # Thai_leknung
	'\u0e52': 0x0df2, //   .   # Thai_leksong
	'\u0e53': 0x0df3, //   .   # Thai_leksam
	'\u0e54': 0x0df4, //   .   # Thai_leksi
	'\u0e55': 0x0df5, //   .   # Thai_lekha
	'\u0e56': 0x0df6, //   .   # Thai_lekhok
	'\u0e57': 0x0df7, //   .   # Thai_lekchet
	'\u0e58': 0x0df8, //   .   # Thai_lekpaet
	'\u0e59': 0x0df9, //   .   # Thai_lekkao
	'\u3131': 0x0ea1, //   f   # Hangul_Kiyeog
	'\u3132': 0x0ea2, //   f   # Hangul_SsangKiyeog
	'\u3133': 0x0ea3, //   f   # Hangul_KiyeogSios
	'\u3134': 0x0ea4, //   f   # Hangul_Nieun
	'\u3135': 0x0ea5, //   f   # Hangul_NieunJieuj
	'\u3136': 0x0ea6, //   f   # Hangul_NieunHieuh
	'\u3137': 0x0ea7, //   f   # Hangul_Dikeud
	'\u3138': 0x0ea8, //   f   # Hangul_SsangDikeud
	'\u3139': 0x0ea9, //   f   # Hangul_Rieul
	'\u313a': 0x0eaa, //   f   # Hangul_RieulKiyeog
	'\u313b': 0x0eab, //   f   # Hangul_RieulMieum
	'\u313c': 0x0eac, //   f   # Hangul_RieulPieub
	'\u313d': 0x0ead, //   f   # Hangul_RieulSios
	'\u313e': 0x0eae, //   f   # Hangul_RieulTieut
	'\u313f': 0x0eaf, //   f   # Hangul_RieulPhieuf
	'\u3140': 0x0eb0, //   f   # Hangul_RieulHieuh
	'\u3141': 0x0eb1, //   f   # Hangul_Mieum
	'\u3142': 0x0eb2, //   f   # Hangul_Pieub
	'\u3143': 0x0eb3, //   f   # Hangul_SsangPieub
	'\u3144': 0x0eb4, //   f   # Hangul_PieubSios
	'\u3145': 0x0eb5, //   f   # Hangul_Sios
	'\u3146': 0x0eb6, //   f   # Hangul_SsangSios
	'\u3147': 0x0eb7, //   f   # Hangul_Ieung
	'\u3148': 0x0eb8, //   f   # Hangul_Jieuj
	'\u3149': 0x0eb9, //   f   # Hangul_SsangJieuj
	'\u314a': 0x0eba, //   f   # Hangul_Cieuc
	'\u314b': 0x0ebb, //   f   # Hangul_Khieuq
	'\u314c': 0x0ebc, //   f   # Hangul_Tieut
	'\u314d': 0x0ebd, //   f   # Hangul_Phieuf
	'\u314e': 0x0ebe, //   f   # Hangul_Hieuh
	'\u314f': 0x0ebf, //   f   # Hangul_A
	'\u3150': 0x0ec0, //   f   # Hangul_AE
	'\u3151': 0x0ec1, //   f   # Hangul_YA
	'\u3152': 0x0ec2, //   f   # Hangul_YAE
	'\u3153': 0x0ec3, //   f   # Hangul_EO
	'\u3154': 0x0ec4, //   f   # Hangul_E
	'\u3155': 0x0ec5, //   f   # Hangul_YEO
	'\u3156': 0x0ec6, //   f   # Hangul_YE
	'\u3157': 0x0ec7, //   f   # Hangul_O
	'\u3158': 0x0ec8, //   f   # Hangul_WA
	'\u3159': 0x0ec9, //   f   # Hangul_WAE
	'\u315a': 0x0eca, //   f   # Hangul_OE
	'\u315b': 0x0ecb, //   f   # Hangul_YO
	'\u315c': 0x0ecc, //   f   # Hangul_U
	'\u315d': 0x0ecd, //   f   # Hangul_WEO
	'\u315e': 0x0ece, //   f   # Hangul_WE
	'\u315f': 0x0ecf, //   f   # Hangul_WI
	'\u3160': 0x0ed0, //   f   # Hangul_YU
	'\u3161': 0x0ed1, //   f   # Hangul_EU
	'\u3162': 0x0ed2, //   f   # Hangul_YI
	'\u3163': 0x0ed3, //   f   # Hangul_I
	'\u11a8': 0x0ed4, //   f   # Hangul_J_Kiyeog
	'\u11a9': 0x0ed5, //   f   # Hangul_J_SsangKiyeog
	'\u11aa': 0x0ed6, //   f   # Hangul_J_KiyeogSios
	'\u11ab': 0x0ed7, //   f   # Hangul_J_Nieun
	'\u11ac': 0x0ed8, //   f   # Hangul_J_NieunJieuj
	'\u11ad': 0x0ed9, //   f   # Hangul_J_NieunHieuh
	'\u11ae': 0x0eda, //   f   # Hangul_J_Dikeud
	'\u11af': 0x0edb, //   f   # Hangul_J_Rieul
	'\u11b0': 0x0edc, //   f   # Hangul_J_RieulKiyeog
	'\u11b1': 0x0edd, //   f   # Hangul_J_RieulMieum
	'\u11b2': 0x0ede, //   f   # Hangul_J_RieulPieub
	'\u11b3': 0x0edf, //   f   # Hangul_J_RieulSios
	'\u11b4': 0x0ee0, //   f   # Hangul_J_RieulTieut
	'\u11b5': 0x0ee1, //   f   # Hangul_J_RieulPhieuf
	'\u11b6': 0x0ee2, //   f   # Hangul_J_RieulHieuh
	'\u11b7': 0x0ee3, //   f   # Hangul_J_Mieum
	'\u11b8': 0x0ee4, //   f   # Hangul_J_Pieub
	'\u11b9': 0x0ee5, //   f   # Hangul_J_PieubSios
	'\u11ba': 0x0ee6, //   f   # Hangul_J_Sios
	'\u11bb': 0x0ee7, //   f   # Hangul_J_SsangSios
	'\u11bc': 0x0ee8, //   f   # Hangul_J_Ieung
	'\u11bd': 0x0ee9, //   f   # Hangul_J_Jieuj
	'\u11be': 0x0eea, //   f   # Hangul_J_Cieuc
	'\u11bf': 0x0eeb, //   f   # Hangul_J_Khieuq
	'\u11c0': 0x0eec, //   f   # Hangul_J_Tieut
	'\u11c1': 0x0eed, //   f   # Hangul_J_Phieuf
	'\u11c2': 0x0eee, //   f   # Hangul_J_Hieuh
	'\u316d': 0x0eef, //   f   # Hangul_RieulYeorinHieuh
	'\u3171': 0x0ef0, //   f   # Hangul_SunkyeongeumMieum
	'\u3178': 0x0ef1, //   f   # Hangul_SunkyeongeumPieub
	'\u317f': 0x0ef2, //   f   # Hangul_PanSios
	'\u3181': 0x0ef3, //   f   # Hangul_KkogjiDalrinIeung
	'\u3184': 0x0ef4, //   f   # Hangul_SunkyeongeumPhieuf
	'\u3186': 0x0ef5, //   f   # Hangul_YeorinHieuh
	'\u318d': 0x0ef6, //   f   # Hangul_AraeA
	'\u318e': 0x0ef7, //   f   # Hangul_AraeAE
	'\u11eb': 0x0ef8, //   f   # Hangul_J_PanSios
	'\u11f0': 0x0ef9, //   f   # Hangul_J_KkogjiDalrinIeung
	'\u11f9': 0x0efa, //   f   # Hangul_J_YeorinHieuh
	'\u20a9': 0x0eff, //   o   # Korean_Won
	'\u0152': 0x13bc, //   .   # OE
	'\u0153': 0x13bd, //   .   # oe
	'\u0178': 0x13be, //   .   # Ydiaeresis
	'\u20a0': 0x20a0, //   u   # EcuSign
	'\u20a1': 0x20a1, //   u   # ColonSign
	'\u20a2': 0x20a2, //   u   # CruzeiroSign
	'\u20a3': 0x20a3, //   u   # FFrancSign
	'\u20a4': 0x20a4, //   u   # LiraSign
	'\u20a5': 0x20a5, //   u   # MillSign
	'\u20a6': 0x20a6, //   u   # NairaSign
	'\u20a7': 0x20a7, //   u   # PesetaSign
	'\u20a8': 0x20a8, //   u   # RupeeSign
	//'\u20a9': 0x20a9,   //   u   # WonSign
	'\u20aa': 0x20aa, //   u   # NewSheqelSign
	'\u20ab': 0x20ab, //   u   # DongSign
	'\u20ac': 0x20ac, //   .   # EuroSign
	//'\u0000': 0xfd01,   //   f   # 3270_Duplicate
	//'\u0000': 0xfd02,   //   f   # 3270_FieldMark
	//'\u0000': 0xfd03,   //   f   # 3270_Right2
	//'\u0000': 0xfd04,   //   f   # 3270_Left2
	//'\u0000': 0xfd05,   //   f   # 3270_BackTab
	//'\u0000': 0xfd06,   //   f   # 3270_EraseEOF
	//'\u0000': 0xfd07,   //   f   # 3270_EraseInput
	//'\u0000': 0xfd08,   //   f   # 3270_Reset
	//'\u0000': 0xfd09,   //   f   # 3270_Quit
	//'\u0000': 0xfd0a,   //   f   # 3270_PA1
	//'\u0000': 0xfd0b,   //   f   # 3270_PA2
	//'\u0000': 0xfd0c,   //   f   # 3270_PA3
	//'\u0000': 0xfd0d,   //   f   # 3270_Test
	//'\u0000': 0xfd0e,   //   f   # 3270_Attn
	//'\u0000': 0xfd0f,   //   f   # 3270_CursorBlink
	//'\u0000': 0xfd10,   //   f   # 3270_AltCursor
	//'\u0000': 0xfd11,   //   f   # 3270_KeyClick
	//'\u0000': 0xfd12,   //   f   # 3270_Jump
	//'\u0000': 0xfd13,   //   f   # 3270_Ident
	//'\u0000': 0xfd14,   //   f   # 3270_Rule
	//'\u0000': 0xfd15,   //   f   # 3270_Copy
	//'\u0000': 0xfd16,   //   f   # 3270_Play
	//'\u0000': 0xfd17,   //   f   # 3270_Setup
	//'\u0000': 0xfd18,   //   f   # 3270_Record
	//'\u0000': 0xfd19,   //   f   # 3270_ChangeScreen
	//'\u0000': 0xfd1a,   //   f   # 3270_DeleteWord
	//'\u0000': 0xfd1b,   //   f   # 3270_ExSelect
	//'\u0000': 0xfd1c,   //   f   # 3270_CursorSelect
	//'\u0000': 0xfd1d,   //   f   # 3270_PrintScreen
	//'\u0000': 0xfd1e,   //   f   # 3270_Enter
	//'\u0000': 0xfe01,   //   f   # ISO_Lock
	//'\u0000': 0xfe02,   //   f   # ISO_Level2_Latch
	//'\u0000': 0xfe03,   //   f   # ISO_Level3_Shift
	//'\u0000': 0xfe04,   //   f   # ISO_Level3_Latch
	//'\u0000': 0xfe05,   //   f   # ISO_Level3_Lock
	//'\u0000': 0xfe06,   //   f   # ISO_Group_Latch
	//'\u0000': 0xfe07,   //   f   # ISO_Group_Lock
	//'\u0000': 0xfe08,   //   f   # ISO_Next_Group
	//'\u0000': 0xfe09,   //   f   # ISO_Next_Group_Lock
	//'\u0000': 0xfe0a,   //   f   # ISO_Prev_Group
	//'\u0000': 0xfe0b,   //   f   # ISO_Prev_Group_Lock
	//'\u0000': 0xfe0c,   //   f   # ISO_First_Group
	//'\u0000': 0xfe0d,   //   f   # ISO_First_Group_Lock
	//'\u0000': 0xfe0e,   //   f   # ISO_Last_Group
	//'\u0000': 0xfe0f,   //   f   # ISO_Last_Group_Lock
	//'\u0000': 0xfe20,   //   f   # ISO_Left_Tab
	//'\u0000': 0xfe21,   //   f   # ISO_Move_Line_Up
	//'\u0000': 0xfe22,   //   f   # ISO_Move_Line_Down
	//'\u0000': 0xfe23,   //   f   # ISO_Partial_Line_Up
	//'\u0000': 0xfe24,   //   f   # ISO_Partial_Line_Down
	//'\u0000': 0xfe25,   //   f   # ISO_Partial_Space_Left
	//'\u0000': 0xfe26,   //   f   # ISO_Partial_Space_Right
	//'\u0000': 0xfe27,   //   f   # ISO_Set_Margin_Left
	//'\u0000': 0xfe28,   //   f   # ISO_Set_Margin_Right
	//'\u0000': 0xfe29,   //   f   # ISO_Release_Margin_Left
	//'\u0000': 0xfe2a,   //   f   # ISO_Release_Margin_Right
	//'\u0000': 0xfe2b,   //   f   # ISO_Release_Both_Margins
	//'\u0000': 0xfe2c,   //   f   # ISO_Fast_Cursor_Left
	//'\u0000': 0xfe2d,   //   f   # ISO_Fast_Cursor_Right
	//'\u0000': 0xfe2e,   //   f   # ISO_Fast_Cursor_Up
	//'\u0000': 0xfe2f,   //   f   # ISO_Fast_Cursor_Down
	//'\u0000': 0xfe30,   //   f   # ISO_Continuous_Underline
	//'\u0000': 0xfe31,   //   f   # ISO_Discontinuous_Underline
	//'\u0000': 0xfe32,   //   f   # ISO_Emphasize
	//'\u0000': 0xfe33,   //   f   # ISO_Center_Object
	//'\u0000': 0xfe34,   //   f   # ISO_Enter
	'\u0300': 0xfe50, //   f   # dead_grave
	'\u0301': 0xfe51, //   f   # dead_acute
	'\u0302': 0xfe52, //   f   # dead_circumflex
	'\u0303': 0xfe53, //   f   # dead_tilde
	'\u0304': 0xfe54, //   f   # dead_macron
	'\u0306': 0xfe55, //   f   # dead_breve
	'\u0307': 0xfe56, //   f   # dead_abovedot
	'\u0308': 0xfe57, //   f   # dead_diaeresis
	'\u030a': 0xfe58, //   f   # dead_abovering
	'\u030b': 0xfe59, //   f   # dead_doubleacute
	'\u030c': 0xfe5a, //   f   # dead_caron
	'\u0327': 0xfe5b, //   f   # dead_cedilla
	'\u0328': 0xfe5c, //   f   # dead_ogonek
	'\u0345': 0xfe5d, //   f   # dead_iota
	'\u3099': 0xfe5e, //   f   # dead_voiced_sound
	'\u309a': 0xfe5f, //   f   # dead_semivoiced_sound
	//'\u0000': 0xfe70,   //   f   # AccessX_Enable
	//'\u0000': 0xfe71,   //   f   # AccessX_Feedback_Enable
	//'\u0000': 0xfe72,   //   f   # RepeatKeys_Enable
	//'\u0000': 0xfe73,   //   f   # SlowKeys_Enable
	//'\u0000': 0xfe74,   //   f   # BounceKeys_Enable
	//'\u0000': 0xfe75,   //   f   # StickyKeys_Enable
	//'\u0000': 0xfe76,   //   f   # MouseKeys_Enable
	//'\u0000': 0xfe77,   //   f   # MouseKeys_Accel_Enable
	//'\u0000': 0xfe78,   //   f   # Overlay1_Enable
	//'\u0000': 0xfe79,   //   f   # Overlay2_Enable
	//'\u0000': 0xfe7a,   //   f   # AudibleBell_Enable
	//'\u0000': 0xfed0,   //   f   # First_Virtual_Screen
	//'\u0000': 0xfed1,   //   f   # Prev_Virtual_Screen
	//'\u0000': 0xfed2,   //   f   # Next_Virtual_Screen
	//'\u0000': 0xfed4,   //   f   # Last_Virtual_Screen
	//'\u0000': 0xfed5,   //   f   # Terminate_Server
	//'\u0000': 0xfee0,   //   f   # Pointer_Left
	//'\u0000': 0xfee1,   //   f   # Pointer_Right
	//'\u0000': 0xfee2,   //   f   # Pointer_Up
	//'\u0000': 0xfee3,   //   f   # Pointer_Down
	//'\u0000': 0xfee4,   //   f   # Pointer_UpLeft
	//'\u0000': 0xfee5,   //   f   # Pointer_UpRight
	//'\u0000': 0xfee6,   //   f   # Pointer_DownLeft
	//'\u0000': 0xfee7,   //   f   # Pointer_DownRight
	//'\u0000': 0xfee8,   //   f   # Pointer_Button_Dflt
	//'\u0000': 0xfee9,   //   f   # Pointer_Button1
	//'\u0000': 0xfeea,   //   f   # Pointer_Button2
	//'\u0000': 0xfeeb,   //   f   # Pointer_Button3
	//'\u0000': 0xfeec,   //   f   # Pointer_Button4
	//'\u0000': 0xfeed,   //   f   # Pointer_Button5
	//'\u0000': 0xfeee,   //   f   # Pointer_DblClick_Dflt
	//'\u0000': 0xfeef,   //   f   # Pointer_DblClick1
	//'\u0000': 0xfef0,   //   f   # Pointer_DblClick2
	//'\u0000': 0xfef1,   //   f   # Pointer_DblClick3
	//'\u0000': 0xfef2,   //   f   # Pointer_DblClick4
	//'\u0000': 0xfef3,   //   f   # Pointer_DblClick5
	//'\u0000': 0xfef4,   //   f   # Pointer_Drag_Dflt
	//'\u0000': 0xfef5,   //   f   # Pointer_Drag1
	//'\u0000': 0xfef6,   //   f   # Pointer_Drag2
	//'\u0000': 0xfef7,   //   f   # Pointer_Drag3
	//'\u0000': 0xfef8,   //   f   # Pointer_Drag4
	//'\u0000': 0xfef9,   //   f   # Pointer_EnableKeys
	//'\u0000': 0xfefa,   //   f   # Pointer_Accelerate
	//'\u0000': 0xfefb,   //   f   # Pointer_DfltBtnNext
	//'\u0000': 0xfefc,   //   f   # Pointer_DfltBtnPrev
	//'\u0000': 0xfefd,   //   f   # Pointer_Drag5
	'\u0008': 0xff08, //   f   # BackSpace	/* back space, back char */
	'\u0009': 0xff09, //   f   # Tab
	'\u000a': 0xff0a, //   f   # Linefeed	/* Linefeed, LF */
	'\u000b': 0xff0b, //   f   # Clear
	'\u000d': 0xff0d, //   f   # Return	/* Return, enter */
	'\u0013': 0xff13, //   f   # Pause	/* Pause, hold */
	'\u0014': 0xff14, //   f   # Scroll_Lock
	'\u0015': 0xff15, //   f   # Sys_Req
	'\u001b': 0xff1b, //   f   # Escape
	//'\u0000': 0xff20,   //   f   # Multi_key
	//'\u0000': 0xff21,   //   f   # Kanji
	//'\u0000': 0xff22,   //   f   # Muhenkan
	//'\u0000': 0xff23,   //   f   # Henkan_Mode
	//'\u0000': 0xff24,   //   f   # Romaji
	//'\u0000': 0xff25,   //   f   # Hiragana
	//'\u0000': 0xff26,   //   f   # Katakana
	//'\u0000': 0xff27,   //   f   # Hiragana_Katakana
	//'\u0000': 0xff28,   //   f   # Zenkaku
	//'\u0000': 0xff29,   //   f   # Hankaku
	//'\u0000': 0xff2a,   //   f   # Zenkaku_Hankaku
	//'\u0000': 0xff2b,   //   f   # Touroku
	//'\u0000': 0xff2c,   //   f   # Massyo
	//'\u0000': 0xff2d,   //   f   # Kana_Lock
	//'\u0000': 0xff2e,   //   f   # Kana_Shift
	//'\u0000': 0xff2f,   //   f   # Eisu_Shift
	//'\u0000': 0xff30,   //   f   # Eisu_toggle
	//'\u0000': 0xff31,   //   f   # Hangul
	//'\u0000': 0xff32,   //   f   # Hangul_Start
	//'\u0000': 0xff33,   //   f   # Hangul_End
	//'\u0000': 0xff34,   //   f   # Hangul_Hanja
	//'\u0000': 0xff35,   //   f   # Hangul_Jamo
	//'\u0000': 0xff36,   //   f   # Hangul_Romaja
	//'\u0000': 0xff37,   //   f   # Codeinput
	//'\u0000': 0xff38,   //   f   # Hangul_Jeonja
	//'\u0000': 0xff39,   //   f   # Hangul_Banja
	//'\u0000': 0xff3a,   //   f   # Hangul_PreHanja
	//'\u0000': 0xff3b,   //   f   # Hangul_PostHanja
	//'\u0000': 0xff3c,   //   f   # SingleCandidate
	//'\u0000': 0xff3d,   //   f   # MultipleCandidate
	//'\u0000': 0xff3e,   //   f   # PreviousCandidate
	//'\u0000': 0xff3f,   //   f   # Hangul_Special
	//'\u0000': 0xff50,   //   f   # Home
	//'\u0000': 0xff51,   //   f   # Left
	//'\u0000': 0xff52,   //   f   # Up
	//'\u0000': 0xff53,   //   f   # Right
	//'\u0000': 0xff54,   //   f   # Down
	//'\u0000': 0xff55,   //   f   # Prior
	//'\u0000': 0xff56,   //   f   # Next
	//'\u0000': 0xff57,   //   f   # End
	//'\u0000': 0xff58,   //   f   # Begin
	//'\u0000': 0xff60,   //   f   # Select
	//'\u0000': 0xff61,   //   f   # Print
	//'\u0000': 0xff62,   //   f   # Execute
	//'\u0000': 0xff63,   //   f   # Insert
	//'\u0000': 0xff65,   //   f   # Undo
	//'\u0000': 0xff66,   //   f   # Redo
	//'\u0000': 0xff67,   //   f   # Menu
	//'\u0000': 0xff68,   //   f   # Find
	//'\u0000': 0xff69,   //   f   # Cancel
	//'\u0000': 0xff6a,   //   f   # Help
	//'\u0000': 0xff6b,   //   f   # Break
	//'\u0000': 0xff7e,   //   f   # Mode_switch
	//'\u0000': 0xff7f,   //   f   # Num_Lock
	//'\u0020': 0xff80,   //   f   # KP_Space	/* space */
	//'\u0009': 0xff89,   //   f   # KP_Tab
	//'\u000d': 0xff8d,   //   f   # KP_Enter	/* enter */
	//'\u0000': 0xff91,   //   f   # KP_F1
	//'\u0000': 0xff92,   //   f   # KP_F2
	//'\u0000': 0xff93,   //   f   # KP_F3
	//'\u0000': 0xff94,   //   f   # KP_F4
	//'\u0000': 0xff95,   //   f   # KP_Home
	//'\u0000': 0xff96,   //   f   # KP_Left
	//'\u0000': 0xff97,   //   f   # KP_Up
	//'\u0000': 0xff98,   //   f   # KP_Right
	//'\u0000': 0xff99,   //   f   # KP_Down
	//'\u0000': 0xff9a,   //   f   # KP_Prior
	//'\u0000': 0xff9b,   //   f   # KP_Next
	//'\u0000': 0xff9c,   //   f   # KP_End
	//'\u0000': 0xff9d,   //   f   # KP_Begin
	//'\u0000': 0xff9e,   //   f   # KP_Insert
	//'\u0000': 0xff9f,   //   f   # KP_Delete
	//'\u002a': 0xffaa,   //   f   # KP_Multiply
	//'\u002b': 0xffab,   //   f   # KP_Add
	//'\u002c': 0xffac,   //   f   # KP_Separator	/* separator, often comma */
	//'\u002d': 0xffad,   //   f   # KP_Subtract
	//'\u002e': 0xffae,   //   f   # KP_Decimal
	//'\u002f': 0xffaf,   //   f   # KP_Divide
	//'\u0030': 0xffb0,   //   f   # KP_0
	//'\u0031': 0xffb1,   //   f   # KP_1
	//'\u0032': 0xffb2,   //   f   # KP_2
	//'\u0033': 0xffb3,   //   f   # KP_3
	//'\u0034': 0xffb4,   //   f   # KP_4
	//'\u0035': 0xffb5,   //   f   # KP_5
	//'\u0036': 0xffb6,   //   f   # KP_6
	//'\u0037': 0xffb7,   //   f   # KP_7
	//'\u0038': 0xffb8,   //   f   # KP_8
	//'\u0039': 0xffb9,   //   f   # KP_9
	//'\u003d': 0xffbd,   //   f   # KP_Equal	/* equals */
	//'\u0000': 0xffbe,   //   f   # F1
	//'\u0000': 0xffbf,   //   f   # F2
	//'\u0000': 0xffc0,   //   f   # F3
	//'\u0000': 0xffc1,   //   f   # F4
	//'\u0000': 0xffc2,   //   f   # F5
	//'\u0000': 0xffc3,   //   f   # F6
	//'\u0000': 0xffc4,   //   f   # F7
	//'\u0000': 0xffc5,   //   f   # F8
	//'\u0000': 0xffc6,   //   f   # F9
	//'\u0000': 0xffc7,   //   f   # F10
	//'\u0000': 0xffc8,   //   f   # F11
	//'\u0000': 0xffc9,   //   f   # F12
	//'\u0000': 0xffca,   //   f   # F13
	//'\u0000': 0xffcb,   //   f   # F14
	//'\u0000': 0xffcc,   //   f   # F15
	//'\u0000': 0xffcd,   //   f   # F16
	//'\u0000': 0xffce,   //   f   # F17
	//'\u0000': 0xffcf,   //   f   # F18
	//'\u0000': 0xffd0,   //   f   # F19
	//'\u0000': 0xffd1,   //   f   # F20
	//'\u0000': 0xffd2,   //   f   # F21
	//'\u0000': 0xffd3,   //   f   # F22
	//'\u0000': 0xffd4,   //   f   # F23
	//'\u0000': 0xffd5,   //   f   # F24
	//'\u0000': 0xffd6,   //   f   # F25
	//'\u0000': 0xffd7,   //   f   # F26
	//'\u0000': 0xffd8,   //   f   # F27
	//'\u0000': 0xffd9,   //   f   # F28
	//'\u0000': 0xffda,   //   f   # F29
	//'\u0000': 0xffdb,   //   f   # F30
	//'\u0000': 0xffdc,   //   f   # F31
	//'\u0000': 0xffdd,   //   f   # F32
	//'\u0000': 0xffde,   //   f   # F33
	//'\u0000': 0xffdf,   //   f   # F34
	//'\u0000': 0xffe0,   //   f   # F35
	//'\u0000': 0xffe1,   //   f   # Shift_L
	//'\u0000': 0xffe2,   //   f   # Shift_R
	//'\u0000': 0xffe3,   //   f   # Control_L
	//'\u0000': 0xffe4,   //   f   # Control_R
	//'\u0000': 0xffe5,   //   f   # Caps_Lock
	//'\u0000': 0xffe6,   //   f   # Shift_Lock
	//'\u0000': 0xffe7,   //   f   # Meta_L
	//'\u0000': 0xffe8,   //   f   # Meta_R
	//'\u0000': 0xffe9,   //   f   # Alt_L
	//'\u0000': 0xffea,   //   f   # Alt_R
	//'\u0000': 0xffeb,   //   f   # Super_L
	//'\u0000': 0xffec,   //   f   # Super_R
	//'\u0000': 0xffed,   //   f   # Hyper_L
	//'\u0000': 0xffee,   //   f   # Hyper_R
	//'\u0000': 0xffff,   //   f   # Delete
	//'\u0000': 0xffffff, //   f   # VoidSymbol

	// # Various XFree86 extensions since X11R6.4
	// # http://cvsweb.xfree86.org/cvsweb/xc/include/keysymdef.h

	// # KOI8-U support (Aleksey Novodvorsky, 1999-05-30)
	// # http://cvsweb.xfree86.org/cvsweb/xc/include/keysymdef.h.diff?r1=1.4&r2=1.5
	// # Used in XFree86's /usr/lib/X11/xkb/symbols/ua mappings

	'\u0491': 0x06ad, //   .   # Ukrainian_ghe_with_upturn
	'\u0490': 0x06bd, //   .   # Ukrainian_GHE_WITH_UPTURN

	// # Support for armscii-8, ibm-cp1133, mulelao-1, viscii1.1-1,
	// # tcvn-5712, georgian-academy, georgian-ps
	// # (#2843, Pablo Saratxaga <pablo@mandrakesoft.com>, 1999-06-06)
	// # http://cvsweb.xfree86.org/cvsweb/xc/include/keysymdef.h.diff?r1=1.6&r2=1.7

	// # Armenian
	// # (not used in any XFree86 4.4 kbd layouts, where /usr/lib/X11/xkb/symbols/am
	// # uses directly Unicode-mapped hexadecimal values instead)
	//'\u0000': 0x14a1, //   r   # Armenian_eternity
	//'\u0587': 0x14a2, //   u   # Armenian_ligature_ew
	//'\u0589': 0x14a3, //   u   # Armenian_verjaket
	//'\u0029': 0x14a4, //   r   # Armenian_parenright
	//'\u0028': 0x14a5, //   r   # Armenian_parenleft
	//'\u00bb': 0x14a6, //   r   # Armenian_guillemotright
	//'\u00ab': 0x14a7, //   r   # Armenian_guillemotleft
	//'\u2014': 0x14a8, //   r   # Armenian_em_dash
	//'\u002e': 0x14a9, //   r   # Armenian_mijaket
	//'\u055d': 0x14aa, //   u   # Armenian_but
	//'\u002c': 0x14ab, //   r   # Armenian_comma
	//'\u2013': 0x14ac, //   r   # Armenian_en_dash
	//'\u058a': 0x14ad, //   u   # Armenian_yentamna
	//'\u2026': 0x14ae, //   r   # Armenian_ellipsis
	//'\u055c': 0x14af, //   u   # Armenian_amanak
	//'\u055b': 0x14b0, //   u   # Armenian_shesht
	//'\u055e': 0x14b1, //   u   # Armenian_paruyk
	//'\u0531': 0x14b2, //   u   # Armenian_AYB
	//'\u0561': 0x14b3, //   u   # Armenian_ayb
	//'\u0532': 0x14b4, //   u   # Armenian_BEN
	//'\u0562': 0x14b5, //   u   # Armenian_ben
	//'\u0533': 0x14b6, //   u   # Armenian_GIM
	//'\u0563': 0x14b7, //   u   # Armenian_gim
	//'\u0534': 0x14b8, //   u   # Armenian_DA
	//'\u0564': 0x14b9, //   u   # Armenian_da
	//'\u0535': 0x14ba, //   u   # Armenian_YECH
	//'\u0565': 0x14bb, //   u   # Armenian_yech
	//'\u0536': 0x14bc, //   u   # Armenian_ZA
	//'\u0566': 0x14bd, //   u   # Armenian_za
	//'\u0537': 0x14be, //   u   # Armenian_E
	//'\u0567': 0x14bf, //   u   # Armenian_e
	//'\u0538': 0x14c0, //   u   # Armenian_AT
	//'\u0568': 0x14c1, //   u   # Armenian_at
	//'\u0539': 0x14c2, //   u   # Armenian_TO
	//'\u0569': 0x14c3, //   u   # Armenian_to
	//'\u053a': 0x14c4, //   u   # Armenian_ZHE
	//'\u056a': 0x14c5, //   u   # Armenian_zhe
	//'\u053b': 0x14c6, //   u   # Armenian_INI
	//'\u056b': 0x14c7, //   u   # Armenian_ini
	//'\u053c': 0x14c8, //   u   # Armenian_LYUN
	//'\u056c': 0x14c9, //   u   # Armenian_lyun
	//'\u053d': 0x14ca, //   u   # Armenian_KHE
	//'\u056d': 0x14cb, //   u   # Armenian_khe
	//'\u053e': 0x14cc, //   u   # Armenian_TSA
	//'\u056e': 0x14cd, //   u   # Armenian_tsa
	//'\u053f': 0x14ce, //   u   # Armenian_KEN
	//'\u056f': 0x14cf, //   u   # Armenian_ken
	//'\u0540': 0x14d0, //   u   # Armenian_HO
	//'\u0570': 0x14d1, //   u   # Armenian_ho
	//'\u0541': 0x14d2, //   u   # Armenian_DZA
	//'\u0571': 0x14d3, //   u   # Armenian_dza
	//'\u0542': 0x14d4, //   u   # Armenian_GHAT
	//'\u0572': 0x14d5, //   u   # Armenian_ghat
	//'\u0543': 0x14d6, //   u   # Armenian_TCHE
	//'\u0573': 0x14d7, //   u   # Armenian_tche
	//'\u0544': 0x14d8, //   u   # Armenian_MEN
	//'\u0574': 0x14d9, //   u   # Armenian_men
	//'\u0545': 0x14da, //   u   # Armenian_HI
	//'\u0575': 0x14db, //   u   # Armenian_hi
	//'\u0546': 0x14dc, //   u   # Armenian_NU
	//'\u0576': 0x14dd, //   u   # Armenian_nu
	//'\u0547': 0x14de, //   u   # Armenian_SHA
	//'\u0577': 0x14df, //   u   # Armenian_sha
	//'\u0548': 0x14e0, //   u   # Armenian_VO
	//'\u0578': 0x14e1, //   u   # Armenian_vo
	//'\u0549': 0x14e2, //   u   # Armenian_CHA
	//'\u0579': 0x14e3, //   u   # Armenian_cha
	//'\u054a': 0x14e4, //   u   # Armenian_PE
	//'\u057a': 0x14e5, //   u   # Armenian_pe
	//'\u054b': 0x14e6, //   u   # Armenian_JE
	//'\u057b': 0x14e7, //   u   # Armenian_je
	//'\u054c': 0x14e8, //   u   # Armenian_RA
	//'\u057c': 0x14e9, //   u   # Armenian_ra
	//'\u054d': 0x14ea, //   u   # Armenian_SE
	//'\u057d': 0x14eb, //   u   # Armenian_se
	//'\u054e': 0x14ec, //   u   # Armenian_VEV
	//'\u057e': 0x14ed, //   u   # Armenian_vev
	//'\u054f': 0x14ee, //   u   # Armenian_TYUN
	//'\u057f': 0x14ef, //   u   # Armenian_tyun
	//'\u0550': 0x14f0, //   u   # Armenian_RE
	//'\u0580': 0x14f1, //   u   # Armenian_re
	//'\u0551': 0x14f2, //   u   # Armenian_TSO
	//'\u0581': 0x14f3, //   u   # Armenian_tso
	//'\u0552': 0x14f4, //   u   # Armenian_VYUN
	//'\u0582': 0x14f5, //   u   # Armenian_vyun
	//'\u0553': 0x14f6, //   u   # Armenian_PYUR
	//'\u0583': 0x14f7, //   u   # Armenian_pyur
	//'\u0554': 0x14f8, //   u   # Armenian_KE
	//'\u0584': 0x14f9, //   u   # Armenian_ke
	//'\u0555': 0x14fa, //   u   # Armenian_O
	//'\u0585': 0x14fb, //   u   # Armenian_o
	//'\u0556': 0x14fc, //   u   # Armenian_FE
	//'\u0586': 0x14fd, //   u   # Armenian_fe
	//'\u055a': 0x14fe, //   u   # Armenian_apostrophe
	//'\u00a7': 0x14ff, //   r   # Armenian_section_sign

	// # Gregorian
	// # (not used in any XFree86 4.4 kbd layouts, were /usr/lib/X11/xkb/symbols/ge_*
	// # uses directly Unicode-mapped hexadecimal values instead)
	'\u10d0': 0x15d0, //   u   # Georgian_an
	'\u10d1': 0x15d1, //   u   # Georgian_ban
	'\u10d2': 0x15d2, //   u   # Georgian_gan
	'\u10d3': 0x15d3, //   u   # Georgian_don
	'\u10d4': 0x15d4, //   u   # Georgian_en
	'\u10d5': 0x15d5, //   u   # Georgian_vin
	'\u10d6': 0x15d6, //   u   # Georgian_zen
	'\u10d7': 0x15d7, //   u   # Georgian_tan
	'\u10d8': 0x15d8, //   u   # Georgian_in
	'\u10d9': 0x15d9, //   u   # Georgian_kan
	'\u10da': 0x15da, //   u   # Georgian_las
	'\u10db': 0x15db, //   u   # Georgian_man
	'\u10dc': 0x15dc, //   u   # Georgian_nar
	'\u10dd': 0x15dd, //   u   # Georgian_on
	'\u10de': 0x15de, //   u   # Georgian_par
	'\u10df': 0x15df, //   u   # Georgian_zhar
	'\u10e0': 0x15e0, //   u   # Georgian_rae
	'\u10e1': 0x15e1, //   u   # Georgian_san
	'\u10e2': 0x15e2, //   u   # Georgian_tar
	'\u10e3': 0x15e3, //   u   # Georgian_un
	'\u10e4': 0x15e4, //   u   # Georgian_phar
	'\u10e5': 0x15e5, //   u   # Georgian_khar
	'\u10e6': 0x15e6, //   u   # Georgian_ghan
	'\u10e7': 0x15e7, //   u   # Georgian_qar
	'\u10e8': 0x15e8, //   u   # Georgian_shin
	'\u10e9': 0x15e9, //   u   # Georgian_chin
	'\u10ea': 0x15ea, //   u   # Georgian_can
	'\u10eb': 0x15eb, //   u   # Georgian_jil
	'\u10ec': 0x15ec, //   u   # Georgian_cil
	'\u10ed': 0x15ed, //   u   # Georgian_char
	'\u10ee': 0x15ee, //   u   # Georgian_xan
	'\u10ef': 0x15ef, //   u   # Georgian_jhan
	'\u10f0': 0x15f0, //   u   # Georgian_hae
	'\u10f1': 0x15f1, //   u   # Georgian_he
	'\u10f2': 0x15f2, //   u   # Georgian_hie
	'\u10f3': 0x15f3, //   u   # Georgian_we
	'\u10f4': 0x15f4, //   u   # Georgian_har
	'\u10f5': 0x15f5, //   u   # Georgian_hoe
	'\u10f6': 0x15f6, //   u   # Georgian_fi

	// # Pablo Saratxaga's i18n updates for XFree86 that are used in Mandrake 7.2.
	// # (#4195, Pablo Saratxaga <pablo@mandrakesoft.com>, 2000-10-27)
	// # http://cvsweb.xfree86.org/cvsweb/xc/include/keysymdef.h.diff?r1=1.9&r2=1.10

	// # Latin-8
	// # (the *abovedot keysyms are used in /usr/lib/X11/xkb/symbols/ie)
	'\u1e02': 0x12a1, //   u   # Babovedot
	'\u1e03': 0x12a2, //   u   # babovedot
	'\u1e0a': 0x12a6, //   u   # Dabovedot
	'\u1e80': 0x12a8, //   u   # Wgrave
	'\u1e82': 0x12aa, //   u   # Wacute
	'\u1e0b': 0x12ab, //   u   # dabovedot
	'\u1ef2': 0x12ac, //   u   # Ygrave
	'\u1e1e': 0x12b0, //   u   # Fabovedot
	'\u1e1f': 0x12b1, //   u   # fabovedot
	'\u1e40': 0x12b4, //   u   # Mabovedot
	'\u1e41': 0x12b5, //   u   # mabovedot
	'\u1e56': 0x12b7, //   u   # Pabovedot
	'\u1e81': 0x12b8, //   u   # wgrave
	'\u1e57': 0x12b9, //   u   # pabovedot
	'\u1e83': 0x12ba, //   u   # wacute
	'\u1e60': 0x12bb, //   u   # Sabovedot
	'\u1ef3': 0x12bc, //   u   # ygrave
	'\u1e84': 0x12bd, //   u   # Wdiaeresis
	'\u1e85': 0x12be, //   u   # wdiaeresis
	'\u1e61': 0x12bf, //   u   # sabovedot
	'\u0174': 0x12d0, //   u   # Wcircumflex
	'\u1e6a': 0x12d7, //   u   # Tabovedot
	'\u0176': 0x12de, //   u   # Ycircumflex
	'\u0175': 0x12f0, //   u   # wcircumflex
	'\u1e6b': 0x12f7, //   u   # tabovedot
	'\u0177': 0x12fe, //   u   # ycircumflex

	// # Arabic
	// # (of these, in XFree86 4.4 only Arabic_superscript_alef, Arabic_madda_above,
	// # Arabic_hamza_* are actually used, e.g. in /usr/lib/X11/xkb/symbols/syr)
	'\u06f0': 0x0590, //   u   # Farsi_0
	'\u06f1': 0x0591, //   u   # Farsi_1
	'\u06f2': 0x0592, //   u   # Farsi_2
	'\u06f3': 0x0593, //   u   # Farsi_3
	'\u06f4': 0x0594, //   u   # Farsi_4
	'\u06f5': 0x0595, //   u   # Farsi_5
	'\u06f6': 0x0596, //   u   # Farsi_6
	'\u06f7': 0x0597, //   u   # Farsi_7
	'\u06f8': 0x0598, //   u   # Farsi_8
	'\u06f9': 0x0599, //   u   # Farsi_9
	'\u066a': 0x05a5, //   u   # Arabic_percent
	'\u0670': 0x05a6, //   u   # Arabic_superscript_alef
	'\u0679': 0x05a7, //   u   # Arabic_tteh
	'\u067e': 0x05a8, //   u   # Arabic_peh
	'\u0686': 0x05a9, //   u   # Arabic_tcheh
	'\u0688': 0x05aa, //   u   # Arabic_ddal
	'\u0691': 0x05ab, //   u   # Arabic_rreh
	'\u06d4': 0x05ae, //   u   # Arabic_fullstop
	'\u0660': 0x05b0, //   u   # Arabic_0
	'\u0661': 0x05b1, //   u   # Arabic_1
	'\u0662': 0x05b2, //   u   # Arabic_2
	'\u0663': 0x05b3, //   u   # Arabic_3
	'\u0664': 0x05b4, //   u   # Arabic_4
	'\u0665': 0x05b5, //   u   # Arabic_5
	'\u0666': 0x05b6, //   u   # Arabic_6
	'\u0667': 0x05b7, //   u   # Arabic_7
	'\u0668': 0x05b8, //   u   # Arabic_8
	'\u0669': 0x05b9, //   u   # Arabic_9
	'\u0653': 0x05f3, //   u   # Arabic_madda_above
	'\u0654': 0x05f4, //   u   # Arabic_hamza_above
	'\u0655': 0x05f5, //   u   # Arabic_hamza_below
	'\u0698': 0x05f6, //   u   # Arabic_jeh
	'\u06a4': 0x05f7, //   u   # Arabic_veh
	'\u06a9': 0x05f8, //   u   # Arabic_keheh
	'\u06af': 0x05f9, //   u   # Arabic_gaf
	'\u06ba': 0x05fa, //   u   # Arabic_noon_ghunna
	'\u06be': 0x05fb, //   u   # Arabic_heh_doachashmee
	'\u06cc': 0x05fc, //   u   # Farsi_yeh
	'\u06d2': 0x05fd, //   u   # Arabic_yeh_baree
	'\u06c1': 0x05fe, //   u   # Arabic_heh_goal

	// # Cyrillic
	// # (none of these are actually used in any XFree86 4.4 kbd layouts)
	'\u0492': 0x0680, //   u   # Cyrillic_GHE_bar
	'\u0496': 0x0681, //   u   # Cyrillic_ZHE_descender
	'\u049a': 0x0682, //   u   # Cyrillic_KA_descender
	'\u049c': 0x0683, //   u   # Cyrillic_KA_vertstroke
	'\u04a2': 0x0684, //   u   # Cyrillic_EN_descender
	'\u04ae': 0x0685, //   u   # Cyrillic_U_straight
	'\u04b0': 0x0686, //   u   # Cyrillic_U_straight_bar
	'\u04b2': 0x0687, //   u   # Cyrillic_HA_descender
	'\u04b6': 0x0688, //   u   # Cyrillic_CHE_descender
	'\u04b8': 0x0689, //   u   # Cyrillic_CHE_vertstroke
	'\u04ba': 0x068a, //   u   # Cyrillic_SHHA
	'\u04d8': 0x068c, //   u   # Cyrillic_SCHWA
	'\u04e2': 0x068d, //   u   # Cyrillic_I_macron
	'\u04e8': 0x068e, //   u   # Cyrillic_O_bar
	'\u04ee': 0x068f, //   u   # Cyrillic_U_macron
	'\u0493': 0x0690, //   u   # Cyrillic_ghe_bar
	'\u0497': 0x0691, //   u   # Cyrillic_zhe_descender
	'\u049b': 0x0692, //   u   # Cyrillic_ka_descender
	'\u049d': 0x0693, //   u   # Cyrillic_ka_vertstroke
	'\u04a3': 0x0694, //   u   # Cyrillic_en_descender
	'\u04af': 0x0695, //   u   # Cyrillic_u_straight
	'\u04b1': 0x0696, //   u   # Cyrillic_u_straight_bar
	'\u04b3': 0x0697, //   u   # Cyrillic_ha_descender
	'\u04b7': 0x0698, //   u   # Cyrillic_che_descender
	'\u04b9': 0x0699, //   u   # Cyrillic_che_vertstroke
	'\u04bb': 0x069a, //   u   # Cyrillic_shha
	'\u04d9': 0x069c, //   u   # Cyrillic_schwa
	'\u04e3': 0x069d, //   u   # Cyrillic_i_macron
	'\u04e9': 0x069e, //   u   # Cyrillic_o_bar
	'\u04ef': 0x069f, //   u   # Cyrillic_u_macron

	// # Caucasus
	// # (of these, in XFree86 4.4 only Gcaron, gcaron are actually used,
	// # e.g. in /usr/lib/X11/xkb/symbols/sapmi; the lack of Unicode
	// # equivalents for the others suggests that they are bogus)
	//'\u0000': 0x16a2, //   r   # Ccedillaabovedot
	'\u1e8a': 0x16a3, //   u   # Xabovedot
	//'\u0000': 0x16a5, //   r   # Qabovedot
	'\u012c': 0x16a6, //   u   # Ibreve
	//'\u0000': 0x16a7, //   r   # IE
	//'\u0000': 0x16a8, //   r   # UO
	'\u01b5': 0x16a9, //   u   # Zstroke
	'\u01e6': 0x16aa, //   u   # Gcaron
	'\u019f': 0x16af, //   u   # Obarred
	//'\u0000': 0x16b2, //   r   # ccedillaabovedot
	'\u1e8b': 0x16b3, //   u   # xabovedot
	//'\u0000': 0x16b4, //   r   # Ocaron
	//'\u0000': 0x16b5, //   r   # qabovedot
	'\u012d': 0x16b6, //   u   # ibreve
	//'\u0000': 0x16b7, //   r   # ie
	//'\u0000': 0x16b8, //   r   # uo
	'\u01b6': 0x16b9, //   u   # zstroke
	'\u01e7': 0x16ba, //   u   # gcaron
	'\u01d2': 0x16bd, //   u   # ocaron
	'\u0275': 0x16bf, //   u   # obarred
	'\u018f': 0x16c6, //   u   # SCHWA
	'\u0259': 0x16f6, //   u   # schwa

	// # Inupiak, Guarani
	// # (none of these are actually used in any XFree86 4.4 kbd layouts,
	// # and the lack of Unicode equivalents suggests that they are bogus)
	'\u1e36': 0x16d1, //   u   # Lbelowdot
	//'\u0000': 0x16d2, //   r   # Lstrokebelowdot
	//'\u0000': 0x16d3, //   r   # Gtilde
	'\u1e37': 0x16e1, //   u   # lbelowdot
	//'\u0000': 0x16e2, //   r   # lstrokebelowdot
	//'\u0000': 0x16e3, //   r   # gtilde

	// # Vietnamese
	// # (none of these are actually used in any XFree86 4.4 kbd layouts; they are
	// # also pointless, as Vietnamese input methods use dead accent keys + ASCII keys)
	'\u1ea0': 0x1ea0, //   u   # Abelowdot
	'\u1ea1': 0x1ea1, //   u   # abelowdot
	'\u1ea2': 0x1ea2, //   u   # Ahook
	'\u1ea3': 0x1ea3, //   u   # ahook
	'\u1ea4': 0x1ea4, //   u   # Acircumflexacute
	'\u1ea5': 0x1ea5, //   u   # acircumflexacute
	'\u1ea6': 0x1ea6, //   u   # Acircumflexgrave
	'\u1ea7': 0x1ea7, //   u   # acircumflexgrave
	'\u1ea8': 0x1ea8, //   u   # Acircumflexhook
	'\u1ea9': 0x1ea9, //   u   # acircumflexhook
	'\u1eaa': 0x1eaa, //   u   # Acircumflextilde
	'\u1eab': 0x1eab, //   u   # acircumflextilde
	'\u1eac': 0x1eac, //   u   # Acircumflexbelowdot
	'\u1ead': 0x1ead, //   u   # acircumflexbelowdot
	'\u1eae': 0x1eae, //   u   # Abreveacute
	'\u1eaf': 0x1eaf, //   u   # abreveacute
	'\u1eb0': 0x1eb0, //   u   # Abrevegrave
	'\u1eb1': 0x1eb1, //   u   # abrevegrave
	'\u1eb2': 0x1eb2, //   u   # Abrevehook
	'\u1eb3': 0x1eb3, //   u   # abrevehook
	'\u1eb4': 0x1eb4, //   u   # Abrevetilde
	'\u1eb5': 0x1eb5, //   u   # abrevetilde
	'\u1eb6': 0x1eb6, //   u   # Abrevebelowdot
	'\u1eb7': 0x1eb7, //   u   # abrevebelowdot
	'\u1eb8': 0x1eb8, //   u   # Ebelowdot
	'\u1eb9': 0x1eb9, //   u   # ebelowdot
	'\u1eba': 0x1eba, //   u   # Ehook
	'\u1ebb': 0x1ebb, //   u   # ehook
	'\u1ebc': 0x1ebc, //   u   # Etilde
	'\u1ebd': 0x1ebd, //   u   # etilde
	'\u1ebe': 0x1ebe, //   u   # Ecircumflexacute
	'\u1ebf': 0x1ebf, //   u   # ecircumflexacute
	'\u1ec0': 0x1ec0, //   u   # Ecircumflexgrave
	'\u1ec1': 0x1ec1, //   u   # ecircumflexgrave
	'\u1ec2': 0x1ec2, //   u   # Ecircumflexhook
	'\u1ec3': 0x1ec3, //   u   # ecircumflexhook
	'\u1ec4': 0x1ec4, //   u   # Ecircumflextilde
	'\u1ec5': 0x1ec5, //   u   # ecircumflextilde
	'\u1ec6': 0x1ec6, //   u   # Ecircumflexbelowdot
	'\u1ec7': 0x1ec7, //   u   # ecircumflexbelowdot
	'\u1ec8': 0x1ec8, //   u   # Ihook
	'\u1ec9': 0x1ec9, //   u   # ihook
	'\u1eca': 0x1eca, //   u   # Ibelowdot
	'\u1ecb': 0x1ecb, //   u   # ibelowdot
	'\u1ecc': 0x1ecc, //   u   # Obelowdot
	'\u1ecd': 0x1ecd, //   u   # obelowdot
	'\u1ece': 0x1ece, //   u   # Ohook
	'\u1ecf': 0x1ecf, //   u   # ohook
	'\u1ed0': 0x1ed0, //   u   # Ocircumflexacute
	'\u1ed1': 0x1ed1, //   u   # ocircumflexacute
	'\u1ed2': 0x1ed2, //   u   # Ocircumflexgrave
	'\u1ed3': 0x1ed3, //   u   # ocircumflexgrave
	'\u1ed4': 0x1ed4, //   u   # Ocircumflexhook
	'\u1ed5': 0x1ed5, //   u   # ocircumflexhook
	'\u1ed6': 0x1ed6, //   u   # Ocircumflextilde
	'\u1ed7': 0x1ed7, //   u   # ocircumflextilde
	'\u1ed8': 0x1ed8, //   u   # Ocircumflexbelowdot
	'\u1ed9': 0x1ed9, //   u   # ocircumflexbelowdot
	'\u1eda': 0x1eda, //   u   # Ohornacute
	'\u1edb': 0x1edb, //   u   # ohornacute
	'\u1edc': 0x1edc, //   u   # Ohorngrave
	'\u1edd': 0x1edd, //   u   # ohorngrave
	'\u1ede': 0x1ede, //   u   # Ohornhook
	'\u1edf': 0x1edf, //   u   # ohornhook
	'\u1ee0': 0x1ee0, //   u   # Ohorntilde
	'\u1ee1': 0x1ee1, //   u   # ohorntilde
	'\u1ee2': 0x1ee2, //   u   # Ohornbelowdot
	'\u1ee3': 0x1ee3, //   u   # ohornbelowdot
	'\u1ee4': 0x1ee4, //   u   # Ubelowdot
	'\u1ee5': 0x1ee5, //   u   # ubelowdot
	'\u1ee6': 0x1ee6, //   u   # Uhook
	'\u1ee7': 0x1ee7, //   u   # uhook
	'\u1ee8': 0x1ee8, //   u   # Uhornacute
	'\u1ee9': 0x1ee9, //   u   # uhornacute
	'\u1eea': 0x1eea, //   u   # Uhorngrave
	'\u1eeb': 0x1eeb, //   u   # uhorngrave
	'\u1eec': 0x1eec, //   u   # Uhornhook
	'\u1eed': 0x1eed, //   u   # uhornhook
	'\u1eee': 0x1eee, //   u   # Uhorntilde
	'\u1eef': 0x1eef, //   u   # uhorntilde
	'\u1ef0': 0x1ef0, //   u   # Uhornbelowdot
	'\u1ef1': 0x1ef1, //   u   # uhornbelowdot
	'\u1ef4': 0x1ef4, //   u   # Ybelowdot
	'\u1ef5': 0x1ef5, //   u   # ybelowdot
	'\u1ef6': 0x1ef6, //   u   # Yhook
	'\u1ef7': 0x1ef7, //   u   # yhook
	'\u1ef8': 0x1ef8, //   u   # Ytilde
	'\u1ef9': 0x1ef9, //   u   # ytilde

	'\u01a0': 0x1efa, //   u   # Ohorn
	'\u01a1': 0x1efb, //   u   # ohorn
	'\u01af': 0x1efc, //   u   # Uhorn
	'\u01b0': 0x1efd, //   u   # uhorn

	// # (Unicode combining characters have no direct equivalence with
	// # keysyms, where dead keys are defined instead)
	//'\u0303': 0x1e9f, //   r   # combining_tilde
	//'\u0300': 0x1ef2, //   r   # combining_grave
	//'\u0301': 0x1ef3, //   r   # combining_acute
	'\u0309': 0x1efe, //   r   # combining_hook
	'\u0323': 0x1eff, //   r   # combining_belowdot

	// # These probably should be added to the X11 standard properly,
	// # as they could be of use for Vietnamese input methods.
	//'\u0323': 0xfe60,    //   f   # dead_belowdot
	//'\u0309': 0xfe61,    //   f   # dead_hook
	'\u031b': 0xfe62,    //   f   # dead_horn
	'\u2080': 0x1002080, //  u   # zerosubscript
	'\u2081': 0x1002081, //  u   # onesubscript
	'\u2082': 0x1002082, //  u   # twosubscript
	'\u2083': 0x1002083, //  u   # threesubscript
	'\u2084': 0x1002084, //  u   # foursubscript
	'\u2085': 0x1002085, //  u   # fivesubscript
	'\u2086': 0x1002086, //  u   # sixsubscript
	'\u2087': 0x1002087, //  u   # sevensubscript
	'\u2088': 0x1002088, //  u   # eightsubscript
	//'\u2088': 0x1002089, //  u   # ninesubscript
}
