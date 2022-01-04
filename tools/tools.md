# Tools

> 提供常用的命令行工具。

## 代码统计工具 CLOC

CLOC(Count Lines of Code)，是一个可以统计多种编程语言中空行、评论行和物理行的工具。这个工具还是蛮实用的，可以帮我们快速了解一个项目中代码的信息。

### windows 10

win10下可以去github上下载其最新版，截止本文时，最新版为1.8.0。下载[链接](https://github.com/AlDanial/cloc/releases)。

### linux mac下安装

在linux下安装就简单的多了，使用你的发行版的包管理器下载安装即可。

```sh
# mac安装Homebrew后
brew install cloc
# ubuntu debian deepin mint等
sudo apt install cloc
# arch manjaro等
sudo pacman -S cloc 
# 或者
sudo yaourt -S cloc
# redhat centOS 
sudo yun install cloc
# Fedora 
sudo dnf install cloc
# ipv6可用时
sudo dnf -6 install cloc
```



## Go 命令行工具

### 将 Go 语言的源代码编译成汇编语言

将 Go 语言的源代码编译成汇编语言，然后通过汇编语言分析程序具体的执行过程。

汇编代码只是 Go 语言编译的结果，作为使用 Go 语言的开发者，已经能够通过上述结果分析程序的性能瓶颈。

```sh
$go1.17 build -gcflags -S main.go 
# command-line-arguments
"".main STEXT size=103 args=0x0 locals=0x40 funcid=0x0
        0x0000 00000 (/Users/rmliu/workspace/golang/src/github.com/Kate-liu/GoBeginner/helloworld/main/main.go:5)       TEXT    "".main(SB), ABIInternal, $64-0
        0x0000 00000 (/Users/rmliu/workspace/golang/src/github.com/Kate-liu/GoBeginner/helloworld/main/main.go:5)       CMPQ    SP, 16(R14)
        0x0004 00004 (/Users/rmliu/workspace/golang/src/github.com/Kate-liu/GoBeginner/helloworld/main/main.go:5)       PCDATA  $0, $-2
        0x0004 00004 (/Users/rmliu/workspace/golang/src/github.com/Kate-liu/GoBeginner/helloworld/main/main.go:5)       JLS     92
        0x0006 00006 (/Users/rmliu/workspace/golang/src/github.com/Kate-liu/GoBeginner/helloworld/main/main.go:5)       PCDATA  $0, $-1
        0x0006 00006 (/Users/rmliu/workspace/golang/src/github.com/Kate-liu/GoBeginner/helloworld/main/main.go:5)       SUBQ    $64, SP
        0x000a 00010 (/Users/rmliu/workspace/golang/src/github.com/Kate-liu/GoBeginner/helloworld/main/main.go:5)       MOVQ    BP, 56(SP)
        0x000f 00015 (/Users/rmliu/workspace/golang/src/github.com/Kate-liu/GoBeginner/helloworld/main/main.go:5)       LEAQ    56(SP), BP
        0x0014 00020 (/Users/rmliu/workspace/golang/src/github.com/Kate-liu/GoBeginner/helloworld/main/main.go:5)       FUNCDATA        $0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
        0x0014 00020 (/Users/rmliu/workspace/golang/src/github.com/Kate-liu/GoBeginner/helloworld/main/main.go:5)       FUNCDATA        $1, gclocals·f207267fbf96a0178e8758c6e3e0ce28(SB)
        0x0014 00020 (/Users/rmliu/workspace/golang/src/github.com/Kate-liu/GoBeginner/helloworld/main/main.go:5)       FUNCDATA        $2, "".main.stkobj(SB)
        0x0014 00020 (/Users/rmliu/workspace/golang/src/github.com/Kate-liu/GoBeginner/helloworld/main/main.go:8)       MOVUPS  X15, ""..autotmp_9+40(SP)
        0x001a 00026 (/Users/rmliu/workspace/golang/src/github.com/Kate-liu/GoBeginner/helloworld/main/main.go:8)       LEAQ    type.string(SB), DX
        0x0021 00033 (/Users/rmliu/workspace/golang/src/github.com/Kate-liu/GoBeginner/helloworld/main/main.go:8)       MOVQ    DX, ""..autotmp_9+40(SP)
        0x0026 00038 (/Users/rmliu/workspace/golang/src/github.com/Kate-liu/GoBeginner/helloworld/main/main.go:8)       LEAQ    ""..stmp_0(SB), DX
        0x002d 00045 (/Users/rmliu/workspace/golang/src/github.com/Kate-liu/GoBeginner/helloworld/main/main.go:8)       MOVQ    DX, ""..autotmp_9+48(SP)
        0x0032 00050 (<unknown line number>)    NOP
        0x0032 00050 ($GOROOT/src/fmt/print.go:274)     MOVQ    os.Stdout(SB), BX
        0x0039 00057 ($GOROOT/src/fmt/print.go:274)     LEAQ    go.itab.*os.File,io.Writer(SB), AX
        0x0040 00064 ($GOROOT/src/fmt/print.go:274)     LEAQ    ""..autotmp_9+40(SP), CX
        0x0045 00069 ($GOROOT/src/fmt/print.go:274)     MOVL    $1, DI
        0x004a 00074 ($GOROOT/src/fmt/print.go:274)     MOVQ    DI, SI
        0x004d 00077 ($GOROOT/src/fmt/print.go:274)     PCDATA  $1, $0
        0x004d 00077 ($GOROOT/src/fmt/print.go:274)     CALL    fmt.Fprintln(SB)
        0x0052 00082 (/Users/rmliu/workspace/golang/src/github.com/Kate-liu/GoBeginner/helloworld/main/main.go:9)       MOVQ    56(SP), BP
        0x0057 00087 (/Users/rmliu/workspace/golang/src/github.com/Kate-liu/GoBeginner/helloworld/main/main.go:9)       ADDQ    $64, SP
        0x005b 00091 (/Users/rmliu/workspace/golang/src/github.com/Kate-liu/GoBeginner/helloworld/main/main.go:9)       RET
        0x005c 00092 (/Users/rmliu/workspace/golang/src/github.com/Kate-liu/GoBeginner/helloworld/main/main.go:9)       NOP
        0x005c 00092 (/Users/rmliu/workspace/golang/src/github.com/Kate-liu/GoBeginner/helloworld/main/main.go:5)       PCDATA  $1, $-1
        0x005c 00092 (/Users/rmliu/workspace/golang/src/github.com/Kate-liu/GoBeginner/helloworld/main/main.go:5)       PCDATA  $0, $-2
        0x005c 00092 (/Users/rmliu/workspace/golang/src/github.com/Kate-liu/GoBeginner/helloworld/main/main.go:5)       NOP
        0x0060 00096 (/Users/rmliu/workspace/golang/src/github.com/Kate-liu/GoBeginner/helloworld/main/main.go:5)       CALL    runtime.morestack_noctxt(SB)
        0x0065 00101 (/Users/rmliu/workspace/golang/src/github.com/Kate-liu/GoBeginner/helloworld/main/main.go:5)       PCDATA  $0, $-1
        0x0065 00101 (/Users/rmliu/workspace/golang/src/github.com/Kate-liu/GoBeginner/helloworld/main/main.go:5)       JMP     0
        0x0000 49 3b 66 10 76 56 48 83 ec 40 48 89 6c 24 38 48  I;f.vVH..@H.l$8H
        0x0010 8d 6c 24 38 44 0f 11 7c 24 28 48 8d 15 00 00 00  .l$8D..|$(H.....
        0x0020 00 48 89 54 24 28 48 8d 15 00 00 00 00 48 89 54  .H.T$(H......H.T
        0x0030 24 30 48 8b 1d 00 00 00 00 48 8d 05 00 00 00 00  $0H......H......
        0x0040 48 8d 4c 24 28 bf 01 00 00 00 48 89 fe e8 00 00  H.L$(.....H.....
        0x0050 00 00 48 8b 6c 24 38 48 83 c4 40 c3 0f 1f 40 00  ..H.l$8H..@...@.
        0x0060 e8 00 00 00 00 eb 99                             .......
        rel 2+0 t=24 type.string+0
        rel 2+0 t=24 type.*os.File+0
        rel 29+4 t=15 type.string+0
        rel 41+4 t=15 ""..stmp_0+0
        rel 53+4 t=15 os.Stdout+0
        rel 60+4 t=15 go.itab.*os.File,io.Writer+0
        rel 78+4 t=7 fmt.Fprintln+0
        rel 97+4 t=7 runtime.morestack_noctxt+0
os.(*File).close STEXT dupok size=86 args=0x8 locals=0x10 funcid=0x16
        0x0000 00000 (<autogenerated>:1)        TEXT    os.(*File).close(SB), DUPOK|WRAPPER|ABIInternal, $16-8
        0x0000 00000 (<autogenerated>:1)        CMPQ    SP, 16(R14)
        0x0004 00004 (<autogenerated>:1)        PCDATA  $0, $-2
        0x0004 00004 (<autogenerated>:1)        JLS     52
        0x0006 00006 (<autogenerated>:1)        PCDATA  $0, $-1
        0x0006 00006 (<autogenerated>:1)        SUBQ    $16, SP
        0x000a 00010 (<autogenerated>:1)        MOVQ    BP, 8(SP)
        0x000f 00015 (<autogenerated>:1)        LEAQ    8(SP), BP
        0x0014 00020 (<autogenerated>:1)        MOVQ    32(R14), R12
        0x0018 00024 (<autogenerated>:1)        TESTQ   R12, R12
        0x001b 00027 (<autogenerated>:1)        JNE     69
        0x001d 00029 (<autogenerated>:1)        NOP
        0x001d 00029 (<autogenerated>:1)        FUNCDATA        $0, gclocals·1a65e721a2ccc325b382662e7ffee780(SB)
        0x001d 00029 (<autogenerated>:1)        FUNCDATA        $1, gclocals·69c1753bd5f81501d95132d08af04464(SB)
        0x001d 00029 (<autogenerated>:1)        FUNCDATA        $5, os.(*File).close.arginfo1(SB)
        0x001d 00029 (<autogenerated>:1)        MOVQ    AX, ""..this+24(SP)
        0x0022 00034 (<autogenerated>:1)        MOVQ    (AX), AX
        0x0025 00037 (<autogenerated>:1)        PCDATA  $1, $1
        0x0025 00037 (<autogenerated>:1)        CALL    os.(*file).close(SB)
        0x002a 00042 (<autogenerated>:1)        MOVQ    8(SP), BP
        0x002f 00047 (<autogenerated>:1)        ADDQ    $16, SP
        0x0033 00051 (<autogenerated>:1)        RET
        0x0034 00052 (<autogenerated>:1)        NOP
        0x0034 00052 (<autogenerated>:1)        PCDATA  $1, $-1
        0x0034 00052 (<autogenerated>:1)        PCDATA  $0, $-2
        0x0034 00052 (<autogenerated>:1)        MOVQ    AX, 8(SP)
        0x0039 00057 (<autogenerated>:1)        CALL    runtime.morestack_noctxt(SB)
        0x003e 00062 (<autogenerated>:1)        MOVQ    8(SP), AX
        0x0043 00067 (<autogenerated>:1)        PCDATA  $0, $-1
        0x0043 00067 (<autogenerated>:1)        JMP     0
        0x0045 00069 (<autogenerated>:1)        LEAQ    24(SP), R13
        0x004a 00074 (<autogenerated>:1)        CMPQ    (R12), R13
        0x004e 00078 (<autogenerated>:1)        JNE     29
        0x0050 00080 (<autogenerated>:1)        MOVQ    SP, (R12)
        0x0054 00084 (<autogenerated>:1)        JMP     29
        0x0000 49 3b 66 10 76 2e 48 83 ec 10 48 89 6c 24 08 48  I;f.v.H...H.l$.H
        0x0010 8d 6c 24 08 4d 8b 66 20 4d 85 e4 75 28 48 89 44  .l$.M.f M..u(H.D
        0x0020 24 18 48 8b 00 e8 00 00 00 00 48 8b 6c 24 08 48  $.H.......H.l$.H
        0x0030 83 c4 10 c3 48 89 44 24 08 e8 00 00 00 00 48 8b  ....H.D$......H.
        0x0040 44 24 08 eb bb 4c 8d 6c 24 18 4d 39 2c 24 75 cd  D$...L.l$.M9,$u.
        0x0050 49 89 24 24 eb c7                                I.$$..
        rel 38+4 t=7 os.(*file).close+0
        rel 58+4 t=7 runtime.morestack_noctxt+0
go.cuinfo.producer.main SDWARFCUINFO dupok size=0
        0x0000 72 65 67 61 62 69                                regabi
go.cuinfo.packagename.main SDWARFCUINFO dupok size=0
        0x0000 6d 61 69 6e                                      main
go.string..gostring.104.a8ff25b8257e3179242ff5e82c37287d79fc396259b6de98ac48845e9bc4b19f SRODATA dupok size=104
        0x0000 30 77 af 0c 92 74 08 02 41 e1 c1 07 e6 d6 18 e6  0w...t..A.......
        0x0010 70 61 74 68 09 63 6f 6d 6d 61 6e 64 2d 6c 69 6e  path.command-lin
        0x0020 65 2d 61 72 67 75 6d 65 6e 74 73 0a 6d 6f 64 09  e-arguments.mod.
        0x0030 67 69 74 68 75 62 2e 63 6f 6d 2f 4b 61 74 65 2d  github.com/Kate-
        0x0040 6c 69 75 2f 47 6f 42 65 67 69 6e 6e 65 72 09 28  liu/GoBeginner.(
        0x0050 64 65 76 65 6c 29 09 0a f9 32 43 31 86 18 20 72  devel)...2C1.. r
        0x0060 00 82 42 10 41 16 d8 f2                          ..B.A...
""..inittask SNOPTRDATA size=32
        0x0000 00 00 00 00 00 00 00 00 01 00 00 00 00 00 00 00  ................
        0x0010 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        rel 24+8 t=1 fmt..inittask+0
go.info.fmt.Println$abstract SDWARFABSFCN dupok size=42
        0x0000 04 66 6d 74 2e 50 72 69 6e 74 6c 6e 00 01 01 11  .fmt.Println....
        0x0010 61 00 00 00 00 00 00 11 6e 00 01 00 00 00 00 11  a.......n.......
        0x0020 65 72 72 00 01 00 00 00 00 00                    err.......
        rel 0+0 t=23 type.[]interface {}+0
        rel 0+0 t=23 type.error+0
        rel 0+0 t=23 type.int+0
        rel 19+4 t=31 go.info.[]interface {}+0
        rel 27+4 t=31 go.info.int+0
        rel 37+4 t=31 go.info.error+0
go.string."你好，世界" SRODATA dupok size=15
        0x0000 e4 bd a0 e5 a5 bd ef bc 8c e4 b8 96 e7 95 8c     ...............
runtime.modinfo SDATA size=16
        0x0000 00 00 00 00 00 00 00 00 68 00 00 00 00 00 00 00  ........h.......
        rel 0+8 t=1 go.string..gostring.104.a8ff25b8257e3179242ff5e82c37287d79fc396259b6de98ac48845e9bc4b19f+0
go.info.runtime.modinfo SDWARFVAR dupok size=32
        0x0000 08 72 75 6e 74 69 6d 65 2e 6d 6f 64 69 6e 66 6f  .runtime.modinfo
        0x0010 00 09 03 00 00 00 00 00 00 00 00 00 00 00 00 01  ................
        rel 19+8 t=1 runtime.modinfo+0
        rel 27+4 t=31 go.info.string+0
""..stmp_0 SRODATA static size=16
        0x0000 00 00 00 00 00 00 00 00 0f 00 00 00 00 00 00 00  ................
        rel 0+8 t=1 go.string."你好，世界"+0
runtime.nilinterequal·f SRODATA dupok size=8
        0x0000 00 00 00 00 00 00 00 00                          ........
        rel 0+8 t=1 runtime.nilinterequal+0
runtime.memequal64·f SRODATA dupok size=8
        0x0000 00 00 00 00 00 00 00 00                          ........
        rel 0+8 t=1 runtime.memequal64+0
runtime.gcbits.01 SRODATA dupok size=1
        0x0000 01                                               .
type..namedata.*interface {}- SRODATA dupok size=15
        0x0000 00 0d 2a 69 6e 74 65 72 66 61 63 65 20 7b 7d     ..*interface {}
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
        rel 44+4 t=-32763 type.*interface {}+0
        rel 56+8 t=1 type.interface {}+80
type..namedata.*[]interface {}- SRODATA dupok size=17
        0x0000 00 0f 2a 5b 5d 69 6e 74 65 72 66 61 63 65 20 7b  ..*[]interface {
        0x0010 7d                                               }
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
        rel 44+4 t=-32763 type.*[]interface {}+0
        rel 48+8 t=1 type.interface {}+0
type..namedata.*[1]interface {}- SRODATA dupok size=18
        0x0000 00 10 2a 5b 31 5d 69 6e 74 65 72 66 61 63 65 20  ..*[1]interface 
        0x0010 7b 7d                                            {}
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
        rel 44+4 t=-32763 type.*[1]interface {}+0
        rel 48+8 t=1 type.interface {}+0
        rel 56+8 t=1 type.[]interface {}+0
go.itab.*os.File,io.Writer SRODATA dupok size=32
        0x0000 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0010 44 b5 f3 33 00 00 00 00 00 00 00 00 00 00 00 00  D..3............
        rel 0+8 t=1 type.io.Writer+0
        rel 8+8 t=1 type.*os.File+0
        rel 24+8 t=-32767 os.(*File).Write+0
type..importpath.fmt. SRODATA dupok size=5
        0x0000 00 03 66 6d 74                                   ..fmt
type..importpath.unsafe. SRODATA dupok size=8
        0x0000 00 06 75 6e 73 61 66 65                          ..unsafe
gclocals·33cdeccccebe80329f1fdbee7f5874cb SRODATA dupok size=8
        0x0000 01 00 00 00 00 00 00 00                          ........
gclocals·f207267fbf96a0178e8758c6e3e0ce28 SRODATA dupok size=9
        0x0000 01 00 00 00 02 00 00 00 00                       .........
"".main.stkobj SRODATA static size=32
        0x0000 01 00 00 00 00 00 00 00 f0 ff ff ff 10 00 00 00  ................
        0x0010 10 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        rel 24+8 t=1 runtime.gcbits.02+0
gclocals·1a65e721a2ccc325b382662e7ffee780 SRODATA dupok size=10
        0x0000 02 00 00 00 01 00 00 00 01 00                    ..........
gclocals·69c1753bd5f81501d95132d08af04464 SRODATA dupok size=8
        0x0000 02 00 00 00 00 00 00 00                          ........
os.(*File).close.arginfo1 SRODATA static dupok size=3
        0x0000 00 08 ff                                         ...
```



### 获取汇编指令的优化过程

如果想要了解 Go 语言更详细的编译过程，可以通过下面的命令获取汇编指令的优化过程。

```sh
$GOSSAFUNC=main go1.17 build main.go
# runtime
dumped SSA to /Users/rmliu/workspace/golang/src/github.com/Kate-liu/GoBeginner/helloworld/main/ssa.html
# command-line-arguments
dumped SSA to ./ssa.html
```

上述命令会在当前文件夹下生成一个 `ssa.html` 文件，打开这个文件后就能看到汇编代码优化的每一个步骤：

![image-20220103112604900](tools.assets/image-20220103112604900.png)

上述 HTML 文件是**可以交互**的，当点击网页上的汇编指令时，页面会使用相同的颜色在 SSA 中间代码生成的不同阶段标识出相关的代码行，更方便开发者分析编译优化的过程。



### Go 代码转换为汇编代码

使用下面的命令，可以在当前文件生成汇编代码。

其中main.s文件中是原go程序的汇编，可以根据(main.go:24)的行号，来对照原文件查找转换后的汇编代码。

```sh
$go tool compile -S main.go > main.s
 
$ls -al
total 120
drwxr-xr-x   5 rmliu  staff    160 Jan  4 19:51 .
drwxr-xr-x  15 rmliu  staff    480 Jan  4 19:51 ..
-rw-r--r--   1 rmliu  staff    308 Jan  4 19:50 main.go
-rw-r--r--   1 rmliu  staff  17826 Jan  4 19:51 main.o
-rw-r--r--   1 rmliu  staff  36214 Jan  4 19:51 main.s
```







## mac 命令行工具

### 获得当前机器的硬件信息

在命令行中输入 `uname -m` 就能获得当前机器的硬件信息：

```sh
$uname -m
x86_64
```

x86 是目前比较常见的指令集，除了 x86 之外，还有 arm 等指令集，苹果最新 Macbook 的自研芯片就使用了 arm 指令集，不同的处理器使用了不同的架构和机器语言，所以很多编程语言为了在不同的机器上运行需要将源代码根据架构翻译成不同的机器代码。

复杂指令集计算机（CISC）和精简指令集计算机（RISC）是两种遵循不同设计理念的指令集，从名字就可以推测出这两种指令集的区别：

- 复杂指令集：通过增加指令的类型减少需要执行的指令数；
- 精简指令集：使用更少的指令类型完成目标的计算任务；

早期的 CPU 为了减少机器语言指令的数量一般使用复杂指令集完成计算任务，这两者并没有绝对的优劣，它们只是在一些设计上的选择不同以达到不同的目的

Go 语言源代码的 [`src/cmd/compile/internal`](https://github.com/golang/go/tree/master/src/cmd/compile/internal) 目录中包含了很多**机器码生成相关的包**，不同类型的 CPU 分别使用了不同的包生成机器码，其中包括 amd64、arm、arm64、mips、mips64、ppc64、s390x、x86 和 wasm，其中比较有趣的就是 [WebAssembly](https://webassembly.org/)（Wasm）了。

作为一种在栈虚拟机上使用的二进制指令格式，它的设计的主要目标就是在 Web 浏览器上提供一种具有高可移植性的目标语言。Go 语言的编译器既然能够生成 Wasm 格式的指令，那么就能够运行在常见的主流浏览器中。

```bash
$ GOARCH=wasm GOOS=js go1.17 build -o lib.wasm main.go
```

可以使用上述的命令将 Go 的源代码编译成能够在浏览器上运行 WebAssembly 文件，当然除了这种新兴的二进制指令格式之外，Go 语言经过编译还可以运行在几乎全部的主流机器上，不过它的兼容性在除 Linux 和 Darwin 之外的机器上可能还有一些问题，例如：Go Plugin 至今仍然不支持 [Windows](https://github.com/golang/go/issues/19282)。





## lex 实现词法分析器

[lex](http://dinosaur.compilertools.net/lex/index.html)是用于生成词法分析器的工具，lex 生成的代码能够将一个文件中的字符分解成 Token 序列，很多语言在设计早期都会使用它快速设计出原型。词法分析作为具有固定模式的任务，出现这种更抽象的工具必然的，lex 作为一个代码生成器，使用了类似 C 语言的语法，将 lex 理解为正则匹配的生成器，它会使用正则匹配扫描输入的字符流，下面是一个 lex 文件的示例：

```c
// simplego.l 文件
%{
#include <stdio.h>
%}

%%
package      printf("PACKAGE ");
import       printf("IMPORT ");
\.           printf("DOT ");
\{           printf("LBRACE ");
\}           printf("RBRACE ");
\(           printf("LPAREN ");
\)           printf("RPAREN ");
\"           printf("QUOTE ");
\n           printf("\n");
[0-9]+       printf("NUMBER ");
[a-zA-Z_]+   printf("IDENT ");
%%
```

这个定义好的文件能够解析 `package` 和 `import` 关键字、常见的特殊字符、数字以及标识符，虽然这里的规则可能有一些简陋和不完善，但是用来解析下面的这一段代码还是比较轻松的：

```go
// main.go 文件
package main

import (
  "fmt"
)

func main() {
  fmt.Println("Hello")
}
```

`.l` 结尾的 lex 代码并不能直接运行，首先需要通过 `lex` 命令将上面的 `simplego.l` 展开成 C 语言代码，这里可以直接执行如下所示的命令编译并打印文件中的内容：

```c
$ lex simplego.l  // 会生成lex.yy.c的新文件
$ cat lex.yy.c
...
int yylex (void) {
  ...
  while ( 1 ) {
    ...
yy_match:
    do {
      register YY_CHAR yy_c = yy_ec[YY_SC_TO_UI(*yy_cp)];
      if ( yy_accept[yy_current_state] ) {
        (yy_last_accepting_state) = yy_current_state;
        (yy_last_accepting_cpos) = yy_cp;
      }
      while ( yy_chk[yy_base[yy_current_state] + yy_c] != yy_current_state ) {
        yy_current_state = (int) yy_def[yy_current_state];
        if ( yy_current_state >= 30 )
          yy_c = yy_meta[(unsigned int) yy_c];
        }
      yy_current_state = yy_nxt[yy_base[yy_current_state] + (unsigned int) yy_c];
      ++yy_cp;
    } while ( yy_base[yy_current_state] != 37 );
    ...

do_action:
    switch ( yy_act )
      case 0:
          ...

      case 1:
          YY_RULE_SETUP
          printf("PACKAGE ");
          YY_BREAK
      ...
}
```

`lex.yy.c` 的前 600 行基本都是宏和函数的声明和定义，后面生成的代码大都是为 `yylex` 这个函数服务的，这个函数使用[有限自动机（Deterministic Finite Automaton、DFA）](https://en.wikipedia.org/wiki/Deterministic_finite_automaton)的程序结构来分析输入的字符流，上述代码中 `while` 循环就是这个有限自动机的主体，如果仔细看这个文件生成的代码会发现当前的文件中并不存在 `main` 函数，`main` 函数是在 liblex 库中定义的，所以在编译时其实需要添加额外的 `-ll` 选项：

```sh
$ cc lex.yy.c -o simplego -ll  # 此时会生成一个 simplego 的二进制文件
$ cat main.go | ./simplego
```

将 C 语言代码通过 gcc 编译成二进制代码之后，就可以使用管道将上面提到的 Go 语言代码作为输入传递到生成的词法分析器中，这个词法分析器会打印出如下的内容：

```go
PACKAGE  IDENT

IMPORT  LPAREN
  QUOTE IDENT QUOTE
RPAREN

IDENT  IDENT LPAREN RPAREN  LBRACE
  IDENT DOT IDENT LPAREN QUOTE IDENT QUOTE RPAREN
RBRACE
```

从上面的输出能够看到 Go 源代码的影子，lex 生成的词法分析器 lexer 通过正则匹配的方式将机器原本很难理解的字符串进行分解成很多的 Token，有利于后面的处理。

到这里已经展示了从定义 `.l` 文件、使用 lex 将 `.l` 文件编译成 C 语言代码以及二进制的全过程，而最后生成的词法分析器也能够将简单的 Go 语言代码进行转换成 Token 序列。















