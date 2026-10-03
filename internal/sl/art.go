package sl

// ASCII art ported from sl.h (SL version 5.02, Copyright 1993,2002,2014 Toyoda Masashi).

const (
	d51Height   = 10
	d51Funnel   = 7
	d51Length   = 83
	d51Patterns = 6

	logoHeight   = 6
	logoFunnel   = 4
	logoLength   = 84
	logoPatterns = 6

	c51Height   = 11
	c51Funnel   = 7
	c51Length   = 87
	c51Patterns = 6
)

const (
	d51Str1 = "      ====        ________                ___________ "
	d51Str2 = "  _D _|  |_______/        \\__I_I_____===__|_________| "
	d51Str3 = "   |(_)---  |   H\\________/ |   |        =|___ ___|   "
	d51Str4 = "   /     |  |   H  |  |     |   |         ||_| |_||   "
	d51Str5 = "  |      |  |   H  |__--------------------| [___] |   "
	d51Str6 = "  | ________|___H__/__|_____/[][]~\\_______|       |   "
	d51Str7 = "  |/ |   |-----------I_____I [][] []  D   |=======|__ "

	d51Whl11 = "__/ =| o |=-~~\\  /~~\\  /~~\\  /~~\\ ____Y___________|__ "
	d51Whl12 = " |/-=|___|=    ||    ||    ||    |_____/~\\___/        "
	d51Whl13 = "  \\_/      \\O=====O=====O=====O_/      \\_/            "

	d51Whl21 = "__/ =| o |=-~~\\  /~~\\  /~~\\  /~~\\ ____Y___________|__ "
	d51Whl22 = " |/-=|___|=O=====O=====O=====O   |_____/~\\___/        "
	d51Whl23 = "  \\_/      \\__/  \\__/  \\__/  \\__/      \\_/            "

	d51Whl31 = "__/ =| o |=-O=====O=====O=====O \\ ____Y___________|__ "
	d51Whl32 = " |/-=|___|=    ||    ||    ||    |_____/~\\___/        "
	d51Whl33 = "  \\_/      \\__/  \\__/  \\__/  \\__/      \\_/            "

	d51Whl41 = "__/ =| o |=-~O=====O=====O=====O\\ ____Y___________|__ "
	d51Whl42 = " |/-=|___|=    ||    ||    ||    |_____/~\\___/        "
	d51Whl43 = "  \\_/      \\__/  \\__/  \\__/  \\__/      \\_/            "

	d51Whl51 = "__/ =| o |=-~~\\  /~~\\  /~~\\  /~~\\ ____Y___________|__ "
	d51Whl52 = " |/-=|___|=   O=====O=====O=====O|_____/~\\___/        "
	d51Whl53 = "  \\_/      \\__/  \\__/  \\__/  \\__/      \\_/            "

	d51Whl61 = "__/ =| o |=-~~\\  /~~\\  /~~\\  /~~\\ ____Y___________|__ "
	d51Whl62 = " |/-=|___|=    ||    ||    ||    |_____/~\\___/        "
	d51Whl63 = "  \\_/      \\_O=====O=====O=====O/      \\_/            "

	d51Del = "                                                      "

	coal01 = "                              "
	coal02 = "                              "
	coal03 = "    _________________         "
	coal04 = "   _|                \\_____A  "
	coal05 = " =|                        |  "
	coal06 = " -|                        |  "
	coal07 = "__|________________________|_ "
	coal08 = "|__________________________|_ "
	coal09 = "   |_D__D__D_|  |_D__D__D_|   "
	coal10 = "    \\_/   \\_/    \\_/   \\_/    "

	coalDel = "                              "

	logo1 = "     ++      +------ "
	logo2 = "     ||      |+-+ |  "
	logo3 = "   /---------|| | |  "
	logo4 = "  + ========  +-+ |  "

	lWhl11 = " _|--O========O~\\-+  "
	lWhl12 = "//// \\_/      \\_/    "

	lWhl21 = " _|--/O========O\\-+  "
	lWhl22 = "//// \\_/      \\_/    "

	lWhl31 = " _|--/~O========O-+  "
	lWhl32 = "//// \\_/      \\_/    "

	lWhl41 = " _|--/~\\------/~\\-+  "
	lWhl42 = "//// \\_O========O    "

	lWhl51 = " _|--/~\\------/~\\-+  "
	lWhl52 = "//// \\O========O/    "

	lWhl61 = " _|--/~\\------/~\\-+  "
	lWhl62 = "//// O========O_/    "

	lCoal1 = "____                 "
	lCoal2 = "|   \\@@@@@@@@@@@     "
	lCoal3 = "|    \\@@@@@@@@@@@@@_ "
	lCoal4 = "|                  | "
	lCoal5 = "|__________________| "
	lCoal6 = "   (O)       (O)     "

	lCar1 = "____________________ "
	lCar2 = "|  ___ ___ ___ ___ | "
	lCar3 = "|  |_| |_| |_| |_| | "
	lCar4 = "|__________________| "
	lCar5 = "|__________________| "
	lCar6 = "   (O)        (O)    "

	delLn = "                     "

	c51Del = "                                                       "

	c51Str1 = "        ___                                            "
	c51Str2 = "       _|_|_  _     __       __             ___________"
	c51Str3 = "    D__/   \\_(_)___|  |__H__|  |_____I_Ii_()|_________|"
	c51Str4 = "     | `---'   |:: `--'  H  `--'         |  |___ ___|  "
	c51Str5 = "    +|~~~~~~~~++::~~~~~~~H~~+=====+~~~~~~|~~||_| |_||  "
	c51Str6 = "    ||        | ::       H  +=====+      |  |::  ...|  "
	c51Str7 = "|    | _______|_::-----------------[][]-----|       |  "

	c51Wh61 = "| /~~ ||   |-----/~~~~\\  /[I_____I][][] --|||_______|__"
	c51Wh62 = "------'|oOo|==[]=-     ||      ||      |  ||=======_|__"
	c51Wh63 = "/~\\____|___|/~\\_|   O=======O=======O  |__|+-/~\\_|     "
	c51Wh64 = "\\_/         \\_/  \\____/  \\____/  \\____/      \\_/       "

	c51Wh51 = "| /~~ ||   |-----/~~~~\\  /[I_____I][][] --|||_______|__"
	c51Wh52 = "------'|oOo|===[]=-    ||      ||      |  ||=======_|__"
	c51Wh53 = "/~\\____|___|/~\\_|    O=======O=======O |__|+-/~\\_|     "
	c51Wh54 = "\\_/         \\_/  \\____/  \\____/  \\____/      \\_/       "

	c51Wh41 = "| /~~ ||   |-----/~~~~\\  /[I_____I][][] --|||_______|__"
	c51Wh42 = "------'|oOo|===[]=- O=======O=======O  |  ||=======_|__"
	c51Wh43 = "/~\\____|___|/~\\_|      ||      ||      |__|+-/~\\_|     "
	c51Wh44 = "\\_/         \\_/  \\____/  \\____/  \\____/      \\_/       "

	c51Wh31 = "| /~~ ||   |-----/~~~~\\  /[I_____I][][] --|||_______|__"
	c51Wh32 = "------'|oOo|==[]=- O=======O=======O   |  ||=======_|__"
	c51Wh33 = "/~\\____|___|/~\\_|      ||      ||      |__|+-/~\\_|     "
	c51Wh34 = "\\_/         \\_/  \\____/  \\____/  \\____/      \\_/       "

	c51Wh21 = "| /~~ ||   |-----/~~~~\\  /[I_____I][][] --|||_______|__"
	c51Wh22 = "------'|oOo|=[]=- O=======O=======O    |  ||=======_|__"
	c51Wh23 = "/~\\____|___|/~\\_|      ||      ||      |__|+-/~\\_|     "
	c51Wh24 = "\\_/         \\_/  \\____/  \\____/  \\____/      \\_/       "

	c51Wh11 = "| /~~ ||   |-----/~~~~\\  /[I_____I][][] --|||_______|__"
	c51Wh12 = "------'|oOo|=[]=-      ||      ||      |  ||=======_|__"
	c51Wh13 = "/~\\____|___|/~\\_|  O=======O=======O   |__|+-/~\\_|     "
	c51Wh14 = "\\_/         \\_/  \\____/  \\____/  \\____/      \\_/       "
)

var d51 = [d51Patterns][d51Height + 1]string{
	{d51Str1, d51Str2, d51Str3, d51Str4, d51Str5, d51Str6, d51Str7, d51Whl11, d51Whl12, d51Whl13, d51Del},
	{d51Str1, d51Str2, d51Str3, d51Str4, d51Str5, d51Str6, d51Str7, d51Whl21, d51Whl22, d51Whl23, d51Del},
	{d51Str1, d51Str2, d51Str3, d51Str4, d51Str5, d51Str6, d51Str7, d51Whl31, d51Whl32, d51Whl33, d51Del},
	{d51Str1, d51Str2, d51Str3, d51Str4, d51Str5, d51Str6, d51Str7, d51Whl41, d51Whl42, d51Whl43, d51Del},
	{d51Str1, d51Str2, d51Str3, d51Str4, d51Str5, d51Str6, d51Str7, d51Whl51, d51Whl52, d51Whl53, d51Del},
	{d51Str1, d51Str2, d51Str3, d51Str4, d51Str5, d51Str6, d51Str7, d51Whl61, d51Whl62, d51Whl63, d51Del},
}

var d51Coal = [d51Height + 1]string{
	coal01, coal02, coal03, coal04, coal05, coal06, coal07, coal08, coal09, coal10, coalDel,
}

var c51 = [c51Patterns][c51Height + 1]string{
	{c51Str1, c51Str2, c51Str3, c51Str4, c51Str5, c51Str6, c51Str7, c51Wh11, c51Wh12, c51Wh13, c51Wh14, c51Del},
	{c51Str1, c51Str2, c51Str3, c51Str4, c51Str5, c51Str6, c51Str7, c51Wh21, c51Wh22, c51Wh23, c51Wh24, c51Del},
	{c51Str1, c51Str2, c51Str3, c51Str4, c51Str5, c51Str6, c51Str7, c51Wh31, c51Wh32, c51Wh33, c51Wh34, c51Del},
	{c51Str1, c51Str2, c51Str3, c51Str4, c51Str5, c51Str6, c51Str7, c51Wh41, c51Wh42, c51Wh43, c51Wh44, c51Del},
	{c51Str1, c51Str2, c51Str3, c51Str4, c51Str5, c51Str6, c51Str7, c51Wh51, c51Wh52, c51Wh53, c51Wh54, c51Del},
	{c51Str1, c51Str2, c51Str3, c51Str4, c51Str5, c51Str6, c51Str7, c51Wh61, c51Wh62, c51Wh63, c51Wh64, c51Del},
}

var c51Coal = [c51Height + 1]string{
	coalDel, coal01, coal02, coal03, coal04, coal05, coal06, coal07, coal08, coal09, coal10, coalDel,
}

var logo = [logoPatterns][logoHeight + 1]string{
	{logo1, logo2, logo3, logo4, lWhl11, lWhl12, delLn},
	{logo1, logo2, logo3, logo4, lWhl21, lWhl22, delLn},
	{logo1, logo2, logo3, logo4, lWhl31, lWhl32, delLn},
	{logo1, logo2, logo3, logo4, lWhl41, lWhl42, delLn},
	{logo1, logo2, logo3, logo4, lWhl51, lWhl52, delLn},
	{logo1, logo2, logo3, logo4, lWhl61, lWhl62, delLn},
}

var logoCoal = [logoHeight + 1]string{lCoal1, lCoal2, lCoal3, lCoal4, lCoal5, lCoal6, delLn}

var logoCar = [logoHeight + 1]string{lCar1, lCar2, lCar3, lCar4, lCar5, lCar6, delLn}

var man = [2][2]string{{"", "(O)"}, {"Help!", "\\O/"}}

const smokePatterns = 16

var smoke = [2][smokePatterns]string{
	{
		"(   )", "(    )", "(    )", "(   )", "(  )",
		"(  )", "( )", "( )", "()", "()",
		"O", "O", "O", "O", "O",
		" ",
	},
	{
		"(@@@)", "(@@@@)", "(@@@@)", "(@@@)", "(@@)",
		"(@@)", "(@)", "(@)", "@@", "@@",
		"@", "@", "@", "@", "@",
		" ",
	},
}

var smokeEraser = [smokePatterns]string{
	"     ", "      ", "      ", "     ", "    ",
	"    ", "   ", "   ", "  ", "  ",
	" ", " ", " ", " ", " ",
	" ",
}

var smokeDY = [smokePatterns]int{2, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

var smokeDX = [smokePatterns]int{-2, -1, 0, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 3, 3, 3}
