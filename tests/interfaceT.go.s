"".Binary.String STEXT size=126 args=0x18 locals=0x40
	0x0000 00000 (interfaceT.go:14)	TEXT	"".Binary.String(SB), ABIInternal, $64-24
	0x0000 00000 (interfaceT.go:14)	MOVQ	(TLS), CX
	0x0009 00009 (interfaceT.go:14)	CMPQ	SP, 16(CX)
	0x000d 00013 (interfaceT.go:14)	PCDATA	$0, $-2
	0x000d 00013 (interfaceT.go:14)	JLS	119
	0x000f 00015 (interfaceT.go:14)	PCDATA	$0, $-1
	0x000f 00015 (interfaceT.go:14)	SUBQ	$64, SP
	0x0013 00019 (interfaceT.go:14)	MOVQ	BP, 56(SP)
	0x0018 00024 (interfaceT.go:14)	LEAQ	56(SP), BP
	0x001d 00029 (interfaceT.go:14)	PCDATA	$0, $-2
	0x001d 00029 (interfaceT.go:14)	PCDATA	$1, $-2
	0x001d 00029 (interfaceT.go:14)	FUNCDATA	$0, gclocals·568470801006e5c0dc3947ea998fe279(SB)
	0x001d 00029 (interfaceT.go:14)	FUNCDATA	$1, gclocals·2589ca35330fc0fce83503f4569854a0(SB)
	0x001d 00029 (interfaceT.go:14)	FUNCDATA	$2, gclocals·568470801006e5c0dc3947ea998fe279(SB)
	0x001d 00029 (interfaceT.go:14)	PCDATA	$0, $0
	0x001d 00029 (interfaceT.go:14)	PCDATA	$1, $0
	0x001d 00029 (interfaceT.go:14)	XORPS	X0, X0
	0x0020 00032 (interfaceT.go:14)	MOVUPS	X0, "".~r0+80(SP)
	0x0025 00037 (interfaceT.go:15)	MOVQ	"".i+72(SP), AX
	0x002a 00042 (interfaceT.go:15)	MOVQ	AX, (SP)
	0x002e 00046 (interfaceT.go:15)	CALL	"".Binary.Get(SB)
	0x0033 00051 (interfaceT.go:15)	MOVQ	8(SP), AX
	0x0038 00056 (interfaceT.go:15)	MOVQ	AX, ""..autotmp_2+32(SP)
	0x003d 00061 (interfaceT.go:15)	MOVQ	AX, (SP)
	0x0041 00065 (interfaceT.go:15)	MOVQ	$2, 8(SP)
	0x004a 00074 (interfaceT.go:15)	CALL	strconv.FormatUint(SB)
	0x004f 00079 (interfaceT.go:15)	MOVQ	24(SP), AX
	0x0054 00084 (interfaceT.go:15)	PCDATA	$0, $1
	0x0054 00084 (interfaceT.go:15)	MOVQ	16(SP), CX
	0x0059 00089 (interfaceT.go:15)	MOVQ	CX, ""..autotmp_3+40(SP)
	0x005e 00094 (interfaceT.go:15)	MOVQ	AX, ""..autotmp_3+48(SP)
	0x0063 00099 (interfaceT.go:15)	PCDATA	$0, $0
	0x0063 00099 (interfaceT.go:15)	PCDATA	$1, $1
	0x0063 00099 (interfaceT.go:15)	MOVQ	CX, "".~r0+80(SP)
	0x0068 00104 (interfaceT.go:15)	MOVQ	AX, "".~r0+88(SP)
	0x006d 00109 (interfaceT.go:15)	MOVQ	56(SP), BP
	0x0072 00114 (interfaceT.go:15)	ADDQ	$64, SP
	0x0076 00118 (interfaceT.go:15)	RET
	0x0077 00119 (interfaceT.go:15)	NOP
	0x0077 00119 (interfaceT.go:14)	PCDATA	$1, $-1
	0x0077 00119 (interfaceT.go:14)	PCDATA	$0, $-2
	0x0077 00119 (interfaceT.go:14)	CALL	runtime.morestack_noctxt(SB)
	0x007c 00124 (interfaceT.go:14)	PCDATA	$0, $-1
	0x007c 00124 (interfaceT.go:14)	JMP	0
	0x0000 65 48 8b 0c 25 00 00 00 00 48 3b 61 10 76 68 48  eH..%....H;a.vhH
	0x0010 83 ec 40 48 89 6c 24 38 48 8d 6c 24 38 0f 57 c0  ..@H.l$8H.l$8.W.
	0x0020 0f 11 44 24 50 48 8b 44 24 48 48 89 04 24 e8 00  ..D$PH.D$HH..$..
	0x0030 00 00 00 48 8b 44 24 08 48 89 44 24 20 48 89 04  ...H.D$.H.D$ H..
	0x0040 24 48 c7 44 24 08 02 00 00 00 e8 00 00 00 00 48  $H.D$..........H
	0x0050 8b 44 24 18 48 8b 4c 24 10 48 89 4c 24 28 48 89  .D$.H.L$.H.L$(H.
	0x0060 44 24 30 48 89 4c 24 50 48 89 44 24 58 48 8b 6c  D$0H.L$PH.D$XH.l
	0x0070 24 38 48 83 c4 40 c3 e8 00 00 00 00 eb 82        $8H..@........
	rel 5+4 t=17 TLS+0
	rel 47+4 t=8 "".Binary.Get+0
	rel 75+4 t=8 strconv.FormatUint+0
	rel 120+4 t=8 runtime.morestack_noctxt+0
"".Binary.Get STEXT nosplit size=20 args=0x10 locals=0x0
	0x0000 00000 (interfaceT.go:18)	TEXT	"".Binary.Get(SB), NOSPLIT|ABIInternal, $0-16
	0x0000 00000 (interfaceT.go:18)	PCDATA	$0, $-2
	0x0000 00000 (interfaceT.go:18)	PCDATA	$1, $-2
	0x0000 00000 (interfaceT.go:18)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0000 00000 (interfaceT.go:18)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0000 00000 (interfaceT.go:18)	FUNCDATA	$2, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0000 00000 (interfaceT.go:18)	PCDATA	$0, $0
	0x0000 00000 (interfaceT.go:18)	PCDATA	$1, $0
	0x0000 00000 (interfaceT.go:18)	MOVQ	$0, "".~r0+16(SP)
	0x0009 00009 (interfaceT.go:19)	MOVQ	"".i+8(SP), AX
	0x000e 00014 (interfaceT.go:19)	MOVQ	AX, "".~r0+16(SP)
	0x0013 00019 (interfaceT.go:19)	RET
	0x0000 48 c7 44 24 10 00 00 00 00 48 8b 44 24 08 48 89  H.D$.....H.D$.H.
	0x0010 44 24 10 c3                                      D$..
"".main STEXT size=286 args=0x0 locals=0xa0
	0x0000 00000 (interfaceT.go:22)	TEXT	"".main(SB), ABIInternal, $160-0
	0x0000 00000 (interfaceT.go:22)	MOVQ	(TLS), CX
	0x0009 00009 (interfaceT.go:22)	LEAQ	-32(SP), AX
	0x000e 00014 (interfaceT.go:22)	CMPQ	AX, 16(CX)
	0x0012 00018 (interfaceT.go:22)	PCDATA	$0, $-2
	0x0012 00018 (interfaceT.go:22)	JLS	276
	0x0018 00024 (interfaceT.go:22)	PCDATA	$0, $-1
	0x0018 00024 (interfaceT.go:22)	SUBQ	$160, SP
	0x001f 00031 (interfaceT.go:22)	MOVQ	BP, 152(SP)
	0x0027 00039 (interfaceT.go:22)	LEAQ	152(SP), BP
	0x002f 00047 (interfaceT.go:22)	PCDATA	$0, $-2
	0x002f 00047 (interfaceT.go:22)	PCDATA	$1, $-2
	0x002f 00047 (interfaceT.go:22)	FUNCDATA	$0, gclocals·7d2d5fca80364273fb07d5820a76fef4(SB)
	0x002f 00047 (interfaceT.go:22)	FUNCDATA	$1, gclocals·154e988fc21ce4c795057534d7421a72(SB)
	0x002f 00047 (interfaceT.go:22)	FUNCDATA	$2, gclocals·f6aec3988379d2bd21c69c093370a150(SB)
	0x002f 00047 (interfaceT.go:22)	FUNCDATA	$3, "".main.stkobj(SB)
	0x002f 00047 (interfaceT.go:23)	PCDATA	$0, $0
	0x002f 00047 (interfaceT.go:23)	PCDATA	$1, $0
	0x002f 00047 (interfaceT.go:23)	MOVQ	$200, "".b+48(SP)
	0x0038 00056 (interfaceT.go:24)	MOVQ	$200, (SP)
	0x0040 00064 (interfaceT.go:24)	CALL	runtime.convT64(SB)
	0x0045 00069 (interfaceT.go:24)	PCDATA	$0, $1
	0x0045 00069 (interfaceT.go:24)	MOVQ	8(SP), AX
	0x004a 00074 (interfaceT.go:24)	MOVQ	AX, ""..autotmp_4+72(SP)
	0x004f 00079 (interfaceT.go:24)	PCDATA	$0, $2
	0x004f 00079 (interfaceT.go:24)	LEAQ	go.itab."".Binary,"".Stringer(SB), CX
	0x0056 00086 (interfaceT.go:24)	MOVQ	CX, "".s+80(SP)
	0x005b 00091 (interfaceT.go:24)	MOVQ	AX, "".s+88(SP)
	0x0060 00096 (interfaceT.go:25)	PCDATA	$0, $1
	0x0060 00096 (interfaceT.go:25)	TESTB	AL, (CX)
	0x0062 00098 (interfaceT.go:25)	MOVQ	go.itab."".Binary,"".Stringer+24(SB), CX
	0x0069 00105 (interfaceT.go:25)	PCDATA	$0, $0
	0x0069 00105 (interfaceT.go:25)	MOVQ	AX, (SP)
	0x006d 00109 (interfaceT.go:25)	CALL	CX
	0x006f 00111 (interfaceT.go:25)	PCDATA	$0, $1
	0x006f 00111 (interfaceT.go:25)	MOVQ	8(SP), AX
	0x0074 00116 (interfaceT.go:25)	MOVQ	16(SP), CX
	0x0079 00121 (interfaceT.go:25)	MOVQ	AX, ""..autotmp_3+96(SP)
	0x007e 00126 (interfaceT.go:25)	MOVQ	CX, ""..autotmp_3+104(SP)
	0x0083 00131 (interfaceT.go:25)	PCDATA	$0, $0
	0x0083 00131 (interfaceT.go:25)	MOVQ	AX, (SP)
	0x0087 00135 (interfaceT.go:25)	MOVQ	CX, 8(SP)
	0x008c 00140 (interfaceT.go:25)	CALL	runtime.convTstring(SB)
	0x0091 00145 (interfaceT.go:25)	PCDATA	$0, $1
	0x0091 00145 (interfaceT.go:25)	MOVQ	16(SP), AX
	0x0096 00150 (interfaceT.go:25)	PCDATA	$0, $0
	0x0096 00150 (interfaceT.go:25)	PCDATA	$1, $1
	0x0096 00150 (interfaceT.go:25)	MOVQ	AX, ""..autotmp_5+64(SP)
	0x009b 00155 (interfaceT.go:25)	PCDATA	$1, $2
	0x009b 00155 (interfaceT.go:25)	XORPS	X0, X0
	0x009e 00158 (interfaceT.go:25)	MOVUPS	X0, ""..autotmp_2+112(SP)
	0x00a3 00163 (interfaceT.go:25)	PCDATA	$0, $1
	0x00a3 00163 (interfaceT.go:25)	PCDATA	$1, $1
	0x00a3 00163 (interfaceT.go:25)	LEAQ	""..autotmp_2+112(SP), AX
	0x00a8 00168 (interfaceT.go:25)	MOVQ	AX, ""..autotmp_7+56(SP)
	0x00ad 00173 (interfaceT.go:25)	TESTB	AL, (AX)
	0x00af 00175 (interfaceT.go:25)	PCDATA	$0, $2
	0x00af 00175 (interfaceT.go:25)	PCDATA	$1, $0
	0x00af 00175 (interfaceT.go:25)	MOVQ	""..autotmp_5+64(SP), CX
	0x00b4 00180 (interfaceT.go:25)	PCDATA	$0, $3
	0x00b4 00180 (interfaceT.go:25)	LEAQ	type.string(SB), DX
	0x00bb 00187 (interfaceT.go:25)	PCDATA	$0, $2
	0x00bb 00187 (interfaceT.go:25)	MOVQ	DX, ""..autotmp_2+112(SP)
	0x00c0 00192 (interfaceT.go:25)	PCDATA	$0, $1
	0x00c0 00192 (interfaceT.go:25)	MOVQ	CX, ""..autotmp_2+120(SP)
	0x00c5 00197 (interfaceT.go:25)	TESTB	AL, (AX)
	0x00c7 00199 (interfaceT.go:25)	JMP	201
	0x00c9 00201 (interfaceT.go:25)	MOVQ	AX, ""..autotmp_6+128(SP)
	0x00d1 00209 (interfaceT.go:25)	MOVQ	$1, ""..autotmp_6+136(SP)
	0x00dd 00221 (interfaceT.go:25)	MOVQ	$1, ""..autotmp_6+144(SP)
	0x00e9 00233 (interfaceT.go:25)	PCDATA	$0, $0
	0x00e9 00233 (interfaceT.go:25)	MOVQ	AX, (SP)
	0x00ed 00237 (interfaceT.go:25)	MOVQ	$1, 8(SP)
	0x00f6 00246 (interfaceT.go:25)	MOVQ	$1, 16(SP)
	0x00ff 00255 (interfaceT.go:25)	CALL	fmt.Println(SB)
	0x0104 00260 (interfaceT.go:26)	MOVQ	152(SP), BP
	0x010c 00268 (interfaceT.go:26)	ADDQ	$160, SP
	0x0113 00275 (interfaceT.go:26)	RET
	0x0114 00276 (interfaceT.go:26)	NOP
	0x0114 00276 (interfaceT.go:22)	PCDATA	$1, $-1
	0x0114 00276 (interfaceT.go:22)	PCDATA	$0, $-2
	0x0114 00276 (interfaceT.go:22)	CALL	runtime.morestack_noctxt(SB)
	0x0119 00281 (interfaceT.go:22)	PCDATA	$0, $-1
	0x0119 00281 (interfaceT.go:22)	JMP	0
	0x0000 65 48 8b 0c 25 00 00 00 00 48 8d 44 24 e0 48 3b  eH..%....H.D$.H;
	0x0010 41 10 0f 86 fc 00 00 00 48 81 ec a0 00 00 00 48  A.......H......H
	0x0020 89 ac 24 98 00 00 00 48 8d ac 24 98 00 00 00 48  ..$....H..$....H
	0x0030 c7 44 24 30 c8 00 00 00 48 c7 04 24 c8 00 00 00  .D$0....H..$....
	0x0040 e8 00 00 00 00 48 8b 44 24 08 48 89 44 24 48 48  .....H.D$.H.D$HH
	0x0050 8d 0d 00 00 00 00 48 89 4c 24 50 48 89 44 24 58  ......H.L$PH.D$X
	0x0060 84 01 48 8b 0d 00 00 00 00 48 89 04 24 ff d1 48  ..H......H..$..H
	0x0070 8b 44 24 08 48 8b 4c 24 10 48 89 44 24 60 48 89  .D$.H.L$.H.D$`H.
	0x0080 4c 24 68 48 89 04 24 48 89 4c 24 08 e8 00 00 00  L$hH..$H.L$.....
	0x0090 00 48 8b 44 24 10 48 89 44 24 40 0f 57 c0 0f 11  .H.D$.H.D$@.W...
	0x00a0 44 24 70 48 8d 44 24 70 48 89 44 24 38 84 00 48  D$pH.D$pH.D$8..H
	0x00b0 8b 4c 24 40 48 8d 15 00 00 00 00 48 89 54 24 70  .L$@H......H.T$p
	0x00c0 48 89 4c 24 78 84 00 eb 00 48 89 84 24 80 00 00  H.L$x....H..$...
	0x00d0 00 48 c7 84 24 88 00 00 00 01 00 00 00 48 c7 84  .H..$........H..
	0x00e0 24 90 00 00 00 01 00 00 00 48 89 04 24 48 c7 44  $........H..$H.D
	0x00f0 24 08 01 00 00 00 48 c7 44 24 10 01 00 00 00 e8  $.....H.D$......
	0x0100 00 00 00 00 48 8b ac 24 98 00 00 00 48 81 c4 a0  ....H..$....H...
	0x0110 00 00 00 c3 e8 00 00 00 00 e9 e2 fe ff ff        ..............
	rel 5+4 t=17 TLS+0
	rel 65+4 t=8 runtime.convT64+0
	rel 82+4 t=16 go.itab."".Binary,"".Stringer+0
	rel 101+4 t=16 go.itab."".Binary,"".Stringer+24
	rel 109+0 t=11 +0
	rel 141+4 t=8 runtime.convTstring+0
	rel 183+4 t=16 type.string+0
	rel 256+4 t=8 fmt.Println+0
	rel 277+4 t=8 runtime.morestack_noctxt+0
"".(*Binary).Get STEXT dupok size=134 args=0x10 locals=0x28
	0x0000 00000 (<autogenerated>:1)	TEXT	"".(*Binary).Get(SB), DUPOK|WRAPPER|ABIInternal, $40-16
	0x0000 00000 (<autogenerated>:1)	MOVQ	(TLS), CX
	0x0009 00009 (<autogenerated>:1)	CMPQ	SP, 16(CX)
	0x000d 00013 (<autogenerated>:1)	PCDATA	$0, $-2
	0x000d 00013 (<autogenerated>:1)	JLS	112
	0x000f 00015 (<autogenerated>:1)	PCDATA	$0, $-1
	0x000f 00015 (<autogenerated>:1)	SUBQ	$40, SP
	0x0013 00019 (<autogenerated>:1)	MOVQ	BP, 32(SP)
	0x0018 00024 (<autogenerated>:1)	LEAQ	32(SP), BP
	0x001d 00029 (<autogenerated>:1)	MOVQ	32(CX), BX
	0x0021 00033 (<autogenerated>:1)	TESTQ	BX, BX
	0x0024 00036 (<autogenerated>:1)	JNE	119
	0x0026 00038 (<autogenerated>:1)	NOP
	0x0026 00038 (<autogenerated>:1)	PCDATA	$0, $-2
	0x0026 00038 (<autogenerated>:1)	PCDATA	$1, $-2
	0x0026 00038 (<autogenerated>:1)	FUNCDATA	$0, gclocals·1a65e721a2ccc325b382662e7ffee780(SB)
	0x0026 00038 (<autogenerated>:1)	FUNCDATA	$1, gclocals·69c1753bd5f81501d95132d08af04464(SB)
	0x0026 00038 (<autogenerated>:1)	FUNCDATA	$2, gclocals·9fb7f0986f647f17cb53dda1484e0f7a(SB)
	0x0026 00038 (<autogenerated>:1)	PCDATA	$0, $0
	0x0026 00038 (<autogenerated>:1)	PCDATA	$1, $0
	0x0026 00038 (<autogenerated>:1)	MOVQ	$0, "".~r0+56(SP)
	0x002f 00047 (<autogenerated>:1)	CMPQ	""..this+48(SP), $0
	0x0035 00053 (<autogenerated>:1)	JNE	57
	0x0037 00055 (<autogenerated>:1)	JMP	106
	0x0039 00057 (<autogenerated>:1)	PCDATA	$0, $1
	0x0039 00057 (<autogenerated>:1)	PCDATA	$1, $1
	0x0039 00057 (<autogenerated>:1)	MOVQ	""..this+48(SP), AX
	0x003e 00062 (<autogenerated>:1)	TESTB	AL, (AX)
	0x0040 00064 (<autogenerated>:1)	PCDATA	$0, $0
	0x0040 00064 (<autogenerated>:1)	MOVQ	(AX), AX
	0x0043 00067 (<autogenerated>:1)	MOVQ	AX, ""..autotmp_3+16(SP)
	0x0048 00072 (<autogenerated>:1)	MOVQ	AX, (SP)
	0x004c 00076 (<autogenerated>:1)	CALL	"".Binary.Get(SB)
	0x0051 00081 (<autogenerated>:1)	MOVQ	8(SP), AX
	0x0056 00086 (<autogenerated>:1)	MOVQ	AX, ""..autotmp_2+24(SP)
	0x005b 00091 (<autogenerated>:1)	MOVQ	AX, "".~r0+56(SP)
	0x0060 00096 (<autogenerated>:1)	MOVQ	32(SP), BP
	0x0065 00101 (<autogenerated>:1)	ADDQ	$40, SP
	0x0069 00105 (<autogenerated>:1)	RET
	0x006a 00106 (<autogenerated>:1)	CALL	runtime.panicwrap(SB)
	0x006f 00111 (<autogenerated>:1)	XCHGL	AX, AX
	0x0070 00112 (<autogenerated>:1)	NOP
	0x0070 00112 (<autogenerated>:1)	PCDATA	$1, $-1
	0x0070 00112 (<autogenerated>:1)	PCDATA	$0, $-2
	0x0070 00112 (<autogenerated>:1)	CALL	runtime.morestack_noctxt(SB)
	0x0075 00117 (<autogenerated>:1)	PCDATA	$0, $-1
	0x0075 00117 (<autogenerated>:1)	JMP	0
	0x0077 00119 (<autogenerated>:1)	LEAQ	48(SP), DI
	0x007c 00124 (<autogenerated>:1)	CMPQ	(BX), DI
	0x007f 00127 (<autogenerated>:1)	JNE	38
	0x0081 00129 (<autogenerated>:1)	MOVQ	SP, (BX)
	0x0084 00132 (<autogenerated>:1)	JMP	38
	0x0000 65 48 8b 0c 25 00 00 00 00 48 3b 61 10 76 61 48  eH..%....H;a.vaH
	0x0010 83 ec 28 48 89 6c 24 20 48 8d 6c 24 20 48 8b 59  ..(H.l$ H.l$ H.Y
	0x0020 20 48 85 db 75 51 48 c7 44 24 38 00 00 00 00 48   H..uQH.D$8....H
	0x0030 83 7c 24 30 00 75 02 eb 31 48 8b 44 24 30 84 00  .|$0.u..1H.D$0..
	0x0040 48 8b 00 48 89 44 24 10 48 89 04 24 e8 00 00 00  H..H.D$.H..$....
	0x0050 00 48 8b 44 24 08 48 89 44 24 18 48 89 44 24 38  .H.D$.H.D$.H.D$8
	0x0060 48 8b 6c 24 20 48 83 c4 28 c3 e8 00 00 00 00 90  H.l$ H..(.......
	0x0070 e8 00 00 00 00 eb 89 48 8d 7c 24 30 48 39 3b 75  .......H.|$0H9;u
	0x0080 a5 48 89 23 eb a0                                .H.#..
	rel 5+4 t=17 TLS+0
	rel 77+4 t=8 "".Binary.Get+0
	rel 107+4 t=8 runtime.panicwrap+0
	rel 113+4 t=8 runtime.morestack_noctxt+0
"".(*Binary).String STEXT dupok size=151 args=0x18 locals=0x38
	0x0000 00000 (<autogenerated>:1)	TEXT	"".(*Binary).String(SB), DUPOK|WRAPPER|ABIInternal, $56-24
	0x0000 00000 (<autogenerated>:1)	MOVQ	(TLS), CX
	0x0009 00009 (<autogenerated>:1)	CMPQ	SP, 16(CX)
	0x000d 00013 (<autogenerated>:1)	PCDATA	$0, $-2
	0x000d 00013 (<autogenerated>:1)	JLS	126
	0x000f 00015 (<autogenerated>:1)	PCDATA	$0, $-1
	0x000f 00015 (<autogenerated>:1)	SUBQ	$56, SP
	0x0013 00019 (<autogenerated>:1)	MOVQ	BP, 48(SP)
	0x0018 00024 (<autogenerated>:1)	LEAQ	48(SP), BP
	0x001d 00029 (<autogenerated>:1)	MOVQ	32(CX), BX
	0x0021 00033 (<autogenerated>:1)	TESTQ	BX, BX
	0x0024 00036 (<autogenerated>:1)	JNE	136
	0x0026 00038 (<autogenerated>:1)	NOP
	0x0026 00038 (<autogenerated>:1)	PCDATA	$0, $-2
	0x0026 00038 (<autogenerated>:1)	PCDATA	$1, $-2
	0x0026 00038 (<autogenerated>:1)	FUNCDATA	$0, gclocals·62420d0a7277934df9079483e4a3e39b(SB)
	0x0026 00038 (<autogenerated>:1)	FUNCDATA	$1, gclocals·d8b28f51bb91e05d264803f0f600a200(SB)
	0x0026 00038 (<autogenerated>:1)	FUNCDATA	$2, gclocals·1cf923758aae2e428391d1783fe59973(SB)
	0x0026 00038 (<autogenerated>:1)	PCDATA	$0, $0
	0x0026 00038 (<autogenerated>:1)	PCDATA	$1, $0
	0x0026 00038 (<autogenerated>:1)	XORPS	X0, X0
	0x0029 00041 (<autogenerated>:1)	MOVUPS	X0, "".~r0+72(SP)
	0x002e 00046 (<autogenerated>:1)	CMPQ	""..this+64(SP), $0
	0x0034 00052 (<autogenerated>:1)	JNE	56
	0x0036 00054 (<autogenerated>:1)	JMP	120
	0x0038 00056 (<autogenerated>:1)	PCDATA	$0, $1
	0x0038 00056 (<autogenerated>:1)	PCDATA	$1, $1
	0x0038 00056 (<autogenerated>:1)	MOVQ	""..this+64(SP), AX
	0x003d 00061 (<autogenerated>:1)	TESTB	AL, (AX)
	0x003f 00063 (<autogenerated>:1)	PCDATA	$0, $0
	0x003f 00063 (<autogenerated>:1)	MOVQ	(AX), AX
	0x0042 00066 (<autogenerated>:1)	MOVQ	AX, ""..autotmp_3+24(SP)
	0x0047 00071 (<autogenerated>:1)	MOVQ	AX, (SP)
	0x004b 00075 (<autogenerated>:1)	CALL	"".Binary.String(SB)
	0x0050 00080 (<autogenerated>:1)	MOVQ	16(SP), AX
	0x0055 00085 (<autogenerated>:1)	PCDATA	$0, $2
	0x0055 00085 (<autogenerated>:1)	MOVQ	8(SP), CX
	0x005a 00090 (<autogenerated>:1)	MOVQ	CX, ""..autotmp_2+32(SP)
	0x005f 00095 (<autogenerated>:1)	MOVQ	AX, ""..autotmp_2+40(SP)
	0x0064 00100 (<autogenerated>:1)	PCDATA	$0, $0
	0x0064 00100 (<autogenerated>:1)	PCDATA	$1, $2
	0x0064 00100 (<autogenerated>:1)	MOVQ	CX, "".~r0+72(SP)
	0x0069 00105 (<autogenerated>:1)	MOVQ	AX, "".~r0+80(SP)
	0x006e 00110 (<autogenerated>:1)	MOVQ	48(SP), BP
	0x0073 00115 (<autogenerated>:1)	ADDQ	$56, SP
	0x0077 00119 (<autogenerated>:1)	RET
	0x0078 00120 (<autogenerated>:1)	PCDATA	$1, $1
	0x0078 00120 (<autogenerated>:1)	CALL	runtime.panicwrap(SB)
	0x007d 00125 (<autogenerated>:1)	XCHGL	AX, AX
	0x007e 00126 (<autogenerated>:1)	NOP
	0x007e 00126 (<autogenerated>:1)	PCDATA	$1, $-1
	0x007e 00126 (<autogenerated>:1)	PCDATA	$0, $-2
	0x007e 00126 (<autogenerated>:1)	CALL	runtime.morestack_noctxt(SB)
	0x0083 00131 (<autogenerated>:1)	PCDATA	$0, $-1
	0x0083 00131 (<autogenerated>:1)	JMP	0
	0x0088 00136 (<autogenerated>:1)	LEAQ	64(SP), DI
	0x008d 00141 (<autogenerated>:1)	CMPQ	(BX), DI
	0x0090 00144 (<autogenerated>:1)	JNE	38
	0x0092 00146 (<autogenerated>:1)	MOVQ	SP, (BX)
	0x0095 00149 (<autogenerated>:1)	JMP	38
	0x0000 65 48 8b 0c 25 00 00 00 00 48 3b 61 10 76 6f 48  eH..%....H;a.voH
	0x0010 83 ec 38 48 89 6c 24 30 48 8d 6c 24 30 48 8b 59  ..8H.l$0H.l$0H.Y
	0x0020 20 48 85 db 75 62 0f 57 c0 0f 11 44 24 48 48 83   H..ub.W...D$HH.
	0x0030 7c 24 40 00 75 02 eb 40 48 8b 44 24 40 84 00 48  |$@.u..@H.D$@..H
	0x0040 8b 00 48 89 44 24 18 48 89 04 24 e8 00 00 00 00  ..H.D$.H..$.....
	0x0050 48 8b 44 24 10 48 8b 4c 24 08 48 89 4c 24 20 48  H.D$.H.L$.H.L$ H
	0x0060 89 44 24 28 48 89 4c 24 48 48 89 44 24 50 48 8b  .D$(H.L$HH.D$PH.
	0x0070 6c 24 30 48 83 c4 38 c3 e8 00 00 00 00 90 e8 00  l$0H..8.........
	0x0080 00 00 00 e9 78 ff ff ff 48 8d 7c 24 40 48 39 3b  ....x...H.|$@H9;
	0x0090 75 94 48 89 23 eb 8f                             u.H.#..
	rel 5+4 t=17 TLS+0
	rel 76+4 t=8 "".Binary.String+0
	rel 121+4 t=8 runtime.panicwrap+0
	rel 127+4 t=8 runtime.morestack_noctxt+0
"".Stringer.String STEXT dupok size=130 args=0x20 locals=0x30
	0x0000 00000 (<autogenerated>:1)	TEXT	"".Stringer.String(SB), DUPOK|WRAPPER|ABIInternal, $48-32
	0x0000 00000 (<autogenerated>:1)	MOVQ	(TLS), CX
	0x0009 00009 (<autogenerated>:1)	CMPQ	SP, 16(CX)
	0x000d 00013 (<autogenerated>:1)	PCDATA	$0, $-2
	0x000d 00013 (<autogenerated>:1)	JLS	108
	0x000f 00015 (<autogenerated>:1)	PCDATA	$0, $-1
	0x000f 00015 (<autogenerated>:1)	SUBQ	$48, SP
	0x0013 00019 (<autogenerated>:1)	MOVQ	BP, 40(SP)
	0x0018 00024 (<autogenerated>:1)	LEAQ	40(SP), BP
	0x001d 00029 (<autogenerated>:1)	MOVQ	32(CX), BX
	0x0021 00033 (<autogenerated>:1)	TESTQ	BX, BX
	0x0024 00036 (<autogenerated>:1)	JNE	115
	0x0026 00038 (<autogenerated>:1)	NOP
	0x0026 00038 (<autogenerated>:1)	PCDATA	$0, $-2
	0x0026 00038 (<autogenerated>:1)	PCDATA	$1, $-2
	0x0026 00038 (<autogenerated>:1)	FUNCDATA	$0, gclocals·e0f6dd6ffe13df6eefebd78fb394216d(SB)
	0x0026 00038 (<autogenerated>:1)	FUNCDATA	$1, gclocals·d8b28f51bb91e05d264803f0f600a200(SB)
	0x0026 00038 (<autogenerated>:1)	FUNCDATA	$2, gclocals·568470801006e5c0dc3947ea998fe279(SB)
	0x0026 00038 (<autogenerated>:1)	PCDATA	$0, $0
	0x0026 00038 (<autogenerated>:1)	PCDATA	$1, $0
	0x0026 00038 (<autogenerated>:1)	XORPS	X0, X0
	0x0029 00041 (<autogenerated>:1)	MOVUPS	X0, "".~r0+72(SP)
	0x002e 00046 (<autogenerated>:1)	MOVQ	""..this+56(SP), AX
	0x0033 00051 (<autogenerated>:1)	TESTB	AL, (AX)
	0x0035 00053 (<autogenerated>:1)	MOVQ	24(AX), AX
	0x0039 00057 (<autogenerated>:1)	PCDATA	$0, $1
	0x0039 00057 (<autogenerated>:1)	PCDATA	$1, $1
	0x0039 00057 (<autogenerated>:1)	MOVQ	""..this+64(SP), CX
	0x003e 00062 (<autogenerated>:1)	PCDATA	$0, $0
	0x003e 00062 (<autogenerated>:1)	MOVQ	CX, (SP)
	0x0042 00066 (<autogenerated>:1)	CALL	AX
	0x0044 00068 (<autogenerated>:1)	MOVQ	16(SP), AX
	0x0049 00073 (<autogenerated>:1)	PCDATA	$0, $1
	0x0049 00073 (<autogenerated>:1)	MOVQ	8(SP), CX
	0x004e 00078 (<autogenerated>:1)	MOVQ	CX, ""..autotmp_2+24(SP)
	0x0053 00083 (<autogenerated>:1)	MOVQ	AX, ""..autotmp_2+32(SP)
	0x0058 00088 (<autogenerated>:1)	PCDATA	$0, $0
	0x0058 00088 (<autogenerated>:1)	PCDATA	$1, $2
	0x0058 00088 (<autogenerated>:1)	MOVQ	CX, "".~r0+72(SP)
	0x005d 00093 (<autogenerated>:1)	MOVQ	AX, "".~r0+80(SP)
	0x0062 00098 (<autogenerated>:1)	MOVQ	40(SP), BP
	0x0067 00103 (<autogenerated>:1)	ADDQ	$48, SP
	0x006b 00107 (<autogenerated>:1)	RET
	0x006c 00108 (<autogenerated>:1)	NOP
	0x006c 00108 (<autogenerated>:1)	PCDATA	$1, $-1
	0x006c 00108 (<autogenerated>:1)	PCDATA	$0, $-2
	0x006c 00108 (<autogenerated>:1)	CALL	runtime.morestack_noctxt(SB)
	0x0071 00113 (<autogenerated>:1)	PCDATA	$0, $-1
	0x0071 00113 (<autogenerated>:1)	JMP	0
	0x0073 00115 (<autogenerated>:1)	LEAQ	56(SP), DI
	0x0078 00120 (<autogenerated>:1)	CMPQ	(BX), DI
	0x007b 00123 (<autogenerated>:1)	JNE	38
	0x007d 00125 (<autogenerated>:1)	MOVQ	SP, (BX)
	0x0080 00128 (<autogenerated>:1)	JMP	38
	0x0000 65 48 8b 0c 25 00 00 00 00 48 3b 61 10 76 5d 48  eH..%....H;a.v]H
	0x0010 83 ec 30 48 89 6c 24 28 48 8d 6c 24 28 48 8b 59  ..0H.l$(H.l$(H.Y
	0x0020 20 48 85 db 75 4d 0f 57 c0 0f 11 44 24 48 48 8b   H..uM.W...D$HH.
	0x0030 44 24 38 84 00 48 8b 40 18 48 8b 4c 24 40 48 89  D$8..H.@.H.L$@H.
	0x0040 0c 24 ff d0 48 8b 44 24 10 48 8b 4c 24 08 48 89  .$..H.D$.H.L$.H.
	0x0050 4c 24 18 48 89 44 24 20 48 89 4c 24 48 48 89 44  L$.H.D$ H.L$HH.D
	0x0060 24 50 48 8b 6c 24 28 48 83 c4 30 c3 e8 00 00 00  $PH.l$(H..0.....
	0x0070 00 eb 8d 48 8d 7c 24 38 48 39 3b 75 a9 48 89 23  ...H.|$8H9;u.H.#
	0x0080 eb a4                                            ..
	rel 5+4 t=17 TLS+0
	rel 66+0 t=11 +0
	rel 109+4 t=8 runtime.morestack_noctxt+0
go.cuinfo.packagename. SDWARFINFO dupok size=0
	0x0000 6d 61 69 6e                                      main
go.loc."".Binary.String SDWARFLOC size=0
go.info."".Binary.String SDWARFINFO size=67
	0x0000 03 22 22 2e 42 69 6e 61 72 79 2e 53 74 72 69 6e  ."".Binary.Strin
	0x0010 67 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  g...............
	0x0020 00 00 01 9c 00 00 00 00 01 0f 69 00 00 0e 00 00  ..........i.....
	0x0030 00 00 01 9c 0f 7e 72 30 00 01 0e 00 00 00 00 02  .....~r0........
	0x0040 91 08 00                                         ...
	rel 0+0 t=24 type."".Binary+0
	rel 0+0 t=24 type.string+0
	rel 0+0 t=24 type.uint64+0
	rel 18+8 t=1 "".Binary.String+0
	rel 26+8 t=1 "".Binary.String+126
	rel 36+4 t=30 gofile../Users/mac/go/src/github.com/kocoler/GoExercises/tests/interfaceT.go+0
	rel 46+4 t=29 go.info."".Binary+0
	rel 59+4 t=29 go.info.string+0
go.range."".Binary.String SDWARFRANGE size=0
go.debuglines."".Binary.String SDWARFMISC size=22
	0x0000 04 02 03 08 14 0a a5 06 b9 06 42 06 41 06 02 35  ..........B.A..5
	0x0010 fe 04 01 03 73 01                                ....s.
go.loc."".Binary.Get SDWARFLOC size=0
go.info."".Binary.Get SDWARFINFO size=64
	0x0000 03 22 22 2e 42 69 6e 61 72 79 2e 47 65 74 00 00  ."".Binary.Get..
	0x0010 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 01  ................
	0x0020 9c 00 00 00 00 01 0f 69 00 00 12 00 00 00 00 01  .......i........
	0x0030 9c 0f 7e 72 30 00 01 12 00 00 00 00 02 91 08 00  ..~r0...........
	rel 0+0 t=24 type."".Binary+0
	rel 0+0 t=24 type.uint64+0
	rel 15+8 t=1 "".Binary.Get+0
	rel 23+8 t=1 "".Binary.Get+20
	rel 33+4 t=30 gofile../Users/mac/go/src/github.com/kocoler/GoExercises/tests/interfaceT.go+0
	rel 43+4 t=29 go.info."".Binary+0
	rel 56+4 t=29 go.info.uint64+0
go.range."".Binary.Get SDWARFRANGE size=0
go.debuglines."".Binary.Get SDWARFMISC size=14
	0x0000 04 02 03 0c 14 6a 06 41 04 01 03 6e 06 01        .....j.A...n..
go.loc."".main SDWARFLOC size=0
go.info."".main SDWARFINFO size=57
	0x0000 03 22 22 2e 6d 61 69 6e 00 00 00 00 00 00 00 00  ."".main........
	0x0010 00 00 00 00 00 00 00 00 00 01 9c 00 00 00 00 01  ................
	0x0020 0a 62 00 17 00 00 00 00 03 91 88 7f 0a 73 00 18  .b...........s..
	0x0030 00 00 00 00 03 91 a8 7f 00                       .........
	rel 0+0 t=24 type."".Binary+0
	rel 0+0 t=24 type."".Stringer+0
	rel 0+0 t=24 type.*[1]interface {}+0
	rel 0+0 t=24 type.[1]interface {}+0
	rel 0+0 t=24 type.[]interface {}+0
	rel 0+0 t=24 type.string+0
	rel 0+0 t=24 type.unsafe.Pointer+0
	rel 9+8 t=1 "".main+0
	rel 17+8 t=1 "".main+286
	rel 27+4 t=30 gofile../Users/mac/go/src/github.com/kocoler/GoExercises/tests/interfaceT.go+0
	rel 36+4 t=29 go.info."".Binary+0
	rel 48+4 t=29 go.info."".Stringer+0
go.range."".main SDWARFRANGE size=0
go.debuglines."".main SDWARFMISC size=27
	0x0000 04 02 03 10 14 0a ff f6 6a 06 5f 06 08 60 06 23  ........j._..`.#
	0x0010 06 02 8b 01 f6 ab 04 01 03 6b 01                 .........k.
runtime.nilinterequal·f SRODATA dupok size=8
	0x0000 00 00 00 00 00 00 00 00                          ........
	rel 0+8 t=1 runtime.nilinterequal+0
runtime.memequal64·f SRODATA dupok size=8
	0x0000 00 00 00 00 00 00 00 00                          ........
	rel 0+8 t=1 runtime.memequal64+0
runtime.gcbits.01 SRODATA dupok size=1
	0x0000 01                                               .
type..namedata.*interface {}- SRODATA dupok size=16
	0x0000 00 00 0d 2a 69 6e 74 65 72 66 61 63 65 20 7b 7d  ...*interface {}
type.*interface {} SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 4f 0f 96 9d 08 08 08 36 00 00 00 00 00 00 00 00  O......6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64·f+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*interface {}-+0
	rel 48+8 t=1 type.interface {}+0
runtime.gcbits.02 SRODATA dupok size=1
	0x0000 02                                               .
type.interface {} SRODATA dupok size=80
	0x0000 10 00 00 00 00 00 00 00 10 00 00 00 00 00 00 00  ................
	0x0010 e7 57 a0 18 02 08 08 14 00 00 00 00 00 00 00 00  .W..............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	rel 24+8 t=1 runtime.nilinterequal·f+0
	rel 32+8 t=1 runtime.gcbits.02+0
	rel 40+4 t=5 type..namedata.*interface {}-+0
	rel 44+4 t=6 type.*interface {}+0
	rel 56+8 t=1 type.interface {}+80
type..namedata.*[]interface {}- SRODATA dupok size=18
	0x0000 00 00 0f 2a 5b 5d 69 6e 74 65 72 66 61 63 65 20  ...*[]interface 
	0x0010 7b 7d                                            {}
type.*[]interface {} SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 f3 04 9a e7 08 08 08 36 00 00 00 00 00 00 00 00  .......6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64·f+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[]interface {}-+0
	rel 48+8 t=1 type.[]interface {}+0
type.[]interface {} SRODATA dupok size=56
	0x0000 18 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 70 93 ea 2f 02 08 08 17 00 00 00 00 00 00 00 00  p../............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[]interface {}-+0
	rel 44+4 t=6 type.*[]interface {}+0
	rel 48+8 t=1 type.interface {}+0
type..namedata.*[1]interface {}- SRODATA dupok size=19
	0x0000 00 00 10 2a 5b 31 5d 69 6e 74 65 72 66 61 63 65  ...*[1]interface
	0x0010 20 7b 7d                                          {}
type.*[1]interface {} SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 bf 03 a8 35 08 08 08 36 00 00 00 00 00 00 00 00  ...5...6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64·f+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[1]interface {}-+0
	rel 48+8 t=1 type.[1]interface {}+0
type.[1]interface {} SRODATA dupok size=72
	0x0000 10 00 00 00 00 00 00 00 10 00 00 00 00 00 00 00  ................
	0x0010 50 91 5b fa 02 08 08 11 00 00 00 00 00 00 00 00  P.[.............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 01 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.nilinterequal·f+0
	rel 32+8 t=1 runtime.gcbits.02+0
	rel 40+4 t=5 type..namedata.*[1]interface {}-+0
	rel 44+4 t=6 type.*[1]interface {}+0
	rel 48+8 t=1 type.interface {}+0
	rel 56+8 t=1 type.[]interface {}+0
""..inittask SNOPTRDATA size=40
	0x0000 00 00 00 00 00 00 00 00 02 00 00 00 00 00 00 00  ................
	0x0010 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0020 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 fmt..inittask+0
	rel 32+8 t=1 strconv..inittask+0
go.loc."".(*Binary).Get SDWARFLOC dupok size=0
go.info."".(*Binary).Get SDWARFINFO dupok size=71
	0x0000 03 22 22 2e 28 2a 42 69 6e 61 72 79 29 2e 47 65  ."".(*Binary).Ge
	0x0010 74 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  t...............
	0x0020 00 00 01 9c 00 00 00 00 01 0f 2e 74 68 69 73 00  ...........this.
	0x0030 00 01 00 00 00 00 01 9c 0f 7e 72 30 00 01 12 00  .........~r0....
	0x0040 00 00 00 02 91 08 00                             .......
	rel 0+0 t=24 type."".Binary+0
	rel 0+0 t=24 type.*"".Binary+0
	rel 0+0 t=24 type.uint64+0
	rel 18+8 t=1 "".(*Binary).Get+0
	rel 26+8 t=1 "".(*Binary).Get+134
	rel 36+4 t=30 gofile..<autogenerated>+0
	rel 50+4 t=29 go.info.*"".Binary+0
	rel 63+4 t=29 go.info.uint64+0
go.range."".(*Binary).Get SDWARFRANGE dupok size=0
go.debuglines."".(*Binary).Get SDWARFMISC dupok size=21
	0x0000 04 01 0f 0a a5 06 08 5f 06 73 06 41 06 02 19 ff  ......._.s.A....
	0x0010 04 01 03 00 01                                   .....
go.loc."".(*Binary).String SDWARFLOC dupok size=0
go.info."".(*Binary).String SDWARFINFO dupok size=74
	0x0000 03 22 22 2e 28 2a 42 69 6e 61 72 79 29 2e 53 74  ."".(*Binary).St
	0x0010 72 69 6e 67 00 00 00 00 00 00 00 00 00 00 00 00  ring............
	0x0020 00 00 00 00 00 01 9c 00 00 00 00 01 0f 2e 74 68  ..............th
	0x0030 69 73 00 00 01 00 00 00 00 01 9c 0f 7e 72 30 00  is..........~r0.
	0x0040 01 0e 00 00 00 00 02 91 08 00                    ..........
	rel 0+0 t=24 type."".Binary+0
	rel 0+0 t=24 type.*"".Binary+0
	rel 0+0 t=24 type.string+0
	rel 21+8 t=1 "".(*Binary).String+0
	rel 29+8 t=1 "".(*Binary).String+151
	rel 39+4 t=30 gofile..<autogenerated>+0
	rel 53+4 t=29 go.info.*"".Binary+0
	rel 66+4 t=29 go.info.string+0
go.range."".(*Binary).String SDWARFRANGE dupok size=0
go.debuglines."".(*Binary).String SDWARFMISC dupok size=21
	0x0000 04 01 0f 0a a5 06 08 23 06 a5 06 41 06 02 28 ff  .......#...A..(.
	0x0010 04 01 03 00 01                                   .....
type..namedata.*main.Binary. SRODATA dupok size=15
	0x0000 01 00 0c 2a 6d 61 69 6e 2e 42 69 6e 61 72 79     ...*main.Binary
type..namedata.*func(*main.Binary) uint64- SRODATA dupok size=29
	0x0000 00 00 1a 2a 66 75 6e 63 28 2a 6d 61 69 6e 2e 42  ...*func(*main.B
	0x0010 69 6e 61 72 79 29 20 75 69 6e 74 36 34           inary) uint64
type.*func(*"".Binary) uint64 SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 af 9a 32 0f 08 08 08 36 00 00 00 00 00 00 00 00  ..2....6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64·f+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*func(*main.Binary) uint64-+0
	rel 48+8 t=1 type.func(*"".Binary) uint64+0
type.func(*"".Binary) uint64 SRODATA dupok size=72
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 a4 61 e5 41 02 08 08 33 00 00 00 00 00 00 00 00  .a.A...3........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 01 00 01 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 00 00 00 00 00 00 00 00                          ........
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*func(*main.Binary) uint64-+0
	rel 44+4 t=6 type.*func(*"".Binary) uint64+0
	rel 56+8 t=1 type.*"".Binary+0
	rel 64+8 t=1 type.uint64+0
type..namedata.*func(*main.Binary) string- SRODATA dupok size=29
	0x0000 00 00 1a 2a 66 75 6e 63 28 2a 6d 61 69 6e 2e 42  ...*func(*main.B
	0x0010 69 6e 61 72 79 29 20 73 74 72 69 6e 67           inary) string
type.*func(*"".Binary) string SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 69 63 71 98 08 08 08 36 00 00 00 00 00 00 00 00  icq....6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64·f+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*func(*main.Binary) string-+0
	rel 48+8 t=1 type.func(*"".Binary) string+0
type.func(*"".Binary) string SRODATA dupok size=72
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 43 b6 3f 72 02 08 08 33 00 00 00 00 00 00 00 00  C.?r...3........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 01 00 01 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 00 00 00 00 00 00 00 00                          ........
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*func(*main.Binary) string-+0
	rel 44+4 t=6 type.*func(*"".Binary) string+0
	rel 56+8 t=1 type.*"".Binary+0
	rel 64+8 t=1 type.string+0
type..namedata.Get. SRODATA dupok size=6
	0x0000 01 00 03 47 65 74                                ...Get
type..namedata.*func() uint64- SRODATA dupok size=17
	0x0000 00 00 0e 2a 66 75 6e 63 28 29 20 75 69 6e 74 36  ...*func() uint6
	0x0010 34                                               4
type.*func() uint64 SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 b0 cf 66 34 08 08 08 36 00 00 00 00 00 00 00 00  ..f4...6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64·f+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*func() uint64-+0
	rel 48+8 t=1 type.func() uint64+0
type.func() uint64 SRODATA dupok size=64
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 0c d3 e7 24 02 08 08 33 00 00 00 00 00 00 00 00  ...$...3........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 01 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*func() uint64-+0
	rel 44+4 t=6 type.*func() uint64+0
	rel 56+8 t=1 type.uint64+0
type..namedata.String. SRODATA dupok size=9
	0x0000 01 00 06 53 74 72 69 6e 67                       ...String
type..namedata.*func() string- SRODATA dupok size=17
	0x0000 00 00 0e 2a 66 75 6e 63 28 29 20 73 74 72 69 6e  ...*func() strin
	0x0010 67                                               g
type.*func() string SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 bc f4 77 69 08 08 08 36 00 00 00 00 00 00 00 00  ..wi...6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64·f+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*func() string-+0
	rel 48+8 t=1 type.func() string+0
type.func() string SRODATA dupok size=64
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 a2 6d cb 1e 02 08 08 33 00 00 00 00 00 00 00 00  .m.....3........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 01 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*func() string-+0
	rel 44+4 t=6 type.*func() string+0
	rel 56+8 t=1 type.string+0
type.*"".Binary SRODATA size=104
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 f1 3c ca 37 09 08 08 36 00 00 00 00 00 00 00 00  .<.7...6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 02 00 02 00  ................
	0x0040 10 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0050 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0060 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64·f+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*main.Binary.+0
	rel 48+8 t=1 type."".Binary+0
	rel 56+4 t=5 type..importpath."".+0
	rel 72+4 t=5 type..namedata.Get.+0
	rel 76+4 t=25 type.func() uint64+0
	rel 80+4 t=25 "".(*Binary).Get+0
	rel 84+4 t=25 "".(*Binary).Get+0
	rel 88+4 t=5 type..namedata.String.+0
	rel 92+4 t=25 type.func() string+0
	rel 96+4 t=25 "".(*Binary).String+0
	rel 100+4 t=25 "".(*Binary).String+0
runtime.gcbits. SRODATA dupok size=0
type..namedata.*func(main.Binary) uint64- SRODATA dupok size=28
	0x0000 00 00 19 2a 66 75 6e 63 28 6d 61 69 6e 2e 42 69  ...*func(main.Bi
	0x0010 6e 61 72 79 29 20 75 69 6e 74 36 34              nary) uint64
type.*func("".Binary) uint64 SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 9a 67 b6 99 08 08 08 36 00 00 00 00 00 00 00 00  .g.....6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64·f+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*func(main.Binary) uint64-+0
	rel 48+8 t=1 type.func("".Binary) uint64+0
type.func("".Binary) uint64 SRODATA dupok size=72
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 7e d0 f0 9c 02 08 08 33 00 00 00 00 00 00 00 00  ~......3........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 01 00 01 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 00 00 00 00 00 00 00 00                          ........
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*func(main.Binary) uint64-+0
	rel 44+4 t=6 type.*func("".Binary) uint64+0
	rel 56+8 t=1 type."".Binary+0
	rel 64+8 t=1 type.uint64+0
type..namedata.*func(main.Binary) string- SRODATA dupok size=28
	0x0000 00 00 19 2a 66 75 6e 63 28 6d 61 69 6e 2e 42 69  ...*func(main.Bi
	0x0010 6e 61 72 79 29 20 73 74 72 69 6e 67              nary) string
type.*func("".Binary) string SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 4a d9 60 10 08 08 08 36 00 00 00 00 00 00 00 00  J.`....6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64·f+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*func(main.Binary) string-+0
	rel 48+8 t=1 type.func("".Binary) string+0
type.func("".Binary) string SRODATA dupok size=72
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 17 36 d7 d8 02 08 08 33 00 00 00 00 00 00 00 00  .6.....3........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 01 00 01 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 00 00 00 00 00 00 00 00                          ........
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*func(main.Binary) string-+0
	rel 44+4 t=6 type.*func("".Binary) string+0
	rel 56+8 t=1 type."".Binary+0
	rel 64+8 t=1 type.string+0
type."".Binary SRODATA size=96
	0x0000 08 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0010 3b 3d 72 44 0f 08 08 0b 00 00 00 00 00 00 00 00  ;=rD............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 02 00 02 00 10 00 00 00 00 00 00 00  ................
	0x0040 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0050 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	rel 24+8 t=1 runtime.memequal64·f+0
	rel 32+8 t=1 runtime.gcbits.+0
	rel 40+4 t=5 type..namedata.*main.Binary.+0
	rel 44+4 t=5 type.*"".Binary+0
	rel 48+4 t=5 type..importpath."".+0
	rel 64+4 t=5 type..namedata.Get.+0
	rel 68+4 t=25 type.func() uint64+0
	rel 72+4 t=25 "".(*Binary).Get+0
	rel 76+4 t=25 "".Binary.Get+0
	rel 80+4 t=5 type..namedata.String.+0
	rel 84+4 t=25 type.func() string+0
	rel 88+4 t=25 "".(*Binary).String+0
	rel 92+4 t=25 "".Binary.String+0
go.loc."".Stringer.String SDWARFLOC dupok size=0
go.info."".Stringer.String SDWARFINFO dupok size=73
	0x0000 03 22 22 2e 53 74 72 69 6e 67 65 72 2e 53 74 72  ."".Stringer.Str
	0x0010 69 6e 67 00 00 00 00 00 00 00 00 00 00 00 00 00  ing.............
	0x0020 00 00 00 00 01 9c 00 00 00 00 01 0f 2e 74 68 69  .............thi
	0x0030 73 00 00 01 00 00 00 00 01 9c 0f 7e 72 30 00 01  s..........~r0..
	0x0040 09 00 00 00 00 02 91 10 00                       .........
	rel 0+0 t=24 type."".Stringer+0
	rel 0+0 t=24 type.string+0
	rel 20+8 t=1 "".Stringer.String+0
	rel 28+8 t=1 "".Stringer.String+130
	rel 38+4 t=30 gofile..<autogenerated>+0
	rel 52+4 t=29 go.info."".Stringer+0
	rel 65+4 t=29 go.info.string+0
go.range."".Stringer.String SDWARFRANGE dupok size=0
go.debuglines."".Stringer.String SDWARFMISC dupok size=17
	0x0000 04 01 0f 0a a5 06 08 23 06 02 2b ff 04 01 03 00  .......#..+.....
	0x0010 01                                               .
runtime.interequal·f SRODATA dupok size=8
	0x0000 00 00 00 00 00 00 00 00                          ........
	rel 0+8 t=1 runtime.interequal+0
type..namedata.*main.Stringer. SRODATA dupok size=17
	0x0000 01 00 0e 2a 6d 61 69 6e 2e 53 74 72 69 6e 67 65  ...*main.Stringe
	0x0010 72                                               r
type.*"".Stringer SRODATA size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 26 d3 e3 8d 08 08 08 36 00 00 00 00 00 00 00 00  &......6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64·f+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*main.Stringer.+0
	rel 48+8 t=1 type."".Stringer+0
type."".Stringer SRODATA size=104
	0x0000 10 00 00 00 00 00 00 00 10 00 00 00 00 00 00 00  ................
	0x0010 57 94 f4 eb 07 08 08 14 00 00 00 00 00 00 00 00  W...............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 01 00 00 00 00 00 00 00 01 00 00 00 00 00 00 00  ................
	0x0050 00 00 00 00 00 00 00 00 18 00 00 00 00 00 00 00  ................
	0x0060 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.interequal·f+0
	rel 32+8 t=1 runtime.gcbits.02+0
	rel 40+4 t=5 type..namedata.*main.Stringer.+0
	rel 44+4 t=5 type.*"".Stringer+0
	rel 48+8 t=1 type..importpath."".+0
	rel 56+8 t=1 type."".Stringer+96
	rel 80+4 t=5 type..importpath."".+0
	rel 96+4 t=5 type..namedata.String.+0
	rel 100+4 t=5 type.func() string+0
go.itab."".Binary,"".Stringer SRODATA dupok size=32
	0x0000 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0010 3b 3d 72 44 00 00 00 00 00 00 00 00 00 00 00 00  ;=rD............
	rel 0+8 t=1 type."".Stringer+0
	rel 8+8 t=1 type."".Binary+0
	rel 24+8 t=1 "".(*Binary).String+0
go.itablink."".Binary,"".Stringer SRODATA dupok size=8
	0x0000 00 00 00 00 00 00 00 00                          ........
	rel 0+8 t=1 go.itab."".Binary,"".Stringer+0
type..importpath.fmt. SRODATA dupok size=6
	0x0000 00 00 03 66 6d 74                                ...fmt
type..importpath.strconv. SRODATA dupok size=10
	0x0000 00 00 07 73 74 72 63 6f 6e 76                    ...strconv
gclocals·568470801006e5c0dc3947ea998fe279 SRODATA dupok size=10
	0x0000 02 00 00 00 02 00 00 00 00 02                    ..........
gclocals·2589ca35330fc0fce83503f4569854a0 SRODATA dupok size=10
	0x0000 02 00 00 00 02 00 00 00 00 00                    ..........
gclocals·33cdeccccebe80329f1fdbee7f5874cb SRODATA dupok size=8
	0x0000 01 00 00 00 00 00 00 00                          ........
gclocals·7d2d5fca80364273fb07d5820a76fef4 SRODATA dupok size=8
	0x0000 03 00 00 00 00 00 00 00                          ........
gclocals·154e988fc21ce4c795057534d7421a72 SRODATA dupok size=14
	0x0000 03 00 00 00 0c 00 00 00 00 00 02 00 02 01        ..............
gclocals·f6aec3988379d2bd21c69c093370a150 SRODATA dupok size=12
	0x0000 04 00 00 00 03 00 00 00 00 01 03 07              ............
"".main.stkobj SRODATA size=24
	0x0000 01 00 00 00 00 00 00 00 d8 ff ff ff ff ff ff ff  ................
	0x0010 00 00 00 00 00 00 00 00                          ........
	rel 16+8 t=1 type.[1]interface {}+0
gclocals·1a65e721a2ccc325b382662e7ffee780 SRODATA dupok size=10
	0x0000 02 00 00 00 01 00 00 00 01 00                    ..........
gclocals·69c1753bd5f81501d95132d08af04464 SRODATA dupok size=8
	0x0000 02 00 00 00 00 00 00 00                          ........
gclocals·9fb7f0986f647f17cb53dda1484e0f7a SRODATA dupok size=10
	0x0000 02 00 00 00 01 00 00 00 00 01                    ..........
gclocals·62420d0a7277934df9079483e4a3e39b SRODATA dupok size=11
	0x0000 03 00 00 00 02 00 00 00 01 00 02                 ...........
gclocals·d8b28f51bb91e05d264803f0f600a200 SRODATA dupok size=11
	0x0000 03 00 00 00 02 00 00 00 00 00 00                 ...........
gclocals·1cf923758aae2e428391d1783fe59973 SRODATA dupok size=11
	0x0000 03 00 00 00 02 00 00 00 00 01 02                 ...........
gclocals·e0f6dd6ffe13df6eefebd78fb394216d SRODATA dupok size=11
	0x0000 03 00 00 00 03 00 00 00 02 00 04                 ...........