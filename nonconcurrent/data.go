package main

// 5v5 => 8.5s
// 10v10 => 5m30s
// 15v15

//var activeTeam League = League{
//    Team{"Simba", 111},
//    Team{"shahreaver", 124},
//    Team{"retrovirus", 126},
//    Team{"ChAmP", 120},
//    Team{"NYRB", 121},
//
//    Team{"BlessMabZ", 116},
//    Team{"daz", 125},
//    Team{"TubyAardvark", 118},
//    //Team{"pop123007", 125},
//    //Team{"Shaunak", 123},
//
//    //Team{"Reagan420", 121},
//    //Team{"BePrepared", 117},
//    //Team{"raiden", 118},
//    //Team{"ed", 122},
//    //Team{"Z", 119},
//
//    //Team{"Kitemafia", 117},
//    //Team{"smithsaga", 125},
//    //Team{"hell", 123},
//    //Team{"Kushal", 118},
//    //Team{"kairoschris", 125},
//
//    //Team{"fcbarca", 111},
//    //Team{"fsugio305", 107},
//    //Team{"Mahdi", 114},
//    //Team{"nads0", 119},
//    //Team{"naman", 111},
//
//    //Team{"kilkr", 118},
//    //Team{"Desquexele", 120},
//    //Team{"tranquilityy", 122},
//    //Team{"jan", 121},
//    //Team{"ryan2985", 121},
//}
//
//var opponentTeam League = League{
//    Team{"Roggo", 119},
//    Team{"Clark", 126},
//    Team{"bluj1025", 122},
//    Team{"BMUVAN", 124},
//    Team{"TylerSmith", 122},
//
//    Team{"Dragon2010", 124},
//    Team{"catson", 116},
//    Team{"Luke", 122},
//    //Team{"xav", 111},
//    //Team{"NewJerseyFC", 114},
//
//    //Team{"Coles-goals", 115},
//    //Team{"Thor", 120},
//    //Team{"dustinj", 121},
//    //Team{"MarcG", 114},
//    //Team{"sajj990", 123},
//
//    //Team{"Sr.Alex", 115},
//    //Team{"discgolf410", 114},
//    //Team{"AllanFraga23", 123},
//    //Team{"kingdrogba", 116},
//    //Team{"SEA60", 117},
//
//    //Team{"shreddinsam", 116},
//    //Team{"R.Baggio", 118},
//    //Team{"heiler", 114},
//    //Team{"SCWolverine", 120},
//    //Team{"jay", 114},
//
//    //Team{"jespinoza", 120},
//    //Team{"bell", 121},
//    //Team{"biv", 125},
//    //Team{"devoe", 129},
//    //Team{"same", 131},
//}

func getActiveTeam(n int) League {
	teams := League{
		Team{"Simba", 111},
		Team{"shahreaver", 124},
		Team{"retrovirus", 126},
		Team{"ChAmP", 120},
		Team{"NYRB", 121},

		Team{"BlessMabZ", 116},
		Team{"daz", 125},
		Team{"TubyAardvark", 118},
		Team{"pop123007", 125},
		Team{"Shaunak", 123},

		Team{"Reagan420", 121},
		Team{"BePrepared", 117},
		Team{"raiden", 118},
		Team{"ed", 122},
		Team{"Z", 119},

		Team{"Kitemafia", 117},
		Team{"smithsaga", 125},
		Team{"hell", 123},
		Team{"Kushal", 118},
		Team{"kairoschris", 125},

		Team{"fcbarca", 111},
		Team{"fsugio305", 107},
		Team{"Mahdi", 114},
		Team{"nads0", 119},
		Team{"naman", 111},

		Team{"kilkr", 118},
		Team{"Desquexele", 120},
		Team{"tranquilityy", 122},
		Team{"jan", 121},
		Team{"ryan2985", 121},
	}

	return (teams)[0:n]
}

func getOpponentTeam(n int) League {
	teams := League{
		Team{"Roggo", 119},
		Team{"Clark", 126},
		Team{"bluj1025", 122},
		Team{"BMUVAN", 124},
		Team{"TylerSmith", 122},

		Team{"Dragon2010", 124},
		Team{"catson", 116},
		Team{"Luke", 122},
		Team{"xav", 111},
		Team{"NewJerseyFC", 114},

		Team{"Coles-goals", 115},
		Team{"Thor", 120},
		Team{"dustinj", 121},
		Team{"MarcG", 114},
		Team{"sajj990", 123},

		Team{"Sr.Alex", 115},
		Team{"discgolf410", 114},
		Team{"AllanFraga23", 123},
		Team{"kingdrogba", 116},
		Team{"SEA60", 117},

		Team{"shreddinsam", 116},
		Team{"R.Baggio", 118},
		Team{"heiler", 114},
		Team{"SCWolverine", 120},
		Team{"jay", 114},

		Team{"jespinoza", 120},
		Team{"bell", 121},
		Team{"biv", 125},
		Team{"devoe", 129},
		Team{"same", 131},
	}

	return (teams)[0:n]
}

var chances Chances = []Chance{
	{16, 4, 40, 60, 0, 0},
	{15, 4, 40, 60, 0, 0},
	{14, 4, 40, 60, 0, 0},
	{13, 4, 40, 60, 0, 0},
	{12, 4, 40, 60, 0, 0},
	{11, 4, 40, 60, 0, 0},
	{10, 4, 40, 60, 0, 0},
	{9, 4, 40, 60, 0, 0},
	{8, 4, 40, 60, 0, 0},
	{7, 4, 40, 60, 0, 0},
	{6, 4, 40, 60, 0, 0},
	{5, 4, 40, 60, 0, 0},
	{4, 3, 40, 60, 0, 0},
	{3, 3, 30, 70, 0, 0},
	{2, 2, 30, 70, 0, 0},
	{1, 1, 30, 60, 10, 0},
	{0, 0, 20, 70, 10, 0},
	{-1, -2, 20, 60, 20, 0},
	{-2, -2, 20, 60, 20, 0},
	{-3, -4, 20, 50, 30, 0},
	{-4, -4, 20, 50, 30, 0},
	{-5, -5, 20, 50, 20, 10},
	{-6, -6, 10, 60, 20, 10},
	{-7, -8, 10, 50, 30, 10},
	{-8, -8, 10, 50, 30, 10},
	{-9, -9, 10, 30, 40, 20},
	{-10, -11, 10, 20, 50, 20},
	{-11, -11, 10, 20, 50, 20},
	{-12, -12, 0, 20, 60, 20},
	{-13, -12, 0, 20, 60, 20},
	{-14, -12, 0, 20, 60, 20},
	{-15, -12, 0, 20, 60, 20},
	{-16, -12, 0, 20, 60, 20},
}

var expectedGoals Goals = []ExpectedGoals{
	{16, 15, 20, 0},
	{15, 15, 20, 0},
	{14, 15, 20, 0},
	{13, 15, 20, 0},
	{12, 15, 20, 0},
	{11, 15, 20, 0},
	{10, 15, 20, 0},
	{9, 15, 20, 0},
	{8, 15, 20, 0},
	{7, 15, 20, 0},
	{6, 15, 20, 0},
	{5, 15, 20, 0},
	{4, 15, 20, 0},
	{3, 14, 20, 0},
	{2, 14, 18, 0},
	{1, 13, 17, 0},
	{0, 13, 15, 0},
	{-1, 12, 15, 0},
	{-2, 12, 15, 0},
	{-3, 11, 13, 0},
	{-4, 11, 13, 0},
	{-5, 10, 12, 0},
	{-6, 9, 10, 0},
	{-7, 6, 10, 0},
	{-8, 5, 10, 0},
	{-9, 4, 8, 0},
	{-10, 3, 8, 0},
	{-11, 3, 7, 0},
	{-12, 2, 7, 0},
	{-13, 2, 7, 0},
	{-14, 2, 7, 0},
	{-15, 2, 7, 0},
	{-16, 2, 7, 0},
}

//var chances Chances = []Chance{
//    {-15,10,20,50,20},
//    {-14,10,20,50,20},
//    {-13,10,20,50,20},
//    {-12,10,50,30,10},
//    {-11,10,50,30,10},
//    {-10,10,50,30,10},
//    {-9,10,50,30,10},
//    {-8,10,60,20,10},
//    {-7,10,60,20,10},
//    {-6,10,60,20,10},
//    {-5,20,50,20,10},
//    {-4,20,50,30,0},
//    {-3,20,50,30,0},
//    {-2,20,60,20,0},
//    {-1,20,60,20,0},
//    {0,20,70,10,0},
//    {1,30,60,10,0},
//    {2,30,70,0,0},
//    {3,30,70,0,0},
//    {4,40,60,0,0},
//    {5,40,60,0,0},
//    {6,40,60,0,0},
//    {7,40,60,0,0},
//    {8,50,50,0,0},
//    {9,50,50,0,0},
//    {10,50,50,0,0},
//    {11,60,40,0,0},
//    {12,60,40,0,0},
//    {13,60,40,0,0},
//    {14,60,40,0,0},
//    {15,60,40,0,0},
//}
