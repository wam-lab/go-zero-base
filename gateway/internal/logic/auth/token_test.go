package logic

import (
	"fmt"
	"testing"
)

func TestAA(t *testing.T) {
	fmt.Println("crashinterceptor.go:40 reflect: reflect.flag.mustBeAssignable using unaddressable value goroutine 44 [running]:\nru\nntime/debug.Stack(0x40, 0xe, 0x0)\n\tD:/Develop/Go/src/runtime/debug/stack.go:24 +0xa4\ngithub.com/tal-tech/go-zero/zrpc/internal/serverinterceptors.toPanicError(0xcd7f60, 0xc00031c340,\n 0x0, 0xc0003c8400)\n\tD:/GoWorkspace/pkg/mod/github.com/tal-tech/go-zero@v1.1.1/zrpc/internal/serverinterceptors/crashinterceptor.go:40 +0x2d\ngithub.com/tal-tech/go-zero/zrpc/internal\n/serverinterceptors.UnaryStatInterceptor.func1.1(0xcd7f60, 0xc00031c340)\n\tD:/GoWorkspace/pkg/mod/github.com/tal-tech/go-zero@v1.1.1/zrpc/internal/serverinterceptors/statinterceptor.go\n:21 +0x45\ngithub.com/tal-tech/go-zero/zrpc/internal/serverinterceptors.handleCrash(0xc000347860)\n\tD:/GoWorkspace/pkg/mod/github.com/tal-tech/go-zero@v1.1.1/zrpc/internal/serverinterc\neptors/crashinterceptor.go:35 +0x5e\npanic(0xcd7f60, 0xc00031c340)\n\tD:/Develop/Go/src/runtime/panic.go:969 +0x174\nreflect.flag.mustBeAssignableSlow(0x87)\n\tD:/Develop/Go/src/reflect\n/value.go:247 +0x13f\nreflect.flag.mustBeAssignable(...)\n\tD:/Develop/Go/src/reflect/value.go:234\nreflect.Value.SetUint(0xcd80a0, 0xc0003b60b0, 0x87, 0x2)\n\tD:/Develop/Go/src/reflect\n/value.go:1687 +0x42\ngorm.io/gorm/schema.(*Field).setupValuerAndSetter.func11(0xde7e40, 0xc0003b60b0, 0x99, 0xcd80a0, 0xc000402890, 0xc000402890, 0xc00008c128)\n\tD:/GoWorkspace/pkg/mo\nd/gorm.io/gorm@v1.20.9/schema/field.go:608 +0xac5\ngorm.io/gorm/schema.(*Field).setupValuerAndSetter.func8(0xde7e40, 0xc0003b60b0, 0x99, 0xcbeec0, 0xc000402858, 0xc0004b9060, 0xc0001101\ne8, 0x0)\n\tD:/GoWorkspace/pkg/mod/gorm.io/gorm@v1.20.9/schema/field.go:501 +0x25f\ngorm.io/gorm/schema.(*Field).setupValuerAndSetter.func11(0xde7e40, 0xc0003b60b0, 0x99, 0xcbeec0, 0xc0\n00402858, 0xc000402858, 0xc000097560)\n\tD:/GoWorkspace/pkg/mod/gorm.io/gorm@v1.20.9/schema/field.go:646 +0xe0\ngorm.io/gorm/schema.(*Field).setupValuerAndSetter.func8(0xde7e40, 0xc0003\nb60b0, 0x99, 0xc00031e3c0, 0xc00008c0c8, 0xc0004b9060, 0xc0003c8330, 0x413781)\n\tD:/GoWorkspace/pkg/mod/gorm.io/gorm@v1.20.9/schema/field.go:501 +0x25f\ngorm.io/gorm/schema.(*Field).se\ntupValuerAndSetter.func11(0xde7e40, 0xc0003b60b0, 0x99, 0xc00031e3c0, 0xc00008c0c8, 0x1, 0x198ffff)\n\tD:/GoWorkspace/pkg/mod/gorm.io/gorm@v1.20.9/schema/field.go:646 +0xe0\ngorm.io/gor\nm.Scan(0xc0003c8300, 0xc000097080, 0xc00002a000)\n\tD:/GoWorkspace/pkg/mod/gorm.io/gorm@v1.20.9/scan.go:221 +0x1953\ngorm.io/gorm/callbacks.Query(0xc000097080)\n\tD:/GoWorkspace/pkg/mod\n/gorm.io/gorm@v1.20.9/callbacks/query.go:26 +0x145\ngorm.io/gorm.(*processor).Execute(0xc0000b1740, 0xc000097080)\n\tD:/GoWorkspace/pkg/mod/gorm.io/gorm@v1.20.9/callbacks.go:105 +0x22b\\nngorm.io/gorm.(*DB).First(0xc000097080, 0xde7e40, 0xc0003b60b0, 0x0, 0x0, 0x0, 0xc000097080)\n\tD:/GoWorkspace/pkg/mod/gorm.io/gorm@v1.20.9/finisher_api.go:115 +0x174\ngithub.com/yguila\ni/timetable-micro/services/user/rpc/internal/logic.(*LoginLogic).Login(0xc000347498, 0xc00008a9b0, 0xc0004092c0, 0x0, 0x0)\n\tD:/GoModules/github.com/yguilai/timetable-micro/services/us\ner/rpc/internal/logic/login_logic.go:64 +0x42d\ngithub.com/yguilai/timetable-micro/services/user/rpc/internal/server.(*UserServer).Login(0xc0003901c8, 0xf9f4c0, 0xc0002a95c0, 0xc00008a9\nb0, 0xc0003901c8, 0xc0002a95c0, 0xc00012c1d0)\n\tD:/GoModules/github.com/yguilai/timetable-micro/services/user/rpc/internal/server/user_server.go:36 +0xc0\ngithub.com/yguilai/timetable-\nmicro/services/user/rpc/user._User_Login_Handler.func1(0xf9f4c0, 0xc0002a95c0, 0xdba800, 0xc00008a9b0, 0xc0002a95c0, 0xc00012c1d0, 0x180e58, 0x0)\n\tD:/GoModules/github.com/yguilai/time\ntable-micro/services/user/rpc/user/user.pb.go:1345 +0x8d\ngithub.com/tal-tech/go-zero/zrpc/internal/serverinterceptors.UnaryTimeoutInterceptor.func1(0xf9f500, 0xc0004092c0, 0xdba800, 0x\nc00008a9b0, 0xc000448400, 0xc000448440, 0x0, 0x0, 0x0, 0x0)\n\tD:/GoWorkspace/pkg/mod/github.com/tal-tech/go-zero@v1.1.1/zrpc/internal/serverinterceptors/timeoutinterceptor.go:16 +0xb0\\nngoogle.golang.org/grpc.getChainUnaryHandler.func1(0xf9f500, 0xc0004092c0, 0xdba800, 0xc00008a9b0, 0x0, 0x40d36f, 0x40, 0xdc12c0)\n\tD:/GoWorkspace/pkg/mod/google.golang.org/grpc@v1.29.\n1/server.go:921 +0xee\ngithub.com/tal-tech/go-zero/zrpc/internal/serverinterceptors.UnarySheddingInterceptor.func1(0xf9f500, 0xc0004092c0, 0xdba800, 0xc00008a9b0, 0xc000448400, 0xc0000b\n06c0, 0x0, 0x0, 0x0, 0x0)\n\tD:/GoWorkspace/pkg/mod/github.com/tal-tech/go-zero@v1.1.1/zrpc/internal/serverinterceptors/sheddinginterceptor.go:42 +0x135\ngoogle.golang.org/grpc.getChain\nUnaryHandler.func1(0xf9f500, 0xc0004092c0, 0xdba800, 0xc00008a9b0, 0x30000, 0xc000111790, 0xc000111758, 0x40d36f)\n\tD:/GoWorkspace/pkg/mod/google.golang.org/grpc@v1.29.1/server.go:921\n+0xee\ngithub.com/tal-tech/go-zero/zrpc/internal/serverinterceptors.UnaryPrometheusInterceptor.func1(0xf9f500, 0xc0004092c0, 0xdba800, 0xc00008a9b0, 0xc000448400, 0xc0000b0680, 0xc0000b\n0680, 0xbff3376625357148, 0x1a2849381, 0x162e380)\n\tD:/GoWorkspace/pkg/mod/github.com/tal-tech/go-zero@v1.1.1/zrpc/internal/serverinterceptors/prometheusinterceptor.go:39 +0xa8\ngoogle\n.golang.org/grpc.getChainUnaryHandler.func1(0xf9f500, 0xc0004092c0, 0xdba800, 0xc00008a9b0, 0x8, 0xc000080000, 0x180e58, 0x0)\n\tD:/GoWorkspace/pkg/mod/google.golang.org/grpc@v1.29.1/se\nrver.go:921 +0xee\ngithub.com/tal-tech/go-zero/zrpc/internal/serverinterceptors.UnaryStatInterceptor.func1(0xf9f500, 0xc0004092c0, 0xdba800, 0xc00008a9b0, 0xc000448400, 0xc0000b0640, 0x\n0, 0x0, 0x0, 0x0)\n\tD:/GoWorkspace/pkg/mod/github.com/tal-tech/go-zero@v1.1.1/zrpc/internal/serverinterceptors/statinterceptor.go:33 +0x170\ngoogle.golang.org/grpc.getChainUnaryHandler\n.func1(0xf9f500, 0xc0004092c0, 0xdba800, 0xc00008a9b0, 0x40d36f, 0x40, 0xdc12c0, 0x1)\n\tD:/GoWorkspace/pkg/mod/google.golang.org/grpc@v1.29.1/server.go:921 +0xee\ngithub.com/tal-tech/g\no-zero/zrpc/internal/serverinterceptors.UnaryCrashInterceptor.func1(0xf9f500, 0xc0004092c0, 0xdba800, 0xc00008a9b0, 0xc000448400, 0xc0000b0600, 0x0, 0x0, 0x0, 0x0)\n\tD:/GoWorkspace/pkg\n/mod/github.com/tal-tech/go-zero@v1.1.1/zrpc/internal/serverinterceptors/crashinterceptor.go:29 +0xbe\ngoogle.golang.org/grpc.getChainUnaryHandler.func1(0xf9f500, 0xc0004092c0, 0xdba800\n, 0xc00008a9b0, 0xc000403f80, 0x8, 0xe53649, 0x10)\n\tD:/GoWorkspace/pkg/mod/google.golang.org/grpc@v1.29.1/server.go:921 +0xee\ngithub.com/tal-tech/go-zero/zrpc/internal/serverintercep\ntors.UnaryTracingInterceptor.func1(0xf9f500, 0xc000408a50, 0xdba800, 0xc00008a9b0, 0xc000448400, 0xc0000b05c0, 0x0, 0x0, 0x0, 0x0)\n\tD:/GoWorkspace/pkg/mod/github.com/tal-tech/go-zero@\nv1.1.1/zrpc/internal/serverinterceptors/tracinginterceptor.go:26 +0x244\ngoogle.golang.org/grpc.chainUnaryServerInterceptors.func1(0xf9f500, 0xc000408a50, 0xdba800, 0xc00008a9b0, 0xc000\n448400, 0xc000448440, 0xc000021ba0, 0x50d07f, 0xdc8fa0, 0xc000408a50)\n\tD:/GoWorkspace/pkg/mod/google.golang.org/grpc@v1.29.1/server.go:907 +0xd1\ngithub.com/yguilai/timetable-micro/se\nrvices/user/rpc/user._User_Login_Handler(0xd8a260, 0xc0003901c8, 0xf9f500, 0xc000408a50, 0xc0002a9380, 0xc0000887e0, 0xf9f500, 0xc000408a50, 0xc00013e260, 0x1f)\n\tD:/GoModules/github.c\nom/yguilai/timetable-micro/services/user/rpc/user/user.pb.go:1347 +0x152\ngoogle.golang.org/grpc.(*Server).processUnaryRPC(0xc00006c000, 0xfaa140, 0xc000524900, 0xc000138900, 0xc0004407\n80, 0x161f490, 0x0, 0x0, 0x0)\n\tD:/GoWorkspace/pkg/mod/google.golang.org/grpc@v1.29.1/server.go:1082 +0x511\ngoogle.golang.org/grpc.(*Server).handleStream(0xc00006c000, 0xfaa140, 0xc00\n0524900, 0xc000138900, 0x0)\n\tD:/GoWorkspace/pkg/mod/google.golang.org/grpc@v1.29.1/server.go:1405 +0xcd4\ngoogle.golang.org/grpc.(*Server).serveStreams.func1.1(0xc0004024f0, 0xc00006c\n000, 0xfaa140, 0xc000524900, 0xc000138900)\n\tD:/GoWorkspace/pkg/mod/google.golang.org/grpc@v1.29.1/server.go:746 +0xa8\ncreated by google.golang.org/grpc.(*Server).serveStreams.func1\n\n\tD:/GoWorkspace/pkg/mod/google.golang.org/grpc@v1.29.1/server.go:744 +0xa8")
}
