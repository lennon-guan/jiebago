package posseg

var (
	ProbStart = make(map[StateTag]float64)
)

func init() {
	ProbStart[StateTag{'B', "a"}] = -4.762305214596967
	ProbStart[StateTag{'B', "ad"}] = -6.680066036784177
	ProbStart[StateTag{'B', "ag"}] = -3.14e+100
	ProbStart[StateTag{'B', "an"}] = -8.697083223018778
	ProbStart[StateTag{'B', "b"}] = -5.018374362109218
	ProbStart[StateTag{'B', "bg"}] = -3.14e+100
	ProbStart[StateTag{'B', "c"}] = -3.423880184954888
	ProbStart[StateTag{'B', "d"}] = -3.9750475297585357
	ProbStart[StateTag{'B', "df"}] = -8.888974230828882
	ProbStart[StateTag{'B', "dg"}] = -3.14e+100
	ProbStart[StateTag{'B', "e"}] = -8.563551830394255
	ProbStart[StateTag{'B', "en"}] = -3.14e+100
	ProbStart[StateTag{'B', "f"}] = -5.491630418482717
	ProbStart[StateTag{'B', "g"}] = -3.14e+100
	ProbStart[StateTag{'B', "h"}] = -13.533365129970255
	ProbStart[StateTag{'B', "i"}] = -6.1157847275557105
	ProbStart[StateTag{'B', "in"}] = -3.14e+100
	ProbStart[StateTag{'B', "j"}] = -5.0576191284681915
	ProbStart[StateTag{'B', "jn"}] = -3.14e+100
	ProbStart[StateTag{'B', "k"}] = -3.14e+100
	ProbStart[StateTag{'B', "l"}] = -4.905883584659895
	ProbStart[StateTag{'B', "ln"}] = -3.14e+100
	ProbStart[StateTag{'B', "m"}] = -3.6524299819046386
	ProbStart[StateTag{'B', "mg"}] = -3.14e+100
	ProbStart[StateTag{'B', "mq"}] = -6.78695300139688
	ProbStart[StateTag{'B', "n"}] = -1.6966257797548328
	ProbStart[StateTag{'B', "ng"}] = -3.14e+100
	ProbStart[StateTag{'B', "nr"}] = -2.2310495913769506
	ProbStart[StateTag{'B', "nrfg"}] = -5.873722175405573
	ProbStart[StateTag{'B', "nrt"}] = -4.985642733519195
	ProbStart[StateTag{'B', "ns"}] = -2.8228438314969213
	ProbStart[StateTag{'B', "nt"}] = -4.846091668182416
	ProbStart[StateTag{'B', "nz"}] = -3.94698846057672
	ProbStart[StateTag{'B', "o"}] = -8.433498702146057
	ProbStart[StateTag{'B', "p"}] = -4.200984132085048
	ProbStart[StateTag{'B', "q"}] = -6.998123858956596
	ProbStart[StateTag{'B', "qe"}] = -3.14e+100
	ProbStart[StateTag{'B', "qg"}] = -3.14e+100
	ProbStart[StateTag{'B', "r"}] = -3.4098187790818413
	ProbStart[StateTag{'B', "rg"}] = -3.14e+100
	ProbStart[StateTag{'B', "rr"}] = -12.434752841302146
	ProbStart[StateTag{'B', "rz"}] = -7.946116471570005
	ProbStart[StateTag{'B', "s"}] = -5.522673590839954
	ProbStart[StateTag{'B', "t"}] = -3.3647479094528574
	ProbStart[StateTag{'B', "tg"}] = -3.14e+100
	ProbStart[StateTag{'B', "u"}] = -9.163917277503234
	ProbStart[StateTag{'B', "ud"}] = -3.14e+100
	ProbStart[StateTag{'B', "ug"}] = -3.14e+100
	ProbStart[StateTag{'B', "uj"}] = -3.14e+100
	ProbStart[StateTag{'B', "ul"}] = -3.14e+100
	ProbStart[StateTag{'B', "uv"}] = -3.14e+100
	ProbStart[StateTag{'B', "uz"}] = -3.14e+100
	ProbStart[StateTag{'B', "v"}] = -2.6740584874265685
	ProbStart[StateTag{'B', "vd"}] = -9.044728760238115
	ProbStart[StateTag{'B', "vg"}] = -3.14e+100
	ProbStart[StateTag{'B', "vi"}] = -12.434752841302146
	ProbStart[StateTag{'B', "vn"}] = -4.3315610890163585
	ProbStart[StateTag{'B', "vq"}] = -12.147070768850364
	ProbStart[StateTag{'B', "w"}] = -3.14e+100
	ProbStart[StateTag{'B', "x"}] = -3.14e+100
	ProbStart[StateTag{'B', "y"}] = -9.844485675856319
	ProbStart[StateTag{'B', "yg"}] = -3.14e+100
	ProbStart[StateTag{'B', "z"}] = -7.045681111485645
	ProbStart[StateTag{'B', "zg"}] = -3.14e+100
	ProbStart[StateTag{'E', "a"}] = -3.14e+100
	ProbStart[StateTag{'E', "ad"}] = -3.14e+100
	ProbStart[StateTag{'E', "ag"}] = -3.14e+100
	ProbStart[StateTag{'E', "an"}] = -3.14e+100
	ProbStart[StateTag{'E', "b"}] = -3.14e+100
	ProbStart[StateTag{'E', "bg"}] = -3.14e+100
	ProbStart[StateTag{'E', "c"}] = -3.14e+100
	ProbStart[StateTag{'E', "d"}] = -3.14e+100
	ProbStart[StateTag{'E', "df"}] = -3.14e+100
	ProbStart[StateTag{'E', "dg"}] = -3.14e+100
	ProbStart[StateTag{'E', "e"}] = -3.14e+100
	ProbStart[StateTag{'E', "en"}] = -3.14e+100
	ProbStart[StateTag{'E', "f"}] = -3.14e+100
	ProbStart[StateTag{'E', "g"}] = -3.14e+100
	ProbStart[StateTag{'E', "h"}] = -3.14e+100
	ProbStart[StateTag{'E', "i"}] = -3.14e+100
	ProbStart[StateTag{'E', "in"}] = -3.14e+100
	ProbStart[StateTag{'E', "j"}] = -3.14e+100
	ProbStart[StateTag{'E', "jn"}] = -3.14e+100
	ProbStart[StateTag{'E', "k"}] = -3.14e+100
	ProbStart[StateTag{'E', "l"}] = -3.14e+100
	ProbStart[StateTag{'E', "ln"}] = -3.14e+100
	ProbStart[StateTag{'E', "m"}] = -3.14e+100
	ProbStart[StateTag{'E', "mg"}] = -3.14e+100
	ProbStart[StateTag{'E', "mq"}] = -3.14e+100
	ProbStart[StateTag{'E', "n"}] = -3.14e+100
	ProbStart[StateTag{'E', "ng"}] = -3.14e+100
	ProbStart[StateTag{'E', "nr"}] = -3.14e+100
	ProbStart[StateTag{'E', "nrfg"}] = -3.14e+100
	ProbStart[StateTag{'E', "nrt"}] = -3.14e+100
	ProbStart[StateTag{'E', "ns"}] = -3.14e+100
	ProbStart[StateTag{'E', "nt"}] = -3.14e+100
	ProbStart[StateTag{'E', "nz"}] = -3.14e+100
	ProbStart[StateTag{'E', "o"}] = -3.14e+100
	ProbStart[StateTag{'E', "p"}] = -3.14e+100
	ProbStart[StateTag{'E', "q"}] = -3.14e+100
	ProbStart[StateTag{'E', "qe"}] = -3.14e+100
	ProbStart[StateTag{'E', "qg"}] = -3.14e+100
	ProbStart[StateTag{'E', "r"}] = -3.14e+100
	ProbStart[StateTag{'E', "rg"}] = -3.14e+100
	ProbStart[StateTag{'E', "rr"}] = -3.14e+100
	ProbStart[StateTag{'E', "rz"}] = -3.14e+100
	ProbStart[StateTag{'E', "s"}] = -3.14e+100
	ProbStart[StateTag{'E', "t"}] = -3.14e+100
	ProbStart[StateTag{'E', "tg"}] = -3.14e+100
	ProbStart[StateTag{'E', "u"}] = -3.14e+100
	ProbStart[StateTag{'E', "ud"}] = -3.14e+100
	ProbStart[StateTag{'E', "ug"}] = -3.14e+100
	ProbStart[StateTag{'E', "uj"}] = -3.14e+100
	ProbStart[StateTag{'E', "ul"}] = -3.14e+100
	ProbStart[StateTag{'E', "uv"}] = -3.14e+100
	ProbStart[StateTag{'E', "uz"}] = -3.14e+100
	ProbStart[StateTag{'E', "v"}] = -3.14e+100
	ProbStart[StateTag{'E', "vd"}] = -3.14e+100
	ProbStart[StateTag{'E', "vg"}] = -3.14e+100
	ProbStart[StateTag{'E', "vi"}] = -3.14e+100
	ProbStart[StateTag{'E', "vn"}] = -3.14e+100
	ProbStart[StateTag{'E', "vq"}] = -3.14e+100
	ProbStart[StateTag{'E', "w"}] = -3.14e+100
	ProbStart[StateTag{'E', "x"}] = -3.14e+100
	ProbStart[StateTag{'E', "y"}] = -3.14e+100
	ProbStart[StateTag{'E', "yg"}] = -3.14e+100
	ProbStart[StateTag{'E', "z"}] = -3.14e+100
	ProbStart[StateTag{'E', "zg"}] = -3.14e+100
	ProbStart[StateTag{'M', "a"}] = -3.14e+100
	ProbStart[StateTag{'M', "ad"}] = -3.14e+100
	ProbStart[StateTag{'M', "ag"}] = -3.14e+100
	ProbStart[StateTag{'M', "an"}] = -3.14e+100
	ProbStart[StateTag{'M', "b"}] = -3.14e+100
	ProbStart[StateTag{'M', "bg"}] = -3.14e+100
	ProbStart[StateTag{'M', "c"}] = -3.14e+100
	ProbStart[StateTag{'M', "d"}] = -3.14e+100
	ProbStart[StateTag{'M', "df"}] = -3.14e+100
	ProbStart[StateTag{'M', "dg"}] = -3.14e+100
	ProbStart[StateTag{'M', "e"}] = -3.14e+100
	ProbStart[StateTag{'M', "en"}] = -3.14e+100
	ProbStart[StateTag{'M', "f"}] = -3.14e+100
	ProbStart[StateTag{'M', "g"}] = -3.14e+100
	ProbStart[StateTag{'M', "h"}] = -3.14e+100
	ProbStart[StateTag{'M', "i"}] = -3.14e+100
	ProbStart[StateTag{'M', "in"}] = -3.14e+100
	ProbStart[StateTag{'M', "j"}] = -3.14e+100
	ProbStart[StateTag{'M', "jn"}] = -3.14e+100
	ProbStart[StateTag{'M', "k"}] = -3.14e+100
	ProbStart[StateTag{'M', "l"}] = -3.14e+100
	ProbStart[StateTag{'M', "ln"}] = -3.14e+100
	ProbStart[StateTag{'M', "m"}] = -3.14e+100
	ProbStart[StateTag{'M', "mg"}] = -3.14e+100
	ProbStart[StateTag{'M', "mq"}] = -3.14e+100
	ProbStart[StateTag{'M', "n"}] = -3.14e+100
	ProbStart[StateTag{'M', "ng"}] = -3.14e+100
	ProbStart[StateTag{'M', "nr"}] = -3.14e+100
	ProbStart[StateTag{'M', "nrfg"}] = -3.14e+100
	ProbStart[StateTag{'M', "nrt"}] = -3.14e+100
	ProbStart[StateTag{'M', "ns"}] = -3.14e+100
	ProbStart[StateTag{'M', "nt"}] = -3.14e+100
	ProbStart[StateTag{'M', "nz"}] = -3.14e+100
	ProbStart[StateTag{'M', "o"}] = -3.14e+100
	ProbStart[StateTag{'M', "p"}] = -3.14e+100
	ProbStart[StateTag{'M', "q"}] = -3.14e+100
	ProbStart[StateTag{'M', "qe"}] = -3.14e+100
	ProbStart[StateTag{'M', "qg"}] = -3.14e+100
	ProbStart[StateTag{'M', "r"}] = -3.14e+100
	ProbStart[StateTag{'M', "rg"}] = -3.14e+100
	ProbStart[StateTag{'M', "rr"}] = -3.14e+100
	ProbStart[StateTag{'M', "rz"}] = -3.14e+100
	ProbStart[StateTag{'M', "s"}] = -3.14e+100
	ProbStart[StateTag{'M', "t"}] = -3.14e+100
	ProbStart[StateTag{'M', "tg"}] = -3.14e+100
	ProbStart[StateTag{'M', "u"}] = -3.14e+100
	ProbStart[StateTag{'M', "ud"}] = -3.14e+100
	ProbStart[StateTag{'M', "ug"}] = -3.14e+100
	ProbStart[StateTag{'M', "uj"}] = -3.14e+100
	ProbStart[StateTag{'M', "ul"}] = -3.14e+100
	ProbStart[StateTag{'M', "uv"}] = -3.14e+100
	ProbStart[StateTag{'M', "uz"}] = -3.14e+100
	ProbStart[StateTag{'M', "v"}] = -3.14e+100
	ProbStart[StateTag{'M', "vd"}] = -3.14e+100
	ProbStart[StateTag{'M', "vg"}] = -3.14e+100
	ProbStart[StateTag{'M', "vi"}] = -3.14e+100
	ProbStart[StateTag{'M', "vn"}] = -3.14e+100
	ProbStart[StateTag{'M', "vq"}] = -3.14e+100
	ProbStart[StateTag{'M', "w"}] = -3.14e+100
	ProbStart[StateTag{'M', "x"}] = -3.14e+100
	ProbStart[StateTag{'M', "y"}] = -3.14e+100
	ProbStart[StateTag{'M', "yg"}] = -3.14e+100
	ProbStart[StateTag{'M', "z"}] = -3.14e+100
	ProbStart[StateTag{'M', "zg"}] = -3.14e+100
	ProbStart[StateTag{'S', "a"}] = -3.9025396831295227
	ProbStart[StateTag{'S', "ad"}] = -11.048458480182255
	ProbStart[StateTag{'S', "ag"}] = -6.954113917960154
	ProbStart[StateTag{'S', "an"}] = -12.84021794941031
	ProbStart[StateTag{'S', "b"}] = -6.472888763970454
	ProbStart[StateTag{'S', "bg"}] = -3.14e+100
	ProbStart[StateTag{'S', "c"}] = -4.786966795861212
	ProbStart[StateTag{'S', "d"}] = -3.903919764181873
	ProbStart[StateTag{'S', "df"}] = -3.14e+100
	ProbStart[StateTag{'S', "dg"}] = -8.948397651299683
	ProbStart[StateTag{'S', "e"}] = -5.942513006281674
	ProbStart[StateTag{'S', "en"}] = -3.14e+100
	ProbStart[StateTag{'S', "f"}] = -5.194820249981676
	ProbStart[StateTag{'S', "g"}] = -6.507826815331734
	ProbStart[StateTag{'S', "h"}] = -8.650563207383884
	ProbStart[StateTag{'S', "i"}] = -3.14e+100
	ProbStart[StateTag{'S', "in"}] = -3.14e+100
	ProbStart[StateTag{'S', "j"}] = -4.911992119644354
	ProbStart[StateTag{'S', "jn"}] = -3.14e+100
	ProbStart[StateTag{'S', "k"}] = -6.940320595827818
	ProbStart[StateTag{'S', "l"}] = -3.14e+100
	ProbStart[StateTag{'S', "ln"}] = -3.14e+100
	ProbStart[StateTag{'S', "m"}] = -3.269200652116097
	ProbStart[StateTag{'S', "mg"}] = -10.825314928868044
	ProbStart[StateTag{'S', "mq"}] = -3.14e+100
	ProbStart[StateTag{'S', "n"}] = -3.8551483897645107
	ProbStart[StateTag{'S', "ng"}] = -4.913434861102905
	ProbStart[StateTag{'S', "nr"}] = -4.483663103956885
	ProbStart[StateTag{'S', "nrfg"}] = -3.14e+100
	ProbStart[StateTag{'S', "nrt"}] = -3.14e+100
	ProbStart[StateTag{'S', "ns"}] = -3.14e+100
	ProbStart[StateTag{'S', "nt"}] = -12.147070768850364
	ProbStart[StateTag{'S', "nz"}] = -3.14e+100
	ProbStart[StateTag{'S', "o"}] = -8.464460927750023
	ProbStart[StateTag{'S', "p"}] = -2.9868401813596317
	ProbStart[StateTag{'S', "q"}] = -4.888658618255058
	ProbStart[StateTag{'S', "qe"}] = -3.14e+100
	ProbStart[StateTag{'S', "qg"}] = -3.14e+100
	ProbStart[StateTag{'S', "r"}] = -2.7635336784127853
	ProbStart[StateTag{'S', "rg"}] = -10.275268591948773
	ProbStart[StateTag{'S', "rr"}] = -3.14e+100
	ProbStart[StateTag{'S', "rz"}] = -3.14e+100
	ProbStart[StateTag{'S', "s"}] = -3.14e+100
	ProbStart[StateTag{'S', "t"}] = -3.14e+100
	ProbStart[StateTag{'S', "tg"}] = -6.272842531880403
	ProbStart[StateTag{'S', "u"}] = -6.940320595827818
	ProbStart[StateTag{'S', "ud"}] = -7.728230161053767
	ProbStart[StateTag{'S', "ug"}] = -7.5394037026636855
	ProbStart[StateTag{'S', "uj"}] = -6.85251045118004
	ProbStart[StateTag{'S', "ul"}] = -8.4153713175535
	ProbStart[StateTag{'S', "uv"}] = -8.15808672228609
	ProbStart[StateTag{'S', "uz"}] = -9.299258625372996
	ProbStart[StateTag{'S', "v"}] = -3.053292303412302
	ProbStart[StateTag{'S', "vd"}] = -3.14e+100
	ProbStart[StateTag{'S', "vg"}] = -5.9430181843676895
	ProbStart[StateTag{'S', "vi"}] = -3.14e+100
	ProbStart[StateTag{'S', "vn"}] = -11.453923588290419
	ProbStart[StateTag{'S', "vq"}] = -3.14e+100
	ProbStart[StateTag{'S', "w"}] = -3.14e+100
	ProbStart[StateTag{'S', "x"}] = -8.427419656069674
	ProbStart[StateTag{'S', "y"}] = -6.1970794699489575
	ProbStart[StateTag{'S', "yg"}] = -13.533365129970255
	ProbStart[StateTag{'S', "z"}] = -3.14e+100
	ProbStart[StateTag{'S', "zg"}] = -3.14e+100
}
