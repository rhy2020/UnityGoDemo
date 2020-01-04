package msg
import (
	"github.com/name5566/leaf/network/protobuf"
)

var (
	Processor = protobuf.NewProcessor()
)

func init() {	// 这里我们注册 protobuf 消息)
    Processor.SetByteOrder(true)
    Processor.Register(&RspHead{})
    Processor.Register(&HeartReq{})
    Processor.Register(&HeartRsp{})
    Processor.Register(&StartFight{})
    Processor.Register(&FightResult{})
    Processor.Register(&EnterFight{})
    Processor.Register(&SignUpResponse{})
    Processor.Register(&TosChat{})
    Processor.Register(&TocChat{})
    Processor.Register(&Login{})
    Processor.Register(&PlayerBaseInfo{})
    Processor.Register(&LoginSuccessfull{})
    Processor.Register(&LoginFaild{})
    Processor.Register(&RoomData{})
    Processor.Register(&RoomListReq{})
    Processor.Register(&RoomListRsp{})
    Processor.Register(&JoinRoomReq{})
    Processor.Register(&JoinRoomRsp{})
    Processor.Register(&LeaveRoomReq{})
    Processor.Register(&LeaveRoomRsp{})
    Processor.Register(&RoomPlayerJoinNtf{})
    Processor.Register(&RoomPlayerLeaveNtf{})
    Processor.Register(&RoomBattleReadyReq{})
    Processor.Register(&RoomBattleReadyRsp{})
    Processor.Register(&RoomBattleReadyNtf{})
    Processor.Register(&RoomBattleCancleReadyReq{})
    Processor.Register(&RoomBattleCancleReadyRsp{})
    Processor.Register(&RoomBattleCancleReadyNtf{})
    Processor.Register(&RoomBattleStartReq{})
    Processor.Register(&RoomBattleStartRsp{})
    Processor.Register(&RoomBattleStartNtf{})

}